package dfabric

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hyperledger/fabric/common/crypto"
	"github.com/hyperledger/fabric/common/localmsp"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/msp/cache"
	"github.com/hyperledger/fabric/msp/mgmt"
	mspmgmt "github.com/hyperledger/fabric/msp/mgmt"
	cb "github.com/hyperledger/fabric/protos/common"
)

func Test_loadMSP(t *testing.T) {
	root := "D:/git/BaaSService/DeployService/app/daction/test/"
	path := filepath.Join(root, "crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp")

	sigHeader := loadLeveltest(t, 5, path, "Org1")
	fmt.Println(sigHeader)

	fmt.Println(strings.Repeat("---", 15))

	path = filepath.Join(root, "crypto-config/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp")

	sigHeader = loadLeveltest(t, 5, path, "Org3")
	fmt.Println(sigHeader)
}

func loadLeveltest(t *testing.T, version int, path, orgID string) *cb.SignatureHeader {

	var oneMSP msp.MSP
	var sigHeader *cb.SignatureHeader
	switch version {
	case 1:
		err := mgmt.LoadLocalMsp(path, nil, orgID)
		if err != nil {
			t.Fatal(err)
		}

		//---

		signer := localmsp.NewSigner()
		sigHeader, err = signer.NewSignatureHeader()
		if err != nil {
			t.Fatal(err)
		}
	case 2:
		conf, err := msp.GetLocalMspConfig(path, nil, "Org1")
		if err != nil {
			t.Fatal(err)
		}

		tmp := mgmt.GetLocalMSP()
		err = tmp.Setup(conf)
		if err != nil {
			t.Fatal(err)
		}

		//---

		one := mspmgmt.GetLocalMSP()
		signer, err := one.GetDefaultSigningIdentity()
		if err != nil {
			t.Fatal(err)
		}

		creatorIdentityRaw, err := signer.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		nonce, err := crypto.GetRandomNonce()
		if err != nil {
			t.Fatal(err)
		}

		sigHeader = &cb.SignatureHeader{}
		sigHeader.Creator = creatorIdentityRaw
		sigHeader.Nonce = nonce
	case 3:
		conf, err := msp.GetLocalMspConfig(path, nil, "Org1")
		if err != nil {
			t.Fatal(err)
		}

		oneMSP = mgmt.GetLocalMSP()
		err = oneMSP.Setup(conf)
		if err != nil {
			t.Fatal(err)
		}

		//---

		signer, err := oneMSP.GetDefaultSigningIdentity()
		if err != nil {
			t.Fatal(err)
		}

		creatorIdentityRaw, err := signer.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		nonce, err := crypto.GetRandomNonce()
		if err != nil {
			t.Fatal(err)
		}

		sigHeader = &cb.SignatureHeader{}
		sigHeader.Creator = creatorIdentityRaw
		sigHeader.Nonce = nonce

	case 4:
		// local bccsp init
		conf, err := msp.GetLocalMspConfig(path, nil, "Org1")
		if err != nil {
			t.Fatal(err)
		}

		// determine the type of MSP (by default, we'll use bccspMSP)
		mspType := msp.ProviderTypeToString(msp.FABRIC)

		var mspOpts = map[string]msp.NewOpts{
			msp.ProviderTypeToString(msp.FABRIC): &msp.BCCSPNewOpts{NewBaseOpts: msp.NewBaseOpts{Version: msp.MSPv1_0}},
			msp.ProviderTypeToString(msp.IDEMIX): &msp.IdemixNewOpts{NewBaseOpts: msp.NewBaseOpts{Version: msp.MSPv1_1}},
		}

		newOpts, found := mspOpts[mspType]
		if !found {
			t.Fatal("msp type " + mspType + " unknown")
		}

		// local bccsp load
		oneMSP, err := msp.New(newOpts)
		if err != nil {
			t.Fatal("Failed to initialize local MSP, received err " + err.Error())
		}
		switch mspType {
		case msp.ProviderTypeToString(msp.FABRIC):
			oneMSP, err = cache.New(oneMSP)
			if err != nil {
				t.Fatal("Failed to initialize local MSP, received err " + err.Error())
			}
		case msp.ProviderTypeToString(msp.IDEMIX):
			// Do nothing
		default:
			t.Fatal("msp type " + mspType + " unknown")
		}

		err = oneMSP.Setup(conf)
		if err != nil {
			t.Fatal(err)
		}

		//---

		signer, err := oneMSP.GetDefaultSigningIdentity()
		if err != nil {
			t.Fatal(err)
		}

		creatorIdentityRaw, err := signer.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		nonce, err := crypto.GetRandomNonce()
		if err != nil {
			t.Fatal(err)
		}

		sigHeader = &cb.SignatureHeader{}
		sigHeader.Creator = creatorIdentityRaw
		sigHeader.Nonce = nonce

	case 5:

		csp, conf, err := msp.GetLocalMspConfigAndBCCSP(path, nil, "Org1")
		if err != nil {
			t.Fatal(err)
		}

		// determine the type of MSP (by default, we'll use bccspMSP)
		mspType := msp.ProviderTypeToString(msp.FABRIC)

		var mspOpts = map[string]msp.NewOpts{
			msp.ProviderTypeToString(msp.FABRIC): &msp.BCCSPNewOpts{NewBaseOpts: msp.NewBaseOpts{Version: msp.MSPv1_0}},
			msp.ProviderTypeToString(msp.IDEMIX): &msp.IdemixNewOpts{NewBaseOpts: msp.NewBaseOpts{Version: msp.MSPv1_1}},
		}

		newOpts, found := mspOpts[mspType]
		if !found {
			t.Fatal("msp type " + mspType + " unknown")
		}

		oneMSP, err := msp.NewByBccsp(newOpts, csp)
		if err != nil {
			t.Fatal("Failed to initialize local MSP, received err " + err.Error())
		}
		switch mspType {
		case msp.ProviderTypeToString(msp.FABRIC):
			oneMSP, err = cache.New(oneMSP)
			if err != nil {
				t.Fatal("Failed to initialize local MSP, received err " + err.Error())
			}
		case msp.ProviderTypeToString(msp.IDEMIX):
			// Do nothing
		default:
			t.Fatal("msp type " + mspType + " unknown")
		}

		err = oneMSP.Setup(conf)
		if err != nil {
			t.Fatal(err)
		}

		//---

		signer, err := oneMSP.GetDefaultSigningIdentity()
		if err != nil {
			t.Fatal(err)
		}

		creatorIdentityRaw, err := signer.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		nonce, err := crypto.GetRandomNonce()
		if err != nil {
			t.Fatal(err)
		}

		sigHeader = &cb.SignatureHeader{}
		sigHeader.Creator = creatorIdentityRaw
		sigHeader.Nonce = nonce
	}

	return sigHeader
}
