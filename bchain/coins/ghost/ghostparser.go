package ghost

import (
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

const (
	MainnetMagic wire.BitcoinNet = 0xD9B4BEF9
	TestnetMagic wire.BitcoinNet = 0xDAB5BFFA
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic

	err := chaincfg.Register(&MainNetParams)
	if err != nil {
		panic(err)
	}
}

// BitcoreParser handle
type GhostParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

// NewBitcoreParser returns new BitcoreParser instance
func NewGhostParser(params *chaincfg.Params, c *btc.Configuration) *GhostParser {
	p := &GhostParser{
		BitcoinParser: btc.NewBitcoinParser(params, c),
		baseparser:    &bchain.BaseParser{},
	}
	return p
}

// GetChainParams contains network parameters for the main Bitcore network,
// and the test Bitcore network
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}

func (p *GhostParser) GetAddrDescFromAddress(address string) (bchain.AddressDescriptor, error) {
	addressByte := []byte(address)
	return bchain.AddressDescriptor(addressByte), nil
}

func (p *GhostParser) GetAddrDescFromVout(output *bchain.Vout) (bchain.AddressDescriptor, error) {
	addrs := output.ScriptPubKey.Addresses
	if addrs == nil || len(addrs) == 0 {
		return nil, nil
	}
	var addressByte []byte
	for i := range output.ScriptPubKey.Addresses {
		addressByte = append(addressByte, output.ScriptPubKey.Addresses[i]...)
	}
	return bchain.AddressDescriptor(addressByte), nil
}

func (p *GhostParser) GetAddressesFromAddrDesc(addrDesc bchain.AddressDescriptor) ([]string, bool, error) {
	var addrs []string
	if addrDesc != nil {
		addrs = append(addrs, string(addrDesc))
	}
	return addrs, true, nil
}
