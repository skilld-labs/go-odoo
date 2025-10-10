package odoo

// IapService represents iap.service model.
type IapService struct {
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	Description    *String   `xmlrpc:"description,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	IntegerBalance *Bool     `xmlrpc:"integer_balance,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	TechnicalName  *String   `xmlrpc:"technical_name,omitempty"`
	UnitName       *String   `xmlrpc:"unit_name,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IapServices represents array of iap.service model.
type IapServices []IapService

// IapServiceModel is the odoo model name.
const IapServiceModel = "iap.service"

// Many2One convert IapService to *Many2One.
func (is *IapService) Many2One() *Many2One {
	return NewMany2One(is.Id.Get(), "")
}

// CreateIapService creates a new iap.service model and returns its id.
func (c *Client) CreateIapService(is *IapService) (int64, error) {
	ids, err := c.CreateIapServices([]*IapService{is})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIapService creates a new iap.service model and returns its id.
func (c *Client) CreateIapServices(iss []*IapService) ([]int64, error) {
	var vv []interface{}
	for _, v := range iss {
		vv = append(vv, v)
	}
	return c.Create(IapServiceModel, vv, nil)
}

// UpdateIapService updates an existing iap.service record.
func (c *Client) UpdateIapService(is *IapService) error {
	return c.UpdateIapServices([]int64{is.Id.Get()}, is)
}

// UpdateIapServices updates existing iap.service records.
// All records (represented by ids) will be updated by is values.
func (c *Client) UpdateIapServices(ids []int64, is *IapService) error {
	return c.Update(IapServiceModel, ids, is, nil)
}

// DeleteIapService deletes an existing iap.service record.
func (c *Client) DeleteIapService(id int64) error {
	return c.DeleteIapServices([]int64{id})
}

// DeleteIapServices deletes existing iap.service records.
func (c *Client) DeleteIapServices(ids []int64) error {
	return c.Delete(IapServiceModel, ids)
}

// GetIapService gets iap.service existing record.
func (c *Client) GetIapService(id int64) (*IapService, error) {
	iss, err := c.GetIapServices([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iss)[0]), nil
}

// GetIapServices gets iap.service existing records.
func (c *Client) GetIapServices(ids []int64) (*IapServices, error) {
	iss := &IapServices{}
	if err := c.Read(IapServiceModel, ids, nil, iss); err != nil {
		return nil, err
	}
	return iss, nil
}

// FindIapService finds iap.service record by querying it with criteria.
func (c *Client) FindIapService(criteria *Criteria) (*IapService, error) {
	iss := &IapServices{}
	if err := c.SearchRead(IapServiceModel, criteria, NewOptions().Limit(1), iss); err != nil {
		return nil, err
	}
	return &((*iss)[0]), nil
}

// FindIapServices finds iap.service records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIapServices(criteria *Criteria, options *Options) (*IapServices, error) {
	iss := &IapServices{}
	if err := c.SearchRead(IapServiceModel, criteria, options, iss); err != nil {
		return nil, err
	}
	return iss, nil
}

// FindIapServiceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIapServiceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IapServiceModel, criteria, options)
}

// FindIapServiceId finds record id by querying it with criteria.
func (c *Client) FindIapServiceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IapServiceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
