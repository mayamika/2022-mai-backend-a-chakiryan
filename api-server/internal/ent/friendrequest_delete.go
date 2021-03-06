// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/friendrequest"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/predicate"
)

// FriendRequestDelete is the builder for deleting a FriendRequest entity.
type FriendRequestDelete struct {
	config
	hooks    []Hook
	mutation *FriendRequestMutation
}

// Where appends a list predicates to the FriendRequestDelete builder.
func (frd *FriendRequestDelete) Where(ps ...predicate.FriendRequest) *FriendRequestDelete {
	frd.mutation.Where(ps...)
	return frd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (frd *FriendRequestDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(frd.hooks) == 0 {
		affected, err = frd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FriendRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			frd.mutation = mutation
			affected, err = frd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(frd.hooks) - 1; i >= 0; i-- {
			if frd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = frd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, frd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (frd *FriendRequestDelete) ExecX(ctx context.Context) int {
	n, err := frd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (frd *FriendRequestDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: friendrequest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: friendrequest.FieldID,
			},
		},
	}
	if ps := frd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, frd.driver, _spec)
}

// FriendRequestDeleteOne is the builder for deleting a single FriendRequest entity.
type FriendRequestDeleteOne struct {
	frd *FriendRequestDelete
}

// Exec executes the deletion query.
func (frdo *FriendRequestDeleteOne) Exec(ctx context.Context) error {
	n, err := frdo.frd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{friendrequest.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (frdo *FriendRequestDeleteOne) ExecX(ctx context.Context) {
	frdo.frd.ExecX(ctx)
}
