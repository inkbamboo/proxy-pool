package main

import (
	"github.com/inkbamboo/proxy-pool/config"
	"github.com/inkbamboo/proxy-pool/internal/services"
	"time"
)

func main() {
	config.InitConfigWithPath()
	for _, item := range config.GetBaseConfig().Spider {
		services.GetSpiderService().Start(item)
	}
	time.Sleep(10 * time.Second)
	//http := utils.VerifyHttp("50.174.145.10:80")
	//fmt.Printf("https%v\n", http)
	//https := utils.VerifyHttps("50.174.145.10:80")
	//fmt.Printf("https%v\n", https)

}
