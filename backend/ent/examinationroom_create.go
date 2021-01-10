// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/dhetporn/team08/ent/examinationroom"
	"github.com/dhetporn/team08/ent/operativerecord"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ExaminationroomCreate is the builder for creating a Examinationroom entity.
type ExaminationroomCreate struct {
	config
	mutation *ExaminationroomMutation
	hooks    []Hook
}

// SetExaminationroomName sets the examinationroom_name field.
func (ec *ExaminationroomCreate) SetExaminationroomName(s string) *ExaminationroomCreate {
	ec.mutation.SetExaminationroomName(s)
	return ec
}

// AddExaminationroomOperativerecordIDs adds the Examinationroom_Operativerecord edge to Operativerecord by ids.
func (ec *ExaminationroomCreate) AddExaminationroomOperativerecordIDs(ids ...int) *ExaminationroomCreate {
	ec.mutation.AddExaminationroomOperativerecordIDs(ids...)
	return ec
}

// AddExaminationroomOperativerecord adds the Examinationroom_Operativerecord edges to Operativerecord.
func (ec *ExaminationroomCreate) AddExaminationroomOperativerecord(o ...*Operativerecord) *ExaminationroomCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ec.AddExaminationroomOperativerecordIDs(ids...)
}

// Mutation returns the ExaminationroomMutation object of the builder.
func (ec *ExaminationroomCreate) Mutation() *ExaminationroomMutation {
	return ec.mutation
}

// Save creates the Examinationroom in the database.
func (ec *ExaminationroomCreate) Save(ctx context.Context) (*Examinationroom, error) {
	if _, ok := ec.mutation.ExaminationroomName(); !ok {
		return nil, &ValidationError{Name: "examinationroom_name", err: errors.New("ent: missing required field \"examinationroom_name\"")}
	}
	if v, ok := ec.mutation.ExaminationroomName(); ok {
		if err := examinationroom.ExaminationroomNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "examinationroom_name", err: fmt.Errorf("ent: validator failed for field \"examinationroom_name\": %w", err)}
		}
	}
	var (
		err  error
		node *Examinationroom
	)
	if len(ec.hooks) == 0 {
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExaminationroomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ec.mutation = mutation
			node, err = ec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *ExaminationroomCreate) SaveX(ctx context.Context) *Examinationroom {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ec *ExaminationroomCreate) sqlSave(ctx context.Context) (*Examinationroom, error) {
	e, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	e.ID = int(id)
	return e, nil
}

func (ec *ExaminationroomCreate) createSpec() (*Examinationroom, *sqlgraph.CreateSpec) {
	var (
		e     = &Examinationroom{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: examinationroom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: examinationroom.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.ExaminationroomName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: examinationroom.FieldExaminationroomName,
		})
		e.ExaminationroomName = value
	}
	if nodes := ec.mutation.ExaminationroomOperativerecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examinationroom.ExaminationroomOperativerecordTable,
			Columns: []string{examinationroom.ExaminationroomOperativerecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operativerecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return e, _spec
}