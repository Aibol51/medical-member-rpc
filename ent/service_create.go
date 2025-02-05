// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-member-rpc/ent/service"
)

// ServiceCreate is the builder for creating a Service entity.
type ServiceCreate struct {
	config
	mutation *ServiceMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *ServiceCreate) SetCreatedAt(t time.Time) *ServiceCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableCreatedAt(t *time.Time) *ServiceCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *ServiceCreate) SetUpdatedAt(t time.Time) *ServiceCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableUpdatedAt(t *time.Time) *ServiceCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetStatus sets the "status" field.
func (sc *ServiceCreate) SetStatus(u uint8) *ServiceCreate {
	sc.mutation.SetStatus(u)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableStatus(u *uint8) *ServiceCreate {
	if u != nil {
		sc.SetStatus(*u)
	}
	return sc
}

// SetSort sets the "sort" field.
func (sc *ServiceCreate) SetSort(u uint32) *ServiceCreate {
	sc.mutation.SetSort(u)
	return sc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableSort(u *uint32) *ServiceCreate {
	if u != nil {
		sc.SetSort(*u)
	}
	return sc
}

// SetNameZh sets the "name_zh" field.
func (sc *ServiceCreate) SetNameZh(s string) *ServiceCreate {
	sc.mutation.SetNameZh(s)
	return sc
}

// SetNameEn sets the "name_en" field.
func (sc *ServiceCreate) SetNameEn(s string) *ServiceCreate {
	sc.mutation.SetNameEn(s)
	return sc
}

// SetNameRu sets the "name_ru" field.
func (sc *ServiceCreate) SetNameRu(s string) *ServiceCreate {
	sc.mutation.SetNameRu(s)
	return sc
}

// SetNameKk sets the "name_kk" field.
func (sc *ServiceCreate) SetNameKk(s string) *ServiceCreate {
	sc.mutation.SetNameKk(s)
	return sc
}

// SetDescriptionZh sets the "description_zh" field.
func (sc *ServiceCreate) SetDescriptionZh(s string) *ServiceCreate {
	sc.mutation.SetDescriptionZh(s)
	return sc
}

// SetNillableDescriptionZh sets the "description_zh" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableDescriptionZh(s *string) *ServiceCreate {
	if s != nil {
		sc.SetDescriptionZh(*s)
	}
	return sc
}

// SetDescriptionEn sets the "description_en" field.
func (sc *ServiceCreate) SetDescriptionEn(s string) *ServiceCreate {
	sc.mutation.SetDescriptionEn(s)
	return sc
}

// SetNillableDescriptionEn sets the "description_en" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableDescriptionEn(s *string) *ServiceCreate {
	if s != nil {
		sc.SetDescriptionEn(*s)
	}
	return sc
}

// SetDescriptionRu sets the "description_ru" field.
func (sc *ServiceCreate) SetDescriptionRu(s string) *ServiceCreate {
	sc.mutation.SetDescriptionRu(s)
	return sc
}

// SetNillableDescriptionRu sets the "description_ru" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableDescriptionRu(s *string) *ServiceCreate {
	if s != nil {
		sc.SetDescriptionRu(*s)
	}
	return sc
}

// SetDescriptionKk sets the "description_kk" field.
func (sc *ServiceCreate) SetDescriptionKk(s string) *ServiceCreate {
	sc.mutation.SetDescriptionKk(s)
	return sc
}

// SetNillableDescriptionKk sets the "description_kk" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableDescriptionKk(s *string) *ServiceCreate {
	if s != nil {
		sc.SetDescriptionKk(*s)
	}
	return sc
}

// SetCoverURL sets the "cover_url" field.
func (sc *ServiceCreate) SetCoverURL(s string) *ServiceCreate {
	sc.mutation.SetCoverURL(s)
	return sc
}

