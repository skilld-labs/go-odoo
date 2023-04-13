package odoo

import (
	"fmt"
)

// WebEditorConverterTestSub represents web_editor.converter.test.sub model.
type WebEditorConverterTestSub struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// WebEditorConverterTestSubs represents array of web_editor.converter.test.sub model.
type WebEditorConverterTestSubs []WebEditorConverterTestSub

// WebEditorConverterTestSubModel is the odoo model name.
const WebEditorConverterTestSubModel = "web_editor.converter.test.sub"

// Many2One convert WebEditorConverterTestSub to *Many2One.
func (wcts *WebEditorConverterTestSub) Many2One() *Many2One {
	return NewMany2One(wcts.Id.Get(), "")
}

// CreateWebEditorConverterTestSub creates a new web_editor.converter.test.sub model and returns its id.
func (c *Client) CreateWebEditorConverterTestSub(wcts *WebEditorConverterTestSub) (int64, error) {
	ids, err := c.Create(WebEditorConverterTestSubModel, []interface{}{wcts})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebEditorConverterTestSub creates a new web_editor.converter.test.sub model and returns its id.
func (c *Client) CreateWebEditorConverterTestSubs(wctss []*WebEditorConverterTestSub) ([]int64, error) {
	var vv []interface{}
	for _, v := range wctss {
		vv = append(vv, v)
	}
	return c.Create(WebEditorConverterTestSubModel, vv)
}

// UpdateWebEditorConverterTestSub updates an existing web_editor.converter.test.sub record.
func (c *Client) UpdateWebEditorConverterTestSub(wcts *WebEditorConverterTestSub) error {
	return c.UpdateWebEditorConverterTestSubs([]int64{wcts.Id.Get()}, wcts)
}

// UpdateWebEditorConverterTestSubs updates existing web_editor.converter.test.sub records.
// All records (represented by ids) will be updated by wcts values.
func (c *Client) UpdateWebEditorConverterTestSubs(ids []int64, wcts *WebEditorConverterTestSub) error {
	return c.Update(WebEditorConverterTestSubModel, ids, wcts)
}

// DeleteWebEditorConverterTestSub deletes an existing web_editor.converter.test.sub record.
func (c *Client) DeleteWebEditorConverterTestSub(id int64) error {
	return c.DeleteWebEditorConverterTestSubs([]int64{id})
}

// DeleteWebEditorConverterTestSubs deletes existing web_editor.converter.test.sub records.
func (c *Client) DeleteWebEditorConverterTestSubs(ids []int64) error {
	return c.Delete(WebEditorConverterTestSubModel, ids)
}

// GetWebEditorConverterTestSub gets web_editor.converter.test.sub existing record.
func (c *Client) GetWebEditorConverterTestSub(id int64) (*WebEditorConverterTestSub, error) {
	wctss, err := c.GetWebEditorConverterTestSubs([]int64{id})
	if err != nil {
		return nil, err
	}
	if wctss != nil && len(*wctss) > 0 {
		return &((*wctss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of web_editor.converter.test.sub not found", id)
}

// GetWebEditorConverterTestSubs gets web_editor.converter.test.sub existing records.
func (c *Client) GetWebEditorConverterTestSubs(ids []int64) (*WebEditorConverterTestSubs, error) {
	wctss := &WebEditorConverterTestSubs{}
	if err := c.Read(WebEditorConverterTestSubModel, ids, nil, wctss); err != nil {
		return nil, err
	}
	return wctss, nil
}

// FindWebEditorConverterTestSub finds web_editor.converter.test.sub record by querying it with criteria.
func (c *Client) FindWebEditorConverterTestSub(criteria *Criteria) (*WebEditorConverterTestSub, error) {
	wctss := &WebEditorConverterTestSubs{}
	if err := c.SearchRead(WebEditorConverterTestSubModel, criteria, NewOptions().Limit(1), wctss); err != nil {
		return nil, err
	}
	if wctss != nil && len(*wctss) > 0 {
		return &((*wctss)[0]), nil
	}
	return nil, fmt.Errorf("web_editor.converter.test.sub was not found with criteria %v", criteria)
}

// FindWebEditorConverterTestSubs finds web_editor.converter.test.sub records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebEditorConverterTestSubs(criteria *Criteria, options *Options) (*WebEditorConverterTestSubs, error) {
	wctss := &WebEditorConverterTestSubs{}
	if err := c.SearchRead(WebEditorConverterTestSubModel, criteria, options, wctss); err != nil {
		return nil, err
	}
	return wctss, nil
}

// FindWebEditorConverterTestSubIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebEditorConverterTestSubIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebEditorConverterTestSubModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebEditorConverterTestSubId finds record id by querying it with criteria.
func (c *Client) FindWebEditorConverterTestSubId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebEditorConverterTestSubModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("web_editor.converter.test.sub was not found with criteria %v and options %v", criteria, options)
}
