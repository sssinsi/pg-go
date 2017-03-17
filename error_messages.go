package pg

import (
	"net/url"
	"strings"
)

func ConvertToErrorResponses(vs url.Values) *ErrorResponses {
	ers := &ErrorResponses{}
	if vs.Get("ErrCode") != "" {
		errCodeStrings := strings.Split(vs.Get("ErrCode"), "|")
		errInfoStrings := strings.Split(vs.Get("ErrInfo"), "|")
		for i := range errCodeStrings {
			er := ErrorResponse{
				Code: errCodeStrings[i],
				Info: errInfoStrings[i],
			}
			ers.Items = append(ers.Items, er)
			ers.Count++
		}
	}
	return ers
}
