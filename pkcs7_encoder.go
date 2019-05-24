package wxcrypt

const blockSize = 32

func pkcg7Encode(textBytes []byte) []byte {

	// 计算需要填充的位数
	amountToPad := blockSize - (len(textBytes) % blockSize)
	if amountToPad == 0 {
		amountToPad = blockSize
	}
	padChar := byte(amountToPad & 0xFF)
	// 获得补位所用的字符
	allPad := make([]byte, amountToPad, amountToPad)
	for i := 0; i < amountToPad; i++ {
		allPad[i] = padChar
	}
	textBytes = append(textBytes, allPad...)
	return textBytes
}

func pkcg7Decode(bytes []byte) []byte {
	bytesLen := len(bytes)
	pad := bytes[bytesLen-1]
	if pad < 1 || pad > byte(blockSize) {
		pad = 0
	}
	return bytes[:bytesLen-int(pad)]
}
