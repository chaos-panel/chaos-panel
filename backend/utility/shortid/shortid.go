package shortid

import (
	"math"
	"strings"
)

// ShortId 结构体
type ShortId struct {
	ID      uint64
	Short36 string
	Short52 string
	Short62 string
}

const (
	SALT_NUM       = "0123456789"
	SALT_ABC_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SALT_ABC_LOWER = "abcdefghijklmnopqrstuvwxyz"
	SALT_36        = SALT_NUM + SALT_ABC_UPPER
	SALT_52        = SALT_ABC_UPPER + SALT_ABC_LOWER
	SALT_62        = SALT_NUM + SALT_ABC_UPPER + SALT_ABC_LOWER
)

// NewShortId 构造函数
func NewShortId(id uint64) *ShortId {
	return &ShortId{
		ID:      id,
		Short36: GenShortId(id, SALT_36),
		Short52: GenShortId(id, SALT_52),
		Short62: GenShortId(id, SALT_62),
	}
}

// encode 编码函数
func GenShortId(id uint64, salt string) string {
	if id < 1 || len(salt) == 0 {
		return ""
	}
	var sb strings.Builder
	for id > 0 {
		sb.WriteByte(salt[id%uint64(len(salt))])
		id /= uint64(len(salt))
	}
	return reverse(sb.String())
}

// decode 解码函数
func ParseShortId(shortId string, salt string) *uint64 {
	if len(shortId) == 0 || len(salt) == 0 {
		return nil
	}
	shortId = reverse(shortId)
	var id uint64
	for i := 0; i < len(shortId); i++ {
		pos := strings.IndexByte(salt, shortId[i])
		if pos < 0 {
			return nil
		}
		id += uint64(pos) * uint64(math.Pow(float64(len(salt)), float64(i)))
	}
	return &id
}

// reverse 反转字符串
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
