// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-member-rpc/ent/service"
)

// Service is the model entity for the Service schema.
type Service struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// Sort Number | 排序编号
	Sort uint32 `json:"sort,omitempty"`
	// Service chinese name | 服务中文名称
	NameZh string `json:"name_zh,omitempty"`
	// Service english name | 服务英文名称
	NameEn string `json:"name_en,omitempty"`
	// Service russian name | 服务俄语名称
	NameRu string `json:"name_ru,omitempty"`
	// Service kazakh name | 服务哈萨克语名称
	NameKk string `json:"name_kk,omitempty"`
	// Description chinese | 服务中文描述
	DescriptionZh string `json:"description_zh,omitempty"`
	// Description english | 服务英文描述
	DescriptionEn string `json:"description_en,omitempty"`
	// Description russian | 服务俄语描述
	DescriptionRu string `json:"description_ru,omitempty"`
	// Description kazakh | 服务哈萨克语描述
	DescriptionKk string `json:"description_kk,omitempty"`
	// Cover image URL | 封面图片URL
	CoverURL     string `json:"cover_url,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Service) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case service.FieldID, service.FieldStatus, service.FieldSort:
			values[i] = new(sql.NullInt64)
		case service.FieldNameZh, service.FieldNameEn, service.FieldNameRu, service.FieldNameKk, service.FieldDescriptionZh, service.FieldDescriptionEn, service.FieldDescriptionRu, service.FieldDescriptionKk, service.FieldCoverURL:
			values[i] = new(sql.NullString)
		case service.FieldCreatedAt, service.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Service fields.
func (s *Service) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case service.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint64(value.Int64)
		case service.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case service.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case service.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = uint8(value.Int64)
			}
		case service.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				s.Sort = uint32(value.Int64)
			}
		case service.FieldNameZh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_zh", values[i])
			} else if value.Valid {
				s.NameZh = value.String
			}
		case service.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				s.NameEn = value.String
			}
		case service.FieldNameRu:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_ru", values[i])
			} else if value.Valid {
				s.NameRu = value.String
			}
		case service.FieldNameKk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_kk", values[i])
			} else if value.Valid {
				s.NameKk = value.String
			}
		case service.FieldDescriptionZh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description_zh", values[i])
			} else if value.Valid {
				s.DescriptionZh = value.String
			}
		case service.FieldDescriptionEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description_en", values[i])
			} else if value.Valid {
				s.DescriptionEn = value.String
			}
		case service.FieldDescriptionRu:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description_ru", values[i])
			} else if value.Valid {
				s.DescriptionRu = value.String
			}
		case service.FieldDescriptionKk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description_kk", values[i])
			} else if value.Valid {
				s.DescriptionKk = value.String
			}
		case service.FieldCoverURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cover_url", values[i])
			} else if value.Valid {
				s.CoverURL = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Service.
// This includes values selected through modifiers, order, etc.
func (s *Service) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Service.
// Note that you need to call Service.Unwrap() before calling this method if this Service
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Service) Update() *ServiceUpdateOne {
	return NewServiceClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Service entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Service) Unwrap() *Service {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Service is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Service) String() string {
	var builder strings.Builder
	builder.WriteString("Service(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", s.Sort))
	builder.WriteString(", ")
	builder.WriteString("name_zh=")
	builder.WriteString(s.NameZh)
	builder.WriteString(", ")
	builder.WriteString("name_en=")
	builder.WriteString(s.NameEn)
	builder.WriteString(", ")
	builder.WriteString("name_ru=")
	builder.WriteString(s.NameRu)
	builder.WriteString(", ")
	builder.WriteString("name_kk=")
	builder.WriteString(s.NameKk)
	builder.WriteString(", ")
	builder.WriteString("description_zh=")
	builder.WriteString(s.DescriptionZh)
	builder.WriteString(", ")
	builder.WriteString("description_en=")
	builder.WriteString(s.DescriptionEn)
	builder.WriteString(", ")
	builder.WriteString("description_ru=")
	builder.WriteString(s.DescriptionRu)
	builder.WriteString(", ")
	builder.WriteString("description_kk=")
	builder.WriteString(s.DescriptionKk)
	builder.WriteString(", ")
	builder.WriteString("cover_url=")
	builder.WriteString(s.CoverURL)
	builder.WriteByte(')')
	return builder.String()
}

// Services is a parsable slice of Service.
type Services []*Service
