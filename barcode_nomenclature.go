package odoo

import (
	"fmt"
)

// BarcodeNomenclature represents barcode.nomenclature model.
type BarcodeNomenclature struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	RuleIds     *Relation  `xmlrpc:"rule_ids,omitempty"`
	UpcEanConv  *Selection `xmlrpc:"upc_ean_conv,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// BarcodeNomenclatures represents array of barcode.nomenclature model.
type BarcodeNomenclatures []BarcodeNomenclature

// BarcodeNomenclatureModel is the odoo model name.
const BarcodeNomenclatureModel = "barcode.nomenclature"

// Many2One convert BarcodeNomenclature to *Many2One.
func (bn *BarcodeNomenclature) Many2One() *Many2One {
	return NewMany2One(bn.Id.Get(), "")
}

// CreateBarcodeNomenclature creates a new barcode.nomenclature model and returns its id.
func (c *Client) CreateBarcodeNomenclature(bn *BarcodeNomenclature) (int64, error) {
	return c.Create(BarcodeNomenclatureModel, bn)
}

// UpdateBarcodeNomenclature updates an existing barcode.nomenclature record.
func (c *Client) UpdateBarcodeNomenclature(bn *BarcodeNomenclature) error {
	return c.UpdateBarcodeNomenclatures([]int64{bn.Id.Get()}, bn)
}

// UpdateBarcodeNomenclatures updates existing barcode.nomenclature records.
// All records (represented by ids) will be updated by bn values.
func (c *Client) UpdateBarcodeNomenclatures(ids []int64, bn *BarcodeNomenclature) error {
	return c.Update(BarcodeNomenclatureModel, ids, bn)
}

// DeleteBarcodeNomenclature deletes an existing barcode.nomenclature record.
func (c *Client) DeleteBarcodeNomenclature(id int64) error {
	return c.DeleteBarcodeNomenclatures([]int64{id})
}

// DeleteBarcodeNomenclatures deletes existing barcode.nomenclature records.
func (c *Client) DeleteBarcodeNomenclatures(ids []int64) error {
	return c.Delete(BarcodeNomenclatureModel, ids)
}

// GetBarcodeNomenclature gets barcode.nomenclature existing record.
func (c *Client) GetBarcodeNomenclature(id int64) (*BarcodeNomenclature, error) {
	bns, err := c.GetBarcodeNomenclatures([]int64{id})
	if err != nil {
		return nil, err
	}
	if bns != nil && len(*bns) > 0 {
		return &((*bns)[0]), nil
	}
	return nil, fmt.Errorf("id %v of barcode.nomenclature not found", id)
}

// GetBarcodeNomenclatures gets barcode.nomenclature existing records.
func (c *Client) GetBarcodeNomenclatures(ids []int64) (*BarcodeNomenclatures, error) {
	bns := &BarcodeNomenclatures{}
	if err := c.Read(BarcodeNomenclatureModel, ids, nil, bns); err != nil {
		return nil, err
	}
	return bns, nil
}

// FindBarcodeNomenclature finds barcode.nomenclature record by querying it with criteria.
func (c *Client) FindBarcodeNomenclature(criteria *Criteria) (*BarcodeNomenclature, error) {
	bns := &BarcodeNomenclatures{}
	if err := c.SearchRead(BarcodeNomenclatureModel, criteria, NewOptions().Limit(1), bns); err != nil {
		return nil, err
	}
	if bns != nil && len(*bns) > 0 {
		return &((*bns)[0]), nil
	}
	return nil, fmt.Errorf("no barcode.nomenclature was found with criteria %v", criteria)
}

// FindBarcodeNomenclatures finds barcode.nomenclature records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBarcodeNomenclatures(criteria *Criteria, options *Options) (*BarcodeNomenclatures, error) {
	bns := &BarcodeNomenclatures{}
	if err := c.SearchRead(BarcodeNomenclatureModel, criteria, options, bns); err != nil {
		return nil, err
	}
	return bns, nil
}

// FindBarcodeNomenclatureIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBarcodeNomenclatureIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BarcodeNomenclatureModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBarcodeNomenclatureId finds record id by querying it with criteria.
func (c *Client) FindBarcodeNomenclatureId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BarcodeNomenclatureModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no barcode.nomenclature was found with criteria %v and options %v", criteria, options)
}
