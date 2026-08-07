package odoo

// HtmlEditorConverterTest represents html_editor.converter.test model.
type HtmlEditorConverterTest struct {
	Binary       *String    `xmlrpc:"binary,omitempty"`
	Char         *String    `xmlrpc:"char,omitempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date         *Time      `xmlrpc:"date,omitempty"`
	Datetime     *Time      `xmlrpc:"datetime,omitempty"`
	DisplayName  *String    `xmlrpc:"display_name,omitempty"`
	Float        *Float     `xmlrpc:"float,omitempty"`
	Html         *String    `xmlrpc:"html,omitempty"`
	Id           *Int       `xmlrpc:"id,omitempty"`
	Integer      *Int       `xmlrpc:"integer,omitempty"`
	Many2One     *Many2One  `xmlrpc:"many2one,omitempty"`
	Numeric      *Float     `xmlrpc:"numeric,omitempty"`
	SelectionStr *Selection `xmlrpc:"selection_str,omitempty"`
	Text         *String    `xmlrpc:"text,omitempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HtmlEditorConverterTests represents array of html_editor.converter.test model.
type HtmlEditorConverterTests []HtmlEditorConverterTest

// HtmlEditorConverterTestModel is the odoo model name.
const HtmlEditorConverterTestModel = "html_editor.converter.test"

// Many2One convert HtmlEditorConverterTest to *Many2One.
func (hct *HtmlEditorConverterTest) Many2One() *Many2One {
	return NewMany2One(hct.Id.Get(), "")
}

// CreateHtmlEditorConverterTest creates a new html_editor.converter.test model and returns its id.
func (c *Client) CreateHtmlEditorConverterTest(hct *HtmlEditorConverterTest) (int64, error) {
	ids, err := c.CreateHtmlEditorConverterTests([]*HtmlEditorConverterTest{hct})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHtmlEditorConverterTest creates a new html_editor.converter.test model and returns its id.
func (c *Client) CreateHtmlEditorConverterTests(hcts []*HtmlEditorConverterTest) ([]int64, error) {
	var vv []interface{}
	for _, v := range hcts {
		vv = append(vv, v)
	}
	return c.Create(HtmlEditorConverterTestModel, vv, nil)
}

// UpdateHtmlEditorConverterTest updates an existing html_editor.converter.test record.
func (c *Client) UpdateHtmlEditorConverterTest(hct *HtmlEditorConverterTest) error {
	return c.UpdateHtmlEditorConverterTests([]int64{hct.Id.Get()}, hct)
}

// UpdateHtmlEditorConverterTests updates existing html_editor.converter.test records.
// All records (represented by ids) will be updated by hct values.
func (c *Client) UpdateHtmlEditorConverterTests(ids []int64, hct *HtmlEditorConverterTest) error {
	return c.Update(HtmlEditorConverterTestModel, ids, hct, nil)
}

// DeleteHtmlEditorConverterTest deletes an existing html_editor.converter.test record.
func (c *Client) DeleteHtmlEditorConverterTest(id int64) error {
	return c.DeleteHtmlEditorConverterTests([]int64{id})
}

// DeleteHtmlEditorConverterTests deletes existing html_editor.converter.test records.
func (c *Client) DeleteHtmlEditorConverterTests(ids []int64) error {
	return c.Delete(HtmlEditorConverterTestModel, ids)
}

// GetHtmlEditorConverterTest gets html_editor.converter.test existing record.
func (c *Client) GetHtmlEditorConverterTest(id int64) (*HtmlEditorConverterTest, error) {
	hcts, err := c.GetHtmlEditorConverterTests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hcts)[0]), nil
}

// GetHtmlEditorConverterTests gets html_editor.converter.test existing records.
func (c *Client) GetHtmlEditorConverterTests(ids []int64) (*HtmlEditorConverterTests, error) {
	hcts := &HtmlEditorConverterTests{}
	if err := c.Read(HtmlEditorConverterTestModel, ids, nil, hcts); err != nil {
		return nil, err
	}
	return hcts, nil
}

// FindHtmlEditorConverterTest finds html_editor.converter.test record by querying it with criteria.
func (c *Client) FindHtmlEditorConverterTest(criteria *Criteria) (*HtmlEditorConverterTest, error) {
	hcts := &HtmlEditorConverterTests{}
	if err := c.SearchRead(HtmlEditorConverterTestModel, criteria, NewOptions().Limit(1), hcts); err != nil {
		return nil, err
	}
	return &((*hcts)[0]), nil
}

// FindHtmlEditorConverterTests finds html_editor.converter.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHtmlEditorConverterTests(criteria *Criteria, options *Options) (*HtmlEditorConverterTests, error) {
	hcts := &HtmlEditorConverterTests{}
	if err := c.SearchRead(HtmlEditorConverterTestModel, criteria, options, hcts); err != nil {
		return nil, err
	}
	return hcts, nil
}

// FindHtmlEditorConverterTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHtmlEditorConverterTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HtmlEditorConverterTestModel, criteria, options)
}

// FindHtmlEditorConverterTestId finds record id by querying it with criteria.
func (c *Client) FindHtmlEditorConverterTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HtmlEditorConverterTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
