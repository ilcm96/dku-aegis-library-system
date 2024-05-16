package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// BookRequest holds the schema definition for the BookRequest entity.
type BookRequest struct {
	ent.Schema
}

// Fields of the BookRequest.
func (BookRequest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("title"),
		field.String("author"),
		field.String("publisher"),
		field.String("reason"),
		field.Enum("approved").
			Values("PENDING", "APPROVED", "REJECTED").
			Default("PENDING"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the BookRequest.
func (BookRequest) Edges() []ent.Edge {
	return nil
}
