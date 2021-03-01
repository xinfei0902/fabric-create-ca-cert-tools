package main

import (
	"fmt"
	"os"

	"../dfabric"
)

func main() {
	handle := os.Args[1]

	switch handle {
	case "peer":
		if len(os.Args) < 6 {
			fmt.Println("Usage: ./tools peer {orgRootPath} {baseDomain} {SubHost of Peer} {UserName}")
			os.Exit(1)
		}
		err := dfabric.BuildNewPeerInMSP(os.Args[2], os.Args[3], os.Args[4], os.Args[5])
		if err != nil {
			fmt.Println(err)
		}

	case "org":
		if len(os.Args) < 5 {
			fmt.Println("Usage: ./tools org {orgRootPath} {orgName} {baseDomain}")
			os.Exit(1)
		}

		err := dfabric.BuildOneOrganization(os.Args[2], os.Args[3], dfabric.DefaultCASpec(os.Args[4]))
		if err != nil {

			fmt.Println(err)
			return
		}
	case "nodes":
		sub := os.Args[2]
		switch sub {
		case "peer":
		case "orderer":
		case "sdk":
		}
	case "help":
	case "version":
	}

}
