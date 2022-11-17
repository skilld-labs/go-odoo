package odoo

import (
	"fmt"
)

// CalendarEvent represents calendar.event model.
type CalendarEvent struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	Active                   *Bool      `xmlrpc:"active,omitempty"`
	ActivityIds              *Relation  `xmlrpc:"activity_ids,omitempty"`
	AlarmIds                 *Relation  `xmlrpc:"alarm_ids,omitempty"`
	Allday                   *Bool      `xmlrpc:"allday,omitempty"`
	AttendeeIds              *Relation  `xmlrpc:"attendee_ids,omitempty"`
	AttendeeStatus           *Selection `xmlrpc:"attendee_status,omitempty"`
	Byday                    *Selection `xmlrpc:"byday,omitempty"`
	CategIds                 *Relation  `xmlrpc:"categ_ids,omitempty"`
	Count                    *Int       `xmlrpc:"count,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	Day                      *Int       `xmlrpc:"day,omitempty"`
	Description              *String    `xmlrpc:"description,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	DisplayStart             *String    `xmlrpc:"display_start,omitempty"`
	DisplayTime              *String    `xmlrpc:"display_time,omitempty"`
	Duration                 *Float     `xmlrpc:"duration,omitempty"`
	EndType                  *Selection `xmlrpc:"end_type,omitempty"`
	FinalDate                *Time      `xmlrpc:"final_date,omitempty"`
	Fr                       *Bool      `xmlrpc:"fr,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	Interval                 *Int       `xmlrpc:"interval,omitempty"`
	IsAttendee               *Bool      `xmlrpc:"is_attendee,omitempty"`
	IsHighlighted            *Bool      `xmlrpc:"is_highlighted,omitempty"`
	Location                 *String    `xmlrpc:"location,omitempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Mo                       *Bool      `xmlrpc:"mo,omitempty"`
	MonthBy                  *Selection `xmlrpc:"month_by,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	OpportunityId            *Many2One  `xmlrpc:"opportunity_id,omitempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerIds               *Relation  `xmlrpc:"partner_ids,omitempty"`
	Privacy                  *Selection `xmlrpc:"privacy,omitempty"`
	Recurrency               *Bool      `xmlrpc:"recurrency,omitempty"`
	RecurrentId              *Int       `xmlrpc:"recurrent_id,omitempty"`
	RecurrentIdDate          *Time      `xmlrpc:"recurrent_id_date,omitempty"`
	ResId                    *Int       `xmlrpc:"res_id,omitempty"`
	ResModel                 *String    `xmlrpc:"res_model,omitempty"`
	ResModelId               *Many2One  `xmlrpc:"res_model_id,omitempty"`
	Rrule                    *String    `xmlrpc:"rrule,omitempty"`
	RruleType                *Selection `xmlrpc:"rrule_type,omitempty"`
	Sa                       *Bool      `xmlrpc:"sa,omitempty"`
	ShowAs                   *Selection `xmlrpc:"show_as,omitempty"`
	Start                    *Time      `xmlrpc:"start,omitempty"`
	StartDate                *Time      `xmlrpc:"start_date,omitempty"`
	StartDatetime            *Time      `xmlrpc:"start_datetime,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty"`
	Stop                     *Time      `xmlrpc:"stop,omitempty"`
	StopDate                 *Time      `xmlrpc:"stop_date,omitempty"`
	StopDatetime             *Time      `xmlrpc:"stop_datetime,omitempty"`
	Su                       *Bool      `xmlrpc:"su,omitempty"`
	Th                       *Bool      `xmlrpc:"th,omitempty"`
	Tu                       *Bool      `xmlrpc:"tu,omitempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omitempty"`
	We                       *Bool      `xmlrpc:"we,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WeekList                 *Selection `xmlrpc:"week_list,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// CalendarEvents represents array of calendar.event model.
type CalendarEvents []CalendarEvent

// CalendarEventModel is the odoo model name.
const CalendarEventModel = "calendar.event"

// Many2One convert CalendarEvent to *Many2One.
func (ce *CalendarEvent) Many2One() *Many2One {
	return NewMany2One(ce.Id.Get(), "")
}

// CreateCalendarEvent creates a new calendar.event model and returns its id.
func (c *Client) CreateCalendarEvent(ce *CalendarEvent) (int64, error) {
	return c.Create(CalendarEventModel, ce)
}

// UpdateCalendarEvent updates an existing calendar.event record.
func (c *Client) UpdateCalendarEvent(ce *CalendarEvent) error {
	return c.UpdateCalendarEvents([]int64{ce.Id.Get()}, ce)
}

// UpdateCalendarEvents updates existing calendar.event records.
// All records (represented by ids) will be updated by ce values.
func (c *Client) UpdateCalendarEvents(ids []int64, ce *CalendarEvent) error {
	return c.Update(CalendarEventModel, ids, ce)
}

// DeleteCalendarEvent deletes an existing calendar.event record.
func (c *Client) DeleteCalendarEvent(id int64) error {
	return c.DeleteCalendarEvents([]int64{id})
}

// DeleteCalendarEvents deletes existing calendar.event records.
func (c *Client) DeleteCalendarEvents(ids []int64) error {
	return c.Delete(CalendarEventModel, ids)
}

// GetCalendarEvent gets calendar.event existing record.
func (c *Client) GetCalendarEvent(id int64) (*CalendarEvent, error) {
	ces, err := c.GetCalendarEvents([]int64{id})
	if err != nil {
		return nil, err
	}
	if ces != nil && len(*ces) > 0 {
		return &((*ces)[0]), nil
	}
	return nil, fmt.Errorf("id %v of calendar.event not found", id)
}

// GetCalendarEvents gets calendar.event existing records.
func (c *Client) GetCalendarEvents(ids []int64) (*CalendarEvents, error) {
	ces := &CalendarEvents{}
	if err := c.Read(CalendarEventModel, ids, nil, ces); err != nil {
		return nil, err
	}
	return ces, nil
}

// FindCalendarEvent finds calendar.event record by querying it with criteria.
func (c *Client) FindCalendarEvent(criteria *Criteria) (*CalendarEvent, error) {
	ces := &CalendarEvents{}
	if err := c.SearchRead(CalendarEventModel, criteria, NewOptions().Limit(1), ces); err != nil {
		return nil, err
	}
	if ces != nil && len(*ces) > 0 {
		return &((*ces)[0]), nil
	}
	return nil, fmt.Errorf("no calendar.event was found with criteria %v", criteria)
}

// FindCalendarEvents finds calendar.event records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarEvents(criteria *Criteria, options *Options) (*CalendarEvents, error) {
	ces := &CalendarEvents{}
	if err := c.SearchRead(CalendarEventModel, criteria, options, ces); err != nil {
		return nil, err
	}
	return ces, nil
}

// FindCalendarEventIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarEventIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CalendarEventModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCalendarEventId finds record id by querying it with criteria.
func (c *Client) FindCalendarEventId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarEventModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no calendar.event was found with criteria %v and options %v", criteria, options)
}
