package alisms

import (
	"fmt"
	"log"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func SendAuthOTP(phoneNumber string, otp string) (res *dysmsapi.SendSmsResponse, err error) {
	if len(phoneNumber) == 0 {
		return nil, fmt.Errorf("phone number not found")
	}

	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ALI_SMS_ACCESSKEY_ID"), os.Getenv("ALI_SMS_ACCESS_SECRET"))
	if err != nil {
		panic(err)
	}
	//client.SetLogger("level", "channel", file)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phoneNumber

	request.SignName = os.Getenv("ALI_SMS_SIGN_NAME")
	request.TemplateCode = os.Getenv("ALI_SMS_OTP_TEMPLATE")
	request.TemplateParam = `{"code":` + otp + `}`
	response, err := client.SendSms(request)
	if err != nil {
		log.Println("Aliyun SMS Error:", err)
		return nil, err
	}
	return response, nil
}
