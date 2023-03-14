package stake

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
)

func InitiateChangeSubTest() {
	initSub := NewInitiateChangeSub()

	err := initSub.InitInitiateChangeSub("ws://127.0.0.1:8545",
		"0xfB45f8B9705FFA9d05D42b72E73afE7E2ceEa860",
	)
	if err != nil {
		log.Crit("New InitInitiateChangeSub error", err)
	}

	for {
		vargs, err := initSub.SubInitiateChange()
		if err != nil {
			log.Warn("SubInitiateChange error: ", err)
		}
		log.Info(fmt.Sprintf("InitiateChange: "))
		var args = vargs[0].([]common.Address)
		for _, arg := range args {
			log.Info(fmt.Sprintf(arg.Hex()))
		}
	}
}

func CalcDiffAddr(addrs1 []common.Address, addrs2 []common.Address) []common.Address {
	resAddrs := []common.Address{}

	for _, addr1 := range addrs1 {
		isExist := false
		for _, addr2 := range addrs2 {
			if addr1 == addr2 {
				isExist = true
			}
		}
		if !isExist {
			resAddrs = append(resAddrs, addr1)
		}
	}

	return resAddrs
}

func NewInitiateChangeSub() *InitiateChangeSub {
	return &InitiateChangeSub{}
}

type InitiateChangeSub struct {
	cli       *ethclient.Client
	parsedAbi *abi.ABI
	query     *ethereum.FilterQuery
	sub       *ethereum.Subscription
	vlogs     *chan types.Log
}

func (is *InitiateChangeSub) InitInitiateChangeSub(wsUrl string, contrAddr string) error {
	err := is.InitEthCli(wsUrl)
	if err != nil {
		return err
	}

	err = is.GenAbi()
	if err != nil {
		return err
	}

	err = is.GenSubFilter(contrAddr)
	if err != nil {
		return err
	}

	return nil
}

func (is *InitiateChangeSub) InitEthCli(url string) error {
	cli, err := ethclient.Dial(url)
	if err != nil {
		errors.Wrap(err, "InitEthCli str")
		return err
	}

	is.cli = cli
	return nil
}

func (is *InitiateChangeSub) GenAbi() error {
	parsedAbi, err := abi.JSON(strings.NewReader(ValidatorsetABI))
	if err != nil {
		errors.Wrap(err, "genAbi2 str")
		return err
	}

	is.parsedAbi = &parsedAbi
	return nil
}

// create filter
func (is *InitiateChangeSub) GenFilter(contrAddr string) {
	var topic = crypto.Keccak256Hash([]byte("InitiateChange(bytes32,address[])"))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contrAddr)},
		Topics: [][]common.Hash{
			{
				common.Hash(topic),
			},
		},
	}

	is.query = &query
}

func (is *InitiateChangeSub) GenSub() error {
	vlogs := make(chan types.Log)
	sub, err := is.cli.SubscribeFilterLogs(context.Background(), *is.query, vlogs)

	if err != nil {
		errors.Wrap(err, "genSub str")
		return err
	}

	is.sub = &sub
	is.vlogs = &vlogs
	return nil
}

func (is *InitiateChangeSub) GenSubFilter(contrAddr string) error {
	is.GenFilter(contrAddr)

	return is.GenSub()
}

// handle subscription events logs
func (is *InitiateChangeSub) SubInitiateChange() ([]interface{}, error) {
	// select阻塞监控多个channel
	// 任意一个channel有消息则解除阻塞，并且执行case内的channel
	select {
	case err := <-(*is.sub).Err():
		errors.Wrap(err, "subEvent1 str")
		return nil, err
	case vlog := <-*is.vlogs:
		//data, err := vlog.MarshalJSON()
		vargs, err := (*is.parsedAbi).Events["InitiateChange"].Inputs.UnpackValues(vlog.Data)
		if err != nil {
			errors.Wrap(err, "subEvent2 str")
			return nil, err
		}
		return vargs, nil
	}
}
