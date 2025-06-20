package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMd5(t *testing.T) {
	assert.Eq(t, "e10adc3949ba59abbe56e057f20f883e", strutil.Md5("123456"))
	assert.Eq(t, "e10adc3949ba59abbe56e057f20f883e", strutil.MD5("123456"))
	assert.Eq(t, "e10adc3949ba59abbe56e057f20f883e", strutil.MD5([]byte("123456")))
	assert.Eq(t, "a906449d5769fa7361d7ecc6aa3f6d28", strutil.GenMd5("123abc"))
	assert.Eq(t, "289dff07669d7a23de0ef88d2f7129e7", strutil.GenMd5(234))

	// short md5
	assert.Eq(t, "ac59075b964b0715", strutil.ShortMd5(123))
	assert.Eq(t, "3cd24fb0d6963f7d", strutil.ShortMd5("abc"))
}

func TestMd5Base62(t *testing.T) {
	aLongStr := strutil.Base64Chars + strutil.Base62Chars
	tests := []struct {
		input    any
		simple   string
		b62, md5 string
	}{
		{123, "wIZAMr7tqd7lJzdO", "YIbZFzz9CzW9dMgNgCUN2", "202cb962ac59075b964b07152d234b70"},
		{"hello", "v32G2dGUZPxlgnbm", "2PY6zCz4ZUQkeY47zJ8txM", "5d41402abc4b2a76b9719d911017c592"},
		{"123abc", "J66xpH2RztOcK1LE", "58WzcshJkGBMniJBznk6ww", "a906449d5769fa7361d7ecc6aa3f6d28"},
		{"strutil.Md5Base62", "TQI9ql4vlqTAiZNj", "7m3F1lWzMyR8HnRJvNcINP", "f1b02c09969180d91558b362127b3151"},
		{aLongStr, "oAnENhaOEK6bEkUd", "Kt5xwUm0jY3trXinINGZ1", "18a055e23111c4ec28aafe49a49076c7"},
	}

	for _, tt := range tests {
		b62 := strutil.Md5Base62(tt.input)
		assert.Eq(t, tt.b62, b62)
		assert.Eq(t, tt.md5, strutil.Md5(tt.input))
		assert.Eq(t, tt.simple, strutil.Md5Simple(tt.input))
	}
}

func TestHashPasswd(t *testing.T) {
	key := "ot54c"
	pwd := "abc123456"

	msgMac := strutil.HashPasswd(pwd, key)
	dump.P(msgMac)
	assert.NotEmpty(t, msgMac)
	assert.True(t, strutil.VerifyPasswd(msgMac, pwd, key))
}
