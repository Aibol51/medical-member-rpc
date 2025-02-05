// Code generated by ent, DO NOT EDIT.

package expert

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-member-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldStatus, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v uint32) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldSort, v))
}

// NameZh applies equality check predicate on the "name_zh" field. It's identical to NameZhEQ.
func NameZh(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldNameZh, v))
}

// NameEn applies equality check predicate on the "name_en" field. It's identical to NameEnEQ.
func NameEn(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldNameEn, v))
}

// NameRu applies equality check predicate on the "name_ru" field. It's identical to NameRuEQ.
func NameRu(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldNameRu, v))
}

// NameKk applies equality check predicate on the "name_kk" field. It's identical to NameKkEQ.
func NameKk(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldNameKk, v))
}

// ContentZh applies equality check predicate on the "content_zh" field. It's identical to ContentZhEQ.
func ContentZh(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldContentZh, v))
}

// ContentEn applies equality check predicate on the "content_en" field. It's identical to ContentEnEQ.
func ContentEn(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldContentEn, v))
}

// ContentRu applies equality check predicate on the "content_ru" field. It's identical to ContentRuEQ.
func ContentRu(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldContentRu, v))
}

// ContentKk applies equality check predicate on the "content_kk" field. It's identical to ContentKkEQ.
func ContentKk(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldContentKk, v))
}

// CoverURL applies equality check predicate on the "cover_url" field. It's identical to CoverURLEQ.
func CoverURL(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldCoverURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Expert {
	return predicate.Expert(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Expert {
	return predicate.Expert(sql.FieldNotNull(FieldStatus))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v uint32) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v uint32) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...uint32) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...uint32) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v uint32) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v uint32) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v uint32) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v uint32) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldSort, v))
}

// NameZhEQ applies the EQ predicate on the "name_zh" field.
func NameZhEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldNameZh, v))
}

// NameZhNEQ applies the NEQ predicate on the "name_zh" field.
func NameZhNEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldNameZh, v))
}

// NameZhIn applies the In predicate on the "name_zh" field.
func NameZhIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldNameZh, vs...))
}

// NameZhNotIn applies the NotIn predicate on the "name_zh" field.
func NameZhNotIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldNameZh, vs...))
}

// NameZhGT applies the GT predicate on the "name_zh" field.
func NameZhGT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldNameZh, v))
}

// NameZhGTE applies the GTE predicate on the "name_zh" field.
func NameZhGTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldNameZh, v))
}

// NameZhLT applies the LT predicate on the "name_zh" field.
func NameZhLT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldNameZh, v))
}

// NameZhLTE applies the LTE predicate on the "name_zh" field.
func NameZhLTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldNameZh, v))
}

// NameZhContains applies the Contains predicate on the "name_zh" field.
func NameZhContains(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContains(FieldNameZh, v))
}

// NameZhHasPrefix applies the HasPrefix predicate on the "name_zh" field.
func NameZhHasPrefix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasPrefix(FieldNameZh, v))
}

// NameZhHasSuffix applies the HasSuffix predicate on the "name_zh" field.
func NameZhHasSuffix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasSuffix(FieldNameZh, v))
}

// NameZhIsNil applies the IsNil predicate on the "name_zh" field.
func NameZhIsNil() predicate.Expert {
	return predicate.Expert(sql.FieldIsNull(FieldNameZh))
}

// NameZhNotNil applies the NotNil predicate on the "name_zh" field.
func NameZhNotNil() predicate.Expert {
	return predicate.Expert(sql.FieldNotNull(FieldNameZh))
}

// NameZhEqualFold applies the EqualFold predicate on the "name_zh" field.
func NameZhEqualFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEqualFold(FieldNameZh, v))
}

// NameZhContainsFold applies the ContainsFold predicate on the "name_zh" field.
func NameZhContainsFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContainsFold(FieldNameZh, v))
}

// NameEnEQ applies the EQ predicate on the "name_en" field.
func NameEnEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldNameEn, v))
}

// NameEnNEQ applies the NEQ predicate on the "name_en" field.
func NameEnNEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldNameEn, v))
}

// NameEnIn applies the In predicate on the "name_en" field.
func NameEnIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldNameEn, vs...))
}

// NameEnNotIn applies the NotIn predicate on the "name_en" field.
func NameEnNotIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldNameEn, vs...))
}

