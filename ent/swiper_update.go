// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-member-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-member-rpc/ent/swiper"
)

// SwiperUpdate is the builder for updating Swiper entities.
type SwiperUpdate struct {
	config
	hooks    []Hook
	mutation *SwiperMutation
}

// Where appends a list predicates to the SwiperUpdate builder.
func (su *SwiperUpdate) Where(ps ...predicate.Swiper) *SwiperUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SwiperUpdate) SetUpdatedAt(t time.Time) *SwiperUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetStatus sets the "status" field.
func (su *SwiperUpdate) SetStatus(u uint8) *SwiperUpdate {
	su.mutation.ResetStatus()
	su.mutation.SetStatus(u)
	return su
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableStatus(u *uint8) *SwiperUpdate {
	if u != nil {
		su.SetStatus(*u)
	}
	return su
}

// AddStatus adds u to the "status" field.
func (su *SwiperUpdate) AddStatus(u int8) *SwiperUpdate {
	su.mutation.AddStatus(u)
	return su
}

// ClearStatus clears the value of the "status" field.
func (su *SwiperUpdate) ClearStatus() *SwiperUpdate {
	su.mutation.ClearStatus()
	return su
}

// SetSort sets the "sort" field.
func (su *SwiperUpdate) SetSort(u uint32) *SwiperUpdate {
	su.mutation.ResetSort()
	su.mutation.SetSort(u)
	return su
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableSort(u *uint32) *SwiperUpdate {
	if u != nil {
		su.SetSort(*u)
	}
	return su
}

// AddSort adds u to the "sort" field.
func (su *SwiperUpdate) AddSort(u int32) *SwiperUpdate {
	su.mutation.AddSort(u)
	return su
}

// SetTitleZh sets the "title_zh" field.
func (su *SwiperUpdate) SetTitleZh(s string) *SwiperUpdate {
	su.mutation.SetTitleZh(s)
	return su
}

// SetNillableTitleZh sets the "title_zh" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableTitleZh(s *string) *SwiperUpdate {
	if s != nil {
		su.SetTitleZh(*s)
	}
	return su
}

// ClearTitleZh clears the value of the "title_zh" field.
func (su *SwiperUpdate) ClearTitleZh() *SwiperUpdate {
	su.mutation.ClearTitleZh()
	return su
}

// SetTitleEn sets the "title_en" field.
func (su *SwiperUpdate) SetTitleEn(s string) *SwiperUpdate {
	su.mutation.SetTitleEn(s)
	return su
}

// SetNillableTitleEn sets the "title_en" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableTitleEn(s *string) *SwiperUpdate {
	if s != nil {
		su.SetTitleEn(*s)
	}
	return su
}

// ClearTitleEn clears the value of the "title_en" field.
func (su *SwiperUpdate) ClearTitleEn() *SwiperUpdate {
	su.mutation.ClearTitleEn()
	return su
}

// SetTitleRu sets the "title_ru" field.
func (su *SwiperUpdate) SetTitleRu(s string) *SwiperUpdate {
	su.mutation.SetTitleRu(s)
	return su
}

// SetNillableTitleRu sets the "title_ru" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableTitleRu(s *string) *SwiperUpdate {
	if s != nil {
		su.SetTitleRu(*s)
	}
	return su
}

// ClearTitleRu clears the value of the "title_ru" field.
func (su *SwiperUpdate) ClearTitleRu() *SwiperUpdate {
	su.mutation.ClearTitleRu()
	return su
}

// SetTitleKk sets the "title_kk" field.
func (su *SwiperUpdate) SetTitleKk(s string) *SwiperUpdate {
	su.mutation.SetTitleKk(s)
	return su
}

// SetNillableTitleKk sets the "title_kk" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableTitleKk(s *string) *SwiperUpdate {
	if s != nil {
		su.SetTitleKk(*s)
	}
	return su
}

// ClearTitleKk clears the value of the "title_kk" field.
func (su *SwiperUpdate) ClearTitleKk() *SwiperUpdate {
	su.mutation.ClearTitleKk()
	return su
}

