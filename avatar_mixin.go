package odoo

// AvatarMixin represents avatar.mixin model.
type AvatarMixin struct {
	Avatar1024  *String `xmlrpc:"avatar_1024,omitempty"`
	Avatar128   *String `xmlrpc:"avatar_128,omitempty"`
	Avatar1920  *String `xmlrpc:"avatar_1920,omitempty"`
	Avatar256   *String `xmlrpc:"avatar_256,omitempty"`
	Avatar512   *String `xmlrpc:"avatar_512,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
	Image1024   *String `xmlrpc:"image_1024,omitempty"`
	Image128    *String `xmlrpc:"image_128,omitempty"`
	Image1920   *String `xmlrpc:"image_1920,omitempty"`
	Image256    *String `xmlrpc:"image_256,omitempty"`
	Image512    *String `xmlrpc:"image_512,omitempty"`
}

// AvatarMixins represents array of avatar.mixin model.
type AvatarMixins []AvatarMixin

// AvatarMixinModel is the odoo model name.
const AvatarMixinModel = "avatar.mixin"

// Many2One convert AvatarMixin to *Many2One.
func (am *AvatarMixin) Many2One() *Many2One {
	return NewMany2One(am.Id.Get(), "")
}

// CreateAvatarMixin creates a new avatar.mixin model and returns its id.
func (c *Client) CreateAvatarMixin(am *AvatarMixin) (int64, error) {
	ids, err := c.CreateAvatarMixins([]*AvatarMixin{am})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAvatarMixin creates a new avatar.mixin model and returns its id.
func (c *Client) CreateAvatarMixins(ams []*AvatarMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range ams {
		vv = append(vv, v)
	}
	return c.Create(AvatarMixinModel, vv, nil)
}

// UpdateAvatarMixin updates an existing avatar.mixin record.
func (c *Client) UpdateAvatarMixin(am *AvatarMixin) error {
	return c.UpdateAvatarMixins([]int64{am.Id.Get()}, am)
}

// UpdateAvatarMixins updates existing avatar.mixin records.
// All records (represented by ids) will be updated by am values.
func (c *Client) UpdateAvatarMixins(ids []int64, am *AvatarMixin) error {
	return c.Update(AvatarMixinModel, ids, am, nil)
}

// DeleteAvatarMixin deletes an existing avatar.mixin record.
func (c *Client) DeleteAvatarMixin(id int64) error {
	return c.DeleteAvatarMixins([]int64{id})
}

// DeleteAvatarMixins deletes existing avatar.mixin records.
func (c *Client) DeleteAvatarMixins(ids []int64) error {
	return c.Delete(AvatarMixinModel, ids)
}

// GetAvatarMixin gets avatar.mixin existing record.
func (c *Client) GetAvatarMixin(id int64) (*AvatarMixin, error) {
	ams, err := c.GetAvatarMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ams)[0]), nil
}

// GetAvatarMixins gets avatar.mixin existing records.
func (c *Client) GetAvatarMixins(ids []int64) (*AvatarMixins, error) {
	ams := &AvatarMixins{}
	if err := c.Read(AvatarMixinModel, ids, nil, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAvatarMixin finds avatar.mixin record by querying it with criteria.
func (c *Client) FindAvatarMixin(criteria *Criteria) (*AvatarMixin, error) {
	ams := &AvatarMixins{}
	if err := c.SearchRead(AvatarMixinModel, criteria, NewOptions().Limit(1), ams); err != nil {
		return nil, err
	}
	return &((*ams)[0]), nil
}

// FindAvatarMixins finds avatar.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAvatarMixins(criteria *Criteria, options *Options) (*AvatarMixins, error) {
	ams := &AvatarMixins{}
	if err := c.SearchRead(AvatarMixinModel, criteria, options, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAvatarMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAvatarMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AvatarMixinModel, criteria, options)
}

// FindAvatarMixinId finds record id by querying it with criteria.
func (c *Client) FindAvatarMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AvatarMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
