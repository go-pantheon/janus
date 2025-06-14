package security

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"

	"github.com/go-pantheon/fabrica-net/xnet"
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/fabrica-util/security/aes"
	rrsa "github.com/go-pantheon/fabrica-util/security/rsa"
	"github.com/go-pantheon/fabrica-util/xrand"
)

var (
	handshakePriKey *rsa.PrivateKey
	tokenAESCipher  *aes.Cipher
)

func Init(aesKey string, priKey string) error {
	var (
		priKeyBytes []byte
		priKeyIface any
		err         error
	)

	if priKeyBytes, err = base64.URLEncoding.DecodeString(priKey); err != nil {
		return errors.Wrap(err, "base64 DecodeString failed")
	}

	if priKeyIface, err = x509.ParsePKCS8PrivateKey(priKeyBytes); err != nil {
		return errors.Wrap(err, "x509 ParsePKCS8PrivateKey failed")
	}

	var ok bool

	handshakePriKey, ok = priKeyIface.(*rsa.PrivateKey)
	if !ok {
		return errors.New("invalid private key")
	}

	tokenAESCipher, err = aes.NewAESCipher([]byte(aesKey))
	if err != nil {
		return errors.Wrap(err, "aes NewAESCipher failed")
	}

	return nil
}

func InitApiCrypto(crypto bool) (xnet.Cryptor, error) {
	if !crypto {
		return xnet.NewNoCryptor(), nil
	}

	str, err := xrand.RandAlphaNumString(32)
	if err != nil {
		return nil, errors.Wrap(err, "xrand RandString failed")
	}

	encryptor, err := xnet.NewCryptor(true, []byte(str))
	if err != nil {
		return nil, errors.Wrap(err, "net NewEncryptor failed")
	}

	return encryptor, nil
}

func DecryptCSHandshake(secret []byte) ([]byte, error) {
	return rrsa.Decrypt(handshakePriKey, secret)
}

func DecryptToken(secret string) ([]byte, error) {
	ser, err := base64.URLEncoding.DecodeString(secret)
	if err != nil {
		return nil, errors.Wrap(err, "base64 DecodeString failed")
	}

	origin, err := tokenAESCipher.Decrypt(ser)
	if err != nil {
		return nil, errors.Wrap(err, "aes Decrypt failed")
	}

	return origin, nil
}
