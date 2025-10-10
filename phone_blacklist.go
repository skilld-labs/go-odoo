package odoo

// PhoneBlacklist represents phone.blacklist model.
type PhoneBlacklist struct {
	Active                   *Bool     `xmlrpc:"active,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	HasMessage               *Bool     `xmlrpc:"has_message,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	Number                   *String   `xmlrpc:"number,omitempty"`
	RatingIds                *Relation `xmlrpc:"rating_ids,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PhoneBlacklists represents array of phone.blacklist model.
type PhoneBlacklists []PhoneBlacklist

// PhoneBlacklistModel is the odoo model name.
const PhoneBlacklistModel = "phone.blacklist"

// Many2One convert PhoneBlacklist to *Many2One.
func (pb *PhoneBlacklist) Many2One() *Many2One {
	return NewMany2One(pb.Id.Get(), "")
}

// CreatePhoneBlacklist creates a new phone.blacklist model and returns its id.
func (c *Client) CreatePhoneBlacklist(pb *PhoneBlacklist) (int64, error) {
	ids, err := c.CreatePhoneBlacklists([]*PhoneBlacklist{pb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePhoneBlacklist creates a new phone.blacklist model and returns its id.
func (c *Client) CreatePhoneBlacklists(pbs []*PhoneBlacklist) ([]int64, error) {
	var vv []interface{}
	for _, v := range pbs {
		vv = append(vv, v)
	}
	return c.Create(PhoneBlacklistModel, vv, nil)
}

// UpdatePhoneBlacklist updates an existing phone.blacklist record.
func (c *Client) UpdatePhoneBlacklist(pb *PhoneBlacklist) error {
	return c.UpdatePhoneBlacklists([]int64{pb.Id.Get()}, pb)
}

// UpdatePhoneBlacklists updates existing phone.blacklist records.
// All records (represented by ids) will be updated by pb values.
func (c *Client) UpdatePhoneBlacklists(ids []int64, pb *PhoneBlacklist) error {
	return c.Update(PhoneBlacklistModel, ids, pb, nil)
}

// DeletePhoneBlacklist deletes an existing phone.blacklist record.
func (c *Client) DeletePhoneBlacklist(id int64) error {
	return c.DeletePhoneBlacklists([]int64{id})
}

// DeletePhoneBlacklists deletes existing phone.blacklist records.
func (c *Client) DeletePhoneBlacklists(ids []int64) error {
	return c.Delete(PhoneBlacklistModel, ids)
}

// GetPhoneBlacklist gets phone.blacklist existing record.
func (c *Client) GetPhoneBlacklist(id int64) (*PhoneBlacklist, error) {
	pbs, err := c.GetPhoneBlacklists([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pbs)[0]), nil
}

// GetPhoneBlacklists gets phone.blacklist existing records.
func (c *Client) GetPhoneBlacklists(ids []int64) (*PhoneBlacklists, error) {
	pbs := &PhoneBlacklists{}
	if err := c.Read(PhoneBlacklistModel, ids, nil, pbs); err != nil {
		return nil, err
	}
	return pbs, nil
}

// FindPhoneBlacklist finds phone.blacklist record by querying it with criteria.
func (c *Client) FindPhoneBlacklist(criteria *Criteria) (*PhoneBlacklist, error) {
	pbs := &PhoneBlacklists{}
	if err := c.SearchRead(PhoneBlacklistModel, criteria, NewOptions().Limit(1), pbs); err != nil {
		return nil, err
	}
	return &((*pbs)[0]), nil
}

// FindPhoneBlacklists finds phone.blacklist records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPhoneBlacklists(criteria *Criteria, options *Options) (*PhoneBlacklists, error) {
	pbs := &PhoneBlacklists{}
	if err := c.SearchRead(PhoneBlacklistModel, criteria, options, pbs); err != nil {
		return nil, err
	}
	return pbs, nil
}

// FindPhoneBlacklistIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPhoneBlacklistIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PhoneBlacklistModel, criteria, options)
}

// FindPhoneBlacklistId finds record id by querying it with criteria.
func (c *Client) FindPhoneBlacklistId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PhoneBlacklistModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
