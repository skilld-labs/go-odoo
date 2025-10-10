package odoo

// CalendarProviderConfig represents calendar.provider.config model.
type CalendarProviderConfig struct {
	CalClientId                      *String    `xmlrpc:"cal_client_id,omitempty"`
	CalClientSecret                  *String    `xmlrpc:"cal_client_secret,omitempty"`
	CalSyncPaused                    *Bool      `xmlrpc:"cal_sync_paused,omitempty"`
	CreateDate                       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                        *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                      *String    `xmlrpc:"display_name,omitempty"`
	ExternalCalendarProvider         *Selection `xmlrpc:"external_calendar_provider,omitempty"`
	Id                               *Int       `xmlrpc:"id,omitempty"`
	MicrosoftOutlookClientIdentifier *String    `xmlrpc:"microsoft_outlook_client_identifier,omitempty"`
	MicrosoftOutlookClientSecret     *String    `xmlrpc:"microsoft_outlook_client_secret,omitempty"`
	MicrosoftOutlookSyncPaused       *Bool      `xmlrpc:"microsoft_outlook_sync_paused,omitempty"`
	WriteDate                        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// CalendarProviderConfigs represents array of calendar.provider.config model.
type CalendarProviderConfigs []CalendarProviderConfig

// CalendarProviderConfigModel is the odoo model name.
const CalendarProviderConfigModel = "calendar.provider.config"

// Many2One convert CalendarProviderConfig to *Many2One.
func (cpc *CalendarProviderConfig) Many2One() *Many2One {
	return NewMany2One(cpc.Id.Get(), "")
}

// CreateCalendarProviderConfig creates a new calendar.provider.config model and returns its id.
func (c *Client) CreateCalendarProviderConfig(cpc *CalendarProviderConfig) (int64, error) {
	ids, err := c.CreateCalendarProviderConfigs([]*CalendarProviderConfig{cpc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarProviderConfig creates a new calendar.provider.config model and returns its id.
func (c *Client) CreateCalendarProviderConfigs(cpcs []*CalendarProviderConfig) ([]int64, error) {
	var vv []interface{}
	for _, v := range cpcs {
		vv = append(vv, v)
	}
	return c.Create(CalendarProviderConfigModel, vv, nil)
}

// UpdateCalendarProviderConfig updates an existing calendar.provider.config record.
func (c *Client) UpdateCalendarProviderConfig(cpc *CalendarProviderConfig) error {
	return c.UpdateCalendarProviderConfigs([]int64{cpc.Id.Get()}, cpc)
}

// UpdateCalendarProviderConfigs updates existing calendar.provider.config records.
// All records (represented by ids) will be updated by cpc values.
func (c *Client) UpdateCalendarProviderConfigs(ids []int64, cpc *CalendarProviderConfig) error {
	return c.Update(CalendarProviderConfigModel, ids, cpc, nil)
}

// DeleteCalendarProviderConfig deletes an existing calendar.provider.config record.
func (c *Client) DeleteCalendarProviderConfig(id int64) error {
	return c.DeleteCalendarProviderConfigs([]int64{id})
}

// DeleteCalendarProviderConfigs deletes existing calendar.provider.config records.
func (c *Client) DeleteCalendarProviderConfigs(ids []int64) error {
	return c.Delete(CalendarProviderConfigModel, ids)
}

// GetCalendarProviderConfig gets calendar.provider.config existing record.
func (c *Client) GetCalendarProviderConfig(id int64) (*CalendarProviderConfig, error) {
	cpcs, err := c.GetCalendarProviderConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cpcs)[0]), nil
}

// GetCalendarProviderConfigs gets calendar.provider.config existing records.
func (c *Client) GetCalendarProviderConfigs(ids []int64) (*CalendarProviderConfigs, error) {
	cpcs := &CalendarProviderConfigs{}
	if err := c.Read(CalendarProviderConfigModel, ids, nil, cpcs); err != nil {
		return nil, err
	}
	return cpcs, nil
}

// FindCalendarProviderConfig finds calendar.provider.config record by querying it with criteria.
func (c *Client) FindCalendarProviderConfig(criteria *Criteria) (*CalendarProviderConfig, error) {
	cpcs := &CalendarProviderConfigs{}
	if err := c.SearchRead(CalendarProviderConfigModel, criteria, NewOptions().Limit(1), cpcs); err != nil {
		return nil, err
	}
	return &((*cpcs)[0]), nil
}

// FindCalendarProviderConfigs finds calendar.provider.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarProviderConfigs(criteria *Criteria, options *Options) (*CalendarProviderConfigs, error) {
	cpcs := &CalendarProviderConfigs{}
	if err := c.SearchRead(CalendarProviderConfigModel, criteria, options, cpcs); err != nil {
		return nil, err
	}
	return cpcs, nil
}

// FindCalendarProviderConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarProviderConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CalendarProviderConfigModel, criteria, options)
}

// FindCalendarProviderConfigId finds record id by querying it with criteria.
func (c *Client) FindCalendarProviderConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarProviderConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
