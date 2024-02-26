package odoo

// CalendarEvent represents calendar.event model.
type CalendarEvent struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	Active                   *Bool      `xmlrpc:"active,omptempty"`
	ActivityIds              *Relation  `xmlrpc:"activity_ids,omptempty"`
	AlarmIds                 *Relation  `xmlrpc:"alarm_ids,omptempty"`
	Allday                   *Bool      `xmlrpc:"allday,omptempty"`
	AttendeeIds              *Relation  `xmlrpc:"attendee_ids,omptempty"`
	AttendeeStatus           *Selection `xmlrpc:"attendee_status,omptempty"`
	Byday                    *Selection `xmlrpc:"byday,omptempty"`
	CategIds                 *Relation  `xmlrpc:"categ_ids,omptempty"`
	Count                    *Int       `xmlrpc:"count,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	Day                      *Int       `xmlrpc:"day,omptempty"`
	Description              *String    `xmlrpc:"description,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	DisplayStart             *String    `xmlrpc:"display_start,omptempty"`
	DisplayTime              *String    `xmlrpc:"display_time,omptempty"`
	Duration                 *Float     `xmlrpc:"duration,omptempty"`
	EndType                  *Selection `xmlrpc:"end_type,omptempty"`
	FinalDate                *Time      `xmlrpc:"final_date,omptempty"`
	Fr                       *Bool      `xmlrpc:"fr,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	Interval                 *Int       `xmlrpc:"interval,omptempty"`
	IsAttendee               *Bool      `xmlrpc:"is_attendee,omptempty"`
	IsHighlighted            *Bool      `xmlrpc:"is_highlighted,omptempty"`
	Location                 *String    `xmlrpc:"location,omptempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Mo                       *Bool      `xmlrpc:"mo,omptempty"`
	MonthBy                  *Selection `xmlrpc:"month_by,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	OpportunityId            *Many2One  `xmlrpc:"opportunity_id,omptempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerIds               *Relation  `xmlrpc:"partner_ids,omptempty"`
	Privacy                  *Selection `xmlrpc:"privacy,omptempty"`
	Recurrency               *Bool      `xmlrpc:"recurrency,omptempty"`
	RecurrentId              *Int       `xmlrpc:"recurrent_id,omptempty"`
	RecurrentIdDate          *Time      `xmlrpc:"recurrent_id_date,omptempty"`
	ResId                    *Int       `xmlrpc:"res_id,omptempty"`
	ResModel                 *String    `xmlrpc:"res_model,omptempty"`
	ResModelId               *Many2One  `xmlrpc:"res_model_id,omptempty"`
	Rrule                    *String    `xmlrpc:"rrule,omptempty"`
	RruleType                *Selection `xmlrpc:"rrule_type,omptempty"`
	Sa                       *Bool      `xmlrpc:"sa,omptempty"`
	ShowAs                   *Selection `xmlrpc:"show_as,omptempty"`
	Start                    *Time      `xmlrpc:"start,omptempty"`
	StartDate                *Time      `xmlrpc:"start_date,omptempty"`
	StartDatetime            *Time      `xmlrpc:"start_datetime,omptempty"`
	State                    *Selection `xmlrpc:"state,omptempty"`
	Stop                     *Time      `xmlrpc:"stop,omptempty"`
	StopDate                 *Time      `xmlrpc:"stop_date,omptempty"`
	StopDatetime             *Time      `xmlrpc:"stop_datetime,omptempty"`
	Su                       *Bool      `xmlrpc:"su,omptempty"`
	Th                       *Bool      `xmlrpc:"th,omptempty"`
	Tu                       *Bool      `xmlrpc:"tu,omptempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omptempty"`
	We                       *Bool      `xmlrpc:"we,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WeekList                 *Selection `xmlrpc:"week_list,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	ids, err := c.CreateCalendarEvents([]*CalendarEvent{ce})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarEvent creates a new calendar.event model and returns its id.
func (c *Client) CreateCalendarEvents(ces []*CalendarEvent) ([]int64, error) {
	var vv []interface{}
	for _, v := range ces {
		vv = append(vv, v)
	}
	return c.Create(CalendarEventModel, vv, nil)
}

// UpdateCalendarEvent updates an existing calendar.event record.
func (c *Client) UpdateCalendarEvent(ce *CalendarEvent) error {
	return c.UpdateCalendarEvents([]int64{ce.Id.Get()}, ce)
}

// UpdateCalendarEvents updates existing calendar.event records.
// All records (represented by ids) will be updated by ce values.
func (c *Client) UpdateCalendarEvents(ids []int64, ce *CalendarEvent) error {
	return c.Update(CalendarEventModel, ids, ce, nil)
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
	return &((*ces)[0]), nil
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
	return &((*ces)[0]), nil
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
	return c.Search(CalendarEventModel, criteria, options)
}

// FindCalendarEventId finds record id by querying it with criteria.
func (c *Client) FindCalendarEventId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarEventModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
