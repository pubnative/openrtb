package openrtb

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {
	var subject *Request

	BeforeEach(func() {
		subject = new(Request)
	})

	It("should should parse requests", func() {
		req, err := ParseRequest(bytes.NewBuffer(testFixtures.expandableCreative))
		Expect(err).NotTo(HaveOccurred())

		Expect(req.Imp).To(Equal([]Impression{
			{
				Id: "1",
				Banner: &Banner{
					W:        300,
					H:        250,
					Wmin:     300,
					Hmin:     250,
					Pos:      1,
					Topframe: 0,
					Battr:    []int{13},
					Expdir:   []int{2, 4},
				},
				Instl:        0,
				Bidfloor:     0,
				Bidfloorcur:  "USD",
				Secure:       0,
				Iframebuster: []string{"vendor1.com", "vendor2.com"},
			},
		}))

		Expect(req.Site).To(Equal(&Site{
			Id:            "1345135123",
			Name:          "Site ABCD",
			Domain:        "siteabcd.com",
			Page:          "http://siteabcd.com/page.htm",
			Privacypolicy: 1,
			Ref:           "http://referringsite.com/referringpage.htm",
			Publisher:     &Publisher{Id: "pub12345", Name: "Publisher A"},
		}))

		Expect(req.Device).To(Equal(&Device{
			Dnt:            0,
			Ua:             "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20110319 Firefox/3.6.16",
			Ip:             "64.124.253.1",
			Os:             "OS X",
			Js:             1,
			Connectiontype: 0,
			Devicetype:     0,
			Flashver:       "10.1",
		}))

		Expect(req.User).To(Equal(&User{
			Id:       "456789876567897654678987656789",
			Buyeruid: "545678765467876567898765678987654",
			Data: []Data{
				{
					Id:   "6",
					Name: "Data Provider 1",
					Segment: []Segment{
						{Id: "12341318394918", Name: "auto intenders"},
						{Id: "1234131839491234", Name: "auto enthusiasts"},
						{Id: "23423424", Name: "data-provider1-age", Value: "30-40"},
					},
				},
			},
		}))
	})

	Describe("ParseRequestBytes()", func() {
		It("should return blank request with defaults when blank", func() {
			req, err := ParseRequestBytes([]byte("{}"))
			Expect(err).NotTo(HaveOccurred())
			Expect(req).To(BeAssignableToTypeOf(&Request{}))
		})

		It("should return accordingly when with simple banner", func() {
			req, err := ParseRequestBytes(testFixtures.simpleBanner)
			Expect(err).NotTo(HaveOccurred())
			Expect(req).To(BeAssignableToTypeOf(&Request{}))

			Expect(req.At).To(Equal(2))
			Expect(req.Id).To(Equal("1234534625254"))

			Expect(len(req.Badv)).To(Equal(2))
			Expect(len(req.Bcat)).To(Equal(0))
			Expect(len(req.Imp)).To(Equal(1))

			Expect(req.Imp[0].Banner.W).To(Equal(300))
			Expect(req.Imp[0].Banner.H).To(Equal(250))
			Expect(req.Site.Name).To(Equal("Site ABCD"))
			Expect(req.Site.Publisher.Name).To(Equal("Publisher A"))
			Expect(req.Device.Ip).To(Equal("64.124.253.1"))
			Expect(req.User.Buyeruid).To(Equal("5df678asd8987656asdf78987654"))
		})

		It("should return accordingly when with expandable creatives", func() {
			req, err := ParseRequestBytes(testFixtures.expandableCreative)
			Expect(err).NotTo(HaveOccurred())
			Expect(req).To(BeAssignableToTypeOf(&Request{}))

			Expect(req.At).To(Equal(2))
			Expect(req.Tmax).To(Equal(120))
			Expect(req.Imp[0].Banner.Expdir).To(Equal([]int{2, 4}))
			Expect(req.Site.Privacypolicy).To(Equal(1))
			Expect(req.Device.Flashver).To(Equal("10.1"))
			Expect(req.User.Data[0].Id).To(Equal("6"))
			Expect(req.User.Data[0].Segment[2].Id).To(Equal("23423424"))

			Expect(len(req.User.Data)).To(Equal(1))
			Expect(len(req.User.Data[0].Segment)).To(Equal(3))
		})
	})

	Describe("Valid()", func() {
		var (
			site       *Site
			app        *App
			impression *Impression
			banner     *Banner
		)

		BeforeEach(func() {
			site = new(Site)
			app = new(App)
			impression = new(Impression)
			banner = new(Banner)
		})

		It("should return error messages when attributes missing", func() {
			ok, err := subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: request ID missing"))

			subject.Id = "RAND_ID" // With ID
			ok, err = subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: no impressions"))

			subject.Site = site // With Site
			ok, err = subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: no impressions"))

			subject.App = app // With App
			ok, err = subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: no impressions"))

			// With Impression
			impression.Id = "IMPID"
			impression.Banner = banner
			subject.Imp = []Impression{*impression}
			ok, err = subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: request has site and app"))

			// with valid attrs
			subject.App = nil
			ok, err = subject.Valid()
			Expect(err).NotTo(HaveOccurred())
			Expect(ok).To(BeTrue())
		})

	})

	Describe("WithDefaults()", func() {

		var (
			site       *Site
			app        *App
			device     *Device
			impression *Impression
			banner     *Banner
			video      *Video
		)

		BeforeEach(func() {
			site = new(Site)
			app = new(App)
			device = new(Device)
			impression = new(Impression)
			banner = new(Banner)
			video = new(Video)

			impression.Banner = banner
			impression.Video = video
			subject.Site = site
			subject.App = app
			subject.Device = device
			subject.Imp = []Impression{*impression}
		})

		It("should return blank request with default values", func() {
			request := subject.WithDefaults()
			Expect(request.At).To(Equal(2))
			Expect(request.App.Privacypolicy).To(Equal(0))
			Expect(request.App.Paid).To(Equal(0))
			Expect(request.Site.Privacypolicy).To(Equal(0))
			Expect(request.Device.Dnt).To(Equal(0))
			Expect(request.Device.Js).To(Equal(0))
			Expect(request.Device.Connectiontype).To(Equal(CONN_TYPE_UNKNOWN))
			Expect(request.Device.Devicetype).To(Equal(DEVICE_TYPE_UNKNOWN))
			Expect(request.Imp[0].Instl).To(Equal(0))
			Expect(request.Imp[0].Bidfloor).To(Equal(float32(0)))
			Expect(request.Imp[0].Bidfloorcur).To(Equal("USD"))
			Expect(request.Imp[0].Banner.Topframe).To(Equal(0))
			Expect(request.Imp[0].Banner.Pos).To(Equal(AD_POS_UNKNOWN))
			Expect(request.Imp[0].Video.Sequence).To(Equal(1))
			Expect(request.Imp[0].Video.Boxingallowed).To(Equal(1))
			Expect(request.Imp[0].Video.Pos).To(Equal(AD_POS_UNKNOWN))
		})
	})
})
