package processor

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"sse_brussels_node/conf"
	"sse_brussels_node/logger"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	_conf, err := conf.InitConf("../conf/conf.yaml")
	if err != nil {
		panic(err)
	}

	_db, err := gorm.Open(postgres.Open(_conf.Dns), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		panic(err)
	}

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

	p := New(_db, _logger, ethCli, _conf.AMMAddr, _conf.FetchBlockInterval, _conf.Signer, _conf.ChainID, _conf.StartBlock)
	p.Process(time.Duration(_conf.FetchInterval * 1_000_000_000))
}
