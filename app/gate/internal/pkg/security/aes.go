package security

import (
	"encoding/base64"

	"github.com/go-pantheon/fabrica-util/errors"
)

func DecryptAccountToken(secret string) ([]byte, error) {
	if !manager.initialized.Load() {
		return nil, errors.New("crypto manager not initialized, call Init() first")
	}

	ser, err := base64.URLEncoding.DecodeString(secret)
	if err != nil {
		return nil, errors.Wrap(err, "base64 DecodeString failed")
	}

	manager.mu.RLock()
	tokenAESCipher := manager.tokenAES
	manager.mu.RUnlock()

	org, err := tokenAESCipher.Decrypt(ser)
	if err != nil {
		return nil, errors.Wrap(err, "aes Decrypt failed")
	}

	return org, nil
}
