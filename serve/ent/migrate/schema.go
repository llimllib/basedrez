// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GamesColumns holds the columns for the "games" table.
	GamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 2147483647},
		{Name: "white_player_id", Type: field.TypeString, Size: 2147483647},
		{Name: "white_player_name", Type: field.TypeString, Size: 2147483647},
		{Name: "white_player_rating_diff", Type: field.TypeString, Size: 2147483647},
		{Name: "black_player_id", Type: field.TypeString, Size: 2147483647},
		{Name: "black_player_name", Type: field.TypeString, Size: 2147483647},
		{Name: "black_player_rating_diff", Type: field.TypeString, Size: 2147483647},
		{Name: "winner", Type: field.TypeString, Size: 2147483647},
		{Name: "moves", Type: field.TypeString, Size: 2147483647},
		{Name: "clock_initial", Type: field.TypeString, Size: 2147483647},
		{Name: "clock_increment", Type: field.TypeString, Size: 2147483647},
		{Name: "clock_total", Type: field.TypeString, Size: 2147483647},
		{Name: "pgn", Type: field.TypeString, Size: 2147483647},
		{Name: "rated", Type: field.TypeString, Size: 2147483647},
		{Name: "variant", Type: field.TypeString, Size: 2147483647},
		{Name: "speed", Type: field.TypeString, Size: 2147483647},
		{Name: "perf", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeString, Size: 2147483647},
		{Name: "last_move_at", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeString, Size: 2147483647},
	}
	// GamesTable holds the schema information for the "games" table.
	GamesTable = &schema.Table{
		Name:       "games",
		Columns:    GamesColumns,
		PrimaryKey: []*schema.Column{GamesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GamesTable,
	}
)

func init() {
}