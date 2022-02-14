package schema

import (
	"enttest/ent/schema/constants"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OAuthUserinfo holds the schema definition for the OAuthUserinfo entity.
type OAuthUserinfo struct {
	ent.Schema
}

type Userinfo struct {
	Nickname string           `bson:"nickname" json:"nickname"`
	Avatar   string           `bson:"avatar" json:"avatar"`
	Gender   constants.Gender `bson:"gender" json:"gender"`
}

// Fields of the OAuthUserinfo.
func (OAuthUserinfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("account").Comment("三方id, 比如 google oauth open id"),
		field.String("nickname").Comment("三方登录的昵称"),
		field.String("avatar").Comment("头像"),
		field.String("gender").GoType(constants.Gender(0)).Comment("性别信息, 获取不到就算了"),
	}
}

// Edges of the OAuthUserinfo.
func (OAuthUserinfo) Edges() []ent.Edge {
	return nil
}
