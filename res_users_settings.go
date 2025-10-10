package odoo

// ResUsersSettings represents res.users.settings model.
type ResUsersSettings struct {
	CalendarDefaultPrivacy              *Selection `xmlrpc:"calendar_default_privacy,omitempty"`
	ChannelNotifications                *Selection `xmlrpc:"channel_notifications,omitempty"`
	CreateDate                          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                         *String    `xmlrpc:"display_name,omitempty"`
	Id                                  *Int       `xmlrpc:"id,omitempty"`
	IsDiscussSidebarCategoryChannelOpen *Bool      `xmlrpc:"is_discuss_sidebar_category_channel_open,omitempty"`
	IsDiscussSidebarCategoryChatOpen    *Bool      `xmlrpc:"is_discuss_sidebar_category_chat_open,omitempty"`
	LivechatLangIds                     *Relation  `xmlrpc:"livechat_lang_ids,omitempty"`
	LivechatUsername                    *String    `xmlrpc:"livechat_username,omitempty"`
	MuteUntilDt                         *Time      `xmlrpc:"mute_until_dt,omitempty"`
	PushToTalkKey                       *String    `xmlrpc:"push_to_talk_key,omitempty"`
	UsePushToTalk                       *Bool      `xmlrpc:"use_push_to_talk,omitempty"`
	UserId                              *Many2One  `xmlrpc:"user_id,omitempty"`
	VoiceActiveDuration                 *Int       `xmlrpc:"voice_active_duration,omitempty"`
	VolumeSettingsIds                   *Relation  `xmlrpc:"volume_settings_ids,omitempty"`
	WriteDate                           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResUsersSettingss represents array of res.users.settings model.
type ResUsersSettingss []ResUsersSettings

// ResUsersSettingsModel is the odoo model name.
const ResUsersSettingsModel = "res.users.settings"

// Many2One convert ResUsersSettings to *Many2One.
func (rus *ResUsersSettings) Many2One() *Many2One {
	return NewMany2One(rus.Id.Get(), "")
}

// CreateResUsersSettings creates a new res.users.settings model and returns its id.
func (c *Client) CreateResUsersSettings(rus *ResUsersSettings) (int64, error) {
	ids, err := c.CreateResUsersSettingss([]*ResUsersSettings{rus})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsersSettings creates a new res.users.settings model and returns its id.
func (c *Client) CreateResUsersSettingss(russ []*ResUsersSettings) ([]int64, error) {
	var vv []interface{}
	for _, v := range russ {
		vv = append(vv, v)
	}
	return c.Create(ResUsersSettingsModel, vv, nil)
}

// UpdateResUsersSettings updates an existing res.users.settings record.
func (c *Client) UpdateResUsersSettings(rus *ResUsersSettings) error {
	return c.UpdateResUsersSettingss([]int64{rus.Id.Get()}, rus)
}

// UpdateResUsersSettingss updates existing res.users.settings records.
// All records (represented by ids) will be updated by rus values.
func (c *Client) UpdateResUsersSettingss(ids []int64, rus *ResUsersSettings) error {
	return c.Update(ResUsersSettingsModel, ids, rus, nil)
}

// DeleteResUsersSettings deletes an existing res.users.settings record.
func (c *Client) DeleteResUsersSettings(id int64) error {
	return c.DeleteResUsersSettingss([]int64{id})
}

// DeleteResUsersSettingss deletes existing res.users.settings records.
func (c *Client) DeleteResUsersSettingss(ids []int64) error {
	return c.Delete(ResUsersSettingsModel, ids)
}

// GetResUsersSettings gets res.users.settings existing record.
func (c *Client) GetResUsersSettings(id int64) (*ResUsersSettings, error) {
	russ, err := c.GetResUsersSettingss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*russ)[0]), nil
}

// GetResUsersSettingss gets res.users.settings existing records.
func (c *Client) GetResUsersSettingss(ids []int64) (*ResUsersSettingss, error) {
	russ := &ResUsersSettingss{}
	if err := c.Read(ResUsersSettingsModel, ids, nil, russ); err != nil {
		return nil, err
	}
	return russ, nil
}

// FindResUsersSettings finds res.users.settings record by querying it with criteria.
func (c *Client) FindResUsersSettings(criteria *Criteria) (*ResUsersSettings, error) {
	russ := &ResUsersSettingss{}
	if err := c.SearchRead(ResUsersSettingsModel, criteria, NewOptions().Limit(1), russ); err != nil {
		return nil, err
	}
	return &((*russ)[0]), nil
}

// FindResUsersSettingss finds res.users.settings records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersSettingss(criteria *Criteria, options *Options) (*ResUsersSettingss, error) {
	russ := &ResUsersSettingss{}
	if err := c.SearchRead(ResUsersSettingsModel, criteria, options, russ); err != nil {
		return nil, err
	}
	return russ, nil
}

// FindResUsersSettingsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersSettingsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResUsersSettingsModel, criteria, options)
}

// FindResUsersSettingsId finds record id by querying it with criteria.
func (c *Client) FindResUsersSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
