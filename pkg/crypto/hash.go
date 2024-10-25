package crypto

import (
	"crypto/sha256"
	"math/big"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Hàm chuyển đổi từ một số lớn (big.Int) thành Base62
func encodeBase62(num *big.Int) string {
	base := big.NewInt(62) // Hệ số của Base62
	zero := big.NewInt(0)
	encoded := ""

	for num.Cmp(zero) > 0 {
		mod := new(big.Int)
		num.DivMod(num, base, mod) // Tìm phần dư khi chia cho 62
		encoded = string(base62Chars[mod.Int64()]) + encoded
	}

	return encoded
}

// Hàm tạo URL rút gọn từ một URL gốc
func ShortenURL(url string) string {
	// Băm URL bằng SHA-256
	hash := sha256.New()
	hash.Write([]byte(url))
	hashBytes := hash.Sum(nil)

	// Chuyển đổi byte slice thành big.Int
	num := new(big.Int).SetBytes(hashBytes)

	// Encode thành chuỗi Base62
	return encodeBase62(num)[:8]
}