// NameEnGT applies the GT predicate on the "name_en" field.
func NameEnGT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldNameEn, v))
}

// NameEnGTE applies the GTE predicate on the "name_en" field.
func NameEnGTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldNameEn, v))
}

// NameEnLT applies the LT predicate on the "name_en" field.
func NameEnLT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldNameEn, v))
}

// NameEnLTE applies the LTE predicate on the "name_en" field.
func NameEnLTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldNameEn, v))
}

// NameEnContains applies the Contains predicate on the "name_en" field.
func NameEnContains(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContains(FieldNameEn, v))
}

// NameEnHasPrefix applies the HasPrefix predicate on the "name_en" field.
func NameEnHasPrefix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasPrefix(FieldNameEn, v))
}

// NameEnHasSuffix applies the HasSuffix predicate on the "name_en" field.
func NameEnHasSuffix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasSuffix(FieldNameEn, v))
}

// NameEnIsNil applies the IsNil predicate on the "name_en" field.
func NameEnIsNil() predicate.Expert {
	return predicate.Expert(sql.FieldIsNull(FieldNameEn))
}

// NameEnNotNil applies the NotNil predicate on the "name_en" field.
func NameEnNotNil() predicate.Expert {
	return predicate.Expert(sql.FieldNotNull(FieldNameEn))
}

// NameEnEqualFold applies the EqualFold predicate on the "name_en" field.
func NameEnEqualFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEqualFold(FieldNameEn, v))
}

// NameEnContainsFold applies the ContainsFold predicate on the "name_en" field.
func NameEnContainsFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContainsFold(FieldNameEn, v))
}

// NameRuEQ applies the EQ predicate on the "name_ru" field.
func NameRuEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldNameRu, v))
}

// NameRuNEQ applies the NEQ predicate on the "name_ru" field.
func NameRuNEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldNameRu, v))
}

// NameRuIn applies the In predicate on the "name_ru" field.
func NameRuIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldNameRu, vs...))
}

// NameRuNotIn applies the NotIn predicate on the "name_ru" field.
func NameRuNotIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldNameRu, vs...))
}

// NameRuGT applies the GT predicate on the "name_ru" field.
func NameRuGT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldNameRu, v))
}

// NameRuGTE applies the GTE predicate on the "name_ru" field.
func NameRuGTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldNameRu, v))
}

// NameRuLT applies the LT predicate on the "name_ru" field.
func NameRuLT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldNameRu, v))
}

// NameRuLTE applies the LTE predicate on the "name_ru" field.
func NameRuLTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldNameRu, v))
}

// NameRuContains applies the Contains predicate on the "name_ru" field.
func NameRuContains(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContains(FieldNameRu, v))
}

// NameRuHasPrefix applies the HasPrefix predicate on the "name_ru" field.
func NameRuHasPrefix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasPrefix(FieldNameRu, v))
}

// NameRuHasSuffix applies the HasSuffix predicate on the "name_ru" field.
func NameRuHasSuffix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasSuffix(FieldNameRu, v))
}

// NameRuIsNil applies the IsNil predicate on the "name_ru" field.
func NameRuIsNil() predicate.Expert {
	return predicate.Expert(sql.FieldIsNull(FieldNameRu))
}

// NameRuNotNil applies the NotNil predicate on the "name_ru" field.
func NameRuNotNil() predicate.Expert {
	return predicate.Expert(sql.FieldNotNull(FieldNameRu))
}

// NameRuEqualFold applies the EqualFold predicate on the "name_ru" field.
func NameRuEqualFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEqualFold(FieldNameRu, v))
}

// NameRuContainsFold applies the ContainsFold predicate on the "name_ru" field.
func NameRuContainsFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContainsFold(FieldNameRu, v))
}

// NameKkEQ applies the EQ predicate on the "name_kk" field.
func NameKkEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldNameKk, v))
}

// NameKkNEQ applies the NEQ predicate on the "name_kk" field.
func NameKkNEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldNameKk, v))
}

// NameKkIn applies the In predicate on the "name_kk" field.
func NameKkIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldNameKk, vs...))
}

// NameKkNotIn applies the NotIn predicate on the "name_kk" field.
func NameKkNotIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldNameKk, vs...))
}

// NameKkGT applies the GT predicate on the "name_kk" field.
func NameKkGT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldNameKk, v))
}

// NameKkGTE applies the GTE predicate on the "name_kk" field.
func NameKkGTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldNameKk, v))
}

