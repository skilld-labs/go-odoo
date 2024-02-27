package odoo

// BarcodesBarcodeEventsMixin represents barcodes.barcode_events_mixin model.
type BarcodesBarcodeEventsMixin struct {
	LastUpdate     *Time   `xmlrpc:"__last_update,omitempty"`
	BarcodeScanned *String `xmlrpc:"_barcode_scanned,omitempty"`
	DisplayName    *String `xmlrpc:"display_name,omitempty"`
	Id             *Int    `xmlrpc:"id,omitempty"`
}

// BarcodesBarcodeEventsMixins represents array of barcodes.barcode_events_mixin model.
type BarcodesBarcodeEventsMixins []BarcodesBarcodeEventsMixin

// BarcodesBarcodeEventsMixinModel is the odoo model name.
const BarcodesBarcodeEventsMixinModel = "barcodes.barcode_events_mixin"

// Many2One convert BarcodesBarcodeEventsMixin to *Many2One.
func (bb *BarcodesBarcodeEventsMixin) Many2One() *Many2One {
	return NewMany2One(bb.Id.Get(), "")
}

// CreateBarcodesBarcodeEventsMixin creates a new barcodes.barcode_events_mixin model and returns its id.
func (c *Client) CreateBarcodesBarcodeEventsMixin(bb *BarcodesBarcodeEventsMixin) (int64, error) {
	ids, err := c.CreateBarcodesBarcodeEventsMixins([]*BarcodesBarcodeEventsMixin{bb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBarcodesBarcodeEventsMixins creates a new barcodes.barcode_events_mixin model and returns its id.
func (c *Client) CreateBarcodesBarcodeEventsMixins(bbs []*BarcodesBarcodeEventsMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range bbs {
		vv = append(vv, v)
	}
	return c.Create(BarcodesBarcodeEventsMixinModel, vv, nil)
}

// UpdateBarcodesBarcodeEventsMixin updates an existing barcodes.barcode_events_mixin record.
func (c *Client) UpdateBarcodesBarcodeEventsMixin(bb *BarcodesBarcodeEventsMixin) error {
	return c.UpdateBarcodesBarcodeEventsMixins([]int64{bb.Id.Get()}, bb)
}

// UpdateBarcodesBarcodeEventsMixins updates existing barcodes.barcode_events_mixin records.
// All records (represented by ids) will be updated by bb values.
func (c *Client) UpdateBarcodesBarcodeEventsMixins(ids []int64, bb *BarcodesBarcodeEventsMixin) error {
	return c.Update(BarcodesBarcodeEventsMixinModel, ids, bb, nil)
}

// DeleteBarcodesBarcodeEventsMixin deletes an existing barcodes.barcode_events_mixin record.
func (c *Client) DeleteBarcodesBarcodeEventsMixin(id int64) error {
	return c.DeleteBarcodesBarcodeEventsMixins([]int64{id})
}

// DeleteBarcodesBarcodeEventsMixins deletes existing barcodes.barcode_events_mixin records.
func (c *Client) DeleteBarcodesBarcodeEventsMixins(ids []int64) error {
	return c.Delete(BarcodesBarcodeEventsMixinModel, ids)
}

// GetBarcodesBarcodeEventsMixin gets barcodes.barcode_events_mixin existing record.
func (c *Client) GetBarcodesBarcodeEventsMixin(id int64) (*BarcodesBarcodeEventsMixin, error) {
	bbs, err := c.GetBarcodesBarcodeEventsMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bbs)[0]), nil
}

// GetBarcodesBarcodeEventsMixins gets barcodes.barcode_events_mixin existing records.
func (c *Client) GetBarcodesBarcodeEventsMixins(ids []int64) (*BarcodesBarcodeEventsMixins, error) {
	bbs := &BarcodesBarcodeEventsMixins{}
	if err := c.Read(BarcodesBarcodeEventsMixinModel, ids, nil, bbs); err != nil {
		return nil, err
	}
	return bbs, nil
}

// FindBarcodesBarcodeEventsMixin finds barcodes.barcode_events_mixin record by querying it with criteria.
func (c *Client) FindBarcodesBarcodeEventsMixin(criteria *Criteria) (*BarcodesBarcodeEventsMixin, error) {
	bbs := &BarcodesBarcodeEventsMixins{}
	if err := c.SearchRead(BarcodesBarcodeEventsMixinModel, criteria, NewOptions().Limit(1), bbs); err != nil {
		return nil, err
	}
	return &((*bbs)[0]), nil
}

// FindBarcodesBarcodeEventsMixins finds barcodes.barcode_events_mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBarcodesBarcodeEventsMixins(criteria *Criteria, options *Options) (*BarcodesBarcodeEventsMixins, error) {
	bbs := &BarcodesBarcodeEventsMixins{}
	if err := c.SearchRead(BarcodesBarcodeEventsMixinModel, criteria, options, bbs); err != nil {
		return nil, err
	}
	return bbs, nil
}

// FindBarcodesBarcodeEventsMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBarcodesBarcodeEventsMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BarcodesBarcodeEventsMixinModel, criteria, options)
}

// FindBarcodesBarcodeEventsMixinId finds record id by querying it with criteria.
func (c *Client) FindBarcodesBarcodeEventsMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BarcodesBarcodeEventsMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
