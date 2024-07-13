package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"math/big"
	"sse_brussels_node/conf"
	"sse_brussels_node/logger"
	"sse_brussels_node/processor"
	"sse_brussels_node/router"
	"sse_brussels_node/server"
	"time"
)

var configFile = flag.String("f", "./conf/conf.yaml", "the config file")

func main() {
	flag.Parse()
	_conf, err := conf.InitConf(*configFile)
	if err != nil {
		panic(err)
	}

	_db, err := gorm.Open(postgres.Open(_conf.Dns), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		panic(err)
	}
	//go startProcess(_conf, _db)

	s := new(server.Server)
	s.DB = _db
	s.BasicTokenSymbol = _conf.BasicTokenSymbol
	s.BasicTokenDecimals, _ = big.NewFloat(0).SetString(_conf.BasicTokenDecimals)
	s.TargetTokenSymbol = _conf.TargetTokenSymbol
	s.TargetTokenDecimals, _ = big.NewFloat(0).SetString(_conf.TargetTokenDecimals)

	r := router.SetUpRouter(s)
	r.Run(fmt.Sprintf(":%s", _conf.Port))
}

func startProcess(_conf *conf.Conf, _db *gorm.DB) {

	var lv logrus.Level

	switch _conf.LogLevel {
	case "info":
		lv = logrus.InfoLevel
	case "error":
		lv = logrus.ErrorLevel
	default:
		lv = logrus.DebugLevel
	}

	logConf := logger.Config{
		Level:        lv,
		ReportCaller: true,
		FilePath:     "log/sse_node",
	}

	_logger, err := logger.New(&logConf)
	if err != nil {
		panic(err)
	}

	ethCli, err := ethclient.Dial(_conf.Rpc)
	if err != nil {
		panic(err)
	}

	p := processor.New(_db, _logger, ethCli, _conf.AMMAddr, _conf.FetchBlockInterval, _conf.Signer, _conf.ChainID, _conf.StartBlock)
	p.Process(time.Duration(_conf.FetchInterval * 1_000_000_000))
}
