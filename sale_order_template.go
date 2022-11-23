package odoo

import (
	"fmt"
)

// SaleOrderTemplate represents sale.order.template model.
type SaleOrderTemplate struct {
	LastUpdate                 *Time     `xmlrpc:"__last_update,omitempty"`
	Active                     *Bool     `xmlrpc:"active,omitempty"`
	CompanyId                  *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate                 *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName                *String   `xmlrpc:"display_name,omitempty"`
	Id                         *Int      `xmlrpc:"id,omitempty"`
	MailTemplateId             *Many2One `xmlrpc:"mail_template_id,omitempty"`
	Name                       *String   `xmlrpc:"name,omitempty"`
	Note                       *String   `xmlrpc:"note,omitempty"`
	NumberOfDays               *Int      `xmlrpc:"number_of_days,omitempty"`
	RequirePayment             *Bool     `xmlrpc:"require_payment,omitempty"`
	RequireSignature           *Bool     `xmlrpc:"require_signature,omitempty"`
	SaleOrderTemplateLineIds   *Relation `xmlrpc:"sale_order_template_line_ids,omitempty"`
	SaleOrderTemplateOptionIds *Relation `xmlrpc:"sale_order_template_option_ids,omitempty"`
	WriteDate                  *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleOrderTemplates represents array of sale.order.template model.
type SaleOrderTemplates []SaleOrderTemplate

// SaleOrderTemplateModel is the odoo model name.
const SaleOrderTemplateModel = "sale.order.template"

// Many2One convert SaleOrderTemplate to *Many2One.
func (sot *SaleOrderTemplate) Many2One() *Many2One {
	return NewMany2One(sot.Id.Get(), "")
}

// CreateSaleOrderTemplate creates a new sale.order.template model and returns its id.
func (c *Client) CreateSaleOrderTemplate(sot *SaleOrderTemplate) (int64, error) {
	return c.Create(SaleOrderTemplateModel, sot)
}

// UpdateSaleOrderTemplate updates an existing sale.order.template record.
func (c *Client) UpdateSaleOrderTemplate(sot *SaleOrderTemplate) error {
	return c.UpdateSaleOrderTemplates([]int64{sot.Id.Get()}, sot)
}

// UpdateSaleOrderTemplates updates existing sale.order.template records.
// All records (represented by IDs) will be updated by sot values.
func (c *Client) UpdateSaleOrderTemplates(ids []int64, sot *SaleOrderTemplate) error {
	return c.Update(SaleOrderTemplateModel, ids, sot)
}

// DeleteSaleOrderTemplate deletes an existing sale.order.template record.
func (c *Client) DeleteSaleOrderTemplate(id int64) error {
	return c.DeleteSaleOrderTemplates([]int64{id})
}

// DeleteSaleOrderTemplates deletes existing sale.order.template records.
func (c *Client) DeleteSaleOrderTemplates(ids []int64) error {
	return c.Delete(SaleOrderTemplateModel, ids)
}

// GetSaleOrderTemplate gets sale.order.template existing record.
func (c *Client) GetSaleOrderTemplate(id int64) (*SaleOrderTemplate, error) {
	sots, err := c.GetSaleOrderTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if sots != nil && len(*sots) > 0 {
		return &((*sots)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.template not found", id)
}

// GetSaleOrderTemplates gets sale.order.template existing records.
func (c *Client) GetSaleOrderTemplates(ids []int64) (*SaleOrderTemplates, error) {
	sots := &SaleOrderTemplates{}
	if err := c.Read(SaleOrderTemplateModel, ids, nil, sots); err != nil {
		return nil, err
	}
	return sots, nil
}

// FindSaleOrderTemplate finds sale.order.template record by querying it with criteria.
func (c *Client) FindSaleOrderTemplate(criteria *Criteria) (*SaleOrderTemplate, error) {
	sots := &SaleOrderTemplates{}
	if err := c.SearchRead(SaleOrderTemplateModel, criteria, NewOptions().Limit(1), sots); err != nil {
		return nil, err
	}
	if sots != nil && len(*sots) > 0 {
		return &((*sots)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.template was not found")
}

// FindSaleOrderTemplates finds sale.order.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderTemplates(criteria *Criteria, options *Options) (*SaleOrderTemplates, error) {
	sots := &SaleOrderTemplates{}
	if err := c.SearchRead(SaleOrderTemplateModel, criteria, options, sots); err != nil {
		return nil, err
	}
	return sots, nil
}

// FindSaleOrderTemplateIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderTemplateId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.template was not found")
}
