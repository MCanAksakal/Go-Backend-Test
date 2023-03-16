package Crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Encrypt(text string) (string, error) {

	err := godotenv.Load("Crypt.env")
	if err != nil {
		log.Fatal("Crypt.env load failed")
	} else {
		log.Printf("Crypt.env loaded successfully")
	}

	keyCode := os.Getenv("CP_Secret")

	block, err := aes.NewCipher([]byte(keyCode))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func Decrypt(text string) (string, error) {

	err := godotenv.Load("Crypt.env")
	if err != nil {
		log.Fatal("Crypt.env load failed")
	} else {
		log.Printf("Crypt.env loaded successfully")
	}

	keyCode := os.Getenv("CP_Secret")

	block, err := aes.NewCipher([]byte(keyCode))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
