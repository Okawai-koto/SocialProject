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

func SendEmail() {
	// AWS kimlik bilgileri ve bölgesi
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("AKIAdedsalk", "dsakşdsşalkd", ""),
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
		</body>
	</html>
	`

	// Ek dosyası
	filePath := "C:/Users/FIRAT/Desktop/social/signupService/aws/attachment.pdf"
	attachmentData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Ek dosyasını base64 ile kodlama
	attachmentContent := base64.StdEncoding.EncodeToString(attachmentData)

	// Raw e-posta mesajı
	rawMessage := fmt.Sprintf("From: sairmisinnesin@gmail.com\r\n" +
		"To: itsbap22@gmail.com\r\n" +
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
