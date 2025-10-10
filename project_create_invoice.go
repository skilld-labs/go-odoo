package odoo

// ProjectCreateInvoice represents project.create.invoice model.
type ProjectCreateInvoice struct {
	CandidateOrders *Relation `xmlrpc:"_candidate_orders,omitempty"`
	AmountToInvoice *Float    `xmlrpc:"amount_to_invoice,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId      *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	ProjectId       *Many2One `xmlrpc:"project_id,omitempty"`
	SaleOrderId     *Many2One `xmlrpc:"sale_order_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectCreateInvoices represents array of project.create.invoice model.
type ProjectCreateInvoices []ProjectCreateInvoice

// ProjectCreateInvoiceModel is the odoo model name.
const ProjectCreateInvoiceModel = "project.create.invoice"

// Many2One convert ProjectCreateInvoice to *Many2One.
func (pci *ProjectCreateInvoice) Many2One() *Many2One {
	return NewMany2One(pci.Id.Get(), "")
}

// CreateProjectCreateInvoice creates a new project.create.invoice model and returns its id.
func (c *Client) CreateProjectCreateInvoice(pci *ProjectCreateInvoice) (int64, error) {
	ids, err := c.CreateProjectCreateInvoices([]*ProjectCreateInvoice{pci})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectCreateInvoice creates a new project.create.invoice model and returns its id.
func (c *Client) CreateProjectCreateInvoices(pcis []*ProjectCreateInvoice) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcis {
		vv = append(vv, v)
	}
	return c.Create(ProjectCreateInvoiceModel, vv, nil)
}

// UpdateProjectCreateInvoice updates an existing project.create.invoice record.
func (c *Client) UpdateProjectCreateInvoice(pci *ProjectCreateInvoice) error {
	return c.UpdateProjectCreateInvoices([]int64{pci.Id.Get()}, pci)
}

// UpdateProjectCreateInvoices updates existing project.create.invoice records.
// All records (represented by ids) will be updated by pci values.
func (c *Client) UpdateProjectCreateInvoices(ids []int64, pci *ProjectCreateInvoice) error {
	return c.Update(ProjectCreateInvoiceModel, ids, pci, nil)
}

// DeleteProjectCreateInvoice deletes an existing project.create.invoice record.
func (c *Client) DeleteProjectCreateInvoice(id int64) error {
	return c.DeleteProjectCreateInvoices([]int64{id})
}

// DeleteProjectCreateInvoices deletes existing project.create.invoice records.
func (c *Client) DeleteProjectCreateInvoices(ids []int64) error {
	return c.Delete(ProjectCreateInvoiceModel, ids)
}

// GetProjectCreateInvoice gets project.create.invoice existing record.
func (c *Client) GetProjectCreateInvoice(id int64) (*ProjectCreateInvoice, error) {
	pcis, err := c.GetProjectCreateInvoices([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcis)[0]), nil
}

// GetProjectCreateInvoices gets project.create.invoice existing records.
func (c *Client) GetProjectCreateInvoices(ids []int64) (*ProjectCreateInvoices, error) {
	pcis := &ProjectCreateInvoices{}
	if err := c.Read(ProjectCreateInvoiceModel, ids, nil, pcis); err != nil {
		return nil, err
	}
	return pcis, nil
}

// FindProjectCreateInvoice finds project.create.invoice record by querying it with criteria.
func (c *Client) FindProjectCreateInvoice(criteria *Criteria) (*ProjectCreateInvoice, error) {
	pcis := &ProjectCreateInvoices{}
	if err := c.SearchRead(ProjectCreateInvoiceModel, criteria, NewOptions().Limit(1), pcis); err != nil {
		return nil, err
	}
	return &((*pcis)[0]), nil
}

// FindProjectCreateInvoices finds project.create.invoice records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectCreateInvoices(criteria *Criteria, options *Options) (*ProjectCreateInvoices, error) {
	pcis := &ProjectCreateInvoices{}
	if err := c.SearchRead(ProjectCreateInvoiceModel, criteria, options, pcis); err != nil {
		return nil, err
	}
	return pcis, nil
}

// FindProjectCreateInvoiceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectCreateInvoiceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectCreateInvoiceModel, criteria, options)
}

// FindProjectCreateInvoiceId finds record id by querying it with criteria.
func (c *Client) FindProjectCreateInvoiceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectCreateInvoiceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
