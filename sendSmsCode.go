package main

//发送短信验证码

import (
	"fmt"
	"net/http"
	"github.com/golang/glog"
	"io/ioutil"
	"unsafe"
	"net/url"
	"bytes"
	"encoding/json"
)

//注册成功短信
func SendRegistSMS(to string, datas string) error {

    if to == "" || datas == ""{
    	fmt.Println("参数错误")
    	return nil
	}

	params := make(map[string]interface{})
	params["account"] = "N1324056"
	params["password"] = "38cjQsmouPe116"
	params["phone"] = to
	params["msg"] = url.QueryEscape(datas)
	params["report"] = "false"
	bytesData, err := json.Marshal(params)

	if err != nil {
		glog.Error(err)
		return nil
	}
	reader := bytes.NewReader(bytesData)
	url := "http://smssh1.253.com/msg/send/json" //请求地址请参考253云通讯自助通平台查看或者询问您的商务负责人获取
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		glog.Error(err)
		return err
	}
	fmt.Println(request,"\n",resp)
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Error(err)
		return err
	}
    fmt.Println(respBytes)
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println("str=",str)

	glog.Error(string(*str))

	return nil
}

func main(){

	fmt.Println(SendRegistSMS("18295514402","很高兴为你服务"))

}
