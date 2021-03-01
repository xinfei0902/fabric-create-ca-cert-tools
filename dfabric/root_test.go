package dfabric

import (
	"testing"
)

func Test_GetCA(t *testing.T) {

	path := "D:/git/codes/test1.3.1/target/crypto-config/ordererOrganizations/ordererorg.example.com/"

	output := "./test/orderer/"

	err := BuildOrdererNodeByFiles(output, path, "ordererorg.example.com", "orderer1")
	if err != nil {
		t.Fatal(err)
	}

	path = "D:/git/BaaSService/DeployService/test/output/aaa/crypto-config/peerOrganizations/org1.basedomain.com"

	output = "./test/peer/"

	err = BuildPeerNodeByFiles(output, path, "org1.basedomain.com", "peer3", "User5")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_BuildNewPeerInMSP(t *testing.T) {
	path := "D:/git/codes/test1.3.1/target/crypto-config/peerOrganizations/org1.example.com"
	err := BuildNewPeerInMSP(path, "org1.example.com", "peer5", "User5")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_BuildOneOrganization(t *testing.T) {
	output := "./test/aorg"
	err := BuildOneOrganization(output, "org1.example.com", DefaultCASpec("org1.example.com"))
	if err != nil {
		t.Fatal(err)
	}
}

func Test_BuildByNew(t *testing.T) {
	output := "./test/bbb"
	err := BuildOneOrganization(output, "org1.example.com", DefaultCASpec("org1.example.com"))
	if err != nil {
		t.Fatal(err)
	}

	err = BuildNewPeerInMSP(output, "org1.example.com", "peer5", "User5")
	if err != nil {
		t.Fatal(err)
	}
}
