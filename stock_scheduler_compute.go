package odoo

// StockSchedulerCompute represents stock.scheduler.compute model.
type StockSchedulerCompute struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockSchedulerComputes represents array of stock.scheduler.compute model.
type StockSchedulerComputes []StockSchedulerCompute

// StockSchedulerComputeModel is the odoo model name.
const StockSchedulerComputeModel = "stock.scheduler.compute"

// Many2One convert StockSchedulerCompute to *Many2One.
func (ssc *StockSchedulerCompute) Many2One() *Many2One {
	return NewMany2One(ssc.Id.Get(), "")
}

// CreateStockSchedulerCompute creates a new stock.scheduler.compute model and returns its id.
func (c *Client) CreateStockSchedulerCompute(ssc *StockSchedulerCompute) (int64, error) {
	ids, err := c.CreateStockSchedulerComputes([]*StockSchedulerCompute{ssc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockSchedulerCompute creates a new stock.scheduler.compute model and returns its id.
func (c *Client) CreateStockSchedulerComputes(sscs []*StockSchedulerCompute) ([]int64, error) {
	var vv []interface{}
	for _, v := range sscs {
		vv = append(vv, v)
	}
	return c.Create(StockSchedulerComputeModel, vv, nil)
}

// UpdateStockSchedulerCompute updates an existing stock.scheduler.compute record.
func (c *Client) UpdateStockSchedulerCompute(ssc *StockSchedulerCompute) error {
	return c.UpdateStockSchedulerComputes([]int64{ssc.Id.Get()}, ssc)
}

// UpdateStockSchedulerComputes updates existing stock.scheduler.compute records.
// All records (represented by ids) will be updated by ssc values.
func (c *Client) UpdateStockSchedulerComputes(ids []int64, ssc *StockSchedulerCompute) error {
	return c.Update(StockSchedulerComputeModel, ids, ssc, nil)
}

// DeleteStockSchedulerCompute deletes an existing stock.scheduler.compute record.
func (c *Client) DeleteStockSchedulerCompute(id int64) error {
	return c.DeleteStockSchedulerComputes([]int64{id})
}

// DeleteStockSchedulerComputes deletes existing stock.scheduler.compute records.
func (c *Client) DeleteStockSchedulerComputes(ids []int64) error {
	return c.Delete(StockSchedulerComputeModel, ids)
}

// GetStockSchedulerCompute gets stock.scheduler.compute existing record.
func (c *Client) GetStockSchedulerCompute(id int64) (*StockSchedulerCompute, error) {
	sscs, err := c.GetStockSchedulerComputes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sscs)[0]), nil
}

// GetStockSchedulerComputes gets stock.scheduler.compute existing records.
func (c *Client) GetStockSchedulerComputes(ids []int64) (*StockSchedulerComputes, error) {
	sscs := &StockSchedulerComputes{}
	if err := c.Read(StockSchedulerComputeModel, ids, nil, sscs); err != nil {
		return nil, err
	}
	return sscs, nil
}

// FindStockSchedulerCompute finds stock.scheduler.compute record by querying it with criteria.
func (c *Client) FindStockSchedulerCompute(criteria *Criteria) (*StockSchedulerCompute, error) {
	sscs := &StockSchedulerComputes{}
	if err := c.SearchRead(StockSchedulerComputeModel, criteria, NewOptions().Limit(1), sscs); err != nil {
		return nil, err
	}
	return &((*sscs)[0]), nil
}

// FindStockSchedulerComputes finds stock.scheduler.compute records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockSchedulerComputes(criteria *Criteria, options *Options) (*StockSchedulerComputes, error) {
	sscs := &StockSchedulerComputes{}
	if err := c.SearchRead(StockSchedulerComputeModel, criteria, options, sscs); err != nil {
		return nil, err
	}
	return sscs, nil
}

// FindStockSchedulerComputeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockSchedulerComputeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockSchedulerComputeModel, criteria, options)
}

// FindStockSchedulerComputeId finds record id by querying it with criteria.
func (c *Client) FindStockSchedulerComputeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockSchedulerComputeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
