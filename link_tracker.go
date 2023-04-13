package odoo

import (
	"fmt"
)

// LinkTracker represents link.tracker model.
type LinkTracker struct {
	LastUpdate            *Time     `xmlrpc:"__last_update,omptempty"`
	CampaignId            *Many2One `xmlrpc:"campaign_id,omptempty"`
	Code                  *String   `xmlrpc:"code,omptempty"`
	Count                 *Int      `xmlrpc:"count,omptempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String   `xmlrpc:"display_name,omptempty"`
	Favicon               *String   `xmlrpc:"favicon,omptempty"`
	IconSrc               *String   `xmlrpc:"icon_src,omptempty"`
	Id                    *Int      `xmlrpc:"id,omptempty"`
	LinkClickIds          *Relation `xmlrpc:"link_click_ids,omptempty"`
	LinkCodeIds           *Relation `xmlrpc:"link_code_ids,omptempty"`
	MassMailingCampaignId *Many2One `xmlrpc:"mass_mailing_campaign_id,omptempty"`
	MassMailingId         *Many2One `xmlrpc:"mass_mailing_id,omptempty"`
	MediumId              *Many2One `xmlrpc:"medium_id,omptempty"`
	RedirectedUrl         *String   `xmlrpc:"redirected_url,omptempty"`
	ShortUrl              *String   `xmlrpc:"short_url,omptempty"`
	ShortUrlHost          *String   `xmlrpc:"short_url_host,omptempty"`
	SourceId              *Many2One `xmlrpc:"source_id,omptempty"`
	Title                 *String   `xmlrpc:"title,omptempty"`
	Url                   *String   `xmlrpc:"url,omptempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omptempty"`
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
	ids, err := c.Create(LinkTrackerModel, []interface{}{lt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLinkTracker creates a new link.tracker model and returns its id.
func (c *Client) CreateLinkTrackers(lts []*LinkTracker) ([]int64, error) {
	var vv []interface{}
	for _, v := range lts {
		vv = append(vv, v)
	}
	return c.Create(LinkTrackerModel, vv)
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
	return nil, fmt.Errorf("link.tracker was not found with criteria %v", criteria)
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
	return -1, fmt.Errorf("link.tracker was not found with criteria %v and options %v", criteria, options)
}
