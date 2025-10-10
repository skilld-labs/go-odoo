package odoo

// StockStorageCategoryCapacity represents stock.storage.category.capacity model.
type StockStorageCategoryCapacity struct {
	CompanyId         *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	PackageTypeId     *Many2One `xmlrpc:"package_type_id,omitempty"`
	ProductId         *Many2One `xmlrpc:"product_id,omitempty"`
	ProductUomId      *Many2One `xmlrpc:"product_uom_id,omitempty"`
	Quantity          *Float    `xmlrpc:"quantity,omitempty"`
	StorageCategoryId *Many2One `xmlrpc:"storage_category_id,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockStorageCategoryCapacitys represents array of stock.storage.category.capacity model.
type StockStorageCategoryCapacitys []StockStorageCategoryCapacity

// StockStorageCategoryCapacityModel is the odoo model name.
const StockStorageCategoryCapacityModel = "stock.storage.category.capacity"

// Many2One convert StockStorageCategoryCapacity to *Many2One.
func (sscc *StockStorageCategoryCapacity) Many2One() *Many2One {
	return NewMany2One(sscc.Id.Get(), "")
}

// CreateStockStorageCategoryCapacity creates a new stock.storage.category.capacity model and returns its id.
func (c *Client) CreateStockStorageCategoryCapacity(sscc *StockStorageCategoryCapacity) (int64, error) {
	ids, err := c.CreateStockStorageCategoryCapacitys([]*StockStorageCategoryCapacity{sscc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockStorageCategoryCapacity creates a new stock.storage.category.capacity model and returns its id.
func (c *Client) CreateStockStorageCategoryCapacitys(ssccs []*StockStorageCategoryCapacity) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssccs {
		vv = append(vv, v)
	}
	return c.Create(StockStorageCategoryCapacityModel, vv, nil)
}

// UpdateStockStorageCategoryCapacity updates an existing stock.storage.category.capacity record.
func (c *Client) UpdateStockStorageCategoryCapacity(sscc *StockStorageCategoryCapacity) error {
	return c.UpdateStockStorageCategoryCapacitys([]int64{sscc.Id.Get()}, sscc)
}

// UpdateStockStorageCategoryCapacitys updates existing stock.storage.category.capacity records.
// All records (represented by ids) will be updated by sscc values.
func (c *Client) UpdateStockStorageCategoryCapacitys(ids []int64, sscc *StockStorageCategoryCapacity) error {
	return c.Update(StockStorageCategoryCapacityModel, ids, sscc, nil)
}

// DeleteStockStorageCategoryCapacity deletes an existing stock.storage.category.capacity record.
func (c *Client) DeleteStockStorageCategoryCapacity(id int64) error {
	return c.DeleteStockStorageCategoryCapacitys([]int64{id})
}

// DeleteStockStorageCategoryCapacitys deletes existing stock.storage.category.capacity records.
func (c *Client) DeleteStockStorageCategoryCapacitys(ids []int64) error {
	return c.Delete(StockStorageCategoryCapacityModel, ids)
}

// GetStockStorageCategoryCapacity gets stock.storage.category.capacity existing record.
func (c *Client) GetStockStorageCategoryCapacity(id int64) (*StockStorageCategoryCapacity, error) {
	ssccs, err := c.GetStockStorageCategoryCapacitys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ssccs)[0]), nil
}

// GetStockStorageCategoryCapacitys gets stock.storage.category.capacity existing records.
func (c *Client) GetStockStorageCategoryCapacitys(ids []int64) (*StockStorageCategoryCapacitys, error) {
	ssccs := &StockStorageCategoryCapacitys{}
	if err := c.Read(StockStorageCategoryCapacityModel, ids, nil, ssccs); err != nil {
		return nil, err
	}
	return ssccs, nil
}

// FindStockStorageCategoryCapacity finds stock.storage.category.capacity record by querying it with criteria.
func (c *Client) FindStockStorageCategoryCapacity(criteria *Criteria) (*StockStorageCategoryCapacity, error) {
	ssccs := &StockStorageCategoryCapacitys{}
	if err := c.SearchRead(StockStorageCategoryCapacityModel, criteria, NewOptions().Limit(1), ssccs); err != nil {
		return nil, err
	}
	return &((*ssccs)[0]), nil
}

// FindStockStorageCategoryCapacitys finds stock.storage.category.capacity records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockStorageCategoryCapacitys(criteria *Criteria, options *Options) (*StockStorageCategoryCapacitys, error) {
	ssccs := &StockStorageCategoryCapacitys{}
	if err := c.SearchRead(StockStorageCategoryCapacityModel, criteria, options, ssccs); err != nil {
		return nil, err
	}
	return ssccs, nil
}

// FindStockStorageCategoryCapacityIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockStorageCategoryCapacityIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockStorageCategoryCapacityModel, criteria, options)
}

// FindStockStorageCategoryCapacityId finds record id by querying it with criteria.
func (c *Client) FindStockStorageCategoryCapacityId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockStorageCategoryCapacityModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
