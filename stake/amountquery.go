package stake

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"strings"
	"time"

	"github.com/pkg/errors"
	"math/big"
)

func AmountQueryTest() {
	amtqry := NewIAmountQuery()

	err := amtqry.InitAmountQuery("http://127.0.0.1:8545",
		"0xfB45f8B9705FFA9d05D42b72E73afE7E2ceEa860",
	)
	if err != nil {
		log.Crit("New AmountQuery error", err)
	}

	for {
		amounts, stakers, err := amtqry.GetStakerAmounts()
		if err == nil {
			log.Warn("GetStakerAmounts error: ", err)
		}
		for idx, val := range stakers {
			log.Info(fmt.Sprintf("stake: ", idx, val))
		}
		for idx, val := range amounts {
			log.Info(fmt.Sprintf("amount: ", idx, val))
		}

		amounts, validators, err := amtqry.GetValidatorAmounts()
		if err == nil {
			log.Warn("GetValidatorAmounts error: ", err)
		}
		for idx, val := range validators {
			log.Info(fmt.Sprintf("stake: ", idx, val))
		}
		for idx, val := range amounts {
			log.Info(fmt.Sprintf("amount: ", idx, val))
		}

		time.Sleep(time.Second)
	}
}

func NewIAmountQuery() *AmountQuery {
	return &AmountQuery{}
}

type AmountQuery struct {
	cli          *ethclient.Client
	parsedAbi    *abi.ABI
	validatorSet *Validatorset
}

func (aq *AmountQuery) InitAmountQuery(httpUrl string, contrAddr string) error {
	err := aq.InitEthCli(httpUrl)
	if err != nil {
		return err
	}

	err = aq.GenAbi()
	if err != nil {
		return err
	}

	err = aq.NewValidatorset(contrAddr)
	if err != nil {
		return err
	}

	return nil
}

func (aq *AmountQuery) InitEthCli(url string) error {
	cli, err := ethclient.Dial(url)
	if err != nil {
		errors.Wrap(err, "InitEthCli str")
		return err
	}

	aq.cli = cli
	return nil
}

func (aq *AmountQuery) GenAbi() error {
	parsedAbi, err := abi.JSON(strings.NewReader(ValidatorsetABI))
	if err != nil {
		errors.Wrap(err, "genAbi2 str")
		return err
	}

	aq.parsedAbi = &parsedAbi
	return nil
}

func (aq *AmountQuery) NewValidatorset(contrAddr string) error {
	validatorSet, err := NewValidatorset(common.HexToAddress(contrAddr), aq.cli)
	if err != nil {
		errors.Wrap(err, "New Validatorset str")
		return err
	}
	aq.validatorSet = validatorSet
	return nil
}

func (aq *AmountQuery) GetStakerAmounts() ([]common.Address, []*big.Int, error) {
	stakers, amouts, err := aq.validatorSet.GetStakerAmounts(
		&bind.CallOpts{
			Pending:     false,
			From:        common.Address{},
			BlockNumber: nil,
			Context:     nil})

	if err != nil {
		errors.Wrap(err, "Get Staker Amounts str")
		return stakers, amouts, err
	}

	return stakers, amouts, err
}

func (aq *AmountQuery) GetValidatorAmounts() ([]common.Address, []*big.Int, error) {
	stakers, amouts, err := aq.validatorSet.GetValidatorAmounts(
		&bind.CallOpts{
			Pending:     false,
			From:        common.Address{},
			BlockNumber: nil,
			Context:     nil})

	if err != nil {
		errors.Wrap(err, "Get Validator Amounts str")
		return stakers, amouts, err
	}

	return stakers, amouts, err
}
