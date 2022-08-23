// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/llimllib/basedrez/ent/game"
)

// GameCreate is the builder for creating a Game entity.
type GameCreate struct {
	config
	mutation *GameMutation
	hooks    []Hook
}

// SetWhitePlayerId sets the "whitePlayerId" field.
func (gc *GameCreate) SetWhitePlayerId(s string) *GameCreate {
	gc.mutation.SetWhitePlayerId(s)
	return gc
}

// SetWhitePlayerName sets the "whitePlayerName" field.
func (gc *GameCreate) SetWhitePlayerName(s string) *GameCreate {
	gc.mutation.SetWhitePlayerName(s)
	return gc
}

// SetWhitePlayerRatingDiff sets the "whitePlayerRatingDiff" field.
func (gc *GameCreate) SetWhitePlayerRatingDiff(s string) *GameCreate {
	gc.mutation.SetWhitePlayerRatingDiff(s)
	return gc
}

// SetBlackPlayerId sets the "blackPlayerId" field.
func (gc *GameCreate) SetBlackPlayerId(s string) *GameCreate {
	gc.mutation.SetBlackPlayerId(s)
	return gc
}

// SetBlackPlayerName sets the "blackPlayerName" field.
func (gc *GameCreate) SetBlackPlayerName(s string) *GameCreate {
	gc.mutation.SetBlackPlayerName(s)
	return gc
}

// SetBlackPlayerRatingDiff sets the "blackPlayerRatingDiff" field.
func (gc *GameCreate) SetBlackPlayerRatingDiff(s string) *GameCreate {
	gc.mutation.SetBlackPlayerRatingDiff(s)
	return gc
}

// SetWinner sets the "winner" field.
func (gc *GameCreate) SetWinner(s string) *GameCreate {
	gc.mutation.SetWinner(s)
	return gc
}

// SetMoves sets the "moves" field.
func (gc *GameCreate) SetMoves(s string) *GameCreate {
	gc.mutation.SetMoves(s)
	return gc
}

// SetClockInitial sets the "clockInitial" field.
func (gc *GameCreate) SetClockInitial(s string) *GameCreate {
	gc.mutation.SetClockInitial(s)
	return gc
}

// SetClockIncrement sets the "clockIncrement" field.
func (gc *GameCreate) SetClockIncrement(s string) *GameCreate {
	gc.mutation.SetClockIncrement(s)
	return gc
}

// SetClockTotal sets the "clockTotal" field.
func (gc *GameCreate) SetClockTotal(s string) *GameCreate {
	gc.mutation.SetClockTotal(s)
	return gc
}

// SetPgn sets the "pgn" field.
func (gc *GameCreate) SetPgn(s string) *GameCreate {
	gc.mutation.SetPgn(s)
	return gc
}

// SetRated sets the "rated" field.
func (gc *GameCreate) SetRated(s string) *GameCreate {
	gc.mutation.SetRated(s)
	return gc
}

// SetVariant sets the "variant" field.
func (gc *GameCreate) SetVariant(s string) *GameCreate {
	gc.mutation.SetVariant(s)
	return gc
}

// SetSpeed sets the "speed" field.
func (gc *GameCreate) SetSpeed(s string) *GameCreate {
	gc.mutation.SetSpeed(s)
	return gc
}

// SetPerf sets the "perf" field.
func (gc *GameCreate) SetPerf(s string) *GameCreate {
	gc.mutation.SetPerf(s)
	return gc
}

// SetCreatedAt sets the "createdAt" field.
func (gc *GameCreate) SetCreatedAt(s string) *GameCreate {
	gc.mutation.SetCreatedAt(s)
	return gc
}

// SetLastMoveAt sets the "lastMoveAt" field.
func (gc *GameCreate) SetLastMoveAt(s string) *GameCreate {
	gc.mutation.SetLastMoveAt(s)
	return gc
}

// SetStatus sets the "status" field.
func (gc *GameCreate) SetStatus(s string) *GameCreate {
	gc.mutation.SetStatus(s)
	return gc
}

// SetID sets the "id" field.
func (gc *GameCreate) SetID(s string) *GameCreate {
	gc.mutation.SetID(s)
	return gc
}

// Mutation returns the GameMutation object of the builder.
func (gc *GameCreate) Mutation() *GameMutation {
	return gc.mutation
}

