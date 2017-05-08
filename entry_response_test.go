package pg

import (
	"fmt"
	"testing"
)

func TestConvertToEntryResponse(t *testing.T) {
	rAccessID := "access_id_is_32_length_strings_!"
	rAccessPass := "access_pass_is_32_length_strings"
	s := fmt.Sprintf("%s=%s&%s=%s", AccessID, rAccessID, AccessPass, rAccessPass)
	er, errors := ConvertToEntryResponse(s)
	if errors != nil && errors.Count > 0 {
		t.Fatal("errors should be nil")
	}
	if er.AccessID != rAccessID {
		t.Fatal(fmt.Sprintf("AccessID should be %s", rAccessID))
	}
	if er.AccessPass != rAccessPass {
		t.Fatal(fmt.Sprintf("AccessPass should be %s", rAccessPass))
	}
}
