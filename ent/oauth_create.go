// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"enttest/ent/oauth"
	"enttest/ent/schema"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OAuthCreate is the builder for creating a OAuth entity.
type OAuthCreate struct {
	config
	mutation *OAuthMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (oc *OAuthCreate) SetName(s string) *OAuthCreate {
	oc.mutation.SetName(s)
	return oc
}

// SetMethod sets the "method" field.
func (oc *OAuthCreate) SetMethod(s string) *OAuthCreate {
	oc.mutation.SetMethod(s)
	return oc
}

// SetIcon sets the "icon" field.
func (oc *OAuthCreate) SetIcon(s string) *OAuthCreate {
	oc.mutation.SetIcon(s)
	return oc
}

// SetAppid sets the "appid" field.
func (oc *OAuthCreate) SetAppid(s string) *OAuthCreate {
	oc.mutation.SetAppid(s)
	return oc
}

// SetSecret sets the "secret" field.
func (oc *OAuthCreate) SetSecret(s string) *OAuthCreate {
	oc.mutation.SetSecret(s)
	return oc
}

// SetResourceURI sets the "resource_uri" field.
func (oc *OAuthCreate) SetResourceURI(su schema.ResourceURI) *OAuthCreate {
	oc.mutation.SetResourceURI(su)
	return oc
}

// SetScope sets the "scope" field.
func (oc *OAuthCreate) SetScope(s []string) *OAuthCreate {
	oc.mutation.SetScope(s)
	return oc
}

// Mutation returns the OAuthMutation object of the builder.
func (oc *OAuthCreate) Mutation() *OAuthMutation {
	return oc.mutation
}

// Save creates the OAuth in the database.
func (oc *OAuthCreate) Save(ctx context.Context) (*OAuth, error) {
	var (
		err  error
		node *OAuth
	)
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			if node, err = oc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			if oc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OAuthCreate) SaveX(ctx context.Context) *OAuth {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OAuthCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OAuthCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OAuthCreate) check() error {
	if _, ok := oc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := oc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "method"`)}
	}
	if _, ok := oc.mutation.Icon(); !ok {
		return &ValidationError{Name: "icon", err: errors.New(`ent: missing required field "icon"`)}
	}
	if _, ok := oc.mutation.Appid(); !ok {
		return &ValidationError{Name: "appid", err: errors.New(`ent: missing required field "appid"`)}
	}
	if _, ok := oc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New(`ent: missing required field "secret"`)}
	}
	if _, ok := oc.mutation.ResourceURI(); !ok {
		return &ValidationError{Name: "resource_uri", err: errors.New(`ent: missing required field "resource_uri"`)}
	}
	if _, ok := oc.mutation.Scope(); !ok {
		return &ValidationError{Name: "scope", err: errors.New(`ent: missing required field "scope"`)}
	}
	return nil
}

func (oc *OAuthCreate) sqlSave(ctx context.Context) (*OAuth, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oc *OAuthCreate) createSpec() (*OAuth, *sqlgraph.CreateSpec) {
	var (
		_node = &OAuth{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: oauth.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oauth.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth.FieldName,
		})
		_node.Name = value
	}
	if value, ok := oc.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := oc.mutation.Icon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth.FieldIcon,
		})
		_node.Icon = value
	}
	if value, ok := oc.mutation.Appid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth.FieldAppid,
		})
		_node.Appid = value
	}
	if value, ok := oc.mutation.Secret(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth.FieldSecret,
		})
		_node.Secret = value
	}
	if value, ok := oc.mutation.ResourceURI(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oauth.FieldResourceURI,
		})
		_node.ResourceURI = value
	}
	if value, ok := oc.mutation.Scope(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oauth.FieldScope,
		})
		_node.Scope = value
	}
	return _node, _spec
}

// OAuthCreateBulk is the builder for creating many OAuth entities in bulk.
type OAuthCreateBulk struct {
	config
	builders []*OAuthCreate
}

// Save creates the OAuth entities in the database.
func (ocb *OAuthCreateBulk) Save(ctx context.Context) ([]*OAuth, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*OAuth, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OAuthMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OAuthCreateBulk) SaveX(ctx context.Context) []*OAuth {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OAuthCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OAuthCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
