package wxcrypt

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestEncrypt(t *testing.T) {
	xml := "<xml><ToUserName><![CDATA[oia2Tj我是中文jewbmiOUlr6X-1crbLOvLw]]></ToUserName><FromUserName><![CDATA[gh_7f083739789a]]></FromUserName><CreateTime>1407743423</CreateTime><MsgType><![CDATA[video]]></MsgType><Video><MediaId><![CDATA[eYJ1MbwPRJtOvIEabaxHs7TX2D-HV71s79GUxqdUkjm6Gs2Ed1KF3ulAOA9H1xG0]]></MediaId><Title><![CDATA[testCallBackReplyVideo]]></Title><Description><![CDATA[testCallBackReplyVideo]]></Description></Video></xml>"
	appID := "wxb11529c136998cb6"
	p := prpcrypt{
		AppID:  appID,
		AesKey: "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFG",
	}
	encrypt, err := p.encrypt([]byte(xml))
	if err != nil {
		panic(err)
	}
	t.Log(encrypt)
}

func TestDecrypt(t *testing.T) {
	encrypt := "ECb94sC68hxnidW1ycQhs0FeImdr6uObTDOLJQFXMtewrM/0umNcgOjLQuXRU40Up76V8iAHR8muCqxC96UBy0fpzKnZi3yI5ECA0SBnUOpEp05uz8VrRcLwt4+8Moq4PDNmRWydNO8a7Sa03pK1MmKeVdxti4m4f3KDK8YPqGB1SXRGesgmepYgroJeWWDNkVdwHYPbLjA1Lqh8t2XVpG/k3iaP4mFFoIwt7LIeqi40tXb8OL2yoCkQ6nKOy/O5SwmdanY8FyTWDly6DWoEJQwUBwk4AuLrUZq1q3VwucQ1TC880B+0pCqA6aPaop1Mv/jy07sjMiQYq4xcywBq17TXD6JPZaJLDlyKOIQrJ4bBVJknjDkuh2h6h72uVaDpdXtmML0OqA+AvUcouszmk0xEc0OaubPllXgzxacQ7Ag2lVg0byuVa6LEx0pQMlD9xWAEbPHaBWzV9JN2kiMZPmoRm0BqwtXGcTuLmLTiECnHGl8+Rsox5cBTOJKP+46IMRFvtgi+pUAqV4q0NEsXLQBIJeGS6NnBeRHBTD6SIjBCXQ9qm0pyJsSWsKmb95fdCGdQwk6h7Jp8YcyDQGF2OWl3CcH+TplE6LlMzQx6wvkQD2Tlzlac1y/F9ingBDCC"
	appID := "wxb11529c136998cb6"
	p := prpcrypt{
		AppID:  appID,
		AesKey: "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFG",
	}
	text, err := p.decrypt(encrypt)
	if err != nil {
		panic(err)
	}
	t.Log(text)

}

func TestGetRandomStr(t *testing.T) {
	rs := getRandomStrBytes()
	t.Log(rs)
}

func BenchmarkGetRandomStr(b *testing.B) {
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		rs := getRandomStrBytes()
		b.Log(rs)
	}
	b.StopTimer()
}

func TestBase64DecodeAesKey(t *testing.T) {
	aesKey := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFG"
	aesBytes, err := base64DecodeAesKey(aesKey)
	if err != nil {
		panic(err)
	}
	// 105 183 29 121 248 33 138 57 37 154 122 41 170 187 45 186 252 49 203 61 53 219 126 57 235 191 61 0 16 131 16 81
	t.Log(aesBytes)
	t.Log(string(aesBytes))

	h, err := hex.DecodeString(string(aesBytes))
	if err != nil {
		panic(err)
	}
	t.Log(h)
}

func TestBase64Decode(t *testing.T) {

	aesKeyBytes, _ := base64.StdEncoding.DecodeString("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFG=")
	t.Log(len(aesKeyBytes))
	t.Log(string(aesKeyBytes))

}
