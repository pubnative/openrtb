package openrtb

import (
	"encoding/json"
	"io"
)

type AltNativeAdm struct {
	Native *AltNativeResponse `json:"native,omitempty"`
}

type AltNativeResponse struct {
	Ver         int             `json:"ver,omitempty"` // Version of the Native Markup version in use
	Assets      []ResponseAsset `json:"assets"`        // Array of Asset Objects
	Link        *Link           `json:"link"`
	Imptrackers []string        `json:"imptrackers,omitempty"`
	Jstracker   string          `json:"jstracker, omitempty"`
	Ext         Extensions      `json:"ext,omitempty"`
}

//Parses an OpenRTB Native Response from an io.Reader
func ParseAltNativeAdm(reader io.Reader) (adm *AltNativeAdm, err error) {
	dec := json.NewDecoder(reader)
	if err = dec.Decode(&adm); err != nil {
		return nil, err
	}
	return altNativeAdmWithDefaults(adm)
}

//Parses an OpenRTB Native Response from bytes
func ParseAltNativeAdmBytes(data []byte) (adm *AltNativeAdm, err error) {
	if err = json.Unmarshal(data, &adm); err != nil {
		return nil, err
	}
	return altNativeAdmWithDefaults(adm)
}

// Applies AltNativeResponse defaults
func (resp *AltNativeResponse) WithDefaults() *AltNativeResponse {
	if resp.Ver == 0 {
		resp.Ver = 1
	}
	for id, asset := range resp.Assets {
		resp.Assets[id] = asset
	}
	return resp
}

func altNativeAdmWithDefaults(adm *AltNativeAdm) (*AltNativeAdm, error) {
	if adm == nil || adm.Native == nil {
		return adm, ErrBlankNativeResponse
	}
	adm.Native = adm.Native.WithDefaults()
	return adm, nil
}
