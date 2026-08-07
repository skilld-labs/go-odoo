package odoo

// HtmlEditorConverterTestSub represents html_editor.converter.test.sub model.
type HtmlEditorConverterTestSub struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HtmlEditorConverterTestSubs represents array of html_editor.converter.test.sub model.
type HtmlEditorConverterTestSubs []HtmlEditorConverterTestSub

// HtmlEditorConverterTestSubModel is the odoo model name.
const HtmlEditorConverterTestSubModel = "html_editor.converter.test.sub"

// Many2One convert HtmlEditorConverterTestSub to *Many2One.
func (hcts *HtmlEditorConverterTestSub) Many2One() *Many2One {
	return NewMany2One(hcts.Id.Get(), "")
}

// CreateHtmlEditorConverterTestSub creates a new html_editor.converter.test.sub model and returns its id.
func (c *Client) CreateHtmlEditorConverterTestSub(hcts *HtmlEditorConverterTestSub) (int64, error) {
	ids, err := c.CreateHtmlEditorConverterTestSubs([]*HtmlEditorConverterTestSub{hcts})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHtmlEditorConverterTestSub creates a new html_editor.converter.test.sub model and returns its id.
func (c *Client) CreateHtmlEditorConverterTestSubs(hctss []*HtmlEditorConverterTestSub) ([]int64, error) {
	var vv []interface{}
	for _, v := range hctss {
		vv = append(vv, v)
	}
	return c.Create(HtmlEditorConverterTestSubModel, vv, nil)
}

// UpdateHtmlEditorConverterTestSub updates an existing html_editor.converter.test.sub record.
func (c *Client) UpdateHtmlEditorConverterTestSub(hcts *HtmlEditorConverterTestSub) error {
	return c.UpdateHtmlEditorConverterTestSubs([]int64{hcts.Id.Get()}, hcts)
}

// UpdateHtmlEditorConverterTestSubs updates existing html_editor.converter.test.sub records.
// All records (represented by ids) will be updated by hcts values.
func (c *Client) UpdateHtmlEditorConverterTestSubs(ids []int64, hcts *HtmlEditorConverterTestSub) error {
	return c.Update(HtmlEditorConverterTestSubModel, ids, hcts, nil)
}

// DeleteHtmlEditorConverterTestSub deletes an existing html_editor.converter.test.sub record.
func (c *Client) DeleteHtmlEditorConverterTestSub(id int64) error {
	return c.DeleteHtmlEditorConverterTestSubs([]int64{id})
}

// DeleteHtmlEditorConverterTestSubs deletes existing html_editor.converter.test.sub records.
func (c *Client) DeleteHtmlEditorConverterTestSubs(ids []int64) error {
	return c.Delete(HtmlEditorConverterTestSubModel, ids)
}

// GetHtmlEditorConverterTestSub gets html_editor.converter.test.sub existing record.
func (c *Client) GetHtmlEditorConverterTestSub(id int64) (*HtmlEditorConverterTestSub, error) {
	hctss, err := c.GetHtmlEditorConverterTestSubs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hctss)[0]), nil
}

// GetHtmlEditorConverterTestSubs gets html_editor.converter.test.sub existing records.
func (c *Client) GetHtmlEditorConverterTestSubs(ids []int64) (*HtmlEditorConverterTestSubs, error) {
	hctss := &HtmlEditorConverterTestSubs{}
	if err := c.Read(HtmlEditorConverterTestSubModel, ids, nil, hctss); err != nil {
		return nil, err
	}
	return hctss, nil
}

// FindHtmlEditorConverterTestSub finds html_editor.converter.test.sub record by querying it with criteria.
func (c *Client) FindHtmlEditorConverterTestSub(criteria *Criteria) (*HtmlEditorConverterTestSub, error) {
	hctss := &HtmlEditorConverterTestSubs{}
	if err := c.SearchRead(HtmlEditorConverterTestSubModel, criteria, NewOptions().Limit(1), hctss); err != nil {
		return nil, err
	}
	return &((*hctss)[0]), nil
}

// FindHtmlEditorConverterTestSubs finds html_editor.converter.test.sub records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHtmlEditorConverterTestSubs(criteria *Criteria, options *Options) (*HtmlEditorConverterTestSubs, error) {
	hctss := &HtmlEditorConverterTestSubs{}
	if err := c.SearchRead(HtmlEditorConverterTestSubModel, criteria, options, hctss); err != nil {
		return nil, err
	}
	return hctss, nil
}

// FindHtmlEditorConverterTestSubIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHtmlEditorConverterTestSubIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HtmlEditorConverterTestSubModel, criteria, options)
}

// FindHtmlEditorConverterTestSubId finds record id by querying it with criteria.
func (c *Client) FindHtmlEditorConverterTestSubId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HtmlEditorConverterTestSubModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
