// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/examples/privacytenant/ent/group"
	"github.com/facebook/ent/examples/privacytenant/ent/predicate"
	"github.com/facebook/ent/examples/privacytenant/ent/tenant"
	"github.com/facebook/ent/examples/privacytenant/ent/user"
	"github.com/facebook/ent/schema/field"
)

// TenantUpdate is the builder for updating Tenant entities.
type TenantUpdate struct {
	config
	hooks    []Hook
	mutation *TenantMutation
}

// Where adds a new predicate for the builder.
func (tu *TenantUpdate) Where(ps ...predicate.Tenant) *TenantUpdate {
	tu.mutation.predicates = append(tu.mutation.predicates, ps...)
	return tu
}

// SetName sets the name field.
func (tu *TenantUpdate) SetName(s string) *TenantUpdate {
	tu.mutation.SetName(s)
	return tu
}

// AddGroupIDs adds the groups edge to Group by ids.
func (tu *TenantUpdate) AddGroupIDs(ids ...int) *TenantUpdate {
	tu.mutation.AddGroupIDs(ids...)
	return tu
}

// AddGroups adds the groups edges to Group.
func (tu *TenantUpdate) AddGroups(g ...*Group) *TenantUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.AddGroupIDs(ids...)
}

// AddUserIDs adds the users edge to User by ids.
func (tu *TenantUpdate) AddUserIDs(ids ...int) *TenantUpdate {
	tu.mutation.AddUserIDs(ids...)
	return tu
}

// AddUsers adds the users edges to User.
func (tu *TenantUpdate) AddUsers(u ...*User) *TenantUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddUserIDs(ids...)
}

// Mutation returns the TenantMutation object of the builder.
func (tu *TenantUpdate) Mutation() *TenantMutation {
	return tu.mutation
}

// ClearGroups clears all "groups" edges to type Group.
func (tu *TenantUpdate) ClearGroups() *TenantUpdate {
	tu.mutation.ClearGroups()
	return tu
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (tu *TenantUpdate) RemoveGroupIDs(ids ...int) *TenantUpdate {
	tu.mutation.RemoveGroupIDs(ids...)
	return tu
}

// RemoveGroups removes groups edges to Group.
func (tu *TenantUpdate) RemoveGroups(g ...*Group) *TenantUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.RemoveGroupIDs(ids...)
}

// ClearUsers clears all "users" edges to type User.
func (tu *TenantUpdate) ClearUsers() *TenantUpdate {
	tu.mutation.ClearUsers()
	return tu
}

// RemoveUserIDs removes the users edge to User by ids.
func (tu *TenantUpdate) RemoveUserIDs(ids ...int) *TenantUpdate {
	tu.mutation.RemoveUserIDs(ids...)
	return tu
}

// RemoveUsers removes users edges to User.
func (tu *TenantUpdate) RemoveUsers(u ...*User) *TenantUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TenantUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TenantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TenantUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TenantUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TenantUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TenantUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := tenant.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (tu *TenantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tenant.Table,
			Columns: tenant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tenant.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tenant.FieldName,
		})
	}
	if tu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tenant.GroupsTable,
			Columns: []string{tenant.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !tu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tenant.GroupsTable,
			Columns: []string{tenant.GroupsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tenant.GroupsTable,
			Columns: []string{tenant.GroupsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tenant.UsersTable,
			Columns: []string{tenant.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !tu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tenant.UsersTable,
			Columns: []string{tenant.UsersColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tenant.UsersTable,
			Columns: []string{tenant.UsersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tenant.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TenantUpdateOne is the builder for updating a single Tenant entity.
type TenantUpdateOne struct {
	config
	hooks    []Hook
	mutation *TenantMutation
}

// SetName sets the name field.
func (tuo *TenantUpdateOne) SetName(s string) *TenantUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// AddGroupIDs adds the groups edge to Group by ids.
func (tuo *TenantUpdateOne) AddGroupIDs(ids ...int) *TenantUpdateOne {
	tuo.mutation.AddGroupIDs(ids...)
	return tuo
}

// AddGroups adds the groups edges to Group.
func (tuo *TenantUpdateOne) AddGroups(g ...*Group) *TenantUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.AddGroupIDs(ids...)
}

// AddUserIDs adds the users edge to User by ids.
func (tuo *TenantUpdateOne) AddUserIDs(ids ...int) *TenantUpdateOne {
	tuo.mutation.AddUserIDs(ids...)
	return tuo
}

// AddUsers adds the users edges to User.
func (tuo *TenantUpdateOne) AddUsers(u ...*User) *TenantUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddUserIDs(ids...)
}

// Mutation returns the TenantMutation object of the builder.
func (tuo *TenantUpdateOne) Mutation() *TenantMutation {
	return tuo.mutation
}

// ClearGroups clears all "groups" edges to type Group.
func (tuo *TenantUpdateOne) ClearGroups() *TenantUpdateOne {
	tuo.mutation.ClearGroups()
	return tuo
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (tuo *TenantUpdateOne) RemoveGroupIDs(ids ...int) *TenantUpdateOne {
	tuo.mutation.RemoveGroupIDs(ids...)
	return tuo
}

// RemoveGroups removes groups edges to Group.
func (tuo *TenantUpdateOne) RemoveGroups(g ...*Group) *TenantUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.RemoveGroupIDs(ids...)
}

// ClearUsers clears all "users" edges to type User.
func (tuo *TenantUpdateOne) ClearUsers() *TenantUpdateOne {
	tuo.mutation.ClearUsers()
	return tuo
}

// RemoveUserIDs removes the users edge to User by ids.
func (tuo *TenantUpdateOne) RemoveUserIDs(ids ...int) *TenantUpdateOne {
	tuo.mutation.RemoveUserIDs(ids...)
	return tuo
}

// RemoveUsers removes users edges to User.
func (tuo *TenantUpdateOne) RemoveUsers(u ...*User) *TenantUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveUserIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (tuo *TenantUpdateOne) Save(ctx context.Context) (*Tenant, error) {
	var (
		err  error
		node *Tenant
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TenantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TenantUpdateOne) SaveX(ctx context.Context) *Tenant {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TenantUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TenantUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TenantUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := tenant.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (tuo *TenantUpdateOne) sqlSave(ctx context.Context) (_node *Tenant, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tenant.Table,
			Columns: tenant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tenant.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Tenant.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tenant.FieldName,
		})
	}
	if tuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tenant.GroupsTable,
			Columns: []string{tenant.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !tuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tenant.GroupsTable,
			Columns: []string{tenant.GroupsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tenant.GroupsTable,
			Columns: []string{tenant.GroupsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tenant.UsersTable,
			Columns: []string{tenant.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !tuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tenant.UsersTable,
			Columns: []string{tenant.UsersColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tenant.UsersTable,
			Columns: []string{tenant.UsersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tenant{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tenant.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
