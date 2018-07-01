package co

import (
	"github.com/yudeguang/gather"
)

// 打开发现值得买的数码商品页面,并且返回前3页内容
func GetHtmls() (htmls []string) {
	gat := gather.NewGather("chrome", false)
	urls := []string{
		"http://www.kiees.com/fenlei/shangpinleibie/shoujishuma/",
		"http://www.kiees.com/fenlei/shangpinleibie/shoujishuma/index2.html",
		"http://www.kiees.com/fenlei/shangpinleibie/shoujishuma/index3.html",
	}
	// 因为只需要读取三个页面的内容,所以没有必要使用多线程等方法
	for k, v := range urls {
		html, _, _ := gat.Get(v, urls[k])
		htmls = append(htmls, html)
	}
	return
}
