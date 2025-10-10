package odoo

// StockStorageCategory represents stock.storage.category model.
type StockStorageCategory struct {
	AllowNewProduct    *Selection `xmlrpc:"allow_new_product,omitempty"`
	CapacityIds        *Relation  `xmlrpc:"capacity_ids,omitempty"`
	CompanyId          *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	LocationIds        *Relation  `xmlrpc:"location_ids,omitempty"`
	MaxWeight          *Float     `xmlrpc:"max_weight,omitempty"`
	Name               *String    `xmlrpc:"name,omitempty"`
	PackageCapacityIds *Relation  `xmlrpc:"package_capacity_ids,omitempty"`
	ProductCapacityIds *Relation  `xmlrpc:"product_capacity_ids,omitempty"`
	WeightUomName      *String    `xmlrpc:"weight_uom_name,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockStorageCategorys represents array of stock.storage.category model.
type StockStorageCategorys []StockStorageCategory

// StockStorageCategoryModel is the odoo model name.
const StockStorageCategoryModel = "stock.storage.category"

// Many2One convert StockStorageCategory to *Many2One.
func (ssc *StockStorageCategory) Many2One() *Many2One {
	return NewMany2One(ssc.Id.Get(), "")
}

// CreateStockStorageCategory creates a new stock.storage.category model and returns its id.
func (c *Client) CreateStockStorageCategory(ssc *StockStorageCategory) (int64, error) {
	ids, err := c.CreateStockStorageCategorys([]*StockStorageCategory{ssc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockStorageCategory creates a new stock.storage.category model and returns its id.
func (c *Client) CreateStockStorageCategorys(sscs []*StockStorageCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range sscs {
		vv = append(vv, v)
	}
	return c.Create(StockStorageCategoryModel, vv, nil)
}

// UpdateStockStorageCategory updates an existing stock.storage.category record.
func (c *Client) UpdateStockStorageCategory(ssc *StockStorageCategory) error {
	return c.UpdateStockStorageCategorys([]int64{ssc.Id.Get()}, ssc)
}

// UpdateStockStorageCategorys updates existing stock.storage.category records.
// All records (represented by ids) will be updated by ssc values.
func (c *Client) UpdateStockStorageCategorys(ids []int64, ssc *StockStorageCategory) error {
	return c.Update(StockStorageCategoryModel, ids, ssc, nil)
}

// DeleteStockStorageCategory deletes an existing stock.storage.category record.
func (c *Client) DeleteStockStorageCategory(id int64) error {
	return c.DeleteStockStorageCategorys([]int64{id})
}

// DeleteStockStorageCategorys deletes existing stock.storage.category records.
func (c *Client) DeleteStockStorageCategorys(ids []int64) error {
	return c.Delete(StockStorageCategoryModel, ids)
}

// GetStockStorageCategory gets stock.storage.category existing record.
func (c *Client) GetStockStorageCategory(id int64) (*StockStorageCategory, error) {
	sscs, err := c.GetStockStorageCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sscs)[0]), nil
}

// GetStockStorageCategorys gets stock.storage.category existing records.
func (c *Client) GetStockStorageCategorys(ids []int64) (*StockStorageCategorys, error) {
	sscs := &StockStorageCategorys{}
	if err := c.Read(StockStorageCategoryModel, ids, nil, sscs); err != nil {
		return nil, err
	}
	return sscs, nil
}

// FindStockStorageCategory finds stock.storage.category record by querying it with criteria.
func (c *Client) FindStockStorageCategory(criteria *Criteria) (*StockStorageCategory, error) {
	sscs := &StockStorageCategorys{}
	if err := c.SearchRead(StockStorageCategoryModel, criteria, NewOptions().Limit(1), sscs); err != nil {
		return nil, err
	}
	return &((*sscs)[0]), nil
}

// FindStockStorageCategorys finds stock.storage.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockStorageCategorys(criteria *Criteria, options *Options) (*StockStorageCategorys, error) {
	sscs := &StockStorageCategorys{}
	if err := c.SearchRead(StockStorageCategoryModel, criteria, options, sscs); err != nil {
		return nil, err
	}
	return sscs, nil
}

// FindStockStorageCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockStorageCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockStorageCategoryModel, criteria, options)
}

// FindStockStorageCategoryId finds record id by querying it with criteria.
func (c *Client) FindStockStorageCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockStorageCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
