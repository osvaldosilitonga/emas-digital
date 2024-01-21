package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
	"topup-service/dto"
)

func CheckHarga(harga int) (bool, error) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	url := os.Getenv("URL_CHECK_HARGA")
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	response, err := client.Do(request)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	data := dto.APIResponse{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return false, errors.New("Fail to get response")
	}

	if data.Data.HargaTopup != harga {
		return false, nil
	}

	return true, nil
}
