package services

import (
	"fmt"
	"github.com/EDDYCJY/fake-useragent"
	"github.com/gocolly/colly/v2"
	"github.com/inkbamboo/proxy-pool/config"
	"github.com/inkbamboo/proxy-pool/internal/utils"
	"regexp"
	"strings"
	"sync"
)

var (
	SpiderService     *Spider
	SpiderServiceOnce sync.Once
)

func GetSpiderService() *Spider {
	SpiderServiceOnce.Do(func() {
		SpiderService = &Spider{}
	})
	return SpiderService
}

type Spider struct {
}

func (s *Spider) Start(cfg *config.Spider) {
	fmt.Println("Start Spider")
	c := colly.NewCollector(
		colly.AllowURLRevisit(),        //允许对同一 URL 进行多次下载
		colly.Async(true),              //设置为异步请求
		colly.MaxDepth(2),              //爬取页面深度,最多为两层
		colly.MaxBodySize(2*1024*1024), //响应正文最大字节数
		colly.IgnoreRobotsTxt(),        //忽略目标机器中的`robots.txt`声明
	)
	if cfg.Proxy && config.GetBaseConfig().Proxy != nil {
		proxy := config.GetBaseConfig().Proxy
		c.SetProxy("http://" + proxy.Host + ":" + proxy.Port)
	}
	c.OnXML(cfg.XPath, func(e *colly.XMLElement) {
		txt := strings.Join(e.ChildTexts("*"), ":")
		ipReg, _ := regexp.Compile(`(\d+?\.\d+?.\d+?\.\d+?):\d+`)
		ips := ipReg.FindAllString(txt, -1)
		for _, ip := range ips {
			if utils.VerifyHttp(ip) {
				fmt.Printf("*****proxy: %s   http//%s\n", cfg.Name, ip)
			} else if utils.VerifyHttps(ip) {
				fmt.Printf("*****proxy: %s   https//%s\n", cfg.Name, ip)
			}
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("proxy: %s Visiting: %v\n", cfg.Name, r.URL)
	})
	for _, url := range cfg.Urls {
		c.UserAgent = browser.Random()
		c.Visit(url)
	}
}
