package images

import (
	"fmt"
	"net/http"

	twisthttp "twistlock/http"
	"twistlock/model"
)

type ImageRequestBody struct {
	RepoTag ImageRepo `json:"repoTag"`
}
type ImageRepo struct {
	Registry string `json:"registry"`
	Repo     string `json:"repo"`
	Tag      string `json:"tag"`
}

func Get(s *model.Scanner) {
	opts := twisthttp.ImageRequestOptions{
		Method: http.MethodGet,
		Paths:  []string{"images"},
		Header: "application/json",
	}

	report := new(model.ScanReport)
	if err := twisthttp.Do(report, nil, opts, s); err != nil {
		panic(fmt.Errorf("can't get report %s", err))
	}

	for _, v := range *report {
		fmt.Println("repoTag", v.RepoTag.Repo)
	}
}

func Scan(s *model.Scanner, req ImageRequestBody) error {
	opts := twisthttp.ImageRequestOptions{
		Method: http.MethodPost,
		Paths:  []string{"images", "scan"},
		Header: "application/json",
	}

	if err := twisthttp.Do(nil, req, opts, s); err != nil {
		return fmt.Errorf("can't get report %s", err)
	}

	return nil
}
