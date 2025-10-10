package odoo

// DigestDigest represents digest.digest model.
type DigestDigest struct {
	AvailableFields               *String    `xmlrpc:"available_fields,omitempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omitempty"`
	Id                            *Int       `xmlrpc:"id,omitempty"`
	IsSubscribed                  *Bool      `xmlrpc:"is_subscribed,omitempty"`
	KpiAccountTotalRevenue        *Bool      `xmlrpc:"kpi_account_total_revenue,omitempty"`
	KpiAccountTotalRevenueValue   *Float     `xmlrpc:"kpi_account_total_revenue_value,omitempty"`
	KpiAllSaleTotal               *Bool      `xmlrpc:"kpi_all_sale_total,omitempty"`
	KpiAllSaleTotalValue          *Float     `xmlrpc:"kpi_all_sale_total_value,omitempty"`
	KpiCrmLeadCreated             *Bool      `xmlrpc:"kpi_crm_lead_created,omitempty"`
	KpiCrmLeadCreatedValue        *Int       `xmlrpc:"kpi_crm_lead_created_value,omitempty"`
	KpiCrmOpportunitiesWon        *Bool      `xmlrpc:"kpi_crm_opportunities_won,omitempty"`
	KpiCrmOpportunitiesWonValue   *Int       `xmlrpc:"kpi_crm_opportunities_won_value,omitempty"`
	KpiLivechatConversations      *Bool      `xmlrpc:"kpi_livechat_conversations,omitempty"`
	KpiLivechatConversationsValue *Int       `xmlrpc:"kpi_livechat_conversations_value,omitempty"`
	KpiLivechatRating             *Bool      `xmlrpc:"kpi_livechat_rating,omitempty"`
	KpiLivechatRatingValue        *Float     `xmlrpc:"kpi_livechat_rating_value,omitempty"`
	KpiLivechatResponse           *Bool      `xmlrpc:"kpi_livechat_response,omitempty"`
	KpiLivechatResponseValue      *Float     `xmlrpc:"kpi_livechat_response_value,omitempty"`
	KpiMailMessageTotal           *Bool      `xmlrpc:"kpi_mail_message_total,omitempty"`
	KpiMailMessageTotalValue      *Int       `xmlrpc:"kpi_mail_message_total_value,omitempty"`
	KpiProjectTaskOpened          *Bool      `xmlrpc:"kpi_project_task_opened,omitempty"`
	KpiProjectTaskOpenedValue     *Int       `xmlrpc:"kpi_project_task_opened_value,omitempty"`
	KpiResUsersConnected          *Bool      `xmlrpc:"kpi_res_users_connected,omitempty"`
	KpiResUsersConnectedValue     *Int       `xmlrpc:"kpi_res_users_connected_value,omitempty"`
	Name                          *String    `xmlrpc:"name,omitempty"`
	NextRunDate                   *Time      `xmlrpc:"next_run_date,omitempty"`
	Periodicity                   *Selection `xmlrpc:"periodicity,omitempty"`
	State                         *Selection `xmlrpc:"state,omitempty"`
	UserIds                       *Relation  `xmlrpc:"user_ids,omitempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// DigestDigests represents array of digest.digest model.
type DigestDigests []DigestDigest

// DigestDigestModel is the odoo model name.
const DigestDigestModel = "digest.digest"

// Many2One convert DigestDigest to *Many2One.
func (dd *DigestDigest) Many2One() *Many2One {
	return NewMany2One(dd.Id.Get(), "")
}

// CreateDigestDigest creates a new digest.digest model and returns its id.
func (c *Client) CreateDigestDigest(dd *DigestDigest) (int64, error) {
	ids, err := c.CreateDigestDigests([]*DigestDigest{dd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDigestDigest creates a new digest.digest model and returns its id.
func (c *Client) CreateDigestDigests(dds []*DigestDigest) ([]int64, error) {
	var vv []interface{}
	for _, v := range dds {
		vv = append(vv, v)
	}
	return c.Create(DigestDigestModel, vv, nil)
}

// UpdateDigestDigest updates an existing digest.digest record.
func (c *Client) UpdateDigestDigest(dd *DigestDigest) error {
	return c.UpdateDigestDigests([]int64{dd.Id.Get()}, dd)
}

// UpdateDigestDigests updates existing digest.digest records.
// All records (represented by ids) will be updated by dd values.
func (c *Client) UpdateDigestDigests(ids []int64, dd *DigestDigest) error {
	return c.Update(DigestDigestModel, ids, dd, nil)
}

// DeleteDigestDigest deletes an existing digest.digest record.
func (c *Client) DeleteDigestDigest(id int64) error {
	return c.DeleteDigestDigests([]int64{id})
}

// DeleteDigestDigests deletes existing digest.digest records.
func (c *Client) DeleteDigestDigests(ids []int64) error {
	return c.Delete(DigestDigestModel, ids)
}

// GetDigestDigest gets digest.digest existing record.
func (c *Client) GetDigestDigest(id int64) (*DigestDigest, error) {
	dds, err := c.GetDigestDigests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dds)[0]), nil
}

// GetDigestDigests gets digest.digest existing records.
func (c *Client) GetDigestDigests(ids []int64) (*DigestDigests, error) {
	dds := &DigestDigests{}
	if err := c.Read(DigestDigestModel, ids, nil, dds); err != nil {
		return nil, err
	}
	return dds, nil
}

// FindDigestDigest finds digest.digest record by querying it with criteria.
func (c *Client) FindDigestDigest(criteria *Criteria) (*DigestDigest, error) {
	dds := &DigestDigests{}
	if err := c.SearchRead(DigestDigestModel, criteria, NewOptions().Limit(1), dds); err != nil {
		return nil, err
	}
	return &((*dds)[0]), nil
}

// FindDigestDigests finds digest.digest records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDigestDigests(criteria *Criteria, options *Options) (*DigestDigests, error) {
	dds := &DigestDigests{}
	if err := c.SearchRead(DigestDigestModel, criteria, options, dds); err != nil {
		return nil, err
	}
	return dds, nil
}

// FindDigestDigestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDigestDigestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DigestDigestModel, criteria, options)
}

// FindDigestDigestId finds record id by querying it with criteria.
func (c *Client) FindDigestDigestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DigestDigestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
