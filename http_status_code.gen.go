package problem

import "net/http"
func (e StatusMultipleChoices) Error() string {
    return string(e)
}

func (e StatusMultipleChoices) StatusCode() StatusCode {
    return StatusCode(http.StatusMultipleChoices)
}

func (e StatusMultipleChoices) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusMovedPermanently) Error() string {
    return string(e)
}

func (e StatusMovedPermanently) StatusCode() StatusCode {
    return StatusCode(http.StatusMovedPermanently)
}

func (e StatusMovedPermanently) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusFound) Error() string {
    return string(e)
}

func (e StatusFound) StatusCode() StatusCode {
    return StatusCode(http.StatusFound)
}

func (e StatusFound) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusSeeOther) Error() string {
    return string(e)
}

func (e StatusSeeOther) StatusCode() StatusCode {
    return StatusCode(http.StatusSeeOther)
}

func (e StatusSeeOther) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusNotModified) Error() string {
    return string(e)
}

func (e StatusNotModified) StatusCode() StatusCode {
    return StatusCode(http.StatusNotModified)
}

func (e StatusNotModified) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusUseProxy) Error() string {
    return string(e)
}

func (e StatusUseProxy) StatusCode() StatusCode {
    return StatusCode(http.StatusUseProxy)
}

func (e StatusUseProxy) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusTemporaryRedirect) Error() string {
    return string(e)
}

func (e StatusTemporaryRedirect) StatusCode() StatusCode {
    return StatusCode(http.StatusTemporaryRedirect)
}

func (e StatusTemporaryRedirect) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusPermanentRedirect) Error() string {
    return string(e)
}

func (e StatusPermanentRedirect) StatusCode() StatusCode {
    return StatusCode(http.StatusPermanentRedirect)
}

func (e StatusPermanentRedirect) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusBadRequest) Error() string {
    return string(e)
}

func (e StatusBadRequest) StatusCode() StatusCode {
    return StatusCode(http.StatusBadRequest)
}

func (e StatusBadRequest) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusUnauthorized) Error() string {
    return string(e)
}

func (e StatusUnauthorized) StatusCode() StatusCode {
    return StatusCode(http.StatusUnauthorized)
}

func (e StatusUnauthorized) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusPaymentRequired) Error() string {
    return string(e)
}

func (e StatusPaymentRequired) StatusCode() StatusCode {
    return StatusCode(http.StatusPaymentRequired)
}

func (e StatusPaymentRequired) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusForbidden) Error() string {
    return string(e)
}

func (e StatusForbidden) StatusCode() StatusCode {
    return StatusCode(http.StatusForbidden)
}

func (e StatusForbidden) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusNotFound) Error() string {
    return string(e)
}

func (e StatusNotFound) StatusCode() StatusCode {
    return StatusCode(http.StatusNotFound)
}

func (e StatusNotFound) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusMethodNotAllowed) Error() string {
    return string(e)
}

func (e StatusMethodNotAllowed) StatusCode() StatusCode {
    return StatusCode(http.StatusMethodNotAllowed)
}

func (e StatusMethodNotAllowed) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusNotAcceptable) Error() string {
    return string(e)
}

func (e StatusNotAcceptable) StatusCode() StatusCode {
    return StatusCode(http.StatusNotAcceptable)
}

func (e StatusNotAcceptable) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusProxyAuthRequired) Error() string {
    return string(e)
}

func (e StatusProxyAuthRequired) StatusCode() StatusCode {
    return StatusCode(http.StatusProxyAuthRequired)
}

func (e StatusProxyAuthRequired) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusRequestTimeout) Error() string {
    return string(e)
}

func (e StatusRequestTimeout) StatusCode() StatusCode {
    return StatusCode(http.StatusRequestTimeout)
}

func (e StatusRequestTimeout) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusConflict) Error() string {
    return string(e)
}

func (e StatusConflict) StatusCode() StatusCode {
    return StatusCode(http.StatusConflict)
}

func (e StatusConflict) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusGone) Error() string {
    return string(e)
}

func (e StatusGone) StatusCode() StatusCode {
    return StatusCode(http.StatusGone)
}

func (e StatusGone) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusLengthRequired) Error() string {
    return string(e)
}

func (e StatusLengthRequired) StatusCode() StatusCode {
    return StatusCode(http.StatusLengthRequired)
}

func (e StatusLengthRequired) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusPreconditionFailed) Error() string {
    return string(e)
}

func (e StatusPreconditionFailed) StatusCode() StatusCode {
    return StatusCode(http.StatusPreconditionFailed)
}

func (e StatusPreconditionFailed) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusRequestEntityTooLarge) Error() string {
    return string(e)
}

func (e StatusRequestEntityTooLarge) StatusCode() StatusCode {
    return StatusCode(http.StatusRequestEntityTooLarge)
}

func (e StatusRequestEntityTooLarge) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusRequestURITooLong) Error() string {
    return string(e)
}

func (e StatusRequestURITooLong) StatusCode() StatusCode {
    return StatusCode(http.StatusRequestURITooLong)
}

