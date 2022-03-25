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
)

//发送短信
func SendMessage(phone, values string) error {
	credential := common.NewCredential(
		"AKIDEI6ewzlGVUCdnhKglZ6ORDiwtZz80f4D",
		"sTnre9ON5WgaDZeS1JJ2VhcFciX5MM5Y",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs([]string{phone})
	request.SmsSdkAppId = common.StringPtr("1400642372")
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
