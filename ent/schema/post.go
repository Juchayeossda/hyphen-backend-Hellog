package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			MaxLen(255).
			NotEmpty(),

		field.String("content").
			SchemaType(map[string]string{
				dialect.MySQL: "longtext",
			}).
			NotEmpty(),

		field.String("preview_image").
			MaxLen(100).
			NotEmpty().
			Unique(),

		field.Bool("is_private").
			Default(false),

		field.Time("created_at").
			Default(time.Now),

		field.Time("updated_at").
			Default(time.Now),

		field.Time("deleted_at").
			Optional(), // NULL able
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("posts").
			Unique(),

		edge.To("comments", Comment.Type),
	}

}
