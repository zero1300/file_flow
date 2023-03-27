// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"file_flow/ent/predicate"
	"file_flow/ent/share"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShareDelete is the builder for deleting a Share entity.
type ShareDelete struct {
	config
	hooks    []Hook
	mutation *ShareMutation
}

// Where appends a list predicates to the ShareDelete builder.
func (sd *ShareDelete) Where(ps ...predicate.Share) *ShareDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *ShareDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ShareMutation](ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *ShareDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *ShareDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(share.Table, sqlgraph.NewFieldSpec(share.FieldID, field.TypeInt))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// ShareDeleteOne is the builder for deleting a single Share entity.
type ShareDeleteOne struct {
	sd *ShareDelete
}

// Where appends a list predicates to the ShareDelete builder.
func (sdo *ShareDeleteOne) Where(ps ...predicate.Share) *ShareDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *ShareDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{share.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *ShareDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
