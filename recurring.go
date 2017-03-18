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

// RegistMemberRecurring store recurring object for member
func (c *Client) RegistMemberRecurring(r *Recurring) (*RecurringResponse, *ErrorResponses) {
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
