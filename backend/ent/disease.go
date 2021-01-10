// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team08/ent/disease"
)

// Disease is the model entity for the Disease schema.
type Disease struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DiseaseName holds the value of the "Disease_Name" field.
	DiseaseName string `json:"Disease_Name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DiseaseQuery when eager-loading is set.
	Edges DiseaseEdges `json:"edges"`
}

// DiseaseEdges holds the relations/edges for other nodes in the graph.
type DiseaseEdges struct {
	// DiseaseDiagnose holds the value of the disease_diagnose edge.
	DiseaseDiagnose []*Diagnose
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DiseaseDiagnoseOrErr returns the DiseaseDiagnose value or an error if the edge
// was not loaded in eager-loading.
func (e DiseaseEdges) DiseaseDiagnoseOrErr() ([]*Diagnose, error) {
	if e.loadedTypes[0] {
		return e.DiseaseDiagnose, nil
	}
	return nil, &NotLoadedError{edge: "disease_diagnose"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Disease) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Disease_Name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Disease fields.
func (d *Disease) assignValues(values ...interface{}) error {
	if m, n := len(values), len(disease.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Disease_Name", values[0])
	} else if value.Valid {
		d.DiseaseName = value.String
	}
	return nil
}

// QueryDiseaseDiagnose queries the disease_diagnose edge of the Disease.
func (d *Disease) QueryDiseaseDiagnose() *DiagnoseQuery {
	return (&DiseaseClient{config: d.config}).QueryDiseaseDiagnose(d)
}

// Update returns a builder for updating this Disease.
// Note that, you need to call Disease.Unwrap() before calling this method, if this Disease
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Disease) Update() *DiseaseUpdateOne {
	return (&DiseaseClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Disease) Unwrap() *Disease {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Disease is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Disease) String() string {
	var builder strings.Builder
	builder.WriteString("Disease(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", Disease_Name=")
	builder.WriteString(d.DiseaseName)
	builder.WriteByte(')')
	return builder.String()
}

// Diseases is a parsable slice of Disease.
type Diseases []*Disease

func (d Diseases) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
