package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	urlPkg "net/url"
)

const (
	DefaultServer = `https://labs.diffbot.com/testdrive/article?token=testdriverehjenztgeil`
)

type Error struct {
	ErrCode    int    `json:"errorCode"` // Description of the error
	ErrMessage string `json:"error"`     // Error code per the chart below
	RawString  string `json:"-"`         // Raw json format error string
}

// ParseJson parses the JSON-encoded error data.
func (p *Error) ParseJson(s string) error {
	if err := json.Unmarshal([]byte(s), p); err != nil {
		return err
	}
	p.RawString = s
	return nil
}

func (p *Error) Error() string {
	d, _ := json.Marshal(p)
	return string(d)
}

// Diffbot uses computer vision, natural language processing
// and machine learning to automatically recognize
// and structure specific page-types.
func Diffbot(method, token, url string, opt *Options) (body []byte, err error) {
	return DiffbotServer(DefaultServer, method, token, url, opt)
}

// DiffbotServer like Diffbot function, but support custom server.
func DiffbotServer(server, method, token, url string, opt *Options) (body []byte, err error) {
	req, err := http.NewRequest("GET", makeRequestUrl(server, method, token, url, opt), nil)
	if err != nil {
		return nil, err
	}
	if opt != nil && opt.CustomHeader != nil {
		req.Header = opt.CustomHeader
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		if len(body) != 0 {
			var apiError Error
			if err = apiError.ParseJson(string(body)); err != nil {
				err = &Error{
					ErrCode:    resp.StatusCode,
					ErrMessage: string(body),
				}
				return
			} else {
				err = &apiError
				return
			}
		} else {
			err = &Error{
				ErrCode:    resp.StatusCode,
				ErrMessage: resp.Status,
			}
			return
		}
	}
	return
}

func makeRequestUrl(server, method, token, webUrl string, opt *Options) string {
	return fmt.Sprintf("%s&url=%s%s",
		server, urlPkg.QueryEscape(webUrl), opt.MethodParamString(method),
	)
}
