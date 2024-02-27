package odoo

// AccountAnalyticAccount represents account.analytic.account model.
type AccountAnalyticAccount struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	Active                   *Bool     `xmlrpc:"active,omitempty"`
	Balance                  *Float    `xmlrpc:"balance,omitempty"`
	Code                     *String   `xmlrpc:"code,omitempty"`
	CompanyId                *Many2One `xmlrpc:"company_id,omitempty"`
	CompanyUomId             *Many2One `xmlrpc:"company_uom_id,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	Credit                   *Float    `xmlrpc:"credit,omitempty"`
	CurrencyId               *Many2One `xmlrpc:"currency_id,omitempty"`
	Debit                    *Float    `xmlrpc:"debit,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	LineIds                  *Relation `xmlrpc:"line_ids,omitempty"`
	MachineInitiativeName    *String   `xmlrpc:"machine_initiative_name,omitempty"`
	MachineProjectName       *String   `xmlrpc:"machine_project_name,omitempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time     `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omitempty"`
	Name                     *String   `xmlrpc:"name,omitempty"`
	PartnerId                *Many2One `xmlrpc:"partner_id,omitempty"`
	ProjectCount             *Int      `xmlrpc:"project_count,omitempty"`
	ProjectCreated           *Bool     `xmlrpc:"project_created,omitempty"`
	ProjectIds               *Relation `xmlrpc:"project_ids,omitempty"`
	TagIds                   *Relation `xmlrpc:"tag_ids,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountAnalyticAccounts represents array of account.analytic.account model.
type AccountAnalyticAccounts []AccountAnalyticAccount

// AccountAnalyticAccountModel is the odoo model name.
const AccountAnalyticAccountModel = "account.analytic.account"

// Many2One convert AccountAnalyticAccount to *Many2One.
func (aaa *AccountAnalyticAccount) Many2One() *Many2One {
	return NewMany2One(aaa.Id.Get(), "")
}

// CreateAccountAnalyticAccount creates a new account.analytic.account model and returns its id.
func (c *Client) CreateAccountAnalyticAccount(aaa *AccountAnalyticAccount) (int64, error) {
	ids, err := c.CreateAccountAnalyticAccounts([]*AccountAnalyticAccount{aaa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAnalyticAccounts creates a new account.analytic.account model and returns its id.
func (c *Client) CreateAccountAnalyticAccounts(aaas []*AccountAnalyticAccount) ([]int64, error) {
	var vv []interface{}
	for _, v := range aaas {
		vv = append(vv, v)
	}
	return c.Create(AccountAnalyticAccountModel, vv, nil)
}

// UpdateAccountAnalyticAccount updates an existing account.analytic.account record.
func (c *Client) UpdateAccountAnalyticAccount(aaa *AccountAnalyticAccount) error {
	return c.UpdateAccountAnalyticAccounts([]int64{aaa.Id.Get()}, aaa)
}

// UpdateAccountAnalyticAccounts updates existing account.analytic.account records.
// All records (represented by ids) will be updated by aaa values.
func (c *Client) UpdateAccountAnalyticAccounts(ids []int64, aaa *AccountAnalyticAccount) error {
	return c.Update(AccountAnalyticAccountModel, ids, aaa, nil)
}

// DeleteAccountAnalyticAccount deletes an existing account.analytic.account record.
func (c *Client) DeleteAccountAnalyticAccount(id int64) error {
	return c.DeleteAccountAnalyticAccounts([]int64{id})
}

// DeleteAccountAnalyticAccounts deletes existing account.analytic.account records.
func (c *Client) DeleteAccountAnalyticAccounts(ids []int64) error {
	return c.Delete(AccountAnalyticAccountModel, ids)
}

// GetAccountAnalyticAccount gets account.analytic.account existing record.
func (c *Client) GetAccountAnalyticAccount(id int64) (*AccountAnalyticAccount, error) {
	aaas, err := c.GetAccountAnalyticAccounts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aaas)[0]), nil
}

// GetAccountAnalyticAccounts gets account.analytic.account existing records.
func (c *Client) GetAccountAnalyticAccounts(ids []int64) (*AccountAnalyticAccounts, error) {
	aaas := &AccountAnalyticAccounts{}
	if err := c.Read(AccountAnalyticAccountModel, ids, nil, aaas); err != nil {
		return nil, err
	}
	return aaas, nil
}

// FindAccountAnalyticAccount finds account.analytic.account record by querying it with criteria.
func (c *Client) FindAccountAnalyticAccount(criteria *Criteria) (*AccountAnalyticAccount, error) {
	aaas := &AccountAnalyticAccounts{}
	if err := c.SearchRead(AccountAnalyticAccountModel, criteria, NewOptions().Limit(1), aaas); err != nil {
		return nil, err
	}
	return &((*aaas)[0]), nil
}

// FindAccountAnalyticAccounts finds account.analytic.account records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticAccounts(criteria *Criteria, options *Options) (*AccountAnalyticAccounts, error) {
	aaas := &AccountAnalyticAccounts{}
	if err := c.SearchRead(AccountAnalyticAccountModel, criteria, options, aaas); err != nil {
		return nil, err
	}
	return aaas, nil
}

// FindAccountAnalyticAccountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticAccountIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAnalyticAccountModel, criteria, options)
}

// FindAccountAnalyticAccountId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticAccountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticAccountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
