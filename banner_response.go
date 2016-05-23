package openrtb

import (
	"golang.org/x/net/html"
	"bytes"
)

type BannerResponse struct {
	ClickUrl    string
	ImageUrl    string
}

//Parses an OpenRTB Banner Response from bytes
func ParseBannerAdmBytes(data []byte) ( *BannerResponse, error) {
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
						break
					}
				}
			}
		}
	}
	return &banner, nil
}
