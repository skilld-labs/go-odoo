package odoo

// StockPackageType represents stock.package.type model.
type StockPackageType struct {
	Barcode                    *String   `xmlrpc:"barcode,omitempty"`
	BaseWeight                 *Float    `xmlrpc:"base_weight,omitempty"`
	CompanyId                  *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate                 *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName                *String   `xmlrpc:"display_name,omitempty"`
	Height                     *Float    `xmlrpc:"height,omitempty"`
	Id                         *Int      `xmlrpc:"id,omitempty"`
	LengthUomName              *String   `xmlrpc:"length_uom_name,omitempty"`
	MaxWeight                  *Float    `xmlrpc:"max_weight,omitempty"`
	Name                       *String   `xmlrpc:"name,omitempty"`
	PackagingLength            *Float    `xmlrpc:"packaging_length,omitempty"`
	Sequence                   *Int      `xmlrpc:"sequence,omitempty"`
	StorageCategoryCapacityIds *Relation `xmlrpc:"storage_category_capacity_ids,omitempty"`
	WeightUomName              *String   `xmlrpc:"weight_uom_name,omitempty"`
	Width                      *Float    `xmlrpc:"width,omitempty"`
	WriteDate                  *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockPackageTypes represents array of stock.package.type model.
type StockPackageTypes []StockPackageType

// StockPackageTypeModel is the odoo model name.
const StockPackageTypeModel = "stock.package.type"

// Many2One convert StockPackageType to *Many2One.
func (spt *StockPackageType) Many2One() *Many2One {
	return NewMany2One(spt.Id.Get(), "")
}

// CreateStockPackageType creates a new stock.package.type model and returns its id.
func (c *Client) CreateStockPackageType(spt *StockPackageType) (int64, error) {
	ids, err := c.CreateStockPackageTypes([]*StockPackageType{spt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockPackageType creates a new stock.package.type model and returns its id.
func (c *Client) CreateStockPackageTypes(spts []*StockPackageType) ([]int64, error) {
	var vv []interface{}
	for _, v := range spts {
		vv = append(vv, v)
	}
	return c.Create(StockPackageTypeModel, vv, nil)
}

// UpdateStockPackageType updates an existing stock.package.type record.
func (c *Client) UpdateStockPackageType(spt *StockPackageType) error {
	return c.UpdateStockPackageTypes([]int64{spt.Id.Get()}, spt)
}

// UpdateStockPackageTypes updates existing stock.package.type records.
// All records (represented by ids) will be updated by spt values.
func (c *Client) UpdateStockPackageTypes(ids []int64, spt *StockPackageType) error {
	return c.Update(StockPackageTypeModel, ids, spt, nil)
}

// DeleteStockPackageType deletes an existing stock.package.type record.
func (c *Client) DeleteStockPackageType(id int64) error {
	return c.DeleteStockPackageTypes([]int64{id})
}

// DeleteStockPackageTypes deletes existing stock.package.type records.
func (c *Client) DeleteStockPackageTypes(ids []int64) error {
	return c.Delete(StockPackageTypeModel, ids)
}

// GetStockPackageType gets stock.package.type existing record.
func (c *Client) GetStockPackageType(id int64) (*StockPackageType, error) {
	spts, err := c.GetStockPackageTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*spts)[0]), nil
}

// GetStockPackageTypes gets stock.package.type existing records.
func (c *Client) GetStockPackageTypes(ids []int64) (*StockPackageTypes, error) {
	spts := &StockPackageTypes{}
	if err := c.Read(StockPackageTypeModel, ids, nil, spts); err != nil {
		return nil, err
	}
	return spts, nil
}

// FindStockPackageType finds stock.package.type record by querying it with criteria.
func (c *Client) FindStockPackageType(criteria *Criteria) (*StockPackageType, error) {
	spts := &StockPackageTypes{}
	if err := c.SearchRead(StockPackageTypeModel, criteria, NewOptions().Limit(1), spts); err != nil {
		return nil, err
	}
	return &((*spts)[0]), nil
}

// FindStockPackageTypes finds stock.package.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageTypes(criteria *Criteria, options *Options) (*StockPackageTypes, error) {
	spts := &StockPackageTypes{}
	if err := c.SearchRead(StockPackageTypeModel, criteria, options, spts); err != nil {
		return nil, err
	}
	return spts, nil
}

// FindStockPackageTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockPackageTypeModel, criteria, options)
}

// FindStockPackageTypeId finds record id by querying it with criteria.
func (c *Client) FindStockPackageTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPackageTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
