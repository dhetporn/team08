// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/dhetporn/team08/ent/coveredperson"
	"github.com/dhetporn/team08/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// CoveredPersonDelete is the builder for deleting a CoveredPerson entity.
type CoveredPersonDelete struct {
	config
	hooks      []Hook
	mutation   *CoveredPersonMutation
	predicates []predicate.CoveredPerson
}

// Where adds a new predicate to the delete builder.
func (cpd *CoveredPersonDelete) Where(ps ...predicate.CoveredPerson) *CoveredPersonDelete {
	cpd.predicates = append(cpd.predicates, ps...)
	return cpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cpd *CoveredPersonDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cpd.hooks) == 0 {
		affected, err = cpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoveredPersonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cpd.mutation = mutation
			affected, err = cpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cpd.hooks) - 1; i >= 0; i-- {
			mut = cpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpd *CoveredPersonDelete) ExecX(ctx context.Context) int {
	n, err := cpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cpd *CoveredPersonDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: coveredperson.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coveredperson.FieldID,
			},
		},
	}
	if ps := cpd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cpd.driver, _spec)
}

// CoveredPersonDeleteOne is the builder for deleting a single CoveredPerson entity.
type CoveredPersonDeleteOne struct {
	cpd *CoveredPersonDelete
}

// Exec executes the deletion query.
func (cpdo *CoveredPersonDeleteOne) Exec(ctx context.Context) error {
	n, err := cpdo.cpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{coveredperson.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cpdo *CoveredPersonDeleteOne) ExecX(ctx context.Context) {
	cpdo.cpd.ExecX(ctx)
}
