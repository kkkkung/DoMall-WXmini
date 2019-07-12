package e

var MsgFlags = map[int]string{
	// Custom
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	// HTTP Status Codes
	// 100
	CONTINUE:            "Continue",
	SWITCHING_PROTOCOLS: "Switching Protocols",
	PROCESSING:          "Processing",
	EARLY_HINTS:         "Early Hints",

	// 200
	OK:                            "OK",
	CREATED:                       "Created",
	ACCEPTED:                      "Accepted",
	NON_AUTHORITATIVE_INFORMATION: "Non-Authoritative Information",
	NO_CONTENT:                    "No Content",
	RESET_CONTENT:                 "Reset Content",
	PARTIAL_CONTENT:               "Partial Content",
	MULTI_STATUS:                  "Multi-Status",
	ALREADY_REPORTED:              "Already Reported",
	IM_USED:                       "IM Used",

	// 300
	MULTIPLE_CHOICES:   "Multiple Choices",
	MOVED_PERMANENTLY:  "Moved Permanently",
	FOUND:              "Found",
	SEE_OTHER:          "See Other",
	NOT_MODIFIED:       "Not Modified",
	USE_PROXY:          "Use Proxy",
	TEMPORARY_REDIRECT: "Temporary Redirect",
	PERMANENT_REDIRECT: "Permanent Redirect",

	// 400
	BAD_REQUEST:                     "Bad Request",
	UNAUTHORIZED:                    "Unauthorized",
	PAYMENT_REQUIRED:                "Payment Required",
	FORBIDDEN:                       "Forbidden",
	NOT_FOUND:                       "Not Found",
	METHOD_NOT_ALLOWED:              "Method Not Allowed",
	NOT_ACCEPTABLE:                  "Not Acceptable",
	PROXY_AUTHENTICATION_REQUIRED:   "Proxy Authentication Required",
	REQUEST_TIMEOUT:                 "Request Timeout",
	CONFLICT:                        "Conflict",
	GONE:                            "Gone",
	LENGTH_REQUIRED:                 "Length Required",
	PRECONDITION_FAILED:             "Precondition Failed",
	PAYLOAD_TOO_LARGE:               "Payload Too Large",
	URI_TOO_LONG:                    "URI Too Long",
	UNSUPPORTED_MEDIA_TYPE:          "Unsupported Media Type",
	RANGE_NOT_SATISFIABLE:           "Range Not Satisfiable",
	EXPECTATION_FAILED:              "Expectation Failed",
	MISDIRECTED_REQUEST:             "Misdirected Request",
	UNPROCESSABLE_ENTITY:            "Unprocessable Entity",
	Locked:                          "Locked",
	FAILED_DEPENDENCY:               "Failed Dependency",
	TOO_EARLY:                       "Too Early",
	UPGRADE_REQUIRED:                "Upgrade Required",
	PRECONDITION_REQUIRED:           "Precondition Required",
	TOO_MANY_REQUESTS:               "Too Many Requests",
	REQUEST_HEADER_FIELDS_TOO_LARGE: "Request Header Fields Too Large",
	UNAVAILABLE_FOR_LEGAL_REASONS:   "Unavailable For Legal Reasons",

	// 500
	INTERNAL_SERVER_ERROR:           "Internal Server Error",
	NOT_IMPLEMENTED:                 "Not Implemented",
	BAD_GATEWAY:                     "Bad Gateway",
	SERVICE_UNAVAILABLE:             "Service Unavailable",
	GATEWAY_TIMEOUT:                 "Gateway Timeout",
	HTTP_VERSION_NOT_SUPPORTED:      "HTTP Version Not Supported",
	VARIANT_ALSO_NEGOTIATES:         "Variant Also Negotiates",
	INSUFFICIENT_STORAGE:            "Insufficient Storage",
	LOOP_DETECTED:                   "Loop Detected",
	NOT_EXTENDED:                    "Not Extended",
	NETWORK_AUTHENTICATION_REQUIRED: "Network Authentication Required",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[INTERNAL_SERVER_ERROR]
}
