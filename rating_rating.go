package odoo

import (
	"fmt"
)

// RatingRating represents rating.rating model.
type RatingRating struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken      *String    `xmlrpc:"access_token,omitempty"`
	Consumed         *Bool      `xmlrpc:"consumed,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	Feedback         *String    `xmlrpc:"feedback,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	MessageId        *Many2One  `xmlrpc:"message_id,omitempty"`
	ParentResId      *Int       `xmlrpc:"parent_res_id,omitempty"`
	ParentResModel   *String    `xmlrpc:"parent_res_model,omitempty"`
	ParentResModelId *Many2One  `xmlrpc:"parent_res_model_id,omitempty"`
	ParentResName    *String    `xmlrpc:"parent_res_name,omitempty"`
	PartnerId        *Many2One  `xmlrpc:"partner_id,omitempty"`
	RatedPartnerId   *Many2One  `xmlrpc:"rated_partner_id,omitempty"`
	Rating           *Float     `xmlrpc:"rating,omitempty"`
	RatingImage      *String    `xmlrpc:"rating_image,omitempty"`
	RatingText       *Selection `xmlrpc:"rating_text,omitempty"`
	ResId            *Int       `xmlrpc:"res_id,omitempty"`
	ResModel         *String    `xmlrpc:"res_model,omitempty"`
	ResModelId       *Many2One  `xmlrpc:"res_model_id,omitempty"`
	ResName          *String    `xmlrpc:"res_name,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// RatingRatings represents array of rating.rating model.
type RatingRatings []RatingRating

// RatingRatingModel is the odoo model name.
const RatingRatingModel = "rating.rating"

// Many2One convert RatingRating to *Many2One.
func (rr *RatingRating) Many2One() *Many2One {
	return NewMany2One(rr.Id.Get(), "")
}

// CreateRatingRating creates a new rating.rating model and returns its id.
func (c *Client) CreateRatingRating(rr *RatingRating) (int64, error) {
	return c.Create(RatingRatingModel, rr)
}

// UpdateRatingRating updates an existing rating.rating record.
func (c *Client) UpdateRatingRating(rr *RatingRating) error {
	return c.UpdateRatingRatings([]int64{rr.Id.Get()}, rr)
}

// UpdateRatingRatings updates existing rating.rating records.
// All records (represented by ids) will be updated by rr values.
func (c *Client) UpdateRatingRatings(ids []int64, rr *RatingRating) error {
	return c.Update(RatingRatingModel, ids, rr)
}

// DeleteRatingRating deletes an existing rating.rating record.
func (c *Client) DeleteRatingRating(id int64) error {
	return c.DeleteRatingRatings([]int64{id})
}

// DeleteRatingRatings deletes existing rating.rating records.
func (c *Client) DeleteRatingRatings(ids []int64) error {
	return c.Delete(RatingRatingModel, ids)
}

// GetRatingRating gets rating.rating existing record.
func (c *Client) GetRatingRating(id int64) (*RatingRating, error) {
	rrs, err := c.GetRatingRatings([]int64{id})
	if err != nil {
		return nil, err
	}
	if rrs != nil && len(*rrs) > 0 {
		return &((*rrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of rating.rating not found", id)
}

// GetRatingRatings gets rating.rating existing records.
func (c *Client) GetRatingRatings(ids []int64) (*RatingRatings, error) {
	rrs := &RatingRatings{}
	if err := c.Read(RatingRatingModel, ids, nil, rrs); err != nil {
		return nil, err
	}
	return rrs, nil
}

// FindRatingRating finds rating.rating record by querying it with criteria.
func (c *Client) FindRatingRating(criteria *Criteria) (*RatingRating, error) {
	rrs := &RatingRatings{}
	if err := c.SearchRead(RatingRatingModel, criteria, NewOptions().Limit(1), rrs); err != nil {
		return nil, err
	}
	if rrs != nil && len(*rrs) > 0 {
		return &((*rrs)[0]), nil
	}
	return nil, fmt.Errorf("no rating.rating was found with criteria %v", criteria)
}

// FindRatingRatings finds rating.rating records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRatingRatings(criteria *Criteria, options *Options) (*RatingRatings, error) {
	rrs := &RatingRatings{}
	if err := c.SearchRead(RatingRatingModel, criteria, options, rrs); err != nil {
		return nil, err
	}
	return rrs, nil
}

// FindRatingRatingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRatingRatingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(RatingRatingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindRatingRatingId finds record id by querying it with criteria.
func (c *Client) FindRatingRatingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RatingRatingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no rating.rating was found with criteria %v and options %v", criteria, options)
}
