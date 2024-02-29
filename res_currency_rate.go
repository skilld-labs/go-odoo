package odoo

// ResCurrencyRate represents res.currency.rate model.
type ResCurrencyRate struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId  *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *Time     `xmlrpc:"name,omitempty"`
	Rate        *Float    `xmlrpc:"rate,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResCurrencyRates represents array of res.currency.rate model.
type ResCurrencyRates []ResCurrencyRate

// ResCurrencyRateModel is the odoo model name.
const ResCurrencyRateModel = "res.currency.rate"

// Many2One convert ResCurrencyRate to *Many2One.
func (rcr *ResCurrencyRate) Many2One() *Many2One {
	return NewMany2One(rcr.Id.Get(), "")
}

// CreateResCurrencyRate creates a new res.currency.rate model and returns its id.
func (c *Client) CreateResCurrencyRate(rcr *ResCurrencyRate) (int64, error) {
	ids, err := c.CreateResCurrencyRates([]*ResCurrencyRate{rcr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResCurrencyRates creates a new res.currency.rate model and returns its id.
func (c *Client) CreateResCurrencyRates(rcrs []*ResCurrencyRate) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcrs {
		vv = append(vv, v)
	}
	return c.Create(ResCurrencyRateModel, vv, nil)
}

// UpdateResCurrencyRate updates an existing res.currency.rate record.
func (c *Client) UpdateResCurrencyRate(rcr *ResCurrencyRate) error {
	return c.UpdateResCurrencyRates([]int64{rcr.Id.Get()}, rcr)
}

// UpdateResCurrencyRates updates existing res.currency.rate records.
// All records (represented by ids) will be updated by rcr values.
func (c *Client) UpdateResCurrencyRates(ids []int64, rcr *ResCurrencyRate) error {
	return c.Update(ResCurrencyRateModel, ids, rcr, nil)
}

// DeleteResCurrencyRate deletes an existing res.currency.rate record.
func (c *Client) DeleteResCurrencyRate(id int64) error {
	return c.DeleteResCurrencyRates([]int64{id})
}

// DeleteResCurrencyRates deletes existing res.currency.rate records.
func (c *Client) DeleteResCurrencyRates(ids []int64) error {
	return c.Delete(ResCurrencyRateModel, ids)
}

// GetResCurrencyRate gets res.currency.rate existing record.
func (c *Client) GetResCurrencyRate(id int64) (*ResCurrencyRate, error) {
	rcrs, err := c.GetResCurrencyRates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcrs)[0]), nil
}

// GetResCurrencyRates gets res.currency.rate existing records.
func (c *Client) GetResCurrencyRates(ids []int64) (*ResCurrencyRates, error) {
	rcrs := &ResCurrencyRates{}
	if err := c.Read(ResCurrencyRateModel, ids, nil, rcrs); err != nil {
		return nil, err
	}
	return rcrs, nil
}

// FindResCurrencyRate finds res.currency.rate record by querying it with criteria.
func (c *Client) FindResCurrencyRate(criteria *Criteria) (*ResCurrencyRate, error) {
	rcrs := &ResCurrencyRates{}
	if err := c.SearchRead(ResCurrencyRateModel, criteria, NewOptions().Limit(1), rcrs); err != nil {
		return nil, err
	}
	return &((*rcrs)[0]), nil
}

// FindResCurrencyRates finds res.currency.rate records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCurrencyRates(criteria *Criteria, options *Options) (*ResCurrencyRates, error) {
	rcrs := &ResCurrencyRates{}
	if err := c.SearchRead(ResCurrencyRateModel, criteria, options, rcrs); err != nil {
		return nil, err
	}
	return rcrs, nil
}

// FindResCurrencyRateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCurrencyRateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResCurrencyRateModel, criteria, options)
}

// FindResCurrencyRateId finds record id by querying it with criteria.
func (c *Client) FindResCurrencyRateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCurrencyRateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
