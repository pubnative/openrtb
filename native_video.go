package openrtb

// The "video" object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type NativeVideo struct {
	Mimes          []string   `json:"mimes"`                    // Content MIME types supported.
	Minduration    int        `json:"minduration"`              // Minimum video ad duration in seconds
	Maxduration    int        `json:"maxduration"`              // Maximum video ad duration in seconds
	Protocol       int        `json:"protocol,omitempty"`       // Video bid response protocols
	Protocols      []int      `json:"protocols"`                // Video bid response protocols
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

// Returns the boxing permission status, with default fallback
func (v *NativeVideo) IsBoxingAllowed() bool {
	return v.Boxingallowed == 1
}

// Validates the object
func (v *NativeVideo) Valid() (bool, error) {
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
func (v *NativeVideo) WithDefaults() *NativeVideo {
	v.Sequence = 1
	v.Boxingallowed = 1

	return v
}
