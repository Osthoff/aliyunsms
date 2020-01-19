//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/1/19 4:09 PM
// 版本：V1.0
//============================================================

package main

import (
	"aliyunsms"
	"fmt"
)

func main()  {
	aliyun_sms, err := aliyunsms.NewAliyunSms("绿洲教育网", "SMS_*****", "*****", "*****")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = aliyun_sms.Send("1314****973", `{"number":"1234"}`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}
