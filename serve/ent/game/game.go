// Code generated by ent, DO NOT EDIT.

package game

const (
	// Label holds the string label denoting the game type in the database.
	Label = "game"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWhitePlayerId holds the string denoting the whiteplayerid field in the database.
	FieldWhitePlayerId = "white_player_id"
	// FieldWhitePlayerName holds the string denoting the whiteplayername field in the database.
	FieldWhitePlayerName = "white_player_name"
	// FieldWhitePlayerRatingDiff holds the string denoting the whiteplayerratingdiff field in the database.
	FieldWhitePlayerRatingDiff = "white_player_rating_diff"
	// FieldBlackPlayerId holds the string denoting the blackplayerid field in the database.
	FieldBlackPlayerId = "black_player_id"
	// FieldBlackPlayerName holds the string denoting the blackplayername field in the database.
	FieldBlackPlayerName = "black_player_name"
	// FieldBlackPlayerRatingDiff holds the string denoting the blackplayerratingdiff field in the database.
	FieldBlackPlayerRatingDiff = "black_player_rating_diff"
	// FieldWinner holds the string denoting the winner field in the database.
	FieldWinner = "winner"
	// FieldMoves holds the string denoting the moves field in the database.
	FieldMoves = "moves"
	// FieldClockInitial holds the string denoting the clockinitial field in the database.
	FieldClockInitial = "clock_initial"
	// FieldClockIncrement holds the string denoting the clockincrement field in the database.
	FieldClockIncrement = "clock_increment"
	// FieldClockTotal holds the string denoting the clocktotal field in the database.
	FieldClockTotal = "clock_total"
	// FieldPgn holds the string denoting the pgn field in the database.
	FieldPgn = "pgn"
	// FieldRated holds the string denoting the rated field in the database.
	FieldRated = "rated"
	// FieldVariant holds the string denoting the variant field in the database.
	FieldVariant = "variant"
	// FieldSpeed holds the string denoting the speed field in the database.
	FieldSpeed = "speed"
	// FieldPerf holds the string denoting the perf field in the database.
	FieldPerf = "perf"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldLastMoveAt holds the string denoting the lastmoveat field in the database.
	FieldLastMoveAt = "last_move_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the game in the database.
	Table = "games"
)

// Columns holds all SQL columns for game fields.
var Columns = []string{
	FieldID,
	FieldWhitePlayerId,
	FieldWhitePlayerName,
	FieldWhitePlayerRatingDiff,
	FieldBlackPlayerId,
	FieldBlackPlayerName,
	FieldBlackPlayerRatingDiff,
	FieldWinner,
	FieldMoves,
	FieldClockInitial,
	FieldClockIncrement,
	FieldClockTotal,
	FieldPgn,
	FieldRated,
	FieldVariant,
	FieldSpeed,
	FieldPerf,
	FieldCreatedAt,
	FieldLastMoveAt,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
