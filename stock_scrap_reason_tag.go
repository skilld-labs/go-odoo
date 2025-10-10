package odoo

// StockScrapReasonTag represents stock.scrap.reason.tag model.
type StockScrapReasonTag struct {
	Color       *String   `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockScrapReasonTags represents array of stock.scrap.reason.tag model.
type StockScrapReasonTags []StockScrapReasonTag

// StockScrapReasonTagModel is the odoo model name.
const StockScrapReasonTagModel = "stock.scrap.reason.tag"

// Many2One convert StockScrapReasonTag to *Many2One.
func (ssrt *StockScrapReasonTag) Many2One() *Many2One {
	return NewMany2One(ssrt.Id.Get(), "")
}

// CreateStockScrapReasonTag creates a new stock.scrap.reason.tag model and returns its id.
func (c *Client) CreateStockScrapReasonTag(ssrt *StockScrapReasonTag) (int64, error) {
	ids, err := c.CreateStockScrapReasonTags([]*StockScrapReasonTag{ssrt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockScrapReasonTag creates a new stock.scrap.reason.tag model and returns its id.
func (c *Client) CreateStockScrapReasonTags(ssrts []*StockScrapReasonTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssrts {
		vv = append(vv, v)
	}
	return c.Create(StockScrapReasonTagModel, vv, nil)
}

// UpdateStockScrapReasonTag updates an existing stock.scrap.reason.tag record.
func (c *Client) UpdateStockScrapReasonTag(ssrt *StockScrapReasonTag) error {
	return c.UpdateStockScrapReasonTags([]int64{ssrt.Id.Get()}, ssrt)
}

// UpdateStockScrapReasonTags updates existing stock.scrap.reason.tag records.
// All records (represented by ids) will be updated by ssrt values.
func (c *Client) UpdateStockScrapReasonTags(ids []int64, ssrt *StockScrapReasonTag) error {
	return c.Update(StockScrapReasonTagModel, ids, ssrt, nil)
}

// DeleteStockScrapReasonTag deletes an existing stock.scrap.reason.tag record.
func (c *Client) DeleteStockScrapReasonTag(id int64) error {
	return c.DeleteStockScrapReasonTags([]int64{id})
}

// DeleteStockScrapReasonTags deletes existing stock.scrap.reason.tag records.
func (c *Client) DeleteStockScrapReasonTags(ids []int64) error {
	return c.Delete(StockScrapReasonTagModel, ids)
}

// GetStockScrapReasonTag gets stock.scrap.reason.tag existing record.
func (c *Client) GetStockScrapReasonTag(id int64) (*StockScrapReasonTag, error) {
	ssrts, err := c.GetStockScrapReasonTags([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ssrts)[0]), nil
}

// GetStockScrapReasonTags gets stock.scrap.reason.tag existing records.
func (c *Client) GetStockScrapReasonTags(ids []int64) (*StockScrapReasonTags, error) {
	ssrts := &StockScrapReasonTags{}
	if err := c.Read(StockScrapReasonTagModel, ids, nil, ssrts); err != nil {
		return nil, err
	}
	return ssrts, nil
}

// FindStockScrapReasonTag finds stock.scrap.reason.tag record by querying it with criteria.
func (c *Client) FindStockScrapReasonTag(criteria *Criteria) (*StockScrapReasonTag, error) {
	ssrts := &StockScrapReasonTags{}
	if err := c.SearchRead(StockScrapReasonTagModel, criteria, NewOptions().Limit(1), ssrts); err != nil {
		return nil, err
	}
	return &((*ssrts)[0]), nil
}

// FindStockScrapReasonTags finds stock.scrap.reason.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockScrapReasonTags(criteria *Criteria, options *Options) (*StockScrapReasonTags, error) {
	ssrts := &StockScrapReasonTags{}
	if err := c.SearchRead(StockScrapReasonTagModel, criteria, options, ssrts); err != nil {
		return nil, err
	}
	return ssrts, nil
}

// FindStockScrapReasonTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockScrapReasonTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockScrapReasonTagModel, criteria, options)
}

// FindStockScrapReasonTagId finds record id by querying it with criteria.
func (c *Client) FindStockScrapReasonTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockScrapReasonTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
