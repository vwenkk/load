package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/vwenkk/load/pkg/ent/schema/mixins"
)

type Group struct {
	ent.Schema
}

func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("group name | 组名"),
		field.String("remark").Default("").Comment("remark | 备注"),
	}
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Group) Edges() []ent.Edge {
	return nil
}

func (Group) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}

func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "load_groups"},
	}
}
