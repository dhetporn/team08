// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/dhetporn/team08/ent/medical"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Medical is the model entity for the Medical schema.
type Medical struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MedicalName holds the value of the "Medical_Name" field.
	MedicalName string `json:"Medical_Name,omitempty"`
	// MedicalEmail holds the value of the "Medical_Email" field.
	MedicalEmail string `json:"Medical_Email,omitempty"`
	// MedicalPassword holds the value of the "Medical_Password" field.
	MedicalPassword string `json:"Medical_Password,omitempty"`
	// MedicalTel holds the value of the "Medical_Tel" field.
	MedicalTel string `json:"Medical_Tel,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Medical) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Medical_Name
		&sql.NullString{}, // Medical_Email
		&sql.NullString{}, // Medical_Password
		&sql.NullString{}, // Medical_Tel
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Medical fields.
func (m *Medical) assignValues(values ...interface{}) error {
	if m, n := len(values), len(medical.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Medical_Name", values[0])
	} else if value.Valid {
		m.MedicalName = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Medical_Email", values[1])
	} else if value.Valid {
		m.MedicalEmail = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Medical_Password", values[2])
	} else if value.Valid {
		m.MedicalPassword = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Medical_Tel", values[3])
	} else if value.Valid {
		m.MedicalTel = value.String
	}
	return nil
}

// Update returns a builder for updating this Medical.
// Note that, you need to call Medical.Unwrap() before calling this method, if this Medical
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Medical) Update() *MedicalUpdateOne {
	return (&MedicalClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Medical) Unwrap() *Medical {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Medical is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Medical) String() string {
	var builder strings.Builder
	builder.WriteString("Medical(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", Medical_Name=")
	builder.WriteString(m.MedicalName)
	builder.WriteString(", Medical_Email=")
	builder.WriteString(m.MedicalEmail)
	builder.WriteString(", Medical_Password=")
	builder.WriteString(m.MedicalPassword)
	builder.WriteString(", Medical_Tel=")
	builder.WriteString(m.MedicalTel)
	builder.WriteByte(')')
	return builder.String()
}

// Medicals is a parsable slice of Medical.
type Medicals []*Medical

func (m Medicals) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
