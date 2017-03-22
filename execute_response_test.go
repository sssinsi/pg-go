package pg

import (
	"fmt"
	"testing"
	"time"
)

func TestConvertToExecuteReponse(t *testing.T) {
	rID := "order_id"
	rForward := "1234567"
	rMethod := "1"
	rPayTimes := 10
	rApprove := "8901234"
	rTranID := "transaction_id12345678901234"
	rTranDate := time.Date(2017, 1, 2, 3, 4, 5, 0, time.UTC) //"20170102030405"
	rCheckString := "length32string123456789012345678"
	const layout = "20060102150405"

	s := fmt.Sprintf(
		"%s=%s&%s=%s&%s=%s&%s=%d&%s=%s&%s=%s&%s=%s&%s=%s",
		OrderID, rID,
		Forward, rForward,
		Method, rMethod,
		PayTimes, rPayTimes,
		Approve, rApprove,
		TransactionID, rTranID,
		TransactionDate, rTranDate.Format(layout),
		CheckString, rCheckString,
	)

	er, errors := ConvertToExecuteResponse(s)
	if errors != nil && errors.Count > 0 {
		t.Fatal("errors should be nil")
	}
	if er.OrderID != rID {
		t.Fatal(fmt.Sprintf("OrderID should be %s", rID))
	}
	if er.Forward != rForward {
		t.Fatal(fmt.Sprintf("Forward should be %s", rForward))
	}
	if er.Method != rMethod {
		t.Fatal(fmt.Sprintf("Method should be %s", rMethod))
	}
	if er.PayTimes != rPayTimes {
		t.Fatal(fmt.Sprintf("PayTimes should be %d", rPayTimes))
	}
	if er.Approve != rApprove {
		t.Fatal(fmt.Sprintf("Approve should be %s", rApprove))
	}
	if er.TransactionID != rTranID {
		t.Fatal(fmt.Sprintf("TransactionID should be %s", rTranID))
	}
	if er.TransactionDate != rTranDate {
		t.Fatal(fmt.Sprintf("TransactionDate(%s) should be %s", er.TransactionDate, rTranDate))
	}
	if er.CheckString != rCheckString {
		t.Fatal(fmt.Sprintf("CheckString should be %s", rCheckString))
	}
}