// SetBannerZh sets the "banner_zh" field.
func (su *SwiperUpdate) SetBannerZh(s string) *SwiperUpdate {
	su.mutation.SetBannerZh(s)
	return su
}

// SetNillableBannerZh sets the "banner_zh" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableBannerZh(s *string) *SwiperUpdate {
	if s != nil {
		su.SetBannerZh(*s)
	}
	return su
}

// ClearBannerZh clears the value of the "banner_zh" field.
func (su *SwiperUpdate) ClearBannerZh() *SwiperUpdate {
	su.mutation.ClearBannerZh()
	return su
}

// SetBannerEn sets the "banner_en" field.
func (su *SwiperUpdate) SetBannerEn(s string) *SwiperUpdate {
	su.mutation.SetBannerEn(s)
	return su
}

// SetNillableBannerEn sets the "banner_en" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableBannerEn(s *string) *SwiperUpdate {
	if s != nil {
		su.SetBannerEn(*s)
	}
	return su
}

// ClearBannerEn clears the value of the "banner_en" field.
func (su *SwiperUpdate) ClearBannerEn() *SwiperUpdate {
	su.mutation.ClearBannerEn()
	return su
}

// SetBannerRu sets the "banner_ru" field.
func (su *SwiperUpdate) SetBannerRu(s string) *SwiperUpdate {
	su.mutation.SetBannerRu(s)
	return su
}

// SetNillableBannerRu sets the "banner_ru" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableBannerRu(s *string) *SwiperUpdate {
	if s != nil {
		su.SetBannerRu(*s)
	}
	return su
}

// ClearBannerRu clears the value of the "banner_ru" field.
func (su *SwiperUpdate) ClearBannerRu() *SwiperUpdate {
	su.mutation.ClearBannerRu()
	return su
}

// SetBannerKk sets the "banner_kk" field.
func (su *SwiperUpdate) SetBannerKk(s string) *SwiperUpdate {
	su.mutation.SetBannerKk(s)
	return su
}

// SetNillableBannerKk sets the "banner_kk" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableBannerKk(s *string) *SwiperUpdate {
	if s != nil {
		su.SetBannerKk(*s)
	}
	return su
}

// ClearBannerKk clears the value of the "banner_kk" field.
func (su *SwiperUpdate) ClearBannerKk() *SwiperUpdate {
	su.mutation.ClearBannerKk()
	return su
}

// SetContentZh sets the "content_zh" field.
func (su *SwiperUpdate) SetContentZh(s string) *SwiperUpdate {
	su.mutation.SetContentZh(s)
	return su
}

// SetNillableContentZh sets the "content_zh" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableContentZh(s *string) *SwiperUpdate {
	if s != nil {
		su.SetContentZh(*s)
	}
	return su
}

// ClearContentZh clears the value of the "content_zh" field.
func (su *SwiperUpdate) ClearContentZh() *SwiperUpdate {
	su.mutation.ClearContentZh()
	return su
}

// SetContentEn sets the "content_en" field.
func (su *SwiperUpdate) SetContentEn(s string) *SwiperUpdate {
	su.mutation.SetContentEn(s)
	return su
}

// SetNillableContentEn sets the "content_en" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableContentEn(s *string) *SwiperUpdate {
	if s != nil {
		su.SetContentEn(*s)
	}
	return su
}

// ClearContentEn clears the value of the "content_en" field.
func (su *SwiperUpdate) ClearContentEn() *SwiperUpdate {
	su.mutation.ClearContentEn()
	return su
}

// SetContentRu sets the "content_ru" field.
func (su *SwiperUpdate) SetContentRu(s string) *SwiperUpdate {
	su.mutation.SetContentRu(s)
	return su
}

// SetNillableContentRu sets the "content_ru" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableContentRu(s *string) *SwiperUpdate {
	if s != nil {
		su.SetContentRu(*s)
	}
	return su
}

// ClearContentRu clears the value of the "content_ru" field.
func (su *SwiperUpdate) ClearContentRu() *SwiperUpdate {
	su.mutation.ClearContentRu()
	return su
}

