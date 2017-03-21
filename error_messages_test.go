package pg

import (
	"fmt"
	"net/url"
	"testing"
)

func TestConvertToErrorResponses(t *testing.T) {
	vs := url.Values{}
	vs.Add(ErrorCode, "E01|E02|E03")
	vs.Add(ErrorInfo, "E01001|E02001|E03001")
	ers := ConvertToErrorResponses(vs)
	if ers.Count != 3 {
		t.Fatal(fmt.Sprintf("Count should be 3"))
	}
	if ers.Items[0].Code != "E01" {
		t.Fatal(fmt.Sprintf("Item0 should be E01"))
	}
	if ers.Items[0].Info != "E01001" {
		t.Fatal(fmt.Sprintf("Info0 should be E01001"))
	}
	if ers.Items[1].Code != "E02" {
		t.Fatal(fmt.Sprintf("Item1 should be E02"))
	}
	if ers.Items[1].Info != "E02001" {
		t.Fatal(fmt.Sprintf("Info1 should be E02001"))
	}
	if ers.Items[2].Code != "E03" {
		t.Fatal(fmt.Sprintf("Item2 should be E01"))
	}
	if ers.Items[2].Info != "E03001" {
		t.Fatal(fmt.Sprintf("Info2 should be E03001"))
	}
}
