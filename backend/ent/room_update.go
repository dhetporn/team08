// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/dhetporn/team08/ent/predicate"
	"github.com/dhetporn/team08/ent/rent"
	"github.com/dhetporn/team08/ent/room"
	"github.com/dhetporn/team08/ent/roomtype"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RoomUpdate is the builder for updating Room entities.
type RoomUpdate struct {
	config
	hooks      []Hook
	mutation   *RoomMutation
	predicates []predicate.Room
}

// Where adds a new predicate for the builder.
func (ru *RoomUpdate) Where(ps ...predicate.Room) *RoomUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetName sets the name field.
func (ru *RoomUpdate) SetName(s string) *RoomUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetBuilding sets the building field.
func (ru *RoomUpdate) SetBuilding(i int) *RoomUpdate {
	ru.mutation.ResetBuilding()
	ru.mutation.SetBuilding(i)
	return ru
}

// AddBuilding adds i to building.
func (ru *RoomUpdate) AddBuilding(i int) *RoomUpdate {
	ru.mutation.AddBuilding(i)
	return ru
}

// SetFloor sets the floor field.
func (ru *RoomUpdate) SetFloor(i int) *RoomUpdate {
	ru.mutation.ResetFloor()
	ru.mutation.SetFloor(i)
	return ru
}

// AddFloor adds i to floor.
func (ru *RoomUpdate) AddFloor(i int) *RoomUpdate {
	ru.mutation.AddFloor(i)
	return ru
}

// SetRentsID sets the rents edge to Rent by id.
func (ru *RoomUpdate) SetRentsID(id int) *RoomUpdate {
	ru.mutation.SetRentsID(id)
	return ru
}

// SetNillableRentsID sets the rents edge to Rent by id if the given value is not nil.
func (ru *RoomUpdate) SetNillableRentsID(id *int) *RoomUpdate {
	if id != nil {
		ru = ru.SetRentsID(*id)
	}
	return ru
}

// SetRents sets the rents edge to Rent.
func (ru *RoomUpdate) SetRents(r *Rent) *RoomUpdate {
	return ru.SetRentsID(r.ID)
}

// SetRoomtypeID sets the roomtype edge to Roomtype by id.
func (ru *RoomUpdate) SetRoomtypeID(id int) *RoomUpdate {
	ru.mutation.SetRoomtypeID(id)
	return ru
}

// SetNillableRoomtypeID sets the roomtype edge to Roomtype by id if the given value is not nil.
func (ru *RoomUpdate) SetNillableRoomtypeID(id *int) *RoomUpdate {
	if id != nil {
		ru = ru.SetRoomtypeID(*id)
	}
	return ru
}

// SetRoomtype sets the roomtype edge to Roomtype.
func (ru *RoomUpdate) SetRoomtype(r *Roomtype) *RoomUpdate {
	return ru.SetRoomtypeID(r.ID)
}

// Mutation returns the RoomMutation object of the builder.
func (ru *RoomUpdate) Mutation() *RoomMutation {
	return ru.mutation
}

// ClearRents clears the rents edge to Rent.
func (ru *RoomUpdate) ClearRents() *RoomUpdate {
	ru.mutation.ClearRents()
	return ru
}

