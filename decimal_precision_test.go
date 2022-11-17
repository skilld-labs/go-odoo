package odoo

import (
	"fmt"
)

// DecimalPrecisionTest represents decimal.precision.test model.
type DecimalPrecisionTest struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Float       *Float    `xmlrpc:"float,omitempty"`
	Float2      *Float    `xmlrpc:"float_2,omitempty"`
	Float4      *Float    `xmlrpc:"float_4,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DecimalPrecisionTests represents array of decimal.precision.test model.
type DecimalPrecisionTests []DecimalPrecisionTest

// DecimalPrecisionTestModel is the odoo model name.
const DecimalPrecisionTestModel = "decimal.precision.test"

// Many2One convert DecimalPrecisionTest to *Many2One.
func (dpt *DecimalPrecisionTest) Many2One() *Many2One {
	return NewMany2One(dpt.Id.Get(), "")
}

// CreateDecimalPrecisionTest creates a new decimal.precision.test model and returns its id.
func (c *Client) CreateDecimalPrecisionTest(dpt *DecimalPrecisionTest) (int64, error) {
	return c.Create(DecimalPrecisionTestModel, dpt)
}

// UpdateDecimalPrecisionTest updates an existing decimal.precision.test record.
func (c *Client) UpdateDecimalPrecisionTest(dpt *DecimalPrecisionTest) error {
	return c.UpdateDecimalPrecisionTests([]int64{dpt.Id.Get()}, dpt)
}

// UpdateDecimalPrecisionTests updates existing decimal.precision.test records.
// All records (represented by ids) will be updated by dpt values.
func (c *Client) UpdateDecimalPrecisionTests(ids []int64, dpt *DecimalPrecisionTest) error {
	return c.Update(DecimalPrecisionTestModel, ids, dpt)
}

// DeleteDecimalPrecisionTest deletes an existing decimal.precision.test record.
func (c *Client) DeleteDecimalPrecisionTest(id int64) error {
	return c.DeleteDecimalPrecisionTests([]int64{id})
}

// DeleteDecimalPrecisionTests deletes existing decimal.precision.test records.
func (c *Client) DeleteDecimalPrecisionTests(ids []int64) error {
	return c.Delete(DecimalPrecisionTestModel, ids)
}

// GetDecimalPrecisionTest gets decimal.precision.test existing record.
func (c *Client) GetDecimalPrecisionTest(id int64) (*DecimalPrecisionTest, error) {
	dpts, err := c.GetDecimalPrecisionTests([]int64{id})
	if err != nil {
		return nil, err
	}
	if dpts != nil && len(*dpts) > 0 {
		return &((*dpts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of decimal.precision.test not found", id)
}

// GetDecimalPrecisionTests gets decimal.precision.test existing records.
func (c *Client) GetDecimalPrecisionTests(ids []int64) (*DecimalPrecisionTests, error) {
	dpts := &DecimalPrecisionTests{}
	if err := c.Read(DecimalPrecisionTestModel, ids, nil, dpts); err != nil {
		return nil, err
	}
	return dpts, nil
}

// FindDecimalPrecisionTest finds decimal.precision.test record by querying it with criteria.
func (c *Client) FindDecimalPrecisionTest(criteria *Criteria) (*DecimalPrecisionTest, error) {
	dpts := &DecimalPrecisionTests{}
	if err := c.SearchRead(DecimalPrecisionTestModel, criteria, NewOptions().Limit(1), dpts); err != nil {
		return nil, err
	}
	if dpts != nil && len(*dpts) > 0 {
		return &((*dpts)[0]), nil
	}
	return nil, fmt.Errorf("no decimal.precision.test was found with criteria %v", criteria)
}

// FindDecimalPrecisionTests finds decimal.precision.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDecimalPrecisionTests(criteria *Criteria, options *Options) (*DecimalPrecisionTests, error) {
	dpts := &DecimalPrecisionTests{}
	if err := c.SearchRead(DecimalPrecisionTestModel, criteria, options, dpts); err != nil {
		return nil, err
	}
	return dpts, nil
}

// FindDecimalPrecisionTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDecimalPrecisionTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DecimalPrecisionTestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDecimalPrecisionTestId finds record id by querying it with criteria.
func (c *Client) FindDecimalPrecisionTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DecimalPrecisionTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no decimal.precision.test was found with criteria %v and options %v", criteria, options)
}
