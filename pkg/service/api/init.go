package api

import (
	"encoding/json"
	"net/http"

	"github.com/Egor123qwe/loggy/pkg/model"
)

func (s service) init(module string) (InitResp, error) {
	req, _ := http.NewRequest(http.MethodGet, s.URL+"/api/module/init?module="+module, nil)
	req.Header.Set("Content-Type", "application/json")

	cookie := &http.Cookie{
		Name:  sessionKey,
		Value: s.token,
	}

	req.AddCookie(cookie)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return InitResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return InitResp{}, model.UnauthorizedErr
		}

		return InitResp{}, model.BadRequestErr
	}

	var result InitResp

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return InitResp{}, err
	}

	return result, nil
}
