package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Range(30000000, 39999999),
		field.String("password").
			Sensitive().
			NotEmpty(),
		field.String("name").
			NotEmpty().
			MinLen(2),
		field.Enum("status").
			Values("WITHDRAW", "PENDING", "APPROVED", "ADMIN").
			Default("PENDING"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("book", Book.Type).
			Ref("user"),
	}
}
