package processor

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/big"
	"sse_brussels_node/abi"
	"sse_brussels_node/logger"
	"sse_brussels_node/model"
	"time"
)

type Processor struct {
	db     *gorm.DB
	ethCli *ethclient.Client
	logger *logger.Logger

	ammAddr common.Address
	amm     *abi.AMM
	txOp    *bind.TransactOpts

	fetchBlockInterval     int64
	winOptionalProbability int64
	optionalDownBound      int64
	optionalUpBound        int64
	marketSwapRate         int64

	block int64
	index uint
}

var ctx = context.Background()

func New(_db *gorm.DB, _logger *logger.Logger, _ethCli *ethclient.Client, _addAddr string, _fetchBlockInterval int64, priKey string, chainId int64, startBlock int64) *Processor {
	block, index, err := model.GetLatestExecBlock(_db)
	if err != nil {
		panic(err)
	}
	if block == 0 {
		block = startBlock
	}

	_privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		panic(err)
	}

	_amm, err := abi.NewAMM(common.HexToAddress(_addAddr), _ethCli)
	if err != nil {
		panic(err)
	}

	_winOptionalProbability, _optionalDownBound, _optionalUpBound, _marketSwapRate, err := _amm.GetConfInfo(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	_txOp, err := bind.NewKeyedTransactorWithChainID(_privateKey, big.NewInt(chainId))
	if err != nil {
		panic(err)
	}

	return &Processor{
		db:                     _db,
		logger:                 _logger,
		ethCli:                 _ethCli,
		ammAddr:                common.HexToAddress(_addAddr),
		fetchBlockInterval:     _fetchBlockInterval,
		block:                  block,
		index:                  index,
		amm:                    _amm,
		winOptionalProbability: _winOptionalProbability.Int64(),
		optionalDownBound:      _optionalDownBound.Int64(),
		optionalUpBound:        _optionalUpBound.Int64(),
		marketSwapRate:         _marketSwapRate.Int64(),
		txOp:                   _txOp,
	}
}

func (p *Processor) Process(updateInterval time.Duration) {
	p.logger.WithFields(logrus.Fields{
		"block": p.block,
		"index": p.index,
	}).Info("start run processor")

	ticker := time.NewTicker(1)

	for {
		select {
		case <-ticker.C:
			ticker.Stop()
			currentBlock, err := p.ethCli.BlockNumber(ctx)
			if err != nil {
				p.logger.WithFields(logrus.Fields{
					"error": err,
				}).Error("fail to get current block num")
				time.Sleep(5 * time.Second)
				ticker.Reset(updateInterval)
			}

			if p.block < int64(currentBlock) {
				p.process(int64(currentBlock))
			}
			ticker.Reset(updateInterval)
		}
	}
}

func (p *Processor) process(currentBlockNum int64) {
	fromBlock := p.block
loop:
	for {
		toBlock := fromBlock + p.fetchBlockInterval
		rawEvents, err := p.ethCli.FilterLogs(ctx, ethereum.FilterQuery{
			FromBlock: big.NewInt(fromBlock),
			ToBlock:   big.NewInt(toBlock),
			Addresses: []common.Address{p.ammAddr},
			Topics:    [][]common.Hash{{GetRandomTopic, AmmSwapSuccessTopic}},
		})
		if err != nil {
			p.logger.WithFields(logrus.Fields{
				"from":  fromBlock,
				"to":    toBlock,
				"error": err,
			}).Error("fail to get amm contract event from chain")
			time.Sleep(10 * time.Second)
			continue
		}

		p.logger.WithFields(logrus.Fields{
			"from":        fromBlock,
			"to":          toBlock,
			"total event": len(rawEvents),
		}).Info()

		for _, e := range rawEvents {
			if uint64(p.block) > e.BlockNumber {
				continue
			} else if uint64(p.block) == e.BlockNumber && p.index >= e.Index {
				continue
			}
			if err = p.dealWithEvent(e); err != nil {
				p.logger.WithFields(logrus.Fields{
					"error": err,
					"block": e.BlockNumber,
					"index": e.Index,
				}).Error("fail to deal with event")
				time.Sleep(10 * time.Second)
				fromBlock = int64(e.BlockNumber)
				continue loop
			}
		}

		if toBlock < currentBlockNum {
			fromBlock = toBlock
			continue
		} else {
			p.block = currentBlockNum
			break
		}

	}
	p.block = currentBlockNum
	p.index = 0
	p.logger.WithFields(logrus.Fields{
		"exec block number": p.block,
		"exec index":        p.index,
		"current block":     currentBlockNum,
	}).Debug("finish one fetch period")
}

func (p *Processor) dealWithEvent(e types.Log) error {
	switch e.Topics[0] {
	case GetRandomTopic:
		if err := p.dealWithGetRandomEvent(e); err != nil {
			return err
		}
	case AmmSwapSuccessTopic:
		if err := p.dealWithSwapSuccessfulEvent(e); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknow event topic:%v", e.Topics[0].String())
	}
	p.block = int64(e.BlockNumber)
	p.index = e.Index
	return nil
}
