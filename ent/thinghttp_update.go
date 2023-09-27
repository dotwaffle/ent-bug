// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/thing"
	"entgo.io/bug/ent/thinghttp"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ThingHTTPUpdate is the builder for updating ThingHTTP entities.
type ThingHTTPUpdate struct {
	config
	hooks    []Hook
	mutation *ThingHTTPMutation
}

// Where appends a list predicates to the ThingHTTPUpdate builder.
func (thu *ThingHTTPUpdate) Where(ps ...predicate.ThingHTTP) *ThingHTTPUpdate {
	thu.mutation.Where(ps...)
	return thu
}

// SetAge sets the "age" field.
func (thu *ThingHTTPUpdate) SetAge(i int) *ThingHTTPUpdate {
	thu.mutation.ResetAge()
	thu.mutation.SetAge(i)
	return thu
}

// AddAge adds i to the "age" field.
func (thu *ThingHTTPUpdate) AddAge(i int) *ThingHTTPUpdate {
	thu.mutation.AddAge(i)
	return thu
}

// SetName sets the "name" field.
func (thu *ThingHTTPUpdate) SetName(s string) *ThingHTTPUpdate {
	thu.mutation.SetName(s)
	return thu
}

// AddProbesHTTPIDs adds the "probes_http" edge to the Thing entity by IDs.
func (thu *ThingHTTPUpdate) AddProbesHTTPIDs(ids ...int) *ThingHTTPUpdate {
	thu.mutation.AddProbesHTTPIDs(ids...)
	return thu
}

// AddProbesHTTP adds the "probes_http" edges to the Thing entity.
func (thu *ThingHTTPUpdate) AddProbesHTTP(t ...*Thing) *ThingHTTPUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return thu.AddProbesHTTPIDs(ids...)
}

// Mutation returns the ThingHTTPMutation object of the builder.
func (thu *ThingHTTPUpdate) Mutation() *ThingHTTPMutation {
	return thu.mutation
}

// ClearProbesHTTP clears all "probes_http" edges to the Thing entity.
func (thu *ThingHTTPUpdate) ClearProbesHTTP() *ThingHTTPUpdate {
	thu.mutation.ClearProbesHTTP()
	return thu
}

// RemoveProbesHTTPIDs removes the "probes_http" edge to Thing entities by IDs.
func (thu *ThingHTTPUpdate) RemoveProbesHTTPIDs(ids ...int) *ThingHTTPUpdate {
	thu.mutation.RemoveProbesHTTPIDs(ids...)
	return thu
}

// RemoveProbesHTTP removes "probes_http" edges to Thing entities.
func (thu *ThingHTTPUpdate) RemoveProbesHTTP(t ...*Thing) *ThingHTTPUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return thu.RemoveProbesHTTPIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (thu *ThingHTTPUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, thu.sqlSave, thu.mutation, thu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (thu *ThingHTTPUpdate) SaveX(ctx context.Context) int {
	affected, err := thu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (thu *ThingHTTPUpdate) Exec(ctx context.Context) error {
	_, err := thu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thu *ThingHTTPUpdate) ExecX(ctx context.Context) {
	if err := thu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (thu *ThingHTTPUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(thinghttp.Table, thinghttp.Columns, sqlgraph.NewFieldSpec(thinghttp.FieldID, field.TypeInt))
	if ps := thu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := thu.mutation.Age(); ok {
		_spec.SetField(thinghttp.FieldAge, field.TypeInt, value)
	}
	if value, ok := thu.mutation.AddedAge(); ok {
		_spec.AddField(thinghttp.FieldAge, field.TypeInt, value)
	}
	if value, ok := thu.mutation.Name(); ok {
		_spec.SetField(thinghttp.FieldName, field.TypeString, value)
	}
	if thu.mutation.ProbesHTTPCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   thinghttp.ProbesHTTPTable,
			Columns: thinghttp.ProbesHTTPPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thing.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := thu.mutation.RemovedProbesHTTPIDs(); len(nodes) > 0 && !thu.mutation.ProbesHTTPCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   thinghttp.ProbesHTTPTable,
			Columns: thinghttp.ProbesHTTPPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thing.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := thu.mutation.ProbesHTTPIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   thinghttp.ProbesHTTPTable,
			Columns: thinghttp.ProbesHTTPPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thing.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, thu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thinghttp.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	thu.mutation.done = true
	return n, nil
}

// ThingHTTPUpdateOne is the builder for updating a single ThingHTTP entity.
type ThingHTTPUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ThingHTTPMutation
}

// SetAge sets the "age" field.
func (thuo *ThingHTTPUpdateOne) SetAge(i int) *ThingHTTPUpdateOne {
	thuo.mutation.ResetAge()
	thuo.mutation.SetAge(i)
	return thuo
}

// AddAge adds i to the "age" field.
func (thuo *ThingHTTPUpdateOne) AddAge(i int) *ThingHTTPUpdateOne {
	thuo.mutation.AddAge(i)
	return thuo
}

// SetName sets the "name" field.
func (thuo *ThingHTTPUpdateOne) SetName(s string) *ThingHTTPUpdateOne {
	thuo.mutation.SetName(s)
	return thuo
}

// AddProbesHTTPIDs adds the "probes_http" edge to the Thing entity by IDs.
func (thuo *ThingHTTPUpdateOne) AddProbesHTTPIDs(ids ...int) *ThingHTTPUpdateOne {
	thuo.mutation.AddProbesHTTPIDs(ids...)
	return thuo
}

// AddProbesHTTP adds the "probes_http" edges to the Thing entity.
func (thuo *ThingHTTPUpdateOne) AddProbesHTTP(t ...*Thing) *ThingHTTPUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return thuo.AddProbesHTTPIDs(ids...)
}

// Mutation returns the ThingHTTPMutation object of the builder.
func (thuo *ThingHTTPUpdateOne) Mutation() *ThingHTTPMutation {
	return thuo.mutation
}

// ClearProbesHTTP clears all "probes_http" edges to the Thing entity.
func (thuo *ThingHTTPUpdateOne) ClearProbesHTTP() *ThingHTTPUpdateOne {
	thuo.mutation.ClearProbesHTTP()
	return thuo
}

// RemoveProbesHTTPIDs removes the "probes_http" edge to Thing entities by IDs.
func (thuo *ThingHTTPUpdateOne) RemoveProbesHTTPIDs(ids ...int) *ThingHTTPUpdateOne {
	thuo.mutation.RemoveProbesHTTPIDs(ids...)
	return thuo
}

// RemoveProbesHTTP removes "probes_http" edges to Thing entities.
func (thuo *ThingHTTPUpdateOne) RemoveProbesHTTP(t ...*Thing) *ThingHTTPUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return thuo.RemoveProbesHTTPIDs(ids...)
}

