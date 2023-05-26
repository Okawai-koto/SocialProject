package aws

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	sender = "descless26@gmail.com"
	//toWho  = "emos004@gmail.com"
)

func SendEmail(toWho string, verify string) {
	// AWS kimlik bilgileri ve bölgesi
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-north-1"),
		Credentials: credentials.NewStaticCredentials("AKIAXBICWBZOLK7BWBDT", "4a3T/TLnomfR5WDd1BFsdd7xwz1U4ET2SxA5HORZ", ""),
	})
	if err != nil {
		log.Fatal(err)
	}

	// E-posta servisi
	svc := ses.New(sess)

	htmlContent :=
		`
	<html>
		<body> 
			<h1>Bu bir test e-postasıdır.</h1>
			<h1>Okawaii KOTTO!!</h1>
			<p>Merhaba, bu e-posta bir test e-postasıdır.</p>
			<br>
		</body>
	</html>
	` + verify

	// Ek dosyası
	filePath := "C:/Users/deneme/Documents/GitHub/SocialProject/mailService/src/aws/attachment.pdf"
	attachmentData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Ek dosyasını base64 ile kodlama
	attachmentContent := base64.StdEncoding.EncodeToString(attachmentData)

	// Raw e-posta mesajı
	rawMessage := fmt.Sprintf("From: " + sender + "\r\n" +
		"To: " + toWho + "\r\n" +
		"Subject: Test Email with Attachment\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: multipart/mixed; boundary=boundary\r\n" +
		"\r\n" +
		"--boundary\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		htmlContent +
		"\r\n" +
		"--boundary\r\n" +
		"Content-Type: application/pdf\r\n" +
		"Content-Disposition: attachment; filename=\"attachment.pdf\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"\r\n" +
		attachmentContent + "\r\n" +
		"--boundary--")

	// Raw e-posta mesajını gönderme isteği
	_, err = svc.SendRawEmail(&ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{
			Data: []byte(rawMessage),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("E-posta gönderildi.")
}