// SetContentKk sets the "content_kk" field.
func (su *SwiperUpdate) SetContentKk(s string) *SwiperUpdate {
	su.mutation.SetContentKk(s)
	return su
}

// SetNillableContentKk sets the "content_kk" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableContentKk(s *string) *SwiperUpdate {
	if s != nil {
		su.SetContentKk(*s)
	}
	return su
}

// ClearContentKk clears the value of the "content_kk" field.
func (su *SwiperUpdate) ClearContentKk() *SwiperUpdate {
	su.mutation.ClearContentKk()
	return su
}

// SetJumpURL sets the "jump_url" field.
func (su *SwiperUpdate) SetJumpURL(s string) *SwiperUpdate {
	su.mutation.SetJumpURL(s)
	return su
}

// SetNillableJumpURL sets the "jump_url" field if the given value is not nil.
func (su *SwiperUpdate) SetNillableJumpURL(s *string) *SwiperUpdate {
	if s != nil {
		su.SetJumpURL(*s)
	}
	return su
}

// ClearJumpURL clears the value of the "jump_url" field.
func (su *SwiperUpdate) ClearJumpURL() *SwiperUpdate {
	su.mutation.ClearJumpURL()
	return su
}

// Mutation returns the SwiperMutation object of the builder.
func (su *SwiperUpdate) Mutation() *SwiperMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SwiperUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SwiperUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SwiperUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SwiperUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SwiperUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := swiper.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *SwiperUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(swiper.Table, swiper.Columns, sqlgraph.NewFieldSpec(swiper.FieldID, field.TypeUint64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(swiper.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(swiper.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := su.mutation.AddedStatus(); ok {
		_spec.AddField(swiper.FieldStatus, field.TypeUint8, value)
	}
	if su.mutation.StatusCleared() {
		_spec.ClearField(swiper.FieldStatus, field.TypeUint8)
	}
	if value, ok := su.mutation.Sort(); ok {
		_spec.SetField(swiper.FieldSort, field.TypeUint32, value)
	}
	if value, ok := su.mutation.AddedSort(); ok {
		_spec.AddField(swiper.FieldSort, field.TypeUint32, value)
	}
	if value, ok := su.mutation.TitleZh(); ok {
		_spec.SetField(swiper.FieldTitleZh, field.TypeString, value)
	}
	if su.mutation.TitleZhCleared() {
		_spec.ClearField(swiper.FieldTitleZh, field.TypeString)
	}
	if value, ok := su.mutation.TitleEn(); ok {
		_spec.SetField(swiper.FieldTitleEn, field.TypeString, value)
	}
	if su.mutation.TitleEnCleared() {
		_spec.ClearField(swiper.FieldTitleEn, field.TypeString)
	}
	if value, ok := su.mutation.TitleRu(); ok {
		_spec.SetField(swiper.FieldTitleRu, field.TypeString, value)
	}
	if su.mutation.TitleRuCleared() {
		_spec.ClearField(swiper.FieldTitleRu, field.TypeString)
	}
	if value, ok := su.mutation.TitleKk(); ok {
		_spec.SetField(swiper.FieldTitleKk, field.TypeString, value)
	}
	if su.mutation.TitleKkCleared() {
		_spec.ClearField(swiper.FieldTitleKk, field.TypeString)
	}
	if value, ok := su.mutation.BannerZh(); ok {
		_spec.SetField(swiper.FieldBannerZh, field.TypeString, value)
	}
	if su.mutation.BannerZhCleared() {
		_spec.ClearField(swiper.FieldBannerZh, field.TypeString)
	}
	if value, ok := su.mutation.BannerEn(); ok {
		_spec.SetField(swiper.FieldBannerEn, field.TypeString, value)
	}
	if su.mutation.BannerEnCleared() {
		_spec.ClearField(swiper.FieldBannerEn, field.TypeString)
	}
	if value, ok := su.mutation.BannerRu(); ok {
		_spec.SetField(swiper.FieldBannerRu, field.TypeString, value)
	}
	if su.mutation.BannerRuCleared() {
		_spec.ClearField(swiper.FieldBannerRu, field.TypeString)
	}
	if value, ok := su.mutation.BannerKk(); ok {
		_spec.SetField(swiper.FieldBannerKk, field.TypeString, value)
	}
	if su.mutation.BannerKkCleared() {
		_spec.ClearField(swiper.FieldBannerKk, field.TypeString)
	}
	if value, ok := su.mutation.ContentZh(); ok {
		_spec.SetField(swiper.FieldContentZh, field.TypeString, value)
	}
	if su.mutation.ContentZhCleared() {
		_spec.ClearField(swiper.FieldContentZh, field.TypeString)
	}
	if value, ok := su.mutation.ContentEn(); ok {
		_spec.SetField(swiper.FieldContentEn, field.TypeString, value)
	}
	if su.mutation.ContentEnCleared() {
		_spec.ClearField(swiper.FieldContentEn, field.TypeString)
	}
	if value, ok := su.mutation.ContentRu(); ok {
		_spec.SetField(swiper.FieldContentRu, field.TypeString, value)
	}
	if su.mutation.ContentRuCleared() {
		_spec.ClearField(swiper.FieldContentRu, field.TypeString)
	}
	if value, ok := su.mutation.ContentKk(); ok {
		_spec.SetField(swiper.FieldContentKk, field.TypeString, value)
	}
	if su.mutation.ContentKkCleared() {
		_spec.ClearField(swiper.FieldContentKk, field.TypeString)
	}
	if value, ok := su.mutation.JumpURL(); ok {
		_spec.SetField(swiper.FieldJumpURL, field.TypeString, value)
	}
	if su.mutation.JumpURLCleared() {
		_spec.ClearField(swiper.FieldJumpURL, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{swiper.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SwiperUpdateOne is the builder for updating a single Swiper entity.
type SwiperUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SwiperMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SwiperUpdateOne) SetUpdatedAt(t time.Time) *SwiperUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetStatus sets the "status" field.
func (suo *SwiperUpdateOne) SetStatus(u uint8) *SwiperUpdateOne {
	suo.mutation.ResetStatus()
	suo.mutation.SetStatus(u)
	return suo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableStatus(u *uint8) *SwiperUpdateOne {
	if u != nil {
		suo.SetStatus(*u)
	}
	return suo
}

// AddStatus adds u to the "status" field.
func (suo *SwiperUpdateOne) AddStatus(u int8) *SwiperUpdateOne {
	suo.mutation.AddStatus(u)
	return suo
}

// ClearStatus clears the value of the "status" field.
func (suo *SwiperUpdateOne) ClearStatus() *SwiperUpdateOne {
	suo.mutation.ClearStatus()
	return suo
}

// SetSort sets the "sort" field.
func (suo *SwiperUpdateOne) SetSort(u uint32) *SwiperUpdateOne {
	suo.mutation.ResetSort()
	suo.mutation.SetSort(u)
	return suo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableSort(u *uint32) *SwiperUpdateOne {
	if u != nil {
		suo.SetSort(*u)
	}
	return suo
}

// AddSort adds u to the "sort" field.
func (suo *SwiperUpdateOne) AddSort(u int32) *SwiperUpdateOne {
	suo.mutation.AddSort(u)
	return suo
}

// SetTitleZh sets the "title_zh" field.
func (suo *SwiperUpdateOne) SetTitleZh(s string) *SwiperUpdateOne {
	suo.mutation.SetTitleZh(s)
	return suo
}

// SetNillableTitleZh sets the "title_zh" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableTitleZh(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetTitleZh(*s)
	}
	return suo
}

// ClearTitleZh clears the value of the "title_zh" field.
func (suo *SwiperUpdateOne) ClearTitleZh() *SwiperUpdateOne {
	suo.mutation.ClearTitleZh()
	return suo
}

// SetTitleEn sets the "title_en" field.
func (suo *SwiperUpdateOne) SetTitleEn(s string) *SwiperUpdateOne {
	suo.mutation.SetTitleEn(s)
	return suo
}

// SetNillableTitleEn sets the "title_en" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableTitleEn(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetTitleEn(*s)
	}
	return suo
}

// ClearTitleEn clears the value of the "title_en" field.
func (suo *SwiperUpdateOne) ClearTitleEn() *SwiperUpdateOne {
	suo.mutation.ClearTitleEn()
	return suo
}

// SetTitleRu sets the "title_ru" field.
func (suo *SwiperUpdateOne) SetTitleRu(s string) *SwiperUpdateOne {
	suo.mutation.SetTitleRu(s)
	return suo
}

// SetNillableTitleRu sets the "title_ru" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableTitleRu(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetTitleRu(*s)
	}
	return suo
}

// ClearTitleRu clears the value of the "title_ru" field.
func (suo *SwiperUpdateOne) ClearTitleRu() *SwiperUpdateOne {
	suo.mutation.ClearTitleRu()
	return suo
}

// SetTitleKk sets the "title_kk" field.
func (suo *SwiperUpdateOne) SetTitleKk(s string) *SwiperUpdateOne {
	suo.mutation.SetTitleKk(s)
	return suo
}

// SetNillableTitleKk sets the "title_kk" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableTitleKk(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetTitleKk(*s)
	}
	return suo
}

