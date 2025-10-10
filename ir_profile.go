package odoo

// IrProfile represents ir.profile model.
type IrProfile struct {
	CreateDate     *Time   `xmlrpc:"create_date,omitempty"`
	DisplayName    *String `xmlrpc:"display_name,omitempty"`
	Duration       *Float  `xmlrpc:"duration,omitempty"`
	EntryCount     *Int    `xmlrpc:"entry_count,omitempty"`
	Id             *Int    `xmlrpc:"id,omitempty"`
	InitStackTrace *String `xmlrpc:"init_stack_trace,omitempty"`
	Name           *String `xmlrpc:"name,omitempty"`
	Qweb           *String `xmlrpc:"qweb,omitempty"`
	Session        *String `xmlrpc:"session,omitempty"`
	Speedscope     *String `xmlrpc:"speedscope,omitempty"`
	SpeedscopeUrl  *String `xmlrpc:"speedscope_url,omitempty"`
	Sql            *String `xmlrpc:"sql,omitempty"`
	SqlCount       *Int    `xmlrpc:"sql_count,omitempty"`
	TracesAsync    *String `xmlrpc:"traces_async,omitempty"`
	TracesSync     *String `xmlrpc:"traces_sync,omitempty"`
}

// IrProfiles represents array of ir.profile model.
type IrProfiles []IrProfile

// IrProfileModel is the odoo model name.
const IrProfileModel = "ir.profile"

// Many2One convert IrProfile to *Many2One.
func (ip *IrProfile) Many2One() *Many2One {
	return NewMany2One(ip.Id.Get(), "")
}

// CreateIrProfile creates a new ir.profile model and returns its id.
func (c *Client) CreateIrProfile(ip *IrProfile) (int64, error) {
	ids, err := c.CreateIrProfiles([]*IrProfile{ip})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrProfile creates a new ir.profile model and returns its id.
func (c *Client) CreateIrProfiles(ips []*IrProfile) ([]int64, error) {
	var vv []interface{}
	for _, v := range ips {
		vv = append(vv, v)
	}
	return c.Create(IrProfileModel, vv, nil)
}

// UpdateIrProfile updates an existing ir.profile record.
func (c *Client) UpdateIrProfile(ip *IrProfile) error {
	return c.UpdateIrProfiles([]int64{ip.Id.Get()}, ip)
}

// UpdateIrProfiles updates existing ir.profile records.
// All records (represented by ids) will be updated by ip values.
func (c *Client) UpdateIrProfiles(ids []int64, ip *IrProfile) error {
	return c.Update(IrProfileModel, ids, ip, nil)
}

// DeleteIrProfile deletes an existing ir.profile record.
func (c *Client) DeleteIrProfile(id int64) error {
	return c.DeleteIrProfiles([]int64{id})
}

// DeleteIrProfiles deletes existing ir.profile records.
func (c *Client) DeleteIrProfiles(ids []int64) error {
	return c.Delete(IrProfileModel, ids)
}

// GetIrProfile gets ir.profile existing record.
func (c *Client) GetIrProfile(id int64) (*IrProfile, error) {
	ips, err := c.GetIrProfiles([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ips)[0]), nil
}

// GetIrProfiles gets ir.profile existing records.
func (c *Client) GetIrProfiles(ids []int64) (*IrProfiles, error) {
	ips := &IrProfiles{}
	if err := c.Read(IrProfileModel, ids, nil, ips); err != nil {
		return nil, err
	}
	return ips, nil
}

// FindIrProfile finds ir.profile record by querying it with criteria.
func (c *Client) FindIrProfile(criteria *Criteria) (*IrProfile, error) {
	ips := &IrProfiles{}
	if err := c.SearchRead(IrProfileModel, criteria, NewOptions().Limit(1), ips); err != nil {
		return nil, err
	}
	return &((*ips)[0]), nil
}

// FindIrProfiles finds ir.profile records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrProfiles(criteria *Criteria, options *Options) (*IrProfiles, error) {
	ips := &IrProfiles{}
	if err := c.SearchRead(IrProfileModel, criteria, options, ips); err != nil {
		return nil, err
	}
	return ips, nil
}

// FindIrProfileIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrProfileIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrProfileModel, criteria, options)
}

// FindIrProfileId finds record id by querying it with criteria.
func (c *Client) FindIrProfileId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrProfileModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
