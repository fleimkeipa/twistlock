package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"twistlock/model"
)

type ImageRequestOptions struct {
	Method string   `json:"method"`
	Paths  []string `json:"paths"`
	Header string   `json:"header"`
}

func Do(responseBody, requestBody interface{}, opts ImageRequestOptions, s *model.Scanner) error {
	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(reqBody)

	u, err := JoinPath(append([]string{s.BaseURL}, opts.Paths...)...)
	if err != nil {
		return err
	}

	p, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return err
	}

	u.RawQuery = p.Encode()

	req, err := http.NewRequest(opts.Method, u.String(), buf)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", opts.Header)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.Token))

	res, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("request failed: %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if responseBody == nil && res.StatusCode == 200 {
		return nil
	}

	if err = json.Unmarshal(body, &responseBody); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return nil
}

// JoinPath joins URL host and paths with by removing all leading slashes from paths and adds a trailing slash to end of paths except last one.
func JoinPath(paths ...string) (*url.URL, error) {
	uri := new(url.URL)
	for i, path := range paths {
		path = strings.TrimLeft(path, "/")
		if i+1 != len(paths) && !strings.HasSuffix(path, "/") {
			path = fmt.Sprintf("%s/", path)
		}

		u, err := url.Parse(path)
		if err != nil {
			return nil, err
		}

		uri = uri.ResolveReference(u)
	}

	return uri, nil
}
