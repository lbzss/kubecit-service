// Code generated by ent, DO NOT EDIT.

package setting

import (
	"kubecit-service/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Setting {
	return predicate.Setting(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Setting {
	return predicate.Setting(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Setting {
	return predicate.Setting(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Setting {
	return predicate.Setting(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Setting {
	return predicate.Setting(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Setting {
	return predicate.Setting(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Setting {
	return predicate.Setting(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldName, v))
}

// Detail applies equality check predicate on the "detail" field. It's identical to DetailEQ.
func Detail(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldDetail, v))
}

// Cover applies equality check predicate on the "cover" field. It's identical to CoverEQ.
func Cover(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldCover, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Setting {
	return predicate.Setting(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Setting {
	return predicate.Setting(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Setting {
	return predicate.Setting(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Setting {
	return predicate.Setting(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Setting {
	return predicate.Setting(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Setting {
	return predicate.Setting(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Setting {
	return predicate.Setting(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Setting {
	return predicate.Setting(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Setting {
	return predicate.Setting(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Setting {
	return predicate.Setting(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Setting {
	return predicate.Setting(sql.FieldContainsFold(FieldName, v))
}

// DetailEQ applies the EQ predicate on the "detail" field.
func DetailEQ(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldDetail, v))
}

// DetailNEQ applies the NEQ predicate on the "detail" field.
func DetailNEQ(v string) predicate.Setting {
	return predicate.Setting(sql.FieldNEQ(FieldDetail, v))
}

// DetailIn applies the In predicate on the "detail" field.
func DetailIn(vs ...string) predicate.Setting {
	return predicate.Setting(sql.FieldIn(FieldDetail, vs...))
}

// DetailNotIn applies the NotIn predicate on the "detail" field.
func DetailNotIn(vs ...string) predicate.Setting {
	return predicate.Setting(sql.FieldNotIn(FieldDetail, vs...))
}

// DetailGT applies the GT predicate on the "detail" field.
func DetailGT(v string) predicate.Setting {
	return predicate.Setting(sql.FieldGT(FieldDetail, v))
}

// DetailGTE applies the GTE predicate on the "detail" field.
func DetailGTE(v string) predicate.Setting {
	return predicate.Setting(sql.FieldGTE(FieldDetail, v))
}

// DetailLT applies the LT predicate on the "detail" field.
func DetailLT(v string) predicate.Setting {
	return predicate.Setting(sql.FieldLT(FieldDetail, v))
}

// DetailLTE applies the LTE predicate on the "detail" field.
func DetailLTE(v string) predicate.Setting {
	return predicate.Setting(sql.FieldLTE(FieldDetail, v))
}

// DetailContains applies the Contains predicate on the "detail" field.
func DetailContains(v string) predicate.Setting {
	return predicate.Setting(sql.FieldContains(FieldDetail, v))
}

// DetailHasPrefix applies the HasPrefix predicate on the "detail" field.
func DetailHasPrefix(v string) predicate.Setting {
	return predicate.Setting(sql.FieldHasPrefix(FieldDetail, v))
}

// DetailHasSuffix applies the HasSuffix predicate on the "detail" field.
func DetailHasSuffix(v string) predicate.Setting {
	return predicate.Setting(sql.FieldHasSuffix(FieldDetail, v))
}

// DetailEqualFold applies the EqualFold predicate on the "detail" field.
func DetailEqualFold(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEqualFold(FieldDetail, v))
}

// DetailContainsFold applies the ContainsFold predicate on the "detail" field.
func DetailContainsFold(v string) predicate.Setting {
	return predicate.Setting(sql.FieldContainsFold(FieldDetail, v))
}

// CoverEQ applies the EQ predicate on the "cover" field.
func CoverEQ(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEQ(FieldCover, v))
}

// CoverNEQ applies the NEQ predicate on the "cover" field.
func CoverNEQ(v string) predicate.Setting {
	return predicate.Setting(sql.FieldNEQ(FieldCover, v))
}

// CoverIn applies the In predicate on the "cover" field.
func CoverIn(vs ...string) predicate.Setting {
	return predicate.Setting(sql.FieldIn(FieldCover, vs...))
}

// CoverNotIn applies the NotIn predicate on the "cover" field.
func CoverNotIn(vs ...string) predicate.Setting {
	return predicate.Setting(sql.FieldNotIn(FieldCover, vs...))
}

// CoverGT applies the GT predicate on the "cover" field.
func CoverGT(v string) predicate.Setting {
	return predicate.Setting(sql.FieldGT(FieldCover, v))
}

// CoverGTE applies the GTE predicate on the "cover" field.
func CoverGTE(v string) predicate.Setting {
	return predicate.Setting(sql.FieldGTE(FieldCover, v))
}

// CoverLT applies the LT predicate on the "cover" field.
func CoverLT(v string) predicate.Setting {
	return predicate.Setting(sql.FieldLT(FieldCover, v))
}

// CoverLTE applies the LTE predicate on the "cover" field.
func CoverLTE(v string) predicate.Setting {
	return predicate.Setting(sql.FieldLTE(FieldCover, v))
}

// CoverContains applies the Contains predicate on the "cover" field.
func CoverContains(v string) predicate.Setting {
	return predicate.Setting(sql.FieldContains(FieldCover, v))
}

// CoverHasPrefix applies the HasPrefix predicate on the "cover" field.
func CoverHasPrefix(v string) predicate.Setting {
	return predicate.Setting(sql.FieldHasPrefix(FieldCover, v))
}

// CoverHasSuffix applies the HasSuffix predicate on the "cover" field.
func CoverHasSuffix(v string) predicate.Setting {
	return predicate.Setting(sql.FieldHasSuffix(FieldCover, v))
}

// CoverEqualFold applies the EqualFold predicate on the "cover" field.
func CoverEqualFold(v string) predicate.Setting {
	return predicate.Setting(sql.FieldEqualFold(FieldCover, v))
}

// CoverContainsFold applies the ContainsFold predicate on the "cover" field.
func CoverContainsFold(v string) predicate.Setting {
	return predicate.Setting(sql.FieldContainsFold(FieldCover, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Setting) predicate.Setting {
	return predicate.Setting(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Setting) predicate.Setting {
	return predicate.Setting(func(s *sql.Selector) {
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
func Not(p predicate.Setting) predicate.Setting {
	return predicate.Setting(func(s *sql.Selector) {
		p(s.Not())
	})
}
