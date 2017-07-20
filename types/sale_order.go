package types 

type NilableType interface {
	Type() interface{}
}

type Type interface {
	NilableType() interface{}
}

type SaleOrder struct {
	__LastUpdate      string  `xmlrpc:"__last_update"`
	AmountTax         float64 `xmlrpc:"amount_tax"`
	AmountTotal       float64 `xmlrpc:"amount_total"`
	AmountUntaxed     float64 `xmlrpc:"amount_untaxed"`
	AutosalesConfigId struct {
		Id   int64
		Name string
	} `xmlrpc:"autosales_config_id"`
	CampaignId struct {
		Id   int64
		Name string
	} `xmlrpc:"campaign_id"`
	ClientOrderRef string `xmlrpc:"client_order_ref"`
	CompanyId      struct {
		Id   int64
		Name string
	} `xmlrpc:"company_id"`
	ConfirmationDate string `xmlrpc:"confirmation_date"`
	CreateDate       string `xmlrpc:"create_date"`
	CreateUid        struct {
		Id   int64
		Name string
	} `xmlrpc:"create_uid"`
	CurrencyId struct {
		Id   int64
		Name string
	} `xmlrpc:"currency_id"`
	DateOrder        string `xmlrpc:"date_order"`
	DeliveryCount    int64  `xmlrpc:"delivery_count"`
	DisplayName      string `xmlrpc:"display_name"`
	FiscalPositionId struct {
		Id   int64
		Name string
	} `xmlrpc:"fiscal_position_id"`
	Id       int64 `xmlrpc:"id"`
	Incoterm struct {
		Id   int64
		Name string
	} `xmlrpc:"incoterm"`
	InvoiceCount  int64   `xmlrpc:"invoice_count"`
	InvoiceIds    []int64 `xmlrpc:"invoice_ids"`
	InvoiceStatus string  `xmlrpc:"invoice_status"`
	MediumId      struct {
		Id   int64
		Name string
	} `xmlrpc:"medium_id"`
	MessageChannelIds        []int64 `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       []int64 `xmlrpc:"message_follower_ids"`
	MessageIds               []int64 `xmlrpc:"message_ids"`
	MessageIsFollower        bool    `xmlrpc:"message_is_follower"`
	MessageLastPost          string  `xmlrpc:"message_last_post"`
	MessageNeedaction        bool    `xmlrpc:"message_needaction"`
	MessageNeedactionCounter int64   `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        []int64 `xmlrpc:"message_partner_ids"`
	MessageUnread            bool    `xmlrpc:"message_unread"`
	MessageUnreadCounter     int64   `xmlrpc:"message_unread_counter"`
	Name                     string  `xmlrpc:"name"`
	Note                     string  `xmlrpc:"note"`
	OpportunityId            struct {
		Id   int64
		Name string
	} `xmlrpc:"opportunity_id"`
	OrderLine []int64 `xmlrpc:"order_line"`
	Origin    string  `xmlrpc:"origin"`
	PartnerId struct {
		Id   int64
		Name string
	} `xmlrpc:"partner_id"`
	PartnerInvoiceId struct {
		Id   int64
		Name string
	} `xmlrpc:"partner_invoice_id"`
	PartnerShippingId struct {
		Id   int64
		Name string
	} `xmlrpc:"partner_shipping_id"`
	PaymentTermId struct {
		Id   int64
		Name string
	} `xmlrpc:"payment_term_id"`
	PickingIds    []int64 `xmlrpc:"picking_ids"`
	PickingPolicy string  `xmlrpc:"picking_policy"`
	PricelistId   struct {
		Id   int64
		Name string
	} `xmlrpc:"pricelist_id"`
	ProcurementGroupId struct {
		Id   int64
		Name string
	} `xmlrpc:"procurement_group_id"`
	ProductId struct {
		Id   int64
		Name string
	} `xmlrpc:"product_id"`
	ProjectId struct {
		Id   int64
		Name string
	} `xmlrpc:"project_id"`
	ProjectProjectId struct {
		Id   int64
		Name string
	} `xmlrpc:"project_project_id"`
	RelatedProjectId struct {
		Id   int64
		Name string
	} `xmlrpc:"related_project_id"`
	SourceId struct {
		Id   int64
		Name string
	} `xmlrpc:"source_id"`
	State      string  `xmlrpc:"state"`
	TagIds     []int64 `xmlrpc:"tag_ids"`
	TasksCount int64   `xmlrpc:"tasks_count"`
	TasksIds   []int64 `xmlrpc:"tasks_ids"`
	TeamId     struct {
		Id   int64
		Name string
	} `xmlrpc:"team_id"`
	TimesheetCount float64 `xmlrpc:"timesheet_count"`
	TimesheetIds   []int64 `xmlrpc:"timesheet_ids"`
	UserId         struct {
		Id   int64
		Name string
	} `xmlrpc:"user_id"`
	ValidityDate string `xmlrpc:"validity_date"`
	WarehouseId  struct {
		Id   int64
		Name string
	} `xmlrpc:"warehouse_id"`
	WriteDate string `xmlrpc:"write_date"`
	WriteUid  struct {
		Id   int64
		Name string
	} `xmlrpc:"write_uid"`
}

