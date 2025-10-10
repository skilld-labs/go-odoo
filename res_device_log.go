package odoo

// ResDeviceLog represents res.device.log model.
type ResDeviceLog struct {
	Browser           *String    `xmlrpc:"browser,omitempty"`
	City              *String    `xmlrpc:"city,omitempty"`
	Country           *String    `xmlrpc:"country,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DeviceType        *Selection `xmlrpc:"device_type,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	FirstActivity     *Time      `xmlrpc:"first_activity,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	IpAddress         *String    `xmlrpc:"ip_address,omitempty"`
	IsCurrent         *Bool      `xmlrpc:"is_current,omitempty"`
	LastActivity      *Time      `xmlrpc:"last_activity,omitempty"`
	LinkedIpAddresses *String    `xmlrpc:"linked_ip_addresses,omitempty"`
	Platform          *String    `xmlrpc:"platform,omitempty"`
	Revoked           *Bool      `xmlrpc:"revoked,omitempty"`
	SessionIdentifier *String    `xmlrpc:"session_identifier,omitempty"`
	UserId            *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResDeviceLogs represents array of res.device.log model.
type ResDeviceLogs []ResDeviceLog

// ResDeviceLogModel is the odoo model name.
const ResDeviceLogModel = "res.device.log"

// Many2One convert ResDeviceLog to *Many2One.
func (rdl *ResDeviceLog) Many2One() *Many2One {
	return NewMany2One(rdl.Id.Get(), "")
}

// CreateResDeviceLog creates a new res.device.log model and returns its id.
func (c *Client) CreateResDeviceLog(rdl *ResDeviceLog) (int64, error) {
	ids, err := c.CreateResDeviceLogs([]*ResDeviceLog{rdl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResDeviceLog creates a new res.device.log model and returns its id.
func (c *Client) CreateResDeviceLogs(rdls []*ResDeviceLog) ([]int64, error) {
	var vv []interface{}
	for _, v := range rdls {
		vv = append(vv, v)
	}
	return c.Create(ResDeviceLogModel, vv, nil)
}

// UpdateResDeviceLog updates an existing res.device.log record.
func (c *Client) UpdateResDeviceLog(rdl *ResDeviceLog) error {
	return c.UpdateResDeviceLogs([]int64{rdl.Id.Get()}, rdl)
}

// UpdateResDeviceLogs updates existing res.device.log records.
// All records (represented by ids) will be updated by rdl values.
func (c *Client) UpdateResDeviceLogs(ids []int64, rdl *ResDeviceLog) error {
	return c.Update(ResDeviceLogModel, ids, rdl, nil)
}

// DeleteResDeviceLog deletes an existing res.device.log record.
func (c *Client) DeleteResDeviceLog(id int64) error {
	return c.DeleteResDeviceLogs([]int64{id})
}

// DeleteResDeviceLogs deletes existing res.device.log records.
func (c *Client) DeleteResDeviceLogs(ids []int64) error {
	return c.Delete(ResDeviceLogModel, ids)
}

// GetResDeviceLog gets res.device.log existing record.
func (c *Client) GetResDeviceLog(id int64) (*ResDeviceLog, error) {
	rdls, err := c.GetResDeviceLogs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rdls)[0]), nil
}

// GetResDeviceLogs gets res.device.log existing records.
func (c *Client) GetResDeviceLogs(ids []int64) (*ResDeviceLogs, error) {
	rdls := &ResDeviceLogs{}
	if err := c.Read(ResDeviceLogModel, ids, nil, rdls); err != nil {
		return nil, err
	}
	return rdls, nil
}

// FindResDeviceLog finds res.device.log record by querying it with criteria.
func (c *Client) FindResDeviceLog(criteria *Criteria) (*ResDeviceLog, error) {
	rdls := &ResDeviceLogs{}
	if err := c.SearchRead(ResDeviceLogModel, criteria, NewOptions().Limit(1), rdls); err != nil {
		return nil, err
	}
	return &((*rdls)[0]), nil
}

// FindResDeviceLogs finds res.device.log records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResDeviceLogs(criteria *Criteria, options *Options) (*ResDeviceLogs, error) {
	rdls := &ResDeviceLogs{}
	if err := c.SearchRead(ResDeviceLogModel, criteria, options, rdls); err != nil {
		return nil, err
	}
	return rdls, nil
}

// FindResDeviceLogIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResDeviceLogIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResDeviceLogModel, criteria, options)
}

// FindResDeviceLogId finds record id by querying it with criteria.
func (c *Client) FindResDeviceLogId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResDeviceLogModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
