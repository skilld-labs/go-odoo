package odoo

// DeliveryCarrier represents delivery.carrier model.
type DeliveryCarrier struct {
	Active                    *Bool      `xmlrpc:"active,omitempty"`
	Amount                    *Float     `xmlrpc:"amount,omitempty"`
	CanGenerateReturn         *Bool      `xmlrpc:"can_generate_return,omitempty"`
	CarrierDescription        *String    `xmlrpc:"carrier_description,omitempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryIds                *Relation  `xmlrpc:"country_ids,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                *Many2One  `xmlrpc:"currency_id,omitempty"`
	DebugLogging              *Bool      `xmlrpc:"debug_logging,omitempty"`
	DeliveryType              *Selection `xmlrpc:"delivery_type,omitempty"`
	DhlSiteID                 *String    `xmlrpc:"dhl_SiteID,omitempty"`
	DhlAccountNumber          *String    `xmlrpc:"dhl_account_number,omitempty"`
	DhlCustomDataRequest      *String    `xmlrpc:"dhl_custom_data_request,omitempty"`
	DhlDefaultPackageTypeId   *Many2One  `xmlrpc:"dhl_default_package_type_id,omitempty"`
	DhlDutiable               *Bool      `xmlrpc:"dhl_dutiable,omitempty"`
	DhlDutyPayment            *Selection `xmlrpc:"dhl_duty_payment,omitempty"`
	DhlLabelImageFormat       *Selection `xmlrpc:"dhl_label_image_format,omitempty"`
	DhlLabelTemplate          *Selection `xmlrpc:"dhl_label_template,omitempty"`
	DhlPackageDimensionUnit   *Selection `xmlrpc:"dhl_package_dimension_unit,omitempty"`
	DhlPackageWeightUnit      *Selection `xmlrpc:"dhl_package_weight_unit,omitempty"`
	DhlPassword               *String    `xmlrpc:"dhl_password,omitempty"`
	DhlProductCode            *Selection `xmlrpc:"dhl_product_code,omitempty"`
	DhlRegionCode             *Selection `xmlrpc:"dhl_region_code,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	FixedMargin               *Float     `xmlrpc:"fixed_margin,omitempty"`
	FixedPrice                *Float     `xmlrpc:"fixed_price,omitempty"`
	FreeOver                  *Bool      `xmlrpc:"free_over,omitempty"`
	GetReturnLabelFromPortal  *Bool      `xmlrpc:"get_return_label_from_portal,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	IntegrationLevel          *Selection `xmlrpc:"integration_level,omitempty"`
	InvoicePolicy             *Selection `xmlrpc:"invoice_policy,omitempty"`
	Margin                    *Float     `xmlrpc:"margin,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	PriceRuleIds              *Relation  `xmlrpc:"price_rule_ids,omitempty"`
	ProdEnvironment           *Bool      `xmlrpc:"prod_environment,omitempty"`
	ProductId                 *Many2One  `xmlrpc:"product_id,omitempty"`
	ReturnLabelOnDelivery     *Bool      `xmlrpc:"return_label_on_delivery,omitempty"`
	RouteIds                  *Relation  `xmlrpc:"route_ids,omitempty"`
	Sequence                  *Int       `xmlrpc:"sequence,omitempty"`
	ShippingInsurance         *Int       `xmlrpc:"shipping_insurance,omitempty"`
	StateIds                  *Relation  `xmlrpc:"state_ids,omitempty"`
	SupportsShippingInsurance *Bool      `xmlrpc:"supports_shipping_insurance,omitempty"`
	UpsAccessToken            *String    `xmlrpc:"ups_access_token,omitempty"`
	UpsBillMyAccount          *Bool      `xmlrpc:"ups_bill_my_account,omitempty"`
	UpsClientId               *String    `xmlrpc:"ups_client_id,omitempty"`
	UpsClientSecret           *String    `xmlrpc:"ups_client_secret,omitempty"`
	UpsCod                    *Bool      `xmlrpc:"ups_cod,omitempty"`
	UpsCodFundsCode           *Selection `xmlrpc:"ups_cod_funds_code,omitempty"`
	UpsDefaultPackagingId     *Many2One  `xmlrpc:"ups_default_packaging_id,omitempty"`
	UpsDefaultServiceType     *Selection `xmlrpc:"ups_default_service_type,omitempty"`
	UpsDutyPayment            *Selection `xmlrpc:"ups_duty_payment,omitempty"`
	UpsLabelFileType          *Selection `xmlrpc:"ups_label_file_type,omitempty"`
	UpsPackageDimensionUnit   *Selection `xmlrpc:"ups_package_dimension_unit,omitempty"`
	UpsPackageWeightUnit      *Selection `xmlrpc:"ups_package_weight_unit,omitempty"`
	UpsSaturdayDelivery       *Bool      `xmlrpc:"ups_saturday_delivery,omitempty"`
	UpsShipperNumber          *String    `xmlrpc:"ups_shipper_number,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
	ZipPrefixIds              *Relation  `xmlrpc:"zip_prefix_ids,omitempty"`
}

