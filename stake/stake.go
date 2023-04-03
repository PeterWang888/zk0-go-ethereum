package stake

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
)

var Staker *Stake = NewStake()

type Stake struct {
	cliquer *Cliquer
	initSub *InitiateChangeSub
	node    *node.Node
	eth     *eth.Ethereum
}

func NewStake() *Stake {
	return &Stake{}
}

func (stake *Stake) Init(node *node.Node, eth *eth.Ethereum) {
	cliquer := NewCliquer(node, eth)
	initSub := NewInitiateChangeSub()
	stake.cliquer = &cliquer
	stake.initSub = initSub
	stake.node = node
	stake.eth = eth

	log.Info(fmt.Sprintf("staker Init: %s", debug.Stack()))
}

func (stake *Stake) EventSub() {
	// pow add code

	var httpUrl string = "ws://127.0.0.1:" + strconv.Itoa(stake.node.Config().HTTPPort)
	var wsUrl string = "ws://127.0.0.1:" + strconv.Itoa(stake.node.Config().WSPort)
	log.Info("-----------------Init Signers Address: ---------------")
	initSigners, err := stake.cliquer.GetSigners()
	for _, signer := range initSigners {
		log.Info(signer.Hex())
	}

	if err != nil {
		log.Warn("cliquer GetSigners error: ", err)
		return
	}

	var ValidatorSetIsEnAble string
	if stake.node.Config().ValidatorSetIsEnAble {
		ValidatorSetIsEnAble = "true"
	} else {
		ValidatorSetIsEnAble = "false"
	}

	log.Info(fmt.Sprintf("ValidatorSetIsEnAble: %s", ValidatorSetIsEnAble))
	log.Info(fmt.Sprintf("ValidatorSetContractAddressFlag: %s", stake.node.Config().ValidatorSetContractAddress))
	log.Info(fmt.Sprintf("HTTP URL: %s", httpUrl))
	log.Info(fmt.Sprintf("WS URL: %s", wsUrl))

	//log.Info("-----------------Signers Address: ---------------")
	//for _, addr := range addrs {
	//	log.Info(addr.Hex())
	//}

	time.Sleep(time.Second * 2)

	if !stake.node.Config().ValidatorSetIsEnAble {
		return
	}

	err = stake.initSub.InitInitiateChangeSub(wsUrl,
		stake.node.Config().ValidatorSetContractAddress,
	)
	if err != nil {
		log.Error(fmt.Sprintf("New InitInitiateChangeSub error: ", err))
		return
	}

	log.Info("-----------------Init InitiateChangeSub---------------")

	for {
		vargs, err := stake.initSub.SubInitiateChange()
		if err != nil {
			log.Warn("Sub InitiateChange error: ", err)
			return
		}
		log.Info("-----------------InitiateChange begin-----------------")
		log.Info("-----------------Signers Address: ---------------")
		signers, err := stake.cliquer.GetSigners()
		signers = CalcDiffAddr(signers, initSigners)
		for _, signer := range signers {
			log.Info(signer.Hex())
		}

		log.Info("-----------------Validator Address: ---------------")
		var validators = vargs[0].([]common.Address)
		for _, validator := range validators {
			log.Info(validator.Hex())
		}
		proposes := CalcDiffAddr(validators, signers)
		log.Info("-----------------Proposes Address: ---------------")
		for _, propose := range proposes {
			log.Info(propose.Hex())
			stake.cliquer.Propose(propose, true)
		}

		log.Info("-----------------Discards Address: ---------------")
		discards := CalcDiffAddr(signers, validators)
		for _, discard := range discards {
			log.Info(discard.Hex())
			stake.cliquer.Discard(discard)
		}

		log.Info("-----------------InitiateChange end-----------------")
	}
}
