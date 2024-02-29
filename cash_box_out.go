package odoo

// CashBoxOut represents cash.box.out model.
type CashBoxOut struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Amount      *Float    `xmlrpc:"amount,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CashBoxOuts represents array of cash.box.out model.
type CashBoxOuts []CashBoxOut

// CashBoxOutModel is the odoo model name.
const CashBoxOutModel = "cash.box.out"

// Many2One convert CashBoxOut to *Many2One.
func (cbo *CashBoxOut) Many2One() *Many2One {
	return NewMany2One(cbo.Id.Get(), "")
}

// CreateCashBoxOut creates a new cash.box.out model and returns its id.
func (c *Client) CreateCashBoxOut(cbo *CashBoxOut) (int64, error) {
	ids, err := c.CreateCashBoxOuts([]*CashBoxOut{cbo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCashBoxOuts creates a new cash.box.out model and returns its id.
func (c *Client) CreateCashBoxOuts(cbos []*CashBoxOut) ([]int64, error) {
	var vv []interface{}
	for _, v := range cbos {
		vv = append(vv, v)
	}
	return c.Create(CashBoxOutModel, vv, nil)
}

// UpdateCashBoxOut updates an existing cash.box.out record.
func (c *Client) UpdateCashBoxOut(cbo *CashBoxOut) error {
	return c.UpdateCashBoxOuts([]int64{cbo.Id.Get()}, cbo)
}

// UpdateCashBoxOuts updates existing cash.box.out records.
// All records (represented by ids) will be updated by cbo values.
func (c *Client) UpdateCashBoxOuts(ids []int64, cbo *CashBoxOut) error {
	return c.Update(CashBoxOutModel, ids, cbo, nil)
}

// DeleteCashBoxOut deletes an existing cash.box.out record.
func (c *Client) DeleteCashBoxOut(id int64) error {
	return c.DeleteCashBoxOuts([]int64{id})
}

// DeleteCashBoxOuts deletes existing cash.box.out records.
func (c *Client) DeleteCashBoxOuts(ids []int64) error {
	return c.Delete(CashBoxOutModel, ids)
}

// GetCashBoxOut gets cash.box.out existing record.
func (c *Client) GetCashBoxOut(id int64) (*CashBoxOut, error) {
	cbos, err := c.GetCashBoxOuts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cbos)[0]), nil
}

// GetCashBoxOuts gets cash.box.out existing records.
func (c *Client) GetCashBoxOuts(ids []int64) (*CashBoxOuts, error) {
	cbos := &CashBoxOuts{}
	if err := c.Read(CashBoxOutModel, ids, nil, cbos); err != nil {
		return nil, err
	}
	return cbos, nil
}

// FindCashBoxOut finds cash.box.out record by querying it with criteria.
func (c *Client) FindCashBoxOut(criteria *Criteria) (*CashBoxOut, error) {
	cbos := &CashBoxOuts{}
	if err := c.SearchRead(CashBoxOutModel, criteria, NewOptions().Limit(1), cbos); err != nil {
		return nil, err
	}
	return &((*cbos)[0]), nil
}

// FindCashBoxOuts finds cash.box.out records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCashBoxOuts(criteria *Criteria, options *Options) (*CashBoxOuts, error) {
	cbos := &CashBoxOuts{}
	if err := c.SearchRead(CashBoxOutModel, criteria, options, cbos); err != nil {
		return nil, err
	}
	return cbos, nil
}

// FindCashBoxOutIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCashBoxOutIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CashBoxOutModel, criteria, options)
}

// FindCashBoxOutId finds record id by querying it with criteria.
func (c *Client) FindCashBoxOutId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CashBoxOutModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
