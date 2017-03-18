package pg

import (
	"net/url"
	"strconv"
)

const (
	OrderID = "OrderID"
	JobCode = "JobCd"
	Amount  = "Amount"
)

func NewEntry(orderID, jobCode string, amount, tax int) *Entry {
	return &Entry{
		OrderID: orderID,
		JobCode: jobCode,
		Amount:  amount,
		Tax:     tax,
	}
}

func (e *Entry) ToValues() url.Values {
	vs := url.Values{}
	vs.Add(OrderID, e.OrderID)
	vs.Add(JobCode, e.JobCode)
	vs.Add(Amount, strconv.Itoa(e.Amount))
	return vs
}