func (e StatusRequestURITooLong) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusUnsupportedMediaType) Error() string {
    return string(e)
}

func (e StatusUnsupportedMediaType) StatusCode() StatusCode {
    return StatusCode(http.StatusUnsupportedMediaType)
}

func (e StatusUnsupportedMediaType) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusRequestedRangeNotSatisfiable) Error() string {
    return string(e)
}

func (e StatusRequestedRangeNotSatisfiable) StatusCode() StatusCode {
    return StatusCode(http.StatusRequestedRangeNotSatisfiable)
}

func (e StatusRequestedRangeNotSatisfiable) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusExpectationFailed) Error() string {
    return string(e)
}

func (e StatusExpectationFailed) StatusCode() StatusCode {
    return StatusCode(http.StatusExpectationFailed)
}

func (e StatusExpectationFailed) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusTeapot) Error() string {
    return string(e)
}

func (e StatusTeapot) StatusCode() StatusCode {
    return StatusCode(http.StatusTeapot)
}

func (e StatusTeapot) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusMisdirectedRequest) Error() string {
    return string(e)
}

func (e StatusMisdirectedRequest) StatusCode() StatusCode {
    return StatusCode(http.StatusMisdirectedRequest)
}

func (e StatusMisdirectedRequest) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusUnprocessableEntity) Error() string {
    return string(e)
}

func (e StatusUnprocessableEntity) StatusCode() StatusCode {
    return StatusCode(http.StatusUnprocessableEntity)
}

func (e StatusUnprocessableEntity) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusLocked) Error() string {
    return string(e)
}

func (e StatusLocked) StatusCode() StatusCode {
    return StatusCode(http.StatusLocked)
}

func (e StatusLocked) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusFailedDependency) Error() string {
    return string(e)
}

func (e StatusFailedDependency) StatusCode() StatusCode {
    return StatusCode(http.StatusFailedDependency)
}

func (e StatusFailedDependency) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusTooEarly) Error() string {
    return string(e)
}

func (e StatusTooEarly) StatusCode() StatusCode {
    return StatusCode(http.StatusTooEarly)
}

func (e StatusTooEarly) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusUpgradeRequired) Error() string {
    return string(e)
}

func (e StatusUpgradeRequired) StatusCode() StatusCode {
    return StatusCode(http.StatusUpgradeRequired)
}

func (e StatusUpgradeRequired) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusPreconditionRequired) Error() string {
    return string(e)
}

func (e StatusPreconditionRequired) StatusCode() StatusCode {
    return StatusCode(http.StatusPreconditionRequired)
}

func (e StatusPreconditionRequired) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusTooManyRequests) Error() string {
    return string(e)
}

func (e StatusTooManyRequests) StatusCode() StatusCode {
    return StatusCode(http.StatusTooManyRequests)
}

func (e StatusTooManyRequests) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusRequestHeaderFieldsTooLarge) Error() string {
    return string(e)
}

func (e StatusRequestHeaderFieldsTooLarge) StatusCode() StatusCode {
    return StatusCode(http.StatusRequestHeaderFieldsTooLarge)
}

func (e StatusRequestHeaderFieldsTooLarge) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusUnavailableForLegalReasons) Error() string {
    return string(e)
}

func (e StatusUnavailableForLegalReasons) StatusCode() StatusCode {
    return StatusCode(http.StatusUnavailableForLegalReasons)
}

func (e StatusUnavailableForLegalReasons) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusInternalServerError) Error() string {
    return string(e)
}

func (e StatusInternalServerError) StatusCode() StatusCode {
    return StatusCode(http.StatusInternalServerError)
}

func (e StatusInternalServerError) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusNotImplemented) Error() string {
    return string(e)
}

func (e StatusNotImplemented) StatusCode() StatusCode {
    return StatusCode(http.StatusNotImplemented)
}

func (e StatusNotImplemented) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusBadGateway) Error() string {
    return string(e)
}

func (e StatusBadGateway) StatusCode() StatusCode {
    return StatusCode(http.StatusBadGateway)
}

func (e StatusBadGateway) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusServiceUnavailable) Error() string {
    return string(e)
}

func (e StatusServiceUnavailable) StatusCode() StatusCode {
    return StatusCode(http.StatusServiceUnavailable)
}

func (e StatusServiceUnavailable) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusGatewayTimeout) Error() string {
    return string(e)
}

func (e StatusGatewayTimeout) StatusCode() StatusCode {
    return StatusCode(http.StatusGatewayTimeout)
}

func (e StatusGatewayTimeout) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusHTTPVersionNotSupported) Error() string {
    return string(e)
}

func (e StatusHTTPVersionNotSupported) StatusCode() StatusCode {
    return StatusCode(http.StatusHTTPVersionNotSupported)
}

func (e StatusHTTPVersionNotSupported) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusVariantAlsoNegotiates) Error() string {
    return string(e)
}

func (e StatusVariantAlsoNegotiates) StatusCode() StatusCode {
    return StatusCode(http.StatusVariantAlsoNegotiates)
}

