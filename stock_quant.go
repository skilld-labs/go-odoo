package odoo

// StockQuant represents stock.quant model.
type StockQuant struct {
	AccountingDate             *Time       `xmlrpc:"accounting_date,omitempty"`
	AvailableQuantity          *Float      `xmlrpc:"available_quantity,omitempty"`
	CompanyId                  *Many2One   `xmlrpc:"company_id,omitempty"`
	CostMethod                 *Selection  `xmlrpc:"cost_method,omitempty"`
	CreateDate                 *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                 *Many2One   `xmlrpc:"currency_id,omitempty"`
	CyclicInventoryFrequency   *Int        `xmlrpc:"cyclic_inventory_frequency,omitempty"`
	DisplayName                *String     `xmlrpc:"display_name,omitempty"`
	DummyId                    *String     `xmlrpc:"dummy_id,omitempty"`
	ExpirationDate             *Time       `xmlrpc:"expiration_date,omitempty"`
	Id                         *Int        `xmlrpc:"id,omitempty"`
	Image1920                  *String     `xmlrpc:"image_1920,omitempty"`
	InDate                     *Time       `xmlrpc:"in_date,omitempty"`
	InventoryDate              *Time       `xmlrpc:"inventory_date,omitempty"`
	InventoryDiffQuantity      *Float      `xmlrpc:"inventory_diff_quantity,omitempty"`
	InventoryQuantity          *Float      `xmlrpc:"inventory_quantity,omitempty"`
	InventoryQuantityAutoApply *Float      `xmlrpc:"inventory_quantity_auto_apply,omitempty"`
	InventoryQuantitySet       *Bool       `xmlrpc:"inventory_quantity_set,omitempty"`
	IsOutdated                 *Bool       `xmlrpc:"is_outdated,omitempty"`
	LastCountDate              *Time       `xmlrpc:"last_count_date,omitempty"`
	LocationId                 *Many2One   `xmlrpc:"location_id,omitempty"`
	LotId                      *Many2One   `xmlrpc:"lot_id,omitempty"`
	LotProperties              interface{} `xmlrpc:"lot_properties,omitempty"`
	OnHand                     *Bool       `xmlrpc:"on_hand,omitempty"`
	OwnerId                    *Many2One   `xmlrpc:"owner_id,omitempty"`
	PackageId                  *Many2One   `xmlrpc:"package_id,omitempty"`
	Priority                   *Selection  `xmlrpc:"priority,omitempty"`
	ProductCategId             *Many2One   `xmlrpc:"product_categ_id,omitempty"`
	ProductId                  *Many2One   `xmlrpc:"product_id,omitempty"`
	ProductReferenceCode       *String     `xmlrpc:"product_reference_code,omitempty"`
	ProductTmplId              *Many2One   `xmlrpc:"product_tmpl_id,omitempty"`
	ProductUomId               *Many2One   `xmlrpc:"product_uom_id,omitempty"`
	Quantity                   *Float      `xmlrpc:"quantity,omitempty"`
	RemovalDate                *Time       `xmlrpc:"removal_date,omitempty"`
	ReservedQuantity           *Float      `xmlrpc:"reserved_quantity,omitempty"`
	SnDuplicated               *Bool       `xmlrpc:"sn_duplicated,omitempty"`
	StorageCategoryId          *Many2One   `xmlrpc:"storage_category_id,omitempty"`
	Tracking                   *Selection  `xmlrpc:"tracking,omitempty"`
	UseExpirationDate          *Bool       `xmlrpc:"use_expiration_date,omitempty"`
	UserId                     *Many2One   `xmlrpc:"user_id,omitempty"`
	Value                      *Float      `xmlrpc:"value,omitempty"`
	WarehouseId                *Many2One   `xmlrpc:"warehouse_id,omitempty"`
	WriteDate                  *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// StockQuants represents array of stock.quant model.
type StockQuants []StockQuant

// StockQuantModel is the odoo model name.
const StockQuantModel = "stock.quant"

// Many2One convert StockQuant to *Many2One.
func (sq *StockQuant) Many2One() *Many2One {
	return NewMany2One(sq.Id.Get(), "")
}

// CreateStockQuant creates a new stock.quant model and returns its id.
func (c *Client) CreateStockQuant(sq *StockQuant) (int64, error) {
	ids, err := c.CreateStockQuants([]*StockQuant{sq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockQuant creates a new stock.quant model and returns its id.
func (c *Client) CreateStockQuants(sqs []*StockQuant) ([]int64, error) {
	var vv []interface{}
	for _, v := range sqs {
		vv = append(vv, v)
	}
	return c.Create(StockQuantModel, vv, nil)
}

// UpdateStockQuant updates an existing stock.quant record.
func (c *Client) UpdateStockQuant(sq *StockQuant) error {
	return c.UpdateStockQuants([]int64{sq.Id.Get()}, sq)
}

// UpdateStockQuants updates existing stock.quant records.
// All records (represented by ids) will be updated by sq values.
func (c *Client) UpdateStockQuants(ids []int64, sq *StockQuant) error {
	return c.Update(StockQuantModel, ids, sq, nil)
}

// DeleteStockQuant deletes an existing stock.quant record.
func (c *Client) DeleteStockQuant(id int64) error {
	return c.DeleteStockQuants([]int64{id})
}

// DeleteStockQuants deletes existing stock.quant records.
func (c *Client) DeleteStockQuants(ids []int64) error {
	return c.Delete(StockQuantModel, ids)
}

// GetStockQuant gets stock.quant existing record.
func (c *Client) GetStockQuant(id int64) (*StockQuant, error) {
	sqs, err := c.GetStockQuants([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sqs)[0]), nil
}

// GetStockQuants gets stock.quant existing records.
func (c *Client) GetStockQuants(ids []int64) (*StockQuants, error) {
	sqs := &StockQuants{}
	if err := c.Read(StockQuantModel, ids, nil, sqs); err != nil {
		return nil, err
	}
	return sqs, nil
}

// FindStockQuant finds stock.quant record by querying it with criteria.
func (c *Client) FindStockQuant(criteria *Criteria) (*StockQuant, error) {
	sqs := &StockQuants{}
	if err := c.SearchRead(StockQuantModel, criteria, NewOptions().Limit(1), sqs); err != nil {
		return nil, err
	}
	return &((*sqs)[0]), nil
}

// FindStockQuants finds stock.quant records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockQuants(criteria *Criteria, options *Options) (*StockQuants, error) {
	sqs := &StockQuants{}
	if err := c.SearchRead(StockQuantModel, criteria, options, sqs); err != nil {
		return nil, err
	}
	return sqs, nil
}

// FindStockQuantIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockQuantIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockQuantModel, criteria, options)
}

// FindStockQuantId finds record id by querying it with criteria.
func (c *Client) FindStockQuantId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockQuantModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
