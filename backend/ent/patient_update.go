// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/dhetporn/team08/ent/bloodtype"
	"github.com/dhetporn/team08/ent/coveredperson"
	"github.com/dhetporn/team08/ent/diagnose"
	"github.com/dhetporn/team08/ent/gender"
	"github.com/dhetporn/team08/ent/patient"
	"github.com/dhetporn/team08/ent/predicate"
	"github.com/dhetporn/team08/ent/prefix"
	"github.com/dhetporn/team08/ent/rent"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PatientUpdate is the builder for updating Patient entities.
type PatientUpdate struct {
	config
	hooks      []Hook
	mutation   *PatientMutation
	predicates []predicate.Patient
}

// Where adds a new predicate for the builder.
func (pu *PatientUpdate) Where(ps ...predicate.Patient) *PatientUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetPatientName sets the Patient_name field.
func (pu *PatientUpdate) SetPatientName(s string) *PatientUpdate {
	pu.mutation.SetPatientName(s)
	return pu
}

// SetPatientAge sets the Patient_age field.
func (pu *PatientUpdate) SetPatientAge(i int) *PatientUpdate {
	pu.mutation.ResetPatientAge()
	pu.mutation.SetPatientAge(i)
	return pu
}

// AddPatientAge adds i to Patient_age.
func (pu *PatientUpdate) AddPatientAge(i int) *PatientUpdate {
	pu.mutation.AddPatientAge(i)
	return pu
}

// SetPatientWeight sets the Patient_weight field.
func (pu *PatientUpdate) SetPatientWeight(f float64) *PatientUpdate {
	pu.mutation.ResetPatientWeight()
	pu.mutation.SetPatientWeight(f)
	return pu
}

// AddPatientWeight adds f to Patient_weight.
func (pu *PatientUpdate) AddPatientWeight(f float64) *PatientUpdate {
	pu.mutation.AddPatientWeight(f)
	return pu
}

// SetPatientHeight sets the Patient_height field.
func (pu *PatientUpdate) SetPatientHeight(f float64) *PatientUpdate {
	pu.mutation.ResetPatientHeight()
	pu.mutation.SetPatientHeight(f)
	return pu
}

// AddPatientHeight adds f to Patient_height.
func (pu *PatientUpdate) AddPatientHeight(f float64) *PatientUpdate {
	pu.mutation.AddPatientHeight(f)
	return pu
}

// SetFrompatientID sets the frompatient edge to Rent by id.
func (pu *PatientUpdate) SetFrompatientID(id int) *PatientUpdate {
	pu.mutation.SetFrompatientID(id)
	return pu
}

// SetNillableFrompatientID sets the frompatient edge to Rent by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableFrompatientID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetFrompatientID(*id)
	}
	return pu
}

// SetFrompatient sets the frompatient edge to Rent.
func (pu *PatientUpdate) SetFrompatient(r *Rent) *PatientUpdate {
	return pu.SetFrompatientID(r.ID)
}

// AddPatientCoveredPersonIDs adds the Patient_CoveredPerson edge to CoveredPerson by ids.
func (pu *PatientUpdate) AddPatientCoveredPersonIDs(ids ...int) *PatientUpdate {
	pu.mutation.AddPatientCoveredPersonIDs(ids...)
	return pu
}

// AddPatientCoveredPerson adds the Patient_CoveredPerson edges to CoveredPerson.
func (pu *PatientUpdate) AddPatientCoveredPerson(c ...*CoveredPerson) *PatientUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddPatientCoveredPersonIDs(ids...)
}

// AddPatientDiagnoseIDs adds the patient_diagnose edge to Diagnose by ids.
func (pu *PatientUpdate) AddPatientDiagnoseIDs(ids ...int) *PatientUpdate {
	pu.mutation.AddPatientDiagnoseIDs(ids...)
	return pu
}

// AddPatientDiagnose adds the patient_diagnose edges to Diagnose.
func (pu *PatientUpdate) AddPatientDiagnose(d ...*Diagnose) *PatientUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.AddPatientDiagnoseIDs(ids...)
}

// SetGenderID sets the gender edge to Gender by id.
func (pu *PatientUpdate) SetGenderID(id int) *PatientUpdate {
	pu.mutation.SetGenderID(id)
	return pu
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableGenderID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetGenderID(*id)
	}
	return pu
}

