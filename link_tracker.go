package odoo

import (
	"fmt"
)

// LinkTracker represents link.tracker model.
type LinkTracker struct {
	LastUpdate            *Time     `xmlrpc:"__last_update,omitempty"`
	CampaignId            *Many2One `xmlrpc:"campaign_id,omitempty"`
	Code                  *String   `xmlrpc:"code,omitempty"`
	Count                 *Int      `xmlrpc:"count,omitempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Favicon               *String   `xmlrpc:"favicon,omitempty"`
	IconSrc               *String   `xmlrpc:"icon_src,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	LinkClickIds          *Relation `xmlrpc:"link_click_ids,omitempty"`
	LinkCodeIds           *Relation `xmlrpc:"link_code_ids,omitempty"`
	MassMailingCampaignId *Many2One `xmlrpc:"mass_mailing_campaign_id,omitempty"`
	MassMailingId         *Many2One `xmlrpc:"mass_mailing_id,omitempty"`
	MediumId              *Many2One `xmlrpc:"medium_id,omitempty"`
	RedirectedUrl         *String   `xmlrpc:"redirected_url,omitempty"`
	ShortUrl              *String   `xmlrpc:"short_url,omitempty"`
	ShortUrlHost          *String   `xmlrpc:"short_url_host,omitempty"`
	SourceId              *Many2One `xmlrpc:"source_id,omitempty"`
	Title                 *String   `xmlrpc:"title,omitempty"`
	Url                   *String   `xmlrpc:"url,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// LinkTrackers represents array of link.tracker model.
type LinkTrackers []LinkTracker

// LinkTrackerModel is the odoo model name.
const LinkTrackerModel = "link.tracker"

// Many2One convert LinkTracker to *Many2One.
func (lt *LinkTracker) Many2One() *Many2One {
	return NewMany2One(lt.Id.Get(), "")
}

// CreateLinkTracker creates a new link.tracker model and returns its id.
func (c *Client) CreateLinkTracker(lt *LinkTracker) (int64, error) {
	return c.Create(LinkTrackerModel, lt)
}

// UpdateLinkTracker updates an existing link.tracker record.
func (c *Client) UpdateLinkTracker(lt *LinkTracker) error {
	return c.UpdateLinkTrackers([]int64{lt.Id.Get()}, lt)
}

// UpdateLinkTrackers updates existing link.tracker records.
// All records (represented by ids) will be updated by lt values.
func (c *Client) UpdateLinkTrackers(ids []int64, lt *LinkTracker) error {
	return c.Update(LinkTrackerModel, ids, lt)
}

// DeleteLinkTracker deletes an existing link.tracker record.
func (c *Client) DeleteLinkTracker(id int64) error {
	return c.DeleteLinkTrackers([]int64{id})
}

// DeleteLinkTrackers deletes existing link.tracker records.
func (c *Client) DeleteLinkTrackers(ids []int64) error {
	return c.Delete(LinkTrackerModel, ids)
}

// GetLinkTracker gets link.tracker existing record.
func (c *Client) GetLinkTracker(id int64) (*LinkTracker, error) {
	lts, err := c.GetLinkTrackers([]int64{id})
	if err != nil {
		return nil, err
	}
	if lts != nil && len(*lts) > 0 {
		return &((*lts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of link.tracker not found", id)
}

// GetLinkTrackers gets link.tracker existing records.
func (c *Client) GetLinkTrackers(ids []int64) (*LinkTrackers, error) {
	lts := &LinkTrackers{}
	if err := c.Read(LinkTrackerModel, ids, nil, lts); err != nil {
		return nil, err
	}
	return lts, nil
}

// FindLinkTracker finds link.tracker record by querying it with criteria.
func (c *Client) FindLinkTracker(criteria *Criteria) (*LinkTracker, error) {
	lts := &LinkTrackers{}
	if err := c.SearchRead(LinkTrackerModel, criteria, NewOptions().Limit(1), lts); err != nil {
		return nil, err
	}
	if lts != nil && len(*lts) > 0 {
		return &((*lts)[0]), nil
	}
	return nil, fmt.Errorf("no link.tracker was found with criteria %v", criteria)
}

// FindLinkTrackers finds link.tracker records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLinkTrackers(criteria *Criteria, options *Options) (*LinkTrackers, error) {
	lts := &LinkTrackers{}
	if err := c.SearchRead(LinkTrackerModel, criteria, options, lts); err != nil {
		return nil, err
	}
	return lts, nil
}

// FindLinkTrackerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLinkTrackerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LinkTrackerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLinkTrackerId finds record id by querying it with criteria.
func (c *Client) FindLinkTrackerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LinkTrackerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no link.tracker was found with criteria %v and options %v", criteria, options)
}
