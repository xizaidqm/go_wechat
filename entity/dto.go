package entity

//公众号搜索结果
type AccountSearchResult struct {
	BaseResp BaseResp            `json:"base_resp"`
	List     []MainlyAccountInfo `json:"list"`
	Total    int                 `json:"total"` //搜索结果总数
}

//基本响应体
type BaseResp struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"err_msg"`
}

//公众号整体信息
type MainlyAccountInfo struct {
	FakeId       string `json:"fakeid"`        //指代公众号ID
	Nickname     string `json:"nickname"`       //公众号名称
	Alias        string `json:"alias"`          //公众号别名
	RoundHeadImg string `json:"round_head_img"` //头像链接
	ServiceType  int    `json:"service_type"`   //公众号类型
	Signature    string `json:"signature"`      //公众号签名
}

//文章搜索结果
type ArticleSearchResult struct {
	BaseResp   BaseResp            `json:"base_resp"`
	AppMsgCnt  int                 `json:"app_msg_cnt"` //文章搜索结果总数
	AppMsgList []MainlyArticleInfo `json:"app_msg_list"`
}

//文章整体信息
type MainlyArticleInfo struct {
	Aid                   string        `json:"aid"`
	AlbumID               string        `json:"album_id"`
	AppmsgAlbumInfos      []interface{} `json:"appmsg_album_infos"`
	Appmsgid              int         `json:"appmsgid"` //文章ID
	Checking              int           `json:"checking"`
	CopyrightType         int           `json:"copyright_type"`
	Cover                 string        `json:"cover"`
	CreateTime            int           `json:"create_time"` //文章创建时间
	Digest                string        `json:"digest"`
	HasRedPacketCover     int           `json:"has_red_packet_cover"`
	IsPaySubscribe        int           `json:"is_pay_subscribe"`
	ItemShowType          int           `json:"item_show_type"`
	ItemIdx               int           `json:"itemidx"`
	Link                  string        `json:"link"` //文章链接
	MediaDuration         string        `json:"media_duration"`
	MediaApiPublishStatus int           `json:"mediaapi_publish_status"`
	PayAlbumInfo          struct {
		AppMsgAlbumInfos []interface{} `json:"appmsg_album_infos"`
	} `json:"pay_album_info"`
	TagId      []interface{} `json:"tagid"`
	Title      string        `json:"title"`       //文章标题app_msg_list
	UpdateTime int           `json:"update_time"` //文章更新时间
}
