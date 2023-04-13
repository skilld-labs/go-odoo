package odoo

import (
	"fmt"
)

// IrAttachment represents ir.attachment model.
type IrAttachment struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken  *String    `xmlrpc:"access_token,omptempty"`
	Checksum     *String    `xmlrpc:"checksum,omptempty"`
	CompanyId    *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omptempty"`
	Datas        *String    `xmlrpc:"datas,omptempty"`
	DatasFname   *String    `xmlrpc:"datas_fname,omptempty"`
	DbDatas      *String    `xmlrpc:"db_datas,omptempty"`
	Description  *String    `xmlrpc:"description,omptempty"`
	DisplayName  *String    `xmlrpc:"display_name,omptempty"`
	FileSize     *Int       `xmlrpc:"file_size,omptempty"`
	Id           *Int       `xmlrpc:"id,omptempty"`
	IndexContent *String    `xmlrpc:"index_content,omptempty"`
	LocalUrl     *String    `xmlrpc:"local_url,omptempty"`
	Mimetype     *String    `xmlrpc:"mimetype,omptempty"`
	Name         *String    `xmlrpc:"name,omptempty"`
	Public       *Bool      `xmlrpc:"public,omptempty"`
	ResField     *String    `xmlrpc:"res_field,omptempty"`
	ResId        *Int       `xmlrpc:"res_id,omptempty"`
	ResModel     *String    `xmlrpc:"res_model,omptempty"`
	ResName      *String    `xmlrpc:"res_name,omptempty"`
	StoreFname   *String    `xmlrpc:"store_fname,omptempty"`
	Type         *Selection `xmlrpc:"type,omptempty"`
	Url          *String    `xmlrpc:"url,omptempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// IrAttachments represents array of ir.attachment model.
type IrAttachments []IrAttachment

// IrAttachmentModel is the odoo model name.
const IrAttachmentModel = "ir.attachment"

// Many2One convert IrAttachment to *Many2One.
func (ia *IrAttachment) Many2One() *Many2One {
	return NewMany2One(ia.Id.Get(), "")
}

// CreateIrAttachment creates a new ir.attachment model and returns its id.
func (c *Client) CreateIrAttachment(ia *IrAttachment) (int64, error) {
	ids, err := c.Create(IrAttachmentModel, []interface{}{ia})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrAttachment creates a new ir.attachment model and returns its id.
func (c *Client) CreateIrAttachments(ias []*IrAttachment) ([]int64, error) {
	var vv []interface{}
	for _, v := range ias {
		vv = append(vv, v)
	}
	return c.Create(IrAttachmentModel, vv)
}

// UpdateIrAttachment updates an existing ir.attachment record.
func (c *Client) UpdateIrAttachment(ia *IrAttachment) error {
	return c.UpdateIrAttachments([]int64{ia.Id.Get()}, ia)
}

// UpdateIrAttachments updates existing ir.attachment records.
// All records (represented by ids) will be updated by ia values.
func (c *Client) UpdateIrAttachments(ids []int64, ia *IrAttachment) error {
	return c.Update(IrAttachmentModel, ids, ia)
}

// DeleteIrAttachment deletes an existing ir.attachment record.
func (c *Client) DeleteIrAttachment(id int64) error {
	return c.DeleteIrAttachments([]int64{id})
}

// DeleteIrAttachments deletes existing ir.attachment records.
func (c *Client) DeleteIrAttachments(ids []int64) error {
	return c.Delete(IrAttachmentModel, ids)
}

// GetIrAttachment gets ir.attachment existing record.
func (c *Client) GetIrAttachment(id int64) (*IrAttachment, error) {
	ias, err := c.GetIrAttachments([]int64{id})
	if err != nil {
		return nil, err
	}
	if ias != nil && len(*ias) > 0 {
		return &((*ias)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.attachment not found", id)
}

// GetIrAttachments gets ir.attachment existing records.
func (c *Client) GetIrAttachments(ids []int64) (*IrAttachments, error) {
	ias := &IrAttachments{}
	if err := c.Read(IrAttachmentModel, ids, nil, ias); err != nil {
		return nil, err
	}
	return ias, nil
}

// FindIrAttachment finds ir.attachment record by querying it with criteria.
func (c *Client) FindIrAttachment(criteria *Criteria) (*IrAttachment, error) {
	ias := &IrAttachments{}
	if err := c.SearchRead(IrAttachmentModel, criteria, NewOptions().Limit(1), ias); err != nil {
		return nil, err
	}
	if ias != nil && len(*ias) > 0 {
		return &((*ias)[0]), nil
	}
	return nil, fmt.Errorf("ir.attachment was not found with criteria %v", criteria)
}

// FindIrAttachments finds ir.attachment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrAttachments(criteria *Criteria, options *Options) (*IrAttachments, error) {
	ias := &IrAttachments{}
	if err := c.SearchRead(IrAttachmentModel, criteria, options, ias); err != nil {
		return nil, err
	}
	return ias, nil
}

// FindIrAttachmentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrAttachmentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrAttachmentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrAttachmentId finds record id by querying it with criteria.
func (c *Client) FindIrAttachmentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrAttachmentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.attachment was not found with criteria %v and options %v", criteria, options)
}
