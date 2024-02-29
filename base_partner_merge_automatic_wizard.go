package odoo

// BasePartnerMergeAutomaticWizard represents base.partner.merge.automatic.wizard model.
type BasePartnerMergeAutomaticWizard struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrentLineId      *Many2One  `xmlrpc:"current_line_id,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	DstPartnerId       *Many2One  `xmlrpc:"dst_partner_id,omitempty"`
	ExcludeContact     *Bool      `xmlrpc:"exclude_contact,omitempty"`
	ExcludeJournalItem *Bool      `xmlrpc:"exclude_journal_item,omitempty"`
	GroupByEmail       *Bool      `xmlrpc:"group_by_email,omitempty"`
	GroupByIsCompany   *Bool      `xmlrpc:"group_by_is_company,omitempty"`
	GroupByName        *Bool      `xmlrpc:"group_by_name,omitempty"`
	GroupByParentId    *Bool      `xmlrpc:"group_by_parent_id,omitempty"`
	GroupByVat         *Bool      `xmlrpc:"group_by_vat,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	LineIds            *Relation  `xmlrpc:"line_ids,omitempty"`
	MaximumGroup       *Int       `xmlrpc:"maximum_group,omitempty"`
	NumberGroup        *Int       `xmlrpc:"number_group,omitempty"`
	PartnerIds         *Relation  `xmlrpc:"partner_ids,omitempty"`
	State              *Selection `xmlrpc:"state,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// BasePartnerMergeAutomaticWizards represents array of base.partner.merge.automatic.wizard model.
type BasePartnerMergeAutomaticWizards []BasePartnerMergeAutomaticWizard

// BasePartnerMergeAutomaticWizardModel is the odoo model name.
const BasePartnerMergeAutomaticWizardModel = "base.partner.merge.automatic.wizard"

// Many2One convert BasePartnerMergeAutomaticWizard to *Many2One.
func (bpmaw *BasePartnerMergeAutomaticWizard) Many2One() *Many2One {
	return NewMany2One(bpmaw.Id.Get(), "")
}

// CreateBasePartnerMergeAutomaticWizard creates a new base.partner.merge.automatic.wizard model and returns its id.
func (c *Client) CreateBasePartnerMergeAutomaticWizard(bpmaw *BasePartnerMergeAutomaticWizard) (int64, error) {
	ids, err := c.CreateBasePartnerMergeAutomaticWizards([]*BasePartnerMergeAutomaticWizard{bpmaw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBasePartnerMergeAutomaticWizards creates a new base.partner.merge.automatic.wizard model and returns its id.
func (c *Client) CreateBasePartnerMergeAutomaticWizards(bpmaws []*BasePartnerMergeAutomaticWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range bpmaws {
		vv = append(vv, v)
	}
	return c.Create(BasePartnerMergeAutomaticWizardModel, vv, nil)
}

// UpdateBasePartnerMergeAutomaticWizard updates an existing base.partner.merge.automatic.wizard record.
func (c *Client) UpdateBasePartnerMergeAutomaticWizard(bpmaw *BasePartnerMergeAutomaticWizard) error {
	return c.UpdateBasePartnerMergeAutomaticWizards([]int64{bpmaw.Id.Get()}, bpmaw)
}

// UpdateBasePartnerMergeAutomaticWizards updates existing base.partner.merge.automatic.wizard records.
// All records (represented by ids) will be updated by bpmaw values.
func (c *Client) UpdateBasePartnerMergeAutomaticWizards(ids []int64, bpmaw *BasePartnerMergeAutomaticWizard) error {
	return c.Update(BasePartnerMergeAutomaticWizardModel, ids, bpmaw, nil)
}

// DeleteBasePartnerMergeAutomaticWizard deletes an existing base.partner.merge.automatic.wizard record.
func (c *Client) DeleteBasePartnerMergeAutomaticWizard(id int64) error {
	return c.DeleteBasePartnerMergeAutomaticWizards([]int64{id})
}

// DeleteBasePartnerMergeAutomaticWizards deletes existing base.partner.merge.automatic.wizard records.
func (c *Client) DeleteBasePartnerMergeAutomaticWizards(ids []int64) error {
	return c.Delete(BasePartnerMergeAutomaticWizardModel, ids)
}

// GetBasePartnerMergeAutomaticWizard gets base.partner.merge.automatic.wizard existing record.
func (c *Client) GetBasePartnerMergeAutomaticWizard(id int64) (*BasePartnerMergeAutomaticWizard, error) {
	bpmaws, err := c.GetBasePartnerMergeAutomaticWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bpmaws)[0]), nil
}

// GetBasePartnerMergeAutomaticWizards gets base.partner.merge.automatic.wizard existing records.
func (c *Client) GetBasePartnerMergeAutomaticWizards(ids []int64) (*BasePartnerMergeAutomaticWizards, error) {
	bpmaws := &BasePartnerMergeAutomaticWizards{}
	if err := c.Read(BasePartnerMergeAutomaticWizardModel, ids, nil, bpmaws); err != nil {
		return nil, err
	}
	return bpmaws, nil
}

// FindBasePartnerMergeAutomaticWizard finds base.partner.merge.automatic.wizard record by querying it with criteria.
func (c *Client) FindBasePartnerMergeAutomaticWizard(criteria *Criteria) (*BasePartnerMergeAutomaticWizard, error) {
	bpmaws := &BasePartnerMergeAutomaticWizards{}
	if err := c.SearchRead(BasePartnerMergeAutomaticWizardModel, criteria, NewOptions().Limit(1), bpmaws); err != nil {
		return nil, err
	}
	return &((*bpmaws)[0]), nil
}

// FindBasePartnerMergeAutomaticWizards finds base.partner.merge.automatic.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBasePartnerMergeAutomaticWizards(criteria *Criteria, options *Options) (*BasePartnerMergeAutomaticWizards, error) {
	bpmaws := &BasePartnerMergeAutomaticWizards{}
	if err := c.SearchRead(BasePartnerMergeAutomaticWizardModel, criteria, options, bpmaws); err != nil {
		return nil, err
	}
	return bpmaws, nil
}

// FindBasePartnerMergeAutomaticWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBasePartnerMergeAutomaticWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BasePartnerMergeAutomaticWizardModel, criteria, options)
}

// FindBasePartnerMergeAutomaticWizardId finds record id by querying it with criteria.
func (c *Client) FindBasePartnerMergeAutomaticWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BasePartnerMergeAutomaticWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