// SetGender sets the gender edge to Gender.
func (pu *PatientUpdate) SetGender(g *Gender) *PatientUpdate {
	return pu.SetGenderID(g.ID)
}

// SetPrefixID sets the prefix edge to Prefix by id.
func (pu *PatientUpdate) SetPrefixID(id int) *PatientUpdate {
	pu.mutation.SetPrefixID(id)
	return pu
}

// SetNillablePrefixID sets the prefix edge to Prefix by id if the given value is not nil.
func (pu *PatientUpdate) SetNillablePrefixID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetPrefixID(*id)
	}
	return pu
}

// SetPrefix sets the prefix edge to Prefix.
func (pu *PatientUpdate) SetPrefix(p *Prefix) *PatientUpdate {
	return pu.SetPrefixID(p.ID)
}

// SetBloodtypeID sets the bloodtype edge to Bloodtype by id.
func (pu *PatientUpdate) SetBloodtypeID(id int) *PatientUpdate {
	pu.mutation.SetBloodtypeID(id)
	return pu
}

// SetNillableBloodtypeID sets the bloodtype edge to Bloodtype by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableBloodtypeID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetBloodtypeID(*id)
	}
	return pu
}

// SetBloodtype sets the bloodtype edge to Bloodtype.
func (pu *PatientUpdate) SetBloodtype(b *Bloodtype) *PatientUpdate {
	return pu.SetBloodtypeID(b.ID)
}

// Mutation returns the PatientMutation object of the builder.
func (pu *PatientUpdate) Mutation() *PatientMutation {
	return pu.mutation
}

// ClearFrompatient clears the frompatient edge to Rent.
func (pu *PatientUpdate) ClearFrompatient() *PatientUpdate {
	pu.mutation.ClearFrompatient()
	return pu
}

// RemovePatientCoveredPersonIDs removes the Patient_CoveredPerson edge to CoveredPerson by ids.
func (pu *PatientUpdate) RemovePatientCoveredPersonIDs(ids ...int) *PatientUpdate {
	pu.mutation.RemovePatientCoveredPersonIDs(ids...)
	return pu
}

// RemovePatientCoveredPerson removes Patient_CoveredPerson edges to CoveredPerson.
func (pu *PatientUpdate) RemovePatientCoveredPerson(c ...*CoveredPerson) *PatientUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemovePatientCoveredPersonIDs(ids...)
}

// RemovePatientDiagnoseIDs removes the patient_diagnose edge to Diagnose by ids.
func (pu *PatientUpdate) RemovePatientDiagnoseIDs(ids ...int) *PatientUpdate {
	pu.mutation.RemovePatientDiagnoseIDs(ids...)
	return pu
}

// RemovePatientDiagnose removes patient_diagnose edges to Diagnose.
func (pu *PatientUpdate) RemovePatientDiagnose(d ...*Diagnose) *PatientUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.RemovePatientDiagnoseIDs(ids...)
}

// ClearGender clears the gender edge to Gender.
func (pu *PatientUpdate) ClearGender() *PatientUpdate {
	pu.mutation.ClearGender()
	return pu
}

// ClearPrefix clears the prefix edge to Prefix.
func (pu *PatientUpdate) ClearPrefix() *PatientUpdate {
	pu.mutation.ClearPrefix()
	return pu
}

