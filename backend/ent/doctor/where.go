// Code generated by entc, DO NOT EDIT.

package doctor

import (
	"github.com/dhetporn/team08/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DoctorName applies equality check predicate on the "Doctor_Name" field. It's identical to DoctorNameEQ.
func DoctorName(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorName), v))
	})
}

// DoctorPassword applies equality check predicate on the "Doctor_Password" field. It's identical to DoctorPasswordEQ.
func DoctorPassword(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorPassword), v))
	})
}

// DoctorEmail applies equality check predicate on the "Doctor_Email" field. It's identical to DoctorEmailEQ.
func DoctorEmail(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorEmail), v))
	})
}

// DoctorTel applies equality check predicate on the "Doctor_tel" field. It's identical to DoctorTelEQ.
func DoctorTel(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorTel), v))
	})
}

// DoctorNameEQ applies the EQ predicate on the "Doctor_Name" field.
func DoctorNameEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorName), v))
	})
}

// DoctorNameNEQ applies the NEQ predicate on the "Doctor_Name" field.
func DoctorNameNEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDoctorName), v))
	})
}

// DoctorNameIn applies the In predicate on the "Doctor_Name" field.
func DoctorNameIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDoctorName), v...))
	})
}

// DoctorNameNotIn applies the NotIn predicate on the "Doctor_Name" field.
func DoctorNameNotIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDoctorName), v...))
	})
}

// DoctorNameGT applies the GT predicate on the "Doctor_Name" field.
func DoctorNameGT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDoctorName), v))
	})
}

// DoctorNameGTE applies the GTE predicate on the "Doctor_Name" field.
func DoctorNameGTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDoctorName), v))
	})
}

// DoctorNameLT applies the LT predicate on the "Doctor_Name" field.
func DoctorNameLT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDoctorName), v))
	})
}

// DoctorNameLTE applies the LTE predicate on the "Doctor_Name" field.
func DoctorNameLTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDoctorName), v))
	})
}

// DoctorNameContains applies the Contains predicate on the "Doctor_Name" field.
func DoctorNameContains(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDoctorName), v))
	})
}

// DoctorNameHasPrefix applies the HasPrefix predicate on the "Doctor_Name" field.
func DoctorNameHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDoctorName), v))
	})
}

// DoctorNameHasSuffix applies the HasSuffix predicate on the "Doctor_Name" field.
func DoctorNameHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDoctorName), v))
	})
}

// DoctorNameEqualFold applies the EqualFold predicate on the "Doctor_Name" field.
func DoctorNameEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDoctorName), v))
	})
}

// DoctorNameContainsFold applies the ContainsFold predicate on the "Doctor_Name" field.
func DoctorNameContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDoctorName), v))
	})
}

// DoctorPasswordEQ applies the EQ predicate on the "Doctor_Password" field.
func DoctorPasswordEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorPassword), v))
	})
}

// DoctorPasswordNEQ applies the NEQ predicate on the "Doctor_Password" field.
func DoctorPasswordNEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDoctorPassword), v))
	})
}

// DoctorPasswordIn applies the In predicate on the "Doctor_Password" field.
func DoctorPasswordIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDoctorPassword), v...))
	})
}

// DoctorPasswordNotIn applies the NotIn predicate on the "Doctor_Password" field.
func DoctorPasswordNotIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDoctorPassword), v...))
	})
}

// DoctorPasswordGT applies the GT predicate on the "Doctor_Password" field.
func DoctorPasswordGT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDoctorPassword), v))
	})
}

// DoctorPasswordGTE applies the GTE predicate on the "Doctor_Password" field.
func DoctorPasswordGTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDoctorPassword), v))
	})
}

// DoctorPasswordLT applies the LT predicate on the "Doctor_Password" field.
func DoctorPasswordLT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDoctorPassword), v))
	})
}

// DoctorPasswordLTE applies the LTE predicate on the "Doctor_Password" field.
func DoctorPasswordLTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDoctorPassword), v))
	})
}

// DoctorPasswordContains applies the Contains predicate on the "Doctor_Password" field.
func DoctorPasswordContains(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDoctorPassword), v))
	})
}

// DoctorPasswordHasPrefix applies the HasPrefix predicate on the "Doctor_Password" field.
func DoctorPasswordHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDoctorPassword), v))
	})
}

// DoctorPasswordHasSuffix applies the HasSuffix predicate on the "Doctor_Password" field.
func DoctorPasswordHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDoctorPassword), v))
	})
}

// DoctorPasswordEqualFold applies the EqualFold predicate on the "Doctor_Password" field.
func DoctorPasswordEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDoctorPassword), v))
	})
}

