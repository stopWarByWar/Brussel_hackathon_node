package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/big"
	"net/http"
	"sse_brussels_node/model"
	"strconv"
)

type Server struct {
	DB                  *gorm.DB
	BasicTokenSymbol    string
	BasicTokenDecimals  *big.Float
	TargetTokenSymbol   string
	TargetTokenDecimals *big.Float
}

func (s *Server) GetOrderList(c *gin.Context) {
	_page := c.Query("page")
	_pageSize := c.Query("page_size")

	page, err := strconv.Atoi(_page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page should be number"})
		return
	}

	pageSize, err := strconv.Atoi(_pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page_size should be number"})
		return
	}

	var orders []model.Order

	if err := s.DB.Where("settle_tx != ?", "").Order("block, index").Offset((page - 1) * pageSize).Limit(pageSize).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fail to get order list"})
		return
	}

	var results []*OrderDetail

	for _, order := range orders {
		res, err := s.dealWithOrderToResp(order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, res)
	}

	c.JSON(http.StatusOK, gin.H{
		"txs": results,
	})
}

func (s *Server) GetTxStatus(c *gin.Context) {
	txHash := c.Query("tx")

	var order model.Order

	if err := s.DB.Where("swap_tx = ?", txHash).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Not find swap transaction hash %s", txHash)})
		return
	}

	res, err := s.dealWithOrderToResp(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order": res,
	})
}

func (s *Server) dealWithOrderToResp(order model.Order) (*OrderDetail, error) {
	var res = &OrderDetail{
		User:     order.User,
		SwapTx:   order.SwapTx,
		SettleTx: order.SettleTx,
		SwapType: order.SwapType,
	}

	if order.BasicTokenIn != "0" {
		res.TokenIn = s.BasicTokenSymbol
		res.TokenOut = s.TargetTokenSymbol

		_tokenInAmount, _ := big.NewFloat(0).SetString(order.BasicTokenIn)
		res.TokenAmountIn, _ = _tokenInAmount.Quo(_tokenInAmount, s.BasicTokenDecimals).Float64()

		_tokenOutAmountExp, _ := big.NewFloat(0).SetString(order.PreTargetTokenOut)
		_tokenOutAmountExp.Quo(_tokenOutAmountExp, s.TargetTokenDecimals)
		res.TokenOutExpectation, _ = _tokenOutAmountExp.Float64()

		_tokenOutFinal, _ := big.NewFloat(0).SetString(order.FinalTargetTokenOut)
		_tokenOutFinal.Quo(_tokenOutFinal, s.TargetTokenDecimals)
		res.TokenOutFinal, _ = _tokenOutFinal.Float64()
	} else if order.TargetTokenIn != "0" {
		res.TokenIn = s.TargetTokenSymbol
		res.TokenOut = s.BasicTokenSymbol

		_tokenInAmount, _ := big.NewFloat(0).SetString(order.TargetTokenIn)
		res.TokenAmountIn, _ = _tokenInAmount.Quo(_tokenInAmount, s.TargetTokenDecimals).Float64()

		_tokenOutAmountExp, _ := big.NewFloat(0).SetString(order.PreBasicTokenOut)
		_tokenOutAmountExp.Quo(_tokenOutAmountExp, s.BasicTokenDecimals)
		res.TokenOutExpectation, _ = _tokenOutAmountExp.Float64()

		_tokenOutFinal, _ := big.NewFloat(0).SetString(order.FinalBasicTokenOut)
		_tokenOutFinal.Quo(_tokenOutFinal, s.BasicTokenDecimals)
		res.TokenOutFinal, _ = _tokenOutFinal.Float64()
	} else {
		return nil, errors.New(fmt.Sprintf("invalid swap order: swap tx %s", order.SwapTx))
	}

	if order.SettleTx != "" {
		res.Status = 0
	}

	res.Rate = (res.TokenOutFinal - res.TokenOutExpectation) / res.TokenOutExpectation
	return res, nil
}

type OrderDetail struct {
	User     string
	TokenIn  string
	TokenOut string

	SwapType            uint8
	SwapTx              string
	SettleTx            string
	TokenAmountIn       float64
	TokenOutExpectation float64
	TokenOutFinal       float64
	Status              uint8
	Rate                float64
}
