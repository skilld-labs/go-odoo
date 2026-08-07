package odoo

// HtmlFieldHistoryMixin represents html.field.history.mixin model.
type HtmlFieldHistoryMixin struct {
	DisplayName              *String     `xmlrpc:"display_name,omitempty"`
	HtmlFieldHistory         interface{} `xmlrpc:"html_field_history,omitempty"`
	HtmlFieldHistoryMetadata interface{} `xmlrpc:"html_field_history_metadata,omitempty"`
	Id                       *Int        `xmlrpc:"id,omitempty"`
}

// HtmlFieldHistoryMixins represents array of html.field.history.mixin model.
type HtmlFieldHistoryMixins []HtmlFieldHistoryMixin

// HtmlFieldHistoryMixinModel is the odoo model name.
const HtmlFieldHistoryMixinModel = "html.field.history.mixin"

// Many2One convert HtmlFieldHistoryMixin to *Many2One.
func (hfhm *HtmlFieldHistoryMixin) Many2One() *Many2One {
	return NewMany2One(hfhm.Id.Get(), "")
}

// CreateHtmlFieldHistoryMixin creates a new html.field.history.mixin model and returns its id.
func (c *Client) CreateHtmlFieldHistoryMixin(hfhm *HtmlFieldHistoryMixin) (int64, error) {
	ids, err := c.CreateHtmlFieldHistoryMixins([]*HtmlFieldHistoryMixin{hfhm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHtmlFieldHistoryMixin creates a new html.field.history.mixin model and returns its id.
func (c *Client) CreateHtmlFieldHistoryMixins(hfhms []*HtmlFieldHistoryMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range hfhms {
		vv = append(vv, v)
	}
	return c.Create(HtmlFieldHistoryMixinModel, vv, nil)
}

// UpdateHtmlFieldHistoryMixin updates an existing html.field.history.mixin record.
func (c *Client) UpdateHtmlFieldHistoryMixin(hfhm *HtmlFieldHistoryMixin) error {
	return c.UpdateHtmlFieldHistoryMixins([]int64{hfhm.Id.Get()}, hfhm)
}

// UpdateHtmlFieldHistoryMixins updates existing html.field.history.mixin records.
// All records (represented by ids) will be updated by hfhm values.
func (c *Client) UpdateHtmlFieldHistoryMixins(ids []int64, hfhm *HtmlFieldHistoryMixin) error {
	return c.Update(HtmlFieldHistoryMixinModel, ids, hfhm, nil)
}

// DeleteHtmlFieldHistoryMixin deletes an existing html.field.history.mixin record.
func (c *Client) DeleteHtmlFieldHistoryMixin(id int64) error {
	return c.DeleteHtmlFieldHistoryMixins([]int64{id})
}

// DeleteHtmlFieldHistoryMixins deletes existing html.field.history.mixin records.
func (c *Client) DeleteHtmlFieldHistoryMixins(ids []int64) error {
	return c.Delete(HtmlFieldHistoryMixinModel, ids)
}

// GetHtmlFieldHistoryMixin gets html.field.history.mixin existing record.
func (c *Client) GetHtmlFieldHistoryMixin(id int64) (*HtmlFieldHistoryMixin, error) {
	hfhms, err := c.GetHtmlFieldHistoryMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hfhms)[0]), nil
}

// GetHtmlFieldHistoryMixins gets html.field.history.mixin existing records.
func (c *Client) GetHtmlFieldHistoryMixins(ids []int64) (*HtmlFieldHistoryMixins, error) {
	hfhms := &HtmlFieldHistoryMixins{}
	if err := c.Read(HtmlFieldHistoryMixinModel, ids, nil, hfhms); err != nil {
		return nil, err
	}
	return hfhms, nil
}

// FindHtmlFieldHistoryMixin finds html.field.history.mixin record by querying it with criteria.
func (c *Client) FindHtmlFieldHistoryMixin(criteria *Criteria) (*HtmlFieldHistoryMixin, error) {
	hfhms := &HtmlFieldHistoryMixins{}
	if err := c.SearchRead(HtmlFieldHistoryMixinModel, criteria, NewOptions().Limit(1), hfhms); err != nil {
		return nil, err
	}
	return &((*hfhms)[0]), nil
}

// FindHtmlFieldHistoryMixins finds html.field.history.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHtmlFieldHistoryMixins(criteria *Criteria, options *Options) (*HtmlFieldHistoryMixins, error) {
	hfhms := &HtmlFieldHistoryMixins{}
	if err := c.SearchRead(HtmlFieldHistoryMixinModel, criteria, options, hfhms); err != nil {
		return nil, err
	}
	return hfhms, nil
}

// FindHtmlFieldHistoryMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHtmlFieldHistoryMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HtmlFieldHistoryMixinModel, criteria, options)
}

// FindHtmlFieldHistoryMixinId finds record id by querying it with criteria.
func (c *Client) FindHtmlFieldHistoryMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HtmlFieldHistoryMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
