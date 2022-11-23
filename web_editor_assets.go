package odoo

import (
	"fmt"
)

// WebEditorAssets represents web_editor.assets model.
type WebEditorAssets struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// WebEditorAssetss represents array of web_editor.assets model.
type WebEditorAssetss []WebEditorAssets

// WebEditorAssetsModel is the odoo model name.
const WebEditorAssetsModel = "web_editor.assets"

// Many2One convert WebEditorAssets to *Many2One.
func (wa *WebEditorAssets) Many2One() *Many2One {
	return NewMany2One(wa.Id.Get(), "")
}

// CreateWebEditorAssets creates a new web_editor.assets model and returns its id.
func (c *Client) CreateWebEditorAssets(wa *WebEditorAssets) (int64, error) {
	return c.Create(WebEditorAssetsModel, wa)
}

// UpdateWebEditorAssets updates an existing web_editor.assets record.
func (c *Client) UpdateWebEditorAssets(wa *WebEditorAssets) error {
	return c.UpdateWebEditorAssetss([]int64{wa.Id.Get()}, wa)
}

// UpdateWebEditorAssetss updates existing web_editor.assets records.
// All records (represented by IDs) will be updated by wa values.
func (c *Client) UpdateWebEditorAssetss(ids []int64, wa *WebEditorAssets) error {
	return c.Update(WebEditorAssetsModel, ids, wa)
}

// DeleteWebEditorAssets deletes an existing web_editor.assets record.
func (c *Client) DeleteWebEditorAssets(id int64) error {
	return c.DeleteWebEditorAssetss([]int64{id})
}

// DeleteWebEditorAssetss deletes existing web_editor.assets records.
func (c *Client) DeleteWebEditorAssetss(ids []int64) error {
	return c.Delete(WebEditorAssetsModel, ids)
}

// GetWebEditorAssets gets web_editor.assets existing record.
func (c *Client) GetWebEditorAssets(id int64) (*WebEditorAssets, error) {
	was, err := c.GetWebEditorAssetss([]int64{id})
	if err != nil {
		return nil, err
	}
	if was != nil && len(*was) > 0 {
		return &((*was)[0]), nil
	}
	return nil, fmt.Errorf("id %v of web_editor.assets not found", id)
}

// GetWebEditorAssetss gets web_editor.assets existing records.
func (c *Client) GetWebEditorAssetss(ids []int64) (*WebEditorAssetss, error) {
	was := &WebEditorAssetss{}
	if err := c.Read(WebEditorAssetsModel, ids, nil, was); err != nil {
		return nil, err
	}
	return was, nil
}

// FindWebEditorAssets finds web_editor.assets record by querying it with criteria.
func (c *Client) FindWebEditorAssets(criteria *Criteria) (*WebEditorAssets, error) {
	was := &WebEditorAssetss{}
	if err := c.SearchRead(WebEditorAssetsModel, criteria, NewOptions().Limit(1), was); err != nil {
		return nil, err
	}
	if was != nil && len(*was) > 0 {
		return &((*was)[0]), nil
	}
	return nil, fmt.Errorf("web_editor.assets was not found")
}

// FindWebEditorAssetss finds web_editor.assets records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebEditorAssetss(criteria *Criteria, options *Options) (*WebEditorAssetss, error) {
	was := &WebEditorAssetss{}
	if err := c.SearchRead(WebEditorAssetsModel, criteria, options, was); err != nil {
		return nil, err
	}
	return was, nil
}

// FindWebEditorAssetsIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebEditorAssetsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebEditorAssetsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebEditorAssetsId finds record id by querying it with criteria.
func (c *Client) FindWebEditorAssetsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebEditorAssetsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("web_editor.assets was not found")
}
