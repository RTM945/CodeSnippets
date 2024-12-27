package client_switch

type Processor interface {
	ProcessRSAKeyExchange(msg *RSAKeyExchange) error
	ProcessKeyExchange(msg *KeyExchange) error
}
