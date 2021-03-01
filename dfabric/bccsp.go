package dfabric

import (
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/msp"
	"github.com/pkg/errors"
)

func NewOneBccsp(keystoreDir string) (bccsp.BCCSP, error) {
	config := msp.SetupBCCSPKeystoreConfig(nil, keystoreDir)

	f := &factory.SWFactory{}

	csp, err := f.Get(config)
	if err != nil {
		return nil, errors.Errorf("Could not initialize BCCSP %s [%s]", f.Name(), err)
	}

	return csp, nil
}
