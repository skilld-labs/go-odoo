package odoo

import (
	"time"
)

type String struct {
	v string
}

func NewString(v string) *String {
	return &String{v: v}
}

func (s *String) Get() string {
	if s == nil {
		return ""
	}
	return s.v
}

type Int struct {
	v int64
}

func NewInt(v int64) *Int {
	return &Int{v: v}
}

func (i *Int) Get() int64 {
	if i == nil {
		return 0
	}
	return i.v
}

type Bool struct {
	v bool
}

func NewBool(v bool) *Bool {
	return &Bool{v: v}
}

func (b *Bool) Get() bool {
	if b == nil {
		return false
	}
	return b.v
}

type Selection struct {
	v interface{}
}

func NewSelection(v interface{}) *Selection {
	return &Selection{v: v}
}

func (s *Selection) Get() interface{} {
	if s == nil {
		return nil
	}
	return s.v
}

type Time struct {
	v time.Time
}

func NewTime(v time.Time) *Time {
	return &Time{v: v}
}

func (t *Time) Get() time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.v
}

type Float struct {
	v float64
}

func NewFloat(v float64) *Float {
	return &Float{v: v}
}

func (f *Float) Get() float64 {
	if f == nil {
		return 0
	}
	return f.v
}

type Many2One struct {
	ID   int64
	Name string
}

func NewMany2One(id int64, name string) *Many2One {
	return &Many2One{ID: id, Name: name}
}

func (m *Many2One) Get() int64 {
	return m.ID
}

type Relation struct {
	ids []int64
	v   []interface{}
}

func NewRelation() *Relation {
	return &Relation{}
}

func (r *Relation) Get() []int64 {
	if r == nil {
		return []int64{}
	}
	return r.ids
}

func (r *Relation) AddNewRecord(values interface{}) {
	r.v = append(r.v, newTuple(0, 0, getValuesFromInterface(values)))
}

func (r *Relation) UpdateRecord(record int64, values interface{}) {
	r.v = append(r.v, newTuple(1, record, getValuesFromInterface(values)))
}

func (r *Relation) DeleteRecord(record int64) {
	r.v = append(r.v, newTuple(2, record, 0))
}

func (r *Relation) RemoveRecord(record int64) {
	r.v = append(r.v, newTuple(3, record, 0))
}

func (r *Relation) AddRecord(record int) {
	r.v = append(r.v, newTuple(4, record, 0))
}

func (r *Relation) RemoveAllRecords() {
	r.v = append(r.v, newTuple(5, 0, 0))
}

func (r *Relation) ReplaceAllRecords(newRecords []int64) {
	r.v = append(r.v, newTuple(6, 0, newRecords))
}
