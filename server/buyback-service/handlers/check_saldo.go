package handlers

import (
	"buyback-service/dto"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

func CheckSaldo(norek string, gram float32) (bool, error) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	reqPayload := map[string]string{
		"norek": norek,
	}

	jsonReq, err := json.Marshal(reqPayload)
	if err != nil {
		return false, err
	}

	url := os.Getenv("URL_CHECK_SALDO")
	request, err := http.NewRequest("POST", url, bytes.NewReader(jsonReq))
	if err != nil {
		return false, err
	}

	response, err := client.Do(request)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	data := dto.CheckSaldoResponse{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return false, errors.New("Fail to get response")
	}

	if data.Saldo.Saldo < gram {
		return false, nil
	}

	return true, nil
}
