package odoo

// MailActivityMixin represents mail.activity.mixin model.
type MailActivityMixin struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omitempty"`
	ActivityDateDeadline *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityIds          *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState        *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary      *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId       *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId       *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
}

// MailActivityMixins represents array of mail.activity.mixin model.
type MailActivityMixins []MailActivityMixin

// MailActivityMixinModel is the odoo model name.
const MailActivityMixinModel = "mail.activity.mixin"

// Many2One convert MailActivityMixin to *Many2One.
func (mam *MailActivityMixin) Many2One() *Many2One {
	return NewMany2One(mam.Id.Get(), "")
}

// CreateMailActivityMixin creates a new mail.activity.mixin model and returns its id.
func (c *Client) CreateMailActivityMixin(mam *MailActivityMixin) (int64, error) {
	ids, err := c.CreateMailActivityMixins([]*MailActivityMixin{mam})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailActivityMixin creates a new mail.activity.mixin model and returns its id.
func (c *Client) CreateMailActivityMixins(mams []*MailActivityMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range mams {
		vv = append(vv, v)
	}
	return c.Create(MailActivityMixinModel, vv, nil)
}

// UpdateMailActivityMixin updates an existing mail.activity.mixin record.
func (c *Client) UpdateMailActivityMixin(mam *MailActivityMixin) error {
	return c.UpdateMailActivityMixins([]int64{mam.Id.Get()}, mam)
}

// UpdateMailActivityMixins updates existing mail.activity.mixin records.
// All records (represented by ids) will be updated by mam values.
func (c *Client) UpdateMailActivityMixins(ids []int64, mam *MailActivityMixin) error {
	return c.Update(MailActivityMixinModel, ids, mam, nil)
}

// DeleteMailActivityMixin deletes an existing mail.activity.mixin record.
func (c *Client) DeleteMailActivityMixin(id int64) error {
	return c.DeleteMailActivityMixins([]int64{id})
}

// DeleteMailActivityMixins deletes existing mail.activity.mixin records.
func (c *Client) DeleteMailActivityMixins(ids []int64) error {
	return c.Delete(MailActivityMixinModel, ids)
}

// GetMailActivityMixin gets mail.activity.mixin existing record.
func (c *Client) GetMailActivityMixin(id int64) (*MailActivityMixin, error) {
	mams, err := c.GetMailActivityMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mams)[0]), nil
}

// GetMailActivityMixins gets mail.activity.mixin existing records.
func (c *Client) GetMailActivityMixins(ids []int64) (*MailActivityMixins, error) {
	mams := &MailActivityMixins{}
	if err := c.Read(MailActivityMixinModel, ids, nil, mams); err != nil {
		return nil, err
	}
	return mams, nil
}

// FindMailActivityMixin finds mail.activity.mixin record by querying it with criteria.
func (c *Client) FindMailActivityMixin(criteria *Criteria) (*MailActivityMixin, error) {
	mams := &MailActivityMixins{}
	if err := c.SearchRead(MailActivityMixinModel, criteria, NewOptions().Limit(1), mams); err != nil {
		return nil, err
	}
	return &((*mams)[0]), nil
}

// FindMailActivityMixins finds mail.activity.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityMixins(criteria *Criteria, options *Options) (*MailActivityMixins, error) {
	mams := &MailActivityMixins{}
	if err := c.SearchRead(MailActivityMixinModel, criteria, options, mams); err != nil {
		return nil, err
	}
	return mams, nil
}

// FindMailActivityMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailActivityMixinModel, criteria, options)
}

// FindMailActivityMixinId finds record id by querying it with criteria.
func (c *Client) FindMailActivityMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailActivityMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
