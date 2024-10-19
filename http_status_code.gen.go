package problem

import "net/http"
func (e StatusBadRequest) Error() string {
    return string(e)
}

func (e StatusBadRequest) StatusCode() int {
    return http.StatusBadRequest
}

func (e StatusBadRequest) New(message string) Error[StatusBadRequest] {
    return New(e).WithMessage(message)
}

func (e StatusUnauthorized) Error() string {
    return string(e)
}

func (e StatusUnauthorized) StatusCode() int {
    return http.StatusUnauthorized
}

func (e StatusUnauthorized) New(message string) Error[StatusUnauthorized] {
    return New(e).WithMessage(message)
}

func (e StatusPaymentRequired) Error() string {
    return string(e)
}

func (e StatusPaymentRequired) StatusCode() int {
    return http.StatusPaymentRequired
}

func (e StatusPaymentRequired) New(message string) Error[StatusPaymentRequired] {
    return New(e).WithMessage(message)
}

func (e StatusForbidden) Error() string {
    return string(e)
}

func (e StatusForbidden) StatusCode() int {
    return http.StatusForbidden
}

func (e StatusForbidden) New(message string) Error[StatusForbidden] {
    return New(e).WithMessage(message)
}

func (e StatusNotFound) Error() string {
    return string(e)
}

func (e StatusNotFound) StatusCode() int {
    return http.StatusNotFound
}

func (e StatusNotFound) New(message string) Error[StatusNotFound] {
    return New(e).WithMessage(message)
}

func (e StatusMethodNotAllowed) Error() string {
    return string(e)
}

func (e StatusMethodNotAllowed) StatusCode() int {
    return http.StatusMethodNotAllowed
}

func (e StatusMethodNotAllowed) New(message string) Error[StatusMethodNotAllowed] {
    return New(e).WithMessage(message)
}

func (e StatusNotAcceptable) Error() string {
    return string(e)
}

func (e StatusNotAcceptable) StatusCode() int {
    return http.StatusNotAcceptable
}

func (e StatusNotAcceptable) New(message string) Error[StatusNotAcceptable] {
    return New(e).WithMessage(message)
}

func (e StatusProxyAuthRequired) Error() string {
    return string(e)
}

func (e StatusProxyAuthRequired) StatusCode() int {
    return http.StatusProxyAuthRequired
}

func (e StatusProxyAuthRequired) New(message string) Error[StatusProxyAuthRequired] {
    return New(e).WithMessage(message)
}

func (e StatusRequestTimeout) Error() string {
    return string(e)
}

func (e StatusRequestTimeout) StatusCode() int {
    return http.StatusRequestTimeout
}

func (e StatusRequestTimeout) New(message string) Error[StatusRequestTimeout] {
    return New(e).WithMessage(message)
}

func (e StatusConflict) Error() string {
    return string(e)
}

func (e StatusConflict) StatusCode() int {
    return http.StatusConflict
}

func (e StatusConflict) New(message string) Error[StatusConflict] {
    return New(e).WithMessage(message)
}

func (e StatusGone) Error() string {
    return string(e)
}

func (e StatusGone) StatusCode() int {
    return http.StatusGone
}

func (e StatusGone) New(message string) Error[StatusGone] {
    return New(e).WithMessage(message)
}

func (e StatusLengthRequired) Error() string {
    return string(e)
}

func (e StatusLengthRequired) StatusCode() int {
    return http.StatusLengthRequired
}

func (e StatusLengthRequired) New(message string) Error[StatusLengthRequired] {
    return New(e).WithMessage(message)
}

func (e StatusPreconditionFailed) Error() string {
    return string(e)
}

func (e StatusPreconditionFailed) StatusCode() int {
    return http.StatusPreconditionFailed
}

func (e StatusPreconditionFailed) New(message string) Error[StatusPreconditionFailed] {
    return New(e).WithMessage(message)
}

func (e StatusRequestEntityTooLarge) Error() string {
    return string(e)
}

func (e StatusRequestEntityTooLarge) StatusCode() int {
    return http.StatusRequestEntityTooLarge
}

func (e StatusRequestEntityTooLarge) New(message string) Error[StatusRequestEntityTooLarge] {
    return New(e).WithMessage(message)
}

