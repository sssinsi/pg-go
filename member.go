package pg

import "net/url"

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
