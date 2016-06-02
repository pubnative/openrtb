package openrtb

import (
	"bytes"

	"golang.org/x/net/html"
)

type BannerResponse struct {
	ClickUrl string
	ImageUrl string
	Title    string
}

//Parses an OpenRTB Banner Response from bytes
func ParseBannerAdmBytes(data []byte) (*BannerResponse, error) {
	banner := BannerResponse{}
	r := bytes.NewReader(data)
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()

		if tt == html.ErrorToken {
			return &banner, nil
		}

		if tt == html.SelfClosingTagToken || tt == html.StartTagToken {
			t := z.Token()
			element := t.Data

			if element == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						banner.ClickUrl = a.Val
						break
					}
				}
				continue
			}

			if element == "img" {
				for _, a := range t.Attr {
					if a.Key == "src" {
						banner.ImageUrl = a.Val
						continue
					}

					if a.Key == "alt" {
						banner.Title = a.Val
						continue
					}
				}
			}
		}
	}
	return &banner, nil
}
