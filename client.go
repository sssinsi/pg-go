package pg

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	SiteID   = "SiteID"
	SitePass = "SitePass"
	ShopID   = "ShopID"
	ShopPass = "ShopPass"
)

func NewClient(siteID, sitePass, shopID, shopPass string, sandBox bool) (*Client, error) {
	if siteID == "" || sitePass == "" || shopID == "" || shopPass == "" {
		return nil, errors.New("NewClient error")
	}
	var url string
	if sandBox {
		url = APISandBoxBaseURL
	} else {
		url = APILiveBaseURL
	}
	return &Client{
		client:     &http.Client{},
		SiteID:     siteID,
		SitePass:   sitePass,
		ShopID:     shopID,
		ShopPass:   shopPass,
		APIBaseURL: url,
	}, nil
}

func (c *Client) ToValues() url.Values {
	values := url.Values{}
	values.Add(SiteID, c.SiteID)
	values.Add(SitePass, c.SitePass)
	values.Add(ShopID, c.ShopID)
	values.Add(ShopPass, c.ShopPass)
	return values
}

func (c *Client) mergeValues(values url.Values) url.Values {
	vs := c.ToValues()
	for key, value := range values {
		for _, v := range value {
			vs.Add(key, v)
		}
	}
	return vs
}

func (c *Client) post(url string, reader *strings.Reader) (string, error) {
	req, err := http.NewRequest(
		"POST",
		url,
		reader,
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	bodyString := string(bodyBytes)

	return bodyString, nil
}
