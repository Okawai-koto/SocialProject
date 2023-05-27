package functions

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostToSignUpService(email string, verifycode string) {
	// API endpoint ve istek verilerini belirleyin
	url := "http://localhost:8081/mail"
	payload := []byte(`{"email": "` + email + `", "verifycode": "` + verifycode + `"}`)

	// POST isteğini oluşturun
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("POST isteği oluşturulamadı:", err)
		return
	}

	// İstek başlıklarını ayarlayın (isteğe bağlı)
	req.Header.Set("Content-Type", "application/json")

	// HTTP isteğini gönderin
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("POST isteği gönderilemedi:", err)
		return
	}
	defer resp.Body.Close()

	// İstek yanıtını okuyun
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("İstek yanıtı okunamadı:", err)
		return
	}

	// İstek yanıtını string olarak yazdırın
	fmt.Println("İstek yanıtı:", string(body))
}
