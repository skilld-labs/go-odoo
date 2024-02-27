package odoo

// ValidateAccountMove represents validate.account.move model.
type ValidateAccountMove struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ValidateAccountMoves represents array of validate.account.move model.
type ValidateAccountMoves []ValidateAccountMove

// ValidateAccountMoveModel is the odoo model name.
const ValidateAccountMoveModel = "validate.account.move"

// Many2One convert ValidateAccountMove to *Many2One.
func (vam *ValidateAccountMove) Many2One() *Many2One {
	return NewMany2One(vam.Id.Get(), "")
}

// CreateValidateAccountMove creates a new validate.account.move model and returns its id.
func (c *Client) CreateValidateAccountMove(vam *ValidateAccountMove) (int64, error) {
	ids, err := c.CreateValidateAccountMoves([]*ValidateAccountMove{vam})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateValidateAccountMove creates a new validate.account.move model and returns its id.
func (c *Client) CreateValidateAccountMoves(vams []*ValidateAccountMove) ([]int64, error) {
	var vv []interface{}
	for _, v := range vams {
		vv = append(vv, v)
	}
	return c.Create(ValidateAccountMoveModel, vv, nil)
}

// UpdateValidateAccountMove updates an existing validate.account.move record.
func (c *Client) UpdateValidateAccountMove(vam *ValidateAccountMove) error {
	return c.UpdateValidateAccountMoves([]int64{vam.Id.Get()}, vam)
}

// UpdateValidateAccountMoves updates existing validate.account.move records.
// All records (represented by ids) will be updated by vam values.
func (c *Client) UpdateValidateAccountMoves(ids []int64, vam *ValidateAccountMove) error {
	return c.Update(ValidateAccountMoveModel, ids, vam, nil)
}

// DeleteValidateAccountMove deletes an existing validate.account.move record.
func (c *Client) DeleteValidateAccountMove(id int64) error {
	return c.DeleteValidateAccountMoves([]int64{id})
}

// DeleteValidateAccountMoves deletes existing validate.account.move records.
func (c *Client) DeleteValidateAccountMoves(ids []int64) error {
	return c.Delete(ValidateAccountMoveModel, ids)
}

// GetValidateAccountMove gets validate.account.move existing record.
func (c *Client) GetValidateAccountMove(id int64) (*ValidateAccountMove, error) {
	vams, err := c.GetValidateAccountMoves([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*vams)[0]), nil
}

// GetValidateAccountMoves gets validate.account.move existing records.
func (c *Client) GetValidateAccountMoves(ids []int64) (*ValidateAccountMoves, error) {
	vams := &ValidateAccountMoves{}
	if err := c.Read(ValidateAccountMoveModel, ids, nil, vams); err != nil {
		return nil, err
	}
	return vams, nil
}

// FindValidateAccountMove finds validate.account.move record by querying it with criteria.
func (c *Client) FindValidateAccountMove(criteria *Criteria) (*ValidateAccountMove, error) {
	vams := &ValidateAccountMoves{}
	if err := c.SearchRead(ValidateAccountMoveModel, criteria, NewOptions().Limit(1), vams); err != nil {
		return nil, err
	}
	return &((*vams)[0]), nil
}

// FindValidateAccountMoves finds validate.account.move records by querying it
// and filtering it with criteria and options.
func (c *Client) FindValidateAccountMoves(criteria *Criteria, options *Options) (*ValidateAccountMoves, error) {
	vams := &ValidateAccountMoves{}
	if err := c.SearchRead(ValidateAccountMoveModel, criteria, options, vams); err != nil {
		return nil, err
	}
	return vams, nil
}

// FindValidateAccountMoveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindValidateAccountMoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ValidateAccountMoveModel, criteria, options)
}

// FindValidateAccountMoveId finds record id by querying it with criteria.
func (c *Client) FindValidateAccountMoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ValidateAccountMoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
