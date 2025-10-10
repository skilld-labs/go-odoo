package odoo

// AccountAnalyticDistributionModel represents account.analytic.distribution.model model.
type AccountAnalyticDistributionModel struct {
	AccountPrefix                  *String     `xmlrpc:"account_prefix,omitempty"`
	AnalyticDistribution           interface{} `xmlrpc:"analytic_distribution,omitempty"`
	AnalyticPrecision              *Int        `xmlrpc:"analytic_precision,omitempty"`
	CompanyId                      *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate                     *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName                    *String     `xmlrpc:"display_name,omitempty"`
	DistributionAnalyticAccountIds *Relation   `xmlrpc:"distribution_analytic_account_ids,omitempty"`
	Id                             *Int        `xmlrpc:"id,omitempty"`
	PartnerCategoryId              *Many2One   `xmlrpc:"partner_category_id,omitempty"`
	PartnerId                      *Many2One   `xmlrpc:"partner_id,omitempty"`
	PrefixPlaceholder              *String     `xmlrpc:"prefix_placeholder,omitempty"`
	ProductCategId                 *Many2One   `xmlrpc:"product_categ_id,omitempty"`
	ProductId                      *Many2One   `xmlrpc:"product_id,omitempty"`
	Sequence                       *Int        `xmlrpc:"sequence,omitempty"`
	WriteDate                      *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountAnalyticDistributionModels represents array of account.analytic.distribution.model model.
type AccountAnalyticDistributionModels []AccountAnalyticDistributionModel

// AccountAnalyticDistributionModelModel is the odoo model name.
const AccountAnalyticDistributionModelModel = "account.analytic.distribution.model"

// Many2One convert AccountAnalyticDistributionModel to *Many2One.
func (aadm *AccountAnalyticDistributionModel) Many2One() *Many2One {
	return NewMany2One(aadm.Id.Get(), "")
}

// CreateAccountAnalyticDistributionModel creates a new account.analytic.distribution.model model and returns its id.
func (c *Client) CreateAccountAnalyticDistributionModel(aadm *AccountAnalyticDistributionModel) (int64, error) {
	ids, err := c.CreateAccountAnalyticDistributionModels([]*AccountAnalyticDistributionModel{aadm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAnalyticDistributionModel creates a new account.analytic.distribution.model model and returns its id.
func (c *Client) CreateAccountAnalyticDistributionModels(aadms []*AccountAnalyticDistributionModel) ([]int64, error) {
	var vv []interface{}
	for _, v := range aadms {
		vv = append(vv, v)
	}
	return c.Create(AccountAnalyticDistributionModelModel, vv, nil)
}

// UpdateAccountAnalyticDistributionModel updates an existing account.analytic.distribution.model record.
func (c *Client) UpdateAccountAnalyticDistributionModel(aadm *AccountAnalyticDistributionModel) error {
	return c.UpdateAccountAnalyticDistributionModels([]int64{aadm.Id.Get()}, aadm)
}

// UpdateAccountAnalyticDistributionModels updates existing account.analytic.distribution.model records.
// All records (represented by ids) will be updated by aadm values.
func (c *Client) UpdateAccountAnalyticDistributionModels(ids []int64, aadm *AccountAnalyticDistributionModel) error {
	return c.Update(AccountAnalyticDistributionModelModel, ids, aadm, nil)
}

// DeleteAccountAnalyticDistributionModel deletes an existing account.analytic.distribution.model record.
func (c *Client) DeleteAccountAnalyticDistributionModel(id int64) error {
	return c.DeleteAccountAnalyticDistributionModels([]int64{id})
}

// DeleteAccountAnalyticDistributionModels deletes existing account.analytic.distribution.model records.
func (c *Client) DeleteAccountAnalyticDistributionModels(ids []int64) error {
	return c.Delete(AccountAnalyticDistributionModelModel, ids)
}

// GetAccountAnalyticDistributionModel gets account.analytic.distribution.model existing record.
func (c *Client) GetAccountAnalyticDistributionModel(id int64) (*AccountAnalyticDistributionModel, error) {
	aadms, err := c.GetAccountAnalyticDistributionModels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aadms)[0]), nil
}

// GetAccountAnalyticDistributionModels gets account.analytic.distribution.model existing records.
func (c *Client) GetAccountAnalyticDistributionModels(ids []int64) (*AccountAnalyticDistributionModels, error) {
	aadms := &AccountAnalyticDistributionModels{}
	if err := c.Read(AccountAnalyticDistributionModelModel, ids, nil, aadms); err != nil {
		return nil, err
	}
	return aadms, nil
}

// FindAccountAnalyticDistributionModel finds account.analytic.distribution.model record by querying it with criteria.
func (c *Client) FindAccountAnalyticDistributionModel(criteria *Criteria) (*AccountAnalyticDistributionModel, error) {
	aadms := &AccountAnalyticDistributionModels{}
	if err := c.SearchRead(AccountAnalyticDistributionModelModel, criteria, NewOptions().Limit(1), aadms); err != nil {
		return nil, err
	}
	return &((*aadms)[0]), nil
}

// FindAccountAnalyticDistributionModels finds account.analytic.distribution.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticDistributionModels(criteria *Criteria, options *Options) (*AccountAnalyticDistributionModels, error) {
	aadms := &AccountAnalyticDistributionModels{}
	if err := c.SearchRead(AccountAnalyticDistributionModelModel, criteria, options, aadms); err != nil {
		return nil, err
	}
	return aadms, nil
}

// FindAccountAnalyticDistributionModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticDistributionModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAnalyticDistributionModelModel, criteria, options)
}

// FindAccountAnalyticDistributionModelId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticDistributionModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticDistributionModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
