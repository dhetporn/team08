// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team08/ent/coveredperson"
	"github.com/sut63/team08/ent/predicate"
	"github.com/sut63/team08/ent/schemetype"
)

// SchemeTypeUpdate is the builder for updating SchemeType entities.
type SchemeTypeUpdate struct {
	config
	hooks      []Hook
	mutation   *SchemeTypeMutation
	predicates []predicate.SchemeType
}

// Where adds a new predicate for the builder.
func (stu *SchemeTypeUpdate) Where(ps ...predicate.SchemeType) *SchemeTypeUpdate {
	stu.predicates = append(stu.predicates, ps...)
	return stu
}

// SetSchemeTypeName sets the SchemeType_Name field.
func (stu *SchemeTypeUpdate) SetSchemeTypeName(s string) *SchemeTypeUpdate {
	stu.mutation.SetSchemeTypeName(s)
	return stu
}

// AddSchemeTypeCoveredPersonIDs adds the SchemeType_CoveredPerson edge to CoveredPerson by ids.
func (stu *SchemeTypeUpdate) AddSchemeTypeCoveredPersonIDs(ids ...int) *SchemeTypeUpdate {
	stu.mutation.AddSchemeTypeCoveredPersonIDs(ids...)
	return stu
}

// AddSchemeTypeCoveredPerson adds the SchemeType_CoveredPerson edges to CoveredPerson.
func (stu *SchemeTypeUpdate) AddSchemeTypeCoveredPerson(c ...*CoveredPerson) *SchemeTypeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return stu.AddSchemeTypeCoveredPersonIDs(ids...)
}

// Mutation returns the SchemeTypeMutation object of the builder.
func (stu *SchemeTypeUpdate) Mutation() *SchemeTypeMutation {
	return stu.mutation
}

// RemoveSchemeTypeCoveredPersonIDs removes the SchemeType_CoveredPerson edge to CoveredPerson by ids.
func (stu *SchemeTypeUpdate) RemoveSchemeTypeCoveredPersonIDs(ids ...int) *SchemeTypeUpdate {
	stu.mutation.RemoveSchemeTypeCoveredPersonIDs(ids...)
	return stu
}

