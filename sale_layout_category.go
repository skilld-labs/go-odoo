package odoo

// SaleLayoutCategory represents sale.layout_category model.
type SaleLayoutCategory struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Pagebreak   *Bool     `xmlrpc:"pagebreak,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	Subtotal    *Bool     `xmlrpc:"subtotal,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleLayoutCategorys represents array of sale.layout_category model.
type SaleLayoutCategorys []SaleLayoutCategory

// SaleLayoutCategoryModel is the odoo model name.
const SaleLayoutCategoryModel = "sale.layout_category"

// Many2One convert SaleLayoutCategory to *Many2One.
func (sl *SaleLayoutCategory) Many2One() *Many2One {
	return NewMany2One(sl.Id.Get(), "")
}

// CreateSaleLayoutCategory creates a new sale.layout_category model and returns its id.
func (c *Client) CreateSaleLayoutCategory(sl *SaleLayoutCategory) (int64, error) {
	ids, err := c.CreateSaleLayoutCategorys([]*SaleLayoutCategory{sl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleLayoutCategorys creates a new sale.layout_category model and returns its id.
func (c *Client) CreateSaleLayoutCategorys(sls []*SaleLayoutCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range sls {
		vv = append(vv, v)
	}
	return c.Create(SaleLayoutCategoryModel, vv, nil)
}

// UpdateSaleLayoutCategory updates an existing sale.layout_category record.
func (c *Client) UpdateSaleLayoutCategory(sl *SaleLayoutCategory) error {
	return c.UpdateSaleLayoutCategorys([]int64{sl.Id.Get()}, sl)
}

// UpdateSaleLayoutCategorys updates existing sale.layout_category records.
// All records (represented by ids) will be updated by sl values.
func (c *Client) UpdateSaleLayoutCategorys(ids []int64, sl *SaleLayoutCategory) error {
	return c.Update(SaleLayoutCategoryModel, ids, sl, nil)
}

// DeleteSaleLayoutCategory deletes an existing sale.layout_category record.
func (c *Client) DeleteSaleLayoutCategory(id int64) error {
	return c.DeleteSaleLayoutCategorys([]int64{id})
}

// DeleteSaleLayoutCategorys deletes existing sale.layout_category records.
func (c *Client) DeleteSaleLayoutCategorys(ids []int64) error {
	return c.Delete(SaleLayoutCategoryModel, ids)
}

// GetSaleLayoutCategory gets sale.layout_category existing record.
func (c *Client) GetSaleLayoutCategory(id int64) (*SaleLayoutCategory, error) {
	sls, err := c.GetSaleLayoutCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sls)[0]), nil
}

// GetSaleLayoutCategorys gets sale.layout_category existing records.
func (c *Client) GetSaleLayoutCategorys(ids []int64) (*SaleLayoutCategorys, error) {
	sls := &SaleLayoutCategorys{}
	if err := c.Read(SaleLayoutCategoryModel, ids, nil, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindSaleLayoutCategory finds sale.layout_category record by querying it with criteria.
func (c *Client) FindSaleLayoutCategory(criteria *Criteria) (*SaleLayoutCategory, error) {
	sls := &SaleLayoutCategorys{}
	if err := c.SearchRead(SaleLayoutCategoryModel, criteria, NewOptions().Limit(1), sls); err != nil {
		return nil, err
	}
	return &((*sls)[0]), nil
}

// FindSaleLayoutCategorys finds sale.layout_category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleLayoutCategorys(criteria *Criteria, options *Options) (*SaleLayoutCategorys, error) {
	sls := &SaleLayoutCategorys{}
	if err := c.SearchRead(SaleLayoutCategoryModel, criteria, options, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindSaleLayoutCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleLayoutCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleLayoutCategoryModel, criteria, options)
}

// FindSaleLayoutCategoryId finds record id by querying it with criteria.
func (c *Client) FindSaleLayoutCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleLayoutCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
