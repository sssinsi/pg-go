package pg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
	values.Add("SiteID", c.SiteID)
	values.Add("SitePass", c.SitePass)
	values.Add("ShopID", c.ShopID)
	values.Add("ShopPass", c.ShopPass)
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

// SaveMember store member account
func (c *Client) SaveMember(m Member) (*MemberResponse, *ErrorResponses) {
	v := c.mergeValues(m.ToValues())
	resp, err := c.post(
		fmt.Sprintf(c.APIBaseURL, "SaveMember"),
		strings.NewReader(v.Encode()),
	)

	if err != nil {
		return nil, nil
	}

	mr, errors := ConvertToMemberResponse(resp)
	if errors != nil && errors.Count > 0 {
		return nil, errors
	}
	return mr, nil
}

// SearchMember search member information from payment gateway
func (c *Client) SearchMember(m Member) (*MemberResponse, *ErrorResponses) {
	v := c.mergeValues(m.ToValues())
	bodyString, err := c.post(fmt.Sprintf(c.APIBaseURL, "SearchMember"), strings.NewReader(v.Encode()))

	if err != nil {
		return nil, nil
	}

	mr, errors := ConvertToMemberResponse(bodyString)
	if errors != nil && errors.Count > 0 {
		return nil, errors
	}
	return mr, nil
}

// SaveCard store credit card to member
func (c *Client) SaveCard(card *CreditCard) (*CreditCardResponse, *ErrorResponses) {
	v := c.mergeValues(card.ToValues())
	bodyString, err := c.post(fmt.Sprintf(c.APIBaseURL, "SaveCard"), strings.NewReader(v.Encode()))

	if err != nil {
		return nil, nil
	}

	cr, errors := ConvertToCreditCardResponse(bodyString)
	if errors != nil && errors.Count > 0 {
		return nil, errors
	}

	return cr, nil
}

// SearchCard return single credit card response
func (c *Client) SearchCard(card *CreditCard) (*CreditCardResponse, *ErrorResponses) {
	v := c.mergeValues(card.ToValues())
	bodyString, err := c.post(fmt.Sprintf(c.APIBaseURL, "SearchCard"), strings.NewReader(v.Encode()))

	if err != nil {
		return nil, nil
	}

	cr, errors := ConvertToCreditCardResponse(bodyString)
	if errors != nil && errors.Count > 0 {
		return nil, errors
	}

	return cr, nil
}
