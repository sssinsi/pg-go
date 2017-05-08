package pg

import (
	"net/url"
	"strings"
)

const (
	// MemberID represents key of Member ID
	MemberID = "MemberId"
	// MemberName represents key of Member Name
	MemberName = "MemberName"
)

// ToValues return ID and Name key-value
func (m *Member) ToValues() url.Values {
	values := url.Values{}
	values.Add(MemberID, m.ID)
	values.Add(MemberName, m.Name)
	return values
}

// SaveMember store member account
func (c *Client) SaveMember(m Member) (*MemberResponse, *ErrorResponses) {
	v := c.mergeValues(m.ToValues())
	resp, err := c.post(
		"SaveMember",
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
	bodyString, err := c.post("SearchMember", strings.NewReader(v.Encode()))

	if err != nil {
		return nil, nil
	}

	mr, errors := ConvertToMemberResponse(bodyString)
	if errors != nil && errors.Count > 0 {
		return nil, errors
	}
	return mr, nil
}
