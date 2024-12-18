package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"sync"
	"sync/atomic"
	"time"
)

type Key struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

var lastCreateTime int64

var instance *Key

const interval = 60 * 1000

var mu sync.Mutex

func GetInstance() *Key {
	now := time.Now().Unix()
	intervalTime := now - interval
	if intervalTime >= atomic.LoadInt64(&lastCreateTime) {
		mu.Lock()
		if intervalTime >= atomic.LoadInt64(&lastCreateTime) {
			tmp := &Key{}
			key, err := generateKey()
			if err != nil {
				println(err)
				return nil
			}
			tmp.privateKey = key
			tmp.publicKey = &key.PublicKey
			instance = tmp
			atomic.StoreInt64(&lastCreateTime, now)
		}
		mu.Unlock()
	}

	return instance
}

func generateKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func (key *Key) GetPublicKey() *rsa.PublicKey {
	return key.publicKey
}

func (key *Key) GetPrivateKey() *rsa.PrivateKey {
	return key.privateKey
}

func (key *Key) GetPublicKeyEncoded() ([]byte, error) {
	return x509.MarshalPKIXPublicKey(key.publicKey)
}

func Encrypt(publicKeyBytes, data []byte) ([]byte, error) {
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey.(*rsa.PublicKey), data, nil)
}

func Decrypt(privateKey *rsa.PrivateKey, cipherData []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, cipherData, nil)
}
