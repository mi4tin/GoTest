package temp

import (
	"encoding/json"
	"fmt"
	"moqikaka.com/goutil/securityUtil"
	"testing"
)

func TestDefer(t *testing.T) {

	fmt.Print(securityUtil.Md5String("amount=600&clientType=1&cpOrderId=20200312141712_523615&extInfo=20200312141712_523615&notifyUrl=http://c.79yougame.com/callback/dzzbt01and_GameFan&orderId=20200312141712734838396059192700&orderStatus=1&paySucTime=20200312141738&payTime=20200312141712&platformId=4&roleId=11963650-13cc-4f6f-aa48-8a7aeb1cccd8&uid=26213553_aFosuU&zoneId=5001fpun6i4ovrfdtmwkbplczkv4rmbd9ikq",false))

	return
	s:=[]byte(`{"StatusCode":0, "Message":"认证成功"}`)

	resultMap := map[string]interface{}{}

	if err := json.Unmarshal(s, &resultMap); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v",resultMap["StatusCode1"])

	if code, ok := resultMap["StatusCode"].(int64); !ok || code != 0 {
		fmt.Println(fmt.Sprintf("code-%#v-%v-%v",resultMap,code,ok))
	}

	fmt.Println("end:",resultMap["StatusCode"]==0)

}


func TestDefer1(t *testing.T) {

	return
	//var i interface{} = "TT"
	var i interface{} = ""
	value, ok := i.(int)
	if ok {
		fmt.Printf("类型匹配int:%v", value)
	} else {
		fmt.Println("类型不匹配int")
	}
}