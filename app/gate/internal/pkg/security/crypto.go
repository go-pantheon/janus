package security

import (
	"crypto/ed25519"
	"sync"
	"sync/atomic"

	"github.com/go-pantheon/fabrica-util/security/aes"
	"github.com/go-pantheon/fabrica-util/security/certificate"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
)

var (
	once    sync.Once
	manager = &CryptoManager{}
)

type CryptoManager struct {
	mu          sync.RWMutex
	initialized atomic.Bool
	svrCertPri  ed25519.PrivateKey
	cliCertPub  ed25519.PublicKey
	tokenAES    *aes.Cipher
}

func Init(config *conf.Secret) (err error) {
	once.Do(func() {
		err = initCryptoManager(config)
	})

	return err
}

func initCryptoManager(c *conf.Secret) (err error) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	aesCipher, err := aes.NewAESCipher([]byte(c.AccountAesKey))
	if err != nil {
		return err
	}

	svrPri, err := certificate.ImportPriFromPEM([]byte(c.ServerCertPrivateKey))
	if err != nil {
		return err
	}

	clientCert, err := certificate.ImportCertFromPEM([]byte(c.ClientCert))
	if err != nil {
		return err
	}

	if err := certificate.VerifyCert(clientCert); err != nil {
		return err
	}

	cliPub, err := certificate.ExportPubFromCert(clientCert)
	if err != nil {
		return err
	}

	manager.tokenAES = aesCipher
	manager.svrCertPri = svrPri
	manager.cliCertPub = cliPub
	manager.initialized.Store(true)

	return nil
}
