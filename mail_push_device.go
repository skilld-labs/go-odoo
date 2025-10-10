package odoo

// MailPushDevice represents mail.push.device model.
type MailPushDevice struct {
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Endpoint       *String   `xmlrpc:"endpoint,omitempty"`
	ExpirationTime *Time     `xmlrpc:"expiration_time,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	Keys           *String   `xmlrpc:"keys,omitempty"`
	PartnerId      *Many2One `xmlrpc:"partner_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailPushDevices represents array of mail.push.device model.
type MailPushDevices []MailPushDevice

// MailPushDeviceModel is the odoo model name.
const MailPushDeviceModel = "mail.push.device"

// Many2One convert MailPushDevice to *Many2One.
func (mpd *MailPushDevice) Many2One() *Many2One {
	return NewMany2One(mpd.Id.Get(), "")
}

// CreateMailPushDevice creates a new mail.push.device model and returns its id.
func (c *Client) CreateMailPushDevice(mpd *MailPushDevice) (int64, error) {
	ids, err := c.CreateMailPushDevices([]*MailPushDevice{mpd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailPushDevice creates a new mail.push.device model and returns its id.
func (c *Client) CreateMailPushDevices(mpds []*MailPushDevice) ([]int64, error) {
	var vv []interface{}
	for _, v := range mpds {
		vv = append(vv, v)
	}
	return c.Create(MailPushDeviceModel, vv, nil)
}

// UpdateMailPushDevice updates an existing mail.push.device record.
func (c *Client) UpdateMailPushDevice(mpd *MailPushDevice) error {
	return c.UpdateMailPushDevices([]int64{mpd.Id.Get()}, mpd)
}

// UpdateMailPushDevices updates existing mail.push.device records.
// All records (represented by ids) will be updated by mpd values.
func (c *Client) UpdateMailPushDevices(ids []int64, mpd *MailPushDevice) error {
	return c.Update(MailPushDeviceModel, ids, mpd, nil)
}

// DeleteMailPushDevice deletes an existing mail.push.device record.
func (c *Client) DeleteMailPushDevice(id int64) error {
	return c.DeleteMailPushDevices([]int64{id})
}

// DeleteMailPushDevices deletes existing mail.push.device records.
func (c *Client) DeleteMailPushDevices(ids []int64) error {
	return c.Delete(MailPushDeviceModel, ids)
}

// GetMailPushDevice gets mail.push.device existing record.
func (c *Client) GetMailPushDevice(id int64) (*MailPushDevice, error) {
	mpds, err := c.GetMailPushDevices([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mpds)[0]), nil
}

// GetMailPushDevices gets mail.push.device existing records.
func (c *Client) GetMailPushDevices(ids []int64) (*MailPushDevices, error) {
	mpds := &MailPushDevices{}
	if err := c.Read(MailPushDeviceModel, ids, nil, mpds); err != nil {
		return nil, err
	}
	return mpds, nil
}

// FindMailPushDevice finds mail.push.device record by querying it with criteria.
func (c *Client) FindMailPushDevice(criteria *Criteria) (*MailPushDevice, error) {
	mpds := &MailPushDevices{}
	if err := c.SearchRead(MailPushDeviceModel, criteria, NewOptions().Limit(1), mpds); err != nil {
		return nil, err
	}
	return &((*mpds)[0]), nil
}

// FindMailPushDevices finds mail.push.device records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailPushDevices(criteria *Criteria, options *Options) (*MailPushDevices, error) {
	mpds := &MailPushDevices{}
	if err := c.SearchRead(MailPushDeviceModel, criteria, options, mpds); err != nil {
		return nil, err
	}
	return mpds, nil
}

// FindMailPushDeviceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailPushDeviceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailPushDeviceModel, criteria, options)
}

// FindMailPushDeviceId finds record id by querying it with criteria.
func (c *Client) FindMailPushDeviceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailPushDeviceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