func (e StatusRequestURITooLong) Error() string {
    return string(e)
}

func (e StatusRequestURITooLong) StatusCode() int {
    return http.StatusRequestURITooLong
}

func (e StatusRequestURITooLong) New(message string) Error[StatusRequestURITooLong] {
    return New(e).WithMessage(message)
}

func (e StatusUnsupportedMediaType) Error() string {
    return string(e)
}

func (e StatusUnsupportedMediaType) StatusCode() int {
    return http.StatusUnsupportedMediaType
}

func (e StatusUnsupportedMediaType) New(message string) Error[StatusUnsupportedMediaType] {
    return New(e).WithMessage(message)
}

func (e StatusRequestedRangeNotSatisfiable) Error() string {
    return string(e)
}

func (e StatusRequestedRangeNotSatisfiable) StatusCode() int {
    return http.StatusRequestedRangeNotSatisfiable
}

func (e StatusRequestedRangeNotSatisfiable) New(message string) Error[StatusRequestedRangeNotSatisfiable] {
    return New(e).WithMessage(message)
}

func (e StatusExpectationFailed) Error() string {
    return string(e)
}

func (e StatusExpectationFailed) StatusCode() int {
    return http.StatusExpectationFailed
}

func (e StatusExpectationFailed) New(message string) Error[StatusExpectationFailed] {
    return New(e).WithMessage(message)
}

func (e StatusTeapot) Error() string {
    return string(e)
}

func (e StatusTeapot) StatusCode() int {
    return http.StatusTeapot
}

func (e StatusTeapot) New(message string) Error[StatusTeapot] {
    return New(e).WithMessage(message)
}

func (e StatusMisdirectedRequest) Error() string {
    return string(e)
}

func (e StatusMisdirectedRequest) StatusCode() int {
    return http.StatusMisdirectedRequest
}

func (e StatusMisdirectedRequest) New(message string) Error[StatusMisdirectedRequest] {
    return New(e).WithMessage(message)
}

func (e StatusUnprocessableEntity) Error() string {
    return string(e)
}

func (e StatusUnprocessableEntity) StatusCode() int {
    return http.StatusUnprocessableEntity
}

func (e StatusUnprocessableEntity) New(message string) Error[StatusUnprocessableEntity] {
    return New(e).WithMessage(message)
}

func (e StatusLocked) Error() string {
    return string(e)
}

func (e StatusLocked) StatusCode() int {
    return http.StatusLocked
}

func (e StatusLocked) New(message string) Error[StatusLocked] {
    return New(e).WithMessage(message)
}

func (e StatusFailedDependency) Error() string {
    return string(e)
}

func (e StatusFailedDependency) StatusCode() int {
    return http.StatusFailedDependency
}

func (e StatusFailedDependency) New(message string) Error[StatusFailedDependency] {
    return New(e).WithMessage(message)
}

func (e StatusTooEarly) Error() string {
    return string(e)
}

func (e StatusTooEarly) StatusCode() int {
    return http.StatusTooEarly
}

func (e StatusTooEarly) New(message string) Error[StatusTooEarly] {
    return New(e).WithMessage(message)
}

func (e StatusUpgradeRequired) Error() string {
    return string(e)
}

func (e StatusUpgradeRequired) StatusCode() int {
    return http.StatusUpgradeRequired
}

func (e StatusUpgradeRequired) New(message string) Error[StatusUpgradeRequired] {
    return New(e).WithMessage(message)
}

func (e StatusPreconditionRequired) Error() string {
    return string(e)
}

func (e StatusPreconditionRequired) StatusCode() int {
    return http.StatusPreconditionRequired
}

func (e StatusPreconditionRequired) New(message string) Error[StatusPreconditionRequired] {
    return New(e).WithMessage(message)
}

func (e StatusTooManyRequests) Error() string {
    return string(e)
}

func (e StatusTooManyRequests) StatusCode() int {
    return http.StatusTooManyRequests
}

func (e StatusTooManyRequests) New(message string) Error[StatusTooManyRequests] {
    return New(e).WithMessage(message)
}

func (e StatusRequestHeaderFieldsTooLarge) Error() string {
    return string(e)
}

