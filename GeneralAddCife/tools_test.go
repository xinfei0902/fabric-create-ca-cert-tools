package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"../dfabric"
	"../objectdefine"
)

func Test_runCustom(t *testing.T) {
	path := "E:/BaasTools/addpeerorigin/DeployService/test/output/addorg.json"
	output := "E:/BaasTools/addpeerorigin/DeployService/test/output/test/crypto-config/peerOrganizations/org3msp.example.com"
	fmt.Println(path)
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(buff))
	indent := &objectdefine.Indent{}
	err = json.Unmarshal(buff, indent)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("build org")
	for orgname, org := range indent.Org {
		err := dfabric.BuildOneOrganization(output, orgname, dfabric.DefaultCASpec(org.OrgDomain))
		if err != nil {

			fmt.Println(err)
			return
		}
		for _, peer := range org.Peer {
			err = dfabric.BuildNewPeerInMSP(output, peer.OrgDomain, peer.Name, "")
			if err != nil {

				fmt.Println(err)
				return
			}
		}

	}

}
