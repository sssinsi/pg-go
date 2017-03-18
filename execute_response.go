package pg

import (
	"net/url"
	"strconv"
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
		er.OrderID = vs.Get(OrderID)
		er.Forward = vs.Get(Forward)
		er.Method = vs.Get(Method)
		er.PayTimes, _ = strconv.Atoi(vs.Get(PayTimes))
		er.Approve = vs.Get(Approve)
		er.TransactionID = vs.Get(TransactionID)
		er.CheckString = vs.Get(CheckString)
	}
	return er, ers
}
