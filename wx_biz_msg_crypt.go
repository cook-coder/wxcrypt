package wxcrypt

import "errors"

// WxCrypt 微信加解密
type WxCrypt struct {
	Token  string
	AesKey string
	AppID  string
}

// EncryptMsg 加密消息
func (c *WxCrypt) EncryptMsg(msgBytes []byte, timestamp, nonce string) (string, error) {
	p := prpcrypt{
		AesKey: c.AesKey,
		AppID:  c.AppID,
	}
	encryptMsg, err := p.encrypt(msgBytes)
	if err != nil {
		return "", err
	}
	return encryptMsg, nil
}

// DecryptMsg 解密消息
func (c *WxCrypt) DecryptMsg(msgSignature, timestamp, nonce, encryptMsg string) ([]byte, error) {
	p := prpcrypt{
		AesKey: c.AesKey,
		AppID:  c.AppID,
	}

	// 验证安全签名
	sha1Sign := GetSHA1(c.Token, timestamp, nonce, encryptMsg)
	if sha1Sign != msgSignature {
		return nil, errors.New("validate signature error")
	}

	plainTextBytes, err := p.decrypt(encryptMsg)
	if err != nil {
		return nil, err
	}

	return plainTextBytes, nil

}
