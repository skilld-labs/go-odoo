package odoo

// AccountOpening represents account.opening model.
type AccountOpening struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId          *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId         *Many2One `xmlrpc:"currency_id,omptempty"`
	Date               *Time     `xmlrpc:"date,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	JournalId          *Many2One `xmlrpc:"journal_id,omptempty"`
	OpeningMoveId      *Many2One `xmlrpc:"opening_move_id,omptempty"`
	OpeningMoveLineIds *Relation `xmlrpc:"opening_move_line_ids,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountOpenings represents array of account.opening model.
type AccountOpenings []AccountOpening

// AccountOpeningModel is the odoo model name.
const AccountOpeningModel = "account.opening"

// Many2One convert AccountOpening to *Many2One.
func (ao *AccountOpening) Many2One() *Many2One {
	return NewMany2One(ao.Id.Get(), "")
}

// CreateAccountOpening creates a new account.opening model and returns its id.
func (c *Client) CreateAccountOpening(ao *AccountOpening) (int64, error) {
	ids, err := c.CreateAccountOpenings([]*AccountOpening{ao})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountOpening creates a new account.opening model and returns its id.
func (c *Client) CreateAccountOpenings(aos []*AccountOpening) ([]int64, error) {
	var vv []interface{}
	for _, v := range aos {
		vv = append(vv, v)
	}
	return c.Create(AccountOpeningModel, vv, nil)
}

// UpdateAccountOpening updates an existing account.opening record.
func (c *Client) UpdateAccountOpening(ao *AccountOpening) error {
	return c.UpdateAccountOpenings([]int64{ao.Id.Get()}, ao)
}

// UpdateAccountOpenings updates existing account.opening records.
// All records (represented by ids) will be updated by ao values.
func (c *Client) UpdateAccountOpenings(ids []int64, ao *AccountOpening) error {
	return c.Update(AccountOpeningModel, ids, ao, nil)
}

// DeleteAccountOpening deletes an existing account.opening record.
func (c *Client) DeleteAccountOpening(id int64) error {
	return c.DeleteAccountOpenings([]int64{id})
}

// DeleteAccountOpenings deletes existing account.opening records.
func (c *Client) DeleteAccountOpenings(ids []int64) error {
	return c.Delete(AccountOpeningModel, ids)
}

// GetAccountOpening gets account.opening existing record.
func (c *Client) GetAccountOpening(id int64) (*AccountOpening, error) {
	aos, err := c.GetAccountOpenings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aos)[0]), nil
}

// GetAccountOpenings gets account.opening existing records.
func (c *Client) GetAccountOpenings(ids []int64) (*AccountOpenings, error) {
	aos := &AccountOpenings{}
	if err := c.Read(AccountOpeningModel, ids, nil, aos); err != nil {
		return nil, err
	}
	return aos, nil
}

// FindAccountOpening finds account.opening record by querying it with criteria.
func (c *Client) FindAccountOpening(criteria *Criteria) (*AccountOpening, error) {
	aos := &AccountOpenings{}
	if err := c.SearchRead(AccountOpeningModel, criteria, NewOptions().Limit(1), aos); err != nil {
		return nil, err
	}
	return &((*aos)[0]), nil
}

// FindAccountOpenings finds account.opening records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOpenings(criteria *Criteria, options *Options) (*AccountOpenings, error) {
	aos := &AccountOpenings{}
	if err := c.SearchRead(AccountOpeningModel, criteria, options, aos); err != nil {
		return nil, err
	}
	return aos, nil
}

// FindAccountOpeningIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOpeningIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountOpeningModel, criteria, options)
}

// FindAccountOpeningId finds record id by querying it with criteria.
func (c *Client) FindAccountOpeningId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountOpeningModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