// ClearTitleKk clears the value of the "title_kk" field.
func (suo *SwiperUpdateOne) ClearTitleKk() *SwiperUpdateOne {
	suo.mutation.ClearTitleKk()
	return suo
}

// SetBannerZh sets the "banner_zh" field.
func (suo *SwiperUpdateOne) SetBannerZh(s string) *SwiperUpdateOne {
	suo.mutation.SetBannerZh(s)
	return suo
}

// SetNillableBannerZh sets the "banner_zh" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableBannerZh(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetBannerZh(*s)
	}
	return suo
}

// ClearBannerZh clears the value of the "banner_zh" field.
func (suo *SwiperUpdateOne) ClearBannerZh() *SwiperUpdateOne {
	suo.mutation.ClearBannerZh()
	return suo
}

// SetBannerEn sets the "banner_en" field.
func (suo *SwiperUpdateOne) SetBannerEn(s string) *SwiperUpdateOne {
	suo.mutation.SetBannerEn(s)
	return suo
}

// SetNillableBannerEn sets the "banner_en" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableBannerEn(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetBannerEn(*s)
	}
	return suo
}

// ClearBannerEn clears the value of the "banner_en" field.
func (suo *SwiperUpdateOne) ClearBannerEn() *SwiperUpdateOne {
	suo.mutation.ClearBannerEn()
	return suo
}

