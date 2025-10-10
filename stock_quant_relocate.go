package odoo

// StockQuantRelocate represents stock.quant.relocate model.
type StockQuantRelocate struct {
	CompanyId           *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omitempty"`
	DestLocationId      *Many2One `xmlrpc:"dest_location_id,omitempty"`
	DestPackageId       *Many2One `xmlrpc:"dest_package_id,omitempty"`
	DestPackageIdDomain *String   `xmlrpc:"dest_package_id_domain,omitempty"`
	DisplayName         *String   `xmlrpc:"display_name,omitempty"`
	Id                  *Int      `xmlrpc:"id,omitempty"`
	IsMultiLocation     *Bool     `xmlrpc:"is_multi_location,omitempty"`
	IsPartialPackage    *Bool     `xmlrpc:"is_partial_package,omitempty"`
	Message             *String   `xmlrpc:"message,omitempty"`
	PartialPackageNames *String   `xmlrpc:"partial_package_names,omitempty"`
	QuantIds            *Relation `xmlrpc:"quant_ids,omitempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockQuantRelocates represents array of stock.quant.relocate model.
type StockQuantRelocates []StockQuantRelocate

// StockQuantRelocateModel is the odoo model name.
const StockQuantRelocateModel = "stock.quant.relocate"

// Many2One convert StockQuantRelocate to *Many2One.
func (sqr *StockQuantRelocate) Many2One() *Many2One {
	return NewMany2One(sqr.Id.Get(), "")
}

// CreateStockQuantRelocate creates a new stock.quant.relocate model and returns its id.
func (c *Client) CreateStockQuantRelocate(sqr *StockQuantRelocate) (int64, error) {
	ids, err := c.CreateStockQuantRelocates([]*StockQuantRelocate{sqr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockQuantRelocate creates a new stock.quant.relocate model and returns its id.
func (c *Client) CreateStockQuantRelocates(sqrs []*StockQuantRelocate) ([]int64, error) {
	var vv []interface{}
	for _, v := range sqrs {
		vv = append(vv, v)
	}
	return c.Create(StockQuantRelocateModel, vv, nil)
}

// UpdateStockQuantRelocate updates an existing stock.quant.relocate record.
func (c *Client) UpdateStockQuantRelocate(sqr *StockQuantRelocate) error {
	return c.UpdateStockQuantRelocates([]int64{sqr.Id.Get()}, sqr)
}

// UpdateStockQuantRelocates updates existing stock.quant.relocate records.
// All records (represented by ids) will be updated by sqr values.
func (c *Client) UpdateStockQuantRelocates(ids []int64, sqr *StockQuantRelocate) error {
	return c.Update(StockQuantRelocateModel, ids, sqr, nil)
}

// DeleteStockQuantRelocate deletes an existing stock.quant.relocate record.
func (c *Client) DeleteStockQuantRelocate(id int64) error {
	return c.DeleteStockQuantRelocates([]int64{id})
}

// DeleteStockQuantRelocates deletes existing stock.quant.relocate records.
func (c *Client) DeleteStockQuantRelocates(ids []int64) error {
	return c.Delete(StockQuantRelocateModel, ids)
}

// GetStockQuantRelocate gets stock.quant.relocate existing record.
func (c *Client) GetStockQuantRelocate(id int64) (*StockQuantRelocate, error) {
	sqrs, err := c.GetStockQuantRelocates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sqrs)[0]), nil
}

// GetStockQuantRelocates gets stock.quant.relocate existing records.
func (c *Client) GetStockQuantRelocates(ids []int64) (*StockQuantRelocates, error) {
	sqrs := &StockQuantRelocates{}
	if err := c.Read(StockQuantRelocateModel, ids, nil, sqrs); err != nil {
		return nil, err
	}
	return sqrs, nil
}

// FindStockQuantRelocate finds stock.quant.relocate record by querying it with criteria.
func (c *Client) FindStockQuantRelocate(criteria *Criteria) (*StockQuantRelocate, error) {
	sqrs := &StockQuantRelocates{}
	if err := c.SearchRead(StockQuantRelocateModel, criteria, NewOptions().Limit(1), sqrs); err != nil {
		return nil, err
	}
	return &((*sqrs)[0]), nil
}

// FindStockQuantRelocates finds stock.quant.relocate records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockQuantRelocates(criteria *Criteria, options *Options) (*StockQuantRelocates, error) {
	sqrs := &StockQuantRelocates{}
	if err := c.SearchRead(StockQuantRelocateModel, criteria, options, sqrs); err != nil {
		return nil, err
	}
	return sqrs, nil
}

// FindStockQuantRelocateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockQuantRelocateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockQuantRelocateModel, criteria, options)
}

// FindStockQuantRelocateId finds record id by querying it with criteria.
func (c *Client) FindStockQuantRelocateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockQuantRelocateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
