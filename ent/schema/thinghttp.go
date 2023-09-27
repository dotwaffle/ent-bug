package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// ThingHTTP holds the schema definition for the ThingHTTP entity.
type ThingHTTP struct {
	ent.Schema
}

// Fields of the ThingHTTP.
func (ThingHTTP) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
	}
}

// Edges of the ThingHTTP.
func (ThingHTTP) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("probes_http", Thing.Type),
	}
}