// SetBannerRu sets the "banner_ru" field.
func (suo *SwiperUpdateOne) SetBannerRu(s string) *SwiperUpdateOne {
	suo.mutation.SetBannerRu(s)
	return suo
}

// SetNillableBannerRu sets the "banner_ru" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableBannerRu(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetBannerRu(*s)
	}
	return suo
}

// ClearBannerRu clears the value of the "banner_ru" field.
func (suo *SwiperUpdateOne) ClearBannerRu() *SwiperUpdateOne {
	suo.mutation.ClearBannerRu()
	return suo
}

// SetBannerKk sets the "banner_kk" field.
func (suo *SwiperUpdateOne) SetBannerKk(s string) *SwiperUpdateOne {
	suo.mutation.SetBannerKk(s)
	return suo
}

// SetNillableBannerKk sets the "banner_kk" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableBannerKk(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetBannerKk(*s)
	}
	return suo
}

// ClearBannerKk clears the value of the "banner_kk" field.
func (suo *SwiperUpdateOne) ClearBannerKk() *SwiperUpdateOne {
	suo.mutation.ClearBannerKk()
	return suo
}

// SetContentZh sets the "content_zh" field.
func (suo *SwiperUpdateOne) SetContentZh(s string) *SwiperUpdateOne {
	suo.mutation.SetContentZh(s)
	return suo
}

