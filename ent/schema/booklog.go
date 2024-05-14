package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// BookLog holds the schema definition for the BookLog entity.
type BookLog struct {
	ent.Schema
}

// Fields of the BookLog.
func (BookLog) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("action").
			Values("CREATE", "UPDATE", "DELETE", "BORROW", "RETURN").
			Immutable(),
		field.Int("user_id").
			Immutable(),
		field.Int("book_id").
			Immutable(),
		field.String("book_title").
			Immutable(),
		field.String("request_id").
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}

}

// Edges of the BookLog.
func (BookLog) Edges() []ent.Edge {
	return nil
}
