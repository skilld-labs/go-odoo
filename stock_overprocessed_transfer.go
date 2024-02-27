package odoo

// StockOverprocessedTransfer represents stock.overprocessed.transfer model.
type StockOverprocessedTransfer struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	OverprocessedProductName *String   `xmlrpc:"overprocessed_product_name,omitempty"`
	PickingId                *Many2One `xmlrpc:"picking_id,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockOverprocessedTransfers represents array of stock.overprocessed.transfer model.
type StockOverprocessedTransfers []StockOverprocessedTransfer

// StockOverprocessedTransferModel is the odoo model name.
const StockOverprocessedTransferModel = "stock.overprocessed.transfer"

// Many2One convert StockOverprocessedTransfer to *Many2One.
func (sot *StockOverprocessedTransfer) Many2One() *Many2One {
	return NewMany2One(sot.Id.Get(), "")
}

// CreateStockOverprocessedTransfer creates a new stock.overprocessed.transfer model and returns its id.
func (c *Client) CreateStockOverprocessedTransfer(sot *StockOverprocessedTransfer) (int64, error) {
	ids, err := c.CreateStockOverprocessedTransfers([]*StockOverprocessedTransfer{sot})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockOverprocessedTransfers creates a new stock.overprocessed.transfer model and returns its id.
func (c *Client) CreateStockOverprocessedTransfers(sots []*StockOverprocessedTransfer) ([]int64, error) {
	var vv []interface{}
	for _, v := range sots {
		vv = append(vv, v)
	}
	return c.Create(StockOverprocessedTransferModel, vv, nil)
}

// UpdateStockOverprocessedTransfer updates an existing stock.overprocessed.transfer record.
func (c *Client) UpdateStockOverprocessedTransfer(sot *StockOverprocessedTransfer) error {
	return c.UpdateStockOverprocessedTransfers([]int64{sot.Id.Get()}, sot)
}

// UpdateStockOverprocessedTransfers updates existing stock.overprocessed.transfer records.
// All records (represented by ids) will be updated by sot values.
func (c *Client) UpdateStockOverprocessedTransfers(ids []int64, sot *StockOverprocessedTransfer) error {
	return c.Update(StockOverprocessedTransferModel, ids, sot, nil)
}

// DeleteStockOverprocessedTransfer deletes an existing stock.overprocessed.transfer record.
func (c *Client) DeleteStockOverprocessedTransfer(id int64) error {
	return c.DeleteStockOverprocessedTransfers([]int64{id})
}

// DeleteStockOverprocessedTransfers deletes existing stock.overprocessed.transfer records.
func (c *Client) DeleteStockOverprocessedTransfers(ids []int64) error {
	return c.Delete(StockOverprocessedTransferModel, ids)
}

// GetStockOverprocessedTransfer gets stock.overprocessed.transfer existing record.
func (c *Client) GetStockOverprocessedTransfer(id int64) (*StockOverprocessedTransfer, error) {
	sots, err := c.GetStockOverprocessedTransfers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sots)[0]), nil
}

// GetStockOverprocessedTransfers gets stock.overprocessed.transfer existing records.
func (c *Client) GetStockOverprocessedTransfers(ids []int64) (*StockOverprocessedTransfers, error) {
	sots := &StockOverprocessedTransfers{}
	if err := c.Read(StockOverprocessedTransferModel, ids, nil, sots); err != nil {
		return nil, err
	}
	return sots, nil
}

// FindStockOverprocessedTransfer finds stock.overprocessed.transfer record by querying it with criteria.
func (c *Client) FindStockOverprocessedTransfer(criteria *Criteria) (*StockOverprocessedTransfer, error) {
	sots := &StockOverprocessedTransfers{}
	if err := c.SearchRead(StockOverprocessedTransferModel, criteria, NewOptions().Limit(1), sots); err != nil {
		return nil, err
	}
	return &((*sots)[0]), nil
}

// FindStockOverprocessedTransfers finds stock.overprocessed.transfer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockOverprocessedTransfers(criteria *Criteria, options *Options) (*StockOverprocessedTransfers, error) {
	sots := &StockOverprocessedTransfers{}
	if err := c.SearchRead(StockOverprocessedTransferModel, criteria, options, sots); err != nil {
		return nil, err
	}
	return sots, nil
}

// FindStockOverprocessedTransferIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockOverprocessedTransferIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockOverprocessedTransferModel, criteria, options)
}

// FindStockOverprocessedTransferId finds record id by querying it with criteria.
func (c *Client) FindStockOverprocessedTransferId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockOverprocessedTransferModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
