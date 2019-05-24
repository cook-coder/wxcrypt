package wxcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"
	"time"
)

type prpcrypt struct {
	AesKey string
	AppID  string
}

// 加密
func (p *prpcrypt) encrypt(text string) (string, error) {
	bytesColl := []byte{}
	randBytes := getRandomStrBytes()
	textBytes := []byte(text)
	networkBytesOrder := getNetworkBytesOrder(len(textBytes))
	// randomStr + networkBytesOrder + text + appid
	bytesColl = append(randBytes, networkBytesOrder...)
	bytesColl = append(bytesColl, textBytes...)
	bytesColl = append(bytesColl, []byte(p.AppID)...)

	// ... + pad: 使用自定义的填充方式对明文进行补位填充
	bytesColl = pkcg7Encode(bytesColl)

	aesKeyBytes, err := base64DecodeAesKey(p.AesKey)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(aesKeyBytes)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(bytesColl))
	iv := aesKeyBytes[0:aes.BlockSize]

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, bytesColl)

	return base64.StdEncoding.EncodeToString(ciphertext), nil

}

// 解密
func (p *prpcrypt) decrypt(encrypt string) ([]byte, error) {
	base64DecodeEncrypt, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return nil, err
	}

	aesKeyBytes, err := base64DecodeAesKey(p.AesKey)
	if err != nil {
		return nil, err
	}

	// aes解密
	block, err := aes.NewCipher(aesKeyBytes)
	if err != nil {
		return nil, err
	}

	iv := aesKeyBytes[0:aes.BlockSize]

	mode := cipher.NewCBCDecrypter(block, iv)
	encryptLen := len(base64DecodeEncrypt)
	decryptedMsg := make([]byte, encryptLen, encryptLen)
	mode.CryptBlocks(decryptedMsg, base64DecodeEncrypt)

	// pkcg7解密
	pkcg7DecodedDecryptedMsg := pkcg7Decode(decryptedMsg)

	// 获取业务明文数据的长度; 因为网络字节序是4个字节,所以+4
	// aes.BlockSize是随机字符串的长度
	contentStartAt := aes.BlockSize + 4
	contentLength := recoverNetworkBytesOrder(pkcg7DecodedDecryptedMsg[aes.BlockSize:contentStartAt])
	contentEndAt := contentStartAt + contentLength
	content := pkcg7DecodedDecryptedMsg[contentStartAt:contentEndAt]
	// 截取APPID
	fromAppID := pkcg7DecodedDecryptedMsg[contentEndAt:]
	if string(fromAppID) != p.AppID {
		return nil, errors.New("app id does not match")
	}

	return content, nil
}

// 获取16位随机字符串
func getRandomStrBytes() []byte {
	strPool := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz")
	max := len(strPool) - 1
	rand.Seed(time.Now().UnixNano())
	var randBytes = make([]byte, 16, 16)
	for i := 0; i < 16; i++ {
		randIndex := rand.Intn(max)
		randBytes[i] = strPool[randIndex]
	}
	return randBytes
}

// 生成4个字节的网络字节序
func getNetworkBytesOrder(sourceNumbser int) []byte {
	orderBytes := make([]byte, 4, 4)
	orderBytes[3] = byte(sourceNumbser & 0xFF)
	orderBytes[2] = byte(sourceNumbser >> 8 & 0xFF)
	orderBytes[1] = byte(sourceNumbser >> 16 & 0xFF)
	orderBytes[0] = byte(sourceNumbser >> 24 & 0xFF)
	return orderBytes
}

// 还原4个字节的网络字节序
func recoverNetworkBytesOrder(orderBytes []byte) int {
	var sourceNumber int
	for i := 0; i < 4; i++ {
		sourceNumber <<= 8
		sourceNumber = sourceNumber | int(orderBytes[i]&0xFF)
	}
	return sourceNumber
}

// 对配置的AES进行Base64解码
func base64DecodeAesKey(aesKey string) ([]byte, error) {
	var buildler strings.Builder
	buildler.WriteString(aesKey)
	buildler.WriteString("=")
	aesKeyBytes, err := base64.StdEncoding.DecodeString(buildler.String())
	if err != nil {
		return nil, err
	}
	return aesKeyBytes, nil
}
