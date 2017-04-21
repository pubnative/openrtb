package openrtb

import (
	"errors"
)

// ID, Impid and Price are required; all other optional.
// If the bidder wins the impression, the exchange calls notice URL (nurl)
// a) to inform the bidder of the win;
// b) to convey certain information using substitution macros.
// Adomain can be used to check advertiser block list compliance.
// Cid can be used to block ads that were previously identified as inappropriate.
// Substitution macros may allow a bidder to use a static notice URL for all of its bids.
type Bid struct {
	Id      string     `json:"id"`
	Impid   string     `json:"impid"`             // Required string ID of the impression object to which this bid applies.
	Price   float32    `json:"price"`             // Bid price in CPM. Suggests using integer math for accounting to avoid rounding errors.
	Adid    string     `json:"adid,omitempty"`    // References the ad to be served if the bid wins.
	Nurl    string     `json:"nurl,omitempty"`    // Win notice URL.
	Adm     string     `json:"adm,omitempty"`     // Actual ad markup. XHTML if a response to a banner object, or VAST XML if a response to a video object.
	Adomain []string   `json:"adomain,omitempty"` // Advertiser’s primary or top-level domain for advertiser checking; or multiple if imp rotating.
	Bundle  string     `json:"bundle,omitempty"`  // Bundle or package name.
	Iurl    string     `json:"iurl,omitempty"`    // Sample image URL.
	Cid     string     `json:"cid,omitempty"`     // Campaign ID that appears with the Ad markup.
	Crid    string     `json:"crid,omitempty"`    // Creative ID for reporting content issues or defects. This could also be used as a reference to a creative ID that is posted with an exchange.
	Attr    []int      `json:"attr,omitempty"`    // Array of creative attributes.
	DealId  string     `json:"deal_id,omitempty"` // DealID extension of private marketplace deals
	Cat     string     `json:"cat,omitempty"`     // IAB content categories of the creative. Refer to List 5.1.
	Ext     Extensions `json:"ext,omitempty"`
}

// Validation errors
var (
	ErrInvalidBidID    = errors.New("openrtb response: bid is missing ID")
	ErrInvalidBidImpID = errors.New("openrtb response: bid is missing impression ID")
	ErrInvalidBidPrice = errors.New("openrtb response: bid is missing price")
)

// Validate Bid required attributes
func (bid *Bid) Valid() (bool, error) {

	if len(bid.Id) == 0 {
		return false, ErrInvalidBidID
	} else if len(bid.Impid) == 0 {
		return false, ErrInvalidBidImpID
	} else if bid.Price == 0 {
		return false, ErrInvalidBidPrice
	}

	return true, nil
}
