package dfabric

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_NewSignerByPath(t *testing.T) {
	root := "D:/git/BaaSService/DeployService/app/daction/test/"
	path := filepath.Join(root, "crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp")

	signer := NewSignerByPath(path, "Org1")
	sigHeader, err := signer.NewSignatureHeader()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(sigHeader)

	path = filepath.Join(root, "crypto-config/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp")

	signer = NewSignerByPath(path, "Org3")
	sigHeader, err = signer.NewSignatureHeader()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(sigHeader)
}
