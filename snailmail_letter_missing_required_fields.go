package odoo

// SnailmailLetterMissingRequiredFields represents snailmail.letter.missing.required.fields model.
type SnailmailLetterMissingRequiredFields struct {
	City        *String   `xmlrpc:"city,omitempty"`
	CountryId   *Many2One `xmlrpc:"country_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	LetterId    *Many2One `xmlrpc:"letter_id,omitempty"`
	PartnerId   *Many2One `xmlrpc:"partner_id,omitempty"`
	StateId     *Many2One `xmlrpc:"state_id,omitempty"`
	Street      *String   `xmlrpc:"street,omitempty"`
	Street2     *String   `xmlrpc:"street2,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
	Zip         *String   `xmlrpc:"zip,omitempty"`
}

// SnailmailLetterMissingRequiredFieldss represents array of snailmail.letter.missing.required.fields model.
type SnailmailLetterMissingRequiredFieldss []SnailmailLetterMissingRequiredFields

// SnailmailLetterMissingRequiredFieldsModel is the odoo model name.
const SnailmailLetterMissingRequiredFieldsModel = "snailmail.letter.missing.required.fields"

// Many2One convert SnailmailLetterMissingRequiredFields to *Many2One.
func (slmrf *SnailmailLetterMissingRequiredFields) Many2One() *Many2One {
	return NewMany2One(slmrf.Id.Get(), "")
}

// CreateSnailmailLetterMissingRequiredFields creates a new snailmail.letter.missing.required.fields model and returns its id.
func (c *Client) CreateSnailmailLetterMissingRequiredFields(slmrf *SnailmailLetterMissingRequiredFields) (int64, error) {
	ids, err := c.CreateSnailmailLetterMissingRequiredFieldss([]*SnailmailLetterMissingRequiredFields{slmrf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSnailmailLetterMissingRequiredFields creates a new snailmail.letter.missing.required.fields model and returns its id.
func (c *Client) CreateSnailmailLetterMissingRequiredFieldss(slmrfs []*SnailmailLetterMissingRequiredFields) ([]int64, error) {
	var vv []interface{}
	for _, v := range slmrfs {
		vv = append(vv, v)
	}
	return c.Create(SnailmailLetterMissingRequiredFieldsModel, vv, nil)
}

// UpdateSnailmailLetterMissingRequiredFields updates an existing snailmail.letter.missing.required.fields record.
func (c *Client) UpdateSnailmailLetterMissingRequiredFields(slmrf *SnailmailLetterMissingRequiredFields) error {
	return c.UpdateSnailmailLetterMissingRequiredFieldss([]int64{slmrf.Id.Get()}, slmrf)
}

// UpdateSnailmailLetterMissingRequiredFieldss updates existing snailmail.letter.missing.required.fields records.
// All records (represented by ids) will be updated by slmrf values.
func (c *Client) UpdateSnailmailLetterMissingRequiredFieldss(ids []int64, slmrf *SnailmailLetterMissingRequiredFields) error {
	return c.Update(SnailmailLetterMissingRequiredFieldsModel, ids, slmrf, nil)
}

// DeleteSnailmailLetterMissingRequiredFields deletes an existing snailmail.letter.missing.required.fields record.
func (c *Client) DeleteSnailmailLetterMissingRequiredFields(id int64) error {
	return c.DeleteSnailmailLetterMissingRequiredFieldss([]int64{id})
}

// DeleteSnailmailLetterMissingRequiredFieldss deletes existing snailmail.letter.missing.required.fields records.
func (c *Client) DeleteSnailmailLetterMissingRequiredFieldss(ids []int64) error {
	return c.Delete(SnailmailLetterMissingRequiredFieldsModel, ids)
}

// GetSnailmailLetterMissingRequiredFields gets snailmail.letter.missing.required.fields existing record.
func (c *Client) GetSnailmailLetterMissingRequiredFields(id int64) (*SnailmailLetterMissingRequiredFields, error) {
	slmrfs, err := c.GetSnailmailLetterMissingRequiredFieldss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*slmrfs)[0]), nil
}

// GetSnailmailLetterMissingRequiredFieldss gets snailmail.letter.missing.required.fields existing records.
func (c *Client) GetSnailmailLetterMissingRequiredFieldss(ids []int64) (*SnailmailLetterMissingRequiredFieldss, error) {
	slmrfs := &SnailmailLetterMissingRequiredFieldss{}
	if err := c.Read(SnailmailLetterMissingRequiredFieldsModel, ids, nil, slmrfs); err != nil {
		return nil, err
	}
	return slmrfs, nil
}

// FindSnailmailLetterMissingRequiredFields finds snailmail.letter.missing.required.fields record by querying it with criteria.
func (c *Client) FindSnailmailLetterMissingRequiredFields(criteria *Criteria) (*SnailmailLetterMissingRequiredFields, error) {
	slmrfs := &SnailmailLetterMissingRequiredFieldss{}
	if err := c.SearchRead(SnailmailLetterMissingRequiredFieldsModel, criteria, NewOptions().Limit(1), slmrfs); err != nil {
		return nil, err
	}
	return &((*slmrfs)[0]), nil
}

// FindSnailmailLetterMissingRequiredFieldss finds snailmail.letter.missing.required.fields records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSnailmailLetterMissingRequiredFieldss(criteria *Criteria, options *Options) (*SnailmailLetterMissingRequiredFieldss, error) {
	slmrfs := &SnailmailLetterMissingRequiredFieldss{}
	if err := c.SearchRead(SnailmailLetterMissingRequiredFieldsModel, criteria, options, slmrfs); err != nil {
		return nil, err
	}
	return slmrfs, nil
}

// FindSnailmailLetterMissingRequiredFieldsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSnailmailLetterMissingRequiredFieldsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SnailmailLetterMissingRequiredFieldsModel, criteria, options)
}

// FindSnailmailLetterMissingRequiredFieldsId finds record id by querying it with criteria.
func (c *Client) FindSnailmailLetterMissingRequiredFieldsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SnailmailLetterMissingRequiredFieldsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
