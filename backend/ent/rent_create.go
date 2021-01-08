// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dhetporn/team08/ent/nurse"
	"github.com/dhetporn/team08/ent/patient"
	"github.com/dhetporn/team08/ent/rent"
	"github.com/dhetporn/team08/ent/room"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RentCreate is the builder for creating a Rent entity.
type RentCreate struct {
	config
	mutation *RentMutation
	hooks    []Hook
}

// SetAddedTime sets the added_time field.
func (rc *RentCreate) SetAddedTime(t time.Time) *RentCreate {
	rc.mutation.SetAddedTime(t)
	return rc
}

// SetRoomID sets the room edge to Room by id.
func (rc *RentCreate) SetRoomID(id int) *RentCreate {
	rc.mutation.SetRoomID(id)
	return rc
}

// SetNillableRoomID sets the room edge to Room by id if the given value is not nil.
func (rc *RentCreate) SetNillableRoomID(id *int) *RentCreate {
	if id != nil {
		rc = rc.SetRoomID(*id)
	}
	return rc
}

// SetRoom sets the room edge to Room.
func (rc *RentCreate) SetRoom(r *Room) *RentCreate {
	return rc.SetRoomID(r.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (rc *RentCreate) SetPatientID(id int) *RentCreate {
	rc.mutation.SetPatientID(id)
	return rc
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (rc *RentCreate) SetNillablePatientID(id *int) *RentCreate {
	if id != nil {
		rc = rc.SetPatientID(*id)
	}
	return rc
}

// SetPatient sets the patient edge to Patient.
func (rc *RentCreate) SetPatient(p *Patient) *RentCreate {
	return rc.SetPatientID(p.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (rc *RentCreate) SetNurseID(id int) *RentCreate {
	rc.mutation.SetNurseID(id)
	return rc
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (rc *RentCreate) SetNillableNurseID(id *int) *RentCreate {
	if id != nil {
		rc = rc.SetNurseID(*id)
	}
	return rc
}

// SetNurse sets the nurse edge to Nurse.
func (rc *RentCreate) SetNurse(n *Nurse) *RentCreate {
	return rc.SetNurseID(n.ID)
}

// Mutation returns the RentMutation object of the builder.
func (rc *RentCreate) Mutation() *RentMutation {
	return rc.mutation
}

// Save creates the Rent in the database.
func (rc *RentCreate) Save(ctx context.Context) (*Rent, error) {
	if _, ok := rc.mutation.AddedTime(); !ok {
		return nil, &ValidationError{Name: "added_time", err: errors.New("ent: missing required field \"added_time\"")}
	}
	var (
		err  error
		node *Rent
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RentCreate) SaveX(ctx context.Context) *Rent {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RentCreate) sqlSave(ctx context.Context) (*Rent, error) {
	r, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}

func (rc *RentCreate) createSpec() (*Rent, *sqlgraph.CreateSpec) {
	var (
		r     = &Rent{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rent.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.AddedTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rent.FieldAddedTime,
		})
		r.AddedTime = value
	}
	if nodes := rc.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rent.RoomTable,
			Columns: []string{rent.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rent.PatientTable,
			Columns: []string{rent.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rent.NurseTable,
			Columns: []string{rent.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return r, _spec
}