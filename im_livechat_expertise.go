package odoo

// ImLivechatExpertise represents im_livechat.expertise model.
type ImLivechatExpertise struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	UserIds     *Relation `xmlrpc:"user_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ImLivechatExpertises represents array of im_livechat.expertise model.
type ImLivechatExpertises []ImLivechatExpertise

// ImLivechatExpertiseModel is the odoo model name.
const ImLivechatExpertiseModel = "im_livechat.expertise"

// Many2One convert ImLivechatExpertise to *Many2One.
func (ie *ImLivechatExpertise) Many2One() *Many2One {
	return NewMany2One(ie.Id.Get(), "")
}

// CreateImLivechatExpertise creates a new im_livechat.expertise model and returns its id.
func (c *Client) CreateImLivechatExpertise(ie *ImLivechatExpertise) (int64, error) {
	ids, err := c.CreateImLivechatExpertises([]*ImLivechatExpertise{ie})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateImLivechatExpertise creates a new im_livechat.expertise model and returns its id.
func (c *Client) CreateImLivechatExpertises(ies []*ImLivechatExpertise) ([]int64, error) {
	var vv []interface{}
	for _, v := range ies {
		vv = append(vv, v)
	}
	return c.Create(ImLivechatExpertiseModel, vv, nil)
}

// UpdateImLivechatExpertise updates an existing im_livechat.expertise record.
func (c *Client) UpdateImLivechatExpertise(ie *ImLivechatExpertise) error {
	return c.UpdateImLivechatExpertises([]int64{ie.Id.Get()}, ie)
}

// UpdateImLivechatExpertises updates existing im_livechat.expertise records.
// All records (represented by ids) will be updated by ie values.
func (c *Client) UpdateImLivechatExpertises(ids []int64, ie *ImLivechatExpertise) error {
	return c.Update(ImLivechatExpertiseModel, ids, ie, nil)
}

// DeleteImLivechatExpertise deletes an existing im_livechat.expertise record.
func (c *Client) DeleteImLivechatExpertise(id int64) error {
	return c.DeleteImLivechatExpertises([]int64{id})
}

// DeleteImLivechatExpertises deletes existing im_livechat.expertise records.
func (c *Client) DeleteImLivechatExpertises(ids []int64) error {
	return c.Delete(ImLivechatExpertiseModel, ids)
}

// GetImLivechatExpertise gets im_livechat.expertise existing record.
func (c *Client) GetImLivechatExpertise(id int64) (*ImLivechatExpertise, error) {
	ies, err := c.GetImLivechatExpertises([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ies)[0]), nil
}

// GetImLivechatExpertises gets im_livechat.expertise existing records.
func (c *Client) GetImLivechatExpertises(ids []int64) (*ImLivechatExpertises, error) {
	ies := &ImLivechatExpertises{}
	if err := c.Read(ImLivechatExpertiseModel, ids, nil, ies); err != nil {
		return nil, err
	}
	return ies, nil
}

// FindImLivechatExpertise finds im_livechat.expertise record by querying it with criteria.
func (c *Client) FindImLivechatExpertise(criteria *Criteria) (*ImLivechatExpertise, error) {
	ies := &ImLivechatExpertises{}
	if err := c.SearchRead(ImLivechatExpertiseModel, criteria, NewOptions().Limit(1), ies); err != nil {
		return nil, err
	}
	return &((*ies)[0]), nil
}

// FindImLivechatExpertises finds im_livechat.expertise records by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatExpertises(criteria *Criteria, options *Options) (*ImLivechatExpertises, error) {
	ies := &ImLivechatExpertises{}
	if err := c.SearchRead(ImLivechatExpertiseModel, criteria, options, ies); err != nil {
		return nil, err
	}
	return ies, nil
}

// FindImLivechatExpertiseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatExpertiseIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ImLivechatExpertiseModel, criteria, options)
}

// FindImLivechatExpertiseId finds record id by querying it with criteria.
func (c *Client) FindImLivechatExpertiseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ImLivechatExpertiseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
