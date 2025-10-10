package odoo

// MailingContactImport represents mailing.contact.import model.
type MailingContactImport struct {
	ContactList    *String   `xmlrpc:"contact_list,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	MailingListIds *Relation `xmlrpc:"mailing_list_ids,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailingContactImports represents array of mailing.contact.import model.
type MailingContactImports []MailingContactImport

// MailingContactImportModel is the odoo model name.
const MailingContactImportModel = "mailing.contact.import"

// Many2One convert MailingContactImport to *Many2One.
func (mci *MailingContactImport) Many2One() *Many2One {
	return NewMany2One(mci.Id.Get(), "")
}

// CreateMailingContactImport creates a new mailing.contact.import model and returns its id.
func (c *Client) CreateMailingContactImport(mci *MailingContactImport) (int64, error) {
	ids, err := c.CreateMailingContactImports([]*MailingContactImport{mci})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingContactImport creates a new mailing.contact.import model and returns its id.
func (c *Client) CreateMailingContactImports(mcis []*MailingContactImport) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcis {
		vv = append(vv, v)
	}
	return c.Create(MailingContactImportModel, vv, nil)
}

// UpdateMailingContactImport updates an existing mailing.contact.import record.
func (c *Client) UpdateMailingContactImport(mci *MailingContactImport) error {
	return c.UpdateMailingContactImports([]int64{mci.Id.Get()}, mci)
}

// UpdateMailingContactImports updates existing mailing.contact.import records.
// All records (represented by ids) will be updated by mci values.
func (c *Client) UpdateMailingContactImports(ids []int64, mci *MailingContactImport) error {
	return c.Update(MailingContactImportModel, ids, mci, nil)
}

// DeleteMailingContactImport deletes an existing mailing.contact.import record.
func (c *Client) DeleteMailingContactImport(id int64) error {
	return c.DeleteMailingContactImports([]int64{id})
}

// DeleteMailingContactImports deletes existing mailing.contact.import records.
func (c *Client) DeleteMailingContactImports(ids []int64) error {
	return c.Delete(MailingContactImportModel, ids)
}

// GetMailingContactImport gets mailing.contact.import existing record.
func (c *Client) GetMailingContactImport(id int64) (*MailingContactImport, error) {
	mcis, err := c.GetMailingContactImports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcis)[0]), nil
}

// GetMailingContactImports gets mailing.contact.import existing records.
func (c *Client) GetMailingContactImports(ids []int64) (*MailingContactImports, error) {
	mcis := &MailingContactImports{}
	if err := c.Read(MailingContactImportModel, ids, nil, mcis); err != nil {
		return nil, err
	}
	return mcis, nil
}

// FindMailingContactImport finds mailing.contact.import record by querying it with criteria.
func (c *Client) FindMailingContactImport(criteria *Criteria) (*MailingContactImport, error) {
	mcis := &MailingContactImports{}
	if err := c.SearchRead(MailingContactImportModel, criteria, NewOptions().Limit(1), mcis); err != nil {
		return nil, err
	}
	return &((*mcis)[0]), nil
}

// FindMailingContactImports finds mailing.contact.import records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingContactImports(criteria *Criteria, options *Options) (*MailingContactImports, error) {
	mcis := &MailingContactImports{}
	if err := c.SearchRead(MailingContactImportModel, criteria, options, mcis); err != nil {
		return nil, err
	}
	return mcis, nil
}

// FindMailingContactImportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingContactImportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailingContactImportModel, criteria, options)
}

// FindMailingContactImportId finds record id by querying it with criteria.
func (c *Client) FindMailingContactImportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingContactImportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
