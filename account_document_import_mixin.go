package odoo

// AccountDocumentImportMixin represents account.document.import.mixin model.
type AccountDocumentImportMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountDocumentImportMixins represents array of account.document.import.mixin model.
type AccountDocumentImportMixins []AccountDocumentImportMixin

// AccountDocumentImportMixinModel is the odoo model name.
const AccountDocumentImportMixinModel = "account.document.import.mixin"

// Many2One convert AccountDocumentImportMixin to *Many2One.
func (adim *AccountDocumentImportMixin) Many2One() *Many2One {
	return NewMany2One(adim.Id.Get(), "")
}

// CreateAccountDocumentImportMixin creates a new account.document.import.mixin model and returns its id.
func (c *Client) CreateAccountDocumentImportMixin(adim *AccountDocumentImportMixin) (int64, error) {
	ids, err := c.CreateAccountDocumentImportMixins([]*AccountDocumentImportMixin{adim})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountDocumentImportMixin creates a new account.document.import.mixin model and returns its id.
func (c *Client) CreateAccountDocumentImportMixins(adims []*AccountDocumentImportMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range adims {
		vv = append(vv, v)
	}
	return c.Create(AccountDocumentImportMixinModel, vv, nil)
}

// UpdateAccountDocumentImportMixin updates an existing account.document.import.mixin record.
func (c *Client) UpdateAccountDocumentImportMixin(adim *AccountDocumentImportMixin) error {
	return c.UpdateAccountDocumentImportMixins([]int64{adim.Id.Get()}, adim)
}

// UpdateAccountDocumentImportMixins updates existing account.document.import.mixin records.
// All records (represented by ids) will be updated by adim values.
func (c *Client) UpdateAccountDocumentImportMixins(ids []int64, adim *AccountDocumentImportMixin) error {
	return c.Update(AccountDocumentImportMixinModel, ids, adim, nil)
}

// DeleteAccountDocumentImportMixin deletes an existing account.document.import.mixin record.
func (c *Client) DeleteAccountDocumentImportMixin(id int64) error {
	return c.DeleteAccountDocumentImportMixins([]int64{id})
}

// DeleteAccountDocumentImportMixins deletes existing account.document.import.mixin records.
func (c *Client) DeleteAccountDocumentImportMixins(ids []int64) error {
	return c.Delete(AccountDocumentImportMixinModel, ids)
}

// GetAccountDocumentImportMixin gets account.document.import.mixin existing record.
func (c *Client) GetAccountDocumentImportMixin(id int64) (*AccountDocumentImportMixin, error) {
	adims, err := c.GetAccountDocumentImportMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*adims)[0]), nil
}

// GetAccountDocumentImportMixins gets account.document.import.mixin existing records.
func (c *Client) GetAccountDocumentImportMixins(ids []int64) (*AccountDocumentImportMixins, error) {
	adims := &AccountDocumentImportMixins{}
	if err := c.Read(AccountDocumentImportMixinModel, ids, nil, adims); err != nil {
		return nil, err
	}
	return adims, nil
}

// FindAccountDocumentImportMixin finds account.document.import.mixin record by querying it with criteria.
func (c *Client) FindAccountDocumentImportMixin(criteria *Criteria) (*AccountDocumentImportMixin, error) {
	adims := &AccountDocumentImportMixins{}
	if err := c.SearchRead(AccountDocumentImportMixinModel, criteria, NewOptions().Limit(1), adims); err != nil {
		return nil, err
	}
	return &((*adims)[0]), nil
}

// FindAccountDocumentImportMixins finds account.document.import.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDocumentImportMixins(criteria *Criteria, options *Options) (*AccountDocumentImportMixins, error) {
	adims := &AccountDocumentImportMixins{}
	if err := c.SearchRead(AccountDocumentImportMixinModel, criteria, options, adims); err != nil {
		return nil, err
	}
	return adims, nil
}

// FindAccountDocumentImportMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountDocumentImportMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountDocumentImportMixinModel, criteria, options)
}

// FindAccountDocumentImportMixinId finds record id by querying it with criteria.
func (c *Client) FindAccountDocumentImportMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountDocumentImportMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
