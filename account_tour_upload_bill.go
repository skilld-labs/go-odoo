package odoo

import (
	"fmt"
)

// AccountTourUploadBill represents account.tour.upload.bill model.
type AccountTourUploadBill struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omitempty"`
	ActiveDomain       *String    `xmlrpc:"active_domain,omitempty"`
	AddSign            *Bool      `xmlrpc:"add_sign,omitempty"`
	AttachmentIds      *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AuthorId           *Many2One  `xmlrpc:"author_id,omitempty"`
	AutoDelete         *Bool      `xmlrpc:"auto_delete,omitempty"`
	AutoDeleteMessage  *Bool      `xmlrpc:"auto_delete_message,omitempty"`
	Body               *String    `xmlrpc:"body,omitempty"`
	ComposerId         *Many2One  `xmlrpc:"composer_id,omitempty"`
	CompositionMode    *Selection `xmlrpc:"composition_mode,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	EmailFrom          *String    `xmlrpc:"email_from,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	IsLog              *Bool      `xmlrpc:"is_log,omitempty"`
	Layout             *String    `xmlrpc:"layout,omitempty"`
	MailActivityTypeId *Many2One  `xmlrpc:"mail_activity_type_id,omitempty"`
	MailServerId       *Many2One  `xmlrpc:"mail_server_id,omitempty"`
	MessageType        *Selection `xmlrpc:"message_type,omitempty"`
	Model              *String    `xmlrpc:"model,omitempty"`
	NoAutoThread       *Bool      `xmlrpc:"no_auto_thread,omitempty"`
	Notify             *Bool      `xmlrpc:"notify,omitempty"`
	ParentId           *Many2One  `xmlrpc:"parent_id,omitempty"`
	PartnerIds         *Relation  `xmlrpc:"partner_ids,omitempty"`
	RecordName         *String    `xmlrpc:"record_name,omitempty"`
	ReplyTo            *String    `xmlrpc:"reply_to,omitempty"`
	ResId              *Int       `xmlrpc:"res_id,omitempty"`
	SampleBillPreview  *String    `xmlrpc:"sample_bill_preview,omitempty"`
	Selection          *Selection `xmlrpc:"selection,omitempty"`
	Subject            *String    `xmlrpc:"subject,omitempty"`
	SubtypeId          *Many2One  `xmlrpc:"subtype_id,omitempty"`
	TemplateId         *Many2One  `xmlrpc:"template_id,omitempty"`
	UseActiveDomain    *Bool      `xmlrpc:"use_active_domain,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountTourUploadBills represents array of account.tour.upload.bill model.
type AccountTourUploadBills []AccountTourUploadBill

// AccountTourUploadBillModel is the odoo model name.
const AccountTourUploadBillModel = "account.tour.upload.bill"

// Many2One convert AccountTourUploadBill to *Many2One.
func (atub *AccountTourUploadBill) Many2One() *Many2One {
	return NewMany2One(atub.Id.Get(), "")
}

// CreateAccountTourUploadBill creates a new account.tour.upload.bill model and returns its id.
func (c *Client) CreateAccountTourUploadBill(atub *AccountTourUploadBill) (int64, error) {
	return c.Create(AccountTourUploadBillModel, atub)
}

// UpdateAccountTourUploadBill updates an existing account.tour.upload.bill record.
func (c *Client) UpdateAccountTourUploadBill(atub *AccountTourUploadBill) error {
	return c.UpdateAccountTourUploadBills([]int64{atub.Id.Get()}, atub)
}

// UpdateAccountTourUploadBills updates existing account.tour.upload.bill records.
// All records (represented by IDs) will be updated by atub values.
func (c *Client) UpdateAccountTourUploadBills(ids []int64, atub *AccountTourUploadBill) error {
	return c.Update(AccountTourUploadBillModel, ids, atub)
}

// DeleteAccountTourUploadBill deletes an existing account.tour.upload.bill record.
func (c *Client) DeleteAccountTourUploadBill(id int64) error {
	return c.DeleteAccountTourUploadBills([]int64{id})
}

// DeleteAccountTourUploadBills deletes existing account.tour.upload.bill records.
func (c *Client) DeleteAccountTourUploadBills(ids []int64) error {
	return c.Delete(AccountTourUploadBillModel, ids)
}

// GetAccountTourUploadBill gets account.tour.upload.bill existing record.
func (c *Client) GetAccountTourUploadBill(id int64) (*AccountTourUploadBill, error) {
	atubs, err := c.GetAccountTourUploadBills([]int64{id})
	if err != nil {
		return nil, err
	}
	if atubs != nil && len(*atubs) > 0 {
		return &((*atubs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.tour.upload.bill not found", id)
}

// GetAccountTourUploadBills gets account.tour.upload.bill existing records.
func (c *Client) GetAccountTourUploadBills(ids []int64) (*AccountTourUploadBills, error) {
	atubs := &AccountTourUploadBills{}
	if err := c.Read(AccountTourUploadBillModel, ids, nil, atubs); err != nil {
		return nil, err
	}
	return atubs, nil
}

// FindAccountTourUploadBill finds account.tour.upload.bill record by querying it with criteria.
func (c *Client) FindAccountTourUploadBill(criteria *Criteria) (*AccountTourUploadBill, error) {
	atubs := &AccountTourUploadBills{}
	if err := c.SearchRead(AccountTourUploadBillModel, criteria, NewOptions().Limit(1), atubs); err != nil {
		return nil, err
	}
	if atubs != nil && len(*atubs) > 0 {
		return &((*atubs)[0]), nil
	}
	return nil, fmt.Errorf("account.tour.upload.bill was not found")
}

// FindAccountTourUploadBills finds account.tour.upload.bill records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTourUploadBills(criteria *Criteria, options *Options) (*AccountTourUploadBills, error) {
	atubs := &AccountTourUploadBills{}
	if err := c.SearchRead(AccountTourUploadBillModel, criteria, options, atubs); err != nil {
		return nil, err
	}
	return atubs, nil
}

// FindAccountTourUploadBillIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTourUploadBillIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountTourUploadBillModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountTourUploadBillId finds record id by querying it with criteria.
func (c *Client) FindAccountTourUploadBillId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTourUploadBillModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.tour.upload.bill was not found")
}
