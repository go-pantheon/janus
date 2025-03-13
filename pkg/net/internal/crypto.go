package internal

import (
	"github.com/go-pantheon/vulcan-util/security/aes"
	"github.com/pkg/errors"
	"github.com/vulcan-frame/vulcan-gate/pkg/net"
)

func encrypt(ss net.Session, data []byte) ([]byte, error) {
	if !ss.IsCrypto() {
		return data, nil
	}

	result, err := aes.Encrypt(ss.Key(), ss.Block(), data)
	if err != nil {
		return nil, errors.Wrapf(err, "packet encrypt failed")
	}
	return result, nil
}

func decrypt(ss net.Session, data []byte) ([]byte, error) {
	if !ss.IsCrypto() {
		return data, nil
	}

	result, err := aes.Decrypt(ss.Key(), ss.Block(), data)
	if err != nil {
		return nil, errors.WithMessage(err, "packet decrypt failed")
	}
	return result, nil
}