// ClearRoomtype clears the roomtype edge to Roomtype.
func (ru *RoomUpdate) ClearRoomtype() *RoomUpdate {
	ru.mutation.ClearRoomtype()
	return ru
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *RoomUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ru.mutation.Name(); ok {
		if err := room.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := ru.mutation.Building(); ok {
		if err := room.BuildingValidator(v); err != nil {
			return 0, &ValidationError{Name: "building", err: fmt.Errorf("ent: validator failed for field \"building\": %w", err)}
		}
	}
	if v, ok := ru.mutation.Floor(); ok {
		if err := room.FloorValidator(v); err != nil {
			return 0, &ValidationError{Name: "floor", err: fmt.Errorf("ent: validator failed for field \"floor\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoomUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoomUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoomUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   room.Table,
			Columns: room.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: room.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldName,
		})
	}
	if value, ok := ru.mutation.Building(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldBuilding,
		})
	}
	if value, ok := ru.mutation.AddedBuilding(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldBuilding,
		})
	}
	if value, ok := ru.mutation.Floor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldFloor,
		})
	}
	if value, ok := ru.mutation.AddedFloor(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldFloor,
		})
	}
	if ru.mutation.RentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   room.RentsTable,
			Columns: []string{room.RentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rent.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   room.RentsTable,
			Columns: []string{room.RentsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.RoomtypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.RoomtypeTable,
			Columns: []string{room.RoomtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RoomtypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.RoomtypeTable,
			Columns: []string{room.RoomtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RoomUpdateOne is the builder for updating a single Room entity.
type RoomUpdateOne struct {
	config
	hooks    []Hook
	mutation *RoomMutation
}

// SetName sets the name field.
func (ruo *RoomUpdateOne) SetName(s string) *RoomUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetBuilding sets the building field.
func (ruo *RoomUpdateOne) SetBuilding(i int) *RoomUpdateOne {
	ruo.mutation.ResetBuilding()
	ruo.mutation.SetBuilding(i)
	return ruo
}

// AddBuilding adds i to building.
func (ruo *RoomUpdateOne) AddBuilding(i int) *RoomUpdateOne {
	ruo.mutation.AddBuilding(i)
	return ruo
}

// SetFloor sets the floor field.
func (ruo *RoomUpdateOne) SetFloor(i int) *RoomUpdateOne {
	ruo.mutation.ResetFloor()
	ruo.mutation.SetFloor(i)
	return ruo
}

// AddFloor adds i to floor.
func (ruo *RoomUpdateOne) AddFloor(i int) *RoomUpdateOne {
	ruo.mutation.AddFloor(i)
	return ruo
}

// SetRentsID sets the rents edge to Rent by id.
func (ruo *RoomUpdateOne) SetRentsID(id int) *RoomUpdateOne {
	ruo.mutation.SetRentsID(id)
	return ruo
}

// SetNillableRentsID sets the rents edge to Rent by id if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableRentsID(id *int) *RoomUpdateOne {
	if id != nil {
		ruo = ruo.SetRentsID(*id)
	}
	return ruo
}

// SetRents sets the rents edge to Rent.
func (ruo *RoomUpdateOne) SetRents(r *Rent) *RoomUpdateOne {
	return ruo.SetRentsID(r.ID)
}

// SetRoomtypeID sets the roomtype edge to Roomtype by id.
func (ruo *RoomUpdateOne) SetRoomtypeID(id int) *RoomUpdateOne {
	ruo.mutation.SetRoomtypeID(id)
	return ruo
}

// SetNillableRoomtypeID sets the roomtype edge to Roomtype by id if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableRoomtypeID(id *int) *RoomUpdateOne {
	if id != nil {
		ruo = ruo.SetRoomtypeID(*id)
	}
	return ruo
}

// SetRoomtype sets the roomtype edge to Roomtype.
func (ruo *RoomUpdateOne) SetRoomtype(r *Roomtype) *RoomUpdateOne {
	return ruo.SetRoomtypeID(r.ID)
}

// Mutation returns the RoomMutation object of the builder.
func (ruo *RoomUpdateOne) Mutation() *RoomMutation {
	return ruo.mutation
}

// ClearRents clears the rents edge to Rent.
func (ruo *RoomUpdateOne) ClearRents() *RoomUpdateOne {
	ruo.mutation.ClearRents()
	return ruo
}

// ClearRoomtype clears the roomtype edge to Roomtype.
func (ruo *RoomUpdateOne) ClearRoomtype() *RoomUpdateOne {
	ruo.mutation.ClearRoomtype()
	return ruo
}

// Save executes the query and returns the updated entity.
func (ruo *RoomUpdateOne) Save(ctx context.Context) (*Room, error) {
	if v, ok := ruo.mutation.Name(); ok {
		if err := room.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.Building(); ok {
		if err := room.BuildingValidator(v); err != nil {
			return nil, &ValidationError{Name: "building", err: fmt.Errorf("ent: validator failed for field \"building\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.Floor(); ok {
		if err := room.FloorValidator(v); err != nil {
			return nil, &ValidationError{Name: "floor", err: fmt.Errorf("ent: validator failed for field \"floor\": %w", err)}
		}
	}

	var (
		err  error
		node *Room
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoomUpdateOne) SaveX(ctx context.Context) *Room {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *RoomUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoomUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RoomUpdateOne) sqlSave(ctx context.Context) (r *Room, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   room.Table,
			Columns: room.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: room.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Room.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldName,
		})
	}
	if value, ok := ruo.mutation.Building(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldBuilding,
		})
	}
	if value, ok := ruo.mutation.AddedBuilding(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldBuilding,
		})
	}
	if value, ok := ruo.mutation.Floor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldFloor,
		})
	}
	if value, ok := ruo.mutation.AddedFloor(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldFloor,
		})
	}
	if ruo.mutation.RentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   room.RentsTable,
			Columns: []string{room.RentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rent.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   room.RentsTable,
			Columns: []string{room.RentsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.RoomtypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.RoomtypeTable,
			Columns: []string{room.RoomtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RoomtypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.RoomtypeTable,
			Columns: []string{room.RoomtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Room{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}
