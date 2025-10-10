package odoo

// SaleOrderTemplateOption represents sale.order.template.option model.
type SaleOrderTemplateOption struct {
	CompanyId            *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	Name                 *String   `xmlrpc:"name,omitempty"`
	ProductId            *Many2One `xmlrpc:"product_id,omitempty"`
	ProductUomCategoryId *Many2One `xmlrpc:"product_uom_category_id,omitempty"`
	Quantity             *Float    `xmlrpc:"quantity,omitempty"`
	SaleOrderTemplateId  *Many2One `xmlrpc:"sale_order_template_id,omitempty"`
	UomId                *Many2One `xmlrpc:"uom_id,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleOrderTemplateOptions represents array of sale.order.template.option model.
type SaleOrderTemplateOptions []SaleOrderTemplateOption

// SaleOrderTemplateOptionModel is the odoo model name.
const SaleOrderTemplateOptionModel = "sale.order.template.option"

// Many2One convert SaleOrderTemplateOption to *Many2One.
func (soto *SaleOrderTemplateOption) Many2One() *Many2One {
	return NewMany2One(soto.Id.Get(), "")
}

// CreateSaleOrderTemplateOption creates a new sale.order.template.option model and returns its id.
func (c *Client) CreateSaleOrderTemplateOption(soto *SaleOrderTemplateOption) (int64, error) {
	ids, err := c.CreateSaleOrderTemplateOptions([]*SaleOrderTemplateOption{soto})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderTemplateOption creates a new sale.order.template.option model and returns its id.
func (c *Client) CreateSaleOrderTemplateOptions(sotos []*SaleOrderTemplateOption) ([]int64, error) {
	var vv []interface{}
	for _, v := range sotos {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderTemplateOptionModel, vv, nil)
}

// UpdateSaleOrderTemplateOption updates an existing sale.order.template.option record.
func (c *Client) UpdateSaleOrderTemplateOption(soto *SaleOrderTemplateOption) error {
	return c.UpdateSaleOrderTemplateOptions([]int64{soto.Id.Get()}, soto)
}

// UpdateSaleOrderTemplateOptions updates existing sale.order.template.option records.
// All records (represented by ids) will be updated by soto values.
func (c *Client) UpdateSaleOrderTemplateOptions(ids []int64, soto *SaleOrderTemplateOption) error {
	return c.Update(SaleOrderTemplateOptionModel, ids, soto, nil)
}

// DeleteSaleOrderTemplateOption deletes an existing sale.order.template.option record.
func (c *Client) DeleteSaleOrderTemplateOption(id int64) error {
	return c.DeleteSaleOrderTemplateOptions([]int64{id})
}

// DeleteSaleOrderTemplateOptions deletes existing sale.order.template.option records.
func (c *Client) DeleteSaleOrderTemplateOptions(ids []int64) error {
	return c.Delete(SaleOrderTemplateOptionModel, ids)
}

// GetSaleOrderTemplateOption gets sale.order.template.option existing record.
func (c *Client) GetSaleOrderTemplateOption(id int64) (*SaleOrderTemplateOption, error) {
	sotos, err := c.GetSaleOrderTemplateOptions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sotos)[0]), nil
}

// GetSaleOrderTemplateOptions gets sale.order.template.option existing records.
func (c *Client) GetSaleOrderTemplateOptions(ids []int64) (*SaleOrderTemplateOptions, error) {
	sotos := &SaleOrderTemplateOptions{}
	if err := c.Read(SaleOrderTemplateOptionModel, ids, nil, sotos); err != nil {
		return nil, err
	}
	return sotos, nil
}

// FindSaleOrderTemplateOption finds sale.order.template.option record by querying it with criteria.
func (c *Client) FindSaleOrderTemplateOption(criteria *Criteria) (*SaleOrderTemplateOption, error) {
	sotos := &SaleOrderTemplateOptions{}
	if err := c.SearchRead(SaleOrderTemplateOptionModel, criteria, NewOptions().Limit(1), sotos); err != nil {
		return nil, err
	}
	return &((*sotos)[0]), nil
}

// FindSaleOrderTemplateOptions finds sale.order.template.option records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderTemplateOptions(criteria *Criteria, options *Options) (*SaleOrderTemplateOptions, error) {
	sotos := &SaleOrderTemplateOptions{}
	if err := c.SearchRead(SaleOrderTemplateOptionModel, criteria, options, sotos); err != nil {
		return nil, err
	}
	return sotos, nil
}

// FindSaleOrderTemplateOptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderTemplateOptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleOrderTemplateOptionModel, criteria, options)
}

// FindSaleOrderTemplateOptionId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderTemplateOptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderTemplateOptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
