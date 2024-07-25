package config

import (
	"bytes"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/viper"
	"net/http"
)

var (
	v  *viper.Viper
	bc *BaseConfig
)

type BaseConfig struct {
	Spider []*Spider `json:"spider"`
	Proxy  *Proxy    `json:"proxy"`
	Config *Config   `json:"config"`
}
type Config struct {
	Ip               string `json:"ip"`
	Port             string `json:"port"`
	HttpTunnelPort   string `json:"httpTunnelPort"`
	SocketTunnelPort string `json:"socketTunnelPort"`
	TunnelTime       int    `json:"tunnelTime"`
	ProxyNum         int    `json:"proxyNum"`
	VerifyTime       int    `json:"verifyTime"`
	ThreadNum        int    `json:"threadNum"`
}
type Spider struct {
	Name  string   `json:"name"`
	Proxy bool     `json:"proxy"`
	Urls  []string `json:"urls"`
	XPath string   `json:"xpath"`
}
type Proxy struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
type ProxyIp struct {
	Ip         string //IP地址
	Port       string //代理端口
	Country    string //代理国家
	Province   string //代理省份
	City       string //代理城市
	Isp        string //IP提供商
	Type       string //代理类型
	Anonymity  string //代理匿名度, 透明：显示真实IP, 普匿：显示假的IP, 高匿：无代理IP特征
	Time       string //代理验证
	Speed      string //代理响应速度
	SuccessNum int    //验证请求成功的次数
	RequestNum int    //验证请求的次数
	Source     string //代理源
}

func SetHeadersConfig(he map[string]string, header *http.Header) *http.Header {
	for k, v := range he {
		header.Add(k, v)
	}
	return header
}

func InitConfigWithPath() {
	v = viper.New()
	configName := "config.yaml"
	v.SetConfigName(configName)
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	cfg := packr.New("config", "./")
	data, err := cfg.Find(configName)
	if err != nil {
		panic(err)
	}
	viper.ReadConfig(bytes.NewBuffer(data))
	if err = v.ReadInConfig(); err != nil {
		fmt.Println(fmt.Sprintf("Viper ReadInConfig err:%s\n", err))
		panic(err)
	}
	err = v.Unmarshal(&bc)
	if err != nil {
		fmt.Println("yaml parse err: ", err)
		panic(err)
	}
}
func GetConfig() *viper.Viper {
	if v == nil {
		panic("Please init Config")
	}
	return v
}
func GetBaseConfig() *BaseConfig {
	if bc == nil {
		panic("Please init Config")
	}
	return bc
}
