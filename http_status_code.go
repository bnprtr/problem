package problem

//go:generate ./bin/generator http_status_code.go

type (
	StatusBadRequest string

	StatusUnauthorized string

	StatusPaymentRequired string

	StatusForbidden string

	StatusNotFound string

	StatusMethodNotAllowed string

	StatusNotAcceptable string

	StatusProxyAuthRequired string

	StatusRequestTimeout string

	StatusConflict string

	StatusGone string

	StatusLengthRequired string

	StatusPreconditionFailed string

	StatusRequestEntityTooLarge string

	StatusRequestURITooLong string

	StatusUnsupportedMediaType string

	StatusRequestedRangeNotSatisfiable string

	StatusExpectationFailed string

	StatusTeapot string

	StatusMisdirectedRequest string

	StatusUnprocessableEntity string

	StatusLocked string

	StatusFailedDependency string

	StatusTooEarly string

	StatusUpgradeRequired string

	StatusPreconditionRequired string

	StatusTooManyRequests string

	StatusRequestHeaderFieldsTooLarge string

	StatusUnavailableForLegalReasons string

	StatusInternalServerError string

	StatusNotImplemented string

	StatusBadGateway string

	StatusServiceUnavailable string

	StatusGatewayTimeout string

	StatusHTTPVersionNotSupported string

	StatusVariantAlsoNegotiates string

	StatusInsufficientStorage string

	StatusLoopDetected string

	StatusNotExtended string

	StatusNetworkAuthenticationRequired string
)
