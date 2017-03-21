package pg

import (
	"fmt"
	"net/url"
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

func TestToValues(t *testing.T) {
	siteID := "site_id"
	sitePass := "site_pass"
	shopID := "shop_id"
	shopPass := "shop_pass"
	sandBox := true
	c, _ := NewClient(siteID, sitePass, shopID, shopPass, sandBox)
	vs := c.ToValues()
	if vs.Get(SiteID) != siteID {
		t.Fatal(fmt.Sprintf("SiteID should be %s", siteID))
	}
	if vs.Get(SitePass) != sitePass {
		t.Fatal(fmt.Sprintf("SitePass should be %s", sitePass))
	}
	if vs.Get(ShopID) != shopID {
		t.Fatal(fmt.Sprintf("ShopID should be %s", shopID))
	}
	if vs.Get(ShopPass) != shopPass {
		t.Fatal(fmt.Sprintf("ShopPass should be %s", shopPass))
	}
}

func TestMergeValues(t *testing.T) {
	siteID := "site_id"
	sitePass := "site_pass"
	shopID := "shop_id"
	shopPass := "shop_pass"
	sandBox := true
	c, _ := NewClient(siteID, sitePass, shopID, shopPass, sandBox)

	k1 := "Key1"
	v1 := "Value1"
	k2 := "Key2"
	v2 := "Value2"
	v := url.Values{}
	v.Add(k1, v1)
	v.Add(k2, v2)
	v.Add(SiteID, "site_id_2")
	vs := c.mergeValues(v)
	if val, ok := vs[SiteID]; ok {
		if len(val) != 2 {
			t.Fatal(fmt.Sprintf("SiteID should have 2 values"))
		}
	}
	if vs.Get(k1) != v1 {
		t.Fatal(fmt.Sprintf("Key1 should be %s", v1))
	}
	if vs.Get(k2) != v2 {
		t.Fatal(fmt.Sprintf("Key2 should be %s", v2))
	}
}
