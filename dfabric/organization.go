package dfabric

import (
	"path/filepath"

	"github.com/hyperledger/fabric/common/tools/cryptogen/ca"
	"github.com/hyperledger/fabric/common/tools/cryptogen/msp"
	"github.com/pkg/errors"
)

func BuildOneOrganization(orgDir, orgName string, orgSpec *NodeSpec) error {
	caDir := filepath.Join(orgDir, "ca")
	tlsCADir := filepath.Join(orgDir, "tlsca")
	mspDir := filepath.Join(orgDir, "msp")
	usrDir := filepath.Join(orgDir, "users")
	adminCertsDir := filepath.Join(mspDir, "admincerts")

	// generate signing CA
	signCA, err := ca.NewCA(caDir, orgName, "ca."+orgSpec.CommonName, orgSpec.Country, orgSpec.Province, orgSpec.Locality, orgSpec.OrganizationalUnit, orgSpec.StreetAddress, orgSpec.PostalCode)
	if err != nil {
		err = errors.WithMessage(err, "Error generating signCA for org "+orgName)
		return err
	}

	// generate TLS CA
	tlsCA, err := ca.NewCA(tlsCADir, orgName, "tlsca."+orgSpec.CommonName, orgSpec.Country, orgSpec.Province, orgSpec.Locality, orgSpec.OrganizationalUnit, orgSpec.StreetAddress, orgSpec.PostalCode)
	if err != nil {
		err = errors.WithMessage(err, "Error generating tlsCA for org "+orgName)
		return err
	}

	err = msp.GenerateVerifyingMSP(mspDir, signCA, tlsCA, false)
	if err != nil {
		err = errors.WithMessage(err, "Error generating MSP for org "+orgName)
		return err
	}

	// admin
	err = msp.GenerateLocalMSP(filepath.Join(usrDir, "Admin@"+orgSpec.CommonName), "Admin@"+orgSpec.CommonName, []string{}, signCA, tlsCA, msp.CLIENT, false)
	if err != nil {
		err = errors.WithMessage(err, "Error generating Admin User for org "+orgSpec.CommonName)
		return err
	}

	// copy the admin cert to the org's MSP admincerts
	err = copyAdminCert(usrDir, adminCertsDir, "Admin@"+orgSpec.CommonName)
	if err != nil {
		err = errors.WithMessage(err, "Error copying Admin cert for org "+orgSpec.CommonName)
		return err
	}

	return nil
}
