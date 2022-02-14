package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OAuth holds the schema definition for the OAuth entity.
type OAuth struct {
	ent.Schema
}

// ResourceURI 资源获取路径
type ResourceURI struct {
	AuthURL     string
	TokenURL    string
	UserinfoURL string
}

// Fields of the OAuth.
func (OAuth) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("登录平台名称"),
		field.String("method").Comment("登录平台标识"),
		field.String("icon").Comment("平台icon"),
		field.String("appid").Comment("client id"),
		field.String("secret").Comment("client secret"),
		field.JSON("resource_uri", ResourceURI{}).Comment("资源路径配置"),
		field.Strings("scope").Comment("scope"),
	}
}

// Edges of the OAuth.
func (OAuth) Edges() []ent.Edge {
	return nil
}
