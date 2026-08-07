package odoo

// StockReference represents stock.reference model.
type StockReference struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	MoveIds     *Relation `xmlrpc:"move_ids,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	PickingIds  *Relation `xmlrpc:"picking_ids,omitempty"`
	PurchaseIds *Relation `xmlrpc:"purchase_ids,omitempty"`
	SaleIds     *Relation `xmlrpc:"sale_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockReferences represents array of stock.reference model.
type StockReferences []StockReference

// StockReferenceModel is the odoo model name.
const StockReferenceModel = "stock.reference"

// Many2One convert StockReference to *Many2One.
func (sr *StockReference) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateStockReference creates a new stock.reference model and returns its id.
func (c *Client) CreateStockReference(sr *StockReference) (int64, error) {
	ids, err := c.CreateStockReferences([]*StockReference{sr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockReference creates a new stock.reference model and returns its id.
func (c *Client) CreateStockReferences(srs []*StockReference) ([]int64, error) {
	var vv []interface{}
	for _, v := range srs {
		vv = append(vv, v)
	}
	return c.Create(StockReferenceModel, vv, nil)
}

// UpdateStockReference updates an existing stock.reference record.
func (c *Client) UpdateStockReference(sr *StockReference) error {
	return c.UpdateStockReferences([]int64{sr.Id.Get()}, sr)
}

// UpdateStockReferences updates existing stock.reference records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateStockReferences(ids []int64, sr *StockReference) error {
	return c.Update(StockReferenceModel, ids, sr, nil)
}

// DeleteStockReference deletes an existing stock.reference record.
func (c *Client) DeleteStockReference(id int64) error {
	return c.DeleteStockReferences([]int64{id})
}

// DeleteStockReferences deletes existing stock.reference records.
func (c *Client) DeleteStockReferences(ids []int64) error {
	return c.Delete(StockReferenceModel, ids)
}

// GetStockReference gets stock.reference existing record.
func (c *Client) GetStockReference(id int64) (*StockReference, error) {
	srs, err := c.GetStockReferences([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*srs)[0]), nil
}

// GetStockReferences gets stock.reference existing records.
func (c *Client) GetStockReferences(ids []int64) (*StockReferences, error) {
	srs := &StockReferences{}
	if err := c.Read(StockReferenceModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockReference finds stock.reference record by querying it with criteria.
func (c *Client) FindStockReference(criteria *Criteria) (*StockReference, error) {
	srs := &StockReferences{}
	if err := c.SearchRead(StockReferenceModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	return &((*srs)[0]), nil
}

// FindStockReferences finds stock.reference records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReferences(criteria *Criteria, options *Options) (*StockReferences, error) {
	srs := &StockReferences{}
	if err := c.SearchRead(StockReferenceModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockReferenceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReferenceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockReferenceModel, criteria, options)
}

// FindStockReferenceId finds record id by querying it with criteria.
func (c *Client) FindStockReferenceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockReferenceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