// SetNillableContentZh sets the "content_zh" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableContentZh(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetContentZh(*s)
	}
	return suo
}

// ClearContentZh clears the value of the "content_zh" field.
func (suo *SwiperUpdateOne) ClearContentZh() *SwiperUpdateOne {
	suo.mutation.ClearContentZh()
	return suo
}

// SetContentEn sets the "content_en" field.
func (suo *SwiperUpdateOne) SetContentEn(s string) *SwiperUpdateOne {
	suo.mutation.SetContentEn(s)
	return suo
}

// SetNillableContentEn sets the "content_en" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableContentEn(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetContentEn(*s)
	}
	return suo
}

// ClearContentEn clears the value of the "content_en" field.
func (suo *SwiperUpdateOne) ClearContentEn() *SwiperUpdateOne {
	suo.mutation.ClearContentEn()
	return suo
}

// SetContentRu sets the "content_ru" field.
func (suo *SwiperUpdateOne) SetContentRu(s string) *SwiperUpdateOne {
	suo.mutation.SetContentRu(s)
	return suo
}

// SetNillableContentRu sets the "content_ru" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableContentRu(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetContentRu(*s)
	}
	return suo
}

// ClearContentRu clears the value of the "content_ru" field.
func (suo *SwiperUpdateOne) ClearContentRu() *SwiperUpdateOne {
	suo.mutation.ClearContentRu()
	return suo
}

// SetContentKk sets the "content_kk" field.
func (suo *SwiperUpdateOne) SetContentKk(s string) *SwiperUpdateOne {
	suo.mutation.SetContentKk(s)
	return suo
}

// SetNillableContentKk sets the "content_kk" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableContentKk(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetContentKk(*s)
	}
	return suo
}

// ClearContentKk clears the value of the "content_kk" field.
func (suo *SwiperUpdateOne) ClearContentKk() *SwiperUpdateOne {
	suo.mutation.ClearContentKk()
	return suo
}

// SetJumpURL sets the "jump_url" field.
func (suo *SwiperUpdateOne) SetJumpURL(s string) *SwiperUpdateOne {
	suo.mutation.SetJumpURL(s)
	return suo
}

// SetNillableJumpURL sets the "jump_url" field if the given value is not nil.
func (suo *SwiperUpdateOne) SetNillableJumpURL(s *string) *SwiperUpdateOne {
	if s != nil {
		suo.SetJumpURL(*s)
	}
	return suo
}

// ClearJumpURL clears the value of the "jump_url" field.
func (suo *SwiperUpdateOne) ClearJumpURL() *SwiperUpdateOne {
	suo.mutation.ClearJumpURL()
	return suo
}

// Mutation returns the SwiperMutation object of the builder.
func (suo *SwiperUpdateOne) Mutation() *SwiperMutation {
	return suo.mutation
}

