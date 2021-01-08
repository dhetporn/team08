// Code generated by entc, DO NOT EDIT.

package coveredperson

import (
	"github.com/dhetporn/team08/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
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
func IDGT(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// HasPatient applies the HasEdge predicate on the "Patient" edge.
func HasPatient() predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "Patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patient) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSchemeType applies the HasEdge predicate on the "SchemeType" edge.
func HasSchemeType() predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SchemeTypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SchemeTypeTable, SchemeTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSchemeTypeWith applies the HasEdge predicate on the "SchemeType" edge with a given conditions (other predicates).
func HasSchemeTypeWith(preds ...predicate.SchemeType) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SchemeTypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SchemeTypeTable, SchemeTypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFund applies the HasEdge predicate on the "Fund" edge.
func HasFund() predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FundTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FundTable, FundColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFundWith applies the HasEdge predicate on the "Fund" edge with a given conditions (other predicates).
func HasFundWith(preds ...predicate.Fund) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FundInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FundTable, FundColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCertificate applies the HasEdge predicate on the "Certificate" edge.
func HasCertificate() predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CertificateTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CertificateTable, CertificateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCertificateWith applies the HasEdge predicate on the "Certificate" edge with a given conditions (other predicates).
func HasCertificateWith(preds ...predicate.Certificate) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CertificateInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CertificateTable, CertificateColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.CoveredPerson) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.CoveredPerson) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
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
func Not(p predicate.CoveredPerson) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		p(s.Not())
	})
}