// RemoveSchemeTypeCoveredPerson removes SchemeType_CoveredPerson edges to CoveredPerson.
func (stu *SchemeTypeUpdate) RemoveSchemeTypeCoveredPerson(c ...*CoveredPerson) *SchemeTypeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return stu.RemoveSchemeTypeCoveredPersonIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (stu *SchemeTypeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := stu.mutation.SchemeTypeName(); ok {
		if err := schemetype.SchemeTypeNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "SchemeType_Name", err: fmt.Errorf("ent: validator failed for field \"SchemeType_Name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(stu.hooks) == 0 {
		affected, err = stu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SchemeTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			stu.mutation = mutation
			affected, err = stu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(stu.hooks) - 1; i >= 0; i-- {
			mut = stu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (stu *SchemeTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := stu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (stu *SchemeTypeUpdate) Exec(ctx context.Context) error {
	_, err := stu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stu *SchemeTypeUpdate) ExecX(ctx context.Context) {
	if err := stu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (stu *SchemeTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   schemetype.Table,
			Columns: schemetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: schemetype.FieldID,
			},
		},
	}
	if ps := stu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stu.mutation.SchemeTypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schemetype.FieldSchemeTypeName,
		})
	}
	if nodes := stu.mutation.RemovedSchemeTypeCoveredPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schemetype.SchemeTypeCoveredPersonTable,
			Columns: []string{schemetype.SchemeTypeCoveredPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coveredperson.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.SchemeTypeCoveredPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schemetype.SchemeTypeCoveredPersonTable,
			Columns: []string{schemetype.SchemeTypeCoveredPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coveredperson.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, stu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schemetype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SchemeTypeUpdateOne is the builder for updating a single SchemeType entity.
type SchemeTypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *SchemeTypeMutation
}

// SetSchemeTypeName sets the SchemeType_Name field.
func (stuo *SchemeTypeUpdateOne) SetSchemeTypeName(s string) *SchemeTypeUpdateOne {
	stuo.mutation.SetSchemeTypeName(s)
	return stuo
}

// AddSchemeTypeCoveredPersonIDs adds the SchemeType_CoveredPerson edge to CoveredPerson by ids.
func (stuo *SchemeTypeUpdateOne) AddSchemeTypeCoveredPersonIDs(ids ...int) *SchemeTypeUpdateOne {
	stuo.mutation.AddSchemeTypeCoveredPersonIDs(ids...)
	return stuo
}

// AddSchemeTypeCoveredPerson adds the SchemeType_CoveredPerson edges to CoveredPerson.
func (stuo *SchemeTypeUpdateOne) AddSchemeTypeCoveredPerson(c ...*CoveredPerson) *SchemeTypeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return stuo.AddSchemeTypeCoveredPersonIDs(ids...)
}

// Mutation returns the SchemeTypeMutation object of the builder.
func (stuo *SchemeTypeUpdateOne) Mutation() *SchemeTypeMutation {
	return stuo.mutation
}

// RemoveSchemeTypeCoveredPersonIDs removes the SchemeType_CoveredPerson edge to CoveredPerson by ids.
func (stuo *SchemeTypeUpdateOne) RemoveSchemeTypeCoveredPersonIDs(ids ...int) *SchemeTypeUpdateOne {
	stuo.mutation.RemoveSchemeTypeCoveredPersonIDs(ids...)
	return stuo
}

// RemoveSchemeTypeCoveredPerson removes SchemeType_CoveredPerson edges to CoveredPerson.
func (stuo *SchemeTypeUpdateOne) RemoveSchemeTypeCoveredPerson(c ...*CoveredPerson) *SchemeTypeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return stuo.RemoveSchemeTypeCoveredPersonIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (stuo *SchemeTypeUpdateOne) Save(ctx context.Context) (*SchemeType, error) {
	if v, ok := stuo.mutation.SchemeTypeName(); ok {
		if err := schemetype.SchemeTypeNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "SchemeType_Name", err: fmt.Errorf("ent: validator failed for field \"SchemeType_Name\": %w", err)}
		}
	}

	var (
		err  error
		node *SchemeType
	)
	if len(stuo.hooks) == 0 {
		node, err = stuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SchemeTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			stuo.mutation = mutation
			node, err = stuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(stuo.hooks) - 1; i >= 0; i-- {
			mut = stuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (stuo *SchemeTypeUpdateOne) SaveX(ctx context.Context) *SchemeType {
	st, err := stuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return st
}

// Exec executes the query on the entity.
func (stuo *SchemeTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := stuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stuo *SchemeTypeUpdateOne) ExecX(ctx context.Context) {
	if err := stuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (stuo *SchemeTypeUpdateOne) sqlSave(ctx context.Context) (st *SchemeType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   schemetype.Table,
			Columns: schemetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: schemetype.FieldID,
			},
		},
	}
	id, ok := stuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SchemeType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := stuo.mutation.SchemeTypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schemetype.FieldSchemeTypeName,
		})
	}
	if nodes := stuo.mutation.RemovedSchemeTypeCoveredPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schemetype.SchemeTypeCoveredPersonTable,
			Columns: []string{schemetype.SchemeTypeCoveredPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coveredperson.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.SchemeTypeCoveredPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schemetype.SchemeTypeCoveredPersonTable,
			Columns: []string{schemetype.SchemeTypeCoveredPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coveredperson.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	st = &SchemeType{config: stuo.config}
	_spec.Assign = st.assignValues
	_spec.ScanValues = st.scanValues()
	if err = sqlgraph.UpdateNode(ctx, stuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schemetype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return st, nil
}
