package odoo

// AccountPaymentTermLine represents account.payment.term.line model.
type AccountPaymentTermLine struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Days        *Int       `xmlrpc:"days,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Option      *Selection `xmlrpc:"option,omptempty"`
	PaymentId   *Many2One  `xmlrpc:"payment_id,omptempty"`
	Sequence    *Int       `xmlrpc:"sequence,omptempty"`
	Value       *Selection `xmlrpc:"value,omptempty"`
	ValueAmount *Float     `xmlrpc:"value_amount,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountPaymentTermLines represents array of account.payment.term.line model.
type AccountPaymentTermLines []AccountPaymentTermLine

// AccountPaymentTermLineModel is the odoo model name.
const AccountPaymentTermLineModel = "account.payment.term.line"

// Many2One convert AccountPaymentTermLine to *Many2One.
func (aptl *AccountPaymentTermLine) Many2One() *Many2One {
	return NewMany2One(aptl.Id.Get(), "")
}

// CreateAccountPaymentTermLine creates a new account.payment.term.line model and returns its id.
func (c *Client) CreateAccountPaymentTermLine(aptl *AccountPaymentTermLine) (int64, error) {
	ids, err := c.CreateAccountPaymentTermLines([]*AccountPaymentTermLine{aptl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPaymentTermLine creates a new account.payment.term.line model and returns its id.
func (c *Client) CreateAccountPaymentTermLines(aptls []*AccountPaymentTermLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range aptls {
		vv = append(vv, v)
	}
	return c.Create(AccountPaymentTermLineModel, vv, nil)
}

// UpdateAccountPaymentTermLine updates an existing account.payment.term.line record.
func (c *Client) UpdateAccountPaymentTermLine(aptl *AccountPaymentTermLine) error {
	return c.UpdateAccountPaymentTermLines([]int64{aptl.Id.Get()}, aptl)
}

// UpdateAccountPaymentTermLines updates existing account.payment.term.line records.
// All records (represented by ids) will be updated by aptl values.
func (c *Client) UpdateAccountPaymentTermLines(ids []int64, aptl *AccountPaymentTermLine) error {
	return c.Update(AccountPaymentTermLineModel, ids, aptl, nil)
}

// DeleteAccountPaymentTermLine deletes an existing account.payment.term.line record.
func (c *Client) DeleteAccountPaymentTermLine(id int64) error {
	return c.DeleteAccountPaymentTermLines([]int64{id})
}

// DeleteAccountPaymentTermLines deletes existing account.payment.term.line records.
func (c *Client) DeleteAccountPaymentTermLines(ids []int64) error {
	return c.Delete(AccountPaymentTermLineModel, ids)
}

// GetAccountPaymentTermLine gets account.payment.term.line existing record.
func (c *Client) GetAccountPaymentTermLine(id int64) (*AccountPaymentTermLine, error) {
	aptls, err := c.GetAccountPaymentTermLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aptls)[0]), nil
}

// GetAccountPaymentTermLines gets account.payment.term.line existing records.
func (c *Client) GetAccountPaymentTermLines(ids []int64) (*AccountPaymentTermLines, error) {
	aptls := &AccountPaymentTermLines{}
	if err := c.Read(AccountPaymentTermLineModel, ids, nil, aptls); err != nil {
		return nil, err
	}
	return aptls, nil
}

// FindAccountPaymentTermLine finds account.payment.term.line record by querying it with criteria.
func (c *Client) FindAccountPaymentTermLine(criteria *Criteria) (*AccountPaymentTermLine, error) {
	aptls := &AccountPaymentTermLines{}
	if err := c.SearchRead(AccountPaymentTermLineModel, criteria, NewOptions().Limit(1), aptls); err != nil {
		return nil, err
	}
	return &((*aptls)[0]), nil
}

// FindAccountPaymentTermLines finds account.payment.term.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentTermLines(criteria *Criteria, options *Options) (*AccountPaymentTermLines, error) {
	aptls := &AccountPaymentTermLines{}
	if err := c.SearchRead(AccountPaymentTermLineModel, criteria, options, aptls); err != nil {
		return nil, err
	}
	return aptls, nil
}

// FindAccountPaymentTermLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentTermLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPaymentTermLineModel, criteria, options)
}

// FindAccountPaymentTermLineId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentTermLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentTermLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
