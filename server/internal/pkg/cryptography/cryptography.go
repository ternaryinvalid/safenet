package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func Encrypt(sharedKey []byte, message []byte) (string, error) {
	block, err := aes.NewCipher(sharedKey)
	if err != nil {
		err = fmt.Errorf("ошибка создания шифра: %v", err)

		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(message))

	iv := ciphertext[:aes.BlockSize]

	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		err = fmt.Errorf("ошибка генерации вектора инициализации: %v", err)

		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], message)

	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(sharedKey []byte, cipheredMessage []byte) (string, error) {
	block, err := aes.NewCipher(sharedKey)
	if err != nil {
		err = fmt.Errorf("ошибка создания шифра: %v", err)

		return "", err
	}

	decodedCiphertext, err := hex.DecodeString(string(cipheredMessage))
	if err != nil {
		err = fmt.Errorf("ошибка декодирования шифртекста: %v", err)

		return "", err
	}

	if len(decodedCiphertext) < aes.BlockSize {
		err = fmt.Errorf("неверный размер шифртекста")

		return "", err
	}

	iv := decodedCiphertext[:aes.BlockSize]
	decodedCiphertext = decodedCiphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decodedCiphertext, decodedCiphertext)

	return string(decodedCiphertext), nil
}

func GetSharedKey(remotePublicKey []byte, localPrivateKey *ecdsa.PrivateKey) []byte {
	x, y := elliptic.P256().ScalarBaseMult(remotePublicKey)

	sharedSecret, _ := elliptic.P256().ScalarMult(x, y, localPrivateKey.D.Bytes())

	return sharedSecret.Bytes()
}

func GenerateKeys() ([]byte, *ecdsa.PrivateKey, error) {
	localPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	return localPrivateKey.PublicKey.X.Bytes(), localPrivateKey, nil
}
