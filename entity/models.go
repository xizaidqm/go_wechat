package entity

//公众号格式
type AccountInfo struct {
	FakeId   string `json:"fake_id"`   //公众号的id
	Nickname string `json:"nickname"` //公众号的名称
}

//爬取的资源
type Resource struct {
	MsgLink string //文章链接
}

type ArticleInfo struct {
	MsgId      string `json:"msg_id"`      //文章ID
	Title      string `json:"title"`       //文章标题
	MsgLink    string `json:"msg_link"`    //文章链接
	CreateTime int `json:"create_time"` //创建时间
	UpdateTime int `json:"update_time"` //更新时间
}

