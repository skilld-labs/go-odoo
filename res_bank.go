package odoo

// ResBank represents res.bank model.
type ResBank struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Active      *Bool     `xmlrpc:"active,omitempty"`
	Bic         *String   `xmlrpc:"bic,omitempty"`
	City        *String   `xmlrpc:"city,omitempty"`
	Country     *Many2One `xmlrpc:"country,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Email       *String   `xmlrpc:"email,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Phone       *String   `xmlrpc:"phone,omitempty"`
	State       *Many2One `xmlrpc:"state,omitempty"`
	Street      *String   `xmlrpc:"street,omitempty"`
	Street2     *String   `xmlrpc:"street2,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
	Zip         *String   `xmlrpc:"zip,omitempty"`
}

// ResBanks represents array of res.bank model.
type ResBanks []ResBank

// ResBankModel is the odoo model name.
const ResBankModel = "res.bank"

// Many2One convert ResBank to *Many2One.
func (rb *ResBank) Many2One() *Many2One {
	return NewMany2One(rb.Id.Get(), "")
}

// CreateResBank creates a new res.bank model and returns its id.
func (c *Client) CreateResBank(rb *ResBank) (int64, error) {
	ids, err := c.CreateResBanks([]*ResBank{rb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResBanks creates a new res.bank model and returns its id.
func (c *Client) CreateResBanks(rbs []*ResBank) ([]int64, error) {
	var vv []interface{}
	for _, v := range rbs {
		vv = append(vv, v)
	}
	return c.Create(ResBankModel, vv, nil)
}

// UpdateResBank updates an existing res.bank record.
func (c *Client) UpdateResBank(rb *ResBank) error {
	return c.UpdateResBanks([]int64{rb.Id.Get()}, rb)
}

// UpdateResBanks updates existing res.bank records.
// All records (represented by ids) will be updated by rb values.
func (c *Client) UpdateResBanks(ids []int64, rb *ResBank) error {
	return c.Update(ResBankModel, ids, rb, nil)
}

// DeleteResBank deletes an existing res.bank record.
func (c *Client) DeleteResBank(id int64) error {
	return c.DeleteResBanks([]int64{id})
}

// DeleteResBanks deletes existing res.bank records.
func (c *Client) DeleteResBanks(ids []int64) error {
	return c.Delete(ResBankModel, ids)
}

// GetResBank gets res.bank existing record.
func (c *Client) GetResBank(id int64) (*ResBank, error) {
	rbs, err := c.GetResBanks([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rbs)[0]), nil
}

// GetResBanks gets res.bank existing records.
func (c *Client) GetResBanks(ids []int64) (*ResBanks, error) {
	rbs := &ResBanks{}
	if err := c.Read(ResBankModel, ids, nil, rbs); err != nil {
		return nil, err
	}
	return rbs, nil
}

// FindResBank finds res.bank record by querying it with criteria.
func (c *Client) FindResBank(criteria *Criteria) (*ResBank, error) {
	rbs := &ResBanks{}
	if err := c.SearchRead(ResBankModel, criteria, NewOptions().Limit(1), rbs); err != nil {
		return nil, err
	}
	return &((*rbs)[0]), nil
}

// FindResBanks finds res.bank records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResBanks(criteria *Criteria, options *Options) (*ResBanks, error) {
	rbs := &ResBanks{}
	if err := c.SearchRead(ResBankModel, criteria, options, rbs); err != nil {
		return nil, err
	}
	return rbs, nil
}

// FindResBankIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResBankIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResBankModel, criteria, options)
}

// FindResBankId finds record id by querying it with criteria.
func (c *Client) FindResBankId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResBankModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
