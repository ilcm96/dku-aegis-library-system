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
			Values("CREATE", "UPDATE", "DELETE", "BORROW", "RETURN"),
		field.Int("user_id"),
		field.Int("book_id"),
		field.String("book_title"),
		field.String("request_id"),
		field.Time("created_at").
			Default(time.Now),
	}

}

// Edges of the BookLog.
func (BookLog) Edges() []ent.Edge {
	return nil
}
