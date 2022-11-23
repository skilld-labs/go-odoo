package odoo

import (
	"time"
)

// String is a string wrapper
type String struct {
	V string
}

// NewString creates a new *String.
func NewString(v string) *String {
	return &String{V: v}
}

// Get *String value.
func (s *String) Get() string {
	if s == nil {
		return ""
	}
	return s.V
}

// Int is an int64 wrapper
type Int struct {
	V int64
}

// NewInt creates a new *Int.
func NewInt(v int64) *Int {
	return &Int{V: v}
}

// Get *Int value.
func (i *Int) Get() int64 {
	if i == nil {
		return 0
	}
	return i.V
}

// Bool is a bool wrapper
type Bool struct {
	V bool
}

// NewBool creates a new *Bool.
func NewBool(v bool) *Bool {
	return &Bool{V: v}
}

// Get *Bool value.
func (b *Bool) Get() bool {
	if b == nil {
		return false
	}
	return b.V
}

// Selection represents selection odoo type.
type Selection struct {
	V interface{}
}

// NewSelection creates a new *Selection.
func NewSelection(v interface{}) *Selection {
	return &Selection{V: v}
}

// Get *Selection value.
func (s *Selection) Get() interface{} {
	if s == nil {
		return nil
	}
	return s.V
}

// Time is a time.Time wrapper.
type Time struct {
	V time.Time
}

// NewTime creates a new *Time.
func NewTime(v time.Time) *Time {
	return &Time{V: v}
}

// Get *Time value.
func (t *Time) Get() time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.V
}

// Float is a float64 wrapper
type Float struct {
	V float64
}

// NewFloat creates a new *Float.
func NewFloat(v float64) *Float {
	return &Float{V: v}
}

// Get *Float value.
func (f *Float) Get() float64 {
	if f == nil {
		return 0
	}
	return f.V
}

// Many2One represents odoo many2one type.
// https://www.odoo.com/documentation/13.0/reference/orm.html#relational-fields
type Many2One struct {
	ID   int64
	Name string
}

// NewMany2One create a new *Many2One.
func NewMany2One(id int64, name string) *Many2One {
	return &Many2One{ID: id, Name: name}
}

// Get *Many2One value.
func (m *Many2One) Get() int64 {
	return m.ID
}

// Relation represents odoo one2many and many2many types.
// https://www.odoo.com/documentation/13.0/reference/orm.html#relational-fields
type Relation struct {
	IDs []int64
	V   []interface{}
}

// NewRelation creates a new *Relation.
func NewRelation() *Relation {
	return &Relation{}
}

// Get *Relation value.
func (r *Relation) Get() []int64 {
	if r == nil {
		return []int64{}
	}
	return r.IDs
}

// AddNewRecord is an helper to create a new record of one2many or many2many.
// https://www.odoo.com/documentation/13.0/reference/orm.html#odoo.models.Model.write
func (r *Relation) AddNewRecord(values interface{}) {
	r.V = append(r.V, newTuple(0, 0, getValuesFromInterface(values)))
}

// UpdateRecord is an helper to update an existing record of one2many or many2many.
// https://www.odoo.com/documentation/13.0/reference/orm.html#odoo.models.Model.write
func (r *Relation) UpdateRecord(record int64, values interface{}) {
	r.V = append(r.V, newTuple(1, record, getValuesFromInterface(values)))
}

// DeleteRecord is an helper to delete an existing record of one2many or many2many.
// https://www.odoo.com/documentation/13.0/reference/orm.html#odoo.models.Model.write
func (r *Relation) DeleteRecord(record int64) {
	r.V = append(r.V, newTuple(2, record, 0))
}

// RemoveRecord is an helper to remove an existing record of one2many or many2many.
// https://www.odoo.com/documentation/13.0/reference/orm.html#odoo.models.Model.write
func (r *Relation) RemoveRecord(record int64) {
	r.V = append(r.V, newTuple(3, record, 0))
}

// AddRecord is an helper to add an existing record of one2many or many2many.
// https://www.odoo.com/documentation/13.0/reference/orm.html#odoo.models.Model.write
func (r *Relation) AddRecord(record int64) {
	r.V = append(r.V, newTuple(4, record, 0))
}

// RemoveAllRecords is an helper to remove all records of one2many or many2many.
// https://www.odoo.com/documentation/13.0/reference/orm.html#odoo.models.Model.write
func (r *Relation) RemoveAllRecords() {
	r.V = append(r.V, newTuple(5, 0, 0))
}

// ReplaceAllRecords is an helper to replace all records of one2many or many2many.
// https://www.odoo.com/documentation/13.0/reference/orm.html#odoo.models.Model.write
func (r *Relation) ReplaceAllRecords(newRecords []int64) {
	r.V = append(r.V, newTuple(6, 0, newRecords))
}
