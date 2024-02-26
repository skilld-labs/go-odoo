package odoo

// MailMassMailingTest represents mail.mass_mailing.test model.
type MailMassMailingTest struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	EmailTo       *String   `xmlrpc:"email_to,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	MassMailingId *Many2One `xmlrpc:"mass_mailing_id,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailMassMailingTests represents array of mail.mass_mailing.test model.
type MailMassMailingTests []MailMassMailingTest

// MailMassMailingTestModel is the odoo model name.
const MailMassMailingTestModel = "mail.mass_mailing.test"

// Many2One convert MailMassMailingTest to *Many2One.
func (mmt *MailMassMailingTest) Many2One() *Many2One {
	return NewMany2One(mmt.Id.Get(), "")
}

// CreateMailMassMailingTest creates a new mail.mass_mailing.test model and returns its id.
func (c *Client) CreateMailMassMailingTest(mmt *MailMassMailingTest) (int64, error) {
	ids, err := c.CreateMailMassMailingTests([]*MailMassMailingTest{mmt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMassMailingTest creates a new mail.mass_mailing.test model and returns its id.
func (c *Client) CreateMailMassMailingTests(mmts []*MailMassMailingTest) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmts {
		vv = append(vv, v)
	}
	return c.Create(MailMassMailingTestModel, vv, nil)
}

// UpdateMailMassMailingTest updates an existing mail.mass_mailing.test record.
func (c *Client) UpdateMailMassMailingTest(mmt *MailMassMailingTest) error {
	return c.UpdateMailMassMailingTests([]int64{mmt.Id.Get()}, mmt)
}

// UpdateMailMassMailingTests updates existing mail.mass_mailing.test records.
// All records (represented by ids) will be updated by mmt values.
func (c *Client) UpdateMailMassMailingTests(ids []int64, mmt *MailMassMailingTest) error {
	return c.Update(MailMassMailingTestModel, ids, mmt, nil)
}

// DeleteMailMassMailingTest deletes an existing mail.mass_mailing.test record.
func (c *Client) DeleteMailMassMailingTest(id int64) error {
	return c.DeleteMailMassMailingTests([]int64{id})
}

// DeleteMailMassMailingTests deletes existing mail.mass_mailing.test records.
func (c *Client) DeleteMailMassMailingTests(ids []int64) error {
	return c.Delete(MailMassMailingTestModel, ids)
}

// GetMailMassMailingTest gets mail.mass_mailing.test existing record.
func (c *Client) GetMailMassMailingTest(id int64) (*MailMassMailingTest, error) {
	mmts, err := c.GetMailMassMailingTests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mmts)[0]), nil
}

// GetMailMassMailingTests gets mail.mass_mailing.test existing records.
func (c *Client) GetMailMassMailingTests(ids []int64) (*MailMassMailingTests, error) {
	mmts := &MailMassMailingTests{}
	if err := c.Read(MailMassMailingTestModel, ids, nil, mmts); err != nil {
		return nil, err
	}
	return mmts, nil
}

// FindMailMassMailingTest finds mail.mass_mailing.test record by querying it with criteria.
func (c *Client) FindMailMassMailingTest(criteria *Criteria) (*MailMassMailingTest, error) {
	mmts := &MailMassMailingTests{}
	if err := c.SearchRead(MailMassMailingTestModel, criteria, NewOptions().Limit(1), mmts); err != nil {
		return nil, err
	}
	return &((*mmts)[0]), nil
}

// FindMailMassMailingTests finds mail.mass_mailing.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingTests(criteria *Criteria, options *Options) (*MailMassMailingTests, error) {
	mmts := &MailMassMailingTests{}
	if err := c.SearchRead(MailMassMailingTestModel, criteria, options, mmts); err != nil {
		return nil, err
	}
	return mmts, nil
}

// FindMailMassMailingTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMassMailingTestModel, criteria, options)
}

// FindMailMassMailingTestId finds record id by querying it with criteria.
func (c *Client) FindMailMassMailingTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMassMailingTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
