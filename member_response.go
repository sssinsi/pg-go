package pg

import (
	"net/url"
	"strconv"
)

func ConvertToMemberResponse(s string) (*MemberResponse, *ErrorResponses) {
	vs, err := url.ParseQuery(s)
	if err != nil {
		return nil, nil
	}
	mr := &MemberResponse{}
	ers := ConvertToErrorResponses(vs)

	if vs.Get(MemberID) != "" {
		mr.ID = vs.Get(MemberID)
		mr.Name = vs.Get(MemberName)
		mr.Deleted, _ = strconv.Atoi(vs.Get(Deleted))
	}

	return mr, ers
}
