// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/ent/fieldtype"
	"github.com/facebook/ent/entc/integration/ent/file"
	"github.com/facebook/ent/entc/integration/ent/filetype"
	"github.com/facebook/ent/entc/integration/ent/group"
	"github.com/facebook/ent/entc/integration/ent/user"
	"github.com/facebook/ent/schema/field"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
}

// SetSize sets the size field.
func (fc *FileCreate) SetSize(i int) *FileCreate {
	fc.mutation.SetSize(i)
	return fc
}

// SetNillableSize sets the size field if the given value is not nil.
func (fc *FileCreate) SetNillableSize(i *int) *FileCreate {
	if i != nil {
		fc.SetSize(*i)
	}
	return fc
}

// SetName sets the name field.
func (fc *FileCreate) SetName(s string) *FileCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetUser sets the user field.
func (fc *FileCreate) SetUser(s string) *FileCreate {
	fc.mutation.SetUser(s)
	return fc
}

// SetNillableUser sets the user field if the given value is not nil.
func (fc *FileCreate) SetNillableUser(s *string) *FileCreate {
	if s != nil {
		fc.SetUser(*s)
	}
	return fc
}

// SetGroup sets the group field.
func (fc *FileCreate) SetGroup(s string) *FileCreate {
	fc.mutation.SetGroup(s)
	return fc
}

// SetNillableGroup sets the group field if the given value is not nil.
func (fc *FileCreate) SetNillableGroup(s *string) *FileCreate {
	if s != nil {
		fc.SetGroup(*s)
	}
	return fc
}

// SetOp sets the op field.
func (fc *FileCreate) SetOp(b bool) *FileCreate {
	fc.mutation.SetOp(b)
	return fc
}

// SetNillableOp sets the op field if the given value is not nil.
func (fc *FileCreate) SetNillableOp(b *bool) *FileCreate {
	if b != nil {
		fc.SetOp(*b)
	}
	return fc
}

// SetOwnerID sets the owner edge to User by id.
func (fc *FileCreate) SetOwnerID(id int) *FileCreate {
	fc.mutation.SetOwnerID(id)
	return fc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (fc *FileCreate) SetNillableOwnerID(id *int) *FileCreate {
	if id != nil {
		fc = fc.SetOwnerID(*id)
	}
	return fc
}

// SetOwner sets the owner edge to User.
func (fc *FileCreate) SetOwner(u *User) *FileCreate {
	return fc.SetOwnerID(u.ID)
}

// SetTypeID sets the type edge to FileType by id.
func (fc *FileCreate) SetTypeID(id int) *FileCreate {
	fc.mutation.SetTypeID(id)
	return fc
}

// SetNillableTypeID sets the type edge to FileType by id if the given value is not nil.
func (fc *FileCreate) SetNillableTypeID(id *int) *FileCreate {
	if id != nil {
		fc = fc.SetTypeID(*id)
	}
	return fc
}

// SetType sets the type edge to FileType.
func (fc *FileCreate) SetType(f *FileType) *FileCreate {
	return fc.SetTypeID(f.ID)
}

// AddFieldIDs adds the field edge to FieldType by ids.
func (fc *FileCreate) AddFieldIDs(ids ...int) *FileCreate {
	fc.mutation.AddFieldIDs(ids...)
	return fc
}

// AddField adds the field edges to FieldType.
func (fc *FileCreate) AddField(f ...*FieldType) *FileCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fc.AddFieldIDs(ids...)
}

// SetFileGroupID sets the file_group edge to Group by id.
func (fc *FileCreate) SetFileGroupID(id int) *FileCreate {
	fc.mutation.SetFileGroupID(id)
	return fc
}

// SetNillableFileGroupID sets the file_group edge to Group by id if the given value is not nil.
func (fc *FileCreate) SetNillableFileGroupID(id *int) *FileCreate {
	if id != nil {
		fc = fc.SetFileGroupID(*id)
	}
	return fc
}

// SetFileGroup sets the file_group edge to Group.
func (fc *FileCreate) SetFileGroup(g *Group) *FileCreate {
	return fc.SetFileGroupID(g.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	var (
		err  error
		node *File
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (fc *FileCreate) defaults() {
	if _, ok := fc.mutation.Size(); !ok {
		v := file.DefaultSize
		fc.mutation.SetSize(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if _, ok := fc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New("ent: missing required field \"size\"")}
	}
	if v, ok := fc.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf("ent: validator failed for field \"size\": %w", err)}
		}
	}
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	return nil
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fc *FileCreate) createSpec() (*File, *sqlgraph.CreateSpec) {
	var (
		_node = &File{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: file.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: file.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: file.FieldSize,
		})
		_node.Size = value
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldName,
		})
		_node.Name = value
	}
	if value, ok := fc.mutation.User(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldUser,
		})
		_node.User = &value
	}
	if value, ok := fc.mutation.Group(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldGroup,
		})
		_node.Group = value
	}
	if value, ok := fc.mutation.GetOp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: file.FieldOp,
		})
		_node.Op = value
	}
	if nodes := fc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.OwnerTable,
			Columns: []string{file.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.TypeTable,
			Columns: []string{file.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: filetype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FieldIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.FieldTable,
			Columns: []string{file.FieldColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fieldtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FileGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.FileGroupTable,
			Columns: []string{file.FileGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FileCreateBulk is the builder for creating a bulk of File entities.
type FileCreateBulk struct {
	config
	builders []*FileCreate
}

// Save creates the File entities in the database.
func (fcb *FileCreateBulk) Save(ctx context.Context) ([]*File, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*File, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (fcb *FileCreateBulk) SaveX(ctx context.Context) []*File {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
