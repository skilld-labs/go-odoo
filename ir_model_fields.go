package odoo

import (
	"fmt"
)

// IrModelFields represents ir.model.fields model
type IrModelFields struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omptempty"`
	Column1          *String    `xmlrpc:"column1,omptempty"`
	Column2          *String    `xmlrpc:"column2,omptempty"`
	CompleteName     *String    `xmlrpc:"complete_name,omptempty"`
	Compute          *String    `xmlrpc:"compute,omptempty"`
	Copied           *Bool      `xmlrpc:"copied,omptempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omptempty"`
	Depends          *String    `xmlrpc:"depends,omptempty"`
	DisplayName      *String    `xmlrpc:"display_name,omptempty"`
	Domain           *String    `xmlrpc:"domain,omptempty"`
	FieldDescription *String    `xmlrpc:"field_description,omptempty"`
	Groups           *Relation  `xmlrpc:"groups,omptempty"`
	Help             *String    `xmlrpc:"help,omptempty"`
	Id               *Int       `xmlrpc:"id,omptempty"`
	Index            *Bool      `xmlrpc:"index,omptempty"`
	Model            *String    `xmlrpc:"model,omptempty"`
	ModelId          *Many2One  `xmlrpc:"model_id,omptempty"`
	Modules          *String    `xmlrpc:"modules,omptempty"`
	Name             *String    `xmlrpc:"name,omptempty"`
	OnDelete         *Selection `xmlrpc:"on_delete,omptempty"`
	Readonly         *Bool      `xmlrpc:"readonly,omptempty"`
	Related          *String    `xmlrpc:"related,omptempty"`
	RelatedFieldId   *Many2One  `xmlrpc:"related_field_id,omptempty"`
	Relation         *String    `xmlrpc:"relation,omptempty"`
	RelationField    *String    `xmlrpc:"relation_field,omptempty"`
	RelationFieldId  *Many2One  `xmlrpc:"relation_field_id,omptempty"`
	RelationTable    *String    `xmlrpc:"relation_table,omptempty"`
	Required         *Bool      `xmlrpc:"required,omptempty"`
	Selectable       *Bool      `xmlrpc:"selectable,omptempty"`
	Selection        *String    `xmlrpc:"selection,omptempty"`
	SelectionIds     *Relation  `xmlrpc:"selection_ids,omptempty"`
	Size             *Int       `xmlrpc:"size,omptempty"`
	State            *Selection `xmlrpc:"state,omptempty"`
	Store            *Bool      `xmlrpc:"store,omptempty"`
	Translate        *Bool      `xmlrpc:"translate,omptempty"`
	Ttype            *Selection `xmlrpc:"ttype,omptempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// IrModelFieldss represents array of ir.model.fields model
type IrModelFieldss []IrModelFields

// IrModelFieldsModel is the odoo model name
const IrModelFieldsModel = "ir.model.fields"

// CreateIrModelFields creates a new ir.model.fields model and returns its id.
func (c *Client) CreateIrModelFields(imf *IrModelFields) (int64, error) {
	return c.Create(IrModelFieldsModel, imf)
}

// UpdateIrModelFields pdates an existing ir.model.fields record.
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
	return nil, fmt.Errorf("id %v of %s not found", id, IrModelFieldsModel)
}

// GetIrModelFieldss gets ir.model.fields existing records.
func (c *Client) GetIrModelFieldss(ids []int64) (*IrModelFieldss, error) {
	imfs := &IrModelFieldss{}
	if err := c.Read(IrModelFieldsModel, ids, nil, imfs); err != nil {
		return nil, err
	}
	return imfs, nil
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
