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
	"github.com/suyuan32/simple-admin-member-rpc/ent/service"
)

// ServiceUpdate is the builder for updating Service entities.
type ServiceUpdate struct {
	config
	hooks    []Hook
	mutation *ServiceMutation
}

// Where appends a list predicates to the ServiceUpdate builder.
func (su *ServiceUpdate) Where(ps ...predicate.Service) *ServiceUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *ServiceUpdate) SetUpdatedAt(t time.Time) *ServiceUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetStatus sets the "status" field.
func (su *ServiceUpdate) SetStatus(u uint8) *ServiceUpdate {
	su.mutation.ResetStatus()
	su.mutation.SetStatus(u)
	return su
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (su *ServiceUpdate) SetNillableStatus(u *uint8) *ServiceUpdate {
	if u != nil {
		su.SetStatus(*u)
	}
	return su
}

// AddStatus adds u to the "status" field.
func (su *ServiceUpdate) AddStatus(u int8) *ServiceUpdate {
	su.mutation.AddStatus(u)
	return su
}

// ClearStatus clears the value of the "status" field.
func (su *ServiceUpdate) ClearStatus() *ServiceUpdate {
	su.mutation.ClearStatus()
	return su
}

// SetSort sets the "sort" field.
func (su *ServiceUpdate) SetSort(u uint32) *ServiceUpdate {
	su.mutation.ResetSort()
	su.mutation.SetSort(u)
	return su
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (su *ServiceUpdate) SetNillableSort(u *uint32) *ServiceUpdate {
	if u != nil {
		su.SetSort(*u)
	}
	return su
}

// AddSort adds u to the "sort" field.
func (su *ServiceUpdate) AddSort(u int32) *ServiceUpdate {
	su.mutation.AddSort(u)
	return su
}

// SetNameZh sets the "name_zh" field.
func (su *ServiceUpdate) SetNameZh(s string) *ServiceUpdate {
	su.mutation.SetNameZh(s)
	return su
}

// SetNillableNameZh sets the "name_zh" field if the given value is not nil.
func (su *ServiceUpdate) SetNillableNameZh(s *string) *ServiceUpdate {
	if s != nil {
		su.SetNameZh(*s)
	}
	return su
}

// SetNameEn sets the "name_en" field.
func (su *ServiceUpdate) SetNameEn(s string) *ServiceUpdate {
	su.mutation.SetNameEn(s)
	return su
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (su *ServiceUpdate) SetNillableNameEn(s *string) *ServiceUpdate {
	if s != nil {
		su.SetNameEn(*s)
	}
	return su
}

// SetNameRu sets the "name_ru" field.
func (su *ServiceUpdate) SetNameRu(s string) *ServiceUpdate {
	su.mutation.SetNameRu(s)
	return su
}

// SetNillableNameRu sets the "name_ru" field if the given value is not nil.
func (su *ServiceUpdate) SetNillableNameRu(s *string) *ServiceUpdate {
	if s != nil {
		su.SetNameRu(*s)
	}
	return su
}

// SetNameKk sets the "name_kk" field.
func (su *ServiceUpdate) SetNameKk(s string) *ServiceUpdate {
	su.mutation.SetNameKk(s)
	return su
}

// SetNillableNameKk sets the "name_kk" field if the given value is not nil.
func (su *ServiceUpdate) SetNillableNameKk(s *string) *ServiceUpdate {
	if s != nil {
		su.SetNameKk(*s)
	}
	return su
}

// SetDescriptionZh sets the "description_zh" field.
func (su *ServiceUpdate) SetDescriptionZh(s string) *ServiceUpdate {
	su.mutation.SetDescriptionZh(s)
	return su
}

// SetNillableDescriptionZh sets the "description_zh" field if the given value is not nil.
func (su *ServiceUpdate) SetNillableDescriptionZh(s *string) *ServiceUpdate {
	if s != nil {
		su.SetDescriptionZh(*s)
	}
	return su
}

// ClearDescriptionZh clears the value of the "description_zh" field.
func (su *ServiceUpdate) ClearDescriptionZh() *ServiceUpdate {
	su.mutation.ClearDescriptionZh()
	return su
}

// SetDescriptionEn sets the "description_en" field.
func (su *ServiceUpdate) SetDescriptionEn(s string) *ServiceUpdate {
	su.mutation.SetDescriptionEn(s)
	return su
}

// SetNillableDescriptionEn sets the "description_en" field if the given value is not nil.
func (su *ServiceUpdate) SetNillableDescriptionEn(s *string) *ServiceUpdate {
	if s != nil {
		su.SetDescriptionEn(*s)
	}
	return su
}

// ClearDescriptionEn clears the value of the "description_en" field.
func (su *ServiceUpdate) ClearDescriptionEn() *ServiceUpdate {
	su.mutation.ClearDescriptionEn()
	return su
}

// SetDescriptionRu sets the "description_ru" field.
func (su *ServiceUpdate) SetDescriptionRu(s string) *ServiceUpdate {
	su.mutation.SetDescriptionRu(s)
	return su
}

// SetNillableDescriptionRu sets the "description_ru" field if the given value is not nil.
func (su *ServiceUpdate) SetNillableDescriptionRu(s *string) *ServiceUpdate {
	if s != nil {
		su.SetDescriptionRu(*s)
	}
	return su
}

// ClearDescriptionRu clears the value of the "description_ru" field.
func (su *ServiceUpdate) ClearDescriptionRu() *ServiceUpdate {
	su.mutation.ClearDescriptionRu()
	return su
}

// SetDescriptionKk sets the "description_kk" field.
func (su *ServiceUpdate) SetDescriptionKk(s string) *ServiceUpdate {
	su.mutation.SetDescriptionKk(s)
	return su
}

// SetNillableDescriptionKk sets the "description_kk" field if the given value is not nil.
func (su *ServiceUpdate) SetNillableDescriptionKk(s *string) *ServiceUpdate {
	if s != nil {
		su.SetDescriptionKk(*s)
	}
	return su
}

// ClearDescriptionKk clears the value of the "description_kk" field.
func (su *ServiceUpdate) ClearDescriptionKk() *ServiceUpdate {
	su.mutation.ClearDescriptionKk()
	return su
}

// SetCoverURL sets the "cover_url" field.
func (su *ServiceUpdate) SetCoverURL(s string) *ServiceUpdate {
	su.mutation.SetCoverURL(s)
	return su
}

// SetNillableCoverURL sets the "cover_url" field if the given value is not nil.
func (su *ServiceUpdate) SetNillableCoverURL(s *string) *ServiceUpdate {
	if s != nil {
		su.SetCoverURL(*s)
	}
	return su
}

// ClearCoverURL clears the value of the "cover_url" field.
func (su *ServiceUpdate) ClearCoverURL() *ServiceUpdate {
	su.mutation.ClearCoverURL()
	return su
}

// Mutation returns the ServiceMutation object of the builder.
func (su *ServiceUpdate) Mutation() *ServiceMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ServiceUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *ServiceUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ServiceUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ServiceUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *ServiceUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := service.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *ServiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(service.Table, service.Columns, sqlgraph.NewFieldSpec(service.FieldID, field.TypeUint64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(service.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(service.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := su.mutation.AddedStatus(); ok {
		_spec.AddField(service.FieldStatus, field.TypeUint8, value)
	}
	if su.mutation.StatusCleared() {
		_spec.ClearField(service.FieldStatus, field.TypeUint8)
	}
	if value, ok := su.mutation.Sort(); ok {
		_spec.SetField(service.FieldSort, field.TypeUint32, value)
	}
	if value, ok := su.mutation.AddedSort(); ok {
		_spec.AddField(service.FieldSort, field.TypeUint32, value)
	}
	if value, ok := su.mutation.NameZh(); ok {
		_spec.SetField(service.FieldNameZh, field.TypeString, value)
	}
	if value, ok := su.mutation.NameEn(); ok {
		_spec.SetField(service.FieldNameEn, field.TypeString, value)
	}
	if value, ok := su.mutation.NameRu(); ok {
		_spec.SetField(service.FieldNameRu, field.TypeString, value)
	}
	if value, ok := su.mutation.NameKk(); ok {
		_spec.SetField(service.FieldNameKk, field.TypeString, value)
	}
	if value, ok := su.mutation.DescriptionZh(); ok {
		_spec.SetField(service.FieldDescriptionZh, field.TypeString, value)
	}
	if su.mutation.DescriptionZhCleared() {
		_spec.ClearField(service.FieldDescriptionZh, field.TypeString)
	}
	if value, ok := su.mutation.DescriptionEn(); ok {
		_spec.SetField(service.FieldDescriptionEn, field.TypeString, value)
	}
	if su.mutation.DescriptionEnCleared() {
		_spec.ClearField(service.FieldDescriptionEn, field.TypeString)
	}
	if value, ok := su.mutation.DescriptionRu(); ok {
		_spec.SetField(service.FieldDescriptionRu, field.TypeString, value)
	}
	if su.mutation.DescriptionRuCleared() {
		_spec.ClearField(service.FieldDescriptionRu, field.TypeString)
	}
	if value, ok := su.mutation.DescriptionKk(); ok {
		_spec.SetField(service.FieldDescriptionKk, field.TypeString, value)
	}
	if su.mutation.DescriptionKkCleared() {
		_spec.ClearField(service.FieldDescriptionKk, field.TypeString)
	}
	if value, ok := su.mutation.CoverURL(); ok {
		_spec.SetField(service.FieldCoverURL, field.TypeString, value)
	}
	if su.mutation.CoverURLCleared() {
		_spec.ClearField(service.FieldCoverURL, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{service.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// ServiceUpdateOne is the builder for updating a single Service entity.
type ServiceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServiceMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *ServiceUpdateOne) SetUpdatedAt(t time.Time) *ServiceUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetStatus sets the "status" field.
func (suo *ServiceUpdateOne) SetStatus(u uint8) *ServiceUpdateOne {
	suo.mutation.ResetStatus()
	suo.mutation.SetStatus(u)
	return suo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableStatus(u *uint8) *ServiceUpdateOne {
	if u != nil {
		suo.SetStatus(*u)
	}
	return suo
}

// AddStatus adds u to the "status" field.
func (suo *ServiceUpdateOne) AddStatus(u int8) *ServiceUpdateOne {
	suo.mutation.AddStatus(u)
	return suo
}

// ClearStatus clears the value of the "status" field.
func (suo *ServiceUpdateOne) ClearStatus() *ServiceUpdateOne {
	suo.mutation.ClearStatus()
	return suo
}

// SetSort sets the "sort" field.
func (suo *ServiceUpdateOne) SetSort(u uint32) *ServiceUpdateOne {
	suo.mutation.ResetSort()
	suo.mutation.SetSort(u)
	return suo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableSort(u *uint32) *ServiceUpdateOne {
	if u != nil {
		suo.SetSort(*u)
	}
	return suo
}

// AddSort adds u to the "sort" field.
func (suo *ServiceUpdateOne) AddSort(u int32) *ServiceUpdateOne {
	suo.mutation.AddSort(u)
	return suo
}

// SetNameZh sets the "name_zh" field.
func (suo *ServiceUpdateOne) SetNameZh(s string) *ServiceUpdateOne {
	suo.mutation.SetNameZh(s)
	return suo
}

// SetNillableNameZh sets the "name_zh" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableNameZh(s *string) *ServiceUpdateOne {
	if s != nil {
		suo.SetNameZh(*s)
	}
	return suo
}

// SetNameEn sets the "name_en" field.
func (suo *ServiceUpdateOne) SetNameEn(s string) *ServiceUpdateOne {
	suo.mutation.SetNameEn(s)
	return suo
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableNameEn(s *string) *ServiceUpdateOne {
	if s != nil {
		suo.SetNameEn(*s)
	}
	return suo
}

// SetNameRu sets the "name_ru" field.
func (suo *ServiceUpdateOne) SetNameRu(s string) *ServiceUpdateOne {
	suo.mutation.SetNameRu(s)
	return suo
}

// SetNillableNameRu sets the "name_ru" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableNameRu(s *string) *ServiceUpdateOne {
	if s != nil {
		suo.SetNameRu(*s)
	}
	return suo
}

// SetNameKk sets the "name_kk" field.
func (suo *ServiceUpdateOne) SetNameKk(s string) *ServiceUpdateOne {
	suo.mutation.SetNameKk(s)
	return suo
}

// SetNillableNameKk sets the "name_kk" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableNameKk(s *string) *ServiceUpdateOne {
	if s != nil {
		suo.SetNameKk(*s)
	}
	return suo
}

// SetDescriptionZh sets the "description_zh" field.
func (suo *ServiceUpdateOne) SetDescriptionZh(s string) *ServiceUpdateOne {
	suo.mutation.SetDescriptionZh(s)
	return suo
}

// SetNillableDescriptionZh sets the "description_zh" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableDescriptionZh(s *string) *ServiceUpdateOne {
	if s != nil {
		suo.SetDescriptionZh(*s)
	}
	return suo
}

// ClearDescriptionZh clears the value of the "description_zh" field.
func (suo *ServiceUpdateOne) ClearDescriptionZh() *ServiceUpdateOne {
	suo.mutation.ClearDescriptionZh()
	return suo
}

// SetDescriptionEn sets the "description_en" field.
func (suo *ServiceUpdateOne) SetDescriptionEn(s string) *ServiceUpdateOne {
	suo.mutation.SetDescriptionEn(s)
	return suo
}

// SetNillableDescriptionEn sets the "description_en" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableDescriptionEn(s *string) *ServiceUpdateOne {
	if s != nil {
		suo.SetDescriptionEn(*s)
	}
	return suo
}

// ClearDescriptionEn clears the value of the "description_en" field.
func (suo *ServiceUpdateOne) ClearDescriptionEn() *ServiceUpdateOne {
	suo.mutation.ClearDescriptionEn()
	return suo
}

// SetDescriptionRu sets the "description_ru" field.
func (suo *ServiceUpdateOne) SetDescriptionRu(s string) *ServiceUpdateOne {
	suo.mutation.SetDescriptionRu(s)
	return suo
}

// SetNillableDescriptionRu sets the "description_ru" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableDescriptionRu(s *string) *ServiceUpdateOne {
	if s != nil {
		suo.SetDescriptionRu(*s)
	}
	return suo
}

// ClearDescriptionRu clears the value of the "description_ru" field.
func (suo *ServiceUpdateOne) ClearDescriptionRu() *ServiceUpdateOne {
	suo.mutation.ClearDescriptionRu()
	return suo
}

// SetDescriptionKk sets the "description_kk" field.
func (suo *ServiceUpdateOne) SetDescriptionKk(s string) *ServiceUpdateOne {
	suo.mutation.SetDescriptionKk(s)
	return suo
}

// SetNillableDescriptionKk sets the "description_kk" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableDescriptionKk(s *string) *ServiceUpdateOne {
	if s != nil {
		suo.SetDescriptionKk(*s)
	}
	return suo
}

// ClearDescriptionKk clears the value of the "description_kk" field.
func (suo *ServiceUpdateOne) ClearDescriptionKk() *ServiceUpdateOne {
	suo.mutation.ClearDescriptionKk()
	return suo
}

// SetCoverURL sets the "cover_url" field.
func (suo *ServiceUpdateOne) SetCoverURL(s string) *ServiceUpdateOne {
	suo.mutation.SetCoverURL(s)
	return suo
}

// SetNillableCoverURL sets the "cover_url" field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableCoverURL(s *string) *ServiceUpdateOne {
	if s != nil {
		suo.SetCoverURL(*s)
	}
	return suo
}

// ClearCoverURL clears the value of the "cover_url" field.
func (suo *ServiceUpdateOne) ClearCoverURL() *ServiceUpdateOne {
	suo.mutation.ClearCoverURL()
	return suo
}

// Mutation returns the ServiceMutation object of the builder.
func (suo *ServiceUpdateOne) Mutation() *ServiceMutation {
	return suo.mutation
}

// Where appends a list predicates to the ServiceUpdate builder.
func (suo *ServiceUpdateOne) Where(ps ...predicate.Service) *ServiceUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ServiceUpdateOne) Select(field string, fields ...string) *ServiceUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Service entity.
func (suo *ServiceUpdateOne) Save(ctx context.Context) (*Service, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ServiceUpdateOne) SaveX(ctx context.Context) *Service {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ServiceUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ServiceUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *ServiceUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := service.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *ServiceUpdateOne) sqlSave(ctx context.Context) (_node *Service, err error) {
	_spec := sqlgraph.NewUpdateSpec(service.Table, service.Columns, sqlgraph.NewFieldSpec(service.FieldID, field.TypeUint64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Service.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, service.FieldID)
		for _, f := range fields {
			if !service.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != service.FieldID {
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
		_spec.SetField(service.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(service.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := suo.mutation.AddedStatus(); ok {
		_spec.AddField(service.FieldStatus, field.TypeUint8, value)
	}
	if suo.mutation.StatusCleared() {
		_spec.ClearField(service.FieldStatus, field.TypeUint8)
	}
	if value, ok := suo.mutation.Sort(); ok {
		_spec.SetField(service.FieldSort, field.TypeUint32, value)
	}
	if value, ok := suo.mutation.AddedSort(); ok {
		_spec.AddField(service.FieldSort, field.TypeUint32, value)
	}
	if value, ok := suo.mutation.NameZh(); ok {
		_spec.SetField(service.FieldNameZh, field.TypeString, value)
	}
	if value, ok := suo.mutation.NameEn(); ok {
		_spec.SetField(service.FieldNameEn, field.TypeString, value)
	}
	if value, ok := suo.mutation.NameRu(); ok {
		_spec.SetField(service.FieldNameRu, field.TypeString, value)
	}
	if value, ok := suo.mutation.NameKk(); ok {
		_spec.SetField(service.FieldNameKk, field.TypeString, value)
	}
	if value, ok := suo.mutation.DescriptionZh(); ok {
		_spec.SetField(service.FieldDescriptionZh, field.TypeString, value)
	}
	if suo.mutation.DescriptionZhCleared() {
		_spec.ClearField(service.FieldDescriptionZh, field.TypeString)
	}
	if value, ok := suo.mutation.DescriptionEn(); ok {
		_spec.SetField(service.FieldDescriptionEn, field.TypeString, value)
	}
	if suo.mutation.DescriptionEnCleared() {
		_spec.ClearField(service.FieldDescriptionEn, field.TypeString)
	}
	if value, ok := suo.mutation.DescriptionRu(); ok {
		_spec.SetField(service.FieldDescriptionRu, field.TypeString, value)
	}
	if suo.mutation.DescriptionRuCleared() {
		_spec.ClearField(service.FieldDescriptionRu, field.TypeString)
	}
	if value, ok := suo.mutation.DescriptionKk(); ok {
		_spec.SetField(service.FieldDescriptionKk, field.TypeString, value)
	}
	if suo.mutation.DescriptionKkCleared() {
		_spec.ClearField(service.FieldDescriptionKk, field.TypeString)
	}
	if value, ok := suo.mutation.CoverURL(); ok {
		_spec.SetField(service.FieldCoverURL, field.TypeString, value)
	}
	if suo.mutation.CoverURLCleared() {
		_spec.ClearField(service.FieldCoverURL, field.TypeString)
	}
	_node = &Service{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{service.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
