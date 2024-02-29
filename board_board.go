package odoo

// BoardBoard represents board.board model.
type BoardBoard struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// BoardBoards represents array of board.board model.
type BoardBoards []BoardBoard

// BoardBoardModel is the odoo model name.
const BoardBoardModel = "board.board"

// Many2One convert BoardBoard to *Many2One.
func (bb *BoardBoard) Many2One() *Many2One {
	return NewMany2One(bb.Id.Get(), "")
}

// CreateBoardBoard creates a new board.board model and returns its id.
func (c *Client) CreateBoardBoard(bb *BoardBoard) (int64, error) {
	ids, err := c.CreateBoardBoards([]*BoardBoard{bb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBoardBoards creates a new board.board model and returns its id.
func (c *Client) CreateBoardBoards(bbs []*BoardBoard) ([]int64, error) {
	var vv []interface{}
	for _, v := range bbs {
		vv = append(vv, v)
	}
	return c.Create(BoardBoardModel, vv, nil)
}

// UpdateBoardBoard updates an existing board.board record.
func (c *Client) UpdateBoardBoard(bb *BoardBoard) error {
	return c.UpdateBoardBoards([]int64{bb.Id.Get()}, bb)
}

// UpdateBoardBoards updates existing board.board records.
// All records (represented by ids) will be updated by bb values.
func (c *Client) UpdateBoardBoards(ids []int64, bb *BoardBoard) error {
	return c.Update(BoardBoardModel, ids, bb, nil)
}

// DeleteBoardBoard deletes an existing board.board record.
func (c *Client) DeleteBoardBoard(id int64) error {
	return c.DeleteBoardBoards([]int64{id})
}

// DeleteBoardBoards deletes existing board.board records.
func (c *Client) DeleteBoardBoards(ids []int64) error {
	return c.Delete(BoardBoardModel, ids)
}

// GetBoardBoard gets board.board existing record.
func (c *Client) GetBoardBoard(id int64) (*BoardBoard, error) {
	bbs, err := c.GetBoardBoards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bbs)[0]), nil
}

// GetBoardBoards gets board.board existing records.
func (c *Client) GetBoardBoards(ids []int64) (*BoardBoards, error) {
	bbs := &BoardBoards{}
	if err := c.Read(BoardBoardModel, ids, nil, bbs); err != nil {
		return nil, err
	}
	return bbs, nil
}

// FindBoardBoard finds board.board record by querying it with criteria.
func (c *Client) FindBoardBoard(criteria *Criteria) (*BoardBoard, error) {
	bbs := &BoardBoards{}
	if err := c.SearchRead(BoardBoardModel, criteria, NewOptions().Limit(1), bbs); err != nil {
		return nil, err
	}
	return &((*bbs)[0]), nil
}

// FindBoardBoards finds board.board records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBoardBoards(criteria *Criteria, options *Options) (*BoardBoards, error) {
	bbs := &BoardBoards{}
	if err := c.SearchRead(BoardBoardModel, criteria, options, bbs); err != nil {
		return nil, err
	}
	return bbs, nil
}

// FindBoardBoardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBoardBoardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BoardBoardModel, criteria, options)
}

// FindBoardBoardId finds record id by querying it with criteria.
func (c *Client) FindBoardBoardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BoardBoardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