// NameKkLT applies the LT predicate on the "name_kk" field.
func NameKkLT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldNameKk, v))
}

// NameKkLTE applies the LTE predicate on the "name_kk" field.
func NameKkLTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldNameKk, v))
}

// NameKkContains applies the Contains predicate on the "name_kk" field.
func NameKkContains(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContains(FieldNameKk, v))
}

// NameKkHasPrefix applies the HasPrefix predicate on the "name_kk" field.
func NameKkHasPrefix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasPrefix(FieldNameKk, v))
}

// NameKkHasSuffix applies the HasSuffix predicate on the "name_kk" field.
func NameKkHasSuffix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasSuffix(FieldNameKk, v))
}

// NameKkIsNil applies the IsNil predicate on the "name_kk" field.
func NameKkIsNil() predicate.Expert {
	return predicate.Expert(sql.FieldIsNull(FieldNameKk))
}

// NameKkNotNil applies the NotNil predicate on the "name_kk" field.
func NameKkNotNil() predicate.Expert {
	return predicate.Expert(sql.FieldNotNull(FieldNameKk))
}

// NameKkEqualFold applies the EqualFold predicate on the "name_kk" field.
func NameKkEqualFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEqualFold(FieldNameKk, v))
}

// NameKkContainsFold applies the ContainsFold predicate on the "name_kk" field.
func NameKkContainsFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContainsFold(FieldNameKk, v))
}

// ContentZhEQ applies the EQ predicate on the "content_zh" field.
func ContentZhEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldContentZh, v))
}

// ContentZhNEQ applies the NEQ predicate on the "content_zh" field.
func ContentZhNEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldContentZh, v))
}

// ContentZhIn applies the In predicate on the "content_zh" field.
func ContentZhIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldContentZh, vs...))
}

// ContentZhNotIn applies the NotIn predicate on the "content_zh" field.
func ContentZhNotIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldContentZh, vs...))
}

// ContentZhGT applies the GT predicate on the "content_zh" field.
func ContentZhGT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldContentZh, v))
}

// ContentZhGTE applies the GTE predicate on the "content_zh" field.
func ContentZhGTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldContentZh, v))
}

// ContentZhLT applies the LT predicate on the "content_zh" field.
func ContentZhLT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldContentZh, v))
}

// ContentZhLTE applies the LTE predicate on the "content_zh" field.
func ContentZhLTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldContentZh, v))
}

// ContentZhContains applies the Contains predicate on the "content_zh" field.
func ContentZhContains(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContains(FieldContentZh, v))
}

// ContentZhHasPrefix applies the HasPrefix predicate on the "content_zh" field.
func ContentZhHasPrefix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasPrefix(FieldContentZh, v))
}

// ContentZhHasSuffix applies the HasSuffix predicate on the "content_zh" field.
func ContentZhHasSuffix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasSuffix(FieldContentZh, v))
}

// ContentZhIsNil applies the IsNil predicate on the "content_zh" field.
func ContentZhIsNil() predicate.Expert {
	return predicate.Expert(sql.FieldIsNull(FieldContentZh))
}

// ContentZhNotNil applies the NotNil predicate on the "content_zh" field.
func ContentZhNotNil() predicate.Expert {
	return predicate.Expert(sql.FieldNotNull(FieldContentZh))
}

// ContentZhEqualFold applies the EqualFold predicate on the "content_zh" field.
func ContentZhEqualFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEqualFold(FieldContentZh, v))
}

// ContentZhContainsFold applies the ContainsFold predicate on the "content_zh" field.
func ContentZhContainsFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContainsFold(FieldContentZh, v))
}

// ContentEnEQ applies the EQ predicate on the "content_en" field.
func ContentEnEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldContentEn, v))
}

// ContentEnNEQ applies the NEQ predicate on the "content_en" field.
func ContentEnNEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldContentEn, v))
}

// ContentEnIn applies the In predicate on the "content_en" field.
func ContentEnIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldContentEn, vs...))
}

// ContentEnNotIn applies the NotIn predicate on the "content_en" field.
func ContentEnNotIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldContentEn, vs...))
}

// ContentEnGT applies the GT predicate on the "content_en" field.
func ContentEnGT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldContentEn, v))
}

// ContentEnGTE applies the GTE predicate on the "content_en" field.
func ContentEnGTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldContentEn, v))
}