type SaleOrderNil struct {
	__LastUpdate             interface{} `xmlrpc:"__last_update"`
	AmountTax                interface{} `xmlrpc:"amount_tax"`
	AmountTotal              interface{} `xmlrpc:"amount_total"`
	AmountUntaxed            interface{} `xmlrpc:"amount_untaxed"`
	AutosalesConfigId        interface{} `xmlrpc:"autosales_config_id"`
	CampaignId               interface{} `xmlrpc:"campaign_id"`
	ClientOrderRef           interface{} `xmlrpc:"client_order_ref"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	ConfirmationDate         interface{} `xmlrpc:"confirmation_date"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	CurrencyId               interface{} `xmlrpc:"currency_id"`
	DateOrder                interface{} `xmlrpc:"date_order"`
	DeliveryCount            interface{} `xmlrpc:"delivery_count"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	FiscalPositionId         interface{} `xmlrpc:"fiscal_position_id"`
	Id                       interface{} `xmlrpc:"id"`
	Incoterm                 interface{} `xmlrpc:"incoterm"`
	InvoiceCount             interface{} `xmlrpc:"invoice_count"`
	InvoiceIds               interface{} `xmlrpc:"invoice_ids"`
	InvoiceStatus            interface{} `xmlrpc:"invoice_status"`
	MediumId                 interface{} `xmlrpc:"medium_id"`
	MessageChannelIds        interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       interface{} `xmlrpc:"message_follower_ids"`
	MessageIds               interface{} `xmlrpc:"message_ids"`
	MessageIsFollower        bool        `xmlrpc:"message_is_follower"`
	MessageLastPost          interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction        bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread            bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter     interface{} `xmlrpc:"message_unread_counter"`
	Name                     interface{} `xmlrpc:"name"`
	Note                     interface{} `xmlrpc:"note"`
	OpportunityId            interface{} `xmlrpc:"opportunity_id"`
	OrderLine                interface{} `xmlrpc:"order_line"`
	Origin                   interface{} `xmlrpc:"origin"`
	PartnerId                interface{} `xmlrpc:"partner_id"`
	PartnerInvoiceId         interface{} `xmlrpc:"partner_invoice_id"`
	PartnerShippingId        interface{} `xmlrpc:"partner_shipping_id"`
	PaymentTermId            interface{} `xmlrpc:"payment_term_id"`
	PickingIds               interface{} `xmlrpc:"picking_ids"`
	PickingPolicy            interface{} `xmlrpc:"picking_policy"`
	PricelistId              interface{} `xmlrpc:"pricelist_id"`
	ProcurementGroupId       interface{} `xmlrpc:"procurement_group_id"`
	ProductId                interface{} `xmlrpc:"product_id"`
	ProjectId                interface{} `xmlrpc:"project_id"`
	ProjectProjectId         interface{} `xmlrpc:"project_project_id"`
	RelatedProjectId         interface{} `xmlrpc:"related_project_id"`
	SourceId                 interface{} `xmlrpc:"source_id"`
	State                    interface{} `xmlrpc:"state"`
	TagIds                   interface{} `xmlrpc:"tag_ids"`
	TasksCount               interface{} `xmlrpc:"tasks_count"`
	TasksIds                 interface{} `xmlrpc:"tasks_ids"`
	TeamId                   interface{} `xmlrpc:"team_id"`
	TimesheetCount           interface{} `xmlrpc:"timesheet_count"`
	TimesheetIds             interface{} `xmlrpc:"timesheet_ids"`
	UserId                   interface{} `xmlrpc:"user_id"`
	ValidityDate             interface{} `xmlrpc:"validity_date"`
	WarehouseId              interface{} `xmlrpc:"warehouse_id"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

type SaleOrders []SaleOrder
type SaleOrdersNil []SaleOrderNil

func (s *SaleOrder) NilableType() interface{} {
	return &SaleOrderNil{}
}

func (ns *SaleOrderNil) Type() interface{} {
	s := &SaleOrder{}
	return convertNil(s, ns)	
}

func (s *SaleOrders) NilableType() interface{} {
	return &SaleOrdersNil{}
}

func (ns *SaleOrdersNil) Type() interface{} {
	s := &SaleOrders{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type().(*SaleOrder))
	}
	return s
}
