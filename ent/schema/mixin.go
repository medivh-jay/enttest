package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"git.geefun.cc/medivh/netswap/pkg/unique"
	"time"
)

// DefaultMixin implements the ent.Mixin for sharing
// 通用 created_at   updated_at
type DefaultMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (m DefaultMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique().DefaultFunc(m.uniqueID).Annotations(m.disableIncremental()),
		field.Int64("created_at").DefaultFunc(m.nowTime).Immutable(),
		field.Int64("updated_at").DefaultFunc(m.nowTime).UpdateDefault(m.nowTime),
	}
}

func (DefaultMixin) uniqueID() uint64 {
	return unique.ID()
}

func (DefaultMixin) disableIncremental() entsql.Annotation {
	incrementalEnabled := false
	return entsql.Annotation{
		Incremental: &incrementalEnabled,
	}
}

func (DefaultMixin) nowTime() int64 {
	return time.Now().Unix()
}
