package odoo

// MailingMailingTest represents mailing.mailing.test model.
type MailingMailingTest struct {
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	EmailTo       *String   `xmlrpc:"email_to,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	MassMailingId *Many2One `xmlrpc:"mass_mailing_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailingMailingTests represents array of mailing.mailing.test model.
type MailingMailingTests []MailingMailingTest

// MailingMailingTestModel is the odoo model name.
const MailingMailingTestModel = "mailing.mailing.test"

// Many2One convert MailingMailingTest to *Many2One.
func (mmt *MailingMailingTest) Many2One() *Many2One {
	return NewMany2One(mmt.Id.Get(), "")
}

// CreateMailingMailingTest creates a new mailing.mailing.test model and returns its id.
func (c *Client) CreateMailingMailingTest(mmt *MailingMailingTest) (int64, error) {
	ids, err := c.CreateMailingMailingTests([]*MailingMailingTest{mmt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingMailingTest creates a new mailing.mailing.test model and returns its id.
func (c *Client) CreateMailingMailingTests(mmts []*MailingMailingTest) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmts {
		vv = append(vv, v)
	}
	return c.Create(MailingMailingTestModel, vv, nil)
}

// UpdateMailingMailingTest updates an existing mailing.mailing.test record.
func (c *Client) UpdateMailingMailingTest(mmt *MailingMailingTest) error {
	return c.UpdateMailingMailingTests([]int64{mmt.Id.Get()}, mmt)
}

// UpdateMailingMailingTests updates existing mailing.mailing.test records.
// All records (represented by ids) will be updated by mmt values.
func (c *Client) UpdateMailingMailingTests(ids []int64, mmt *MailingMailingTest) error {
	return c.Update(MailingMailingTestModel, ids, mmt, nil)
}

// DeleteMailingMailingTest deletes an existing mailing.mailing.test record.
func (c *Client) DeleteMailingMailingTest(id int64) error {
	return c.DeleteMailingMailingTests([]int64{id})
}

// DeleteMailingMailingTests deletes existing mailing.mailing.test records.
func (c *Client) DeleteMailingMailingTests(ids []int64) error {
	return c.Delete(MailingMailingTestModel, ids)
}

// GetMailingMailingTest gets mailing.mailing.test existing record.
func (c *Client) GetMailingMailingTest(id int64) (*MailingMailingTest, error) {
	mmts, err := c.GetMailingMailingTests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mmts)[0]), nil
}

// GetMailingMailingTests gets mailing.mailing.test existing records.
func (c *Client) GetMailingMailingTests(ids []int64) (*MailingMailingTests, error) {
	mmts := &MailingMailingTests{}
	if err := c.Read(MailingMailingTestModel, ids, nil, mmts); err != nil {
		return nil, err
	}
	return mmts, nil
}

// FindMailingMailingTest finds mailing.mailing.test record by querying it with criteria.
func (c *Client) FindMailingMailingTest(criteria *Criteria) (*MailingMailingTest, error) {
	mmts := &MailingMailingTests{}
	if err := c.SearchRead(MailingMailingTestModel, criteria, NewOptions().Limit(1), mmts); err != nil {
		return nil, err
	}
	return &((*mmts)[0]), nil
}

// FindMailingMailingTests finds mailing.mailing.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingMailingTests(criteria *Criteria, options *Options) (*MailingMailingTests, error) {
	mmts := &MailingMailingTests{}
	if err := c.SearchRead(MailingMailingTestModel, criteria, options, mmts); err != nil {
		return nil, err
	}
	return mmts, nil
}

// FindMailingMailingTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingMailingTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailingMailingTestModel, criteria, options)
}

// FindMailingMailingTestId finds record id by querying it with criteria.
func (c *Client) FindMailingMailingTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingMailingTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
