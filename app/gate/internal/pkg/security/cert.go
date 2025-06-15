package security

import (
	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/go-pantheon/fabrica-util/security/certificate"
)

func VerifyECDHCliPubKey(data, sign []byte) error {
	if !manager.initialized.Load() {
		return errors.New("crypto manager not initialized, call Init() first")
	}

	manager.mu.RLock()
	cliPub := manager.cliCertPub
	manager.mu.RUnlock()

	if !certificate.Verify(cliPub, data, sign) {
		return errors.New("certificate Verify failed")
	}

	return nil
}

func SignECDHSvrPubKey(key []byte) ([]byte, error) {
	if !manager.initialized.Load() {
		return nil, errors.New("crypto manager not initialized, call Init() first")
	}

	manager.mu.RLock()
	svrPri := manager.svrCertPri
	manager.mu.RUnlock()

	ret, err := certificate.Sign(svrPri, key)
	if err != nil {
		return nil, err
	}

	return ret.Sign, nil
}
