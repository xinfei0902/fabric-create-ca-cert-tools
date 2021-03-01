package dfabric

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/common/tools/configtxgen/localconfig"
)

func Test_BuildGenesisBlockJ(t *testing.T) {

	output := "./test/org1"
	err := BuildOneOrganization(output, "org1.example.com", DefaultCASpec("org1.example.com"))
	if err != nil {
		t.Fatal(err)
	}

	err = BuildNewPeerInMSP(output, "org1.example.com", "peer5", "User5")
	if err != nil {
		t.Fatal(err)
	}

	output = "./test/orderer"
	err = BuildOneOrganization(output, "example.com", DefaultCASpec("example.com"))
	if err != nil {
		t.Fatal(err)
	}

	err = BuildOrdererNodeByFiles(output, output, "example.com", "orderer0")
	if err != nil {
		t.Fatal(err)
	}

	org := []*localconfig.Organization{
		&localconfig.Organization{
			ID:             "Orderer",
			Name:           "Orderer",
			AdminPrincipal: "",
			MSPType:        "bccsp",
			Policies: map[string]*localconfig.Policy{
				"Readers": &localconfig.Policy{
					Type: "ImplicitMeta",
					Rule: "ANY Readers",
				},
				"Writers": &localconfig.Policy{
					Type: "ImplicitMeta",
					Rule: "ANY Writers",
				},
				"Admins": &localconfig.Policy{
					Type: "ImplicitMeta",
					Rule: "MAJORITY Admins",
				},
				"BlockValidation": &localconfig.Policy{
					Type: "ImplicitMeta",
					Rule: "MAJORITY Writers",
				},
			},
			MSPDir: "./test/orderer/msp",
		},
		&localconfig.Organization{
			ID:             "Org1",
			Name:           "Org1",
			AdminPrincipal: "",
			MSPType:        "bccsp",
			Policies: map[string]*localconfig.Policy{
				"Readers": &localconfig.Policy{
					Type: "Signature",
					Rule: "OR('Org1.member')",
				},
				"Writers": &localconfig.Policy{
					Type: "Signature",
					Rule: "OR('Org1.member')",
				},
				"Admins": &localconfig.Policy{
					Type: "Signature",
					Rule: "OR('Org1.member')",
				},
			},
			MSPDir: "./test/org1/msp",
			AnchorPeers: []*localconfig.AnchorPeer{
				&localconfig.AnchorPeer{
					Host: "peer0.org1.example.com",
					Port: 7051,
				},
			},
		},
	}

	one := &localconfig.Profile{}

	one.Consortium = "mychannel"

	one.Capabilities = map[string]bool{
		"V1_1": true,
	}
	one.Policies = map[string]*localconfig.Policy{
		"Readers": &localconfig.Policy{
			Type: "ImplicitMeta",
			Rule: "ANY Readers",
		},
		"Writers": &localconfig.Policy{
			Type: "ImplicitMeta",
			Rule: "ANY Writers",
		},
		"Admins": &localconfig.Policy{
			Type: "ImplicitMeta",
			Rule: "MAJORITY Admins",
		},
	}
	one.Orderer = &localconfig.Orderer{
		OrdererType: "kafka",
		Addresses:   []string{"orderer0.example.com:7053"},
		BatchSize: localconfig.BatchSize{
			AbsoluteMaxBytes:  10 * 1024 * 1024,
			PreferredMaxBytes: 512 * 1024,
			MaxMessageCount:   10,
		},
		Kafka: localconfig.Kafka{
			Brokers: []string{"kafka0.example.com:9092"},
		},
		Organizations: org[0:1],
	}
	one.Application = &localconfig.Application{
		ACLs: map[string]string{
			"lscc/ChaincodeExists": "/Channel/Application/Readers",

			// # ACL policy for lscc's "getdepspec" function
			"lscc/GetDeploymentSpec": "/Channel/Application/Readers",

			// # ACL policy for lscc's "getccdata" function
			"lscc/GetChaincodeData": "/Channel/Application/Readers",

			// # ACL Policy for lscc's "getchaincodes" function
			"lscc/GetInstantiatedChaincodes": "/Channel/Application/Readers",

			// #---Query System Chaincode (qscc) function to policy mapping for access control---// #

			// # ACL policy for qscc's "GetChainInfo" function
			"qscc/GetChainInfo": "/Channel/Application/Readers",

			// # ACL policy for qscc's "GetBlockByNumber" function
			"qscc/GetBlockByNumber": "/Channel/Application/Readers",

			// # ACL policy for qscc's  "GetBlockByHash" function
			"qscc/GetBlockByHash": "/Channel/Application/Readers",

			// # ACL policy for qscc's "GetTransactionByID" function
			"qscc/GetTransactionByID": "/Channel/Application/Readers",

			// # ACL policy for qscc's "GetBlockByTxID" function
			"qscc/GetBlockByTxID": "/Channel/Application/Readers",

			// #---Configuration System Chaincode (cscc) function to policy mapping for access control---// #

			// # ACL policy for cscc's "GetConfigBlock" function
			"cscc/GetConfigBlock": "/Channel/Application/Readers",

			// # ACL policy for cscc's "GetConfigTree" function
			"cscc/GetConfigTree": "/Channel/Application/Readers",

			// # ACL policy for cscc's "SimulateConfigTreeUpdate" function
			"cscc/SimulateConfigTreeUpdate": "/Channel/Application/Readers",

			// #---Miscellanesous peer function to policy mapping for access control---// #

			// # ACL policy for invoking chaincodes on peer
			"peer/Propose": "/Channel/Application/Writers",

			// # ACL policy for chaincode to chaincode invocation
			"peer/ChaincodeToChaincode": "/Channel/Application/Readers",

			// #---Events resource to policy mapping for access control// #// #// #---// #

			// # ACL policy for sending block events
			"event/Block": "/Channel/Application/Readers",

			// # ACL policy for sending filtered block events
			"event/FilteredBlock": "/Channel/Application/Readers",
		},
		Organizations: org,
	}

	one.Consortiums = map[string]*localconfig.Consortium{
		"mychannel": &localconfig.Consortium{
			Organizations: org,
		},
	}

	buff, err := BuildGenesisBlock(one, "mychannel")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(buff))

	buff, err = BuildChannelCreateTx(one, "mychannel")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(buff))
}
