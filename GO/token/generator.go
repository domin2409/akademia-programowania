package token

import "github.com/thanhpk/randstr"

func Generate(size int) string {
	return randstr.String(size)
}
