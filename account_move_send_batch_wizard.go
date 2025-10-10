package odoo

// AccountMoveSendBatchWizard represents account.move.send.batch.wizard model.
type AccountMoveSendBatchWizard struct {
	Alerts           interface{} `xmlrpc:"alerts,omitempty"`
	CreateDate       *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String     `xmlrpc:"display_name,omitempty"`
	Id               *Int        `xmlrpc:"id,omitempty"`
	MoveIds          *Relation   `xmlrpc:"move_ids,omitempty"`
	SendByPostStamps *Int        `xmlrpc:"send_by_post_stamps,omitempty"`
	SummaryData      interface{} `xmlrpc:"summary_data,omitempty"`
	WriteDate        *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountMoveSendBatchWizards represents array of account.move.send.batch.wizard model.
type AccountMoveSendBatchWizards []AccountMoveSendBatchWizard

// AccountMoveSendBatchWizardModel is the odoo model name.
const AccountMoveSendBatchWizardModel = "account.move.send.batch.wizard"

// Many2One convert AccountMoveSendBatchWizard to *Many2One.
func (amsbw *AccountMoveSendBatchWizard) Many2One() *Many2One {
	return NewMany2One(amsbw.Id.Get(), "")
}

// CreateAccountMoveSendBatchWizard creates a new account.move.send.batch.wizard model and returns its id.
func (c *Client) CreateAccountMoveSendBatchWizard(amsbw *AccountMoveSendBatchWizard) (int64, error) {
	ids, err := c.CreateAccountMoveSendBatchWizards([]*AccountMoveSendBatchWizard{amsbw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMoveSendBatchWizard creates a new account.move.send.batch.wizard model and returns its id.
func (c *Client) CreateAccountMoveSendBatchWizards(amsbws []*AccountMoveSendBatchWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range amsbws {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveSendBatchWizardModel, vv, nil)
}

// UpdateAccountMoveSendBatchWizard updates an existing account.move.send.batch.wizard record.
func (c *Client) UpdateAccountMoveSendBatchWizard(amsbw *AccountMoveSendBatchWizard) error {
	return c.UpdateAccountMoveSendBatchWizards([]int64{amsbw.Id.Get()}, amsbw)
}

// UpdateAccountMoveSendBatchWizards updates existing account.move.send.batch.wizard records.
// All records (represented by ids) will be updated by amsbw values.
func (c *Client) UpdateAccountMoveSendBatchWizards(ids []int64, amsbw *AccountMoveSendBatchWizard) error {
	return c.Update(AccountMoveSendBatchWizardModel, ids, amsbw, nil)
}

// DeleteAccountMoveSendBatchWizard deletes an existing account.move.send.batch.wizard record.
func (c *Client) DeleteAccountMoveSendBatchWizard(id int64) error {
	return c.DeleteAccountMoveSendBatchWizards([]int64{id})
}

// DeleteAccountMoveSendBatchWizards deletes existing account.move.send.batch.wizard records.
func (c *Client) DeleteAccountMoveSendBatchWizards(ids []int64) error {
	return c.Delete(AccountMoveSendBatchWizardModel, ids)
}

// GetAccountMoveSendBatchWizard gets account.move.send.batch.wizard existing record.
func (c *Client) GetAccountMoveSendBatchWizard(id int64) (*AccountMoveSendBatchWizard, error) {
	amsbws, err := c.GetAccountMoveSendBatchWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amsbws)[0]), nil
}

// GetAccountMoveSendBatchWizards gets account.move.send.batch.wizard existing records.
func (c *Client) GetAccountMoveSendBatchWizards(ids []int64) (*AccountMoveSendBatchWizards, error) {
	amsbws := &AccountMoveSendBatchWizards{}
	if err := c.Read(AccountMoveSendBatchWizardModel, ids, nil, amsbws); err != nil {
		return nil, err
	}
	return amsbws, nil
}

// FindAccountMoveSendBatchWizard finds account.move.send.batch.wizard record by querying it with criteria.
func (c *Client) FindAccountMoveSendBatchWizard(criteria *Criteria) (*AccountMoveSendBatchWizard, error) {
	amsbws := &AccountMoveSendBatchWizards{}
	if err := c.SearchRead(AccountMoveSendBatchWizardModel, criteria, NewOptions().Limit(1), amsbws); err != nil {
		return nil, err
	}
	return &((*amsbws)[0]), nil
}

// FindAccountMoveSendBatchWizards finds account.move.send.batch.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveSendBatchWizards(criteria *Criteria, options *Options) (*AccountMoveSendBatchWizards, error) {
	amsbws := &AccountMoveSendBatchWizards{}
	if err := c.SearchRead(AccountMoveSendBatchWizardModel, criteria, options, amsbws); err != nil {
		return nil, err
	}
	return amsbws, nil
}

// FindAccountMoveSendBatchWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveSendBatchWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMoveSendBatchWizardModel, criteria, options)
}

// FindAccountMoveSendBatchWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveSendBatchWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveSendBatchWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
