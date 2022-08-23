package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.Text("id"),
		field.Text("white_player_id"),
		field.Text("white_player_name"),
		field.Text("white_player_rating"),
		field.Text("white_player_rating_diff"),
		field.Text("black_player_id"),
		field.Text("black_player_name"),
		field.Text("black_player_rating"),
		field.Text("black_player_rating_diff"),
		field.Text("winner"),
		field.Text("moves"),
		field.Text("clock_initial"),
		field.Text("clock_increment"),
		field.Text("clock_total"),
		field.Text("pgn"),
		field.Text("rated"),
		field.Text("variant"),
		field.Text("speed"),
		field.Text("perf"),
		field.Text("created_at"),
		field.Text("lastMove_at"),
		field.Text("status"),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return nil
}
