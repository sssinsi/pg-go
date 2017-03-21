package pg

import (
	"fmt"
	"testing"
)

func TestConvertToMemberResponse(t *testing.T) {
	mID := "member_id"
	mName := "member_name"
	mDelete := 0
	s := fmt.Sprintf("%s=%s&%s=%s&%s=%d", MemberID, mID, MemberName, mName, Deleted, mDelete)
	mr, errors := ConvertToMemberResponse(s)
	if errors != nil && errors.Count > 0 {
		t.Fatal("errors should be nil")
	}
	if mr.ID != mID {
		t.Fatal(fmt.Sprintf("ID should be %s", mID))
	}
	if mr.Name != mName {
		t.Fatal(fmt.Sprintf("Name should be %s", mName))
	}
	if mr.Deleted != mDelete {
		t.Fatal(fmt.Sprintf("Deleted should be %d", mDelete))
	}
}
