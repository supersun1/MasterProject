package main

import "encoding/json"

type Article struct {
	Url           string                 `json:"url"`
	ResolvedUrl   string                 `json:"resolved_url"`
	Icon          string                 `json:"icon"`
	Meta          map[string]interface{} `json:"meta,omitempty"`        // Returned with fields.
	QueryString   string                 `json:"querystring,omitempty"` // Returned with fields.
	Links         []string               `json:"links,omitempty"`       // Returned with fields.
	Type          string                 `json:"type"`
	Title         string                 `json:"title"`
	Text          string                 `json:"text"`
	Html          string                 `json:"html"`
	NumPages      string                 `json:"numPages"`
	Date          string                 `json:"date"`
	Author        string                 `json:"author"`
	Tags          []string               `json:"tags,omitempty"`          // Returned with fields.
	HumanLanguage string                 `json:"humanLanguage,omitempty"` // Returned with fields.
	Images        []struct {
		Url         string `json:"url"`
		PixelHeight int    `json:"pixelHeight"`
		PixelWidth  int    `json:"pixelWidth"`
		Caption     string `json:"caption"`
		Primary     string `json:"primary"`
	} `json:"images"`
	Videos []struct {
		Url         string `json:"url"`
		PixelHeight int    `json:"pixelHeight"`
		PixelWidth  int    `json:"pixelWidth"`
		Primary     string `json:"primary"`
	} `json:"videos"`
}

// type of Article.Images[?]
type articleImageType struct {
	Url         string `json:"url"`
	PixelHeight int    `json:"pixelHeight"`
	PixelWidth  int    `json:"pixelWidth"`
	Caption     string `json:"caption"`
	Primary     string `json:"primary"`
}

// type of Article.Videos[?]
type articleVideoType struct {
	Url         string `json:"url"`
	PixelHeight int    `json:"pixelHeight"`
	PixelWidth  int    `json:"pixelWidth"`
	Primary     string `json:"primary"`
}

func ParseIt(token, url string, opt *Options) (*Article, error) {
	body, err := Diffbot("article", token, url, opt)
	if err != nil {
		return nil, err
	}
	var result Article
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *Article) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}
