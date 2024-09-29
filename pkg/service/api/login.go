package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Egor123qwe/loggy/pkg/model"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s service) login(credentials Credentials) (string, error) {
	data, _ := json.Marshal(credentials)

	req, _ := http.NewRequest(http.MethodPost, s.URL+"/auth/login", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", model.BadCredentialsErr
	}

	cookies := resp.Cookies()
	for _, c := range cookies {
		if c.Name == sessionKey {
			return c.Value, nil
		}
	}

	return "", model.AuthHeadersInvalidErr
}
