package odoo

// StockLocationPath represents stock.location.path model.
type StockLocationPath struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	Active         *Bool      `xmlrpc:"active,omitempty"`
	Auto           *Selection `xmlrpc:"auto,omitempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	Delay          *Int       `xmlrpc:"delay,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	LocationDestId *Many2One  `xmlrpc:"location_dest_id,omitempty"`
	LocationFromId *Many2One  `xmlrpc:"location_from_id,omitempty"`
	Name           *String    `xmlrpc:"name,omitempty"`
	PickingTypeId  *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	Propagate      *Bool      `xmlrpc:"propagate,omitempty"`
	RouteId        *Many2One  `xmlrpc:"route_id,omitempty"`
	RouteSequence  *Int       `xmlrpc:"route_sequence,omitempty"`
	Sequence       *Int       `xmlrpc:"sequence,omitempty"`
	WarehouseId    *Many2One  `xmlrpc:"warehouse_id,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockLocationPaths represents array of stock.location.path model.
type StockLocationPaths []StockLocationPath

// StockLocationPathModel is the odoo model name.
const StockLocationPathModel = "stock.location.path"

// Many2One convert StockLocationPath to *Many2One.
func (slp *StockLocationPath) Many2One() *Many2One {
	return NewMany2One(slp.Id.Get(), "")
}

// CreateStockLocationPath creates a new stock.location.path model and returns its id.
func (c *Client) CreateStockLocationPath(slp *StockLocationPath) (int64, error) {
	ids, err := c.CreateStockLocationPaths([]*StockLocationPath{slp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockLocationPath creates a new stock.location.path model and returns its id.
func (c *Client) CreateStockLocationPaths(slps []*StockLocationPath) ([]int64, error) {
	var vv []interface{}
	for _, v := range slps {
		vv = append(vv, v)
	}
	return c.Create(StockLocationPathModel, vv, nil)
}

// UpdateStockLocationPath updates an existing stock.location.path record.
func (c *Client) UpdateStockLocationPath(slp *StockLocationPath) error {
	return c.UpdateStockLocationPaths([]int64{slp.Id.Get()}, slp)
}

// UpdateStockLocationPaths updates existing stock.location.path records.
// All records (represented by ids) will be updated by slp values.
func (c *Client) UpdateStockLocationPaths(ids []int64, slp *StockLocationPath) error {
	return c.Update(StockLocationPathModel, ids, slp, nil)
}

// DeleteStockLocationPath deletes an existing stock.location.path record.
func (c *Client) DeleteStockLocationPath(id int64) error {
	return c.DeleteStockLocationPaths([]int64{id})
}

// DeleteStockLocationPaths deletes existing stock.location.path records.
func (c *Client) DeleteStockLocationPaths(ids []int64) error {
	return c.Delete(StockLocationPathModel, ids)
}

// GetStockLocationPath gets stock.location.path existing record.
func (c *Client) GetStockLocationPath(id int64) (*StockLocationPath, error) {
	slps, err := c.GetStockLocationPaths([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*slps)[0]), nil
}

// GetStockLocationPaths gets stock.location.path existing records.
func (c *Client) GetStockLocationPaths(ids []int64) (*StockLocationPaths, error) {
	slps := &StockLocationPaths{}
	if err := c.Read(StockLocationPathModel, ids, nil, slps); err != nil {
		return nil, err
	}
	return slps, nil
}

// FindStockLocationPath finds stock.location.path record by querying it with criteria.
func (c *Client) FindStockLocationPath(criteria *Criteria) (*StockLocationPath, error) {
	slps := &StockLocationPaths{}
	if err := c.SearchRead(StockLocationPathModel, criteria, NewOptions().Limit(1), slps); err != nil {
		return nil, err
	}
	return &((*slps)[0]), nil
}

// FindStockLocationPaths finds stock.location.path records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLocationPaths(criteria *Criteria, options *Options) (*StockLocationPaths, error) {
	slps := &StockLocationPaths{}
	if err := c.SearchRead(StockLocationPathModel, criteria, options, slps); err != nil {
		return nil, err
	}
	return slps, nil
}

// FindStockLocationPathIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLocationPathIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockLocationPathModel, criteria, options)
}

// FindStockLocationPathId finds record id by querying it with criteria.
func (c *Client) FindStockLocationPathId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockLocationPathModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
