package security

import (
	"crypto/cipher"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"

	"github.com/go-pantheon/fabrica-util/security/aes"
	rrsa "github.com/go-pantheon/fabrica-util/security/rsa"
	"github.com/go-pantheon/fabrica-util/xrand"
	"github.com/pkg/errors"
)

var (
	handshakePriKey *rsa.PrivateKey
	tokenAESKey     []byte
	tokenAESBlock   cipher.Block
)

func Init(aesKey string, priKey string) error {
	var (
		priKeyBytes []byte
		priKeyIface interface{}
		err         error
	)

	if priKeyBytes, err = base64.URLEncoding.DecodeString(priKey); err != nil {
		return errors.Wrapf(err, "base64 DecodeString failed.")
	}
	if priKeyIface, err = x509.ParsePKCS8PrivateKey(priKeyBytes); err != nil {
		return errors.Wrapf(err, "x509 ParsePKCS8PrivateKey failed.")
	}
	handshakePriKey = priKeyIface.(*rsa.PrivateKey)

	tokenAESKey = []byte(aesKey)

	var block cipher.Block
	if block, err = aes.NewBlock(tokenAESKey); err != nil {
		return errors.Wrapf(err, "aes NewBlock failed.")
	}
	tokenAESBlock = block

	return nil
}

func InitApiCrypto() (cipher.Block, []byte, error) {
	str, err := xrand.RandAlphaNumString(32)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "xrand RandString failed.")
	}

	key := []byte(str)
	block, err := aes.NewBlock(key)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "create aes block failed.")
	}
	return block, key, nil
}

func DecryptCSHandshake(secret []byte) ([]byte, error) {
	return rrsa.Decrypt(handshakePriKey, secret)
}

func DecryptToken(secret string) ([]byte, error) {
	ser, err := base64.URLEncoding.DecodeString(secret)
	if err != nil {
		return nil, errors.Wrapf(err, "base64 DecodeString failed.")
	}

	origin, err := aes.Decrypt(tokenAESKey, tokenAESBlock, ser)
	if err != nil {
		return nil, errors.Wrapf(err, "aes Decrypt failed.")
	}
	return origin, nil
}
