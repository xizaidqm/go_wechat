package task

import (
	"bufio"
	"entity"
	"fmt"
	"os"
	"sync"
)

type SpiderTask struct {
	name string	//任务名称
	articleInfos chan entity.ArticleInfo	//需要爬取的文章链接
	Group *sync.WaitGroup
}

func (task *SpiderTask) Start() {
	defer task.Group.Add(-1)
	index := 0
	filePath := "E:/go_Projects/wechat_spider/links.txt"
	file,err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE,0666)
	if err!= nil{
		fmt.Println("open file failed",err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)

	for {
		url,ok := <-task.articleInfos
		//写入文件
		fmt.Fprintf(write,"第%d个链接：%s\n",index,url)
		index++
		if !ok{
			break
		}
	}
	write.Flush()
}

func CreateSpiderTask(articleInfos chan entity.ArticleInfo, group *sync.WaitGroup) Task{
	task :=  SpiderTask{
		name:"default spider task",
		articleInfos : articleInfos,
		Group: group,
	}
	return &task
}