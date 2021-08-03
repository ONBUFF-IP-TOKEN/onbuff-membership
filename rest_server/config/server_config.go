package config

import (
	"sync"

	baseconf "github.com/ONBUFF-IP-TOKEN/baseapp/config"
)

var once sync.Once
var currentConfig *ServerConfig

type IpblockServer struct {
	ApplicationName string `json:"application_name" yaml:"application_name"`
	APIDocs         bool   `json:"api_docs" yaml:"api_docs"`
}

type TokenInfo struct {
	MainnetHost      string   `yaml:"mainnet_host"`
	ServerWalletAddr string   `yaml:"server_wallet_address"`
	ServerPrivateKey string   `yaml:"server_private_key"`
	TokenAddrs       []string `yaml:"token_address"`
	NftUriDomain     string   `yaml:"nft_uri_domain"`
}

type ApiAuth struct {
	AuthEnable        bool   `yaml:"auth_enable"`
	JwtSecretKey      string `yaml:"jwt_secret_key"`
	TokenExpiryPeriod int64  `yaml:"token_expiry_period"`
	SignExpiryPeriod  int64  `yaml:"sign_expiry_period"`
	AesKey            string `yaml:"aes_key"`
}

type ServerConfig struct {
	baseconf.Config `yaml:",inline"`

	IPServer    IpblockServer   `yaml:"ipblock_server"`
	MysqlDBAuth baseconf.DBAuth `yaml:"mysql_db_auth"`
	Token       TokenInfo       `yaml:"token_info"`
	Auth        ApiAuth         `yaml:"api_auth"`
}

func GetInstance(filepath ...string) *ServerConfig {
	once.Do(func() {
		if len(filepath) <= 0 {
			panic(baseconf.ErrInitConfigFailed)
		}
		currentConfig = &ServerConfig{}
		if err := baseconf.Load(filepath[0], currentConfig); err != nil {
			currentConfig = nil
		}
	})

	return currentConfig
}
