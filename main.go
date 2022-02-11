package main

import (
	"log"
	"net/http"

	"twistlock/authenticate"
	"twistlock/images"
	"twistlock/model"
)

var (
	BaseURL  = ""
	Username = ""
	Password = ""
	Registry = ""
	Repo     = ""
	Tag      = ""
)

func main() {
	twistlock := model.Scanner{
		BaseURL: BaseURL,
		Client:  &http.Client{},
	}

	twistlock.User = model.User{
		Username: Username,
		Password: Password,
	}

	if err := authenticate.Login(&twistlock); err != nil {
		log.Fatalf("can't authenticate %s", err)
	}

	req := images.ImageRequestBody{
		RepoTag: images.ImageRepo{
			Registry: Registry,
			Repo:     Repo,
			Tag:      Tag,
		},
	}
	// images.Get(&twistlock)
	if err := images.Scan(&twistlock, req); err != nil {
		log.Fatalf("images can't create %s", err)
	}
}