func (e StatusRequestHeaderFieldsTooLarge) StatusCode() int {
    return http.StatusRequestHeaderFieldsTooLarge
}

func (e StatusRequestHeaderFieldsTooLarge) New(message string) Error[StatusRequestHeaderFieldsTooLarge] {
    return New(e).WithMessage(message)
}

func (e StatusUnavailableForLegalReasons) Error() string {
    return string(e)
}

func (e StatusUnavailableForLegalReasons) StatusCode() int {
    return http.StatusUnavailableForLegalReasons
}

func (e StatusUnavailableForLegalReasons) New(message string) Error[StatusUnavailableForLegalReasons] {
    return New(e).WithMessage(message)
}

func (e StatusInternalServerError) Error() string {
    return string(e)
}

func (e StatusInternalServerError) StatusCode() int {
    return http.StatusInternalServerError
}

func (e StatusInternalServerError) New(message string) Error[StatusInternalServerError] {
    return New(e).WithMessage(message)
}

func (e StatusNotImplemented) Error() string {
    return string(e)
}

func (e StatusNotImplemented) StatusCode() int {
    return http.StatusNotImplemented
}

func (e StatusNotImplemented) New(message string) Error[StatusNotImplemented] {
    return New(e).WithMessage(message)
}

func (e StatusBadGateway) Error() string {
    return string(e)
}

func (e StatusBadGateway) StatusCode() int {
    return http.StatusBadGateway
}

func (e StatusBadGateway) New(message string) Error[StatusBadGateway] {
    return New(e).WithMessage(message)
}

func (e StatusServiceUnavailable) Error() string {
    return string(e)
}

func (e StatusServiceUnavailable) StatusCode() int {
    return http.StatusServiceUnavailable
}

func (e StatusServiceUnavailable) New(message string) Error[StatusServiceUnavailable] {
    return New(e).WithMessage(message)
}

func (e StatusGatewayTimeout) Error() string {
    return string(e)
}

func (e StatusGatewayTimeout) StatusCode() int {
    return http.StatusGatewayTimeout
}

func (e StatusGatewayTimeout) New(message string) Error[StatusGatewayTimeout] {
    return New(e).WithMessage(message)
}

func (e StatusHTTPVersionNotSupported) Error() string {
    return string(e)
}

func (e StatusHTTPVersionNotSupported) StatusCode() int {
    return http.StatusHTTPVersionNotSupported
}

func (e StatusHTTPVersionNotSupported) New(message string) Error[StatusHTTPVersionNotSupported] {
    return New(e).WithMessage(message)
}

func (e StatusVariantAlsoNegotiates) Error() string {
    return string(e)
}

func (e StatusVariantAlsoNegotiates) StatusCode() int {
    return http.StatusVariantAlsoNegotiates
}

func (e StatusVariantAlsoNegotiates) New(message string) Error[StatusVariantAlsoNegotiates] {
    return New(e).WithMessage(message)
}

func (e StatusInsufficientStorage) Error() string {
    return string(e)
}

func (e StatusInsufficientStorage) StatusCode() int {
    return http.StatusInsufficientStorage
}

func (e StatusInsufficientStorage) New(message string) Error[StatusInsufficientStorage] {
    return New(e).WithMessage(message)
}

func (e StatusLoopDetected) Error() string {
    return string(e)
}

func (e StatusLoopDetected) StatusCode() int {
    return http.StatusLoopDetected
}

func (e StatusLoopDetected) New(message string) Error[StatusLoopDetected] {
    return New(e).WithMessage(message)
}

func (e StatusNotExtended) Error() string {
    return string(e)
}

func (e StatusNotExtended) StatusCode() int {
    return http.StatusNotExtended
}

func (e StatusNotExtended) New(message string) Error[StatusNotExtended] {
    return New(e).WithMessage(message)
}

func (e StatusNetworkAuthenticationRequired) Error() string {
    return string(e)
}

func (e StatusNetworkAuthenticationRequired) StatusCode() int {
    return http.StatusNetworkAuthenticationRequired
}

func (e StatusNetworkAuthenticationRequired) New(message string) Error[StatusNetworkAuthenticationRequired] {
    return New(e).WithMessage(message)
}

