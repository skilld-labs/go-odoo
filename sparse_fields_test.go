package odoo

import (
	"fmt"
)

// SparseFieldsTest represents sparse_fields.test model.
type SparseFieldsTest struct {
	LastUpdate  *Time       `xmlrpc:"__last_update,omptempty"`
	Boolean     *Bool       `xmlrpc:"boolean,omptempty"`
	Char        *String     `xmlrpc:"char,omptempty"`
	CreateDate  *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One   `xmlrpc:"create_uid,omptempty"`
	Data        interface{} `xmlrpc:"data,omptempty"`
	DisplayName *String     `xmlrpc:"display_name,omptempty"`
	Float       *Float      `xmlrpc:"float,omptempty"`
	Id          *Int        `xmlrpc:"id,omptempty"`
	Integer     *Int        `xmlrpc:"integer,omptempty"`
	Partner     *Many2One   `xmlrpc:"partner,omptempty"`
	Selection   *Selection  `xmlrpc:"selection,omptempty"`
	WriteDate   *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// SparseFieldsTests represents array of sparse_fields.test model.
type SparseFieldsTests []SparseFieldsTest

// SparseFieldsTestModel is the odoo model name.
const SparseFieldsTestModel = "sparse_fields.test"

// Many2One convert SparseFieldsTest to *Many2One.
func (st *SparseFieldsTest) Many2One() *Many2One {
	return NewMany2One(st.Id.Get(), "")
}

// CreateSparseFieldsTest creates a new sparse_fields.test model and returns its id.
func (c *Client) CreateSparseFieldsTest(st *SparseFieldsTest) (int64, error) {
	return c.Create(SparseFieldsTestModel, st)
}

// UpdateSparseFieldsTest updates an existing sparse_fields.test record.
func (c *Client) UpdateSparseFieldsTest(st *SparseFieldsTest) error {
	return c.UpdateSparseFieldsTests([]int64{st.Id.Get()}, st)
}

// UpdateSparseFieldsTests updates existing sparse_fields.test records.
// All records (represented by ids) will be updated by st values.
func (c *Client) UpdateSparseFieldsTests(ids []int64, st *SparseFieldsTest) error {
	return c.Update(SparseFieldsTestModel, ids, st)
}

// DeleteSparseFieldsTest deletes an existing sparse_fields.test record.
func (c *Client) DeleteSparseFieldsTest(id int64) error {
	return c.DeleteSparseFieldsTests([]int64{id})
}

// DeleteSparseFieldsTests deletes existing sparse_fields.test records.
func (c *Client) DeleteSparseFieldsTests(ids []int64) error {
	return c.Delete(SparseFieldsTestModel, ids)
}

// GetSparseFieldsTest gets sparse_fields.test existing record.
func (c *Client) GetSparseFieldsTest(id int64) (*SparseFieldsTest, error) {
	sts, err := c.GetSparseFieldsTests([]int64{id})
	if err != nil {
		return nil, err
	}
	if sts != nil && len(*sts) > 0 {
		return &((*sts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sparse_fields.test not found", id)
}

// GetSparseFieldsTests gets sparse_fields.test existing records.
func (c *Client) GetSparseFieldsTests(ids []int64) (*SparseFieldsTests, error) {
	sts := &SparseFieldsTests{}
	if err := c.Read(SparseFieldsTestModel, ids, nil, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSparseFieldsTest finds sparse_fields.test record by querying it with criteria.
func (c *Client) FindSparseFieldsTest(criteria *Criteria) (*SparseFieldsTest, error) {
	sts := &SparseFieldsTests{}
	if err := c.SearchRead(SparseFieldsTestModel, criteria, NewOptions().Limit(1), sts); err != nil {
		return nil, err
	}
	if sts != nil && len(*sts) > 0 {
		return &((*sts)[0]), nil
	}
	return nil, fmt.Errorf("no sparse_fields.test was found with criteria %v", criteria)
}

// FindSparseFieldsTests finds sparse_fields.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSparseFieldsTests(criteria *Criteria, options *Options) (*SparseFieldsTests, error) {
	sts := &SparseFieldsTests{}
	if err := c.SearchRead(SparseFieldsTestModel, criteria, options, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSparseFieldsTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSparseFieldsTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SparseFieldsTestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSparseFieldsTestId finds record id by querying it with criteria.
func (c *Client) FindSparseFieldsTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SparseFieldsTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no sparse_fields.test was found with criteria %v and options %v", criteria, options)
}