// DoctorPasswordContainsFold applies the ContainsFold predicate on the "Doctor_Password" field.
func DoctorPasswordContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDoctorPassword), v))
	})
}

// DoctorEmailEQ applies the EQ predicate on the "Doctor_Email" field.
func DoctorEmailEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailNEQ applies the NEQ predicate on the "Doctor_Email" field.
func DoctorEmailNEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailIn applies the In predicate on the "Doctor_Email" field.
func DoctorEmailIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDoctorEmail), v...))
	})
}

// DoctorEmailNotIn applies the NotIn predicate on the "Doctor_Email" field.
func DoctorEmailNotIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDoctorEmail), v...))
	})
}

// DoctorEmailGT applies the GT predicate on the "Doctor_Email" field.
func DoctorEmailGT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailGTE applies the GTE predicate on the "Doctor_Email" field.
func DoctorEmailGTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailLT applies the LT predicate on the "Doctor_Email" field.
func DoctorEmailLT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailLTE applies the LTE predicate on the "Doctor_Email" field.
func DoctorEmailLTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailContains applies the Contains predicate on the "Doctor_Email" field.
func DoctorEmailContains(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailHasPrefix applies the HasPrefix predicate on the "Doctor_Email" field.
func DoctorEmailHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailHasSuffix applies the HasSuffix predicate on the "Doctor_Email" field.
func DoctorEmailHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailEqualFold applies the EqualFold predicate on the "Doctor_Email" field.
func DoctorEmailEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailContainsFold applies the ContainsFold predicate on the "Doctor_Email" field.
func DoctorEmailContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDoctorEmail), v))
	})
}

// DoctorTelEQ applies the EQ predicate on the "Doctor_tel" field.
func DoctorTelEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorTel), v))
	})
}

// DoctorTelNEQ applies the NEQ predicate on the "Doctor_tel" field.
func DoctorTelNEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDoctorTel), v))
	})
}

// DoctorTelIn applies the In predicate on the "Doctor_tel" field.
func DoctorTelIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDoctorTel), v...))
	})
}

// DoctorTelNotIn applies the NotIn predicate on the "Doctor_tel" field.
func DoctorTelNotIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDoctorTel), v...))
	})
}

// DoctorTelGT applies the GT predicate on the "Doctor_tel" field.
func DoctorTelGT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDoctorTel), v))
	})
}

// DoctorTelGTE applies the GTE predicate on the "Doctor_tel" field.
func DoctorTelGTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDoctorTel), v))
	})
}

// DoctorTelLT applies the LT predicate on the "Doctor_tel" field.
func DoctorTelLT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDoctorTel), v))
	})
}

// DoctorTelLTE applies the LTE predicate on the "Doctor_tel" field.
func DoctorTelLTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDoctorTel), v))
	})
}

// DoctorTelContains applies the Contains predicate on the "Doctor_tel" field.
func DoctorTelContains(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDoctorTel), v))
	})
}

// DoctorTelHasPrefix applies the HasPrefix predicate on the "Doctor_tel" field.
func DoctorTelHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDoctorTel), v))
	})
}

// DoctorTelHasSuffix applies the HasSuffix predicate on the "Doctor_tel" field.
func DoctorTelHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDoctorTel), v))
	})
}

// DoctorTelEqualFold applies the EqualFold predicate on the "Doctor_tel" field.
func DoctorTelEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDoctorTel), v))
	})
}

// DoctorTelContainsFold applies the ContainsFold predicate on the "Doctor_tel" field.
func DoctorTelContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDoctorTel), v))
	})
}

// HasDoctorDiagnose applies the HasEdge predicate on the "doctor_diagnose" edge.
func HasDoctorDiagnose() predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorDiagnoseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DoctorDiagnoseTable, DoctorDiagnoseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDoctorDiagnoseWith applies the HasEdge predicate on the "doctor_diagnose" edge with a given conditions (other predicates).
func HasDoctorDiagnoseWith(preds ...predicate.Diagnose) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorDiagnoseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DoctorDiagnoseTable, DoctorDiagnoseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDoctorPrescription applies the HasEdge predicate on the "doctor_prescription" edge.
func HasDoctorPrescription() predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorPrescriptionTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DoctorPrescriptionTable, DoctorPrescriptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDoctorPrescriptionWith applies the HasEdge predicate on the "doctor_prescription" edge with a given conditions (other predicates).
func HasDoctorPrescriptionWith(preds ...predicate.Prescription) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorPrescriptionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DoctorPrescriptionTable, DoctorPrescriptionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Doctor) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Doctor) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Doctor) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		p(s.Not())
	})
}
