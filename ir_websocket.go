package odoo

// IrWebsocket represents ir.websocket model.
type IrWebsocket struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrWebsockets represents array of ir.websocket model.
type IrWebsockets []IrWebsocket

// IrWebsocketModel is the odoo model name.
const IrWebsocketModel = "ir.websocket"

// Many2One convert IrWebsocket to *Many2One.
func (iw *IrWebsocket) Many2One() *Many2One {
	return NewMany2One(iw.Id.Get(), "")
}

// CreateIrWebsocket creates a new ir.websocket model and returns its id.
func (c *Client) CreateIrWebsocket(iw *IrWebsocket) (int64, error) {
	ids, err := c.CreateIrWebsockets([]*IrWebsocket{iw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrWebsocket creates a new ir.websocket model and returns its id.
func (c *Client) CreateIrWebsockets(iws []*IrWebsocket) ([]int64, error) {
	var vv []interface{}
	for _, v := range iws {
		vv = append(vv, v)
	}
	return c.Create(IrWebsocketModel, vv, nil)
}

// UpdateIrWebsocket updates an existing ir.websocket record.
func (c *Client) UpdateIrWebsocket(iw *IrWebsocket) error {
	return c.UpdateIrWebsockets([]int64{iw.Id.Get()}, iw)
}

// UpdateIrWebsockets updates existing ir.websocket records.
// All records (represented by ids) will be updated by iw values.
func (c *Client) UpdateIrWebsockets(ids []int64, iw *IrWebsocket) error {
	return c.Update(IrWebsocketModel, ids, iw, nil)
}

// DeleteIrWebsocket deletes an existing ir.websocket record.
func (c *Client) DeleteIrWebsocket(id int64) error {
	return c.DeleteIrWebsockets([]int64{id})
}

// DeleteIrWebsockets deletes existing ir.websocket records.
func (c *Client) DeleteIrWebsockets(ids []int64) error {
	return c.Delete(IrWebsocketModel, ids)
}

// GetIrWebsocket gets ir.websocket existing record.
func (c *Client) GetIrWebsocket(id int64) (*IrWebsocket, error) {
	iws, err := c.GetIrWebsockets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iws)[0]), nil
}

// GetIrWebsockets gets ir.websocket existing records.
func (c *Client) GetIrWebsockets(ids []int64) (*IrWebsockets, error) {
	iws := &IrWebsockets{}
	if err := c.Read(IrWebsocketModel, ids, nil, iws); err != nil {
		return nil, err
	}
	return iws, nil
}

// FindIrWebsocket finds ir.websocket record by querying it with criteria.
func (c *Client) FindIrWebsocket(criteria *Criteria) (*IrWebsocket, error) {
	iws := &IrWebsockets{}
	if err := c.SearchRead(IrWebsocketModel, criteria, NewOptions().Limit(1), iws); err != nil {
		return nil, err
	}
	return &((*iws)[0]), nil
}

// FindIrWebsockets finds ir.websocket records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrWebsockets(criteria *Criteria, options *Options) (*IrWebsockets, error) {
	iws := &IrWebsockets{}
	if err := c.SearchRead(IrWebsocketModel, criteria, options, iws); err != nil {
		return nil, err
	}
	return iws, nil
}

// FindIrWebsocketIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrWebsocketIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrWebsocketModel, criteria, options)
}

// FindIrWebsocketId finds record id by querying it with criteria.
func (c *Client) FindIrWebsocketId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrWebsocketModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
