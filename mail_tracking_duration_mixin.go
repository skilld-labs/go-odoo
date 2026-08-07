package odoo

// MailTrackingDurationMixin represents mail.tracking.duration.mixin model.
type MailTrackingDurationMixin struct {
	DisplayName              *String     `xmlrpc:"display_name,omitempty"`
	DurationTracking         interface{} `xmlrpc:"duration_tracking,omitempty"`
	HasMessage               *Bool       `xmlrpc:"has_message,omitempty"`
	Id                       *Int        `xmlrpc:"id,omitempty"`
	IsRotting                *Bool       `xmlrpc:"is_rotting,omitempty"`
	MessageAttachmentCount   *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	RatingIds                *Relation   `xmlrpc:"rating_ids,omitempty"`
	RottingDays              *Int        `xmlrpc:"rotting_days,omitempty"`
	WebsiteMessageIds        *Relation   `xmlrpc:"website_message_ids,omitempty"`
}

// MailTrackingDurationMixins represents array of mail.tracking.duration.mixin model.
type MailTrackingDurationMixins []MailTrackingDurationMixin

// MailTrackingDurationMixinModel is the odoo model name.
const MailTrackingDurationMixinModel = "mail.tracking.duration.mixin"

// Many2One convert MailTrackingDurationMixin to *Many2One.
func (mtdm *MailTrackingDurationMixin) Many2One() *Many2One {
	return NewMany2One(mtdm.Id.Get(), "")
}

// CreateMailTrackingDurationMixin creates a new mail.tracking.duration.mixin model and returns its id.
func (c *Client) CreateMailTrackingDurationMixin(mtdm *MailTrackingDurationMixin) (int64, error) {
	ids, err := c.CreateMailTrackingDurationMixins([]*MailTrackingDurationMixin{mtdm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailTrackingDurationMixin creates a new mail.tracking.duration.mixin model and returns its id.
func (c *Client) CreateMailTrackingDurationMixins(mtdms []*MailTrackingDurationMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range mtdms {
		vv = append(vv, v)
	}
	return c.Create(MailTrackingDurationMixinModel, vv, nil)
}

// UpdateMailTrackingDurationMixin updates an existing mail.tracking.duration.mixin record.
func (c *Client) UpdateMailTrackingDurationMixin(mtdm *MailTrackingDurationMixin) error {
	return c.UpdateMailTrackingDurationMixins([]int64{mtdm.Id.Get()}, mtdm)
}

// UpdateMailTrackingDurationMixins updates existing mail.tracking.duration.mixin records.
// All records (represented by ids) will be updated by mtdm values.
func (c *Client) UpdateMailTrackingDurationMixins(ids []int64, mtdm *MailTrackingDurationMixin) error {
	return c.Update(MailTrackingDurationMixinModel, ids, mtdm, nil)
}

// DeleteMailTrackingDurationMixin deletes an existing mail.tracking.duration.mixin record.
func (c *Client) DeleteMailTrackingDurationMixin(id int64) error {
	return c.DeleteMailTrackingDurationMixins([]int64{id})
}

// DeleteMailTrackingDurationMixins deletes existing mail.tracking.duration.mixin records.
func (c *Client) DeleteMailTrackingDurationMixins(ids []int64) error {
	return c.Delete(MailTrackingDurationMixinModel, ids)
}

// GetMailTrackingDurationMixin gets mail.tracking.duration.mixin existing record.
func (c *Client) GetMailTrackingDurationMixin(id int64) (*MailTrackingDurationMixin, error) {
	mtdms, err := c.GetMailTrackingDurationMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mtdms)[0]), nil
}

// GetMailTrackingDurationMixins gets mail.tracking.duration.mixin existing records.
func (c *Client) GetMailTrackingDurationMixins(ids []int64) (*MailTrackingDurationMixins, error) {
	mtdms := &MailTrackingDurationMixins{}
	if err := c.Read(MailTrackingDurationMixinModel, ids, nil, mtdms); err != nil {
		return nil, err
	}
	return mtdms, nil
}

// FindMailTrackingDurationMixin finds mail.tracking.duration.mixin record by querying it with criteria.
func (c *Client) FindMailTrackingDurationMixin(criteria *Criteria) (*MailTrackingDurationMixin, error) {
	mtdms := &MailTrackingDurationMixins{}
	if err := c.SearchRead(MailTrackingDurationMixinModel, criteria, NewOptions().Limit(1), mtdms); err != nil {
		return nil, err
	}
	return &((*mtdms)[0]), nil
}

// FindMailTrackingDurationMixins finds mail.tracking.duration.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTrackingDurationMixins(criteria *Criteria, options *Options) (*MailTrackingDurationMixins, error) {
	mtdms := &MailTrackingDurationMixins{}
	if err := c.SearchRead(MailTrackingDurationMixinModel, criteria, options, mtdms); err != nil {
		return nil, err
	}
	return mtdms, nil
}

// FindMailTrackingDurationMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTrackingDurationMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailTrackingDurationMixinModel, criteria, options)
}

// FindMailTrackingDurationMixinId finds record id by querying it with criteria.
func (c *Client) FindMailTrackingDurationMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailTrackingDurationMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
