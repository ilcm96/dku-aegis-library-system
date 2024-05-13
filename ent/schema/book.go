package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.String("author").
			Optional(),
		field.String("publisher").
			Optional(),
		field.Int("quantity").
			Default(1),
		field.Int("borrow").
			Default(0),
		field.String("cover"),
		field.String("category"),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
	}
}
