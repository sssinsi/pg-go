package pg

import (
	"net/url"
	"strconv"
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
