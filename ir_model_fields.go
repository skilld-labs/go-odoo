package odoo

import (
	"errors"
	"fmt"
)

type IrModelFields struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omptempty"`
	Column1              *String    `xmlrpc:"column1,omptempty"`
	Column2              *String    `xmlrpc:"column2,omptempty"`
	CompleteName         *String    `xmlrpc:"complete_name,omptempty"`
	Compute              *String    `xmlrpc:"compute,omptempty"`
	Copy                 *Bool      `xmlrpc:"copy,omptempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omptempty"`
	Depends              *String    `xmlrpc:"depends,omptempty"`
	DisplayName          *String    `xmlrpc:"display_name,omptempty"`
	Domain               *String    `xmlrpc:"domain,omptempty"`
	FieldDescription     *String    `xmlrpc:"field_description,omptempty"`
	Groups               *Relation  `xmlrpc:"groups,omptempty"`
	Help                 *String    `xmlrpc:"help,omptempty"`
	Id                   *Int       `xmlrpc:"id,omptempty"`
	Index                *Bool      `xmlrpc:"index,omptempty"`
	Model                *String    `xmlrpc:"model,omptempty"`
	ModelId              *Many2One  `xmlrpc:"model_id,omptempty"`
	Modules              *String    `xmlrpc:"modules,omptempty"`
	Name                 *String    `xmlrpc:"name,omptempty"`
	OnDelete             *Selection `xmlrpc:"on_delete,omptempty"`
	Readonly             *Bool      `xmlrpc:"readonly,omptempty"`
	Related              *String    `xmlrpc:"related,omptempty"`
	Relation             *String    `xmlrpc:"relation,omptempty"`
	RelationField        *String    `xmlrpc:"relation_field,omptempty"`
	RelationTable        *String    `xmlrpc:"relation_table,omptempty"`
	Required             *Bool      `xmlrpc:"required,omptempty"`
	Selectable           *Bool      `xmlrpc:"selectable,omptempty"`
	Selection            *String    `xmlrpc:"selection,omptempty"`
	SerializationFieldId *Many2One  `xmlrpc:"serialization_field_id,omptempty"`
	Size                 *Int       `xmlrpc:"size,omptempty"`
	State                *Selection `xmlrpc:"state,omptempty"`
	Store                *Bool      `xmlrpc:"store,omptempty"`
	TrackVisibility      *Selection `xmlrpc:"track_visibility,omptempty"`
	Translate            *Bool      `xmlrpc:"translate,omptempty"`
	Ttype                *Selection `xmlrpc:"ttype,omptempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omptempty"`
}

type IrModelFieldss []IrModelFields

const IrModelFieldsModel = "ir.model.fields"

func (c *Client) CreateIrModelFields(imf *IrModelFields) (int64, error) {
	return c.Create(IrModelFieldsModel, imf)
}

func (c *Client) UpdateIrModelFields(imf *IrModelFields) error {
	return c.UpdateIrModelFieldss([]int64{imf.Id.Get()}, imf)
}

func (c *Client) UpdateIrModelFieldss(ids []int64, imf *IrModelFields) error {
	return c.Update(IrModelFieldsModel, ids, imf)
}

func (c *Client) DeleteIrModelFields(id int64) error {
	return c.DeleteIrModelFieldss([]int64{id})
}

func (c *Client) DeleteIrModelFieldss(ids []int64) error {
	return c.Delete(IrModelFieldsModel, ids)
}

func (c *Client) GetIrModelFields(id int64) (*IrModelFields, error) {
	imfs, err := c.GetIrModelFieldss([]int64{id})
	if err != nil {
		return nil, err
	}
	if imfs != nil && len(*imfs) > 0 {
		return &((*imfs)[0]), nil
	}
	return nil, errors.New(fmt.Sprintf("id %v of %s not found", id, IrModelFieldsModel))
}

func (c *Client) GetIrModelFieldss(ids []int64) (*IrModelFieldss, error) {
	imfs := &IrModelFieldss{}
	if err := c.Read(IrModelFieldsModel, ids, nil, imfs); err != nil {
		return nil, err
	}
	return imfs, nil
}

func (c *Client) FindIrModelFieldss(criteria *Criteria, options *Options) (*IrModelFieldss, error) {
	imfs := &IrModelFieldss{}
	if err := c.SearchRead(IrModelFieldsModel, criteria, options, imfs); err != nil {
		return nil, err
	}
	return imfs, nil
}
