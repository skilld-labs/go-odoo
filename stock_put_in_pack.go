package odoo

// StockPutInPack represents stock.put.in.pack model.
type StockPutInPack struct {
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	LocationDestId        *Many2One `xmlrpc:"location_dest_id,omitempty"`
	MoveLineIds           *Relation `xmlrpc:"move_line_ids,omitempty"`
	OriginPackageIds      *Relation `xmlrpc:"origin_package_ids,omitempty"`
	PackageIds            *Relation `xmlrpc:"package_ids,omitempty"`
	PackageTypeId         *Many2One `xmlrpc:"package_type_id,omitempty"`
	PackageTypeSequenceId *Many2One `xmlrpc:"package_type_sequence_id,omitempty"`
	ResultPackageId       *Many2One `xmlrpc:"result_package_id,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockPutInPacks represents array of stock.put.in.pack model.
type StockPutInPacks []StockPutInPack

// StockPutInPackModel is the odoo model name.
const StockPutInPackModel = "stock.put.in.pack"

// Many2One convert StockPutInPack to *Many2One.
func (spip *StockPutInPack) Many2One() *Many2One {
	return NewMany2One(spip.Id.Get(), "")
}

// CreateStockPutInPack creates a new stock.put.in.pack model and returns its id.
func (c *Client) CreateStockPutInPack(spip *StockPutInPack) (int64, error) {
	ids, err := c.CreateStockPutInPacks([]*StockPutInPack{spip})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockPutInPack creates a new stock.put.in.pack model and returns its id.
func (c *Client) CreateStockPutInPacks(spips []*StockPutInPack) ([]int64, error) {
	var vv []interface{}
	for _, v := range spips {
		vv = append(vv, v)
	}
	return c.Create(StockPutInPackModel, vv, nil)
}

// UpdateStockPutInPack updates an existing stock.put.in.pack record.
func (c *Client) UpdateStockPutInPack(spip *StockPutInPack) error {
	return c.UpdateStockPutInPacks([]int64{spip.Id.Get()}, spip)
}

// UpdateStockPutInPacks updates existing stock.put.in.pack records.
// All records (represented by ids) will be updated by spip values.
func (c *Client) UpdateStockPutInPacks(ids []int64, spip *StockPutInPack) error {
	return c.Update(StockPutInPackModel, ids, spip, nil)
}

// DeleteStockPutInPack deletes an existing stock.put.in.pack record.
func (c *Client) DeleteStockPutInPack(id int64) error {
	return c.DeleteStockPutInPacks([]int64{id})
}

// DeleteStockPutInPacks deletes existing stock.put.in.pack records.
func (c *Client) DeleteStockPutInPacks(ids []int64) error {
	return c.Delete(StockPutInPackModel, ids)
}

// GetStockPutInPack gets stock.put.in.pack existing record.
func (c *Client) GetStockPutInPack(id int64) (*StockPutInPack, error) {
	spips, err := c.GetStockPutInPacks([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*spips)[0]), nil
}

// GetStockPutInPacks gets stock.put.in.pack existing records.
func (c *Client) GetStockPutInPacks(ids []int64) (*StockPutInPacks, error) {
	spips := &StockPutInPacks{}
	if err := c.Read(StockPutInPackModel, ids, nil, spips); err != nil {
		return nil, err
	}
	return spips, nil
}

// FindStockPutInPack finds stock.put.in.pack record by querying it with criteria.
func (c *Client) FindStockPutInPack(criteria *Criteria) (*StockPutInPack, error) {
	spips := &StockPutInPacks{}
	if err := c.SearchRead(StockPutInPackModel, criteria, NewOptions().Limit(1), spips); err != nil {
		return nil, err
	}
	return &((*spips)[0]), nil
}

// FindStockPutInPacks finds stock.put.in.pack records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPutInPacks(criteria *Criteria, options *Options) (*StockPutInPacks, error) {
	spips := &StockPutInPacks{}
	if err := c.SearchRead(StockPutInPackModel, criteria, options, spips); err != nil {
		return nil, err
	}
	return spips, nil
}

// FindStockPutInPackIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPutInPackIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockPutInPackModel, criteria, options)
}

// FindStockPutInPackId finds record id by querying it with criteria.
func (c *Client) FindStockPutInPackId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPutInPackModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