// ContentEnLT applies the LT predicate on the "content_en" field.
func ContentEnLT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldContentEn, v))
}

// ContentEnLTE applies the LTE predicate on the "content_en" field.
func ContentEnLTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldContentEn, v))
}

// ContentEnContains applies the Contains predicate on the "content_en" field.
func ContentEnContains(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContains(FieldContentEn, v))
}

// ContentEnHasPrefix applies the HasPrefix predicate on the "content_en" field.
func ContentEnHasPrefix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasPrefix(FieldContentEn, v))
}

// ContentEnHasSuffix applies the HasSuffix predicate on the "content_en" field.
func ContentEnHasSuffix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasSuffix(FieldContentEn, v))
}

// ContentEnIsNil applies the IsNil predicate on the "content_en" field.
func ContentEnIsNil() predicate.Expert {
	return predicate.Expert(sql.FieldIsNull(FieldContentEn))
}

// ContentEnNotNil applies the NotNil predicate on the "content_en" field.
func ContentEnNotNil() predicate.Expert {
	return predicate.Expert(sql.FieldNotNull(FieldContentEn))
}

// ContentEnEqualFold applies the EqualFold predicate on the "content_en" field.
func ContentEnEqualFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEqualFold(FieldContentEn, v))
}

// ContentEnContainsFold applies the ContainsFold predicate on the "content_en" field.
func ContentEnContainsFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContainsFold(FieldContentEn, v))
}

// ContentRuEQ applies the EQ predicate on the "content_ru" field.
func ContentRuEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldContentRu, v))
}

// ContentRuNEQ applies the NEQ predicate on the "content_ru" field.
func ContentRuNEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldContentRu, v))
}

// ContentRuIn applies the In predicate on the "content_ru" field.
func ContentRuIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldContentRu, vs...))
}

// ContentRuNotIn applies the NotIn predicate on the "content_ru" field.
func ContentRuNotIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldContentRu, vs...))
}

// ContentRuGT applies the GT predicate on the "content_ru" field.
func ContentRuGT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldContentRu, v))
}

// ContentRuGTE applies the GTE predicate on the "content_ru" field.
func ContentRuGTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldContentRu, v))
}

// ContentRuLT applies the LT predicate on the "content_ru" field.
func ContentRuLT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldContentRu, v))
}

// ContentRuLTE applies the LTE predicate on the "content_ru" field.
func ContentRuLTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldContentRu, v))
}

// ContentRuContains applies the Contains predicate on the "content_ru" field.
func ContentRuContains(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContains(FieldContentRu, v))
}

// ContentRuHasPrefix applies the HasPrefix predicate on the "content_ru" field.
func ContentRuHasPrefix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasPrefix(FieldContentRu, v))
}

// ContentRuHasSuffix applies the HasSuffix predicate on the "content_ru" field.
func ContentRuHasSuffix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasSuffix(FieldContentRu, v))
}

// ContentRuIsNil applies the IsNil predicate on the "content_ru" field.
func ContentRuIsNil() predicate.Expert {
	return predicate.Expert(sql.FieldIsNull(FieldContentRu))
}

// ContentRuNotNil applies the NotNil predicate on the "content_ru" field.
func ContentRuNotNil() predicate.Expert {
	return predicate.Expert(sql.FieldNotNull(FieldContentRu))
}

// ContentRuEqualFold applies the EqualFold predicate on the "content_ru" field.
func ContentRuEqualFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEqualFold(FieldContentRu, v))
}

// ContentRuContainsFold applies the ContainsFold predicate on the "content_ru" field.
func ContentRuContainsFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContainsFold(FieldContentRu, v))
}

// ContentKkEQ applies the EQ predicate on the "content_kk" field.
func ContentKkEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldContentKk, v))
}

// ContentKkNEQ applies the NEQ predicate on the "content_kk" field.
func ContentKkNEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldContentKk, v))
}

// ContentKkIn applies the In predicate on the "content_kk" field.
func ContentKkIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldContentKk, vs...))
}

// ContentKkNotIn applies the NotIn predicate on the "content_kk" field.
func ContentKkNotIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldContentKk, vs...))
}

// ContentKkGT applies the GT predicate on the "content_kk" field.
func ContentKkGT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldContentKk, v))
}

// ContentKkGTE applies the GTE predicate on the "content_kk" field.
func ContentKkGTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldContentKk, v))
}

