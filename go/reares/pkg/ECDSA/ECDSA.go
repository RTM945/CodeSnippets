package ECDSA

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const interval = 60 * 1000

type ECDSA struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

var mu sync.RWMutex
var lastCreateTime int64
var instance *ECDSA

func GetInstance() *ECDSA {
	now := time.Now().Unix()
	intervalTime := now - interval
	if intervalTime >= atomic.LoadInt64(&lastCreateTime) {
		mu.Lock()
		if intervalTime >= atomic.LoadInt64(&lastCreateTime) {
			publicKey, privateKey := genKeyPair()
			tmp := &ECDSA{
				privateKey: privateKey,
				publicKey:  publicKey,
			}
			instance = tmp
			atomic.StoreInt64(&lastCreateTime, now)
		}
		mu.Unlock()
	}
	return instance
}

func genKeyPair() (*ecdsa.PublicKey, *ecdsa.PrivateKey) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(fmt.Errorf("failed to generate key: %v", err))
	}
	publicKey := &privateKey.PublicKey
	return publicKey, privateKey
}

func (e *ECDSA) GetPublicKey() *ecdsa.PublicKey {
	return e.publicKey
}

func (e *ECDSA) GetPrivateKey() *ecdsa.PrivateKey {
	return e.privateKey
}

func (e *ECDSA) GetPublicKeyEncoded() []byte {
	key, err := x509.MarshalPKIXPublicKey(e.publicKey)
	if err != nil {
		return nil
	}
	return key
}
