package response

import "go-playground/m/v1/usecase/rule"

// ErrorCode ...
type ErrorCode int

// Text ...
func (a ErrorCode) Text() string {
	return errorTextMap[a]
}

var errorTextMap = map[ErrorCode]string{
	rule.NotFound:            "レコードが見つかりませんでした。",
	rule.InternalServerError: "予期せぬエラーが発生しました。",
	rule.BadRequest:          "入力内容に誤りがあります。",
	rule.ShortBalance:        "残高が不足しています。",
}