// DeliveryCarriers represents array of delivery.carrier model.
type DeliveryCarriers []DeliveryCarrier

// DeliveryCarrierModel is the odoo model name.
const DeliveryCarrierModel = "delivery.carrier"

// Many2One convert DeliveryCarrier to *Many2One.
func (dc *DeliveryCarrier) Many2One() *Many2One {
	return NewMany2One(dc.Id.Get(), "")
}

// CreateDeliveryCarrier creates a new delivery.carrier model and returns its id.
func (c *Client) CreateDeliveryCarrier(dc *DeliveryCarrier) (int64, error) {
	ids, err := c.CreateDeliveryCarriers([]*DeliveryCarrier{dc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDeliveryCarrier creates a new delivery.carrier model and returns its id.
func (c *Client) CreateDeliveryCarriers(dcs []*DeliveryCarrier) ([]int64, error) {
	var vv []interface{}
	for _, v := range dcs {
		vv = append(vv, v)
	}
	return c.Create(DeliveryCarrierModel, vv, nil)
}

// UpdateDeliveryCarrier updates an existing delivery.carrier record.
func (c *Client) UpdateDeliveryCarrier(dc *DeliveryCarrier) error {
	return c.UpdateDeliveryCarriers([]int64{dc.Id.Get()}, dc)
}

// UpdateDeliveryCarriers updates existing delivery.carrier records.
// All records (represented by ids) will be updated by dc values.
func (c *Client) UpdateDeliveryCarriers(ids []int64, dc *DeliveryCarrier) error {
	return c.Update(DeliveryCarrierModel, ids, dc, nil)
}

// DeleteDeliveryCarrier deletes an existing delivery.carrier record.
func (c *Client) DeleteDeliveryCarrier(id int64) error {
	return c.DeleteDeliveryCarriers([]int64{id})
}

// DeleteDeliveryCarriers deletes existing delivery.carrier records.
func (c *Client) DeleteDeliveryCarriers(ids []int64) error {
	return c.Delete(DeliveryCarrierModel, ids)
}

// GetDeliveryCarrier gets delivery.carrier existing record.
func (c *Client) GetDeliveryCarrier(id int64) (*DeliveryCarrier, error) {
	dcs, err := c.GetDeliveryCarriers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dcs)[0]), nil
}

// GetDeliveryCarriers gets delivery.carrier existing records.
func (c *Client) GetDeliveryCarriers(ids []int64) (*DeliveryCarriers, error) {
	dcs := &DeliveryCarriers{}
	if err := c.Read(DeliveryCarrierModel, ids, nil, dcs); err != nil {
		return nil, err
	}
	return dcs, nil
}

// FindDeliveryCarrier finds delivery.carrier record by querying it with criteria.
func (c *Client) FindDeliveryCarrier(criteria *Criteria) (*DeliveryCarrier, error) {
	dcs := &DeliveryCarriers{}
	if err := c.SearchRead(DeliveryCarrierModel, criteria, NewOptions().Limit(1), dcs); err != nil {
		return nil, err
	}
	return &((*dcs)[0]), nil
}

// FindDeliveryCarriers finds delivery.carrier records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDeliveryCarriers(criteria *Criteria, options *Options) (*DeliveryCarriers, error) {
	dcs := &DeliveryCarriers{}
	if err := c.SearchRead(DeliveryCarrierModel, criteria, options, dcs); err != nil {
		return nil, err
	}
	return dcs, nil
}

// FindDeliveryCarrierIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDeliveryCarrierIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DeliveryCarrierModel, criteria, options)
}

// FindDeliveryCarrierId finds record id by querying it with criteria.
func (c *Client) FindDeliveryCarrierId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DeliveryCarrierModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
