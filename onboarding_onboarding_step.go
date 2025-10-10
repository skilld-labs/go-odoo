package odoo

// OnboardingOnboardingStep represents onboarding.onboarding.step model.
type OnboardingOnboardingStep struct {
	ButtonText              *String    `xmlrpc:"button_text,omitempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrentProgressStepId   *Many2One  `xmlrpc:"current_progress_step_id,omitempty"`
	CurrentStepState        *Selection `xmlrpc:"current_step_state,omitempty"`
	Description             *String    `xmlrpc:"description,omitempty"`
	DisplayName             *String    `xmlrpc:"display_name,omitempty"`
	DoneIcon                *String    `xmlrpc:"done_icon,omitempty"`
	DoneText                *String    `xmlrpc:"done_text,omitempty"`
	Id                      *Int       `xmlrpc:"id,omitempty"`
	IsPerCompany            *Bool      `xmlrpc:"is_per_company,omitempty"`
	OnboardingIds           *Relation  `xmlrpc:"onboarding_ids,omitempty"`
	PanelStepOpenActionName *String    `xmlrpc:"panel_step_open_action_name,omitempty"`
	ProgressIds             *Relation  `xmlrpc:"progress_ids,omitempty"`
	Sequence                *Int       `xmlrpc:"sequence,omitempty"`
	StepImage               *String    `xmlrpc:"step_image,omitempty"`
	StepImageAlt            *String    `xmlrpc:"step_image_alt,omitempty"`
	StepImageFilename       *String    `xmlrpc:"step_image_filename,omitempty"`
	Title                   *String    `xmlrpc:"title,omitempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// OnboardingOnboardingSteps represents array of onboarding.onboarding.step model.
type OnboardingOnboardingSteps []OnboardingOnboardingStep

// OnboardingOnboardingStepModel is the odoo model name.
const OnboardingOnboardingStepModel = "onboarding.onboarding.step"

// Many2One convert OnboardingOnboardingStep to *Many2One.
func (oos *OnboardingOnboardingStep) Many2One() *Many2One {
	return NewMany2One(oos.Id.Get(), "")
}

// CreateOnboardingOnboardingStep creates a new onboarding.onboarding.step model and returns its id.
func (c *Client) CreateOnboardingOnboardingStep(oos *OnboardingOnboardingStep) (int64, error) {
	ids, err := c.CreateOnboardingOnboardingSteps([]*OnboardingOnboardingStep{oos})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateOnboardingOnboardingStep creates a new onboarding.onboarding.step model and returns its id.
func (c *Client) CreateOnboardingOnboardingSteps(ooss []*OnboardingOnboardingStep) ([]int64, error) {
	var vv []interface{}
	for _, v := range ooss {
		vv = append(vv, v)
	}
	return c.Create(OnboardingOnboardingStepModel, vv, nil)
}

// UpdateOnboardingOnboardingStep updates an existing onboarding.onboarding.step record.
func (c *Client) UpdateOnboardingOnboardingStep(oos *OnboardingOnboardingStep) error {
	return c.UpdateOnboardingOnboardingSteps([]int64{oos.Id.Get()}, oos)
}

// UpdateOnboardingOnboardingSteps updates existing onboarding.onboarding.step records.
// All records (represented by ids) will be updated by oos values.
func (c *Client) UpdateOnboardingOnboardingSteps(ids []int64, oos *OnboardingOnboardingStep) error {
	return c.Update(OnboardingOnboardingStepModel, ids, oos, nil)
}

// DeleteOnboardingOnboardingStep deletes an existing onboarding.onboarding.step record.
func (c *Client) DeleteOnboardingOnboardingStep(id int64) error {
	return c.DeleteOnboardingOnboardingSteps([]int64{id})
}

// DeleteOnboardingOnboardingSteps deletes existing onboarding.onboarding.step records.
func (c *Client) DeleteOnboardingOnboardingSteps(ids []int64) error {
	return c.Delete(OnboardingOnboardingStepModel, ids)
}

// GetOnboardingOnboardingStep gets onboarding.onboarding.step existing record.
func (c *Client) GetOnboardingOnboardingStep(id int64) (*OnboardingOnboardingStep, error) {
	ooss, err := c.GetOnboardingOnboardingSteps([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ooss)[0]), nil
}

// GetOnboardingOnboardingSteps gets onboarding.onboarding.step existing records.
func (c *Client) GetOnboardingOnboardingSteps(ids []int64) (*OnboardingOnboardingSteps, error) {
	ooss := &OnboardingOnboardingSteps{}
	if err := c.Read(OnboardingOnboardingStepModel, ids, nil, ooss); err != nil {
		return nil, err
	}
	return ooss, nil
}

// FindOnboardingOnboardingStep finds onboarding.onboarding.step record by querying it with criteria.
func (c *Client) FindOnboardingOnboardingStep(criteria *Criteria) (*OnboardingOnboardingStep, error) {
	ooss := &OnboardingOnboardingSteps{}
	if err := c.SearchRead(OnboardingOnboardingStepModel, criteria, NewOptions().Limit(1), ooss); err != nil {
		return nil, err
	}
	return &((*ooss)[0]), nil
}

// FindOnboardingOnboardingSteps finds onboarding.onboarding.step records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOnboardingOnboardingSteps(criteria *Criteria, options *Options) (*OnboardingOnboardingSteps, error) {
	ooss := &OnboardingOnboardingSteps{}
	if err := c.SearchRead(OnboardingOnboardingStepModel, criteria, options, ooss); err != nil {
		return nil, err
	}
	return ooss, nil
}

// FindOnboardingOnboardingStepIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOnboardingOnboardingStepIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(OnboardingOnboardingStepModel, criteria, options)
}

// FindOnboardingOnboardingStepId finds record id by querying it with criteria.
func (c *Client) FindOnboardingOnboardingStepId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OnboardingOnboardingStepModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
