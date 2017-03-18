package pg

import (
	"net/url"
	"strconv"
)

func ConvertToCreditCardResponse(s string) (*CreditCardResponse, *ErrorResponses) {
	vs, err := url.ParseQuery(s)
	if err != nil {
		return nil, nil
	}
	cr := &CreditCardResponse{}
	ers := ConvertToErrorResponses(vs)

	if vs.Get(SequenceNumber) != "" {
		d := vs.Get(Default) == "1"
		del := vs.Get(Deleted) == "1"
		cr.SequenceNumber, _ = strconv.Atoi(vs.Get(SequenceNumber))
		cr.Default = d
		cr.Name = vs.Get(Name)
		cr.Number = vs.Get(Number)
		cr.Expire = vs.Get(Expire)
		cr.HolderName = vs.Get(HolderName)
		cr.Deleted = del
	}

	return cr, ers
}
