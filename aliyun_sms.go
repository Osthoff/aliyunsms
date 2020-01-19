//============================================================
// 描述:
// 作者: Yang
// 日期: 2020/1/19 3:44 PM
// 版本：V1.0
//============================================================

package aliyunsms

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type AliyunSms struct {
	AccessKeyID  string //用户AccessKey
	AccessSecret string //用户AccessKey
	SignName     string //签名名称
	TemplateCode string //模板code
}



func NewAliyunSms(sign_name string, template_code string, access_key_id string, access_secret string) (*AliyunSms, error) {
	var a AliyunSms
	a.SignName = sign_name
	a.TemplateCode = template_code
	a.AccessKeyID = access_key_id
	a.AccessSecret = access_secret

	return &a, nil
}

func (this *AliyunSms) Send(numbers string, params string) error {
	var request Request
	request.Format = "JSON"
	request.Version = "2017-05-25"
	request.AccessKeyId = this.AccessKeyID
	request.SignatureMethod = "HMAC-SHA1"
	request.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	request.SignatureVersion = "1.0"
	request.SignatureNonce = uuid.New().String()

	request.Action = "SendSms"
	request.SignName = this.SignName
	request.TemplateCode = this.TemplateCode
	request.PhoneNumbers = numbers
	request.TemplateParam = params
	request.RegionId = "cn-hangzhou"

	url := request.ComposeUrl("GET", this.AccessSecret)
	var resp *http.Response
	var err error
	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_m := make(map[string](string))
	err = json.Unmarshal(b, &_m)
	if err != nil {
		return err
	}
	message, ok := _m["Message"]
	if ok && strings.ToUpper(message) == "OK" {
		return nil
	}
	if ok {
		return errors.New(message)
	}
	return errors.New("send sms error")
}