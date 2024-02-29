package odoo

// MailTrackingValue represents mail.tracking.value model.
type MailTrackingValue struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Field            *String   `xmlrpc:"field,omitempty"`
	FieldDesc        *String   `xmlrpc:"field_desc,omitempty"`
	FieldType        *String   `xmlrpc:"field_type,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	MailMessageId    *Many2One `xmlrpc:"mail_message_id,omitempty"`
	NewValueChar     *String   `xmlrpc:"new_value_char,omitempty"`
	NewValueDatetime *Time     `xmlrpc:"new_value_datetime,omitempty"`
	NewValueFloat    *Float    `xmlrpc:"new_value_float,omitempty"`
	NewValueInteger  *Int      `xmlrpc:"new_value_integer,omitempty"`
	NewValueMonetary *Float    `xmlrpc:"new_value_monetary,omitempty"`
	NewValueText     *String   `xmlrpc:"new_value_text,omitempty"`
	OldValueChar     *String   `xmlrpc:"old_value_char,omitempty"`
	OldValueDatetime *Time     `xmlrpc:"old_value_datetime,omitempty"`
	OldValueFloat    *Float    `xmlrpc:"old_value_float,omitempty"`
	OldValueInteger  *Int      `xmlrpc:"old_value_integer,omitempty"`
	OldValueMonetary *Float    `xmlrpc:"old_value_monetary,omitempty"`
	OldValueText     *String   `xmlrpc:"old_value_text,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailTrackingValues represents array of mail.tracking.value model.
type MailTrackingValues []MailTrackingValue

// MailTrackingValueModel is the odoo model name.
const MailTrackingValueModel = "mail.tracking.value"

// Many2One convert MailTrackingValue to *Many2One.
func (mtv *MailTrackingValue) Many2One() *Many2One {
	return NewMany2One(mtv.Id.Get(), "")
}

// CreateMailTrackingValue creates a new mail.tracking.value model and returns its id.
func (c *Client) CreateMailTrackingValue(mtv *MailTrackingValue) (int64, error) {
	ids, err := c.CreateMailTrackingValues([]*MailTrackingValue{mtv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailTrackingValues creates a new mail.tracking.value model and returns its id.
func (c *Client) CreateMailTrackingValues(mtvs []*MailTrackingValue) ([]int64, error) {
	var vv []interface{}
	for _, v := range mtvs {
		vv = append(vv, v)
	}
	return c.Create(MailTrackingValueModel, vv, nil)
}

// UpdateMailTrackingValue updates an existing mail.tracking.value record.
func (c *Client) UpdateMailTrackingValue(mtv *MailTrackingValue) error {
	return c.UpdateMailTrackingValues([]int64{mtv.Id.Get()}, mtv)
}

// UpdateMailTrackingValues updates existing mail.tracking.value records.
// All records (represented by ids) will be updated by mtv values.
func (c *Client) UpdateMailTrackingValues(ids []int64, mtv *MailTrackingValue) error {
	return c.Update(MailTrackingValueModel, ids, mtv, nil)
}

// DeleteMailTrackingValue deletes an existing mail.tracking.value record.
func (c *Client) DeleteMailTrackingValue(id int64) error {
	return c.DeleteMailTrackingValues([]int64{id})
}

// DeleteMailTrackingValues deletes existing mail.tracking.value records.
func (c *Client) DeleteMailTrackingValues(ids []int64) error {
	return c.Delete(MailTrackingValueModel, ids)
}

// GetMailTrackingValue gets mail.tracking.value existing record.
func (c *Client) GetMailTrackingValue(id int64) (*MailTrackingValue, error) {
	mtvs, err := c.GetMailTrackingValues([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mtvs)[0]), nil
}

// GetMailTrackingValues gets mail.tracking.value existing records.
func (c *Client) GetMailTrackingValues(ids []int64) (*MailTrackingValues, error) {
	mtvs := &MailTrackingValues{}
	if err := c.Read(MailTrackingValueModel, ids, nil, mtvs); err != nil {
		return nil, err
	}
	return mtvs, nil
}

// FindMailTrackingValue finds mail.tracking.value record by querying it with criteria.
func (c *Client) FindMailTrackingValue(criteria *Criteria) (*MailTrackingValue, error) {
	mtvs := &MailTrackingValues{}
	if err := c.SearchRead(MailTrackingValueModel, criteria, NewOptions().Limit(1), mtvs); err != nil {
		return nil, err
	}
	return &((*mtvs)[0]), nil
}

// FindMailTrackingValues finds mail.tracking.value records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTrackingValues(criteria *Criteria, options *Options) (*MailTrackingValues, error) {
	mtvs := &MailTrackingValues{}
	if err := c.SearchRead(MailTrackingValueModel, criteria, options, mtvs); err != nil {
		return nil, err
	}
	return mtvs, nil
}

// FindMailTrackingValueIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTrackingValueIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailTrackingValueModel, criteria, options)
}

// FindMailTrackingValueId finds record id by querying it with criteria.
func (c *Client) FindMailTrackingValueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailTrackingValueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