// ClearBloodtype clears the bloodtype edge to Bloodtype.
func (pu *PatientUpdate) ClearBloodtype() *PatientUpdate {
	pu.mutation.ClearBloodtype()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PatientUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.PatientName(); ok {
		if err := patient.PatientNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Patient_name", err: fmt.Errorf("ent: validator failed for field \"Patient_name\": %w", err)}
		}
	}
	if v, ok := pu.mutation.PatientAge(); ok {
		if err := patient.PatientAgeValidator(v); err != nil {
			return 0, &ValidationError{Name: "Patient_age", err: fmt.Errorf("ent: validator failed for field \"Patient_age\": %w", err)}
		}
	}
	if v, ok := pu.mutation.PatientWeight(); ok {
		if err := patient.PatientWeightValidator(v); err != nil {
			return 0, &ValidationError{Name: "Patient_weight", err: fmt.Errorf("ent: validator failed for field \"Patient_weight\": %w", err)}
		}
	}
	if v, ok := pu.mutation.PatientHeight(); ok {
		if err := patient.PatientHeightValidator(v); err != nil {
			return 0, &ValidationError{Name: "Patient_height", err: fmt.Errorf("ent: validator failed for field \"Patient_height\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PatientUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PatientUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PatientUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PatientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.PatientName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientName,
		})
	}
	if value, ok := pu.mutation.PatientAge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientAge,
		})
	}
	if value, ok := pu.mutation.AddedPatientAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientAge,
		})
	}
	if value, ok := pu.mutation.PatientWeight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: patient.FieldPatientWeight,
		})
	}
	if value, ok := pu.mutation.AddedPatientWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: patient.FieldPatientWeight,
		})
	}
	if value, ok := pu.mutation.PatientHeight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: patient.FieldPatientHeight,
		})
	}
	if value, ok := pu.mutation.AddedPatientHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: patient.FieldPatientHeight,
		})
	}
	if pu.mutation.FrompatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   patient.FrompatientTable,
			Columns: []string{patient.FrompatientColumn},
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
	if nodes := pu.mutation.FrompatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   patient.FrompatientTable,
			Columns: []string{patient.FrompatientColumn},
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
	if nodes := pu.mutation.RemovedPatientCoveredPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientCoveredPersonTable,
			Columns: []string{patient.PatientCoveredPersonColumn},
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
	if nodes := pu.mutation.PatientCoveredPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientCoveredPersonTable,
			Columns: []string{patient.PatientCoveredPersonColumn},
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
	if nodes := pu.mutation.RemovedPatientDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientDiagnoseTable,
			Columns: []string{patient.PatientDiagnoseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PatientDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientDiagnoseTable,
			Columns: []string{patient.PatientDiagnoseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.GenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PrefixCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.PrefixTable,
			Columns: []string{patient.PrefixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prefix.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PrefixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.PrefixTable,
			Columns: []string{patient.PrefixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prefix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.BloodtypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.BloodtypeTable,
			Columns: []string{patient.BloodtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bloodtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.BloodtypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.BloodtypeTable,
			Columns: []string{patient.BloodtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bloodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PatientUpdateOne is the builder for updating a single Patient entity.
type PatientUpdateOne struct {
	config
	hooks    []Hook
	mutation *PatientMutation
}

// SetPatientName sets the Patient_name field.
func (puo *PatientUpdateOne) SetPatientName(s string) *PatientUpdateOne {
	puo.mutation.SetPatientName(s)
	return puo
}

// SetPatientAge sets the Patient_age field.
func (puo *PatientUpdateOne) SetPatientAge(i int) *PatientUpdateOne {
	puo.mutation.ResetPatientAge()
	puo.mutation.SetPatientAge(i)
	return puo
}

// AddPatientAge adds i to Patient_age.
func (puo *PatientUpdateOne) AddPatientAge(i int) *PatientUpdateOne {
	puo.mutation.AddPatientAge(i)
	return puo
}

// SetPatientWeight sets the Patient_weight field.
func (puo *PatientUpdateOne) SetPatientWeight(f float64) *PatientUpdateOne {
	puo.mutation.ResetPatientWeight()
	puo.mutation.SetPatientWeight(f)
	return puo
}

// AddPatientWeight adds f to Patient_weight.
func (puo *PatientUpdateOne) AddPatientWeight(f float64) *PatientUpdateOne {
	puo.mutation.AddPatientWeight(f)
	return puo
}

// SetPatientHeight sets the Patient_height field.
func (puo *PatientUpdateOne) SetPatientHeight(f float64) *PatientUpdateOne {
	puo.mutation.ResetPatientHeight()
	puo.mutation.SetPatientHeight(f)
	return puo
}

// AddPatientHeight adds f to Patient_height.
func (puo *PatientUpdateOne) AddPatientHeight(f float64) *PatientUpdateOne {
	puo.mutation.AddPatientHeight(f)
	return puo
}

// SetFrompatientID sets the frompatient edge to Rent by id.
func (puo *PatientUpdateOne) SetFrompatientID(id int) *PatientUpdateOne {
	puo.mutation.SetFrompatientID(id)
	return puo
}

// SetNillableFrompatientID sets the frompatient edge to Rent by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableFrompatientID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetFrompatientID(*id)
	}
	return puo
}

// SetFrompatient sets the frompatient edge to Rent.
func (puo *PatientUpdateOne) SetFrompatient(r *Rent) *PatientUpdateOne {
	return puo.SetFrompatientID(r.ID)
}

// AddPatientCoveredPersonIDs adds the Patient_CoveredPerson edge to CoveredPerson by ids.
func (puo *PatientUpdateOne) AddPatientCoveredPersonIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.AddPatientCoveredPersonIDs(ids...)
	return puo
}

// AddPatientCoveredPerson adds the Patient_CoveredPerson edges to CoveredPerson.
func (puo *PatientUpdateOne) AddPatientCoveredPerson(c ...*CoveredPerson) *PatientUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddPatientCoveredPersonIDs(ids...)
}

// AddPatientDiagnoseIDs adds the patient_diagnose edge to Diagnose by ids.
func (puo *PatientUpdateOne) AddPatientDiagnoseIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.AddPatientDiagnoseIDs(ids...)
	return puo
}

// AddPatientDiagnose adds the patient_diagnose edges to Diagnose.
func (puo *PatientUpdateOne) AddPatientDiagnose(d ...*Diagnose) *PatientUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.AddPatientDiagnoseIDs(ids...)
}