// SetNillableCoverURL sets the "cover_url" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableCoverURL(s *string) *ServiceCreate {
	if s != nil {
		sc.SetCoverURL(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *ServiceCreate) SetID(u uint64) *ServiceCreate {
	sc.mutation.SetID(u)
	return sc
}

// Mutation returns the ServiceMutation object of the builder.
func (sc *ServiceCreate) Mutation() *ServiceMutation {
	return sc.mutation
}

// Save creates the Service in the database.
func (sc *ServiceCreate) Save(ctx context.Context) (*Service, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ServiceCreate) SaveX(ctx context.Context) *Service {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ServiceCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ServiceCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ServiceCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := service.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := service.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.Status(); !ok {
		v := service.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.Sort(); !ok {
		v := service.DefaultSort
		sc.mutation.SetSort(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ServiceCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Service.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Service.updated_at"`)}
	}
	if _, ok := sc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Service.sort"`)}
	}
	if _, ok := sc.mutation.NameZh(); !ok {
		return &ValidationError{Name: "name_zh", err: errors.New(`ent: missing required field "Service.name_zh"`)}
	}
	if _, ok := sc.mutation.NameEn(); !ok {
		return &ValidationError{Name: "name_en", err: errors.New(`ent: missing required field "Service.name_en"`)}
	}
	if _, ok := sc.mutation.NameRu(); !ok {
		return &ValidationError{Name: "name_ru", err: errors.New(`ent: missing required field "Service.name_ru"`)}
	}
	if _, ok := sc.mutation.NameKk(); !ok {
		return &ValidationError{Name: "name_kk", err: errors.New(`ent: missing required field "Service.name_kk"`)}
	}
	return nil
}

func (sc *ServiceCreate) sqlSave(ctx context.Context) (*Service, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ServiceCreate) createSpec() (*Service, *sqlgraph.CreateSpec) {
	var (
		_node = &Service{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(service.Table, sqlgraph.NewFieldSpec(service.FieldID, field.TypeUint64))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(service.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(service.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(service.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.Sort(); ok {
		_spec.SetField(service.FieldSort, field.TypeUint32, value)
		_node.Sort = value
	}
	if value, ok := sc.mutation.NameZh(); ok {
		_spec.SetField(service.FieldNameZh, field.TypeString, value)
		_node.NameZh = value
	}
	if value, ok := sc.mutation.NameEn(); ok {
		_spec.SetField(service.FieldNameEn, field.TypeString, value)
		_node.NameEn = value
	}
	if value, ok := sc.mutation.NameRu(); ok {
		_spec.SetField(service.FieldNameRu, field.TypeString, value)
		_node.NameRu = value
	}
	if value, ok := sc.mutation.NameKk(); ok {
		_spec.SetField(service.FieldNameKk, field.TypeString, value)
		_node.NameKk = value
	}
	if value, ok := sc.mutation.DescriptionZh(); ok {
		_spec.SetField(service.FieldDescriptionZh, field.TypeString, value)
		_node.DescriptionZh = value
	}
	if value, ok := sc.mutation.DescriptionEn(); ok {
		_spec.SetField(service.FieldDescriptionEn, field.TypeString, value)
		_node.DescriptionEn = value
	}
	if value, ok := sc.mutation.DescriptionRu(); ok {
		_spec.SetField(service.FieldDescriptionRu, field.TypeString, value)
		_node.DescriptionRu = value
	}
	if value, ok := sc.mutation.DescriptionKk(); ok {
		_spec.SetField(service.FieldDescriptionKk, field.TypeString, value)
		_node.DescriptionKk = value
	}
	if value, ok := sc.mutation.CoverURL(); ok {
		_spec.SetField(service.FieldCoverURL, field.TypeString, value)
		_node.CoverURL = value
	}
	return _node, _spec
}

// ServiceCreateBulk is the builder for creating many Service entities in bulk.
type ServiceCreateBulk struct {
	config
	err      error
	builders []*ServiceCreate
}

// Save creates the Service entities in the database.
func (scb *ServiceCreateBulk) Save(ctx context.Context) ([]*Service, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Service, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServiceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ServiceCreateBulk) SaveX(ctx context.Context) []*Service {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ServiceCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ServiceCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
