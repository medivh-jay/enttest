// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"enttest/ent/oauth"
	"enttest/ent/schema"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// OAuth is the model entity for the OAuth schema.
type OAuth struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	// 登录平台名称
	Name string `json:"name,omitempty"`
	// Method holds the value of the "method" field.
	// 登录平台标识
	Method string `json:"method,omitempty"`
	// Icon holds the value of the "icon" field.
	// 平台icon
	Icon string `json:"icon,omitempty"`
	// Appid holds the value of the "appid" field.
	// client id
	Appid string `json:"appid,omitempty"`
	// Secret holds the value of the "secret" field.
	// client secret
	Secret string `json:"secret,omitempty"`
	// ResourceURI holds the value of the "resource_uri" field.
	// 资源路径配置
	ResourceURI schema.ResourceURI `json:"resource_uri,omitempty"`
	// Scope holds the value of the "scope" field.
	// scope
	Scope []string `json:"scope,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OAuth) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case oauth.FieldResourceURI, oauth.FieldScope:
			values[i] = new([]byte)
		case oauth.FieldID:
			values[i] = new(sql.NullInt64)
		case oauth.FieldName, oauth.FieldMethod, oauth.FieldIcon, oauth.FieldAppid, oauth.FieldSecret:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OAuth", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OAuth fields.
func (o *OAuth) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oauth.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case oauth.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		case oauth.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				o.Method = value.String
			}
		case oauth.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				o.Icon = value.String
			}
		case oauth.FieldAppid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field appid", values[i])
			} else if value.Valid {
				o.Appid = value.String
			}
		case oauth.FieldSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value.Valid {
				o.Secret = value.String
			}
		case oauth.FieldResourceURI:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field resource_uri", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.ResourceURI); err != nil {
					return fmt.Errorf("unmarshal field resource_uri: %w", err)
				}
			}
		case oauth.FieldScope:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field scope", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.Scope); err != nil {
					return fmt.Errorf("unmarshal field scope: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this OAuth.
// Note that you need to call OAuth.Unwrap() before calling this method if this OAuth
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *OAuth) Update() *OAuthUpdateOne {
	return (&OAuthClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the OAuth entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *OAuth) Unwrap() *OAuth {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: OAuth is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *OAuth) String() string {
	var builder strings.Builder
	builder.WriteString("OAuth(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", name=")
	builder.WriteString(o.Name)
	builder.WriteString(", method=")
	builder.WriteString(o.Method)
	builder.WriteString(", icon=")
	builder.WriteString(o.Icon)
	builder.WriteString(", appid=")
	builder.WriteString(o.Appid)
	builder.WriteString(", secret=")
	builder.WriteString(o.Secret)
	builder.WriteString(", resource_uri=")
	builder.WriteString(fmt.Sprintf("%v", o.ResourceURI))
	builder.WriteString(", scope=")
	builder.WriteString(fmt.Sprintf("%v", o.Scope))
	builder.WriteByte(')')
	return builder.String()
}

// OAuths is a parsable slice of OAuth.
type OAuths []*OAuth

func (o OAuths) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}