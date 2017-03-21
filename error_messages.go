package pg

import (
	"net/url"
	"strings"
)

const (
	ErrorCode = "ErrCode"
	ErrorInfo = "ErrInfo"
)

func ConvertToErrorResponses(vs url.Values) *ErrorResponses {
	ers := &ErrorResponses{}
	if vs.Get(ErrorCode) != "" {
		errCodeStrings := strings.Split(vs.Get(ErrorCode), "|")
		errInfoStrings := strings.Split(vs.Get(ErrorInfo), "|")
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
