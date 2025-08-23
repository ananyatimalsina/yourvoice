package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Expression struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Data string `json:"data" gorm:"not null"`
}

type RSAPrivateKey struct {
	Key rsa.PrivateKey
}

func (r RSAPrivateKey) Value() (driver.Value, error) {
	if r.Key.D == nil {
		return nil, nil
	}
	der := x509.MarshalPKCS1PrivateKey(&r.Key)
	return der, nil
}

func (r *RSAPrivateKey) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte, got %T", value)
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return err
	}
	r.Key = *key
	return nil
}

type Event struct {
	gorm.Model
	Name       string        `json:"name" gorm:"not null"`
	StartDate  time.Time     `json:"start_date" gorm:"not null"`
	EndDate    time.Time     `json:"end_date" gorm:"not null"`
	PrivateKey RSAPrivateKey `json:"private_key" gorm:"not null"`
}

// templ
type InputOption struct {
	Value string
	Label string
}
