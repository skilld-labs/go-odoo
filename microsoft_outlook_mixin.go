package odoo

// MicrosoftOutlookMixin represents microsoft.outlook.mixin model.
type MicrosoftOutlookMixin struct {
	Active                                *Bool   `xmlrpc:"active,omitempty"`
	DisplayName                           *String `xmlrpc:"display_name,omitempty"`
	Id                                    *Int    `xmlrpc:"id,omitempty"`
	MicrosoftOutlookAccessToken           *String `xmlrpc:"microsoft_outlook_access_token,omitempty"`
	MicrosoftOutlookAccessTokenExpiration *Int    `xmlrpc:"microsoft_outlook_access_token_expiration,omitempty"`
	MicrosoftOutlookRefreshToken          *String `xmlrpc:"microsoft_outlook_refresh_token,omitempty"`
	MicrosoftOutlookUri                   *String `xmlrpc:"microsoft_outlook_uri,omitempty"`
}

// MicrosoftOutlookMixins represents array of microsoft.outlook.mixin model.
type MicrosoftOutlookMixins []MicrosoftOutlookMixin

// MicrosoftOutlookMixinModel is the odoo model name.
const MicrosoftOutlookMixinModel = "microsoft.outlook.mixin"

// Many2One convert MicrosoftOutlookMixin to *Many2One.
func (mom *MicrosoftOutlookMixin) Many2One() *Many2One {
	return NewMany2One(mom.Id.Get(), "")
}

// CreateMicrosoftOutlookMixin creates a new microsoft.outlook.mixin model and returns its id.
func (c *Client) CreateMicrosoftOutlookMixin(mom *MicrosoftOutlookMixin) (int64, error) {
	ids, err := c.CreateMicrosoftOutlookMixins([]*MicrosoftOutlookMixin{mom})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMicrosoftOutlookMixin creates a new microsoft.outlook.mixin model and returns its id.
func (c *Client) CreateMicrosoftOutlookMixins(moms []*MicrosoftOutlookMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range moms {
		vv = append(vv, v)
	}
	return c.Create(MicrosoftOutlookMixinModel, vv, nil)
}

// UpdateMicrosoftOutlookMixin updates an existing microsoft.outlook.mixin record.
func (c *Client) UpdateMicrosoftOutlookMixin(mom *MicrosoftOutlookMixin) error {
	return c.UpdateMicrosoftOutlookMixins([]int64{mom.Id.Get()}, mom)
}

// UpdateMicrosoftOutlookMixins updates existing microsoft.outlook.mixin records.
// All records (represented by ids) will be updated by mom values.
func (c *Client) UpdateMicrosoftOutlookMixins(ids []int64, mom *MicrosoftOutlookMixin) error {
	return c.Update(MicrosoftOutlookMixinModel, ids, mom, nil)
}

// DeleteMicrosoftOutlookMixin deletes an existing microsoft.outlook.mixin record.
func (c *Client) DeleteMicrosoftOutlookMixin(id int64) error {
	return c.DeleteMicrosoftOutlookMixins([]int64{id})
}

// DeleteMicrosoftOutlookMixins deletes existing microsoft.outlook.mixin records.
func (c *Client) DeleteMicrosoftOutlookMixins(ids []int64) error {
	return c.Delete(MicrosoftOutlookMixinModel, ids)
}

// GetMicrosoftOutlookMixin gets microsoft.outlook.mixin existing record.
func (c *Client) GetMicrosoftOutlookMixin(id int64) (*MicrosoftOutlookMixin, error) {
	moms, err := c.GetMicrosoftOutlookMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*moms)[0]), nil
}

// GetMicrosoftOutlookMixins gets microsoft.outlook.mixin existing records.
func (c *Client) GetMicrosoftOutlookMixins(ids []int64) (*MicrosoftOutlookMixins, error) {
	moms := &MicrosoftOutlookMixins{}
	if err := c.Read(MicrosoftOutlookMixinModel, ids, nil, moms); err != nil {
		return nil, err
	}
	return moms, nil
}

// FindMicrosoftOutlookMixin finds microsoft.outlook.mixin record by querying it with criteria.
func (c *Client) FindMicrosoftOutlookMixin(criteria *Criteria) (*MicrosoftOutlookMixin, error) {
	moms := &MicrosoftOutlookMixins{}
	if err := c.SearchRead(MicrosoftOutlookMixinModel, criteria, NewOptions().Limit(1), moms); err != nil {
		return nil, err
	}
	return &((*moms)[0]), nil
}

// FindMicrosoftOutlookMixins finds microsoft.outlook.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMicrosoftOutlookMixins(criteria *Criteria, options *Options) (*MicrosoftOutlookMixins, error) {
	moms := &MicrosoftOutlookMixins{}
	if err := c.SearchRead(MicrosoftOutlookMixinModel, criteria, options, moms); err != nil {
		return nil, err
	}
	return moms, nil
}

// FindMicrosoftOutlookMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMicrosoftOutlookMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MicrosoftOutlookMixinModel, criteria, options)
}

// FindMicrosoftOutlookMixinId finds record id by querying it with criteria.
func (c *Client) FindMicrosoftOutlookMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MicrosoftOutlookMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