// ContentKkLT applies the LT predicate on the "content_kk" field.
func ContentKkLT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldContentKk, v))
}

// ContentKkLTE applies the LTE predicate on the "content_kk" field.
func ContentKkLTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldContentKk, v))
}

// ContentKkContains applies the Contains predicate on the "content_kk" field.
func ContentKkContains(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContains(FieldContentKk, v))
}

// ContentKkHasPrefix applies the HasPrefix predicate on the "content_kk" field.
func ContentKkHasPrefix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasPrefix(FieldContentKk, v))
}

// ContentKkHasSuffix applies the HasSuffix predicate on the "content_kk" field.
func ContentKkHasSuffix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasSuffix(FieldContentKk, v))
}

// ContentKkIsNil applies the IsNil predicate on the "content_kk" field.
func ContentKkIsNil() predicate.Expert {
	return predicate.Expert(sql.FieldIsNull(FieldContentKk))
}

// ContentKkNotNil applies the NotNil predicate on the "content_kk" field.
func ContentKkNotNil() predicate.Expert {
	return predicate.Expert(sql.FieldNotNull(FieldContentKk))
}

// ContentKkEqualFold applies the EqualFold predicate on the "content_kk" field.
func ContentKkEqualFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEqualFold(FieldContentKk, v))
}

// ContentKkContainsFold applies the ContainsFold predicate on the "content_kk" field.
func ContentKkContainsFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContainsFold(FieldContentKk, v))
}

// CoverURLEQ applies the EQ predicate on the "cover_url" field.
func CoverURLEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEQ(FieldCoverURL, v))
}

// CoverURLNEQ applies the NEQ predicate on the "cover_url" field.
func CoverURLNEQ(v string) predicate.Expert {
	return predicate.Expert(sql.FieldNEQ(FieldCoverURL, v))
}

// CoverURLIn applies the In predicate on the "cover_url" field.
func CoverURLIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldIn(FieldCoverURL, vs...))
}

// CoverURLNotIn applies the NotIn predicate on the "cover_url" field.
func CoverURLNotIn(vs ...string) predicate.Expert {
	return predicate.Expert(sql.FieldNotIn(FieldCoverURL, vs...))
}

// CoverURLGT applies the GT predicate on the "cover_url" field.
func CoverURLGT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGT(FieldCoverURL, v))
}

// CoverURLGTE applies the GTE predicate on the "cover_url" field.
func CoverURLGTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldGTE(FieldCoverURL, v))
}

// CoverURLLT applies the LT predicate on the "cover_url" field.
func CoverURLLT(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLT(FieldCoverURL, v))
}

// CoverURLLTE applies the LTE predicate on the "cover_url" field.
func CoverURLLTE(v string) predicate.Expert {
	return predicate.Expert(sql.FieldLTE(FieldCoverURL, v))
}

// CoverURLContains applies the Contains predicate on the "cover_url" field.
func CoverURLContains(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContains(FieldCoverURL, v))
}

// CoverURLHasPrefix applies the HasPrefix predicate on the "cover_url" field.
func CoverURLHasPrefix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasPrefix(FieldCoverURL, v))
}

// CoverURLHasSuffix applies the HasSuffix predicate on the "cover_url" field.
func CoverURLHasSuffix(v string) predicate.Expert {
	return predicate.Expert(sql.FieldHasSuffix(FieldCoverURL, v))
}

// CoverURLIsNil applies the IsNil predicate on the "cover_url" field.
func CoverURLIsNil() predicate.Expert {
	return predicate.Expert(sql.FieldIsNull(FieldCoverURL))
}

// CoverURLNotNil applies the NotNil predicate on the "cover_url" field.
func CoverURLNotNil() predicate.Expert {
	return predicate.Expert(sql.FieldNotNull(FieldCoverURL))
}

// CoverURLEqualFold applies the EqualFold predicate on the "cover_url" field.
func CoverURLEqualFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldEqualFold(FieldCoverURL, v))
}

// CoverURLContainsFold applies the ContainsFold predicate on the "cover_url" field.
func CoverURLContainsFold(v string) predicate.Expert {
	return predicate.Expert(sql.FieldContainsFold(FieldCoverURL, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Expert) predicate.Expert {
	return predicate.Expert(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Expert) predicate.Expert {
	return predicate.Expert(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Expert) predicate.Expert {
	return predicate.Expert(sql.NotPredicates(p))
}
