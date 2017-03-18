package pg

import "net/url"

const (
	AccessID   = "AccessID"
	AccessPass = "AccessPass"
)

func ConvertToEntryResponse(s string) (*EntryResponse, *ErrorResponses) {
	vs, err := url.ParseQuery(s)
	if err != nil {
		return nil, nil
	}
	er := &EntryResponse{}
	ers := ConvertToErrorResponses(vs)
	if vs.Get(AccessID) != "" {
		er.AccessID = vs.Get(AccessID)
		er.AccessPass = vs.Get(AccessPass)
	}
	return er, ers
}
