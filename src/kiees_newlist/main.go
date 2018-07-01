/*
小爬虫
作用: 定时获取当前 发现值得买 页面上数码分区新出现可能会感兴趣的内容并且通过邮件把相关连接发送到邮箱中.
特殊: 添加黑名单机制,读取配置文件中预设的黑名单,使用正则筛选出真正感兴趣的内容而不是每日重复的那些.
作者: 死猪panndora(胖哆啦)
*/
package main

import (
	"co"
	"config"
	"log"
	"os"
	"regexp"
	"sendmail"
	"strings"
)

// 处理获取到的链接,根据黑名单返回需要的url地址列表.
func dataProcessing(hs []string, baklist []string) (c []string) {
	regpostbox, _ := regexp.Compile("<div class=\"postbox\">[\\s\\S]+?\"clear:both\"></div></div>") // 框架
	regalt, _ := regexp.Compile("alt=\"(.*?)\"")                                                    // 标题
	regurl, _ := regexp.Compile("/\\d{4}/\\d{1,2}/\\d{1,2}/\\d+\\.html")                            // 网址
	for _, v := range hs {
		for _, m := range regpostbox.FindAllString(v, -1) {
			t := regalt.FindAllStringSubmatch(m, -1)
			k := true
			// 查找当前标题中是否包含关键字
			for _, w := range baklist {
				if strings.Contains(t[0][1], w) {
					k = false
				}
			}
			if k {
				// 保存为超链接,可直接写入邮件
				c = append(c, "<a href=\"http://www.kiees.com"+regurl.FindString(m)+"\" target=\"_black\">"+t[0][1]+"</a>")
			}
		}
	}
	return c
}

// 发送邮件返回成功与否
func sendEmail(l string) bool {
	if err := sendmail.Send("overpan@qq.com", "********", "panndora@deadpig.cc", l); err != nil {
		return false
	}
	return true
}

// 读取配置文件黑名单返回列表
func readConfigFile() []string {
	return config.LoadConfig()
}

func main() {
	logFile, _ := os.Create("kiees.log")
	defer logFile.Close()
	runLog := log.New(logFile, "[优惠卷过滤器运行日志]", log.LstdFlags)
	Couponlist := dataProcessing(co.GetHtmls(), readConfigFile())
	mailmsg := strings.Join(Couponlist, "<br>")
	if sendEmail(mailmsg) {
		runLog.Println("邮件发送成功")
	} else {
		runLog.Println("邮件发送失败")
	}
	return
}
