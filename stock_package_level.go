package odoo

// StockPackageLevel represents stock.package_level model.
type StockPackageLevel struct {
	CompanyId       *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	IsDone          *Bool      `xmlrpc:"is_done,omitempty"`
	IsFreshPackage  *Bool      `xmlrpc:"is_fresh_package,omitempty"`
	LocationDestId  *Many2One  `xmlrpc:"location_dest_id,omitempty"`
	LocationId      *Many2One  `xmlrpc:"location_id,omitempty"`
	MoveIds         *Relation  `xmlrpc:"move_ids,omitempty"`
	MoveLineIds     *Relation  `xmlrpc:"move_line_ids,omitempty"`
	PackageId       *Many2One  `xmlrpc:"package_id,omitempty"`
	PickingId       *Many2One  `xmlrpc:"picking_id,omitempty"`
	PickingTypeCode *Selection `xmlrpc:"picking_type_code,omitempty"`
	ShowLotsM2O     *Bool      `xmlrpc:"show_lots_m2o,omitempty"`
	ShowLotsText    *Bool      `xmlrpc:"show_lots_text,omitempty"`
	State           *Selection `xmlrpc:"state,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockPackageLevels represents array of stock.package_level model.
type StockPackageLevels []StockPackageLevel

// StockPackageLevelModel is the odoo model name.
const StockPackageLevelModel = "stock.package_level"

// Many2One convert StockPackageLevel to *Many2One.
func (sp *StockPackageLevel) Many2One() *Many2One {
	return NewMany2One(sp.Id.Get(), "")
}

// CreateStockPackageLevel creates a new stock.package_level model and returns its id.
func (c *Client) CreateStockPackageLevel(sp *StockPackageLevel) (int64, error) {
	ids, err := c.CreateStockPackageLevels([]*StockPackageLevel{sp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockPackageLevel creates a new stock.package_level model and returns its id.
func (c *Client) CreateStockPackageLevels(sps []*StockPackageLevel) ([]int64, error) {
	var vv []interface{}
	for _, v := range sps {
		vv = append(vv, v)
	}
	return c.Create(StockPackageLevelModel, vv, nil)
}

// UpdateStockPackageLevel updates an existing stock.package_level record.
func (c *Client) UpdateStockPackageLevel(sp *StockPackageLevel) error {
	return c.UpdateStockPackageLevels([]int64{sp.Id.Get()}, sp)
}

// UpdateStockPackageLevels updates existing stock.package_level records.
// All records (represented by ids) will be updated by sp values.
func (c *Client) UpdateStockPackageLevels(ids []int64, sp *StockPackageLevel) error {
	return c.Update(StockPackageLevelModel, ids, sp, nil)
}

// DeleteStockPackageLevel deletes an existing stock.package_level record.
func (c *Client) DeleteStockPackageLevel(id int64) error {
	return c.DeleteStockPackageLevels([]int64{id})
}

// DeleteStockPackageLevels deletes existing stock.package_level records.
func (c *Client) DeleteStockPackageLevels(ids []int64) error {
	return c.Delete(StockPackageLevelModel, ids)
}

// GetStockPackageLevel gets stock.package_level existing record.
func (c *Client) GetStockPackageLevel(id int64) (*StockPackageLevel, error) {
	sps, err := c.GetStockPackageLevels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sps)[0]), nil
}

// GetStockPackageLevels gets stock.package_level existing records.
func (c *Client) GetStockPackageLevels(ids []int64) (*StockPackageLevels, error) {
	sps := &StockPackageLevels{}
	if err := c.Read(StockPackageLevelModel, ids, nil, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindStockPackageLevel finds stock.package_level record by querying it with criteria.
func (c *Client) FindStockPackageLevel(criteria *Criteria) (*StockPackageLevel, error) {
	sps := &StockPackageLevels{}
	if err := c.SearchRead(StockPackageLevelModel, criteria, NewOptions().Limit(1), sps); err != nil {
		return nil, err
	}
	return &((*sps)[0]), nil
}

// FindStockPackageLevels finds stock.package_level records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageLevels(criteria *Criteria, options *Options) (*StockPackageLevels, error) {
	sps := &StockPackageLevels{}
	if err := c.SearchRead(StockPackageLevelModel, criteria, options, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindStockPackageLevelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageLevelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockPackageLevelModel, criteria, options)
}

// FindStockPackageLevelId finds record id by querying it with criteria.
func (c *Client) FindStockPackageLevelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPackageLevelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
