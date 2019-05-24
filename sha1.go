package wxcrypt

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

// GetSHA1 用SHA1算法生成安全签名
func GetSHA1(strs ...string) string {
	// 排序
	sort.Strings(strs)

	// 生成排序好的字符串
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
	}
	sortedStr := builder.String()

	// SHA1签名
	sha1Arr := sha1.Sum([]byte(sortedStr))
	// 转为16进制字符串
	return hex.EncodeToString(sha1Arr[:])
}
