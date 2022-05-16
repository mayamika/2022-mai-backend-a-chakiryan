// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/predicate"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/userauth"
)

// UserAuthUpdate is the builder for updating UserAuth entities.
type UserAuthUpdate struct {
	config
	hooks    []Hook
	mutation *UserAuthMutation
}

// Where appends a list predicates to the UserAuthUpdate builder.
func (uau *UserAuthUpdate) Where(ps ...predicate.UserAuth) *UserAuthUpdate {
	uau.mutation.Where(ps...)
	return uau
}

// Mutation returns the UserAuthMutation object of the builder.
func (uau *UserAuthUpdate) Mutation() *UserAuthMutation {
	return uau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uau *UserAuthUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uau.hooks) == 0 {
		affected, err = uau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uau.mutation = mutation
			affected, err = uau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uau.hooks) - 1; i >= 0; i-- {
			if uau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uau *UserAuthUpdate) SaveX(ctx context.Context) int {
	affected, err := uau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uau *UserAuthUpdate) Exec(ctx context.Context) error {
	_, err := uau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uau *UserAuthUpdate) ExecX(ctx context.Context) {
	if err := uau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uau *UserAuthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userauth.Table,
			Columns: userauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userauth.FieldID,
			},
		},
	}
	if ps := uau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserAuthUpdateOne is the builder for updating a single UserAuth entity.
type UserAuthUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserAuthMutation
}

// Mutation returns the UserAuthMutation object of the builder.
func (uauo *UserAuthUpdateOne) Mutation() *UserAuthMutation {
	return uauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uauo *UserAuthUpdateOne) Select(field string, fields ...string) *UserAuthUpdateOne {
	uauo.fields = append([]string{field}, fields...)
	return uauo
}

// Save executes the query and returns the updated UserAuth entity.
func (uauo *UserAuthUpdateOne) Save(ctx context.Context) (*UserAuth, error) {
	var (
		err  error
		node *UserAuth
	)
	if len(uauo.hooks) == 0 {
		node, err = uauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uauo.mutation = mutation
			node, err = uauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uauo.hooks) - 1; i >= 0; i-- {
			if uauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserAuth)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserAuthMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uauo *UserAuthUpdateOne) SaveX(ctx context.Context) *UserAuth {
	node, err := uauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uauo *UserAuthUpdateOne) Exec(ctx context.Context) error {
	_, err := uauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uauo *UserAuthUpdateOne) ExecX(ctx context.Context) {
	if err := uauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uauo *UserAuthUpdateOne) sqlSave(ctx context.Context) (_node *UserAuth, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userauth.Table,
			Columns: userauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userauth.FieldID,
			},
		},
	}
	id, ok := uauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserAuth.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userauth.FieldID)
		for _, f := range fields {
			if !userauth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &UserAuth{config: uauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
