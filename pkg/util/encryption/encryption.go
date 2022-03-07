package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
)

const (
	key = "passphrasewhichneedstobe32bytes!"
)

func Encrypt(plainText []byte) (string, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	log.Println("cipherText", cipherText)

	encodedText := hex.EncodeToString(cipherText)

	return encodedText, nil
}

func Decrypt(encodedText string) ([]byte, error) {
	cipherText, err := hex.DecodeString(encodedText)
	if err != nil {
		return nil, err
	}

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	plainText, err := gcm.Open(nil, cipherText[:nonceSize], cipherText[nonceSize:], nil)
	if err != nil {
		return nil, err
	}

	log.Println("plainText", plainText)

	return plainText, nil
}
