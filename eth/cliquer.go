package eth

import (
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
)

type Cliquer struct {
	Node *node.Node
	Eth  *Ethereum
}

func NewCliquer(node *node.Node, eth *Ethereum) Cliquer {
	return Cliquer{
		Node: node,
		Eth:  eth,
	}
}

func (cliquer *Cliquer) GetSigners() ([]common.Address, error) {
	var apis = cliquer.Eth.Engine().APIs(cliquer.Eth.BlockChain())
	var api = apis[0]

	if api.Namespace == "clique" {
		var cliqueApi = api.Service.(*clique.API)
		var blockNum rpc.BlockNumber = rpc.BlockNumber(cliquer.Eth.BlockChain().CurrentBlock().Number.Int64())
		addrs, err := cliqueApi.GetSigners(&blockNum)
		return addrs, err
	} else {
		var err1 = errors.New("Cliquer err")
		errors.Wrap(err1, "GetSigners str")
		return nil, err1
	}
}

func (cliquer *Cliquer) Propose(address common.Address, auth bool) error {
	var apis = cliquer.Eth.Engine().APIs(cliquer.Eth.BlockChain())
	var api = apis[0]

	if api.Namespace == "clique" {
		var cliqueApi = api.Service.(*clique.API)
		cliqueApi.Propose(address, auth)
		return nil
	} else {
		var err1 = errors.New("Cliquer err")
		errors.Wrap(err1, "Propose str")
		return err1
	}
}

// Discard drops a currently running proposal, stopping the signer from casting
// further votes (either for or against).
func (cliquer *Cliquer) Discard(address common.Address) error {
	var apis = cliquer.Eth.Engine().APIs(cliquer.Eth.BlockChain())
	var api = apis[0]

	if api.Namespace == "clique" {
		var cliqueApi = api.Service.(*clique.API)
		cliqueApi.Discard(address)
		return nil
	} else {
		var err1 = errors.New("Cliquer err")
		errors.Wrap(err1, "Discard str")
		return err1
	}
}
