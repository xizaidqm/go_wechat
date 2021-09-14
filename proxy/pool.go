package proxy

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)


func VisitUrl(method, url string, body io.Reader) (resp *http.Response, err error) {
	request,err:= http.NewRequest(method, url, body)
	if err!=nil{
		log.Fatal("request error!",err)
		return nil, err
	}
	fillHeaderInfo(request)
	client := &http.Client{}
	return client.Do(request)
}

func fillHeaderInfo(req *http.Request) {
	req.Header.Set("Accept","*/*")
	req.Header.Set("accept-language"," zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")

	//req.Header.Set("Referer"," https://mp.weixin.qq.com/cgi-bin/appmsg?t=media/appmsg_edit_v2&action=edit&isNew=1&type=10&token=1921878132&lang=zh_CN")
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/92.0.902.84")
	req.Header.Set("Cookie","appmsglist_action_3930279374=card; pgv_pvid=5924932349; ptcz=86efab138b777ff6b1e4f2f32f6e285a7cb58d88280d7ad981d3266329f9748b; pgv_pvi=4271604736; RK=ACC1vE9WS8; pt_sms_phone=151******68; _ga=GA1.2.49328439.1599103251; eas_sid=e106i0q0S699v9i2R1Z5T6e6j2; ied_qq=o1051234004; o_cookie=1051234004; pac_uid=1_1051234004; tvfe_boss_uuid=c6d02811c0c4d53f; uin_cookie=o1051234004; fqm_pvqid=c76c04be-fdef-4492-9668-2092a88c9d01; pgg_uid=736163321; pgg_appid=101503919; pgg_openid=4FD5E6D0B8F7292A9347227C00F03754; pgg_access_token=8A87F4C9F9535E4AAEA96F9D8DE46AAC; pgg_type=1; pgg_user_type=5; ua_id=0Qxif6f3j9bSjfmQAAAAAJ8knPRgVhk0c6mo_VjIka0=; wxuin=29994081402622; mm_lang=zh_CN; ptui_loginuin=2804007837; rand_info=CAESIDlL2EUWDcFqNVfavBTW0oTZsPyxsVOZYQoDfqY7kFYh; slave_bizuin=3930279374; data_bizuin=3930279374; bizuin=3930279374; data_ticket=Lr1RwhWY0YDPG5M2hAMO3ysM3mSRk+YY+qRlffnKTRzK8PqJK4leHLEEdtU3zH+2; slave_sid=bGxWYmo4SjJjUDhtUFdNcmZuYzJtR0E0SjM0azA0dlRRdWpZQ0Vfb3VDVmZaQ2g4NEFFM0tNNEUwdDgxQ0VPR2lVWXI2SWdVQVM2SjhrQTAzYmlmRDFIc01yODVOUEFqclVCeU80bnBrcXFuZFNvUUJGTWNLT1dwNVVQc2NkdVl6OG1hS0EzREJVeEd4eFQ3; slave_user=gh_bce985735f9f; xid=cde15ba508f271c299bc22256465b7e1")

}

func UserAgent() string{
	agent := [...]string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent,Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"User-Agent,Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return agent[r.Intn(len(agent))]
}