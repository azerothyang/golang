package policy

import (
	"encoding/json"
	"net/http"
	"errors"
	"io/ioutil"
)

const uri  =  "http://talent.qh-1.cn/pc/home/newlist?version=personal&cur_code=0086&page=0&pageSize=1"

type RetData struct {
	Code int `json:"code"`
	Data data `json:"data"`
	Message string `json:"message"`
}


type data struct {
	NewPolicies []newPolicy `json:"newPolicies"`
	Total uint8 `json:"total"`
}

type newPolicy struct {
	AreaName string `json:"area_name"`
	CityName string `json:"city_name"`
	CommunityName string `json:"community_name"`
	CurName string `json:"cur_name"`
	FromSource string `json:"from_source"`
	ImageUrl string `json:"image_url"`
	NewsTitle string `json:"news_title"`
	PolicyCode string `json:"policy_code"`
	ProvinceName string `json:"province_name"`
	PublishTime string `json:"publish_time"`
	StreetName string `json:"street_name"`
}

//并行发送请求并解析, 将解析后的数据写入通道
func RequestAndAnalyze(chRetData chan <- *RetData){
	retData, err := Request()
	if err != nil {
		return
	}
	chRetData <- retData
}

//发送请求函数, 并且将请求解析的内容返回
func Request() (*RetData, error) {
	var retData RetData //retData
	var retBytes []byte
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//判断请求如果不正常, 则打印日志退出
	if res.StatusCode != 200 {
		return nil, errors.New("返回http码错误" + string(res.StatusCode))
	}

	if retBytes, err = ioutil.ReadAll(res.Body); err != nil {
		return &retData, err
	}
	//解析内容
	err = json.Unmarshal(retBytes, &retData) //retData返回的结构体 必需是地址, 也不能为空指针nil
	if err != nil {
		return nil, err
	}
	return &retData, nil
}