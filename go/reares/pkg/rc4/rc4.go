package rc4

import (
	"crypto/rc4"
)

type RC4 struct {
	key    []byte
	cipher *rc4.Cipher
}

func NewRC4(key []byte) *RC4 {
	cipher, _ := rc4.NewCipher(key)
	res := &RC4{
		key:    key,
		cipher: cipher,
	}

	return res
}

func (r *RC4) DoUpdate(data []byte) {
	r.cipher.XORKeyStream(data, data)
}
