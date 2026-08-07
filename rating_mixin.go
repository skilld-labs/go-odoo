package odoo

// RatingMixin represents rating.mixin model.
type RatingMixin struct {
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	HasMessage                   *Bool      `xmlrpc:"has_message,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	RatingAvg                    *Float     `xmlrpc:"rating_avg,omitempty"`
	RatingAvgText                *Selection `xmlrpc:"rating_avg_text,omitempty"`
	RatingCount                  *Int       `xmlrpc:"rating_count,omitempty"`
	RatingIds                    *Relation  `xmlrpc:"rating_ids,omitempty"`
	RatingLastFeedback           *String    `xmlrpc:"rating_last_feedback,omitempty"`
	RatingLastImage              *String    `xmlrpc:"rating_last_image,omitempty"`
	RatingLastText               *Selection `xmlrpc:"rating_last_text,omitempty"`
	RatingLastValue              *Float     `xmlrpc:"rating_last_value,omitempty"`
	RatingPercentageSatisfaction *Float     `xmlrpc:"rating_percentage_satisfaction,omitempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omitempty"`
}

// RatingMixins represents array of rating.mixin model.
type RatingMixins []RatingMixin

// RatingMixinModel is the odoo model name.
const RatingMixinModel = "rating.mixin"

// Many2One convert RatingMixin to *Many2One.
func (rm *RatingMixin) Many2One() *Many2One {
	return NewMany2One(rm.Id.Get(), "")
}

// CreateRatingMixin creates a new rating.mixin model and returns its id.
func (c *Client) CreateRatingMixin(rm *RatingMixin) (int64, error) {
	ids, err := c.CreateRatingMixins([]*RatingMixin{rm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateRatingMixin creates a new rating.mixin model and returns its id.
func (c *Client) CreateRatingMixins(rms []*RatingMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range rms {
		vv = append(vv, v)
	}
	return c.Create(RatingMixinModel, vv, nil)
}

// UpdateRatingMixin updates an existing rating.mixin record.
func (c *Client) UpdateRatingMixin(rm *RatingMixin) error {
	return c.UpdateRatingMixins([]int64{rm.Id.Get()}, rm)
}

// UpdateRatingMixins updates existing rating.mixin records.
// All records (represented by ids) will be updated by rm values.
func (c *Client) UpdateRatingMixins(ids []int64, rm *RatingMixin) error {
	return c.Update(RatingMixinModel, ids, rm, nil)
}

// DeleteRatingMixin deletes an existing rating.mixin record.
func (c *Client) DeleteRatingMixin(id int64) error {
	return c.DeleteRatingMixins([]int64{id})
}

// DeleteRatingMixins deletes existing rating.mixin records.
func (c *Client) DeleteRatingMixins(ids []int64) error {
	return c.Delete(RatingMixinModel, ids)
}

// GetRatingMixin gets rating.mixin existing record.
func (c *Client) GetRatingMixin(id int64) (*RatingMixin, error) {
	rms, err := c.GetRatingMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rms)[0]), nil
}

// GetRatingMixins gets rating.mixin existing records.
func (c *Client) GetRatingMixins(ids []int64) (*RatingMixins, error) {
	rms := &RatingMixins{}
	if err := c.Read(RatingMixinModel, ids, nil, rms); err != nil {
		return nil, err
	}
	return rms, nil
}

// FindRatingMixin finds rating.mixin record by querying it with criteria.
func (c *Client) FindRatingMixin(criteria *Criteria) (*RatingMixin, error) {
	rms := &RatingMixins{}
	if err := c.SearchRead(RatingMixinModel, criteria, NewOptions().Limit(1), rms); err != nil {
		return nil, err
	}
	return &((*rms)[0]), nil
}

// FindRatingMixins finds rating.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRatingMixins(criteria *Criteria, options *Options) (*RatingMixins, error) {
	rms := &RatingMixins{}
	if err := c.SearchRead(RatingMixinModel, criteria, options, rms); err != nil {
		return nil, err
	}
	return rms, nil
}

// FindRatingMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRatingMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(RatingMixinModel, criteria, options)
}

// FindRatingMixinId finds record id by querying it with criteria.
func (c *Client) FindRatingMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RatingMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
