package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
)

func main() {

	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

	host := "mail.ciichr.com"
	username := "zhongzhibaoxian2.zzbx@ciichr.com"
	password := "abcd.1234"
	// host := "mail.ciichr.com"
	// username := "zhongzhibaoxian2.zzbx@ciichr.com"
	// password := "abcd.1234"

	title := "这是邮件标题"
	from := mail.Address{"发送人名称", "zhongzhibaoxian2.zzbx@ciichr.com"}
	to := mail.Address{"接收人名称", "fangmingwang.zzbx@ciichr.com"}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(title)))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	body := "这是一封测试邮件!"
	message += "\r\n" + b64.EncodeToString([]byte(body))

	auth := smtp.PlainAuth(
		"",
		username,
		password,
		host,
	)

	err := smtp.SendMail(
		host+":587",
		auth,
		username,
		[]string{to.Address},
		[]byte(message),
	)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("OK")
}
