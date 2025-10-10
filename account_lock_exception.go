package odoo

// AccountLockException represents account.lock_exception model.
type AccountLockException struct {
	Active             *Bool      `xmlrpc:"active,omitempty"`
	CompanyId          *Many2One  `xmlrpc:"company_id,omitempty"`
	CompanyLockDate    *Time      `xmlrpc:"company_lock_date,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	EndDatetime        *Time      `xmlrpc:"end_datetime,omitempty"`
	FiscalyearLockDate *Time      `xmlrpc:"fiscalyear_lock_date,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	LockDate           *Time      `xmlrpc:"lock_date,omitempty"`
	LockDateField      *Selection `xmlrpc:"lock_date_field,omitempty"`
	PurchaseLockDate   *Time      `xmlrpc:"purchase_lock_date,omitempty"`
	Reason             *String    `xmlrpc:"reason,omitempty"`
	SaleLockDate       *Time      `xmlrpc:"sale_lock_date,omitempty"`
	State              *Selection `xmlrpc:"state,omitempty"`
	TaxLockDate        *Time      `xmlrpc:"tax_lock_date,omitempty"`
	UserId             *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountLockExceptions represents array of account.lock_exception model.
type AccountLockExceptions []AccountLockException

// AccountLockExceptionModel is the odoo model name.
const AccountLockExceptionModel = "account.lock_exception"

// Many2One convert AccountLockException to *Many2One.
func (al *AccountLockException) Many2One() *Many2One {
	return NewMany2One(al.Id.Get(), "")
}

// CreateAccountLockException creates a new account.lock_exception model and returns its id.
func (c *Client) CreateAccountLockException(al *AccountLockException) (int64, error) {
	ids, err := c.CreateAccountLockExceptions([]*AccountLockException{al})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountLockException creates a new account.lock_exception model and returns its id.
func (c *Client) CreateAccountLockExceptions(als []*AccountLockException) ([]int64, error) {
	var vv []interface{}
	for _, v := range als {
		vv = append(vv, v)
	}
	return c.Create(AccountLockExceptionModel, vv, nil)
}

// UpdateAccountLockException updates an existing account.lock_exception record.
func (c *Client) UpdateAccountLockException(al *AccountLockException) error {
	return c.UpdateAccountLockExceptions([]int64{al.Id.Get()}, al)
}

// UpdateAccountLockExceptions updates existing account.lock_exception records.
// All records (represented by ids) will be updated by al values.
func (c *Client) UpdateAccountLockExceptions(ids []int64, al *AccountLockException) error {
	return c.Update(AccountLockExceptionModel, ids, al, nil)
}

// DeleteAccountLockException deletes an existing account.lock_exception record.
func (c *Client) DeleteAccountLockException(id int64) error {
	return c.DeleteAccountLockExceptions([]int64{id})
}

// DeleteAccountLockExceptions deletes existing account.lock_exception records.
func (c *Client) DeleteAccountLockExceptions(ids []int64) error {
	return c.Delete(AccountLockExceptionModel, ids)
}

// GetAccountLockException gets account.lock_exception existing record.
func (c *Client) GetAccountLockException(id int64) (*AccountLockException, error) {
	als, err := c.GetAccountLockExceptions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*als)[0]), nil
}

// GetAccountLockExceptions gets account.lock_exception existing records.
func (c *Client) GetAccountLockExceptions(ids []int64) (*AccountLockExceptions, error) {
	als := &AccountLockExceptions{}
	if err := c.Read(AccountLockExceptionModel, ids, nil, als); err != nil {
		return nil, err
	}
	return als, nil
}

// FindAccountLockException finds account.lock_exception record by querying it with criteria.
func (c *Client) FindAccountLockException(criteria *Criteria) (*AccountLockException, error) {
	als := &AccountLockExceptions{}
	if err := c.SearchRead(AccountLockExceptionModel, criteria, NewOptions().Limit(1), als); err != nil {
		return nil, err
	}
	return &((*als)[0]), nil
}

// FindAccountLockExceptions finds account.lock_exception records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLockExceptions(criteria *Criteria, options *Options) (*AccountLockExceptions, error) {
	als := &AccountLockExceptions{}
	if err := c.SearchRead(AccountLockExceptionModel, criteria, options, als); err != nil {
		return nil, err
	}
	return als, nil
}

// FindAccountLockExceptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLockExceptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountLockExceptionModel, criteria, options)
}

// FindAccountLockExceptionId finds record id by querying it with criteria.
func (c *Client) FindAccountLockExceptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountLockExceptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
