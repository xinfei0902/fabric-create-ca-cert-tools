package dfabric

import (
	"github.com/hyperledger/fabric/common/tools/configtxgen/encoder"
	"github.com/hyperledger/fabric/common/tools/configtxgen/localconfig"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/pkg/errors"
)

func BuildGenesisBlock(config *localconfig.Profile, channelID string) (buff []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.Errorf("Genesis block error: %v", e)
		}
	}()

	pgen := encoder.New(config)
	if len(config.Consortiums) == 0 {
		return nil, errors.New("Genesis block does not contain a consortiums group definition")
	}
	genesisBlock := pgen.GenesisBlockForChannel(channelID)

	return utils.Marshal(genesisBlock)
}

func BuildChannelCreateTx(conf *localconfig.Profile, channelID string) (buff []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.Errorf("Genesis ChannelTx error: %v", e)
		}
	}()

	configtx, err := encoder.MakeChannelCreationTransaction(channelID, nil, conf)
	if err != nil {
		return nil, err
	}

	return utils.Marshal(configtx)
}
