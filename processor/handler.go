package processor

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math"
	"math/big"
	"sse_brussels_node/model"
	"time"
)

func (p *Processor) dealWithGetRandomEvent(e types.Log) error {
	event, err := p.amm.ParseGetRandoms(e)
	if err != nil {
		return errors.New(fmt.Sprintf("fail to unpack getRandom event with error %v", err))
	}

	var percentage float64
	uniRand := unifyRandom(event.Randomness)
	//option swap
	if event.SellType == 0 {
		if int64(uniRand*10000) < p.winOptionalProbability {
			//win optional type swap
			percentage = getRandomAccordingDistribution(uniRand, float64(2*p.optionalDownBound)/10_000, float64(2*p.optionalUpBound)/10_000)
		} else {
			percentage = getRandomAccordingDistribution(uniRand, float64(p.optionalDownBound)/10_000, float64(p.optionalUpBound)/10_000)
		}
		//range order
	} else if event.SellType == 1 {
		percentage = getRandomAccordingDistribution(uniRand, 1-float64(p.marketSwapRate)/10_000, 1+float64(p.marketSwapRate)/10_000)
	}

	signedTx, err := p.amm.Settle(p.txOp, event.RequestId, big.NewInt(int64(math.Round(percentage*10_000))))
	if err != nil {
		return err
	}

	tx, err := waitTx(p.ethCli, signedTx)
	if err != nil {
		return err
	}

	p.logger.WithFields(logrus.Fields{
		"tx":     tx,
		"req id": event.RequestId.String(),
	}).Info("settle order")

	Order := model.Order{
		RequestID:         event.RequestId.String(),
		User:              event.User.String(),
		BasicTokenIn:      event.BasicToken.String(),
		TargetTokenIn:     event.TargetToken.String(),
		SwapType:          event.SellType,
		ResultRate:        int(math.Round(percentage * 10_000)),
		SwapTx:            event.Raw.TxHash.String(),
		PreBasicTokenOut:  "0",
		PreTargetTokenOut: "0",
		SettleTx:          "",
		Block:             event.Raw.BlockNumber,
		Index:             event.Raw.Index,
		CreateTime:        time.Now(),
	}
	return p.db.Save(&Order).Error
}

func (p *Processor) dealWithSwapSuccessfulEvent(e types.Log) error {
	event, err := p.amm.ParseSwapSuceesfully(e)
	if err != nil {
		return errors.New(fmt.Sprintf("fail to unpack swapSuccessful event with error %v", err))
	}

	var order *model.Order

	if err = p.db.Where("request_id = ?", event.RequestId).First(&order).Error; err != nil {
		return err
	}

	order.Block = event.Raw.BlockNumber
	order.Index = event.Raw.TxIndex

	// basic in > 0 => target out > 0
	if order.BasicTokenIn == big.NewInt(0).String() {
		order.PreBasicTokenOut = event.PreOutTokenAmount.String()
		order.FinalBasicTokenOut = event.FinalOutTokenAmount.String()

		// target in > 0 => basic out > 0
	} else if order.TargetTokenIn == big.NewInt(0).String() {
		order.PreTargetTokenOut = event.PreOutTokenAmount.String()
		order.FinalTargetTokenOut = event.FinalOutTokenAmount.String()

	}
	order.SettleTx = event.Raw.TxHash.String()
	return p.db.Save(&order).Error
}

func waitTx(conn *ethclient.Client, tx *types.Transaction) (string, error) {
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		return "", err
	}
	return receipt.TxHash.String(), nil
}
