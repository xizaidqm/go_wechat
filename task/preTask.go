package task

import (
	"encoding/json"
	"entity"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"proxy"
	"strconv"
	"strings"
	"sync"
	"time"
)

//PreTask任务用于获取 特定公众号 的所有文章链接
type PreTask struct {
	Name  string
	UrlContents  chan entity.ArticleInfo //需要爬取的链接
	Group *sync.WaitGroup
}

//实现Task接口
func (task *PreTask) Start() {
	//线程结束后终止waitGroup等待
	defer task.Group.Add(-1)

	//想要爬取的公众号名称
	name := "老韭菜生存日记"
	//count默认是5，改了也没有效果
	begin, count := 0, 5
	token := 1921878132

	targetAccount := *searchTargetAccount(name, begin, count, token)
	if &targetAccount == nil {
		log.Fatal("find no target account")
		return
	}
	//第一次获取前5篇内容
	msgResult := *searchScopeMsg(begin, count, token, targetAccount.FakeId, task.UrlContents)
	if &msgResult == nil {
		log.Fatal("find no msg")
	}
	log.Println("总计获取文章个数为：",msgResult.AppMsgCnt)

	scrapeCount := 5
	//爬取剩余的文章链接
	//TODO 目前是单机版本，后续需要改成并发
	for i := 5; i < msgResult.AppMsgCnt; i += 5 {
		searchScopeMsg(i, 5, token, targetAccount.FakeId, task.UrlContents)
		//随机休眠1s
		randTime := rand.Intn(5)
		time.Sleep( time.Second * time.Duration(randTime))

		//测试前30条
		if scrapeCount>5{
			break
		}
		scrapeCount+=1
	}
	//显示关闭channel输入
	defer close(task.UrlContents)

}

var msgIndex=0

func searchTargetAccount(name string, begin, count, token int) *entity.AccountInfo {

	var targetAccount entity.AccountInfo

	url := "https://mp.weixin.qq.com/cgi-bin/searchbiz?action=search_biz&begin=" + strconv.Itoa(begin) + "&count=" + strconv.Itoa(count) + "&query=" + name + "&token=" + strconv.Itoa(token) + "&lang=zh_CN&f=json&ajax=1"

	resp, err := proxy.VisitUrl(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("request url error", err)
		return nil
	}
	if resp == nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println("查询特定公众号链接：",url)

	var accountResult entity.AccountSearchResult
	err = json.Unmarshal(body, &accountResult)
	if err != nil {
		log.Fatalf("unmarshal failed: %v", err)
		return nil
	}

	//选取指定名称的公众号即可
	//TODO 后续需进行拓展，因目前的count只有5个，可能会小于total的数值
	for i := 0; i < accountResult.Total; i++ {
		if name == accountResult.List[i].Nickname {
			targetAccount.Nickname = name
			targetAccount.FakeId = accountResult.List[i].FakeId
			break
		}
	}
	return &targetAccount
}

func searchScopeMsg(begin, count, token int, fake_id string, ch chan entity.ArticleInfo) *entity.ArticleSearchResult {

	//搜索该公众号下的文章
	url := "https://mp.weixin.qq.com/cgi-bin/appmsg?action=list_ex&begin=" + strconv.Itoa(begin) + "&count=" + strconv.Itoa(count) + "&fakeid=" + fake_id + "&type=9&query=&token=" + strconv.Itoa(token) + "&lang=zh_CN&f=json&ajax=1"
	resp, err := proxy.VisitUrl(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("request url error", err)
		return nil
	}
	if resp == nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var articleSearchResult entity.ArticleSearchResult
	err = json.Unmarshal(body, &articleSearchResult)
	if err != nil {
		log.Fatalf("unmarshal failed: %v", err)
		return nil
	}

	//将文章链接输入channel，用于下一步爬取任务
	for _, msg := range articleSearchResult.AppMsgList {
		//若文章有标题，则截取换行符前的部分
		msgTitle := strings.Fields(msg.Title)[0]
		articleInfo :=  entity.ArticleInfo {MsgId: strconv.Itoa(msg.Appmsgid), Title: msgTitle, MsgLink:msg.Link, CreateTime: msg.CreateTime, UpdateTime: msg.UpdateTime}
		ch <- articleInfo
		//fmt.Println("第",msgIndex,"篇文章标题：",msg.Title,"；链接：",msg.Link)
		msgIndex+=1
	}

	return &articleSearchResult
}

//创建PreTask的工厂方法
func CreatePreTask(urlContents chan entity.ArticleInfo, group *sync.WaitGroup) Task {
	task := PreTask{
		Name:  "crawling msg links...",
		UrlContents:  urlContents,
		Group: group,
	}
	return &task
}
