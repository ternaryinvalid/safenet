package cryptography

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"math/big"
)

// Encrypt шифрует данные с использованием алгоритма AES в режиме CBC.
func Encrypt(plainText, keyString string) (string, error) {
	key := sha256.Sum256([]byte(keyString))
	keySlice := key[:] // Преобразуем массив в срез

	block, err := aes.NewCipher(keySlice)
	if err != nil {
		return "", err
	}

	// Генерируем случайный IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Добавляем паддинг до кратности длины блока
	plainTextBytes := []byte(plainText)
	plainTextBytes = pad(plainTextBytes)

	// Шифруем данные
	cipherText := make([]byte, aes.BlockSize+len(plainTextBytes))
	copy(cipherText[:aes.BlockSize], iv)
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainTextBytes)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func pad(buf []byte) []byte {
	padding := aes.BlockSize - len(buf)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(buf, padText...)
}

// Decrypt дешифрует данные с использованием алгоритма AES в режиме CBC.
func Decrypt(cipherTextString, keyString string) (string, error) {
	key := sha256.Sum256([]byte(keyString))
	keySlice := key[:] // Преобразуем массив в срез

	cipherText, err := base64.StdEncoding.DecodeString(cipherTextString)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("Invalid ciphertext length: too short")
	}
	if len(cipherText)%aes.BlockSize != 0 {
		return "", errors.New("Invalid ciphertext length: not a multiple of blocksize")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	block, err := aes.NewCipher(keySlice)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText = unpad(cipherText)

	return string(cipherText), nil
}

func unpad(buf []byte) []byte {
	// Предполагаем допустимую длину и дополнение.
	// Нужно добавить проверки
	padding := int(buf[len(buf)-1])
	return buf[:len(buf)-padding]
}

// GetTransportKey генерирует общий транспортный ключ на основе приватного ключа отправителя и публичного ключа получателя.
func GetTransportKey(senderPrivKeyHex, receiverPubKeyHex string) ([]byte, error) {
	// Декодирование приватного ключа отправителя
	senderPrivKeyBytes, err := hex.DecodeString(senderPrivKeyHex)
	if err != nil {
		return nil, err
	}

	// Декодирование публичного ключа получателя
	receiverPubKeyBytes, err := hex.DecodeString(receiverPubKeyHex)
	if err != nil {
		return nil, err
	}

	// Проверка длины приватного ключа
	if len(senderPrivKeyBytes) != 32 {
		return nil, errors.New("invalid length for private key, need 256 bits")
	}

	// Проверка длины публичного ключа
	if len(receiverPubKeyBytes) != 65 {
		return nil, errors.New("invalid length for public key, need 65 bytes")
	}

	// Преобразование строки ключа в структуру приватного ключа
	privateKey := new(ecdsa.PrivateKey)
	privateKey.PublicKey.Curve = elliptic.P256()
	privateKey.D = new(big.Int).SetBytes(senderPrivKeyBytes)
	privateKey.PublicKey.X, privateKey.PublicKey.Y = privateKey.PublicKey.Curve.ScalarBaseMult(senderPrivKeyBytes)

	// Преобразование строки ключа в структуру публичного ключа
	x, y := elliptic.Unmarshal(elliptic.P256(), receiverPubKeyBytes)
	if x == nil {
		return nil, errors.New("invalid public key")
	}
	receiverPubKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	// Выполнение ECDH key exchange
	sharedX, _ := receiverPubKey.ScalarMult(receiverPubKey.X, receiverPubKey.Y, privateKey.D.Bytes())

	// Преобразование общего секрета в срез байтов
	sharedSecret := sharedX.Bytes()

	// Хэширование общего секрета для получения транспортного ключа
	transportKey := sha256.Sum256(sharedSecret)

	return transportKey[:], nil
}

// GenerateKeys создает подходящие приватный и публичный ключи ECDSA.
func GenerateKeys() (privateKeyHex string, publicKeyHex string, err error) {
	// Генерация приватного ключа
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}

	// Приватный ключ должен быть 32 байта
	privateKeyBytes := privateKey.D.Bytes()
	if len(privateKeyBytes) != 32 {
		return "", "", errors.New("invalid length for private key, need 256 bits")
	}
	privateKeyHex = hex.EncodeToString(privateKeyBytes)

	// Публичный ключ должен быть 65 байт
	publicKeyBytes := elliptic.Marshal(privateKey.Curve, privateKey.X, privateKey.Y)
	if len(publicKeyBytes) != 65 {
		return "", "", errors.New("invalid length for public key, need 65 bytes")
	}
	publicKeyHex = hex.EncodeToString(publicKeyBytes)

	return privateKeyHex, publicKeyHex, nil
}

// GetPublicKeyAndAddressByPrivateKey возвращает публичный ключ и адрес на основе приватного ключа.
func GetPublicKeyAndAddressByPrivateKey(privateKeyHex string) (publicKeyHex string, err error) {
	// Удаление префикса "0x", если он существует
	if len(privateKeyHex) > 2 && privateKeyHex[:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}

	// Декодирование приватного ключа
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", err
	}

	// Проверка длины приватного ключа
	if len(privateKeyBytes) != 32 {
		return "", errors.New("invalid length for private key, need 256 bits")
	}

	// Преобразование байтового среза в структуру приватного ключа
	privateKey := new(ecdsa.PrivateKey)
	privateKey.PublicKey.Curve = elliptic.P256()
	privateKey.D = new(big.Int).SetBytes(privateKeyBytes)
	privateKey.PublicKey.X, privateKey.PublicKey.Y = privateKey.PublicKey.Curve.ScalarBaseMult(privateKeyBytes)

	// Генерация публичного ключа в байтовом формате
	publicKeyBytes := elliptic.Marshal(privateKey.Curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)
	if len(publicKeyBytes) != 65 {
		return "", errors.New("invalid length for public key, need 65 bytes")
	}
	publicKeyHex = hex.EncodeToString(publicKeyBytes)

	return publicKeyHex, nil
}
