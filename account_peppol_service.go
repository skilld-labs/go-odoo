package odoo

// AccountPeppolService represents account_peppol.service model.
type AccountPeppolService struct {
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	DocumentIdentifier *String   `xmlrpc:"document_identifier,omitempty"`
	DocumentName       *String   `xmlrpc:"document_name,omitempty"`
	Enabled            *Bool     `xmlrpc:"enabled,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	WizardId           *Many2One `xmlrpc:"wizard_id,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountPeppolServices represents array of account_peppol.service model.
type AccountPeppolServices []AccountPeppolService

// AccountPeppolServiceModel is the odoo model name.
const AccountPeppolServiceModel = "account_peppol.service"

// Many2One convert AccountPeppolService to *Many2One.
func (as *AccountPeppolService) Many2One() *Many2One {
	return NewMany2One(as.Id.Get(), "")
}

// CreateAccountPeppolService creates a new account_peppol.service model and returns its id.
func (c *Client) CreateAccountPeppolService(as *AccountPeppolService) (int64, error) {
	ids, err := c.CreateAccountPeppolServices([]*AccountPeppolService{as})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPeppolService creates a new account_peppol.service model and returns its id.
func (c *Client) CreateAccountPeppolServices(ass []*AccountPeppolService) ([]int64, error) {
	var vv []interface{}
	for _, v := range ass {
		vv = append(vv, v)
	}
	return c.Create(AccountPeppolServiceModel, vv, nil)
}

// UpdateAccountPeppolService updates an existing account_peppol.service record.
func (c *Client) UpdateAccountPeppolService(as *AccountPeppolService) error {
	return c.UpdateAccountPeppolServices([]int64{as.Id.Get()}, as)
}

// UpdateAccountPeppolServices updates existing account_peppol.service records.
// All records (represented by ids) will be updated by as values.
func (c *Client) UpdateAccountPeppolServices(ids []int64, as *AccountPeppolService) error {
	return c.Update(AccountPeppolServiceModel, ids, as, nil)
}

// DeleteAccountPeppolService deletes an existing account_peppol.service record.
func (c *Client) DeleteAccountPeppolService(id int64) error {
	return c.DeleteAccountPeppolServices([]int64{id})
}

// DeleteAccountPeppolServices deletes existing account_peppol.service records.
func (c *Client) DeleteAccountPeppolServices(ids []int64) error {
	return c.Delete(AccountPeppolServiceModel, ids)
}

// GetAccountPeppolService gets account_peppol.service existing record.
func (c *Client) GetAccountPeppolService(id int64) (*AccountPeppolService, error) {
	ass, err := c.GetAccountPeppolServices([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ass)[0]), nil
}

// GetAccountPeppolServices gets account_peppol.service existing records.
func (c *Client) GetAccountPeppolServices(ids []int64) (*AccountPeppolServices, error) {
	ass := &AccountPeppolServices{}
	if err := c.Read(AccountPeppolServiceModel, ids, nil, ass); err != nil {
		return nil, err
	}
	return ass, nil
}

// FindAccountPeppolService finds account_peppol.service record by querying it with criteria.
func (c *Client) FindAccountPeppolService(criteria *Criteria) (*AccountPeppolService, error) {
	ass := &AccountPeppolServices{}
	if err := c.SearchRead(AccountPeppolServiceModel, criteria, NewOptions().Limit(1), ass); err != nil {
		return nil, err
	}
	return &((*ass)[0]), nil
}

// FindAccountPeppolServices finds account_peppol.service records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPeppolServices(criteria *Criteria, options *Options) (*AccountPeppolServices, error) {
	ass := &AccountPeppolServices{}
	if err := c.SearchRead(AccountPeppolServiceModel, criteria, options, ass); err != nil {
		return nil, err
	}
	return ass, nil
}

// FindAccountPeppolServiceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPeppolServiceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPeppolServiceModel, criteria, options)
}

// FindAccountPeppolServiceId finds record id by querying it with criteria.
func (c *Client) FindAccountPeppolServiceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPeppolServiceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
