package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Thing holds the schema definition for the Thing entity.
type Thing struct {
	ent.Schema
}

// Fields of the Thing.
func (Thing) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
	}
}

// Edges of the Thing.
func (Thing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("probed_by", ThingHTTP.Type).Ref("probes_http"),
	}
}
