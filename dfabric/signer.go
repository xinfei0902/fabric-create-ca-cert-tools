package dfabric

import (
	"fmt"

	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/common/crypto"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/msp/cache"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/pkg/errors"
)

type mspSignerPath struct {
	dir         string
	bccspConfig *factory.FactoryOpts
	mspID       string
	signer      msp.SigningIdentity
}

// NewSignerByPath returns a new instance of the msp-based LocalSigner.
// It assumes that the local msp has been already initialized.
// Look at mspmgmt.LoadLocalMsp for further information.
func NewSignerByPath(path, mspID string) crypto.LocalSigner {
	return &mspSignerPath{
		dir:         path,
		bccspConfig: nil,
		mspID:       mspID,
	}
}

// NewSignatureHeader creates a SignatureHeader with the correct signing identity and a valid nonce
func (s *mspSignerPath) NewSignatureHeader() (*cb.SignatureHeader, error) {
	err := s.getSigner()
	if err != nil {
		return nil, err
	}

	creatorIdentityRaw, err := s.signer.Serialize()
	if err != nil {
		return nil, fmt.Errorf("Failed serializing creator public identity [%s]", err)
	}

	nonce, err := crypto.GetRandomNonce()
	if err != nil {
		return nil, fmt.Errorf("Failed creating nonce [%s]", err)
	}

	sh := &cb.SignatureHeader{}
	sh.Creator = creatorIdentityRaw
	sh.Nonce = nonce

	return sh, nil
}

// Sign a message which should embed a signature header created by NewSignatureHeader
func (s *mspSignerPath) Sign(message []byte) ([]byte, error) {
	err := s.getSigner()
	if err != nil {
		return nil, err
	}

	signature, err := s.signer.Sign(message)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed generating signature")
	}

	return signature, nil
}

func (s *mspSignerPath) getSigner() (err error) {
	if s.signer != nil {
		return nil
	}

	csp, conf, err := msp.GetLocalMspConfigAndBCCSP(s.dir, s.bccspConfig, s.mspID)
	if err != nil {
		err = errors.WithMessage(err, "Local signer MSP and BCCSP error")
		return err
	}

	// determine the type of MSP (by default, we'll use bccspMSP)
	mspType := msp.ProviderTypeToString(msp.FABRIC)

	var mspOpts = map[string]msp.NewOpts{
		msp.ProviderTypeToString(msp.FABRIC): &msp.BCCSPNewOpts{NewBaseOpts: msp.NewBaseOpts{Version: msp.MSPv1_0}},
		msp.ProviderTypeToString(msp.IDEMIX): &msp.IdemixNewOpts{NewBaseOpts: msp.NewBaseOpts{Version: msp.MSPv1_1}},
	}

	newOpts, found := mspOpts[mspType]
	if !found {
		err = errors.New("msp type " + mspType + " unknown")
		return err
	}

	oneMSP, err := msp.NewByBccsp(newOpts, csp)
	if err != nil {
		err = errors.WithMessage(err, "Failed to initialize local MSP, received error")
		return err
	}
	switch mspType {
	case msp.ProviderTypeToString(msp.FABRIC):
		oneMSP, err = cache.New(oneMSP)
		if err != nil {
			err = errors.WithMessage(err, "Failed to initialize local MSP, received error")
			return err
		}
	case msp.ProviderTypeToString(msp.IDEMIX):
		// Do nothing
	default:
		if err != nil {
			err = errors.New("msp type " + mspType + " unknown")
			return err
		}
	}

	err = oneMSP.Setup(conf)
	if err != nil {
		err = errors.WithMessage(err, "Failed to setup local MSP, received error")
		return err
	}

	signer, err := oneMSP.GetDefaultSigningIdentity()
	if err != nil {
		err = errors.WithMessage(err, "Get default signer error")
		return err
	}

	s.signer = signer
	return nil
}
