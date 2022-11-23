package odoo

import (
	"fmt"
)

// SaleOrderTemplateLine represents sale.order.template.line model.
type SaleOrderTemplateLine struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omitempty"`
	CompanyId            *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	DisplayType          *Selection `xmlrpc:"display_type,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	Name                 *String    `xmlrpc:"name,omitempty"`
	ProductId            *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductUomCategoryId *Many2One  `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomId         *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	ProductUomQty        *Float     `xmlrpc:"product_uom_qty,omitempty"`
	SaleOrderTemplateId  *Many2One  `xmlrpc:"sale_order_template_id,omitempty"`
	Sequence             *Int       `xmlrpc:"sequence,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SaleOrderTemplateLines represents array of sale.order.template.line model.
type SaleOrderTemplateLines []SaleOrderTemplateLine

// SaleOrderTemplateLineModel is the odoo model name.
const SaleOrderTemplateLineModel = "sale.order.template.line"

// Many2One convert SaleOrderTemplateLine to *Many2One.
func (sotl *SaleOrderTemplateLine) Many2One() *Many2One {
	return NewMany2One(sotl.Id.Get(), "")
}

// CreateSaleOrderTemplateLine creates a new sale.order.template.line model and returns its id.
func (c *Client) CreateSaleOrderTemplateLine(sotl *SaleOrderTemplateLine) (int64, error) {
	return c.Create(SaleOrderTemplateLineModel, sotl)
}

// UpdateSaleOrderTemplateLine updates an existing sale.order.template.line record.
func (c *Client) UpdateSaleOrderTemplateLine(sotl *SaleOrderTemplateLine) error {
	return c.UpdateSaleOrderTemplateLines([]int64{sotl.Id.Get()}, sotl)
}

// UpdateSaleOrderTemplateLines updates existing sale.order.template.line records.
// All records (represented by IDs) will be updated by sotl values.
func (c *Client) UpdateSaleOrderTemplateLines(ids []int64, sotl *SaleOrderTemplateLine) error {
	return c.Update(SaleOrderTemplateLineModel, ids, sotl)
}

// DeleteSaleOrderTemplateLine deletes an existing sale.order.template.line record.
func (c *Client) DeleteSaleOrderTemplateLine(id int64) error {
	return c.DeleteSaleOrderTemplateLines([]int64{id})
}

// DeleteSaleOrderTemplateLines deletes existing sale.order.template.line records.
func (c *Client) DeleteSaleOrderTemplateLines(ids []int64) error {
	return c.Delete(SaleOrderTemplateLineModel, ids)
}

// GetSaleOrderTemplateLine gets sale.order.template.line existing record.
func (c *Client) GetSaleOrderTemplateLine(id int64) (*SaleOrderTemplateLine, error) {
	sotls, err := c.GetSaleOrderTemplateLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if sotls != nil && len(*sotls) > 0 {
		return &((*sotls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.template.line not found", id)
}

// GetSaleOrderTemplateLines gets sale.order.template.line existing records.
func (c *Client) GetSaleOrderTemplateLines(ids []int64) (*SaleOrderTemplateLines, error) {
	sotls := &SaleOrderTemplateLines{}
	if err := c.Read(SaleOrderTemplateLineModel, ids, nil, sotls); err != nil {
		return nil, err
	}
	return sotls, nil
}

// FindSaleOrderTemplateLine finds sale.order.template.line record by querying it with criteria.
func (c *Client) FindSaleOrderTemplateLine(criteria *Criteria) (*SaleOrderTemplateLine, error) {
	sotls := &SaleOrderTemplateLines{}
	if err := c.SearchRead(SaleOrderTemplateLineModel, criteria, NewOptions().Limit(1), sotls); err != nil {
		return nil, err
	}
	if sotls != nil && len(*sotls) > 0 {
		return &((*sotls)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.template.line was not found")
}

// FindSaleOrderTemplateLines finds sale.order.template.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderTemplateLines(criteria *Criteria, options *Options) (*SaleOrderTemplateLines, error) {
	sotls := &SaleOrderTemplateLines{}
	if err := c.SearchRead(SaleOrderTemplateLineModel, criteria, options, sotls); err != nil {
		return nil, err
	}
	return sotls, nil
}

// FindSaleOrderTemplateLineIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderTemplateLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderTemplateLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderTemplateLineId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderTemplateLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderTemplateLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.template.line was not found")
}
