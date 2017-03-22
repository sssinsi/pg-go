package pg

import (
	"fmt"
	"testing"
)

func TestConvertToMemberRecurringResponse(t *testing.T) {
	rID := "recurring_id"
	rAmount := 1000
	rTax := 1
	rChargeDay := "02"
	rChargeMonth := "01|02|03|04|05|06|07|08|09|10|11|12"
	rChargeStartDate := "20170401"
	rChargeStopDate := ""
	rNextChargeDate := "20170501"
	rMemberID := "member_id"
	rCreditCardNumber := ""
	rCreditCardExpire := ""
	s := fmt.Sprintf(
		"%s=%s&%s=%d&%s=%d&%s=%s&%s=%s&%s=%s&%s=%s&%s=%s&%s=%s&%s=%s&%s=%s",
		RecurringID, rID,
		Amount, rAmount,
		Tax, rTax,
		ChargeDay, rChargeDay,
		ChargeMonth, rChargeMonth,
		ChargeStartDate, rChargeStartDate,
		ChargeStopDate, rChargeStopDate,
		NextChargeDate, rNextChargeDate,
		MemberID, rMemberID,
		Number, rCreditCardNumber,
		Expire, rCreditCardExpire,
	)
	vs, errors := ConvertToMemberRecurringResponse(s)
	if errors != nil && errors.Count > 0 {
		t.Fatal("errors should be nil")
	}
	if vs.Recurring.ID != rID {
		t.Fatal(fmt.Printf("ID should be %s", rID))
	}
	if vs.Recurring.Amount != rAmount {
		t.Fatal(fmt.Printf("Amount should be %d", rAmount))
	}
	if vs.Recurring.Tax != rTax {
		t.Fatal(fmt.Printf("ID should be %d", rTax))
	}
	if vs.Recurring.ChargeDay != rChargeDay {
		t.Fatal(fmt.Printf("ChargeDay should be %s", rChargeDay))
	}
	if vs.Recurring.ChargeMonth != rChargeMonth {
		t.Fatal(fmt.Printf("ChargeMonth should be %s", rChargeMonth))
	}
	if vs.Recurring.ChargeStartDate != rChargeStartDate {
		t.Fatal(fmt.Printf("ChargeStartDate should be %s", rChargeStartDate))
	}
	if vs.Recurring.ChargeStopDate != rChargeStopDate {
		t.Fatal(fmt.Printf("ChargeStopDate should be %s", rChargeStopDate))
	}
	if vs.NextChargeDate != rNextChargeDate {
		t.Fatal(fmt.Printf("NextChargeDate should be %s", rNextChargeDate))
	}
	if vs.Member.ID != rMemberID {
		t.Fatal(fmt.Printf("MemberID should be %s", rMemberID))
	}
	if vs.CreditCard.Number != rCreditCardNumber {
		t.Fatal(fmt.Printf("Number should be %s", rCreditCardNumber))
	}
	if vs.CreditCard.Expire != rCreditCardExpire {
		t.Fatal(fmt.Printf("Expire should be %s", rCreditCardExpire))
	}
}
