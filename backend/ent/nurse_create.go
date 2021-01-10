// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team08/ent/nurse"
	"github.com/sut63/team08/ent/operativerecord"
	"github.com/sut63/team08/ent/prescription"
	"github.com/sut63/team08/ent/rent"
)

// NurseCreate is the builder for creating a Nurse entity.
type NurseCreate struct {
	config
	mutation *NurseMutation
	hooks    []Hook
}

// SetNurseName sets the nurse_name field.
func (nc *NurseCreate) SetNurseName(s string) *NurseCreate {
	nc.mutation.SetNurseName(s)
	return nc
}

// SetNurseEmail sets the nurse_email field.
func (nc *NurseCreate) SetNurseEmail(s string) *NurseCreate {
	nc.mutation.SetNurseEmail(s)
	return nc
}

// SetNursePassword sets the nurse_password field.
func (nc *NurseCreate) SetNursePassword(s string) *NurseCreate {
	nc.mutation.SetNursePassword(s)
	return nc
}

// SetNurseTel sets the nurse_tel field.
func (nc *NurseCreate) SetNurseTel(s string) *NurseCreate {
	nc.mutation.SetNurseTel(s)
	return nc
}

// AddFromnurseIDs adds the fromnurse edge to Rent by ids.
func (nc *NurseCreate) AddFromnurseIDs(ids ...int) *NurseCreate {
	nc.mutation.AddFromnurseIDs(ids...)
	return nc
}

// AddFromnurse adds the fromnurse edges to Rent.
func (nc *NurseCreate) AddFromnurse(r ...*Rent) *NurseCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return nc.AddFromnurseIDs(ids...)
}

// AddNursePrescriptionIDs adds the nurse_prescription edge to Prescription by ids.
func (nc *NurseCreate) AddNursePrescriptionIDs(ids ...int) *NurseCreate {
	nc.mutation.AddNursePrescriptionIDs(ids...)
	return nc
}

// AddNursePrescription adds the nurse_prescription edges to Prescription.
func (nc *NurseCreate) AddNursePrescription(p ...*Prescription) *NurseCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nc.AddNursePrescriptionIDs(ids...)
}

// AddNurseOperativerecordIDs adds the Nurse_Operativerecord edge to Operativerecord by ids.
func (nc *NurseCreate) AddNurseOperativerecordIDs(ids ...int) *NurseCreate {
	nc.mutation.AddNurseOperativerecordIDs(ids...)
	return nc
}

// AddNurseOperativerecord adds the Nurse_Operativerecord edges to Operativerecord.
func (nc *NurseCreate) AddNurseOperativerecord(o ...*Operativerecord) *NurseCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return nc.AddNurseOperativerecordIDs(ids...)
}

// Mutation returns the NurseMutation object of the builder.
func (nc *NurseCreate) Mutation() *NurseMutation {
	return nc.mutation
}

// Save creates the Nurse in the database.
func (nc *NurseCreate) Save(ctx context.Context) (*Nurse, error) {
	if _, ok := nc.mutation.NurseName(); !ok {
		return nil, &ValidationError{Name: "nurse_name", err: errors.New("ent: missing required field \"nurse_name\"")}
	}
	if v, ok := nc.mutation.NurseName(); ok {
		if err := nurse.NurseNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_name", err: fmt.Errorf("ent: validator failed for field \"nurse_name\": %w", err)}
		}
	}
	if _, ok := nc.mutation.NurseEmail(); !ok {
		return nil, &ValidationError{Name: "nurse_email", err: errors.New("ent: missing required field \"nurse_email\"")}
	}
	if v, ok := nc.mutation.NurseEmail(); ok {
		if err := nurse.NurseEmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_email", err: fmt.Errorf("ent: validator failed for field \"nurse_email\": %w", err)}
		}
	}
	if _, ok := nc.mutation.NursePassword(); !ok {
		return nil, &ValidationError{Name: "nurse_password", err: errors.New("ent: missing required field \"nurse_password\"")}
	}
	if v, ok := nc.mutation.NursePassword(); ok {
		if err := nurse.NursePasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_password", err: fmt.Errorf("ent: validator failed for field \"nurse_password\": %w", err)}
		}
	}
	if _, ok := nc.mutation.NurseTel(); !ok {
		return nil, &ValidationError{Name: "nurse_tel", err: errors.New("ent: missing required field \"nurse_tel\"")}
	}
	if v, ok := nc.mutation.NurseTel(); ok {
		if err := nurse.NurseTelValidator(v); err != nil {
			return nil, &ValidationError{Name: "nurse_tel", err: fmt.Errorf("ent: validator failed for field \"nurse_tel\": %w", err)}
		}
	}
	var (
		err  error
		node *Nurse
	)
	if len(nc.hooks) == 0 {
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NurseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nc.mutation = mutation
			node, err = nc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NurseCreate) SaveX(ctx context.Context) *Nurse {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nc *NurseCreate) sqlSave(ctx context.Context) (*Nurse, error) {
	n, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	n.ID = int(id)
	return n, nil
}

func (nc *NurseCreate) createSpec() (*Nurse, *sqlgraph.CreateSpec) {
	var (
		n     = &Nurse{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: nurse.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nurse.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.NurseName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseName,
		})
		n.NurseName = value
	}
	if value, ok := nc.mutation.NurseEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseEmail,
		})
		n.NurseEmail = value
	}
	if value, ok := nc.mutation.NursePassword(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNursePassword,
		})
		n.NursePassword = value
	}
	if value, ok := nc.mutation.NurseTel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nurse.FieldNurseTel,
		})
		n.NurseTel = value
	}
	if nodes := nc.mutation.FromnurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.FromnurseTable,
			Columns: []string{nurse.FromnurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.NursePrescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.NursePrescriptionTable,
			Columns: []string{nurse.NursePrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prescription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.NurseOperativerecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nurse.NurseOperativerecordTable,
			Columns: []string{nurse.NurseOperativerecordColumn},
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
	return n, _spec
}
