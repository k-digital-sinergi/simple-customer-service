package db

import (
	"database/sql/driver"
	"simple-customer-service/pkg/util/encryption"
)

type AES256 string

func (s AES256) Value() (driver.Value, error) {
	cipherText, err := encryption.Encrypt([]byte(s))
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

func (s *AES256) Scan(value interface{}) error {
	cipherText, _ := value.([]byte)
	plainText, err := encryption.Decrypt(string(cipherText))
	if err != nil {
		return err
	}
	*s = AES256(plainText)
	return nil
}
