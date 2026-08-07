package odoo

// AccountPeppolResponse represents account.peppol.response model.
type AccountPeppolResponse struct {
	CompanyId          *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate         *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String     `xmlrpc:"display_name,omitempty"`
	Id                 *Int        `xmlrpc:"id,omitempty"`
	MoveId             *Many2One   `xmlrpc:"move_id,omitempty"`
	PdpFlowNumber      *Selection  `xmlrpc:"pdp_flow_number,omitempty"`
	PdpIssueDate       *Time       `xmlrpc:"pdp_issue_date,omitempty"`
	PdpPaymentInfo     interface{} `xmlrpc:"pdp_payment_info,omitempty"`
	PdpPpfState        *Selection  `xmlrpc:"pdp_ppf_state,omitempty"`
	PdpRefResponseCode *Selection  `xmlrpc:"pdp_ref_response_code,omitempty"`
	PdpRefUuid         *String     `xmlrpc:"pdp_ref_uuid,omitempty"`
	PdpStatusInfo      *String     `xmlrpc:"pdp_status_info,omitempty"`
	PeppolMessageUuid  *String     `xmlrpc:"peppol_message_uuid,omitempty"`
	PeppolState        *Selection  `xmlrpc:"peppol_state,omitempty"`
	ResponseCode       *Selection  `xmlrpc:"response_code,omitempty"`
	WriteDate          *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountPeppolResponses represents array of account.peppol.response model.
type AccountPeppolResponses []AccountPeppolResponse

// AccountPeppolResponseModel is the odoo model name.
const AccountPeppolResponseModel = "account.peppol.response"

// Many2One convert AccountPeppolResponse to *Many2One.
func (apr *AccountPeppolResponse) Many2One() *Many2One {
	return NewMany2One(apr.Id.Get(), "")
}

// CreateAccountPeppolResponse creates a new account.peppol.response model and returns its id.
func (c *Client) CreateAccountPeppolResponse(apr *AccountPeppolResponse) (int64, error) {
	ids, err := c.CreateAccountPeppolResponses([]*AccountPeppolResponse{apr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPeppolResponse creates a new account.peppol.response model and returns its id.
func (c *Client) CreateAccountPeppolResponses(aprs []*AccountPeppolResponse) ([]int64, error) {
	var vv []interface{}
	for _, v := range aprs {
		vv = append(vv, v)
	}
	return c.Create(AccountPeppolResponseModel, vv, nil)
}

// UpdateAccountPeppolResponse updates an existing account.peppol.response record.
func (c *Client) UpdateAccountPeppolResponse(apr *AccountPeppolResponse) error {
	return c.UpdateAccountPeppolResponses([]int64{apr.Id.Get()}, apr)
}

// UpdateAccountPeppolResponses updates existing account.peppol.response records.
// All records (represented by ids) will be updated by apr values.
func (c *Client) UpdateAccountPeppolResponses(ids []int64, apr *AccountPeppolResponse) error {
	return c.Update(AccountPeppolResponseModel, ids, apr, nil)
}

// DeleteAccountPeppolResponse deletes an existing account.peppol.response record.
func (c *Client) DeleteAccountPeppolResponse(id int64) error {
	return c.DeleteAccountPeppolResponses([]int64{id})
}

// DeleteAccountPeppolResponses deletes existing account.peppol.response records.
func (c *Client) DeleteAccountPeppolResponses(ids []int64) error {
	return c.Delete(AccountPeppolResponseModel, ids)
}

// GetAccountPeppolResponse gets account.peppol.response existing record.
func (c *Client) GetAccountPeppolResponse(id int64) (*AccountPeppolResponse, error) {
	aprs, err := c.GetAccountPeppolResponses([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aprs)[0]), nil
}

// GetAccountPeppolResponses gets account.peppol.response existing records.
func (c *Client) GetAccountPeppolResponses(ids []int64) (*AccountPeppolResponses, error) {
	aprs := &AccountPeppolResponses{}
	if err := c.Read(AccountPeppolResponseModel, ids, nil, aprs); err != nil {
		return nil, err
	}
	return aprs, nil
}

// FindAccountPeppolResponse finds account.peppol.response record by querying it with criteria.
func (c *Client) FindAccountPeppolResponse(criteria *Criteria) (*AccountPeppolResponse, error) {
	aprs := &AccountPeppolResponses{}
	if err := c.SearchRead(AccountPeppolResponseModel, criteria, NewOptions().Limit(1), aprs); err != nil {
		return nil, err
	}
	return &((*aprs)[0]), nil
}

// FindAccountPeppolResponses finds account.peppol.response records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPeppolResponses(criteria *Criteria, options *Options) (*AccountPeppolResponses, error) {
	aprs := &AccountPeppolResponses{}
	if err := c.SearchRead(AccountPeppolResponseModel, criteria, options, aprs); err != nil {
		return nil, err
	}
	return aprs, nil
}

// FindAccountPeppolResponseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPeppolResponseIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPeppolResponseModel, criteria, options)
}

// FindAccountPeppolResponseId finds record id by querying it with criteria.
func (c *Client) FindAccountPeppolResponseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPeppolResponseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
