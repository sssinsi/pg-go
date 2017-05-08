package pg

import (
	"net/url"
	"strconv"
)

const (
	NextChargeDate = "NextChargeDate"
)

func ConvertToMemberRecurringResponse(s string) (*RecurringResponse, *ErrorResponses) {
	vs, err := url.ParseQuery(s)
	if err != nil {
		return nil, nil
	}
	rr := &RecurringResponse{}
	errors := ConvertToErrorResponses(vs)

	//ignore ShopID and SiteID
	if vs.Get(RecurringID) != "" {
		rr.Recurring.ID = vs.Get(RecurringID)
		rr.Recurring.Amount, _ = strconv.Atoi(vs.Get(Amount))
		rr.Recurring.Tax, _ = strconv.Atoi(vs.Get(Tax))
		rr.Recurring.ChargeDay = vs.Get(ChargeDay)
		rr.Recurring.ChargeMonth = vs.Get(ChargeMonth)
		rr.Recurring.ChargeStartDate = vs.Get(ChargeStartDate)
		rr.Recurring.ChargeStopDate = vs.Get(ChargeStopDate)
		rr.NextChargeDate = vs.Get(NextChargeDate)
		rr.Member.ID = vs.Get(MemberID)
		rr.CreditCard.Number = vs.Get(Number)
		rr.CreditCard.Expire = vs.Get(Expire)
	}
	return rr, errors
}
