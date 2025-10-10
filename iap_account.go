package odoo

// IapAccount represents iap.account model.
type IapAccount struct {
	AccountToken             *String    `xmlrpc:"account_token,omitempty"`
	Balance                  *String    `xmlrpc:"balance,omitempty"`
	CompanyIds               *Relation  `xmlrpc:"company_ids,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description              *String    `xmlrpc:"description,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	RatingIds                *Relation  `xmlrpc:"rating_ids,omitempty"`
	SenderName               *String    `xmlrpc:"sender_name,omitempty"`
	ServiceId                *Many2One  `xmlrpc:"service_id,omitempty"`
	ServiceLocked            *Bool      `xmlrpc:"service_locked,omitempty"`
	ServiceName              *String    `xmlrpc:"service_name,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty"`
	WarningThreshold         *Float     `xmlrpc:"warning_threshold,omitempty"`
	WarningUserIds           *Relation  `xmlrpc:"warning_user_ids,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// IapAccounts represents array of iap.account model.
type IapAccounts []IapAccount

// IapAccountModel is the odoo model name.
const IapAccountModel = "iap.account"

// Many2One convert IapAccount to *Many2One.
func (ia *IapAccount) Many2One() *Many2One {
	return NewMany2One(ia.Id.Get(), "")
}

// CreateIapAccount creates a new iap.account model and returns its id.
func (c *Client) CreateIapAccount(ia *IapAccount) (int64, error) {
	ids, err := c.CreateIapAccounts([]*IapAccount{ia})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIapAccount creates a new iap.account model and returns its id.
func (c *Client) CreateIapAccounts(ias []*IapAccount) ([]int64, error) {
	var vv []interface{}
	for _, v := range ias {
		vv = append(vv, v)
	}
	return c.Create(IapAccountModel, vv, nil)
}

// UpdateIapAccount updates an existing iap.account record.
func (c *Client) UpdateIapAccount(ia *IapAccount) error {
	return c.UpdateIapAccounts([]int64{ia.Id.Get()}, ia)
}

// UpdateIapAccounts updates existing iap.account records.
// All records (represented by ids) will be updated by ia values.
func (c *Client) UpdateIapAccounts(ids []int64, ia *IapAccount) error {
	return c.Update(IapAccountModel, ids, ia, nil)
}

// DeleteIapAccount deletes an existing iap.account record.
func (c *Client) DeleteIapAccount(id int64) error {
	return c.DeleteIapAccounts([]int64{id})
}

// DeleteIapAccounts deletes existing iap.account records.
func (c *Client) DeleteIapAccounts(ids []int64) error {
	return c.Delete(IapAccountModel, ids)
}

// GetIapAccount gets iap.account existing record.
func (c *Client) GetIapAccount(id int64) (*IapAccount, error) {
	ias, err := c.GetIapAccounts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ias)[0]), nil
}

// GetIapAccounts gets iap.account existing records.
func (c *Client) GetIapAccounts(ids []int64) (*IapAccounts, error) {
	ias := &IapAccounts{}
	if err := c.Read(IapAccountModel, ids, nil, ias); err != nil {
		return nil, err
	}
	return ias, nil
}

// FindIapAccount finds iap.account record by querying it with criteria.
func (c *Client) FindIapAccount(criteria *Criteria) (*IapAccount, error) {
	ias := &IapAccounts{}
	if err := c.SearchRead(IapAccountModel, criteria, NewOptions().Limit(1), ias); err != nil {
		return nil, err
	}
	return &((*ias)[0]), nil
}

// FindIapAccounts finds iap.account records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIapAccounts(criteria *Criteria, options *Options) (*IapAccounts, error) {
	ias := &IapAccounts{}
	if err := c.SearchRead(IapAccountModel, criteria, options, ias); err != nil {
		return nil, err
	}
	return ias, nil
}

// FindIapAccountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIapAccountIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IapAccountModel, criteria, options)
}

// FindIapAccountId finds record id by querying it with criteria.
func (c *Client) FindIapAccountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IapAccountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
