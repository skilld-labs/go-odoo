package odoo

import (
	"fmt"
)

// DigestDigest represents digest.digest model.
type DigestDigest struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	AvailableFields             *String    `xmlrpc:"available_fields,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IsSubscribed                *Bool      `xmlrpc:"is_subscribed,omitempty"`
	KpiAccountTotalRevenue      *Bool      `xmlrpc:"kpi_account_total_revenue,omitempty"`
	KpiAccountTotalRevenueValue *Float     `xmlrpc:"kpi_account_total_revenue_value,omitempty"`
	KpiAllSaleTotal             *Bool      `xmlrpc:"kpi_all_sale_total,omitempty"`
	KpiAllSaleTotalValue        *Float     `xmlrpc:"kpi_all_sale_total_value,omitempty"`
	KpiMailMessageTotal         *Bool      `xmlrpc:"kpi_mail_message_total,omitempty"`
	KpiMailMessageTotalValue    *Int       `xmlrpc:"kpi_mail_message_total_value,omitempty"`
	KpiResUsersConnected        *Bool      `xmlrpc:"kpi_res_users_connected,omitempty"`
	KpiResUsersConnectedValue   *Int       `xmlrpc:"kpi_res_users_connected_value,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	NextRunDate                 *Time      `xmlrpc:"next_run_date,omitempty"`
	Periodicity                 *Selection `xmlrpc:"periodicity,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	UserIds                     *Relation  `xmlrpc:"user_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// DigestDigests represents array of digest.digest model.
type DigestDigests []DigestDigest

// DigestDigestModel is the odoo model name.
const DigestDigestModel = "digest.digest"

// Many2One convert DigestDigest to *Many2One.
func (dd *DigestDigest) Many2One() *Many2One {
	return NewMany2One(dd.Id.Get(), "")
}

// CreateDigestDigest creates a new digest.digest model and returns its id.
func (c *Client) CreateDigestDigest(dd *DigestDigest) (int64, error) {
	return c.Create(DigestDigestModel, dd)
}

// UpdateDigestDigest updates an existing digest.digest record.
func (c *Client) UpdateDigestDigest(dd *DigestDigest) error {
	return c.UpdateDigestDigests([]int64{dd.Id.Get()}, dd)
}

// UpdateDigestDigests updates existing digest.digest records.
// All records (represented by IDs) will be updated by dd values.
func (c *Client) UpdateDigestDigests(ids []int64, dd *DigestDigest) error {
	return c.Update(DigestDigestModel, ids, dd)
}

// DeleteDigestDigest deletes an existing digest.digest record.
func (c *Client) DeleteDigestDigest(id int64) error {
	return c.DeleteDigestDigests([]int64{id})
}

// DeleteDigestDigests deletes existing digest.digest records.
func (c *Client) DeleteDigestDigests(ids []int64) error {
	return c.Delete(DigestDigestModel, ids)
}

// GetDigestDigest gets digest.digest existing record.
func (c *Client) GetDigestDigest(id int64) (*DigestDigest, error) {
	dds, err := c.GetDigestDigests([]int64{id})
	if err != nil {
		return nil, err
	}
	if dds != nil && len(*dds) > 0 {
		return &((*dds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of digest.digest not found", id)
}

// GetDigestDigests gets digest.digest existing records.
func (c *Client) GetDigestDigests(ids []int64) (*DigestDigests, error) {
	dds := &DigestDigests{}
	if err := c.Read(DigestDigestModel, ids, nil, dds); err != nil {
		return nil, err
	}
	return dds, nil
}

// FindDigestDigest finds digest.digest record by querying it with criteria.
func (c *Client) FindDigestDigest(criteria *Criteria) (*DigestDigest, error) {
	dds := &DigestDigests{}
	if err := c.SearchRead(DigestDigestModel, criteria, NewOptions().Limit(1), dds); err != nil {
		return nil, err
	}
	if dds != nil && len(*dds) > 0 {
		return &((*dds)[0]), nil
	}
	return nil, fmt.Errorf("digest.digest was not found")
}

// FindDigestDigests finds digest.digest records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDigestDigests(criteria *Criteria, options *Options) (*DigestDigests, error) {
	dds := &DigestDigests{}
	if err := c.SearchRead(DigestDigestModel, criteria, options, dds); err != nil {
		return nil, err
	}
	return dds, nil
}

// FindDigestDigestIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindDigestDigestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DigestDigestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDigestDigestId finds record id by querying it with criteria.
func (c *Client) FindDigestDigestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DigestDigestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("digest.digest was not found")
}
