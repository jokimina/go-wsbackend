package _36kr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)



func RunInformationOverview() *JsonOverviewResponse{
	data := Payload{
		Page: 1,
		PerPage: 100,
		Sort: "date",
		EntityType: "post",
		Keyword: "垃圾",
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://36kr.com/pp/api/search/entity-search", body)
	if err != nil {
		// handle err
	}

	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Origin", "https://36kr.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Referer", "https://36kr.com/search/articles/%E5%9E%83%E5%9C%BE")
	req.Header.Set("Cookie", "acw_tc=276aedde15625920483078361e700b8e27fa9e4d25a931789c335a28f3f752; kr_stat_uuid=BNbww26043200; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22BNbww26043200%22%2C%22%24device_id%22%3A%2216bd1be705729c-0f7f138782a671-e343166-2073600-16bd1be70582b9%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_referrer_host%22%3A%22%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%7D%2C%22first_id%22%3A%2216bd1be705729c-0f7f138782a671-e343166-2073600-16bd1be70582b9%22%7D; device-uid=ac0e6e60-a185-11e9-9a0b-51c874806a80; krnewsfrontss=fbee6d1a40ea9779069cd74601ec2158; M-XSRF-TOKEN=c90c2ab6b7bc5ece75e15d51ab83bb0dbac09443a1fb12e44faba28754351b43; Hm_lvt_1684191ccae0314c6254306a8333d090=1562592047,1563852409; Hm_lpvt_1684191ccae0314c6254306a8333d090=1563852409; Hm_lvt_713123c60a0e86982326bae1a51083e1=1562592047,1563852409; Hm_lpvt_713123c60a0e86982326bae1a51083e1=1563852409; SERVERID=d36083915ff24d6bb8cb3b8490c52181|1563852411|1563852411")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("M-X-Xsrf-Token", "c90c2ab6b7bc5ece75e15d51ab83bb0dbac09443a1fb12e44faba28754351b43")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//re := regexp.MustCompile("initialState=(.*?)</script>")
	//r := re.FindSubmatch(rbody)[1]
	var Json JsonOverviewResponse
	err = json.Unmarshal(rbody, &Json)
	if err != nil {
		log.Println(err.Error())
	}
	return &Json
}

func RunInformationContent(originID string) *JsonContentResponse {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://36kr.com/p/%s", originID), nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Cookie", "acw_tc=276aedde15625920483078361e700b8e27fa9e4d25a931789c335a28f3f752; kr_stat_uuid=BNbww26043200; Hm_lvt_713123c60a0e86982326bae1a51083e1=1562592047; Hm_lvt_1684191ccae0314c6254306a8333d090=1562592047; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22BNbww26043200%22%2C%22%24device_id%22%3A%2216bd1be705729c-0f7f138782a671-e343166-2073600-16bd1be70582b9%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_referrer_host%22%3A%22%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%7D%2C%22first_id%22%3A%2216bd1be705729c-0f7f138782a671-e343166-2073600-16bd1be70582b9%22%7D; device-uid=ac0e6e60-a185-11e9-9a0b-51c874806a80; krnewsfrontss=737e4daed6959ed4a89fa5edec3df984; M-XSRF-TOKEN=e788830bbb5a294747d86522e65d17ff0c4ba9fcf5168cb9c56086b3addbc8f7; Hm_lpvt_1684191ccae0314c6254306a8333d090=1562657086; Hm_lpvt_713123c60a0e86982326bae1a51083e1=1562657086; SERVERID=6754aaff36cb16c614a357bbc08228ea|1562657086|1562653894")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile("initialState=(.*?)</script>")
	r := re.FindSubmatch(rbody)[1]
	var Json JsonContentResponse
	err = json.Unmarshal(r, &Json)
	if err != nil {
		log.Println(err.Error())
	}
	return &Json
}
