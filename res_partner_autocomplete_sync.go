package odoo

import (
	"fmt"
)

// ResPartnerAutocompleteSync represents res.partner.autocomplete.sync model.
type ResPartnerAutocompleteSync struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	PartnerId   *Many2One `xmlrpc:"partner_id,omitempty"`
	Synched     *Bool     `xmlrpc:"synched,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResPartnerAutocompleteSyncs represents array of res.partner.autocomplete.sync model.
type ResPartnerAutocompleteSyncs []ResPartnerAutocompleteSync

// ResPartnerAutocompleteSyncModel is the odoo model name.
const ResPartnerAutocompleteSyncModel = "res.partner.autocomplete.sync"

// Many2One convert ResPartnerAutocompleteSync to *Many2One.
func (rpas *ResPartnerAutocompleteSync) Many2One() *Many2One {
	return NewMany2One(rpas.Id.Get(), "")
}

// CreateResPartnerAutocompleteSync creates a new res.partner.autocomplete.sync model and returns its id.
func (c *Client) CreateResPartnerAutocompleteSync(rpas *ResPartnerAutocompleteSync) (int64, error) {
	return c.Create(ResPartnerAutocompleteSyncModel, rpas)
}

// UpdateResPartnerAutocompleteSync updates an existing res.partner.autocomplete.sync record.
func (c *Client) UpdateResPartnerAutocompleteSync(rpas *ResPartnerAutocompleteSync) error {
	return c.UpdateResPartnerAutocompleteSyncs([]int64{rpas.Id.Get()}, rpas)
}

// UpdateResPartnerAutocompleteSyncs updates existing res.partner.autocomplete.sync records.
// All records (represented by ids) will be updated by rpas values.
func (c *Client) UpdateResPartnerAutocompleteSyncs(ids []int64, rpas *ResPartnerAutocompleteSync) error {
	return c.Update(ResPartnerAutocompleteSyncModel, ids, rpas)
}

// DeleteResPartnerAutocompleteSync deletes an existing res.partner.autocomplete.sync record.
func (c *Client) DeleteResPartnerAutocompleteSync(id int64) error {
	return c.DeleteResPartnerAutocompleteSyncs([]int64{id})
}

// DeleteResPartnerAutocompleteSyncs deletes existing res.partner.autocomplete.sync records.
func (c *Client) DeleteResPartnerAutocompleteSyncs(ids []int64) error {
	return c.Delete(ResPartnerAutocompleteSyncModel, ids)
}

// GetResPartnerAutocompleteSync gets res.partner.autocomplete.sync existing record.
func (c *Client) GetResPartnerAutocompleteSync(id int64) (*ResPartnerAutocompleteSync, error) {
	rpass, err := c.GetResPartnerAutocompleteSyncs([]int64{id})
	if err != nil {
		return nil, err
	}
	if rpass != nil && len(*rpass) > 0 {
		return &((*rpass)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.partner.autocomplete.sync not found", id)
}

// GetResPartnerAutocompleteSyncs gets res.partner.autocomplete.sync existing records.
func (c *Client) GetResPartnerAutocompleteSyncs(ids []int64) (*ResPartnerAutocompleteSyncs, error) {
	rpass := &ResPartnerAutocompleteSyncs{}
	if err := c.Read(ResPartnerAutocompleteSyncModel, ids, nil, rpass); err != nil {
		return nil, err
	}
	return rpass, nil
}

// FindResPartnerAutocompleteSync finds res.partner.autocomplete.sync record by querying it with criteria.
func (c *Client) FindResPartnerAutocompleteSync(criteria *Criteria) (*ResPartnerAutocompleteSync, error) {
	rpass := &ResPartnerAutocompleteSyncs{}
	if err := c.SearchRead(ResPartnerAutocompleteSyncModel, criteria, NewOptions().Limit(1), rpass); err != nil {
		return nil, err
	}
	if rpass != nil && len(*rpass) > 0 {
		return &((*rpass)[0]), nil
	}
	return nil, fmt.Errorf("res.partner.autocomplete.sync was not found")
}

// FindResPartnerAutocompleteSyncs finds res.partner.autocomplete.sync records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerAutocompleteSyncs(criteria *Criteria, options *Options) (*ResPartnerAutocompleteSyncs, error) {
	rpass := &ResPartnerAutocompleteSyncs{}
	if err := c.SearchRead(ResPartnerAutocompleteSyncModel, criteria, options, rpass); err != nil {
		return nil, err
	}
	return rpass, nil
}

// FindResPartnerAutocompleteSyncIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerAutocompleteSyncIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResPartnerAutocompleteSyncModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResPartnerAutocompleteSyncId finds record id by querying it with criteria.
func (c *Client) FindResPartnerAutocompleteSyncId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerAutocompleteSyncModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.partner.autocomplete.sync was not found")
}
