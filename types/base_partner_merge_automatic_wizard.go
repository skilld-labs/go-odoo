package types

import (
	"time"
)

type BasePartnerMergeAutomaticWizard struct {
	LastUpdate         time.Time `xmlrpc:"__last_update"`
	CreateDate         time.Time `xmlrpc:"create_date"`
	CreateUid          Many2One  `xmlrpc:"create_uid"`
	CurrentLineId      Many2One  `xmlrpc:"current_line_id"`
	DisplayName        string    `xmlrpc:"display_name"`
	DstPartnerId       Many2One  `xmlrpc:"dst_partner_id"`
	ExcludeContact     bool      `xmlrpc:"exclude_contact"`
	ExcludeJournalItem bool      `xmlrpc:"exclude_journal_item"`
	GroupByEmail       bool      `xmlrpc:"group_by_email"`
	GroupByIsCompany   bool      `xmlrpc:"group_by_is_company"`
	GroupByName        bool      `xmlrpc:"group_by_name"`
	GroupByParentId    bool      `xmlrpc:"group_by_parent_id"`
	GroupByVat         bool      `xmlrpc:"group_by_vat"`
	Id                 int64     `xmlrpc:"id"`
	LineIds            []int64   `xmlrpc:"line_ids"`
	MaximumGroup       int64     `xmlrpc:"maximum_group"`
	NumberGroup        int64     `xmlrpc:"number_group"`
	PartnerIds         []int64   `xmlrpc:"partner_ids"`
	State              string    `xmlrpc:"state"`
	WriteDate          time.Time `xmlrpc:"write_date"`
	WriteUid           Many2One  `xmlrpc:"write_uid"`
}

type BasePartnerMergeAutomaticWizardNil struct {
	LastUpdate         interface{} `xmlrpc:"__last_update"`
	CreateDate         interface{} `xmlrpc:"create_date"`
	CreateUid          interface{} `xmlrpc:"create_uid"`
	CurrentLineId      interface{} `xmlrpc:"current_line_id"`
	DisplayName        interface{} `xmlrpc:"display_name"`
	DstPartnerId       interface{} `xmlrpc:"dst_partner_id"`
	ExcludeContact     bool        `xmlrpc:"exclude_contact"`
	ExcludeJournalItem bool        `xmlrpc:"exclude_journal_item"`
	GroupByEmail       bool        `xmlrpc:"group_by_email"`
	GroupByIsCompany   bool        `xmlrpc:"group_by_is_company"`
	GroupByName        bool        `xmlrpc:"group_by_name"`
	GroupByParentId    bool        `xmlrpc:"group_by_parent_id"`
	GroupByVat         bool        `xmlrpc:"group_by_vat"`
	Id                 interface{} `xmlrpc:"id"`
	LineIds            interface{} `xmlrpc:"line_ids"`
	MaximumGroup       interface{} `xmlrpc:"maximum_group"`
	NumberGroup        interface{} `xmlrpc:"number_group"`
	PartnerIds         interface{} `xmlrpc:"partner_ids"`
	State              interface{} `xmlrpc:"state"`
	WriteDate          interface{} `xmlrpc:"write_date"`
	WriteUid           interface{} `xmlrpc:"write_uid"`
}

var BasePartnerMergeAutomaticWizardModel string = "base.partner.merge.automatic.wizard"

type BasePartnerMergeAutomaticWizards []BasePartnerMergeAutomaticWizard

type BasePartnerMergeAutomaticWizardsNil []BasePartnerMergeAutomaticWizardNil

func (s *BasePartnerMergeAutomaticWizard) NilableType_() interface{} {
	return &BasePartnerMergeAutomaticWizardNil{}
}

func (ns *BasePartnerMergeAutomaticWizardNil) Type_() interface{} {
	s := &BasePartnerMergeAutomaticWizard{}
	return load(ns, s)
}

func (s *BasePartnerMergeAutomaticWizards) NilableType_() interface{} {
	return &BasePartnerMergeAutomaticWizardsNil{}
}

func (ns *BasePartnerMergeAutomaticWizardsNil) Type_() interface{} {
	s := &BasePartnerMergeAutomaticWizards{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BasePartnerMergeAutomaticWizard))
	}
	return s
}
