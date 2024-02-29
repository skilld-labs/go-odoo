package odoo

// DecimalPrecision represents decimal.precision model.
type DecimalPrecision struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Digits      *Int      `xmlrpc:"digits,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DecimalPrecisions represents array of decimal.precision model.
type DecimalPrecisions []DecimalPrecision

// DecimalPrecisionModel is the odoo model name.
const DecimalPrecisionModel = "decimal.precision"

// Many2One convert DecimalPrecision to *Many2One.
func (dp *DecimalPrecision) Many2One() *Many2One {
	return NewMany2One(dp.Id.Get(), "")
}

// CreateDecimalPrecision creates a new decimal.precision model and returns its id.
func (c *Client) CreateDecimalPrecision(dp *DecimalPrecision) (int64, error) {
	ids, err := c.CreateDecimalPrecisions([]*DecimalPrecision{dp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDecimalPrecisions creates a new decimal.precision model and returns its id.
func (c *Client) CreateDecimalPrecisions(dps []*DecimalPrecision) ([]int64, error) {
	var vv []interface{}
	for _, v := range dps {
		vv = append(vv, v)
	}
	return c.Create(DecimalPrecisionModel, vv, nil)
}

// UpdateDecimalPrecision updates an existing decimal.precision record.
func (c *Client) UpdateDecimalPrecision(dp *DecimalPrecision) error {
	return c.UpdateDecimalPrecisions([]int64{dp.Id.Get()}, dp)
}

// UpdateDecimalPrecisions updates existing decimal.precision records.
// All records (represented by ids) will be updated by dp values.
func (c *Client) UpdateDecimalPrecisions(ids []int64, dp *DecimalPrecision) error {
	return c.Update(DecimalPrecisionModel, ids, dp, nil)
}

// DeleteDecimalPrecision deletes an existing decimal.precision record.
func (c *Client) DeleteDecimalPrecision(id int64) error {
	return c.DeleteDecimalPrecisions([]int64{id})
}

// DeleteDecimalPrecisions deletes existing decimal.precision records.
func (c *Client) DeleteDecimalPrecisions(ids []int64) error {
	return c.Delete(DecimalPrecisionModel, ids)
}

// GetDecimalPrecision gets decimal.precision existing record.
func (c *Client) GetDecimalPrecision(id int64) (*DecimalPrecision, error) {
	dps, err := c.GetDecimalPrecisions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dps)[0]), nil
}

// GetDecimalPrecisions gets decimal.precision existing records.
func (c *Client) GetDecimalPrecisions(ids []int64) (*DecimalPrecisions, error) {
	dps := &DecimalPrecisions{}
	if err := c.Read(DecimalPrecisionModel, ids, nil, dps); err != nil {
		return nil, err
	}
	return dps, nil
}

// FindDecimalPrecision finds decimal.precision record by querying it with criteria.
func (c *Client) FindDecimalPrecision(criteria *Criteria) (*DecimalPrecision, error) {
	dps := &DecimalPrecisions{}
	if err := c.SearchRead(DecimalPrecisionModel, criteria, NewOptions().Limit(1), dps); err != nil {
		return nil, err
	}
	return &((*dps)[0]), nil
}

// FindDecimalPrecisions finds decimal.precision records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDecimalPrecisions(criteria *Criteria, options *Options) (*DecimalPrecisions, error) {
	dps := &DecimalPrecisions{}
	if err := c.SearchRead(DecimalPrecisionModel, criteria, options, dps); err != nil {
		return nil, err
	}
	return dps, nil
}

// FindDecimalPrecisionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDecimalPrecisionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DecimalPrecisionModel, criteria, options)
}

// FindDecimalPrecisionId finds record id by querying it with criteria.
func (c *Client) FindDecimalPrecisionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DecimalPrecisionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
