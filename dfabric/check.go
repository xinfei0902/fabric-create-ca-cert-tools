package dfabric

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/common/tools/cryptogen/ca"
)

func DebugCAKeys(name string, input *ca.CA) {

	fmt.Println(strings.Repeat("-", 10), name, strings.Repeat("-", 10))

	fmt.Println(input.Signer.Public())
	fmt.Println(input.SignCert.Subject.Organization)

	fmt.Println(strings.Repeat("=", 20))

}
