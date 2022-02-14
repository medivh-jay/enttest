package constants

import (
	"database/sql/driver"
	"strconv"
)

// AccountProvider 用户登录类型
type AccountProvider uint8

const (
	// Password 用户名和密码认证
	Password AccountProvider = iota
	// Mobile 手机号和验证码认证
	Mobile
	// MobilePassword 手机号码密码验证
	MobilePassword
	// Email 邮箱验证码认证
	Email
	// EmailPassword 邮箱和密码验证
	EmailPassword
	// OAuth 认证
	OAuth
)

// Values provides list valid values for Enum.
func (AccountProvider) Values() (kinds []string) {
	for _, s := range []AccountProvider{Password, Mobile, MobilePassword, Email, EmailPassword, OAuth} {
		kinds = append(kinds, s.String())
	}
	return
}

// Value provides the DB a string from int.
func (p AccountProvider) Value() (driver.Value, error) {
	return p.String(), nil
}

func (p AccountProvider) String() string {
	switch p {
	case Password, Mobile, MobilePassword, Email, EmailPassword, OAuth:
		return strconv.Itoa(int(p))
	default:
		return "0"
	}
}

// Scan tells our code how to read the enum into our type.
func (p *AccountProvider) Scan(val interface{}) error {
	var s string
	switch v := val.(type) {
	case nil:
		return nil
	case string:
		s = v
	case []uint8:
		s = string(v)
	}

	switch s {
	case "0":
		*p = Password
	case "1":
		*p = Mobile
	case "2":
		*p = MobilePassword
	case "3":
		*p = Email
	case "4":
		*p = EmailPassword
	case "5":
		*p = OAuth
	}
	return nil
}
