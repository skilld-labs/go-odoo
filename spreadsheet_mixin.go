package odoo

// SpreadsheetMixin represents spreadsheet.mixin model.
type SpreadsheetMixin struct {
	DisplayName           *String `xmlrpc:"display_name,omitempty"`
	Id                    *Int    `xmlrpc:"id,omitempty"`
	SpreadsheetBinaryData *String `xmlrpc:"spreadsheet_binary_data,omitempty"`
	SpreadsheetData       *String `xmlrpc:"spreadsheet_data,omitempty"`
	SpreadsheetFileName   *String `xmlrpc:"spreadsheet_file_name,omitempty"`
	Thumbnail             *String `xmlrpc:"thumbnail,omitempty"`
}

// SpreadsheetMixins represents array of spreadsheet.mixin model.
type SpreadsheetMixins []SpreadsheetMixin

// SpreadsheetMixinModel is the odoo model name.
const SpreadsheetMixinModel = "spreadsheet.mixin"

// Many2One convert SpreadsheetMixin to *Many2One.
func (sm *SpreadsheetMixin) Many2One() *Many2One {
	return NewMany2One(sm.Id.Get(), "")
}

// CreateSpreadsheetMixin creates a new spreadsheet.mixin model and returns its id.
func (c *Client) CreateSpreadsheetMixin(sm *SpreadsheetMixin) (int64, error) {
	ids, err := c.CreateSpreadsheetMixins([]*SpreadsheetMixin{sm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSpreadsheetMixin creates a new spreadsheet.mixin model and returns its id.
func (c *Client) CreateSpreadsheetMixins(sms []*SpreadsheetMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range sms {
		vv = append(vv, v)
	}
	return c.Create(SpreadsheetMixinModel, vv, nil)
}

// UpdateSpreadsheetMixin updates an existing spreadsheet.mixin record.
func (c *Client) UpdateSpreadsheetMixin(sm *SpreadsheetMixin) error {
	return c.UpdateSpreadsheetMixins([]int64{sm.Id.Get()}, sm)
}

// UpdateSpreadsheetMixins updates existing spreadsheet.mixin records.
// All records (represented by ids) will be updated by sm values.
func (c *Client) UpdateSpreadsheetMixins(ids []int64, sm *SpreadsheetMixin) error {
	return c.Update(SpreadsheetMixinModel, ids, sm, nil)
}

// DeleteSpreadsheetMixin deletes an existing spreadsheet.mixin record.
func (c *Client) DeleteSpreadsheetMixin(id int64) error {
	return c.DeleteSpreadsheetMixins([]int64{id})
}

// DeleteSpreadsheetMixins deletes existing spreadsheet.mixin records.
func (c *Client) DeleteSpreadsheetMixins(ids []int64) error {
	return c.Delete(SpreadsheetMixinModel, ids)
}

// GetSpreadsheetMixin gets spreadsheet.mixin existing record.
func (c *Client) GetSpreadsheetMixin(id int64) (*SpreadsheetMixin, error) {
	sms, err := c.GetSpreadsheetMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sms)[0]), nil
}

// GetSpreadsheetMixins gets spreadsheet.mixin existing records.
func (c *Client) GetSpreadsheetMixins(ids []int64) (*SpreadsheetMixins, error) {
	sms := &SpreadsheetMixins{}
	if err := c.Read(SpreadsheetMixinModel, ids, nil, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindSpreadsheetMixin finds spreadsheet.mixin record by querying it with criteria.
func (c *Client) FindSpreadsheetMixin(criteria *Criteria) (*SpreadsheetMixin, error) {
	sms := &SpreadsheetMixins{}
	if err := c.SearchRead(SpreadsheetMixinModel, criteria, NewOptions().Limit(1), sms); err != nil {
		return nil, err
	}
	return &((*sms)[0]), nil
}

// FindSpreadsheetMixins finds spreadsheet.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetMixins(criteria *Criteria, options *Options) (*SpreadsheetMixins, error) {
	sms := &SpreadsheetMixins{}
	if err := c.SearchRead(SpreadsheetMixinModel, criteria, options, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindSpreadsheetMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SpreadsheetMixinModel, criteria, options)
}

// FindSpreadsheetMixinId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
