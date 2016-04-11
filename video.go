package openrtb

import (
	"errors"
)

// The "video" object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type Video struct {
	Mimes          []string   `json:"mimes,omitempty"`          // Content MIME types supported.
	Minduration    int        `json:"minduration,omitempty"`    // Minimum video ad duration in seconds
	Maxduration    int        `json:"maxduration,omitempty"`    // Maximum video ad duration in seconds
	Protocol       int        `json:"protocol,omitempty"`       // Video bid response protocols
	Protocols      []int      `json:"protocols,omitempty"`      // Video bid response protocols
	W              int        `json:"w,omitempty"`              // Width of the player in pixels
	H              int        `json:"h,omitempty"`              // Height of the player in pixels
	Startdelay     int        `json:"startdelay,omitempty"`     // Indicates the start delay in seconds
	Linearity      int        `json:"linearity,omitempty"`      // Indicates whether the ad impression is linear or non-linear
	Sequence       int        `json:"sequence,omitempty"`       // Default: 1
	Battr          []int      `json:"battr,omitempty"`          // Blocked creative attributes
	Maxextended    int        `json:"maxextended,omitempty"`    // Maximum extended video ad duration
	Minbitrate     int        `json:"minbitrate,omitempty"`     // Minimum bit rate in Kbps
	Maxbitrate     int        `json:"maxbitrate,omitempty"`     // Maximum bit rate in Kbps
	Boxingallowed  int        `json:"boxingallowed,omitempty"`  // If exchange publisher has rules preventing letter boxing
	Playbackmethod []int      `json:"playbackmethod,omitempty"` // List of allowed playback methods
	Delivery       []int      `json:"delivery,omitempty"`       // List of supported delivery methods
	Pos            int        `json:"pos,omitempty"`            // Ad Position
	Companionad    []Banner   `json:"companionad,omitempty"`
	Api            []int      `json:"api,omitempty"` // List of supported API frameworks
	Companiontype  []int      `json:"companiontype,omitempty"`
	Ext            Extensions `json:"ext,omitempty"`
}

// Returns the sequence number, with default fallback
func (v *Video) Seq() int {
	return v.Sequence
}

// Returns the boxing permission status, with default fallback
func (v *Video) IsBoxingAllowed() bool {
	return v.Boxingallowed == 1
}

// Validation errors
var (
	ErrInvalidVideoMimes       = errors.New("openrtb parse: video has no mimes")
	ErrInvalidVideoLinearity   = errors.New("openrtb parse: video linearity missing")
	ErrInvalidVideoMinduration = errors.New("openrtb parse: video minduration missing")
	ErrInvalidVideoMaxduration = errors.New("openrtb parse: video maxduration missing")
	ErrInvalidVideoProtocol    = errors.New("openrtb parse: video protocol missing")
)

// Validates the object
func (v *Video) Valid() (bool, error) {
	if len(v.Mimes) == 0 {
		return false, ErrInvalidVideoMimes
	} else if v.Linearity == 0 {
		return false, ErrInvalidVideoLinearity
	} else if v.Minduration == 0 {
		return false, ErrInvalidVideoMinduration
	} else if v.Maxduration == 0 {
		return false, ErrInvalidVideoMaxduration
	} else if v.Protocol == 0 {
		return false, ErrInvalidVideoProtocol
	}
	return true, nil
}

// Applies defaults
func (v *Video) WithDefaults() *Video {
	v.Sequence = 1
	v.Boxingallowed = 1
	return v
}
