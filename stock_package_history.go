package odoo

// StockPackageHistory represents stock.package.history model.
type StockPackageHistory struct {
	CompanyId       *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	LocationDestId  *Many2One `xmlrpc:"location_dest_id,omitempty"`
	LocationId      *Many2One `xmlrpc:"location_id,omitempty"`
	MoveLineIds     *Relation `xmlrpc:"move_line_ids,omitempty"`
	OutermostDestId *Many2One `xmlrpc:"outermost_dest_id,omitempty"`
	PackageId       *Many2One `xmlrpc:"package_id,omitempty"`
	PackageName     *String   `xmlrpc:"package_name,omitempty"`
	PackageTypeId   *Many2One `xmlrpc:"package_type_id,omitempty"`
	ParentDestId    *Many2One `xmlrpc:"parent_dest_id,omitempty"`
	ParentDestName  *String   `xmlrpc:"parent_dest_name,omitempty"`
	ParentOrigId    *Many2One `xmlrpc:"parent_orig_id,omitempty"`
	ParentOrigName  *String   `xmlrpc:"parent_orig_name,omitempty"`
	PickingIds      *Relation `xmlrpc:"picking_ids,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockPackageHistorys represents array of stock.package.history model.
type StockPackageHistorys []StockPackageHistory

// StockPackageHistoryModel is the odoo model name.
const StockPackageHistoryModel = "stock.package.history"

// Many2One convert StockPackageHistory to *Many2One.
func (sph *StockPackageHistory) Many2One() *Many2One {
	return NewMany2One(sph.Id.Get(), "")
}

// CreateStockPackageHistory creates a new stock.package.history model and returns its id.
func (c *Client) CreateStockPackageHistory(sph *StockPackageHistory) (int64, error) {
	ids, err := c.CreateStockPackageHistorys([]*StockPackageHistory{sph})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockPackageHistory creates a new stock.package.history model and returns its id.
func (c *Client) CreateStockPackageHistorys(sphs []*StockPackageHistory) ([]int64, error) {
	var vv []interface{}
	for _, v := range sphs {
		vv = append(vv, v)
	}
	return c.Create(StockPackageHistoryModel, vv, nil)
}

// UpdateStockPackageHistory updates an existing stock.package.history record.
func (c *Client) UpdateStockPackageHistory(sph *StockPackageHistory) error {
	return c.UpdateStockPackageHistorys([]int64{sph.Id.Get()}, sph)
}

// UpdateStockPackageHistorys updates existing stock.package.history records.
// All records (represented by ids) will be updated by sph values.
func (c *Client) UpdateStockPackageHistorys(ids []int64, sph *StockPackageHistory) error {
	return c.Update(StockPackageHistoryModel, ids, sph, nil)
}

// DeleteStockPackageHistory deletes an existing stock.package.history record.
func (c *Client) DeleteStockPackageHistory(id int64) error {
	return c.DeleteStockPackageHistorys([]int64{id})
}

// DeleteStockPackageHistorys deletes existing stock.package.history records.
func (c *Client) DeleteStockPackageHistorys(ids []int64) error {
	return c.Delete(StockPackageHistoryModel, ids)
}

// GetStockPackageHistory gets stock.package.history existing record.
func (c *Client) GetStockPackageHistory(id int64) (*StockPackageHistory, error) {
	sphs, err := c.GetStockPackageHistorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sphs)[0]), nil
}

// GetStockPackageHistorys gets stock.package.history existing records.
func (c *Client) GetStockPackageHistorys(ids []int64) (*StockPackageHistorys, error) {
	sphs := &StockPackageHistorys{}
	if err := c.Read(StockPackageHistoryModel, ids, nil, sphs); err != nil {
		return nil, err
	}
	return sphs, nil
}

// FindStockPackageHistory finds stock.package.history record by querying it with criteria.
func (c *Client) FindStockPackageHistory(criteria *Criteria) (*StockPackageHistory, error) {
	sphs := &StockPackageHistorys{}
	if err := c.SearchRead(StockPackageHistoryModel, criteria, NewOptions().Limit(1), sphs); err != nil {
		return nil, err
	}
	return &((*sphs)[0]), nil
}

// FindStockPackageHistorys finds stock.package.history records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageHistorys(criteria *Criteria, options *Options) (*StockPackageHistorys, error) {
	sphs := &StockPackageHistorys{}
	if err := c.SearchRead(StockPackageHistoryModel, criteria, options, sphs); err != nil {
		return nil, err
	}
	return sphs, nil
}

// FindStockPackageHistoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageHistoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockPackageHistoryModel, criteria, options)
}

// FindStockPackageHistoryId finds record id by querying it with criteria.
func (c *Client) FindStockPackageHistoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPackageHistoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
