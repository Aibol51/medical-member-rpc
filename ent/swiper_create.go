// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-member-rpc/ent/swiper"
)

// SwiperCreate is the builder for creating a Swiper entity.
type SwiperCreate struct {
	config
	mutation *SwiperMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *SwiperCreate) SetCreatedAt(t time.Time) *SwiperCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableCreatedAt(t *time.Time) *SwiperCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SwiperCreate) SetUpdatedAt(t time.Time) *SwiperCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableUpdatedAt(t *time.Time) *SwiperCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetStatus sets the "status" field.
func (sc *SwiperCreate) SetStatus(u uint8) *SwiperCreate {
	sc.mutation.SetStatus(u)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableStatus(u *uint8) *SwiperCreate {
	if u != nil {
		sc.SetStatus(*u)
	}
	return sc
}

// SetSort sets the "sort" field.
func (sc *SwiperCreate) SetSort(u uint32) *SwiperCreate {
	sc.mutation.SetSort(u)
	return sc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableSort(u *uint32) *SwiperCreate {
	if u != nil {
		sc.SetSort(*u)
	}
	return sc
}

// SetTitleZh sets the "title_zh" field.
func (sc *SwiperCreate) SetTitleZh(s string) *SwiperCreate {
	sc.mutation.SetTitleZh(s)
	return sc
}

// SetNillableTitleZh sets the "title_zh" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableTitleZh(s *string) *SwiperCreate {
	if s != nil {
		sc.SetTitleZh(*s)
	}
	return sc
}

// SetTitleEn sets the "title_en" field.
func (sc *SwiperCreate) SetTitleEn(s string) *SwiperCreate {
	sc.mutation.SetTitleEn(s)
	return sc
}

// SetNillableTitleEn sets the "title_en" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableTitleEn(s *string) *SwiperCreate {
	if s != nil {
		sc.SetTitleEn(*s)
	}
	return sc
}

// SetTitleRu sets the "title_ru" field.
func (sc *SwiperCreate) SetTitleRu(s string) *SwiperCreate {
	sc.mutation.SetTitleRu(s)
	return sc
}

// SetNillableTitleRu sets the "title_ru" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableTitleRu(s *string) *SwiperCreate {
	if s != nil {
		sc.SetTitleRu(*s)
	}
	return sc
}

// SetTitleKk sets the "title_kk" field.
func (sc *SwiperCreate) SetTitleKk(s string) *SwiperCreate {
	sc.mutation.SetTitleKk(s)
	return sc
}

// SetNillableTitleKk sets the "title_kk" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableTitleKk(s *string) *SwiperCreate {
	if s != nil {
		sc.SetTitleKk(*s)
	}
	return sc
}

// SetBannerZh sets the "banner_zh" field.
func (sc *SwiperCreate) SetBannerZh(s string) *SwiperCreate {
	sc.mutation.SetBannerZh(s)
	return sc
}

// SetNillableBannerZh sets the "banner_zh" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableBannerZh(s *string) *SwiperCreate {
	if s != nil {
		sc.SetBannerZh(*s)
	}
	return sc
}

// SetBannerEn sets the "banner_en" field.
func (sc *SwiperCreate) SetBannerEn(s string) *SwiperCreate {
	sc.mutation.SetBannerEn(s)
	return sc
}

// SetNillableBannerEn sets the "banner_en" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableBannerEn(s *string) *SwiperCreate {
	if s != nil {
		sc.SetBannerEn(*s)
	}
	return sc
}

// SetBannerRu sets the "banner_ru" field.
func (sc *SwiperCreate) SetBannerRu(s string) *SwiperCreate {
	sc.mutation.SetBannerRu(s)
	return sc
}

// SetNillableBannerRu sets the "banner_ru" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableBannerRu(s *string) *SwiperCreate {
	if s != nil {
		sc.SetBannerRu(*s)
	}
	return sc
}

// SetBannerKk sets the "banner_kk" field.
func (sc *SwiperCreate) SetBannerKk(s string) *SwiperCreate {
	sc.mutation.SetBannerKk(s)
	return sc
}

// SetNillableBannerKk sets the "banner_kk" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableBannerKk(s *string) *SwiperCreate {
	if s != nil {
		sc.SetBannerKk(*s)
	}
	return sc
}

// SetContentZh sets the "content_zh" field.
func (sc *SwiperCreate) SetContentZh(s string) *SwiperCreate {
	sc.mutation.SetContentZh(s)
	return sc
}

// SetNillableContentZh sets the "content_zh" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableContentZh(s *string) *SwiperCreate {
	if s != nil {
		sc.SetContentZh(*s)
	}
	return sc
}

// SetContentEn sets the "content_en" field.
func (sc *SwiperCreate) SetContentEn(s string) *SwiperCreate {
	sc.mutation.SetContentEn(s)
	return sc
}

// SetNillableContentEn sets the "content_en" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableContentEn(s *string) *SwiperCreate {
	if s != nil {
		sc.SetContentEn(*s)
	}
	return sc
}

// SetContentRu sets the "content_ru" field.
func (sc *SwiperCreate) SetContentRu(s string) *SwiperCreate {
	sc.mutation.SetContentRu(s)
	return sc
}

