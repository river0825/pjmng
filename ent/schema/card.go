package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("name"),
		field.String("description"),
		field.Int("status"),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return nil
}