// Save creates the Game in the database.
func (gc *GameCreate) Save(ctx context.Context) (*Game, error) {
	var (
		err  error
		node *Game
	)
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			if node, err = gc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			if gc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Game)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GameMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GameCreate) SaveX(ctx context.Context) *Game {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GameCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GameCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GameCreate) check() error {
	if _, ok := gc.mutation.WhitePlayerId(); !ok {
		return &ValidationError{Name: "whitePlayerId", err: errors.New(`ent: missing required field "Game.whitePlayerId"`)}
	}
	if _, ok := gc.mutation.WhitePlayerName(); !ok {
		return &ValidationError{Name: "whitePlayerName", err: errors.New(`ent: missing required field "Game.whitePlayerName"`)}
	}
	if _, ok := gc.mutation.WhitePlayerRatingDiff(); !ok {
		return &ValidationError{Name: "whitePlayerRatingDiff", err: errors.New(`ent: missing required field "Game.whitePlayerRatingDiff"`)}
	}
	if _, ok := gc.mutation.BlackPlayerId(); !ok {
		return &ValidationError{Name: "blackPlayerId", err: errors.New(`ent: missing required field "Game.blackPlayerId"`)}
	}
	if _, ok := gc.mutation.BlackPlayerName(); !ok {
		return &ValidationError{Name: "blackPlayerName", err: errors.New(`ent: missing required field "Game.blackPlayerName"`)}
	}
	if _, ok := gc.mutation.BlackPlayerRatingDiff(); !ok {
		return &ValidationError{Name: "blackPlayerRatingDiff", err: errors.New(`ent: missing required field "Game.blackPlayerRatingDiff"`)}
	}
	if _, ok := gc.mutation.Winner(); !ok {
		return &ValidationError{Name: "winner", err: errors.New(`ent: missing required field "Game.winner"`)}
	}
	if _, ok := gc.mutation.Moves(); !ok {
		return &ValidationError{Name: "moves", err: errors.New(`ent: missing required field "Game.moves"`)}
	}
	if _, ok := gc.mutation.ClockInitial(); !ok {
		return &ValidationError{Name: "clockInitial", err: errors.New(`ent: missing required field "Game.clockInitial"`)}
	}
	if _, ok := gc.mutation.ClockIncrement(); !ok {
		return &ValidationError{Name: "clockIncrement", err: errors.New(`ent: missing required field "Game.clockIncrement"`)}
	}
	if _, ok := gc.mutation.ClockTotal(); !ok {
		return &ValidationError{Name: "clockTotal", err: errors.New(`ent: missing required field "Game.clockTotal"`)}
	}
	if _, ok := gc.mutation.Pgn(); !ok {
		return &ValidationError{Name: "pgn", err: errors.New(`ent: missing required field "Game.pgn"`)}
	}
	if _, ok := gc.mutation.Rated(); !ok {
		return &ValidationError{Name: "rated", err: errors.New(`ent: missing required field "Game.rated"`)}
	}
	if _, ok := gc.mutation.Variant(); !ok {
		return &ValidationError{Name: "variant", err: errors.New(`ent: missing required field "Game.variant"`)}
	}
	if _, ok := gc.mutation.Speed(); !ok {
		return &ValidationError{Name: "speed", err: errors.New(`ent: missing required field "Game.speed"`)}
	}
	if _, ok := gc.mutation.Perf(); !ok {
		return &ValidationError{Name: "perf", err: errors.New(`ent: missing required field "Game.perf"`)}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Game.createdAt"`)}
	}
	if _, ok := gc.mutation.LastMoveAt(); !ok {
		return &ValidationError{Name: "lastMoveAt", err: errors.New(`ent: missing required field "Game.lastMoveAt"`)}
	}
	if _, ok := gc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Game.status"`)}
	}
	return nil
}

func (gc *GameCreate) sqlSave(ctx context.Context) (*Game, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Game.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (gc *GameCreate) createSpec() (*Game, *sqlgraph.CreateSpec) {
	var (
		_node = &Game{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: game.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: game.FieldID,
			},
		}
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.WhitePlayerId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldWhitePlayerId,
		})
		_node.WhitePlayerId = value
	}
	if value, ok := gc.mutation.WhitePlayerName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldWhitePlayerName,
		})
		_node.WhitePlayerName = value
	}
	if value, ok := gc.mutation.WhitePlayerRatingDiff(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldWhitePlayerRatingDiff,
		})
		_node.WhitePlayerRatingDiff = value
	}
	if value, ok := gc.mutation.BlackPlayerId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldBlackPlayerId,
		})
		_node.BlackPlayerId = value
	}
	if value, ok := gc.mutation.BlackPlayerName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldBlackPlayerName,
		})
		_node.BlackPlayerName = value
	}
	if value, ok := gc.mutation.BlackPlayerRatingDiff(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldBlackPlayerRatingDiff,
		})
		_node.BlackPlayerRatingDiff = value
	}
	if value, ok := gc.mutation.Winner(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldWinner,
		})
		_node.Winner = value
	}
	if value, ok := gc.mutation.Moves(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldMoves,
		})
		_node.Moves = value
	}
	if value, ok := gc.mutation.ClockInitial(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldClockInitial,
		})
		_node.ClockInitial = value
	}
	if value, ok := gc.mutation.ClockIncrement(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldClockIncrement,
		})
		_node.ClockIncrement = value
	}
	if value, ok := gc.mutation.ClockTotal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldClockTotal,
		})
		_node.ClockTotal = value
	}
	if value, ok := gc.mutation.Pgn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldPgn,
		})
		_node.Pgn = value
	}
	if value, ok := gc.mutation.Rated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldRated,
		})
		_node.Rated = value
	}
	if value, ok := gc.mutation.Variant(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldVariant,
		})
		_node.Variant = value
	}
	if value, ok := gc.mutation.Speed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldSpeed,
		})
		_node.Speed = value
	}
	if value, ok := gc.mutation.Perf(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldPerf,
		})
		_node.Perf = value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.LastMoveAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldLastMoveAt,
		})
		_node.LastMoveAt = value
	}
	if value, ok := gc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldStatus,
		})
		_node.Status = value
	}
	return _node, _spec
}

// GameCreateBulk is the builder for creating many Game entities in bulk.
type GameCreateBulk struct {
	config
	builders []*GameCreate
}

// Save creates the Game entities in the database.
func (gcb *GameCreateBulk) Save(ctx context.Context) ([]*Game, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Game, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GameMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GameCreateBulk) SaveX(ctx context.Context) []*Game {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GameCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GameCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}