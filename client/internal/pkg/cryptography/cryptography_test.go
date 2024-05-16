package cryptography

import (
	"log"
	"testing"
)

func TestGetSharedKey(t *testing.T) {
	//publicKeyA := "04af99eb56d35ff5e843d9f82b63f435f676430bc0665efae04a09a502534cb2eeaded9aa4199d3d3839e69f5efef9ed09b963acdf7951e19b9f3eeffc893d0dba"
	//privateKeyA := "41a7f601c049055b3a9da94422b1db2b243bbba0b4a545f15d6a6b671cdce736"
	//
	//publicKeyB := "044f18c9a3dcdb13877d46bf02d4dad6e6da1f9893cd2df759c22609a6ff744b0f2c16c1b48ac0e374237626b3a806f4605d64c379638d5fbf2db8f70198f3ebb9"
	//privateKeyB := "ff791d94aef0efb5f0466b21e8391f68157f798c4555282a996eef957482965a"

	sharedKey := "hhhh_asdq0x"

	text := "hello world...a dd"

	encrypted, err := Encrypt(text, sharedKey)
	if err != nil {
		t.Errorf("encrypt err %v", err)
		t.Fail()
	}

	decrypted, err := Decrypt(encrypted, sharedKey)
	if err != nil {
		t.Errorf("decrypt err %v", err)
		t.Fail()
	}

	if string(text) != decrypted {
		t.Errorf("неправильная дешифровка")
	}
}

func TestGetTransportKey(t *testing.T) {
	privA, pubA, err := GenerateKeys()
	if err != nil {
		t.Errorf("[GenerateKeys] ERR %v", err)
		t.Fail()
	}

	log.Println(len(pubA), pubA)
	log.Println(len(privA), privA)

	privB, pubB, err := GenerateKeys()
	if err != nil {
		t.Errorf("[GenerateKeys] ERR %v", err)
		t.Fail()
	}

	log.Println(len(pubB))
	log.Println(len(privB))

	transportA, err := GetTransportKey(privA, pubA)
	if err != nil {
		t.Errorf("[GenerateTransportKey] ERR [A] %v", err)
		t.Fail()
	}

	transportAA, err := GetTransportKey(privA, pubA)
	if err != nil {
		t.Errorf("[GenerateTransportKey] ERR [B] %v", err)
		t.Fail()
	}

	message := "hello"

	en, err := Encrypt(message, string(transportA))
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}

	dec, err := Decrypt(en, string(transportAA))
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}

	if message != dec {
		t.Fail()
	}
}

func TestGetPublicKeyAndAddressByPrivateKey(t *testing.T) {
	priv := "c2e136ddbfe07da18ee01993c2d9f93deb669face627bba4b702fa1ba17fcfb4"

	pub, err := GetPublicKeyAndAddressByPrivateKey(priv)
	if err != nil {
		t.Errorf("[GetPublicKeyAndAddressByPrivateKey] ERR %v", err)
		t.Fail()
	}

	log.Println(len(pub), pub)
}
