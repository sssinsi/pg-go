package pg

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	SequenceNumber = "CardSeq"
	SequenceMode   = "SeqMode"
	Default        = "DefaultFlag"
	Name           = "CardName"
	Number         = "CardNo"
	Expire         = "Expire"
	HolderName     = "HolderName"
	Token          = " Token"
	// deleted        = "DeleteFlag"
)

// ToValues return key-value
func (c *CreditCard) ToValues() url.Values {
	var d string
	if c.Default {
		d = "1"
	} else {
		d = "0"
	}

	vs := url.Values{}
	vs.Add(SequenceNumber, strconv.Itoa(c.SequenceNumber))
	vs.Add(SequenceMode, strconv.Itoa(c.SequenceMode))
	vs.Add(Default, d)
	vs.Add(Name, c.Name)
	vs.Add(Number, c.Number)
	vs.Add(Expire, c.Expire)
	vs.Add(HolderName, c.HolderName)

	for k, value := range c.Member.ToValues() {
		for _, v := range value {
			vs.Add(k, v)
		}
	}
	return vs
}

// SaveCard store credit card to member
func (c *Client) SaveCard(card *CreditCard) (*CreditCardResponse, *ErrorResponses) {
	v := c.mergeValues(card.ToValues())
	bodyString, err := c.post("SaveCard", strings.NewReader(v.Encode()))

	if err != nil {
		return nil, nil
	}

	cr, errors := ConvertToCreditCardResponse(bodyString)
	if errors != nil && errors.Count > 0 {
		return nil, errors
	}

	return cr, nil
}

// SearchCard return single credit card response
func (c *Client) SearchCard(card *CreditCard) (*CreditCardResponse, *ErrorResponses) {
	v := c.mergeValues(card.ToValues())
	bodyString, err := c.post("SearchCard", strings.NewReader(v.Encode()))

	if err != nil {
		return nil, nil
	}

	cr, errors := ConvertToCreditCardResponse(bodyString)
	if errors != nil && errors.Count > 0 {
		return nil, errors
	}

	return cr, nil
}

// CardCharge store charge by credit card.
func (c *Client) CardCharge(card *CreditCard, amount, tax int) (*CardChargeResponse, *ErrorResponses) {
	e := NewEntry("", "1", amount, tax)
	//entry
	er, errors := c.entry(e)
	if errors != nil && errors.Count > 0 {
		return nil, errors
	}
	charge := &Charge{card, e, er}
	//exec
	exr, errors := c.execute(charge)
	if errors != nil && errors.Count > 0 {
		return nil, errors
	}

	return &CardChargeResponse{er, exr}, nil
}

func (c *Client) entry(e *Entry) (*EntryResponse, *ErrorResponses) {
	vs := c.mergeValues(e.ToValues())
	bodyString, err := c.post("EntryTran", strings.NewReader(vs.Encode()))
	if err != nil {
		return nil, nil
	}

	er, errors := ConvertToEntryResponse(bodyString)
	if errors != nil && errors.Count > 0 {
		return nil, errors
	}

	return er, nil
}

func (c *Client) execute(charge *Charge) (*ExecuteResponse, *ErrorResponses) {
	vs := c.ToValues()
	vs.Add(MemberID, charge.CreditCard.Member.ID)
	vs.Add(SequenceNumber, strconv.Itoa(charge.CreditCard.SequenceNumber))
	vs.Add(OrderID, charge.Entry.OrderID)
	vs.Add(AccessID, charge.EntryResponse.AccessID)
	vs.Add(AccessPass, charge.EntryResponse.AccessPass)
	vs.Add("Method", "1")

	bodyString, err := c.post("ExecTran", strings.NewReader(vs.Encode()))
	if err != nil {
		return nil, nil
	}

	exr, errors := ConvertToExecuteResponse(bodyString)
	if errors != nil && errors.Count > 0 {
		return nil, errors
	}

	return exr, nil
}
