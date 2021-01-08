// Code generated by entc, DO NOT EDIT.

package roomtype

import (
	"github.com/dhetporn/team08/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Bathroom applies equality check predicate on the "bathroom" field. It's identical to BathroomEQ.
func Bathroom(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBathroom), v))
	})
}

// Toilets applies equality check predicate on the "toilets" field. It's identical to ToiletsEQ.
func Toilets(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToilets), v))
	})
}

// Areasize applies equality check predicate on the "areasize" field. It's identical to AreasizeEQ.
func Areasize(v float64) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAreasize), v))
	})
}

// Etc applies equality check predicate on the "etc" field. It's identical to EtcEQ.
func Etc(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEtc), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Roomtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Roomtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// BathroomEQ applies the EQ predicate on the "bathroom" field.
func BathroomEQ(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBathroom), v))
	})
}

// BathroomNEQ applies the NEQ predicate on the "bathroom" field.
func BathroomNEQ(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBathroom), v))
	})
}

// BathroomIn applies the In predicate on the "bathroom" field.
func BathroomIn(vs ...int) predicate.Roomtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBathroom), v...))
	})
}

// BathroomNotIn applies the NotIn predicate on the "bathroom" field.
func BathroomNotIn(vs ...int) predicate.Roomtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBathroom), v...))
	})
}

// BathroomGT applies the GT predicate on the "bathroom" field.
func BathroomGT(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBathroom), v))
	})
}

// BathroomGTE applies the GTE predicate on the "bathroom" field.
func BathroomGTE(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBathroom), v))
	})
}

// BathroomLT applies the LT predicate on the "bathroom" field.
func BathroomLT(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBathroom), v))
	})
}

// BathroomLTE applies the LTE predicate on the "bathroom" field.
func BathroomLTE(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBathroom), v))
	})
}

// ToiletsEQ applies the EQ predicate on the "toilets" field.
func ToiletsEQ(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToilets), v))
	})
}

// ToiletsNEQ applies the NEQ predicate on the "toilets" field.
func ToiletsNEQ(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToilets), v))
	})
}

// ToiletsIn applies the In predicate on the "toilets" field.
func ToiletsIn(vs ...int) predicate.Roomtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldToilets), v...))
	})
}

// ToiletsNotIn applies the NotIn predicate on the "toilets" field.
func ToiletsNotIn(vs ...int) predicate.Roomtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldToilets), v...))
	})
}

// ToiletsGT applies the GT predicate on the "toilets" field.
func ToiletsGT(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToilets), v))
	})
}

// ToiletsGTE applies the GTE predicate on the "toilets" field.
func ToiletsGTE(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToilets), v))
	})
}

// ToiletsLT applies the LT predicate on the "toilets" field.
func ToiletsLT(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToilets), v))
	})
}

// ToiletsLTE applies the LTE predicate on the "toilets" field.
func ToiletsLTE(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToilets), v))
	})
}

// AreasizeEQ applies the EQ predicate on the "areasize" field.
func AreasizeEQ(v float64) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAreasize), v))
	})
}

// AreasizeNEQ applies the NEQ predicate on the "areasize" field.
func AreasizeNEQ(v float64) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAreasize), v))
	})
}

// AreasizeIn applies the In predicate on the "areasize" field.
func AreasizeIn(vs ...float64) predicate.Roomtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAreasize), v...))
	})
}

// AreasizeNotIn applies the NotIn predicate on the "areasize" field.
func AreasizeNotIn(vs ...float64) predicate.Roomtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAreasize), v...))
	})
}

// AreasizeGT applies the GT predicate on the "areasize" field.
func AreasizeGT(v float64) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAreasize), v))
	})
}

// AreasizeGTE applies the GTE predicate on the "areasize" field.
func AreasizeGTE(v float64) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAreasize), v))
	})
}

// AreasizeLT applies the LT predicate on the "areasize" field.
func AreasizeLT(v float64) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAreasize), v))
	})
}

// AreasizeLTE applies the LTE predicate on the "areasize" field.
func AreasizeLTE(v float64) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAreasize), v))
	})
}

// EtcEQ applies the EQ predicate on the "etc" field.
func EtcEQ(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEtc), v))
	})
}

// EtcNEQ applies the NEQ predicate on the "etc" field.
func EtcNEQ(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEtc), v))
	})
}

// EtcIn applies the In predicate on the "etc" field.
func EtcIn(vs ...string) predicate.Roomtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEtc), v...))
	})
}

// EtcNotIn applies the NotIn predicate on the "etc" field.
func EtcNotIn(vs ...string) predicate.Roomtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEtc), v...))
	})
}

// EtcGT applies the GT predicate on the "etc" field.
func EtcGT(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEtc), v))
	})
}

// EtcGTE applies the GTE predicate on the "etc" field.
func EtcGTE(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEtc), v))
	})
}

// EtcLT applies the LT predicate on the "etc" field.
func EtcLT(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEtc), v))
	})
}

// EtcLTE applies the LTE predicate on the "etc" field.
func EtcLTE(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEtc), v))
	})
}

// EtcContains applies the Contains predicate on the "etc" field.
func EtcContains(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEtc), v))
	})
}

// EtcHasPrefix applies the HasPrefix predicate on the "etc" field.
func EtcHasPrefix(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEtc), v))
	})
}

// EtcHasSuffix applies the HasSuffix predicate on the "etc" field.
func EtcHasSuffix(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEtc), v))
	})
}

// EtcEqualFold applies the EqualFold predicate on the "etc" field.
func EtcEqualFold(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEtc), v))
	})
}

// EtcContainsFold applies the ContainsFold predicate on the "etc" field.
func EtcContainsFold(v string) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEtc), v))
	})
}

// HasRooms applies the HasEdge predicate on the "rooms" edge.
func HasRooms() predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoomsTable, RoomsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomsWith applies the HasEdge predicate on the "rooms" edge with a given conditions (other predicates).
func HasRoomsWith(preds ...predicate.Room) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoomsTable, RoomsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Roomtype) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Roomtype) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
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
func Not(p predicate.Roomtype) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		p(s.Not())
	})
}
