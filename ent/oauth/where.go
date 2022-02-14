// Code generated by entc, DO NOT EDIT.

package oauth

import (
	"enttest/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Method applies equality check predicate on the "method" field. It's identical to MethodEQ.
func Method(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMethod), v))
	})
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIcon), v))
	})
}

// Appid applies equality check predicate on the "appid" field. It's identical to AppidEQ.
func Appid(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppid), v))
	})
}

// Secret applies equality check predicate on the "secret" field. It's identical to SecretEQ.
func Secret(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecret), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.OAuth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuth(func(s *sql.Selector) {
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
func NameNotIn(vs ...string) predicate.OAuth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuth(func(s *sql.Selector) {
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
func NameGT(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMethod), v))
	})
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMethod), v))
	})
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...string) predicate.OAuth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMethod), v...))
	})
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...string) predicate.OAuth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMethod), v...))
	})
}

// MethodGT applies the GT predicate on the "method" field.
func MethodGT(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMethod), v))
	})
}

// MethodGTE applies the GTE predicate on the "method" field.
func MethodGTE(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMethod), v))
	})
}

// MethodLT applies the LT predicate on the "method" field.
func MethodLT(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMethod), v))
	})
}

// MethodLTE applies the LTE predicate on the "method" field.
func MethodLTE(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMethod), v))
	})
}

// MethodContains applies the Contains predicate on the "method" field.
func MethodContains(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMethod), v))
	})
}

// MethodHasPrefix applies the HasPrefix predicate on the "method" field.
func MethodHasPrefix(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMethod), v))
	})
}

// MethodHasSuffix applies the HasSuffix predicate on the "method" field.
func MethodHasSuffix(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMethod), v))
	})
}

// MethodEqualFold applies the EqualFold predicate on the "method" field.
func MethodEqualFold(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMethod), v))
	})
}

// MethodContainsFold applies the ContainsFold predicate on the "method" field.
func MethodContainsFold(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMethod), v))
	})
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIcon), v))
	})
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIcon), v))
	})
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.OAuth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIcon), v...))
	})
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.OAuth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIcon), v...))
	})
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIcon), v))
	})
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIcon), v))
	})
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIcon), v))
	})
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIcon), v))
	})
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIcon), v))
	})
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIcon), v))
	})
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIcon), v))
	})
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIcon), v))
	})
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIcon), v))
	})
}

// AppidEQ applies the EQ predicate on the "appid" field.
func AppidEQ(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppid), v))
	})
}

// AppidNEQ applies the NEQ predicate on the "appid" field.
func AppidNEQ(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppid), v))
	})
}

// AppidIn applies the In predicate on the "appid" field.
func AppidIn(vs ...string) predicate.OAuth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppid), v...))
	})
}

// AppidNotIn applies the NotIn predicate on the "appid" field.
func AppidNotIn(vs ...string) predicate.OAuth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppid), v...))
	})
}

// AppidGT applies the GT predicate on the "appid" field.
func AppidGT(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppid), v))
	})
}

// AppidGTE applies the GTE predicate on the "appid" field.
func AppidGTE(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppid), v))
	})
}

// AppidLT applies the LT predicate on the "appid" field.
func AppidLT(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppid), v))
	})
}

// AppidLTE applies the LTE predicate on the "appid" field.
func AppidLTE(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppid), v))
	})
}

// AppidContains applies the Contains predicate on the "appid" field.
func AppidContains(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAppid), v))
	})
}

// AppidHasPrefix applies the HasPrefix predicate on the "appid" field.
func AppidHasPrefix(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAppid), v))
	})
}

// AppidHasSuffix applies the HasSuffix predicate on the "appid" field.
func AppidHasSuffix(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAppid), v))
	})
}

// AppidEqualFold applies the EqualFold predicate on the "appid" field.
func AppidEqualFold(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAppid), v))
	})
}

// AppidContainsFold applies the ContainsFold predicate on the "appid" field.
func AppidContainsFold(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAppid), v))
	})
}

// SecretEQ applies the EQ predicate on the "secret" field.
func SecretEQ(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecret), v))
	})
}

// SecretNEQ applies the NEQ predicate on the "secret" field.
func SecretNEQ(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSecret), v))
	})
}

// SecretIn applies the In predicate on the "secret" field.
func SecretIn(vs ...string) predicate.OAuth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSecret), v...))
	})
}

// SecretNotIn applies the NotIn predicate on the "secret" field.
func SecretNotIn(vs ...string) predicate.OAuth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSecret), v...))
	})
}

// SecretGT applies the GT predicate on the "secret" field.
func SecretGT(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSecret), v))
	})
}

// SecretGTE applies the GTE predicate on the "secret" field.
func SecretGTE(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSecret), v))
	})
}

// SecretLT applies the LT predicate on the "secret" field.
func SecretLT(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSecret), v))
	})
}

// SecretLTE applies the LTE predicate on the "secret" field.
func SecretLTE(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSecret), v))
	})
}

// SecretContains applies the Contains predicate on the "secret" field.
func SecretContains(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSecret), v))
	})
}

// SecretHasPrefix applies the HasPrefix predicate on the "secret" field.
func SecretHasPrefix(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSecret), v))
	})
}

// SecretHasSuffix applies the HasSuffix predicate on the "secret" field.
func SecretHasSuffix(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSecret), v))
	})
}

// SecretEqualFold applies the EqualFold predicate on the "secret" field.
func SecretEqualFold(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSecret), v))
	})
}

// SecretContainsFold applies the ContainsFold predicate on the "secret" field.
func SecretContainsFold(v string) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSecret), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OAuth) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OAuth) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
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
func Not(p predicate.OAuth) predicate.OAuth {
	return predicate.OAuth(func(s *sql.Selector) {
		p(s.Not())
	})
}
