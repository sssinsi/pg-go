package pg

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	siteID := "site_id"
	sitePass := "site_pass"
	shopID := "shop_id"
	shopPass := "shop_pass"
	sandBox := true
	c, err := NewClient(siteID, sitePass, shopID, shopPass, sandBox)

	if err != nil {
		t.Fatal("err should be null")
	}

	if c.SiteID != siteID {
		t.Fatal(fmt.Sprintf("SiteID should be %s", siteID))
	}
	if c.SitePass != sitePass {
		t.Fatal(fmt.Sprintf("SitePass should be %s", sitePass))
	}
	if c.ShopID != shopID {
		t.Fatal(fmt.Sprintf("ShopID should be %s", shopID))
	}
	if c.ShopPass != shopPass {
		t.Fatal(fmt.Sprintf("ShopPass should be %s", shopPass))
	}
	if c.APIBaseURL != APISandBoxBaseURL {
		t.Fatal(fmt.Sprintf("APIBaseURL should be %s", APISandBoxBaseURL))
	}
}
