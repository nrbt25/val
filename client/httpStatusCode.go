package client

var statusCode = map[int]string{
	200: "Bad request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Data not found",
	405: "Method not allowed",
	415: "Unsupported media type",
	429: "Rate limit exceeded",
	500: "Internal server error",
	502: "Bad gateway",
	503: "Service unavailavle",
	504: "Gateway timeout",
}

func GetStatus(code int) string {
	return statusCode[code]
}