// SetNillableContentRu sets the "content_ru" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableContentRu(s *string) *SwiperCreate {
	if s != nil {
		sc.SetContentRu(*s)
	}
	return sc
}

// SetContentKk sets the "content_kk" field.
func (sc *SwiperCreate) SetContentKk(s string) *SwiperCreate {
	sc.mutation.SetContentKk(s)
	return sc
}

// SetNillableContentKk sets the "content_kk" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableContentKk(s *string) *SwiperCreate {
	if s != nil {
		sc.SetContentKk(*s)
	}
	return sc
}

// SetJumpURL sets the "jump_url" field.
func (sc *SwiperCreate) SetJumpURL(s string) *SwiperCreate {
	sc.mutation.SetJumpURL(s)
	return sc
}

// SetNillableJumpURL sets the "jump_url" field if the given value is not nil.
func (sc *SwiperCreate) SetNillableJumpURL(s *string) *SwiperCreate {
	if s != nil {
		sc.SetJumpURL(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SwiperCreate) SetID(u uint64) *SwiperCreate {
	sc.mutation.SetID(u)
	return sc
}

// Mutation returns the SwiperMutation object of the builder.
func (sc *SwiperCreate) Mutation() *SwiperMutation {
	return sc.mutation
}

// Save creates the Swiper in the database.
func (sc *SwiperCreate) Save(ctx context.Context) (*Swiper, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SwiperCreate) SaveX(ctx context.Context) *Swiper {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SwiperCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SwiperCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SwiperCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := swiper.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := swiper.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.Status(); !ok {
		v := swiper.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.Sort(); !ok {
		v := swiper.DefaultSort
		sc.mutation.SetSort(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SwiperCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Swiper.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Swiper.updated_at"`)}
	}
	if _, ok := sc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Swiper.sort"`)}
	}
	return nil
}

func (sc *SwiperCreate) sqlSave(ctx context.Context) (*Swiper, error) {
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

func (sc *SwiperCreate) createSpec() (*Swiper, *sqlgraph.CreateSpec) {
	var (
		_node = &Swiper{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(swiper.Table, sqlgraph.NewFieldSpec(swiper.FieldID, field.TypeUint64))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(swiper.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(swiper.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(swiper.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.Sort(); ok {
		_spec.SetField(swiper.FieldSort, field.TypeUint32, value)
		_node.Sort = value
	}
	if value, ok := sc.mutation.TitleZh(); ok {
		_spec.SetField(swiper.FieldTitleZh, field.TypeString, value)
		_node.TitleZh = value
	}
	if value, ok := sc.mutation.TitleEn(); ok {
		_spec.SetField(swiper.FieldTitleEn, field.TypeString, value)
		_node.TitleEn = value
	}
	if value, ok := sc.mutation.TitleRu(); ok {
		_spec.SetField(swiper.FieldTitleRu, field.TypeString, value)
		_node.TitleRu = value
	}
	if value, ok := sc.mutation.TitleKk(); ok {
		_spec.SetField(swiper.FieldTitleKk, field.TypeString, value)
		_node.TitleKk = value
	}
	if value, ok := sc.mutation.BannerZh(); ok {
		_spec.SetField(swiper.FieldBannerZh, field.TypeString, value)
		_node.BannerZh = value
	}
	if value, ok := sc.mutation.BannerEn(); ok {
		_spec.SetField(swiper.FieldBannerEn, field.TypeString, value)
		_node.BannerEn = value
	}
	if value, ok := sc.mutation.BannerRu(); ok {
		_spec.SetField(swiper.FieldBannerRu, field.TypeString, value)
		_node.BannerRu = value
	}
	if value, ok := sc.mutation.BannerKk(); ok {
		_spec.SetField(swiper.FieldBannerKk, field.TypeString, value)
		_node.BannerKk = value
	}
	if value, ok := sc.mutation.ContentZh(); ok {
		_spec.SetField(swiper.FieldContentZh, field.TypeString, value)
		_node.ContentZh = value
	}
	if value, ok := sc.mutation.ContentEn(); ok {
		_spec.SetField(swiper.FieldContentEn, field.TypeString, value)
		_node.ContentEn = value
	}
	if value, ok := sc.mutation.ContentRu(); ok {
		_spec.SetField(swiper.FieldContentRu, field.TypeString, value)
		_node.ContentRu = value
	}
	if value, ok := sc.mutation.ContentKk(); ok {
		_spec.SetField(swiper.FieldContentKk, field.TypeString, value)
		_node.ContentKk = value
	}
	if value, ok := sc.mutation.JumpURL(); ok {
		_spec.SetField(swiper.FieldJumpURL, field.TypeString, value)
		_node.JumpURL = value
	}
	return _node, _spec
}

// SwiperCreateBulk is the builder for creating many Swiper entities in bulk.
type SwiperCreateBulk struct {
	config
	err      error
	builders []*SwiperCreate
}

// Save creates the Swiper entities in the database.
func (scb *SwiperCreateBulk) Save(ctx context.Context) ([]*Swiper, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Swiper, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SwiperMutation)
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
func (scb *SwiperCreateBulk) SaveX(ctx context.Context) []*Swiper {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SwiperCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SwiperCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
