package openrtb

import (
	"bytes"
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
	var subject *Response

	BeforeEach(func() {
		subject = new(Response)
	})

	It("should validate", func() {
		ok, err := subject.Valid()
		Expect(err).To(Equal(ErrInvalidResID))
		Expect(ok).To(BeFalse())

		subject.Id = "RES_ID"
		bid := Bid{Id: "BIDID", Impid: "IMPID", Price: 0.01}
		sb := Seatbid{}
		sb.Bid = append(sb.Bid, bid)
		subject.Seatbid = append(subject.Seatbid, sb)

		ok, err = subject.Valid()
		Expect(err).NotTo(HaveOccurred())
		Expect(ok).To(BeTrue())
	})

	It("should generate responses", func() {
		nobid := Response{
			Id:      "32a69c6ba388f110487f9d1e63f77b22d86e916b",
			Nbr:     0,
			Seatbid: []Seatbid{},
		}
		bin, err := json.Marshal(nobid)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(bin)).To(Equal(`{"id":"32a69c6ba388f110487f9d1e63f77b22d86e916b","seatbid":[]}`))
	})

	It("should parse responses", func() {
		resp, err := ParseResponse(bytes.NewBuffer(testFixtures.simpleResponse))
		Expect(err).NotTo(HaveOccurred())

		bid := Bid{
			Id:      "32a69c6ba388f110487f9d1e63f77b22d86e916b",
			Impid:   "32a69c6ba388f110487f9d1e63f77b22d86e916b",
			Price:   0.065445,
			Adid:    "529833ce55314b19e8796116",
			Nurl:    "http://ads.com/win/529833ce55314b19e8796116?won=${auction_price}",
			Adm:     "<iframe src=\"foo.bar\"/>",
			Adomain: []string{},
			Attr:    []int{},
			Cid:     "529833ce55314b19e8796116",
			Crid:    "529833ce55314b19e8796116_1385706446",
		}

		Expect(resp.Seatbid[0].Bid[0]).To(Equal(bid))
	})
})
