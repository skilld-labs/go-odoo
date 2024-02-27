package odoo

// IrQwebFieldContact represents ir.qweb.field.contact model.
type IrQwebFieldContact struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldContacts represents array of ir.qweb.field.contact model.
type IrQwebFieldContacts []IrQwebFieldContact

// IrQwebFieldContactModel is the odoo model name.
const IrQwebFieldContactModel = "ir.qweb.field.contact"

// Many2One convert IrQwebFieldContact to *Many2One.
func (iqfc *IrQwebFieldContact) Many2One() *Many2One {
	return NewMany2One(iqfc.Id.Get(), "")
}

// CreateIrQwebFieldContact creates a new ir.qweb.field.contact model and returns its id.
func (c *Client) CreateIrQwebFieldContact(iqfc *IrQwebFieldContact) (int64, error) {
	ids, err := c.CreateIrQwebFieldContacts([]*IrQwebFieldContact{iqfc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrQwebFieldContacts creates a new ir.qweb.field.contact model and returns its id.
func (c *Client) CreateIrQwebFieldContacts(iqfcs []*IrQwebFieldContact) ([]int64, error) {
	var vv []interface{}
	for _, v := range iqfcs {
		vv = append(vv, v)
	}
	return c.Create(IrQwebFieldContactModel, vv, nil)
}

// UpdateIrQwebFieldContact updates an existing ir.qweb.field.contact record.
func (c *Client) UpdateIrQwebFieldContact(iqfc *IrQwebFieldContact) error {
	return c.UpdateIrQwebFieldContacts([]int64{iqfc.Id.Get()}, iqfc)
}

// UpdateIrQwebFieldContacts updates existing ir.qweb.field.contact records.
// All records (represented by ids) will be updated by iqfc values.
func (c *Client) UpdateIrQwebFieldContacts(ids []int64, iqfc *IrQwebFieldContact) error {
	return c.Update(IrQwebFieldContactModel, ids, iqfc, nil)
}

// DeleteIrQwebFieldContact deletes an existing ir.qweb.field.contact record.
func (c *Client) DeleteIrQwebFieldContact(id int64) error {
	return c.DeleteIrQwebFieldContacts([]int64{id})
}

// DeleteIrQwebFieldContacts deletes existing ir.qweb.field.contact records.
func (c *Client) DeleteIrQwebFieldContacts(ids []int64) error {
	return c.Delete(IrQwebFieldContactModel, ids)
}

// GetIrQwebFieldContact gets ir.qweb.field.contact existing record.
func (c *Client) GetIrQwebFieldContact(id int64) (*IrQwebFieldContact, error) {
	iqfcs, err := c.GetIrQwebFieldContacts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iqfcs)[0]), nil
}

// GetIrQwebFieldContacts gets ir.qweb.field.contact existing records.
func (c *Client) GetIrQwebFieldContacts(ids []int64) (*IrQwebFieldContacts, error) {
	iqfcs := &IrQwebFieldContacts{}
	if err := c.Read(IrQwebFieldContactModel, ids, nil, iqfcs); err != nil {
		return nil, err
	}
	return iqfcs, nil
}

// FindIrQwebFieldContact finds ir.qweb.field.contact record by querying it with criteria.
func (c *Client) FindIrQwebFieldContact(criteria *Criteria) (*IrQwebFieldContact, error) {
	iqfcs := &IrQwebFieldContacts{}
	if err := c.SearchRead(IrQwebFieldContactModel, criteria, NewOptions().Limit(1), iqfcs); err != nil {
		return nil, err
	}
	return &((*iqfcs)[0]), nil
}

// FindIrQwebFieldContacts finds ir.qweb.field.contact records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldContacts(criteria *Criteria, options *Options) (*IrQwebFieldContacts, error) {
	iqfcs := &IrQwebFieldContacts{}
	if err := c.SearchRead(IrQwebFieldContactModel, criteria, options, iqfcs); err != nil {
		return nil, err
	}
	return iqfcs, nil
}

// FindIrQwebFieldContactIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldContactIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrQwebFieldContactModel, criteria, options)
}

// FindIrQwebFieldContactId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldContactId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldContactModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
