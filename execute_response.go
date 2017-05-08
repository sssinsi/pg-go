package pg

import (
	"net/url"
	"strconv"
	"time"
)

const (
	Forward         = "Forward"
	Method          = "Method"
	PayTimes        = "PayTimes"
	Approve         = "Approve"
	TransactionID   = "TransactionId"
	TransactionDate = "TranDate"
	CheckString     = "CheckString"
)

func ConvertToExecuteResponse(s string) (*ExecuteResponse, *ErrorResponses) {
	vs, err := url.ParseQuery(s)
	if err != nil {
		return nil, nil
	}
	er := &ExecuteResponse{}
	ers := ConvertToErrorResponses(vs)
	if vs.Get(OrderID) != "" {
		const layout = "20060102150405"
		er.OrderID = vs.Get(OrderID)
		er.Forward = vs.Get(Forward)
		er.Method = vs.Get(Method)
		er.PayTimes, _ = strconv.Atoi(vs.Get(PayTimes))
		er.Approve = vs.Get(Approve)
		er.TransactionID = vs.Get(TransactionID)
		er.TransactionDate, _ = time.Parse(layout, vs.Get(TransactionDate))
		er.CheckString = vs.Get(CheckString)
	}
	return er, ers
}
