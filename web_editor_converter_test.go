package odoo

import (
	"fmt"
)

// WebEditorConverterTest represents web_editor.converter.test model.
type WebEditorConverterTest struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omitempty"`
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
	return c.Create(WebEditorConverterTestModel, wct)
}

// UpdateWebEditorConverterTest updates an existing web_editor.converter.test record.
func (c *Client) UpdateWebEditorConverterTest(wct *WebEditorConverterTest) error {
	return c.UpdateWebEditorConverterTests([]int64{wct.Id.Get()}, wct)
}

// UpdateWebEditorConverterTests updates existing web_editor.converter.test records.
// All records (represented by ids) will be updated by wct values.
func (c *Client) UpdateWebEditorConverterTests(ids []int64, wct *WebEditorConverterTest) error {
	return c.Update(WebEditorConverterTestModel, ids, wct)
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
	if wcts != nil && len(*wcts) > 0 {
		return &((*wcts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of web_editor.converter.test not found", id)
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
	if wcts != nil && len(*wcts) > 0 {
		return &((*wcts)[0]), nil
	}
	return nil, fmt.Errorf("web_editor.converter.test was not found")
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
	ids, err := c.Search(WebEditorConverterTestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebEditorConverterTestId finds record id by querying it with criteria.
func (c *Client) FindWebEditorConverterTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebEditorConverterTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("web_editor.converter.test was not found")
}