// SetGenderID sets the gender edge to Gender by id.
func (puo *PatientUpdateOne) SetGenderID(id int) *PatientUpdateOne {
	puo.mutation.SetGenderID(id)
	return puo
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableGenderID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetGenderID(*id)
	}
	return puo
}

// SetGender sets the gender edge to Gender.
func (puo *PatientUpdateOne) SetGender(g *Gender) *PatientUpdateOne {
	return puo.SetGenderID(g.ID)
}

// SetPrefixID sets the prefix edge to Prefix by id.
func (puo *PatientUpdateOne) SetPrefixID(id int) *PatientUpdateOne {
	puo.mutation.SetPrefixID(id)
	return puo
}

// SetNillablePrefixID sets the prefix edge to Prefix by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillablePrefixID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetPrefixID(*id)
	}
	return puo
}

// SetPrefix sets the prefix edge to Prefix.
func (puo *PatientUpdateOne) SetPrefix(p *Prefix) *PatientUpdateOne {
	return puo.SetPrefixID(p.ID)
}

// SetBloodtypeID sets the bloodtype edge to Bloodtype by id.
func (puo *PatientUpdateOne) SetBloodtypeID(id int) *PatientUpdateOne {
	puo.mutation.SetBloodtypeID(id)
	return puo
}

// SetNillableBloodtypeID sets the bloodtype edge to Bloodtype by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableBloodtypeID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetBloodtypeID(*id)
	}
	return puo
}

// SetBloodtype sets the bloodtype edge to Bloodtype.
func (puo *PatientUpdateOne) SetBloodtype(b *Bloodtype) *PatientUpdateOne {
	return puo.SetBloodtypeID(b.ID)
}

// Mutation returns the PatientMutation object of the builder.
func (puo *PatientUpdateOne) Mutation() *PatientMutation {
	return puo.mutation
}

// ClearFrompatient clears the frompatient edge to Rent.
func (puo *PatientUpdateOne) ClearFrompatient() *PatientUpdateOne {
	puo.mutation.ClearFrompatient()
	return puo
}

// RemovePatientCoveredPersonIDs removes the Patient_CoveredPerson edge to CoveredPerson by ids.
func (puo *PatientUpdateOne) RemovePatientCoveredPersonIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.RemovePatientCoveredPersonIDs(ids...)
	return puo
}

// RemovePatientCoveredPerson removes Patient_CoveredPerson edges to CoveredPerson.
func (puo *PatientUpdateOne) RemovePatientCoveredPerson(c ...*CoveredPerson) *PatientUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemovePatientCoveredPersonIDs(ids...)
}

// RemovePatientDiagnoseIDs removes the patient_diagnose edge to Diagnose by ids.
func (puo *PatientUpdateOne) RemovePatientDiagnoseIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.RemovePatientDiagnoseIDs(ids...)
	return puo
}

