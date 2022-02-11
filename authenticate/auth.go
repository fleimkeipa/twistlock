package authenticate

import (
	"fmt"
	"net/http"

	twisthttp "twistlock/http"
	"twistlock/model"
)

func Login(s *model.Scanner) error {
	resp := &model.Authenticate{}

	req := s.User

	opts := twisthttp.ImageRequestOptions{
		Method: http.MethodPost,
		Paths:  []string{"authenticate"},
		Header: "application/json",
	}

	if err := twisthttp.Do(resp, req, opts, s); err != nil {
		return fmt.Errorf("can't login %s", err)
	}

	s.Token = resp.Token
	return nil
}
