package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Get(endpoint string, token string, data interface{}) error {
	client := http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	return json.NewDecoder(resp.Body).Decode(&data)

}
