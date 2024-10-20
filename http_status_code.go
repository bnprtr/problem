package problem

//go:generate ./bin/generator http_status_code.go

// list out status codes names from the net/http package that can be mapped to an error.
// the type name must match exactly the status code const in net/http. The generator will
// generate methods, including a StatusCode() int method based on the type name's matching
// const in net/http.

type (
	// 300 level
	StatusMultipleChoices   string
	StatusMovedPermanently  string
	StatusFound             string
	StatusSeeOther          string
	StatusNotModified       string
	StatusUseProxy          string
	StatusTemporaryRedirect string
	StatusPermanentRedirect string

	// 400 level
	StatusBadRequest                   string
	StatusUnauthorized                 string
	StatusPaymentRequired              string
	StatusForbidden                    string
	StatusNotFound                     string
	StatusMethodNotAllowed             string
	StatusNotAcceptable                string
	StatusProxyAuthRequired            string
	StatusRequestTimeout               string
	StatusConflict                     string
	StatusGone                         string
	StatusLengthRequired               string
	StatusPreconditionFailed           string
	StatusRequestEntityTooLarge        string
	StatusRequestURITooLong            string
	StatusUnsupportedMediaType         string
	StatusRequestedRangeNotSatisfiable string
	StatusExpectationFailed            string
	StatusTeapot                       string
	StatusMisdirectedRequest           string
	StatusUnprocessableEntity          string
	StatusLocked                       string
	StatusFailedDependency             string
	StatusTooEarly                     string
	StatusUpgradeRequired              string
	StatusPreconditionRequired         string
	StatusTooManyRequests              string
	StatusRequestHeaderFieldsTooLarge  string
	StatusUnavailableForLegalReasons   string

	// 500 level
	StatusInternalServerError           string
	StatusNotImplemented                string
	StatusBadGateway                    string
	StatusServiceUnavailable            string
	StatusGatewayTimeout                string
	StatusHTTPVersionNotSupported       string
	StatusVariantAlsoNegotiates         string
	StatusInsufficientStorage           string
	StatusLoopDetected                  string
	StatusNotExtended                   string
	StatusNetworkAuthenticationRequired string
)
