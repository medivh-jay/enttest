// Code generated by entc, DO NOT EDIT.

package oauthuserinfo

const (
	// Label holds the string label denoting the oauthuserinfo type in the database.
	Label = "oauth_userinfo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccount holds the string denoting the account field in the database.
	FieldAccount = "account"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// Table holds the table name of the oauthuserinfo in the database.
	Table = "oauth_userinfos"
)

// Columns holds all SQL columns for oauthuserinfo fields.
var Columns = []string{
	FieldID,
	FieldAccount,
	FieldNickname,
	FieldAvatar,
	FieldGender,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
