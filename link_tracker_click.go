package odoo

import (
	"fmt"
)

// LinkTrackerClick represents link.tracker.click model.
type LinkTrackerClick struct {
	LastUpdate            *Time     `xmlrpc:"__last_update,omptempty"`
	ClickDate             *Time     `xmlrpc:"click_date,omptempty"`
	CountryId             *Many2One `xmlrpc:"country_id,omptempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String   `xmlrpc:"display_name,omptempty"`
	Id                    *Int      `xmlrpc:"id,omptempty"`
	Ip                    *String   `xmlrpc:"ip,omptempty"`
	LinkId                *Many2One `xmlrpc:"link_id,omptempty"`
	MailStatId            *Many2One `xmlrpc:"mail_stat_id,omptempty"`
	MassMailingCampaignId *Many2One `xmlrpc:"mass_mailing_campaign_id,omptempty"`
	MassMailingId         *Many2One `xmlrpc:"mass_mailing_id,omptempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omptempty"`
}

// LinkTrackerClicks represents array of link.tracker.click model.
type LinkTrackerClicks []LinkTrackerClick

// LinkTrackerClickModel is the odoo model name.
const LinkTrackerClickModel = "link.tracker.click"

// Many2One convert LinkTrackerClick to *Many2One.
func (ltc *LinkTrackerClick) Many2One() *Many2One {
	return NewMany2One(ltc.Id.Get(), "")
}

// CreateLinkTrackerClick creates a new link.tracker.click model and returns its id.
func (c *Client) CreateLinkTrackerClick(ltc *LinkTrackerClick) (int64, error) {
	ids, err := c.CreateLinkTrackerClicks([]*LinkTrackerClick{ltc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLinkTrackerClick creates a new link.tracker.click model and returns its id.
func (c *Client) CreateLinkTrackerClicks(ltcs []*LinkTrackerClick) ([]int64, error) {
	var vv []interface{}
	for _, v := range ltcs {
		vv = append(vv, v)
	}
	return c.Create(LinkTrackerClickModel, vv)
}

// UpdateLinkTrackerClick updates an existing link.tracker.click record.
func (c *Client) UpdateLinkTrackerClick(ltc *LinkTrackerClick) error {
	return c.UpdateLinkTrackerClicks([]int64{ltc.Id.Get()}, ltc)
}

// UpdateLinkTrackerClicks updates existing link.tracker.click records.
// All records (represented by ids) will be updated by ltc values.
func (c *Client) UpdateLinkTrackerClicks(ids []int64, ltc *LinkTrackerClick) error {
	return c.Update(LinkTrackerClickModel, ids, ltc)
}

// DeleteLinkTrackerClick deletes an existing link.tracker.click record.
func (c *Client) DeleteLinkTrackerClick(id int64) error {
	return c.DeleteLinkTrackerClicks([]int64{id})
}

// DeleteLinkTrackerClicks deletes existing link.tracker.click records.
func (c *Client) DeleteLinkTrackerClicks(ids []int64) error {
	return c.Delete(LinkTrackerClickModel, ids)
}

// GetLinkTrackerClick gets link.tracker.click existing record.
func (c *Client) GetLinkTrackerClick(id int64) (*LinkTrackerClick, error) {
	ltcs, err := c.GetLinkTrackerClicks([]int64{id})
	if err != nil {
		return nil, err
	}
	if ltcs != nil && len(*ltcs) > 0 {
		return &((*ltcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of link.tracker.click not found", id)
}

// GetLinkTrackerClicks gets link.tracker.click existing records.
func (c *Client) GetLinkTrackerClicks(ids []int64) (*LinkTrackerClicks, error) {
	ltcs := &LinkTrackerClicks{}
	if err := c.Read(LinkTrackerClickModel, ids, nil, ltcs); err != nil {
		return nil, err
	}
	return ltcs, nil
}

// FindLinkTrackerClick finds link.tracker.click record by querying it with criteria.
func (c *Client) FindLinkTrackerClick(criteria *Criteria) (*LinkTrackerClick, error) {
	ltcs := &LinkTrackerClicks{}
	if err := c.SearchRead(LinkTrackerClickModel, criteria, NewOptions().Limit(1), ltcs); err != nil {
		return nil, err
	}
	if ltcs != nil && len(*ltcs) > 0 {
		return &((*ltcs)[0]), nil
	}
	return nil, fmt.Errorf("link.tracker.click was not found with criteria %v", criteria)
}

// FindLinkTrackerClicks finds link.tracker.click records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLinkTrackerClicks(criteria *Criteria, options *Options) (*LinkTrackerClicks, error) {
	ltcs := &LinkTrackerClicks{}
	if err := c.SearchRead(LinkTrackerClickModel, criteria, options, ltcs); err != nil {
		return nil, err
	}
	return ltcs, nil
}

// FindLinkTrackerClickIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLinkTrackerClickIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LinkTrackerClickModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLinkTrackerClickId finds record id by querying it with criteria.
func (c *Client) FindLinkTrackerClickId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LinkTrackerClickModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("link.tracker.click was not found with criteria %v and options %v", criteria, options)
}
