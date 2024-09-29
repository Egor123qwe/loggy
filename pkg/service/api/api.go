package api

import (
	"errors"

	"github.com/Egor123qwe/loggy/pkg/model"
)

const (
	sessionKey = "logs_viewer_session"
)

type Service interface {
	Init(module string) (InitResp, error)
}

type service struct {
	URL   string
	token string

	credentials Credentials
}

func New(URL string, credentials Credentials) (Service, error) {
	srv := &service{
		URL:         URL,
		credentials: credentials,
	}

	token, err := srv.login(srv.credentials)
	if err != nil {
		return nil, err
	}

	srv.token = token

	return srv, nil
}

type InitResp struct {
	ModuleID int64 `json:"module_id"`
}

func (s service) Init(module string) (InitResp, error) {
	resp, err := s.init(module)
	if err != nil && !errors.Is(err, model.UnauthorizedErr) {
		return InitResp{}, err
	}

	if err == nil {
		return resp, nil
	}

	if errors.Is(err, model.UnauthorizedErr) {
		s.token, err = s.login(s.credentials)
		if err != nil {
			return InitResp{}, err
		}
	}

	return s.init(module)
}
