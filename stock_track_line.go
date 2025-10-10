package odoo

// StockTrackLine represents stock.track.line model.
type StockTrackLine struct {
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	ProductDisplayName *String    `xmlrpc:"product_display_name,omitempty"`
	ProductId          *Many2One  `xmlrpc:"product_id,omitempty"`
	Tracking           *Selection `xmlrpc:"tracking,omitempty"`
	WizardId           *Many2One  `xmlrpc:"wizard_id,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockTrackLines represents array of stock.track.line model.
type StockTrackLines []StockTrackLine

// StockTrackLineModel is the odoo model name.
const StockTrackLineModel = "stock.track.line"

// Many2One convert StockTrackLine to *Many2One.
func (stl *StockTrackLine) Many2One() *Many2One {
	return NewMany2One(stl.Id.Get(), "")
}

// CreateStockTrackLine creates a new stock.track.line model and returns its id.
func (c *Client) CreateStockTrackLine(stl *StockTrackLine) (int64, error) {
	ids, err := c.CreateStockTrackLines([]*StockTrackLine{stl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockTrackLine creates a new stock.track.line model and returns its id.
func (c *Client) CreateStockTrackLines(stls []*StockTrackLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range stls {
		vv = append(vv, v)
	}
	return c.Create(StockTrackLineModel, vv, nil)
}

// UpdateStockTrackLine updates an existing stock.track.line record.
func (c *Client) UpdateStockTrackLine(stl *StockTrackLine) error {
	return c.UpdateStockTrackLines([]int64{stl.Id.Get()}, stl)
}

// UpdateStockTrackLines updates existing stock.track.line records.
// All records (represented by ids) will be updated by stl values.
func (c *Client) UpdateStockTrackLines(ids []int64, stl *StockTrackLine) error {
	return c.Update(StockTrackLineModel, ids, stl, nil)
}

// DeleteStockTrackLine deletes an existing stock.track.line record.
func (c *Client) DeleteStockTrackLine(id int64) error {
	return c.DeleteStockTrackLines([]int64{id})
}

// DeleteStockTrackLines deletes existing stock.track.line records.
func (c *Client) DeleteStockTrackLines(ids []int64) error {
	return c.Delete(StockTrackLineModel, ids)
}

// GetStockTrackLine gets stock.track.line existing record.
func (c *Client) GetStockTrackLine(id int64) (*StockTrackLine, error) {
	stls, err := c.GetStockTrackLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*stls)[0]), nil
}

// GetStockTrackLines gets stock.track.line existing records.
func (c *Client) GetStockTrackLines(ids []int64) (*StockTrackLines, error) {
	stls := &StockTrackLines{}
	if err := c.Read(StockTrackLineModel, ids, nil, stls); err != nil {
		return nil, err
	}
	return stls, nil
}

// FindStockTrackLine finds stock.track.line record by querying it with criteria.
func (c *Client) FindStockTrackLine(criteria *Criteria) (*StockTrackLine, error) {
	stls := &StockTrackLines{}
	if err := c.SearchRead(StockTrackLineModel, criteria, NewOptions().Limit(1), stls); err != nil {
		return nil, err
	}
	return &((*stls)[0]), nil
}

// FindStockTrackLines finds stock.track.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockTrackLines(criteria *Criteria, options *Options) (*StockTrackLines, error) {
	stls := &StockTrackLines{}
	if err := c.SearchRead(StockTrackLineModel, criteria, options, stls); err != nil {
		return nil, err
	}
	return stls, nil
}

// FindStockTrackLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockTrackLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockTrackLineModel, criteria, options)
}

// FindStockTrackLineId finds record id by querying it with criteria.
func (c *Client) FindStockTrackLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockTrackLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
