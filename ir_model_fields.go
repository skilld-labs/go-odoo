package odoo

import (
	"fmt"
)

// IrModelFields represents ir.model.fields model.
type IrModelFields struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omitempty"`
	Column1          *String    `xmlrpc:"column1,omitempty"`
	Column2          *String    `xmlrpc:"column2,omitempty"`
	CompleteName     *String    `xmlrpc:"complete_name,omitempty"`
	Compute          *String    `xmlrpc:"compute,omitempty"`
	Copied           *Bool      `xmlrpc:"copied,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	Depends          *String    `xmlrpc:"depends,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	Domain           *String    `xmlrpc:"domain,omitempty"`
	FieldDescription *String    `xmlrpc:"field_description,omitempty"`
	GroupExpand      *Bool      `xmlrpc:"group_expand,omitempty"`
	Groups           *Relation  `xmlrpc:"groups,omitempty"`
	Help             *String    `xmlrpc:"help,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	Index            *Bool      `xmlrpc:"index,omitempty"`
	Model            *String    `xmlrpc:"model,omitempty"`
	ModelId          *Many2One  `xmlrpc:"model_id,omitempty"`
	Modules          *String    `xmlrpc:"modules,omitempty"`
	Name             *String    `xmlrpc:"name,omitempty"`
	OnDelete         *Selection `xmlrpc:"on_delete,omitempty"`
	Readonly         *Bool      `xmlrpc:"readonly,omitempty"`
	Related          *String    `xmlrpc:"related,omitempty"`
	RelatedFieldId   *Many2One  `xmlrpc:"related_field_id,omitempty"`
	Relation         *String    `xmlrpc:"relation,omitempty"`
	RelationField    *String    `xmlrpc:"relation_field,omitempty"`
	RelationFieldId  *Many2One  `xmlrpc:"relation_field_id,omitempty"`
	RelationTable    *String    `xmlrpc:"relation_table,omitempty"`
	Required         *Bool      `xmlrpc:"required,omitempty"`
	Selectable       *Bool      `xmlrpc:"selectable,omitempty"`
	Selection        *String    `xmlrpc:"selection,omitempty"`
	SelectionIds     *Relation  `xmlrpc:"selection_ids,omitempty"`
	Size             *Int       `xmlrpc:"size,omitempty"`
	State            *Selection `xmlrpc:"state,omitempty"`
	Store            *Bool      `xmlrpc:"store,omitempty"`
	Tracking         *Int       `xmlrpc:"tracking,omitempty"`
	Translate        *Bool      `xmlrpc:"translate,omitempty"`
	Ttype            *Selection `xmlrpc:"ttype,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// IrModelFieldss represents array of ir.model.fields model.
type IrModelFieldss []IrModelFields

// IrModelFieldsModel is the odoo model name.
const IrModelFieldsModel = "ir.model.fields"

// Many2One convert IrModelFields to *Many2One.
func (imf *IrModelFields) Many2One() *Many2One {
	return NewMany2One(imf.Id.Get(), "")
}

// CreateIrModelFields creates a new ir.model.fields model and returns its id.
func (c *Client) CreateIrModelFields(imf *IrModelFields) (int64, error) {
	return c.Create(IrModelFieldsModel, imf)
}

// UpdateIrModelFields updates an existing ir.model.fields record.
func (c *Client) UpdateIrModelFields(imf *IrModelFields) error {
	return c.UpdateIrModelFieldss([]int64{imf.Id.Get()}, imf)
}

// UpdateIrModelFieldss updates existing ir.model.fields records.
// All records (represented by ids) will be updated by imf values.
func (c *Client) UpdateIrModelFieldss(ids []int64, imf *IrModelFields) error {
	return c.Update(IrModelFieldsModel, ids, imf)
}

// DeleteIrModelFields deletes an existing ir.model.fields record.
func (c *Client) DeleteIrModelFields(id int64) error {
	return c.DeleteIrModelFieldss([]int64{id})
}

// DeleteIrModelFieldss deletes existing ir.model.fields records.
func (c *Client) DeleteIrModelFieldss(ids []int64) error {
	return c.Delete(IrModelFieldsModel, ids)
}

// GetIrModelFields gets ir.model.fields existing record.
func (c *Client) GetIrModelFields(id int64) (*IrModelFields, error) {
	imfs, err := c.GetIrModelFieldss([]int64{id})
	if err != nil {
		return nil, err
	}
	if imfs != nil && len(*imfs) > 0 {
		return &((*imfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.model.fields not found", id)
}

// GetIrModelFieldss gets ir.model.fields existing records.
func (c *Client) GetIrModelFieldss(ids []int64) (*IrModelFieldss, error) {
	imfs := &IrModelFieldss{}
	if err := c.Read(IrModelFieldsModel, ids, nil, imfs); err != nil {
		return nil, err
	}
	return imfs, nil
}

// FindIrModelFields finds ir.model.fields record by querying it with criteria.
func (c *Client) FindIrModelFields(criteria *Criteria) (*IrModelFields, error) {
	imfs := &IrModelFieldss{}
	if err := c.SearchRead(IrModelFieldsModel, criteria, NewOptions().Limit(1), imfs); err != nil {
		return nil, err
	}
	if imfs != nil && len(*imfs) > 0 {
		return &((*imfs)[0]), nil
	}
	return nil, fmt.Errorf("ir.model.fields was not found")
}

// FindIrModelFieldss finds ir.model.fields records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelFieldss(criteria *Criteria, options *Options) (*IrModelFieldss, error) {
	imfs := &IrModelFieldss{}
	if err := c.SearchRead(IrModelFieldsModel, criteria, options, imfs); err != nil {
		return nil, err
	}
	return imfs, nil
}

// FindIrModelFieldsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelFieldsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrModelFieldsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrModelFieldsId finds record id by querying it with criteria.
func (c *Client) FindIrModelFieldsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModelFieldsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.model.fields was not found")
}
