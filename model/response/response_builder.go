package response

import (
	"net/http"
	"strings"

	"sbs-be/model/constant"
)

func BuildSuccessResponse(data interface{}) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusOK,
		ErrorCode:       &constant.ERROR_CODE_SUCCESS,
		ResponseCode:    &constant.RESPONSE_CODE_SUCCESS,
		ResponseMessage: &constant.RESPONSE_MESSAGE_SUCCESS,
		Errors:          nil,
		Data:            data,
		Info:            nil,
	}
}

func BuildSuccessResponseWithInfo(data interface{}, info *ResponseInfo) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusOK,
		ErrorCode:       &constant.ERROR_CODE_SUCCESS,
		ResponseCode:    &constant.RESPONSE_CODE_SUCCESS,
		ResponseMessage: &constant.RESPONSE_MESSAGE_SUCCESS,
		Errors:          nil,
		Data:            data,
		Info:            info,
	}
}

func BuildDataNotFoundResponse() *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusNotFound,
		ErrorCode:       &constant.ERROR_CODE_DATA_NOT_FOUND,
		ResponseCode:    &constant.RESPONSE_CODE_NOT_FOUND,
		ResponseMessage: &constant.RESPONSE_MESSAGE_DATA_NOT_FOUND,
		Errors:          nil,
		Data:            nil,
		Info:            nil,
	}
}

func BuildDataNotFoundResponseWithMessage(msg string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusNotFound,
		ErrorCode:       &constant.ERROR_CODE_DATA_NOT_FOUND,
		ResponseCode:    &constant.RESPONSE_CODE_NOT_FOUND,
		ResponseMessage: &constant.RESPONSE_MESSAGE_DATA_NOT_FOUND,
		Errors:          strings.Split(msg, "\n"),
		Data:            nil,
		Info:            nil,
	}
}

func BuildBadRequestResponse(errCode, respCode, errMessage, throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusBadRequest,
		ErrorCode:       &errCode,
		ResponseCode:    &respCode,
		ResponseMessage: &errMessage,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
		Info:            nil,
	}
}

func BuildInternalErrorResponse(errCode, respCode, errMessage, throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusInternalServerError,
		ErrorCode:       &errCode,
		ResponseCode:    &respCode,
		ResponseMessage: &errMessage,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
		Info:            nil,
	}
}

func BuildRouteNotFoundResponse() *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusNotFound,
		ErrorCode:       &constant.ERROR_CODE_DATA_NOT_FOUND,
		ResponseCode:    &constant.RESPONSE_CODE_NOT_FOUND,
		ResponseMessage: &constant.RESPONSE_MESSAGE_ROUTE_NOT_FOUND,
		Errors:          nil,
		Data:            nil,
		Info:            nil,
	}
}

func BuildEmptyBodyReqResponse(errMessage, throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusBadRequest,
		ErrorCode:       &constant.ERROR_CODE_BODY_REQUEST_EMPTY,
		ResponseCode:    &constant.RESPONSE_CODE_BAD_REQUEST,
		ResponseMessage: &errMessage,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
	}
}

func BuildInvalidTypeResponse(errMessage, throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusBadRequest,
		ErrorCode:       &constant.ERROR_CODE_INVALID_DATA_TYPE,
		ResponseCode:    &constant.RESPONSE_CODE_BAD_REQUEST,
		ResponseMessage: &errMessage,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
	}
}

func BuildUnauthorizedResponse(errMessage, throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusUnauthorized,
		ErrorCode:       &constant.ERROR_CODE_INVALID_AUTH_TOKEN,
		ResponseCode:    &constant.RESPONSE_CODE_AUTH_ERROR,
		ResponseMessage: &errMessage,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
	}
}

func BuildTimeoutResponse(throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusRequestTimeout,
		ErrorCode:       &constant.ERROR_CODE_TIMEOUT,
		ResponseCode:    &constant.RESPONSE_CODE_INTERNAL_ERROR,
		ResponseMessage: &constant.RESPONSE_MESSAGE_TIMEOUT,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
	}
}

func BuildForbiddenAccessResponse(throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusForbidden,
		ErrorCode:       &constant.ERROR_CODE_FORBIDDEN_ACCESS,
		ResponseCode:    &constant.RESPONSE_CODE_AUTH_ERROR,
		ResponseMessage: &constant.RESPONSE_MESSAGE_FORBIDDEN_ACCESS,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
	}
}
