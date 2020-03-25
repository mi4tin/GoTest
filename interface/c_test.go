package _interface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"moqikaka.com/X1/PlatformHelper/timeHelper"
	"moqikaka.com/goutil/timeUtil"
	"testing"
)

type AB struct {
	C interface{}
}

type AB1 struct {
	C string
}

func TestA(t *testing.T){
	//测试json无科学计数转换
	s:=`[{"C":65811045}]`
	v := &[]AB{}
	d := json.NewDecoder(bytes.NewReader([]byte(s)))
	d.UseNumber()
	err := d.Decode(&v)
	if err != nil {
		// 错误处理
	}
	fmt.Println((*v)[0].C)
	return

	abObj:=map[string]interface{}{}

	json.Unmarshal([]byte(s),abObj)


	fmt.Println(abObj["C"])


	//fmt.Println(fmt.Sprintf("%.0f",abObj.C))

	return
	//时间格式化测试
	fmt.Println(timeUtil.ToDateTimeString2(timeHelper.GetNowAddMinTime(10)))
}
