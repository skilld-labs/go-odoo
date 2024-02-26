package odoo

// WebEditorConverterTest represents web_editor.converter.test model.
type WebEditorConverterTest struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omptempty"`
	Binary       *String    `xmlrpc:"binary,omptempty"`
	Char         *String    `xmlrpc:"char,omptempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omptempty"`
	Date         *Time      `xmlrpc:"date,omptempty"`
	Datetime     *Time      `xmlrpc:"datetime,omptempty"`
	DisplayName  *String    `xmlrpc:"display_name,omptempty"`
	Float        *Float     `xmlrpc:"float,omptempty"`
	Html         *String    `xmlrpc:"html,omptempty"`
	Id           *Int       `xmlrpc:"id,omptempty"`
	Integer      *Int       `xmlrpc:"integer,omptempty"`
	Many2One     *Many2One  `xmlrpc:"many2one,omptempty"`
	Numeric      *Float     `xmlrpc:"numeric,omptempty"`
	Selection    *Selection `xmlrpc:"selection,omptempty"`
	SelectionStr *Selection `xmlrpc:"selection_str,omptempty"`
	Text         *String    `xmlrpc:"text,omptempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// WebEditorConverterTests represents array of web_editor.converter.test model.
type WebEditorConverterTests []WebEditorConverterTest

// WebEditorConverterTestModel is the odoo model name.
const WebEditorConverterTestModel = "web_editor.converter.test"

// Many2One convert WebEditorConverterTest to *Many2One.
func (wct *WebEditorConverterTest) Many2One() *Many2One {
	return NewMany2One(wct.Id.Get(), "")
}

// CreateWebEditorConverterTest creates a new web_editor.converter.test model and returns its id.
func (c *Client) CreateWebEditorConverterTest(wct *WebEditorConverterTest) (int64, error) {
	ids, err := c.CreateWebEditorConverterTests([]*WebEditorConverterTest{wct})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebEditorConverterTest creates a new web_editor.converter.test model and returns its id.
func (c *Client) CreateWebEditorConverterTests(wcts []*WebEditorConverterTest) ([]int64, error) {
	var vv []interface{}
	for _, v := range wcts {
		vv = append(vv, v)
	}
	return c.Create(WebEditorConverterTestModel, vv, nil)
}

// UpdateWebEditorConverterTest updates an existing web_editor.converter.test record.
func (c *Client) UpdateWebEditorConverterTest(wct *WebEditorConverterTest) error {
	return c.UpdateWebEditorConverterTests([]int64{wct.Id.Get()}, wct)
}

// UpdateWebEditorConverterTests updates existing web_editor.converter.test records.
// All records (represented by ids) will be updated by wct values.
func (c *Client) UpdateWebEditorConverterTests(ids []int64, wct *WebEditorConverterTest) error {
	return c.Update(WebEditorConverterTestModel, ids, wct, nil)
}

// DeleteWebEditorConverterTest deletes an existing web_editor.converter.test record.
func (c *Client) DeleteWebEditorConverterTest(id int64) error {
	return c.DeleteWebEditorConverterTests([]int64{id})
}

// DeleteWebEditorConverterTests deletes existing web_editor.converter.test records.
func (c *Client) DeleteWebEditorConverterTests(ids []int64) error {
	return c.Delete(WebEditorConverterTestModel, ids)
}

// GetWebEditorConverterTest gets web_editor.converter.test existing record.
func (c *Client) GetWebEditorConverterTest(id int64) (*WebEditorConverterTest, error) {
	wcts, err := c.GetWebEditorConverterTests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wcts)[0]), nil
}

// GetWebEditorConverterTests gets web_editor.converter.test existing records.
func (c *Client) GetWebEditorConverterTests(ids []int64) (*WebEditorConverterTests, error) {
	wcts := &WebEditorConverterTests{}
	if err := c.Read(WebEditorConverterTestModel, ids, nil, wcts); err != nil {
		return nil, err
	}
	return wcts, nil
}

// FindWebEditorConverterTest finds web_editor.converter.test record by querying it with criteria.
func (c *Client) FindWebEditorConverterTest(criteria *Criteria) (*WebEditorConverterTest, error) {
	wcts := &WebEditorConverterTests{}
	if err := c.SearchRead(WebEditorConverterTestModel, criteria, NewOptions().Limit(1), wcts); err != nil {
		return nil, err
	}
	return &((*wcts)[0]), nil
}

// FindWebEditorConverterTests finds web_editor.converter.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebEditorConverterTests(criteria *Criteria, options *Options) (*WebEditorConverterTests, error) {
	wcts := &WebEditorConverterTests{}
	if err := c.SearchRead(WebEditorConverterTestModel, criteria, options, wcts); err != nil {
		return nil, err
	}
	return wcts, nil
}

// FindWebEditorConverterTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebEditorConverterTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebEditorConverterTestModel, criteria, options)
}

// FindWebEditorConverterTestId finds record id by querying it with criteria.
func (c *Client) FindWebEditorConverterTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebEditorConverterTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