// RemovePatientDiagnose removes patient_diagnose edges to Diagnose.
func (puo *PatientUpdateOne) RemovePatientDiagnose(d ...*Diagnose) *PatientUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.RemovePatientDiagnoseIDs(ids...)
}

// ClearGender clears the gender edge to Gender.
func (puo *PatientUpdateOne) ClearGender() *PatientUpdateOne {
	puo.mutation.ClearGender()
	return puo
}

// ClearPrefix clears the prefix edge to Prefix.
func (puo *PatientUpdateOne) ClearPrefix() *PatientUpdateOne {
	puo.mutation.ClearPrefix()
	return puo
}

// ClearBloodtype clears the bloodtype edge to Bloodtype.
func (puo *PatientUpdateOne) ClearBloodtype() *PatientUpdateOne {
	puo.mutation.ClearBloodtype()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PatientUpdateOne) Save(ctx context.Context) (*Patient, error) {
	if v, ok := puo.mutation.PatientName(); ok {
		if err := patient.PatientNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Patient_name", err: fmt.Errorf("ent: validator failed for field \"Patient_name\": %w", err)}
		}
	}
	if v, ok := puo.mutation.PatientAge(); ok {
		if err := patient.PatientAgeValidator(v); err != nil {
			return nil, &ValidationError{Name: "Patient_age", err: fmt.Errorf("ent: validator failed for field \"Patient_age\": %w", err)}
		}
	}
	if v, ok := puo.mutation.PatientWeight(); ok {
		if err := patient.PatientWeightValidator(v); err != nil {
			return nil, &ValidationError{Name: "Patient_weight", err: fmt.Errorf("ent: validator failed for field \"Patient_weight\": %w", err)}
		}
	}
	if v, ok := puo.mutation.PatientHeight(); ok {
		if err := patient.PatientHeightValidator(v); err != nil {
			return nil, &ValidationError{Name: "Patient_height", err: fmt.Errorf("ent: validator failed for field \"Patient_height\": %w", err)}
		}
	}

	var (
		err  error
		node *Patient
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PatientUpdateOne) SaveX(ctx context.Context) *Patient {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PatientUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PatientUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PatientUpdateOne) sqlSave(ctx context.Context) (pa *Patient, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Patient.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.PatientName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientName,
		})
	}
	if value, ok := puo.mutation.PatientAge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientAge,
		})
	}
	if value, ok := puo.mutation.AddedPatientAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientAge,
		})
	}
	if value, ok := puo.mutation.PatientWeight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: patient.FieldPatientWeight,
		})
	}
	if value, ok := puo.mutation.AddedPatientWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: patient.FieldPatientWeight,
		})
	}
	if value, ok := puo.mutation.PatientHeight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: patient.FieldPatientHeight,
		})
	}
	if value, ok := puo.mutation.AddedPatientHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: patient.FieldPatientHeight,
		})
	}
	if puo.mutation.FrompatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   patient.FrompatientTable,
			Columns: []string{patient.FrompatientColumn},
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
	if nodes := puo.mutation.FrompatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   patient.FrompatientTable,
			Columns: []string{patient.FrompatientColumn},
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
	if nodes := puo.mutation.RemovedPatientCoveredPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientCoveredPersonTable,
			Columns: []string{patient.PatientCoveredPersonColumn},
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
	if nodes := puo.mutation.PatientCoveredPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientCoveredPersonTable,
			Columns: []string{patient.PatientCoveredPersonColumn},
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
	if nodes := puo.mutation.RemovedPatientDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientDiagnoseTable,
			Columns: []string{patient.PatientDiagnoseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PatientDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientDiagnoseTable,
			Columns: []string{patient.PatientDiagnoseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.GenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PrefixCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.PrefixTable,
			Columns: []string{patient.PrefixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prefix.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PrefixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.PrefixTable,
			Columns: []string{patient.PrefixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prefix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.BloodtypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.BloodtypeTable,
			Columns: []string{patient.BloodtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bloodtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.BloodtypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.BloodtypeTable,
			Columns: []string{patient.BloodtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bloodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pa = &Patient{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}
