package openrtb

// The "device" object provides information pertaining to the device including its hardware,
// platform, location, and carrier. This device can refer to a mobile handset, a desktop computer,
// set top box or other digital device.
type Device struct {
	Dnt            int        `json:"dnt,omitempty"` // "1": Do not track
	Ua             string     `json:"ua,omitempty"`  // User agent
	Ip             string     `json:"ip,omitempty"`  // IPv4
	Geo            *Geo       `json:"geo,omitempty"`
	Didsha1        string     `json:"didsha1,omitempty"`  // SHA1 hashed device ID
	Didmd5         string     `json:"didmd5,omitempty"`   // MD5 hashed device ID
	Dpidsha1       string     `json:"dpidsha1,omitempty"` // SHA1 hashed platform device ID
	Dpidmd5        string     `json:"dpidmd5,omitempty"`  // MD5 hashed platform device ID
	Macsha1        string     `json:"macsha1,omitempty"`  // SHA1 hashed device ID; IMEI when available, else MEID or ESN
	Macmd5         string     `json:"macmd5,omitempty"`   // MD5 hashed device ID; IMEI when available, else MEID or ESN
	Ipv6           string     `json:"ipv6,omitempty"`     // IPv6
	Carrier        string     `json:"carrier,omitempty"`  // Carrier or ISP derived from the IP address
	Language       string     `json:"language,omitempty"` // Browser language
	Make           string     `json:"make,omitempty"`     // Device make
	Model          string     `json:"model,omitempty"`    // Device model
	Os             string     `json:"os,omitempty"`       // Device OS
	Osv            string     `json:"osv,omitempty"`      // Device OS version
	Js             int        `json:"js,omitempty"`       // Javascript status ("0": Disabled, "1": Enabled)
	Connectiontype int        `json:"connectiontype,omitempty"`
	Devicetype     int        `json:"devicetype,omitempty"`
	Flashver       string     `json:"flashver,omitempty"` // Flash version
	Ifa            string     `json:"ifa,omitempty"`      // Native identifier for advertisers
	Ext            Extensions `json:"ext,omitempty"`
}

// Returns the DNT status, with default fallback
func (d *Device) IsDnt() bool {
	return d.Dnt == 1
}

// Returns the JS status, with default fallback
func (d *Device) IsJs() bool {
	return d.Js == 1
}
