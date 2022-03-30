/*******
* @Author:qingmeng
* @Description:
* @File:sms
* @Date2022/3/25
 */

package service

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"shop/secret"
)

//发送短信
func SendMessage(phone, values string) error {
	credential := common.NewCredential(
		secret.MyConf.SecretId,
		secret.MyConf.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs([]string{phone})
	request.SmsSdkAppId = common.StringPtr("1400642372")
	request.SignName = common.StringPtr("邹云个人测试使用")
	request.TemplateId = common.StringPtr("1327589")
	request.TemplateParamSet = common.StringPtrs([]string{values, "2"})

	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err
	}
	if err != nil {
		return err
	}
	fmt.Printf("%s", response.ToJsonString())
	return nil
}
