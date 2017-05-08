package pg

import (
	"fmt"
	"testing"
)

func TetstConvetToCreditCardResponse(t *testing.T) {
	cName := "card_name"
	cNumber := "4242424242424242"
	cExpire := "2203"
	cHoldName := "hold name"
	cDefault := 1
	cDeleted := 0
	cSeqNo := 1
	s := fmt.Sprintf(
		"%s=%s&%s=%s&%s=%s&%s=%s&%s=%d&%s=%d&%s=%d",
		Name, cName,
		Number, cNumber,
		Expire, cExpire,
		HolderName, cHoldName,
		Default, cDefault,
		Deleted, cDeleted,
		SequenceNumber, cSeqNo,
	)
	cr, ers := ConvertToCreditCardResponse(s)
	if ers != nil && ers.Count > 0 {
		t.Fatal("errors should be nil")
	}
	if cr.Name != cName {
		t.Fatal(fmt.Sprintf("Name should be %s", cName))
	}
	if cr.Number != cNumber {
		t.Fatal(fmt.Sprintf("Number should be %s", cNumber))
	}
	if cr.Expire != cExpire {
		t.Fatal(fmt.Sprintf("Expire should be %s", cExpire))
	}
	if cr.HolderName != cHoldName {
		t.Fatal(fmt.Sprintf("HolderName should be %s", cHoldName))
	}
	if !cr.Default {
		t.Fatal("Deleted should be true")
	}
	if cr.Deleted {
		t.Fatal("Deleted should be false")
	}
	if cr.SequenceNumber != cSeqNo {
		t.Fatal(fmt.Sprintf("SequenceNumber should be %d", cSeqNo))
	}
}
