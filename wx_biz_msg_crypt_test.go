package wxcrypt

import (
	"encoding/json"
	"testing"
)

func TestWxBizMsgCrypt(t *testing.T) {

	encodingAesKey := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFG"
	token := "pamtest"
	timeStamp := "1409304348"
	nonce := "xxxxxx"
	appID := "wxb11529c136998cb6"

	textMap := map[string]interface{}{
		"ToUserName":   "oia2Tj我是中文jewbmiOUlr6X-1crbLOvLw",
		"FromUserName": "gh_7f083739789a",
		"CreateTime":   "1407743423",
		"MsgType":      "video",
		"Video": map[string]string{
			"MediaID":     "eYJ1MbwPRJtOvIEabaxHs7TX2D-HV71s79GUxqdUkjm6Gs2Ed1KF3ulAOA9H1xG0",
			"Title":       "testCallbackReplyVideo",
			"Description": "testCallbackReplyVideo",
		},
	}
	jsonBytes, err := json.Marshal(textMap)
	if err != nil {
		panic(err)
	}

	c := WxCrypt{
		AesKey: encodingAesKey,
		AppID:  appID,
		Token:  token,
	}

	encryptMsg, err := c.EncryptMsg(jsonBytes, timeStamp, nonce)
	if err != nil {
		panic(err)
	}
	t.Log(encryptMsg)

	sha1Sign := GetSHA1(token, timeStamp, nonce, encryptMsg)

	jBytes, err := c.DecryptMsg(sha1Sign, timeStamp, nonce, encryptMsg)
	if err != nil {
		panic(err)
	}

	t.Log(string(jBytes))
}
