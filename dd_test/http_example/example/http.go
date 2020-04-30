package example

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

//ReponseBody ...
type ReponseBody struct {
	Data    interface{} `json:"_data"`
	ErrCode string      `json:"_errCode"`
	ErrStr  string      `json:"_errStr"`
}

type item struct {
	Avatar    string `json:"avatar"`
	Amount    int    `json:"amount"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

// type items []item

//HTTPGet ...
func HTTPGet() {
	// resp, err := http.Get("http://localhost:8081/v1/order/getYesterdayLoopList?pid=1001")
	resp, err := http.Get("http://localhost:8081/v1/order/getHistoryListForH5?pid=1001")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// var res ReponseBody
	// err = json.Unmarshal(body, &res)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// var items []item
	// fmt.Printf("%s", body)
	// fmt.Printf("%+v", res)

	// fmt.Printf("%s", res.Data)
	// fmt.Println(reflect.TypeOf(res.Data))
	// // info := res.Data.([]interface{})[0].(map[string]interface{})["name"]
	// value, ok := (res.Data).([]interface{})
	// // fmt.Println(ok)
	// if !ok {
	// 	fmt.Println(ok)
	// }
	// fmt.Printf("%+v", value[0].(map[string]interface{})["name"])

	// value, ok := (res.Data).([]map[string]interface{})
	// // fmt.Println(ok)
	// if !ok {
	// 	fmt.Println(ok)
	// }
	// fmt.Printf("%+v", value[0]["name"])
	// var items []item
	js, err := simplejson.NewJson([]byte(body)) //反序列化

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(js.Get("_data").GetIndex(0).Get("amount").String())

}

//HTTPPost ...
func HTTPPost() {
	resp, err := http.Post("https://www.baidu.com", "application/json", strings.NewReader("name=wt"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
}

//HTTPPostForm ...
func HTTPPostForm() {
	resp, err := http.PostForm("https://www.baidu.com", url.Values{"name": {"wt"}})
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

//HTTPDo ...
func HTTPDo() {
	client := http.DefaultClient
	var reader io.Reader
	reader = strings.NewReader("name=wt")
	reader = bytes.NewReader([]byte("name=wt1"))
	req, err := http.NewRequest("POST", "https://www.baidu.com", reader)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Cookie", "name=wt")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
