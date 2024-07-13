package conf

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Conf struct {
	Port    string `yaml:"port"`
	Signer  string `yaml:"signer"`
	ChainID int64  `yaml:"chain_id"`

	StartBlock int64 `yaml:"start_block"`

	Rpc                string `yaml:"rpc"`
	Dns                string `yaml:"dns"`
	LogLevel           string `yaml:"log_level"`
	LogPath            string `yaml:"log_path"`
	AMMAddr            string `yaml:"amm_addr"`
	FetchBlockInterval int64  `yaml:"fetch_block_interval"`
	FetchInterval      int64  `yaml:"fetch_interval"` //second

	BasicTokenSymbol   string `yaml:"basic_token_symbol"`
	BasicTokenDecimals string `yaml:"basic_token_decimals"`

	TargetTokenSymbol   string `yaml:"target_token_symbol"`
	TargetTokenDecimals string `yaml:"target_token_decimals"`
}

func InitConf(confPath string) (*Conf, error) {
	var config Conf
	dataBytes, err := os.ReadFile(confPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		return nil, err
	}

	if config.Signer == "" {
		config.Signer = os.Getenv("PRIVATE_KEY")
	}

	return &config, nil
}
