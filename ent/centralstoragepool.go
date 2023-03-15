// Code generated by ent, DO NOT EDIT.

package ent

import (
	"file_flow/ent/centralstoragepool"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// CentralStoragePool is the model entity for the CentralStoragePool schema.
type CentralStoragePool struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt *time.Time `json:"delete_at,omitempty"`
	// 文件名
	Filename string `json:"filename,omitempty"`
	// 文件扩展名
	Ext string `json:"ext,omitempty"`
	// 文件大小
	Size float64 `json:"size,omitempty"`
	// 文件路径
	Path string `json:"path,omitempty"`
	// 文件哈希
	Hash string `json:"hash,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CentralStoragePool) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case centralstoragepool.FieldSize:
			values[i] = new(sql.NullFloat64)
		case centralstoragepool.FieldID:
			values[i] = new(sql.NullInt64)
		case centralstoragepool.FieldFilename, centralstoragepool.FieldExt, centralstoragepool.FieldPath, centralstoragepool.FieldHash:
			values[i] = new(sql.NullString)
		case centralstoragepool.FieldDeleteAt, centralstoragepool.FieldCreateAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CentralStoragePool", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CentralStoragePool fields.
func (csp *CentralStoragePool) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case centralstoragepool.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			csp.ID = int(value.Int64)
		case centralstoragepool.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				csp.DeleteAt = new(time.Time)
				*csp.DeleteAt = value.Time
			}
		case centralstoragepool.FieldFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field filename", values[i])
			} else if value.Valid {
				csp.Filename = value.String
			}
		case centralstoragepool.FieldExt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ext", values[i])
			} else if value.Valid {
				csp.Ext = value.String
			}
		case centralstoragepool.FieldSize:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				csp.Size = value.Float64
			}
		case centralstoragepool.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				csp.Path = value.String
			}
		case centralstoragepool.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				csp.Hash = value.String
			}
		case centralstoragepool.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				csp.CreateAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CentralStoragePool.
// Note that you need to call CentralStoragePool.Unwrap() before calling this method if this CentralStoragePool
// was returned from a transaction, and the transaction was committed or rolled back.
func (csp *CentralStoragePool) Update() *CentralStoragePoolUpdateOne {
	return NewCentralStoragePoolClient(csp.config).UpdateOne(csp)
}

// Unwrap unwraps the CentralStoragePool entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (csp *CentralStoragePool) Unwrap() *CentralStoragePool {
	_tx, ok := csp.config.driver.(*txDriver)
	if !ok {
		panic("ent: CentralStoragePool is not a transactional entity")
	}
	csp.config.driver = _tx.drv
	return csp
}

// String implements the fmt.Stringer.
func (csp *CentralStoragePool) String() string {
	var builder strings.Builder
	builder.WriteString("CentralStoragePool(")
	builder.WriteString(fmt.Sprintf("id=%v, ", csp.ID))
	if v := csp.DeleteAt; v != nil {
		builder.WriteString("delete_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("filename=")
	builder.WriteString(csp.Filename)
	builder.WriteString(", ")
	builder.WriteString("ext=")
	builder.WriteString(csp.Ext)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", csp.Size))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(csp.Path)
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(csp.Hash)
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(csp.CreateAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CentralStoragePools is a parsable slice of CentralStoragePool.
type CentralStoragePools []*CentralStoragePool
