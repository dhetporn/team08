// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team08/ent/roomtype"
)

// Roomtype is the model entity for the Roomtype schema.
type Roomtype struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Bathroom holds the value of the "bathroom" field.
	Bathroom int `json:"bathroom,omitempty"`
	// Toilets holds the value of the "toilets" field.
	Toilets int `json:"toilets,omitempty"`
	// Areasize holds the value of the "areasize" field.
	Areasize float64 `json:"areasize,omitempty"`
	// Etc holds the value of the "etc" field.
	Etc string `json:"etc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomtypeQuery when eager-loading is set.
	Edges RoomtypeEdges `json:"edges"`
}

// RoomtypeEdges holds the relations/edges for other nodes in the graph.
type RoomtypeEdges struct {
	// Rooms holds the value of the rooms edge.
	Rooms []*Room
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoomsOrErr returns the Rooms value or an error if the edge
// was not loaded in eager-loading.
func (e RoomtypeEdges) RoomsOrErr() ([]*Room, error) {
	if e.loadedTypes[0] {
		return e.Rooms, nil
	}
	return nil, &NotLoadedError{edge: "rooms"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Roomtype) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},   // id
		&sql.NullString{},  // name
		&sql.NullInt64{},   // bathroom
		&sql.NullInt64{},   // toilets
		&sql.NullFloat64{}, // areasize
		&sql.NullString{},  // etc
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Roomtype fields.
func (r *Roomtype) assignValues(values ...interface{}) error {
	if m, n := len(values), len(roomtype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		r.Name = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field bathroom", values[1])
	} else if value.Valid {
		r.Bathroom = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field toilets", values[2])
	} else if value.Valid {
		r.Toilets = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field areasize", values[3])
	} else if value.Valid {
		r.Areasize = value.Float64
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field etc", values[4])
	} else if value.Valid {
		r.Etc = value.String
	}
	return nil
}

// QueryRooms queries the rooms edge of the Roomtype.
func (r *Roomtype) QueryRooms() *RoomQuery {
	return (&RoomtypeClient{config: r.config}).QueryRooms(r)
}

// Update returns a builder for updating this Roomtype.
// Note that, you need to call Roomtype.Unwrap() before calling this method, if this Roomtype
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Roomtype) Update() *RoomtypeUpdateOne {
	return (&RoomtypeClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Roomtype) Unwrap() *Roomtype {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Roomtype is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Roomtype) String() string {
	var builder strings.Builder
	builder.WriteString("Roomtype(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteString(", bathroom=")
	builder.WriteString(fmt.Sprintf("%v", r.Bathroom))
	builder.WriteString(", toilets=")
	builder.WriteString(fmt.Sprintf("%v", r.Toilets))
	builder.WriteString(", areasize=")
	builder.WriteString(fmt.Sprintf("%v", r.Areasize))
	builder.WriteString(", etc=")
	builder.WriteString(r.Etc)
	builder.WriteByte(')')
	return builder.String()
}

// Roomtypes is a parsable slice of Roomtype.
type Roomtypes []*Roomtype

func (r Roomtypes) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
