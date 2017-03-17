package pg

import "net/url"

// ToValues return ID and Name key-value
func (m *Member) ToValues() url.Values {
	values := url.Values{}
	values.Add("MemberID", m.ID)
	values.Add("MemberName", m.Name)
	return values
}
