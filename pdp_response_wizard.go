package odoo

// PdpResponseWizard represents pdp.response.wizard model.
type PdpResponseWizard struct {
	AvailableStatuses *String    `xmlrpc:"available_statuses,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId        *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	FullyPaid         *Bool      `xmlrpc:"fully_paid,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	MoveCount         *Int       `xmlrpc:"move_count,omitempty"`
	MoveIds           *Relation  `xmlrpc:"move_ids,omitempty"`
	Note              *String    `xmlrpc:"note,omitempty"`
	PaidAmount        *Float     `xmlrpc:"paid_amount,omitempty"`
	ReasonCode        *Selection `xmlrpc:"reason_code,omitempty"`
	ShowReasonCode    *Bool      `xmlrpc:"show_reason_code,omitempty"`
	Status            *Selection `xmlrpc:"status,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PdpResponseWizards represents array of pdp.response.wizard model.
type PdpResponseWizards []PdpResponseWizard

// PdpResponseWizardModel is the odoo model name.
const PdpResponseWizardModel = "pdp.response.wizard"

// Many2One convert PdpResponseWizard to *Many2One.
func (prw *PdpResponseWizard) Many2One() *Many2One {
	return NewMany2One(prw.Id.Get(), "")
}

// CreatePdpResponseWizard creates a new pdp.response.wizard model and returns its id.
func (c *Client) CreatePdpResponseWizard(prw *PdpResponseWizard) (int64, error) {
	ids, err := c.CreatePdpResponseWizards([]*PdpResponseWizard{prw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePdpResponseWizard creates a new pdp.response.wizard model and returns its id.
func (c *Client) CreatePdpResponseWizards(prws []*PdpResponseWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range prws {
		vv = append(vv, v)
	}
	return c.Create(PdpResponseWizardModel, vv, nil)
}

// UpdatePdpResponseWizard updates an existing pdp.response.wizard record.
func (c *Client) UpdatePdpResponseWizard(prw *PdpResponseWizard) error {
	return c.UpdatePdpResponseWizards([]int64{prw.Id.Get()}, prw)
}

// UpdatePdpResponseWizards updates existing pdp.response.wizard records.
// All records (represented by ids) will be updated by prw values.
func (c *Client) UpdatePdpResponseWizards(ids []int64, prw *PdpResponseWizard) error {
	return c.Update(PdpResponseWizardModel, ids, prw, nil)
}

// DeletePdpResponseWizard deletes an existing pdp.response.wizard record.
func (c *Client) DeletePdpResponseWizard(id int64) error {
	return c.DeletePdpResponseWizards([]int64{id})
}

// DeletePdpResponseWizards deletes existing pdp.response.wizard records.
func (c *Client) DeletePdpResponseWizards(ids []int64) error {
	return c.Delete(PdpResponseWizardModel, ids)
}

// GetPdpResponseWizard gets pdp.response.wizard existing record.
func (c *Client) GetPdpResponseWizard(id int64) (*PdpResponseWizard, error) {
	prws, err := c.GetPdpResponseWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*prws)[0]), nil
}

// GetPdpResponseWizards gets pdp.response.wizard existing records.
func (c *Client) GetPdpResponseWizards(ids []int64) (*PdpResponseWizards, error) {
	prws := &PdpResponseWizards{}
	if err := c.Read(PdpResponseWizardModel, ids, nil, prws); err != nil {
		return nil, err
	}
	return prws, nil
}

// FindPdpResponseWizard finds pdp.response.wizard record by querying it with criteria.
func (c *Client) FindPdpResponseWizard(criteria *Criteria) (*PdpResponseWizard, error) {
	prws := &PdpResponseWizards{}
	if err := c.SearchRead(PdpResponseWizardModel, criteria, NewOptions().Limit(1), prws); err != nil {
		return nil, err
	}
	return &((*prws)[0]), nil
}

// FindPdpResponseWizards finds pdp.response.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPdpResponseWizards(criteria *Criteria, options *Options) (*PdpResponseWizards, error) {
	prws := &PdpResponseWizards{}
	if err := c.SearchRead(PdpResponseWizardModel, criteria, options, prws); err != nil {
		return nil, err
	}
	return prws, nil
}

// FindPdpResponseWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPdpResponseWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PdpResponseWizardModel, criteria, options)
}

// FindPdpResponseWizardId finds record id by querying it with criteria.
func (c *Client) FindPdpResponseWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PdpResponseWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
