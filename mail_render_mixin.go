package odoo

// MailRenderMixin represents mail.render.mixin model.
type MailRenderMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
	Lang        *String `xmlrpc:"lang,omitempty"`
	RenderModel *String `xmlrpc:"render_model,omitempty"`
}

// MailRenderMixins represents array of mail.render.mixin model.
type MailRenderMixins []MailRenderMixin

// MailRenderMixinModel is the odoo model name.
const MailRenderMixinModel = "mail.render.mixin"

// Many2One convert MailRenderMixin to *Many2One.
func (mrm *MailRenderMixin) Many2One() *Many2One {
	return NewMany2One(mrm.Id.Get(), "")
}

// CreateMailRenderMixin creates a new mail.render.mixin model and returns its id.
func (c *Client) CreateMailRenderMixin(mrm *MailRenderMixin) (int64, error) {
	ids, err := c.CreateMailRenderMixins([]*MailRenderMixin{mrm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailRenderMixin creates a new mail.render.mixin model and returns its id.
func (c *Client) CreateMailRenderMixins(mrms []*MailRenderMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range mrms {
		vv = append(vv, v)
	}
	return c.Create(MailRenderMixinModel, vv, nil)
}

// UpdateMailRenderMixin updates an existing mail.render.mixin record.
func (c *Client) UpdateMailRenderMixin(mrm *MailRenderMixin) error {
	return c.UpdateMailRenderMixins([]int64{mrm.Id.Get()}, mrm)
}

// UpdateMailRenderMixins updates existing mail.render.mixin records.
// All records (represented by ids) will be updated by mrm values.
func (c *Client) UpdateMailRenderMixins(ids []int64, mrm *MailRenderMixin) error {
	return c.Update(MailRenderMixinModel, ids, mrm, nil)
}

// DeleteMailRenderMixin deletes an existing mail.render.mixin record.
func (c *Client) DeleteMailRenderMixin(id int64) error {
	return c.DeleteMailRenderMixins([]int64{id})
}

// DeleteMailRenderMixins deletes existing mail.render.mixin records.
func (c *Client) DeleteMailRenderMixins(ids []int64) error {
	return c.Delete(MailRenderMixinModel, ids)
}

// GetMailRenderMixin gets mail.render.mixin existing record.
func (c *Client) GetMailRenderMixin(id int64) (*MailRenderMixin, error) {
	mrms, err := c.GetMailRenderMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mrms)[0]), nil
}

// GetMailRenderMixins gets mail.render.mixin existing records.
func (c *Client) GetMailRenderMixins(ids []int64) (*MailRenderMixins, error) {
	mrms := &MailRenderMixins{}
	if err := c.Read(MailRenderMixinModel, ids, nil, mrms); err != nil {
		return nil, err
	}
	return mrms, nil
}

// FindMailRenderMixin finds mail.render.mixin record by querying it with criteria.
func (c *Client) FindMailRenderMixin(criteria *Criteria) (*MailRenderMixin, error) {
	mrms := &MailRenderMixins{}
	if err := c.SearchRead(MailRenderMixinModel, criteria, NewOptions().Limit(1), mrms); err != nil {
		return nil, err
	}
	return &((*mrms)[0]), nil
}

// FindMailRenderMixins finds mail.render.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailRenderMixins(criteria *Criteria, options *Options) (*MailRenderMixins, error) {
	mrms := &MailRenderMixins{}
	if err := c.SearchRead(MailRenderMixinModel, criteria, options, mrms); err != nil {
		return nil, err
	}
	return mrms, nil
}

// FindMailRenderMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailRenderMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailRenderMixinModel, criteria, options)
}

// FindMailRenderMixinId finds record id by querying it with criteria.
func (c *Client) FindMailRenderMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailRenderMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