// Where appends a list predicates to the ThingHTTPUpdate builder.
func (thuo *ThingHTTPUpdateOne) Where(ps ...predicate.ThingHTTP) *ThingHTTPUpdateOne {
	thuo.mutation.Where(ps...)
	return thuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (thuo *ThingHTTPUpdateOne) Select(field string, fields ...string) *ThingHTTPUpdateOne {
	thuo.fields = append([]string{field}, fields...)
	return thuo
}

// Save executes the query and returns the updated ThingHTTP entity.
func (thuo *ThingHTTPUpdateOne) Save(ctx context.Context) (*ThingHTTP, error) {
	return withHooks(ctx, thuo.sqlSave, thuo.mutation, thuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (thuo *ThingHTTPUpdateOne) SaveX(ctx context.Context) *ThingHTTP {
	node, err := thuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (thuo *ThingHTTPUpdateOne) Exec(ctx context.Context) error {
	_, err := thuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thuo *ThingHTTPUpdateOne) ExecX(ctx context.Context) {
	if err := thuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (thuo *ThingHTTPUpdateOne) sqlSave(ctx context.Context) (_node *ThingHTTP, err error) {
	_spec := sqlgraph.NewUpdateSpec(thinghttp.Table, thinghttp.Columns, sqlgraph.NewFieldSpec(thinghttp.FieldID, field.TypeInt))
	id, ok := thuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ThingHTTP.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := thuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, thinghttp.FieldID)
		for _, f := range fields {
			if !thinghttp.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != thinghttp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := thuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := thuo.mutation.Age(); ok {
		_spec.SetField(thinghttp.FieldAge, field.TypeInt, value)
	}
	if value, ok := thuo.mutation.AddedAge(); ok {
		_spec.AddField(thinghttp.FieldAge, field.TypeInt, value)
	}
	if value, ok := thuo.mutation.Name(); ok {
		_spec.SetField(thinghttp.FieldName, field.TypeString, value)
	}
	if thuo.mutation.ProbesHTTPCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   thinghttp.ProbesHTTPTable,
			Columns: thinghttp.ProbesHTTPPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thing.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := thuo.mutation.RemovedProbesHTTPIDs(); len(nodes) > 0 && !thuo.mutation.ProbesHTTPCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   thinghttp.ProbesHTTPTable,
			Columns: thinghttp.ProbesHTTPPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thing.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := thuo.mutation.ProbesHTTPIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   thinghttp.ProbesHTTPTable,
			Columns: thinghttp.ProbesHTTPPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thing.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ThingHTTP{config: thuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, thuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thinghttp.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	thuo.mutation.done = true
	return _node, nil
}