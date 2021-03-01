package dfabric

import (
	"github.com/hyperledger/fabric/common/tools/cryptogen/ca"
	"github.com/hyperledger/fabric/common/tools/cryptogen/csp"
	"github.com/pkg/errors"
)

type NodeSpec struct {
	Hostname           string   `yaml:"Hostname"`
	CommonName         string   `yaml:"CommonName"`
	Country            string   `yaml:"Country"`
	Province           string   `yaml:"Province"`
	Locality           string   `yaml:"Locality"`
	OrganizationalUnit string   `yaml:"OrganizationalUnit"`
	StreetAddress      string   `yaml:"StreetAddress"`
	PostalCode         string   `yaml:"PostalCode"`
	SANS               []string `yaml:"SANS"`
}

func DefaultCASpec(cn string) *NodeSpec {
	return &NodeSpec{
		CommonName:         cn,
		Country:            "",
		Province:           "",
		Locality:           "",
		OrganizationalUnit: "",
		StreetAddress:      "",
		PostalCode:         "",
	}
}

func MakeCA(caDir string, spec *NodeSpec, orgName string) (*ca.CA, error) {
	return ca.NewCA(caDir, orgName, spec.CommonName, spec.Country, spec.Province, spec.Locality, spec.OrganizationalUnit, spec.StreetAddress, spec.PostalCode)
}

func GetCA(caDir string, spec *NodeSpec) (*ca.CA, error) {
	_, signer, err := csp.LoadPrivateKey(caDir)
	if err != nil {
		err = errors.WithMessage(err, "Load private key error")
		return nil, err
	}
	cert, err := ca.LoadCertificateECDSA(caDir)
	if err != nil {
		err = errors.WithMessage(err, "Load cert key error")
		return nil, err
	}

	return &ca.CA{
		Name:               spec.CommonName,
		Signer:             signer,
		SignCert:           cert,
		Country:            spec.Country,
		Province:           spec.Province,
		Locality:           spec.Locality,
		OrganizationalUnit: spec.OrganizationalUnit,
		StreetAddress:      spec.StreetAddress,
		PostalCode:         spec.PostalCode,
	}, nil
}
