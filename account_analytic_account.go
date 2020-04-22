package odoo

import (
	"fmt"
)

// AccountAnalyticAccount represents account.analytic.account model
type AccountAnalyticAccount struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	Active                   *Bool     `xmlrpc:"active,omptempty"`
	Balance                  *Float    `xmlrpc:"balance,omptempty"`
	Code                     *String   `xmlrpc:"code,omptempty"`
	CompanyId                *Many2One `xmlrpc:"company_id,omptempty"`
	CompanyUomId             *Many2One `xmlrpc:"company_uom_id,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	Credit                   *Float    `xmlrpc:"credit,omptempty"`
	CurrencyId               *Many2One `xmlrpc:"currency_id,omptempty"`
	Debit                    *Float    `xmlrpc:"debit,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	LineIds                  *Relation `xmlrpc:"line_ids,omptempty"`
	MachineProjectName       *String   `xmlrpc:"machine_project_name,omptempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost          *Time     `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	PartnerId                *Many2One `xmlrpc:"partner_id,omptempty"`
	ProjectCount             *Int      `xmlrpc:"project_count,omptempty"`
	ProjectCreated           *Bool     `xmlrpc:"project_created,omptempty"`
	ProjectIds               *Relation `xmlrpc:"project_ids,omptempty"`
	TagIds                   *Relation `xmlrpc:"tag_ids,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountAnalyticAccounts represents array of account.analytic.account model
type AccountAnalyticAccounts []AccountAnalyticAccount

// AccountAnalyticAccountModel is the odoo model name
const AccountAnalyticAccountModel = "account.analytic.account"

// CreateAccountAnalyticAccount creates a new account.analytic.account model and returns its id.
func (c *Client) CreateAccountAnalyticAccount(aaa *AccountAnalyticAccount) (int64, error) {
	return c.Create(AccountAnalyticAccountModel, aaa)
}

// UpdateAccountAnalyticAccount pdates an existing account.analytic.account record.
func (c *Client) UpdateAccountAnalyticAccount(aaa *AccountAnalyticAccount) error {
	return c.UpdateAccountAnalyticAccounts([]int64{aaa.Id.Get()}, aaa)
}

// UpdateAccountAnalyticAccounts updates existing account.analytic.account records.
// All records (represented by ids) will be updated by aaa values.
func (c *Client) UpdateAccountAnalyticAccounts(ids []int64, aaa *AccountAnalyticAccount) error {
	return c.Update(AccountAnalyticAccountModel, ids, aaa)
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
	if aaas != nil && len(*aaas) > 0 {
		return &((*aaas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of %s not found", id, AccountAnalyticAccountModel)
}

// GetAccountAnalyticAccounts gets account.analytic.account existing records.
func (c *Client) GetAccountAnalyticAccounts(ids []int64) (*AccountAnalyticAccounts, error) {
	aaas := &AccountAnalyticAccounts{}
	if err := c.Read(AccountAnalyticAccountModel, ids, nil, aaas); err != nil {
		return nil, err
	}
	return aaas, nil
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
