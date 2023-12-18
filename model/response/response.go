package response

type ResponseContainer struct {
	StatusCode      int           `json:"statusCode"`
	ErrorCode       *string       `json:"errorCode"`
	ResponseCode    *string       `json:"responseCode"`
	ResponseMessage *string       `json:"responseMessage"`
	Errors          []string      `json:"errors"`
	Data            interface{}   `json:"data"`
	Info            *ResponseInfo `json:"info,omitempty"`
}

type ResponseInfo struct {
	Limit    int `json:"limit"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

type ResponseJSONContainer struct {
	Response DataHubBodyRespond `json:"RESPONSE"`
}

type DataHubBodyRespond struct {
	StatusCode      int    `json:"STATUS_CODE"`
	ErrorCode       string `json:"ERROR_CODE"`
	ResponseCode    string `json:"RESPONSE_CODE"`
	ResponseMessage string `json:"RESPONSE_MESSAGE"`
}