func (e StatusVariantAlsoNegotiates) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusInsufficientStorage) Error() string {
    return string(e)
}

func (e StatusInsufficientStorage) StatusCode() StatusCode {
    return StatusCode(http.StatusInsufficientStorage)
}

func (e StatusInsufficientStorage) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusLoopDetected) Error() string {
    return string(e)
}

func (e StatusLoopDetected) StatusCode() StatusCode {
    return StatusCode(http.StatusLoopDetected)
}

func (e StatusLoopDetected) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusNotExtended) Error() string {
    return string(e)
}

func (e StatusNotExtended) StatusCode() StatusCode {
    return StatusCode(http.StatusNotExtended)
}

func (e StatusNotExtended) New(message string) Error {
    return New(e).WithMessage(message)
}

func (e StatusNetworkAuthenticationRequired) Error() string {
    return string(e)
}

func (e StatusNetworkAuthenticationRequired) StatusCode() StatusCode {
    return StatusCode(http.StatusNetworkAuthenticationRequired)
}

func (e StatusNetworkAuthenticationRequired) New(message string) Error {
    return New(e).WithMessage(message)
}

func convertToTitle(statusCode int, title string) error {
    switch statusCode {
    case http.StatusMultipleChoices: return StatusMultipleChoices(title)
    case http.StatusMovedPermanently: return StatusMovedPermanently(title)
    case http.StatusFound: return StatusFound(title)
    case http.StatusSeeOther: return StatusSeeOther(title)
    case http.StatusNotModified: return StatusNotModified(title)
    case http.StatusUseProxy: return StatusUseProxy(title)
    case http.StatusTemporaryRedirect: return StatusTemporaryRedirect(title)
    case http.StatusPermanentRedirect: return StatusPermanentRedirect(title)
    case http.StatusBadRequest: return StatusBadRequest(title)
    case http.StatusUnauthorized: return StatusUnauthorized(title)
    case http.StatusPaymentRequired: return StatusPaymentRequired(title)
    case http.StatusForbidden: return StatusForbidden(title)
    case http.StatusNotFound: return StatusNotFound(title)
    case http.StatusMethodNotAllowed: return StatusMethodNotAllowed(title)
    case http.StatusNotAcceptable: return StatusNotAcceptable(title)
    case http.StatusProxyAuthRequired: return StatusProxyAuthRequired(title)
    case http.StatusRequestTimeout: return StatusRequestTimeout(title)
    case http.StatusConflict: return StatusConflict(title)
    case http.StatusGone: return StatusGone(title)
    case http.StatusLengthRequired: return StatusLengthRequired(title)
    case http.StatusPreconditionFailed: return StatusPreconditionFailed(title)
    case http.StatusRequestEntityTooLarge: return StatusRequestEntityTooLarge(title)
    case http.StatusRequestURITooLong: return StatusRequestURITooLong(title)
    case http.StatusUnsupportedMediaType: return StatusUnsupportedMediaType(title)
    case http.StatusRequestedRangeNotSatisfiable: return StatusRequestedRangeNotSatisfiable(title)
    case http.StatusExpectationFailed: return StatusExpectationFailed(title)
    case http.StatusTeapot: return StatusTeapot(title)
    case http.StatusMisdirectedRequest: return StatusMisdirectedRequest(title)
    case http.StatusUnprocessableEntity: return StatusUnprocessableEntity(title)
    case http.StatusLocked: return StatusLocked(title)
    case http.StatusFailedDependency: return StatusFailedDependency(title)
    case http.StatusTooEarly: return StatusTooEarly(title)
    case http.StatusUpgradeRequired: return StatusUpgradeRequired(title)
    case http.StatusPreconditionRequired: return StatusPreconditionRequired(title)
    case http.StatusTooManyRequests: return StatusTooManyRequests(title)
    case http.StatusRequestHeaderFieldsTooLarge: return StatusRequestHeaderFieldsTooLarge(title)
    case http.StatusUnavailableForLegalReasons: return StatusUnavailableForLegalReasons(title)
    case http.StatusInternalServerError: return StatusInternalServerError(title)
    case http.StatusNotImplemented: return StatusNotImplemented(title)
    case http.StatusBadGateway: return StatusBadGateway(title)
    case http.StatusServiceUnavailable: return StatusServiceUnavailable(title)
    case http.StatusGatewayTimeout: return StatusGatewayTimeout(title)
    case http.StatusHTTPVersionNotSupported: return StatusHTTPVersionNotSupported(title)
    case http.StatusVariantAlsoNegotiates: return StatusVariantAlsoNegotiates(title)
    case http.StatusInsufficientStorage: return StatusInsufficientStorage(title)
    case http.StatusLoopDetected: return StatusLoopDetected(title)
    case http.StatusNotExtended: return StatusNotExtended(title)
    case http.StatusNetworkAuthenticationRequired: return StatusNetworkAuthenticationRequired(title)
    default: return StatusInternalServerError(title)
    }
}

