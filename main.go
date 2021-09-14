package main

import (
	"entity"
	"fmt"
	"sync"
	"task"
	"time"
)

func main() {
	start := time.Now()
	//保障所有goroutine执行完毕
	var wg sync.WaitGroup
	contents := make(chan entity.ArticleInfo, 10)

	routines := make(chan struct{},10)

	preTask := task.CreatePreTask(contents, &wg)
	wg.Add(1)
	go preTask.Start()


	//spiderTask:= task.CreateSpiderTask(contents, &wg)
	//wg.Add(1)
	//go spiderTask.Start()

	downloadTask := task.CreateDownloadTask(contents, &wg, routines)
	wg.Add(1)
	go downloadTask.Start()

	wg.Wait()

	end := time.Now()

	fmt.Printf("整体耗时：%s",end.Sub(start).String())

	//a1 := entity.ArticleInfo{Title: "大晚上想起自己以前的炒股经历\n大一本来就巨穷，炒股还10000块亏了7000\n身上也没啥零花钱，也不好意思跟家里要钱\n就吃南大六食堂6块钱的铁板饭吃了3个月\n唉 一定要用闲钱炒股",
	//	MsgId: strconv.Itoa(2247486560), MsgLink: "http://mp.weixin.qq.com/s?__biz=MzA4NzYwNjMwNg==&mid=2247486560&idx=1&sn=9308346b9648ad269360852a40846958&chksm=9037992da740103bb59717a0d39166734284e3503d280f984b7893f9d7afa12a3418935d56e6#rd",
	//	CreateTime: 1626274014, UpdateTime: 1626274014}
	//a2 := entity.ArticleInfo{Title: "在导致泡沫破灭的原因切实出现之前，你怎么诅咒泡沫都是无用的，就像我们去年看不惯白酒酱油涨一样，怎么骂都没用，但是现在导致其泡沫破裂的原因已经出现了，也用不着我们诅咒它们了。当下的科技新能源，说没有泡沫肯定是忽悠人的，但是导致这个泡沫破裂的原因请大家想一下有什么，然后在这个原因出现之前，你再怎么看不惯科技新能源也是没用的。就算有一天它们跌了，也是底层支撑的逻辑变化了，而不是你唱空唱下来的。",
	//	MsgId: strconv.Itoa(2247486558), MsgLink: "http://mp.weixin.qq.com/s?__biz=MzA4NzYwNjMwNg==&mid=2247486558&idx=1&sn=12bcb9fe732e683c987097a87403738d&chksm=90379913a7401005cd3183ec5b239358547dc134d7f4f781618bc90b4ee62852aa91796b7eb6#rd",
	//	CreateTime: 1625737632, UpdateTime: 1625737632}
	//a3:=entity.ArticleInfo{Title: "肯定不是你唱空唱下来的。",
	//	MsgId: strconv.Itoa(2247486558), MsgLink: "https://mp.weixin.qq.com/s?__biz=MzA4NzYwNjMwNg==&mid=2247485322&idx=1&sn=5b5fa3440fcdc4ef0e0218c5a419652e&chksm=903792c7a7401bd1243daf4f14bff275bdbc2f42090ad8396d65d286dc58013ab8f4c13f040e#rd",
	//	CreateTime: 1525733215, UpdateTime: 1525733215}
	//contents <- a1
	//contents <- a2
	//contents<-a3
	//close(contents)
	//downloadTask := task.CreateDownloadTask(contents, &wg)
	//downloadTask.Start()



}
