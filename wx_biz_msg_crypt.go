package wxcrypt

import "errors"

// WxCrypt 微信加解密
type WxCrypt struct {
	Token  string
	AesKey string
	AppID  string
}

// EncryptMsg 加密消息
func (c *WxCrypt) EncryptMsg(msg, timestamp, nonce string) (map[string]string, error) {
	p := prpcrypt{
		AesKey: c.AesKey,
		AppID:  c.AppID,
	}
	encryptMsg, err := p.encrypt(msg)
	if err != nil {
		return nil, err
	}
	sha1Sign := GetSHA1(c.Token, timestamp, nonce, encryptMsg)
	return map[string]string{
		"Encrypt":      encryptMsg,
		"MsgSignature": sha1Sign,
		"TimeStamp":    timestamp,
		"Nonce":        nonce,
	}, nil
}

// DecryptMsg 解密消息
func (c *WxCrypt) DecryptMsg(msgSignature, timestamp, nonce, encryptMsg string) (string, error) {
	p := prpcrypt{
		AesKey: c.AesKey,
		AppID:  c.AppID,
	}

	// 验证安全签名
	sha1Sign := GetSHA1(c.Token, timestamp, nonce, encryptMsg)
	if sha1Sign != msgSignature {
		return "", errors.New("validate signature error")
	}

	plainText, err := p.decrypt(encryptMsg)
	if err != nil {
		return "", err
	}

	return plainText, nil

}
