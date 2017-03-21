package pg

import "net/url"

func ConvertToMemberResponse(s string) (*MemberResponse, *ErrorResponses) {
	vs, err := url.ParseQuery(s)
	if err != nil {
		return nil, nil
	}
	mr := &MemberResponse{}
	ers := ConvertToErrorResponses(vs)

	if vs.Get(MemberID) != "" {
		del := vs.Get(Deleted) == "1"
		mr.ID = vs.Get(MemberID)
		mr.Name = vs.Get(MemberName)
		mr.Deleted = del
	}

	return mr, ers
}
