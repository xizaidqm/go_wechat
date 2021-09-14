package task

import (
	"entity"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/petermattis/goid"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"proxy"
	"strings"
	"sync"
	"time"
)

type DownloadTask struct {
	Name string
	articleInfos chan entity.ArticleInfo
	group *sync.WaitGroup
	routines chan struct{}
}

var wkhtmltoxPath, fileSavePath string

func (task *DownloadTask) Start() {

	defer task.group.Add(-1)

	pdfGenerator := SaveAsPDF()

	wkhtmltoxPath = "E:\\wkhtmltox\\bin\\wkhtmltopdf.exe"
	fileSavePath = "E:\\go_Projects\\wechat_files\\zbx"

	for {
		content,ok := <-task.articleInfos
		if !ok{
			break
		}
		task.routines<- struct{}{}
		go func() {

			resp, err := proxy.VisitUrl(http.MethodGet, content.MsgLink, nil)
			if err != nil {
				log.Fatal("visit url error!", err)
				return
			}
			defer resp.Body.Close()

			baseFileName:= time.Unix(int64(content.CreateTime),0).Format("2006-01-02")
			if len(baseFileName)>250{
				baseFileName = baseFileName[:100]
			}
			htmlPath:=EditHTML(resp.Body, baseFileName+".html")

			InsertHTMLNode(htmlPath, content.Title)

			pdfGenerator(wkhtmltoxPath,fileSavePath+"\\"+baseFileName+".html",fileSavePath+"\\"+baseFileName+".pdf")
			fmt.Printf("【%d】完成文章：%s\n",goid.Get(),baseFileName)

			<- task.routines
		}()
	}

}



//编辑HTML,并保存HTML
func EditHTML(respBody io.Reader, fileName string) string {
	msgContent, err := ioutil.ReadAll(respBody)
	if err != nil {
		log.Fatal(err)
	}
	content := string(msgContent)

	content = strings.ReplaceAll(content, "data-src", "src")

	//判断文件夹是否存在
	if _,err:=os.Stat(fileSavePath); os.IsNotExist(err){
		//创建文件夹
		os.Mkdir(fileSavePath, os.ModePerm)
	}

	return saveStrContent(content, fileSavePath, fileName)
}

func InsertHTMLNode(htmlPath,nodeText string) {
	//为了确保PDF完整，有时候只有标题
	//root,_ := htmlquery.Parse(content)
	root,_ := htmlquery.LoadDoc(htmlPath)
	xpath:="//div[@id=\"js_content\"]"
	//xpath:="//body"
	x := htmlquery.Find(root,xpath)
	if len(x)>1{
		log.Println("找到多个js_content属性标签!")
		return
	}else if len(x)==1{
		insertNode := &html.Node{Data: "p",Type: html.ElementNode}
		insertNode.InsertBefore(&html.Node{Data:nodeText,Type: html.TextNode},nil)
		x[0].InsertBefore( insertNode,x[0].FirstChild)

	}

	content := htmlquery.OutputHTML(root,true)
	saveFile(content, htmlPath)
}

//保存文章中的图片，并返回保存路径
//func saveImages(url, savePath string) int {
//	fmt.Println("保存路径为：", savePath)
//	resp, err := proxy.VisitUrl(http.MethodGet, url, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	content, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return saveFile(content, savePath)
//}

//保存至文件
func saveFile(content string, savePath string) int {
	f, err := os.OpenFile(savePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	n, err := f.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func saveStrContent(content, savePath, fileName string) string {
	f, err := os.OpenFile(savePath+"\\"+fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
	return savePath+"\\"+fileName
}

//wkhtmltox安装位置："E:\\wkhtmltox\\bin\\wkhtmltopdf.exe"
func SaveAsPDF() func(wkhtmltox, htmlPath, pdfPath string) {
	count, index:=0,0

 	return func (wkhtmltox, htmlPath, pdfPath string){
 		count++
		//调用cmd执行wkhtmltopdf
		c := exec.Command(wkhtmltox, "--enable-local-file-access", htmlPath, pdfPath)
		_, err:=c.Output()
		if err!=nil{
			//跳过error处理
			fmt.Printf("总计调用%d次，第%d次未正常退出\n",count,index)
			index++
		}

	}
}

func CreateDownloadTask(articleInfos chan entity.ArticleInfo, group *sync.WaitGroup,routines chan struct{}) Task {
	task := DownloadTask{
		Name: "default Download Task",
		articleInfos:  articleInfos,
		group: group,
		routines: routines,
	}
	return &task
}
