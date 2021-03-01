package dfabric

import (
	"path/filepath"

	"github.com/hyperledger/fabric/common/tools/cryptogen/ca"
	"github.com/hyperledger/fabric/common/tools/cryptogen/msp"
	"github.com/pkg/errors"
)

func BuildOrdererNodeByFiles(output, MSPPath, baseDomain, host string) (err error) {
	signca, tlsca, err := LoadCAObjectFromFiles(MSPPath, baseDomain)
	if err != nil {
		return err
	}

	err = BuildOrdererNode(output, signca, tlsca, baseDomain, host)
	if err != nil {
		err = errors.WithMessage(err, "build MSP error")
		return
	}

	err = copyAdminCert(filepath.Join(MSPPath, "users"), filepath.Join(output, "msp", "admincerts"), "Admin@"+baseDomain)
	if err != nil {
		err = errors.WithMessage(err, "copy user cert error")
		return
	}
	return nil
}

func BuildOrdererNode(output string, signca, tlsca *ca.CA, baseDomain, host string) (err error) {
	url := host + "." + baseDomain

	return msp.GenerateLocalMSP(filepath.Join(output, "orderer", url), url, []string{url}, signca, tlsca, msp.ORDERER, false)
}

func BuildOrdererAdminUser(output string, signca, tlsca *ca.CA, baseDomain string) (err error) {
	return msp.GenerateLocalMSP(filepath.Join(output, "users"), "Admin@"+baseDomain, []string{}, signca, tlsca, msp.CLIENT, false)
}
