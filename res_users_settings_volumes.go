package odoo

// ResUsersSettingsVolumes represents res.users.settings.volumes model.
type ResUsersSettingsVolumes struct {
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	GuestId       *Many2One `xmlrpc:"guest_id,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	PartnerId     *Many2One `xmlrpc:"partner_id,omitempty"`
	UserSettingId *Many2One `xmlrpc:"user_setting_id,omitempty"`
	Volume        *Float    `xmlrpc:"volume,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResUsersSettingsVolumess represents array of res.users.settings.volumes model.
type ResUsersSettingsVolumess []ResUsersSettingsVolumes

// ResUsersSettingsVolumesModel is the odoo model name.
const ResUsersSettingsVolumesModel = "res.users.settings.volumes"

// Many2One convert ResUsersSettingsVolumes to *Many2One.
func (rusv *ResUsersSettingsVolumes) Many2One() *Many2One {
	return NewMany2One(rusv.Id.Get(), "")
}

// CreateResUsersSettingsVolumes creates a new res.users.settings.volumes model and returns its id.
func (c *Client) CreateResUsersSettingsVolumes(rusv *ResUsersSettingsVolumes) (int64, error) {
	ids, err := c.CreateResUsersSettingsVolumess([]*ResUsersSettingsVolumes{rusv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsersSettingsVolumes creates a new res.users.settings.volumes model and returns its id.
func (c *Client) CreateResUsersSettingsVolumess(rusvs []*ResUsersSettingsVolumes) ([]int64, error) {
	var vv []interface{}
	for _, v := range rusvs {
		vv = append(vv, v)
	}
	return c.Create(ResUsersSettingsVolumesModel, vv, nil)
}

// UpdateResUsersSettingsVolumes updates an existing res.users.settings.volumes record.
func (c *Client) UpdateResUsersSettingsVolumes(rusv *ResUsersSettingsVolumes) error {
	return c.UpdateResUsersSettingsVolumess([]int64{rusv.Id.Get()}, rusv)
}

// UpdateResUsersSettingsVolumess updates existing res.users.settings.volumes records.
// All records (represented by ids) will be updated by rusv values.
func (c *Client) UpdateResUsersSettingsVolumess(ids []int64, rusv *ResUsersSettingsVolumes) error {
	return c.Update(ResUsersSettingsVolumesModel, ids, rusv, nil)
}

// DeleteResUsersSettingsVolumes deletes an existing res.users.settings.volumes record.
func (c *Client) DeleteResUsersSettingsVolumes(id int64) error {
	return c.DeleteResUsersSettingsVolumess([]int64{id})
}

// DeleteResUsersSettingsVolumess deletes existing res.users.settings.volumes records.
func (c *Client) DeleteResUsersSettingsVolumess(ids []int64) error {
	return c.Delete(ResUsersSettingsVolumesModel, ids)
}

// GetResUsersSettingsVolumes gets res.users.settings.volumes existing record.
func (c *Client) GetResUsersSettingsVolumes(id int64) (*ResUsersSettingsVolumes, error) {
	rusvs, err := c.GetResUsersSettingsVolumess([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rusvs)[0]), nil
}

// GetResUsersSettingsVolumess gets res.users.settings.volumes existing records.
func (c *Client) GetResUsersSettingsVolumess(ids []int64) (*ResUsersSettingsVolumess, error) {
	rusvs := &ResUsersSettingsVolumess{}
	if err := c.Read(ResUsersSettingsVolumesModel, ids, nil, rusvs); err != nil {
		return nil, err
	}
	return rusvs, nil
}

// FindResUsersSettingsVolumes finds res.users.settings.volumes record by querying it with criteria.
func (c *Client) FindResUsersSettingsVolumes(criteria *Criteria) (*ResUsersSettingsVolumes, error) {
	rusvs := &ResUsersSettingsVolumess{}
	if err := c.SearchRead(ResUsersSettingsVolumesModel, criteria, NewOptions().Limit(1), rusvs); err != nil {
		return nil, err
	}
	return &((*rusvs)[0]), nil
}

// FindResUsersSettingsVolumess finds res.users.settings.volumes records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersSettingsVolumess(criteria *Criteria, options *Options) (*ResUsersSettingsVolumess, error) {
	rusvs := &ResUsersSettingsVolumess{}
	if err := c.SearchRead(ResUsersSettingsVolumesModel, criteria, options, rusvs); err != nil {
		return nil, err
	}
	return rusvs, nil
}

// FindResUsersSettingsVolumesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersSettingsVolumesIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResUsersSettingsVolumesModel, criteria, options)
}

// FindResUsersSettingsVolumesId finds record id by querying it with criteria.
func (c *Client) FindResUsersSettingsVolumesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersSettingsVolumesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
