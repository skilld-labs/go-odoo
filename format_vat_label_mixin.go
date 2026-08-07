package odoo

// FormatVatLabelMixin represents format.vat.label.mixin model.
type FormatVatLabelMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// FormatVatLabelMixins represents array of format.vat.label.mixin model.
type FormatVatLabelMixins []FormatVatLabelMixin

// FormatVatLabelMixinModel is the odoo model name.
const FormatVatLabelMixinModel = "format.vat.label.mixin"

// Many2One convert FormatVatLabelMixin to *Many2One.
func (fvlm *FormatVatLabelMixin) Many2One() *Many2One {
	return NewMany2One(fvlm.Id.Get(), "")
}

// CreateFormatVatLabelMixin creates a new format.vat.label.mixin model and returns its id.
func (c *Client) CreateFormatVatLabelMixin(fvlm *FormatVatLabelMixin) (int64, error) {
	ids, err := c.CreateFormatVatLabelMixins([]*FormatVatLabelMixin{fvlm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFormatVatLabelMixin creates a new format.vat.label.mixin model and returns its id.
func (c *Client) CreateFormatVatLabelMixins(fvlms []*FormatVatLabelMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range fvlms {
		vv = append(vv, v)
	}
	return c.Create(FormatVatLabelMixinModel, vv, nil)
}

// UpdateFormatVatLabelMixin updates an existing format.vat.label.mixin record.
func (c *Client) UpdateFormatVatLabelMixin(fvlm *FormatVatLabelMixin) error {
	return c.UpdateFormatVatLabelMixins([]int64{fvlm.Id.Get()}, fvlm)
}

// UpdateFormatVatLabelMixins updates existing format.vat.label.mixin records.
// All records (represented by ids) will be updated by fvlm values.
func (c *Client) UpdateFormatVatLabelMixins(ids []int64, fvlm *FormatVatLabelMixin) error {
	return c.Update(FormatVatLabelMixinModel, ids, fvlm, nil)
}

// DeleteFormatVatLabelMixin deletes an existing format.vat.label.mixin record.
func (c *Client) DeleteFormatVatLabelMixin(id int64) error {
	return c.DeleteFormatVatLabelMixins([]int64{id})
}

// DeleteFormatVatLabelMixins deletes existing format.vat.label.mixin records.
func (c *Client) DeleteFormatVatLabelMixins(ids []int64) error {
	return c.Delete(FormatVatLabelMixinModel, ids)
}

// GetFormatVatLabelMixin gets format.vat.label.mixin existing record.
func (c *Client) GetFormatVatLabelMixin(id int64) (*FormatVatLabelMixin, error) {
	fvlms, err := c.GetFormatVatLabelMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*fvlms)[0]), nil
}

// GetFormatVatLabelMixins gets format.vat.label.mixin existing records.
func (c *Client) GetFormatVatLabelMixins(ids []int64) (*FormatVatLabelMixins, error) {
	fvlms := &FormatVatLabelMixins{}
	if err := c.Read(FormatVatLabelMixinModel, ids, nil, fvlms); err != nil {
		return nil, err
	}
	return fvlms, nil
}

// FindFormatVatLabelMixin finds format.vat.label.mixin record by querying it with criteria.
func (c *Client) FindFormatVatLabelMixin(criteria *Criteria) (*FormatVatLabelMixin, error) {
	fvlms := &FormatVatLabelMixins{}
	if err := c.SearchRead(FormatVatLabelMixinModel, criteria, NewOptions().Limit(1), fvlms); err != nil {
		return nil, err
	}
	return &((*fvlms)[0]), nil
}

// FindFormatVatLabelMixins finds format.vat.label.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFormatVatLabelMixins(criteria *Criteria, options *Options) (*FormatVatLabelMixins, error) {
	fvlms := &FormatVatLabelMixins{}
	if err := c.SearchRead(FormatVatLabelMixinModel, criteria, options, fvlms); err != nil {
		return nil, err
	}
	return fvlms, nil
}

// FindFormatVatLabelMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFormatVatLabelMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(FormatVatLabelMixinModel, criteria, options)
}

// FindFormatVatLabelMixinId finds record id by querying it with criteria.
func (c *Client) FindFormatVatLabelMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FormatVatLabelMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
