package main


import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
)


func main() {

	path := chromedp.ExecPath("C:\\Program Files\\Mozilla Firefox\\firefox.exe")
	fmt.Println(path)
	// 创建新的cdp上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()


	// 此处以360搜索首页为例
	urlstr := `https://home.firefoxchina.cn/`
	var buf []byte
	// 需要截图的元素，支持CSS selector以及XPath query
	selector := `#main`
	if err := chromedp.Run(ctx, elementScreenshot(urlstr, selector, &buf)); err != nil {
		log.Fatal("run error:",err)
	}
	// 写入文件
	if err := ioutil.WriteFile("360_so.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}


// 截图方法
func elementScreenshot(urlstr, sel string, res *[]byte)chromedp.Tasks {
	return chromedp.Tasks{
		// 打开url指向的页面
		chromedp.Navigate(urlstr),


		// 等待待截图的元素渲染完成
		chromedp.WaitVisible(sel, chromedp.ByID),
		// 也可以等待一定的时间
		//cdp.Sleep(time.Duration(3) * time.Second),


		// 执行截图
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
	}
}
