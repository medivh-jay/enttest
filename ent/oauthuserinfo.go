// Code generated by entc, DO NOT EDIT.

package ent

import (
	"enttest/ent/oauthuserinfo"
	"enttest/ent/schema/constants"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// OAuthUserinfo is the model entity for the OAuthUserinfo schema.
type OAuthUserinfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Account holds the value of the "account" field.
	// 三方id, 比如 google oauth open id
	Account string `json:"account,omitempty"`
	// Nickname holds the value of the "nickname" field.
	// 三方登录的昵称
	Nickname string `json:"nickname,omitempty"`
	// Avatar holds the value of the "avatar" field.
	// 头像
	Avatar string `json:"avatar,omitempty"`
	// Gender holds the value of the "gender" field.
	// 性别信息, 获取不到就算了
	Gender constants.Gender `json:"gender,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OAuthUserinfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case oauthuserinfo.FieldGender:
			values[i] = new(constants.Gender)
		case oauthuserinfo.FieldID:
			values[i] = new(sql.NullInt64)
		case oauthuserinfo.FieldAccount, oauthuserinfo.FieldNickname, oauthuserinfo.FieldAvatar:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OAuthUserinfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OAuthUserinfo fields.
func (ou *OAuthUserinfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oauthuserinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ou.ID = int(value.Int64)
		case oauthuserinfo.FieldAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account", values[i])
			} else if value.Valid {
				ou.Account = value.String
			}
		case oauthuserinfo.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				ou.Nickname = value.String
			}
		case oauthuserinfo.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				ou.Avatar = value.String
			}
		case oauthuserinfo.FieldGender:
			if value, ok := values[i].(*constants.Gender); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value != nil {
				ou.Gender = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this OAuthUserinfo.
// Note that you need to call OAuthUserinfo.Unwrap() before calling this method if this OAuthUserinfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ou *OAuthUserinfo) Update() *OAuthUserinfoUpdateOne {
	return (&OAuthUserinfoClient{config: ou.config}).UpdateOne(ou)
}

// Unwrap unwraps the OAuthUserinfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ou *OAuthUserinfo) Unwrap() *OAuthUserinfo {
	tx, ok := ou.config.driver.(*txDriver)
	if !ok {
		panic("ent: OAuthUserinfo is not a transactional entity")
	}
	ou.config.driver = tx.drv
	return ou
}

// String implements the fmt.Stringer.
func (ou *OAuthUserinfo) String() string {
	var builder strings.Builder
	builder.WriteString("OAuthUserinfo(")
	builder.WriteString(fmt.Sprintf("id=%v", ou.ID))
	builder.WriteString(", account=")
	builder.WriteString(ou.Account)
	builder.WriteString(", nickname=")
	builder.WriteString(ou.Nickname)
	builder.WriteString(", avatar=")
	builder.WriteString(ou.Avatar)
	builder.WriteString(", gender=")
	builder.WriteString(fmt.Sprintf("%v", ou.Gender))
	builder.WriteByte(')')
	return builder.String()
}

// OAuthUserinfos is a parsable slice of OAuthUserinfo.
type OAuthUserinfos []*OAuthUserinfo

func (ou OAuthUserinfos) config(cfg config) {
	for _i := range ou {
		ou[_i].config = cfg
	}
}
