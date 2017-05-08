package pg

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const (
	RecurringID     = "RecurringID"
	ChargeDay       = "ChargeDay"
	ChargeMonth     = "ChargeMonth"
	ChargeStartDate = "ChargeStartDate"
	ChargeStopDate  = "ChargeStopDate"
	RegistType      = "RegistType"
)

func (r *Recurring) ToValues() url.Values {
	vs := url.Values{}
	vs.Add(RecurringID, r.ID)
	vs.Add(Amount, strconv.Itoa(r.Amount))
	vs.Add(Tax, strconv.Itoa(r.Tax))
	vs.Add(ChargeDay, r.ChargeDay)
	vs.Add(ChargeMonth, r.ChargeMonth)
	vs.Add(ChargeStartDate, r.ChargeStartDate)
	vs.Add(ChargeStopDate, r.ChargeStopDate)
	vs.Add(RegistType, r.RegistType)
	vs.Add(MemberID, r.Member.ID)
	return vs
}

// RegisterMemberRecurring store recurring object for member
func (c *Client) RegisterMemberRecurring(r *Recurring) (*RecurringResponse, *ErrorResponses) {
	v := c.mergeValues(r.ToValues())
	bodyString, err := c.post(fmt.Sprintf(c.APIBaseURL, "RegisterRecurringCredit"), strings.NewReader(v.Encode()))

	if err != nil {
		return nil, nil
	}

	rr, errors := ConvertToMemberRrecurring(bodyString)

	if errors != nil && errors.Count > 0 {
		return nil, errors
	}

	return rr, nil
}

// UnregisterRecurring cancel recurring object
func (c *Client) UnregisterRecurring(r *Recurring) (*RecurringResponse, *ErrorResponses) {
	//only ShopID,ShopPass,RecurringID
	v := url.Values{}
	v.Add(ShopID, c.ShopID)
	v.Add(ShopPass, c.ShopPass)
	v.Add(RecurringID, r.ID)

	bodyString, err := c.post(fmt.Sprintf(c.APIBaseURL, "UnregisterRecurring"), strings.NewReader(v.Encode()))

	if err != nil {
		return nil, nil
	}

	rr, errors := ConvertToMemberRrecurring(bodyString)

	if errors != nil && errors.Count > 0 {
		return nil, errors
	}

	return rr, nil
}

// SearchRecurring search recurring object.
func (c *Client) SearchRecurring(r *Recurring) (*RecurringResponse, *ErrorResponses) {
	//only ShopID,ShopPass,RecurringID
	v := url.Values{}
	v.Add(ShopID, c.ShopID)
	v.Add(ShopPass, c.ShopPass)
	v.Add(RecurringID, r.ID)

	bodyString, err := c.post(fmt.Sprintf(c.APIBaseURL, "SearchRecurring"), strings.NewReader(v.Encode()))

	if err != nil {
		return nil, nil
	}

	rr, errors := ConvertToMemberRrecurring(bodyString)

	if errors != nil && errors.Count > 0 {
		return nil, errors
	}

	return rr, nil
}

func (c *Client) SearchRecurringResult(r *Recurring) (*RecurringResultResponse, *ErrorResponses) {
	//only ShopID,ShopPass,RecurringID
	v := url.Values{}
	v.Add(ShopID, c.ShopID)
	v.Add(ShopPass, c.ShopPass)
	v.Add(RecurringID, r.ID)

	bodyString, err := c.post(fmt.Sprintf(c.APIBaseURL, "SearchRecurringResult"), strings.NewReader(v.Encode()))

	if err != nil {
		return nil, nil
	}

	//fixme
	rr, errors := ConvertToMemberRrecurring(bodyString)

	if errors != nil && errors.Count > 0 {
		return nil, errors
	}

	return rr, nil
}
