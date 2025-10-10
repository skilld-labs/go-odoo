package odoo

// SalePdfFormField represents sale.pdf.form.field model.
type SalePdfFormField struct {
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	DocumentType         *Selection `xmlrpc:"document_type,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	Name                 *String    `xmlrpc:"name,omitempty"`
	Path                 *String    `xmlrpc:"path,omitempty"`
	ProductDocumentIds   *Relation  `xmlrpc:"product_document_ids,omitempty"`
	QuotationDocumentIds *Relation  `xmlrpc:"quotation_document_ids,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SalePdfFormFields represents array of sale.pdf.form.field model.
type SalePdfFormFields []SalePdfFormField

// SalePdfFormFieldModel is the odoo model name.
const SalePdfFormFieldModel = "sale.pdf.form.field"

// Many2One convert SalePdfFormField to *Many2One.
func (spff *SalePdfFormField) Many2One() *Many2One {
	return NewMany2One(spff.Id.Get(), "")
}

// CreateSalePdfFormField creates a new sale.pdf.form.field model and returns its id.
func (c *Client) CreateSalePdfFormField(spff *SalePdfFormField) (int64, error) {
	ids, err := c.CreateSalePdfFormFields([]*SalePdfFormField{spff})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSalePdfFormField creates a new sale.pdf.form.field model and returns its id.
func (c *Client) CreateSalePdfFormFields(spffs []*SalePdfFormField) ([]int64, error) {
	var vv []interface{}
	for _, v := range spffs {
		vv = append(vv, v)
	}
	return c.Create(SalePdfFormFieldModel, vv, nil)
}

// UpdateSalePdfFormField updates an existing sale.pdf.form.field record.
func (c *Client) UpdateSalePdfFormField(spff *SalePdfFormField) error {
	return c.UpdateSalePdfFormFields([]int64{spff.Id.Get()}, spff)
}

// UpdateSalePdfFormFields updates existing sale.pdf.form.field records.
// All records (represented by ids) will be updated by spff values.
func (c *Client) UpdateSalePdfFormFields(ids []int64, spff *SalePdfFormField) error {
	return c.Update(SalePdfFormFieldModel, ids, spff, nil)
}

// DeleteSalePdfFormField deletes an existing sale.pdf.form.field record.
func (c *Client) DeleteSalePdfFormField(id int64) error {
	return c.DeleteSalePdfFormFields([]int64{id})
}

// DeleteSalePdfFormFields deletes existing sale.pdf.form.field records.
func (c *Client) DeleteSalePdfFormFields(ids []int64) error {
	return c.Delete(SalePdfFormFieldModel, ids)
}

// GetSalePdfFormField gets sale.pdf.form.field existing record.
func (c *Client) GetSalePdfFormField(id int64) (*SalePdfFormField, error) {
	spffs, err := c.GetSalePdfFormFields([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*spffs)[0]), nil
}

// GetSalePdfFormFields gets sale.pdf.form.field existing records.
func (c *Client) GetSalePdfFormFields(ids []int64) (*SalePdfFormFields, error) {
	spffs := &SalePdfFormFields{}
	if err := c.Read(SalePdfFormFieldModel, ids, nil, spffs); err != nil {
		return nil, err
	}
	return spffs, nil
}

// FindSalePdfFormField finds sale.pdf.form.field record by querying it with criteria.
func (c *Client) FindSalePdfFormField(criteria *Criteria) (*SalePdfFormField, error) {
	spffs := &SalePdfFormFields{}
	if err := c.SearchRead(SalePdfFormFieldModel, criteria, NewOptions().Limit(1), spffs); err != nil {
		return nil, err
	}
	return &((*spffs)[0]), nil
}

// FindSalePdfFormFields finds sale.pdf.form.field records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSalePdfFormFields(criteria *Criteria, options *Options) (*SalePdfFormFields, error) {
	spffs := &SalePdfFormFields{}
	if err := c.SearchRead(SalePdfFormFieldModel, criteria, options, spffs); err != nil {
		return nil, err
	}
	return spffs, nil
}

// FindSalePdfFormFieldIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSalePdfFormFieldIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SalePdfFormFieldModel, criteria, options)
}

// FindSalePdfFormFieldId finds record id by querying it with criteria.
func (c *Client) FindSalePdfFormFieldId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SalePdfFormFieldModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
