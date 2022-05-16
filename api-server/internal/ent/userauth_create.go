// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/userauth"
)

// UserAuthCreate is the builder for creating a UserAuth entity.
type UserAuthCreate struct {
	config
	mutation *UserAuthMutation
	hooks    []Hook
}

// Mutation returns the UserAuthMutation object of the builder.
func (uac *UserAuthCreate) Mutation() *UserAuthMutation {
	return uac.mutation
}

// Save creates the UserAuth in the database.
func (uac *UserAuthCreate) Save(ctx context.Context) (*UserAuth, error) {
	var (
		err  error
		node *UserAuth
	)
	if len(uac.hooks) == 0 {
		if err = uac.check(); err != nil {
			return nil, err
		}
		node, err = uac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uac.check(); err != nil {
				return nil, err
			}
			uac.mutation = mutation
			if node, err = uac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uac.hooks) - 1; i >= 0; i-- {
			if uac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uac.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (uac *UserAuthCreate) SaveX(ctx context.Context) *UserAuth {
	v, err := uac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uac *UserAuthCreate) Exec(ctx context.Context) error {
	_, err := uac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uac *UserAuthCreate) ExecX(ctx context.Context) {
	if err := uac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uac *UserAuthCreate) check() error {
	return nil
}

func (uac *UserAuthCreate) sqlSave(ctx context.Context) (*UserAuth, error) {
	_node, _spec := uac.createSpec()
	if err := sqlgraph.CreateNode(ctx, uac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uac *UserAuthCreate) createSpec() (*UserAuth, *sqlgraph.CreateSpec) {
	var (
		_node = &UserAuth{config: uac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userauth.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userauth.FieldID,
			},
		}
	)
	return _node, _spec
}

// UserAuthCreateBulk is the builder for creating many UserAuth entities in bulk.
type UserAuthCreateBulk struct {
	config
	builders []*UserAuthCreate
}

// Save creates the UserAuth entities in the database.
func (uacb *UserAuthCreateBulk) Save(ctx context.Context) ([]*UserAuth, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uacb.builders))
	nodes := make([]*UserAuth, len(uacb.builders))
	mutators := make([]Mutator, len(uacb.builders))
	for i := range uacb.builders {
		func(i int, root context.Context) {
			builder := uacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserAuthMutation)
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
					_, err = mutators[i+1].Mutate(root, uacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uacb *UserAuthCreateBulk) SaveX(ctx context.Context) []*UserAuth {
	v, err := uacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uacb *UserAuthCreateBulk) Exec(ctx context.Context) error {
	_, err := uacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uacb *UserAuthCreateBulk) ExecX(ctx context.Context) {
	if err := uacb.Exec(ctx); err != nil {
		panic(err)
	}
}