// Where appends a list predicates to the SwiperUpdate builder.
func (suo *SwiperUpdateOne) Where(ps ...predicate.Swiper) *SwiperUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SwiperUpdateOne) Select(field string, fields ...string) *SwiperUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Swiper entity.
func (suo *SwiperUpdateOne) Save(ctx context.Context) (*Swiper, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SwiperUpdateOne) SaveX(ctx context.Context) *Swiper {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SwiperUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SwiperUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SwiperUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := swiper.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *SwiperUpdateOne) sqlSave(ctx context.Context) (_node *Swiper, err error) {
	_spec := sqlgraph.NewUpdateSpec(swiper.Table, swiper.Columns, sqlgraph.NewFieldSpec(swiper.FieldID, field.TypeUint64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Swiper.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, swiper.FieldID)
		for _, f := range fields {
			if !swiper.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != swiper.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(swiper.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(swiper.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := suo.mutation.AddedStatus(); ok {
		_spec.AddField(swiper.FieldStatus, field.TypeUint8, value)
	}
	if suo.mutation.StatusCleared() {
		_spec.ClearField(swiper.FieldStatus, field.TypeUint8)
	}
	if value, ok := suo.mutation.Sort(); ok {
		_spec.SetField(swiper.FieldSort, field.TypeUint32, value)
	}
	if value, ok := suo.mutation.AddedSort(); ok {
		_spec.AddField(swiper.FieldSort, field.TypeUint32, value)
	}
	if value, ok := suo.mutation.TitleZh(); ok {
		_spec.SetField(swiper.FieldTitleZh, field.TypeString, value)
	}
	if suo.mutation.TitleZhCleared() {
		_spec.ClearField(swiper.FieldTitleZh, field.TypeString)
	}
	if value, ok := suo.mutation.TitleEn(); ok {
		_spec.SetField(swiper.FieldTitleEn, field.TypeString, value)
	}
	if suo.mutation.TitleEnCleared() {
		_spec.ClearField(swiper.FieldTitleEn, field.TypeString)
	}
	if value, ok := suo.mutation.TitleRu(); ok {
		_spec.SetField(swiper.FieldTitleRu, field.TypeString, value)
	}
	if suo.mutation.TitleRuCleared() {
		_spec.ClearField(swiper.FieldTitleRu, field.TypeString)
	}
	if value, ok := suo.mutation.TitleKk(); ok {
		_spec.SetField(swiper.FieldTitleKk, field.TypeString, value)
	}
	if suo.mutation.TitleKkCleared() {
		_spec.ClearField(swiper.FieldTitleKk, field.TypeString)
	}
	if value, ok := suo.mutation.BannerZh(); ok {
		_spec.SetField(swiper.FieldBannerZh, field.TypeString, value)
	}
	if suo.mutation.BannerZhCleared() {
		_spec.ClearField(swiper.FieldBannerZh, field.TypeString)
	}
	if value, ok := suo.mutation.BannerEn(); ok {
		_spec.SetField(swiper.FieldBannerEn, field.TypeString, value)
	}
	if suo.mutation.BannerEnCleared() {
		_spec.ClearField(swiper.FieldBannerEn, field.TypeString)
	}
	if value, ok := suo.mutation.BannerRu(); ok {
		_spec.SetField(swiper.FieldBannerRu, field.TypeString, value)
	}
	if suo.mutation.BannerRuCleared() {
		_spec.ClearField(swiper.FieldBannerRu, field.TypeString)
	}
	if value, ok := suo.mutation.BannerKk(); ok {
		_spec.SetField(swiper.FieldBannerKk, field.TypeString, value)
	}
	if suo.mutation.BannerKkCleared() {
		_spec.ClearField(swiper.FieldBannerKk, field.TypeString)
	}
	if value, ok := suo.mutation.ContentZh(); ok {
		_spec.SetField(swiper.FieldContentZh, field.TypeString, value)
	}
	if suo.mutation.ContentZhCleared() {
		_spec.ClearField(swiper.FieldContentZh, field.TypeString)
	}
	if value, ok := suo.mutation.ContentEn(); ok {
		_spec.SetField(swiper.FieldContentEn, field.TypeString, value)
	}
	if suo.mutation.ContentEnCleared() {
		_spec.ClearField(swiper.FieldContentEn, field.TypeString)
	}
	if value, ok := suo.mutation.ContentRu(); ok {
		_spec.SetField(swiper.FieldContentRu, field.TypeString, value)
	}
	if suo.mutation.ContentRuCleared() {
		_spec.ClearField(swiper.FieldContentRu, field.TypeString)
	}
	if value, ok := suo.mutation.ContentKk(); ok {
		_spec.SetField(swiper.FieldContentKk, field.TypeString, value)
	}
	if suo.mutation.ContentKkCleared() {
		_spec.ClearField(swiper.FieldContentKk, field.TypeString)
	}
	if value, ok := suo.mutation.JumpURL(); ok {
		_spec.SetField(swiper.FieldJumpURL, field.TypeString, value)
	}
	if suo.mutation.JumpURLCleared() {
		_spec.ClearField(swiper.FieldJumpURL, field.TypeString)
	}
	_node = &Swiper{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{swiper.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
