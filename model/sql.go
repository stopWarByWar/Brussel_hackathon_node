package model

import (
	"gorm.io/gorm"
	"time"
)

func GetLatestExecBlock(db *gorm.DB) (int64, uint, error) {
	var index struct {
		Block int64
		Index uint
	}

	err := db.Table(Order{}.TableName()).Select("block, index").Order("block desc,index desc").Find(&index).Error
	return index.Block, index.Index, err
}

type Order struct {
	RequestID           string `gorm:"column:request_id; primary_key"`
	User                string `gorm:"column:user_addr"`
	BasicTokenIn        string
	TargetTokenIn       string
	SwapType            uint8
	ResultRate          int
	SwapTx              string
	PreBasicTokenOut    string
	PreTargetTokenOut   string
	FinalBasicTokenOut  string
	FinalTargetTokenOut string
	SettleTx            string
	Block               uint64
	Index               uint
	CreateTime          time.Time
}

func (Order) TableName() string {
	return "orders"
}
