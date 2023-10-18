package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id_from_user_ms").
			Positive().
			Unique(),

		field.String("name").
			MaxLen(12). // name min(2) max(12) {2 ~ 12}
			NotEmpty(). // Not Empty
			Unique(),

		field.String("profile_image").
			MaxLen(64). // siss get endpoint + uuid max len
			NotEmpty().
			Unique(),

		field.Time("joined_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("comments", Comment.Type),
	}
}
