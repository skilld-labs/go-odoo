package odoo

// MailComposerMixin represents mail.composer.mixin model.
type MailComposerMixin struct {
	Body                 *String   `xmlrpc:"body,omitempty"`
	BodyHasTemplateValue *Bool     `xmlrpc:"body_has_template_value,omitempty"`
	CanEditBody          *Bool     `xmlrpc:"can_edit_body,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	IsMailTemplateEditor *Bool     `xmlrpc:"is_mail_template_editor,omitempty"`
	Lang                 *String   `xmlrpc:"lang,omitempty"`
	RenderModel          *String   `xmlrpc:"render_model,omitempty"`
	Subject              *String   `xmlrpc:"subject,omitempty"`
	TemplateId           *Many2One `xmlrpc:"template_id,omitempty"`
}

// MailComposerMixins represents array of mail.composer.mixin model.
type MailComposerMixins []MailComposerMixin

// MailComposerMixinModel is the odoo model name.
const MailComposerMixinModel = "mail.composer.mixin"

// Many2One convert MailComposerMixin to *Many2One.
func (mcm *MailComposerMixin) Many2One() *Many2One {
	return NewMany2One(mcm.Id.Get(), "")
}

// CreateMailComposerMixin creates a new mail.composer.mixin model and returns its id.
func (c *Client) CreateMailComposerMixin(mcm *MailComposerMixin) (int64, error) {
	ids, err := c.CreateMailComposerMixins([]*MailComposerMixin{mcm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailComposerMixin creates a new mail.composer.mixin model and returns its id.
func (c *Client) CreateMailComposerMixins(mcms []*MailComposerMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcms {
		vv = append(vv, v)
	}
	return c.Create(MailComposerMixinModel, vv, nil)
}

// UpdateMailComposerMixin updates an existing mail.composer.mixin record.
func (c *Client) UpdateMailComposerMixin(mcm *MailComposerMixin) error {
	return c.UpdateMailComposerMixins([]int64{mcm.Id.Get()}, mcm)
}

// UpdateMailComposerMixins updates existing mail.composer.mixin records.
// All records (represented by ids) will be updated by mcm values.
func (c *Client) UpdateMailComposerMixins(ids []int64, mcm *MailComposerMixin) error {
	return c.Update(MailComposerMixinModel, ids, mcm, nil)
}

// DeleteMailComposerMixin deletes an existing mail.composer.mixin record.
func (c *Client) DeleteMailComposerMixin(id int64) error {
	return c.DeleteMailComposerMixins([]int64{id})
}

// DeleteMailComposerMixins deletes existing mail.composer.mixin records.
func (c *Client) DeleteMailComposerMixins(ids []int64) error {
	return c.Delete(MailComposerMixinModel, ids)
}

// GetMailComposerMixin gets mail.composer.mixin existing record.
func (c *Client) GetMailComposerMixin(id int64) (*MailComposerMixin, error) {
	mcms, err := c.GetMailComposerMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcms)[0]), nil
}

// GetMailComposerMixins gets mail.composer.mixin existing records.
func (c *Client) GetMailComposerMixins(ids []int64) (*MailComposerMixins, error) {
	mcms := &MailComposerMixins{}
	if err := c.Read(MailComposerMixinModel, ids, nil, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailComposerMixin finds mail.composer.mixin record by querying it with criteria.
func (c *Client) FindMailComposerMixin(criteria *Criteria) (*MailComposerMixin, error) {
	mcms := &MailComposerMixins{}
	if err := c.SearchRead(MailComposerMixinModel, criteria, NewOptions().Limit(1), mcms); err != nil {
		return nil, err
	}
	return &((*mcms)[0]), nil
}

// FindMailComposerMixins finds mail.composer.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailComposerMixins(criteria *Criteria, options *Options) (*MailComposerMixins, error) {
	mcms := &MailComposerMixins{}
	if err := c.SearchRead(MailComposerMixinModel, criteria, options, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailComposerMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailComposerMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailComposerMixinModel, criteria, options)
}

// FindMailComposerMixinId finds record id by querying it with criteria.
func (c *Client) FindMailComposerMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailComposerMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
