package openrtb

import (
	"encoding/json"
	"errors"
	"io"
)

var ErrBlankNativeResponse = errors.New(`Native response doesn't contain "native" node`)

type NativeAdm struct {
	Native *NativeResponse `json:"native,omitempty"`
}

type NativeResponse struct {
	Ver         string          `json:"ver,omitempty"` // Version of the Native Markup version in use
	Assets      []ResponseAsset `json:"assets"`        // Array of Asset Objects
	Link        *Link           `json:"link"`
	Imptrackers []string        `json:"imptrackers,omitempty"`
	Jstracker   string          `json:"jstracker,omitempty"`
	Ext         Extensions      `json:"ext,omitempty"`
}

type ResponseAsset struct {
	Id       int            `json:"id"`                 // Unique asset ID, assigned by exchange
	Required int            `json:"required,omitempty"` // Set to 1 if asset is required
	Title    *ResponseTitle `json:"title,omitempty"`    // Title object for title assets
	Img      *ResponseImg   `json:"img,omitempty"`      // Image object for image assets
	Video    *ResponseVideo `json:"video,omitempty"`    // Video object for video assets
	Data     *ResponseData  `json:"data,omitempty"`     // Data object for ratings, price, etc.
	Link     *Link          `json:"link,omitempty"`
	Ext      Extensions     `json:"ext,omitempty"`
}

type Link struct {
	Url           string     `json:"url"`
	Clicktrackers []string   `json:"clicktrackers,omitempty"`
	Fallback      string     `json:"fallback,omitempty"`
	Ext           Extensions `json:"ext,omitempty"`
}

type ResponseTitle struct {
	Text string     `json:"text"`
	Ext  Extensions `json:"ext,omitempty"`
}

type ResponseImg struct {
	Url string     `json:"url"`
	W   int        `json:"w,omitempty"` // Width
	H   int        `json:"h,omitempty"` // Height
	Ext Extensions `json:"ext,omitempty"`
}

type ResponseData struct {
	Label string     `json:"label,omitempty"`
	Value string     `json:"value"`
	Ext   Extensions `json:"ext,omitempty"`
}

type ResponseVideo struct {
	Vasttag string `json:"vasttag"`
}

//Parses an OpenRTB Native Response from an io.Reader
func ParseNativeAdm(reader io.Reader) (adm *NativeAdm, err error) {
	dec := json.NewDecoder(reader)
	if err = dec.Decode(&adm); err != nil {
		return nil, err
	}
	return nativeAdmWithDefaults(adm)
}

//Parses an OpenRTB Native Response from bytes
func ParseNativeAdmBytes(data []byte) (adm *NativeAdm, err error) {
	if err = json.Unmarshal(data, &adm); err != nil {
		return nil, err
	}
	return nativeAdmWithDefaults(adm)
}

// Applies NativeResponse defaults
func (resp *NativeResponse) WithDefaults() *NativeResponse {
	if len(resp.Ver) == 0 {
		resp.Ver = "1"
	}
	for id, asset := range resp.Assets {
		resp.Assets[id] = asset
	}
	return resp
}

func nativeAdmWithDefaults(adm *NativeAdm) (*NativeAdm, error) {
	if adm == nil || adm.Native == nil {
		return adm, ErrBlankNativeResponse
	}
	adm.Native = adm.Native.WithDefaults()
	return adm, nil
}
