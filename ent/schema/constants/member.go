package constants

import (
	"database/sql/driver"
	"strconv"
)

// Gender 用户性别
type Gender uint8

// 性别枚举
const (
	// GenderUnknown 未知
	GenderUnknown Gender = iota
	// GenderFemale 女性
	GenderFemale
	// GenderMale 男性
	GenderMale
)

// Values provides list valid values for Enum.
func (Gender) Values() (kinds []string) {
	for _, s := range []Gender{GenderUnknown, GenderFemale, GenderMale} {
		kinds = append(kinds, s.String())
	}
	return
}

// Value provides the DB a string from int.
func (p Gender) Value() (driver.Value, error) {
	return p.String(), nil
}

func (p Gender) String() string {
	switch p {
	case GenderUnknown, GenderFemale, GenderMale:
		return strconv.Itoa(int(p))
	default:
		return "0"
	}
}

// Scan tells our code how to read the enum into our type.
func (p *Gender) Scan(val interface{}) error {
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
		*p = GenderUnknown
	case "1":
		*p = GenderFemale
	case "2":
		*p = GenderMale
	}
	return nil
}
