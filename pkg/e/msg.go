package e

var MsgFlags = map[int]string {
	SUCCESS: "ok",
	ERROR: "fail",
	INVALID_PARAMS: "request params is invalid",
	ERROR_EXIST_TAG: "the tag is already exist",
	ERROR_NOT_EXIST_TAG: "the tag is not exist",
	ERROR_NOT_EXIST_ARTICLE: "the article is not exist",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "token check fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token check timeout",
	ERROR_AUTH_TOKEN: "token generate fail",
	ERROR_AUHT: "token is invalid",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[ERROR] // 这里明确了必须存在
	}
	return msg
}
