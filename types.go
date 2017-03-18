package pg

import "net/http"

const (
	// APISandBoxBaseURL points to test API
	APISandBoxBaseURL = "https://pt01.mul-pay.jp/payment/%s.idPass"
	// APILiveBaseURL points to production API
	APILiveBaseURL = "https://p01.mul-pay.jp/payment/%s.idPass"

	Deleted = "DeleteFlag"
)

type (
	// Client represents payment gateway API client
	Client struct {
		client     *http.Client
		SiteID     string
		SitePass   string
		ShopID     string
		ShopPass   string
		APIBaseURL string
	}

	// Member represents member of payment service
	Member struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Deleted bool   `json:"deleted"`
	}

	// MemberResponse represents response of member request
	MemberResponse struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Deleted int    `json:"deleted"`
	}

	// CreditCard represents credit card for member
	CreditCard struct {
		Member         *Member `json:"id"`
		SequenceNumber int     `json:"sequence_number"`
		SequenceMode   int
		Default        bool   `json:"default"`
		Name           string `json:"name"`
		Number         string `json:"number"`
		Expire         string `json:"expire"`
		HolderName     string `json:"holder_name"`
		Token          string `json:"token,omitempty"`
		Deleted        bool   `json:"deleted"`
	}

	// CreditCardResponse represents response of card request
	CreditCardResponse struct {
		SequenceNumber int    `json:"sequence_number"`
		Default        bool   `json:"default"`
		Name           string `json:"name"`
		Number         string `json:"number"`
		Expire         string `json:"expire"`
		HolderName     string `json:"holder_name"`
		Deleted        bool   `json:"deleted"`
	}

	// CreditCards represents credit card list
	CreditCards struct {
		Items []CreditCard `json:"items"`
		Count int          `json:"count"`
	}

	// Entry represents payment transaction
	Entry struct {
		OrderID string
		JobCode string
		Amount  int
		Tax     int
	}

	// EntryResponse represents response of entry transaction
	EntryResponse struct {
		AccessID   string
		AccessPass string
	}

	// Charge represents payment with sequence number of credit card
	Charge struct {
		CreditCard
		Entry
		EntryResponse
	}

	// ChargeResponse represents response of charge
	ChargeResponse struct {
	}

	// CardCharge represents payment with credit card number
	CardCharge struct {
		Entry
		EntryResponse
		Method     int
		PayTimes   int
		CardNumber string
		Expire     string
		Token      string
	}

	// CardChargeResponse represents response of card charge
	CardChargeResponse struct {
	}

	// CardRecurring represents recurring object of member
	CardRecurring struct {
		ID              string
		Tax             int
		ChargeDay       string
		ChargeMonth     string
		ChargeStartDate string
		ChargeStopDate  string
		RegistType      string
		Member
		// CreditCard
		// Entry
	}

	// CardRecurringResponse represents response of register card recurring
	CardRecurringResponse struct {
		CardRecurring
		Member
		CreditCard
		NextChargeDate string
	}

	// ErrorResponse represents single error response
	ErrorResponse struct {
		Code string
		Info string
	}

	// ErrorResponses represents list of error response
	ErrorResponses struct {
		Items []ErrorResponse `json:"items"`
		Count int             `json:"count"`
	}
)
