package schema

import (
	"context"
	"regexp"
	"time"

	"enttest/ent/schema/constants"

	"enttest/ent/hook"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

const (
	// UIDDefaultIncrement uid 默认增长初始值
	UIDDefaultIncrement = 100000
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("uid").Unique(),
		field.Uint("register_at").DefaultFunc(func() uint { return uint(time.Now().Unix()) }).Immutable(),
		field.String("register_ip").MaxLen(64).Immutable(),
		field.String("avatar").MaxLen(1024).Optional(),
		field.Enum("gender").GoType(constants.Gender(0)),
		field.String("description").MaxLen(1024),
		field.String("nickname").MaxLen(20),
		field.Uint8("area_code").Optional(),
		field.Uint64("mobile").Optional(),
		field.String("email").Match(regexp.MustCompile("\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*")).Optional(),
	}
}

// Hooks of the Card.
func (m Member) Hooks() []ent.Hook {
	return []ent.Hook{
		// 生成一个用户id
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
					if s, ok := mutation.(interface{ SetUID(uint) }); ok {
						s.SetUID(uint(1))
					}
					return next.Mutate(ctx, mutation)
				})
			},
			ent.OpCreate,
		),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}
