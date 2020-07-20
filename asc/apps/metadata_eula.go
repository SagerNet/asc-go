package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
	"github.com/aaronsky/asc-go/v1/asc/pricing"
)

// EndUserLicenseAgreement defines model for EndUserLicenseAgreement.
type EndUserLicenseAgreement struct {
	Attributes *struct {
		AgreementText *string `json:"agreementText,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *common.RelationshipsData  `json:"data,omitempty"`
			Links *common.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		Territories *struct {
			Data  *[]common.RelationshipsData `json:"data,omitempty"`
			Links *common.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *common.PagingInformation   `json:"meta,omitempty"`
		} `json:"territories,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// EndUserLicenseAgreementCreateRequest defines model for EndUserLicenseAgreementCreateRequest.
type EndUserLicenseAgreementCreateRequest struct {
	Data struct {
		Attributes struct {
			AgreementText string `json:"agreementText"`
		} `json:"attributes"`
		Relationships struct {
			App struct {
				Data common.RelationshipsData `json:"data"`
			} `json:"app"`
			Territories struct {
				Data []common.RelationshipsData `json:"data"`
			} `json:"territories"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// EndUserLicenseAgreementUpdateRequest defines model for EndUserLicenseAgreementUpdateRequest.
type EndUserLicenseAgreementUpdateRequest struct {
	Data struct {
		Attributes *struct {
			AgreementText *string `json:"agreementText,omitempty"`
		} `json:"attributes,omitempty"`
		ID            string `json:"id"`
		Relationships *struct {
			Territories *struct {
				Data *[]common.RelationshipsData `json:"data,omitempty"`
			} `json:"territories,omitempty"`
		} `json:"relationships,omitempty"`
		Type string `json:"type"`
	} `json:"data"`
}

// EndUserLicenseAgreementResponse defines model for EndUserLicenseAgreementResponse.
type EndUserLicenseAgreementResponse struct {
	Data     EndUserLicenseAgreement `json:"data"`
	Included *[]pricing.Territory    `json:"included,omitempty"`
	Links    common.DocumentLinks    `json:"links"`
}

type GetEULAQuery struct {
	FieldsEndUserLicenseAgreements *[]string `url:"fields[endUserLicenseAgreements],omitempty"`
	FieldsTerritories              *[]string `url:"fields[territories],omitempty"`
	Include                        *[]string `url:"include,omitempty"`
	LimitTerritories               *int      `url:"limit[territories],omitempty"`
}

type GetEULAForAppQuery struct {
	FieldsEndUserLicenseAgreements *[]string `url:"fields[endUserLicenseAgreements],omitempty"`
}

// Add a custom end user license agreement (EULA) to an app and configure the territories to which it applies.
func (s *Service) CreateEULA(id string, body *EndUserLicenseAgreementCreateRequest) (*EndUserLicenseAgreementResponse, *common.Response, error) {
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.Post("endUserLicenseAgreements", body, res)
	return res, resp, err
}

// Update the text or territories for your custom end user license agreement.
func (s *Service) UpdateEULA(id string, body *EndUserLicenseAgreementUpdateRequest) (*EndUserLicenseAgreementResponse, *common.Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s", id)
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// Delete the custom end user license agreement that is associated with an app.
func (s *Service) DeleteEULA(id string) (*common.Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s", id)
	return s.Delete(url, nil)
}

// Get the custom end user license agreement associated with an app, and the territories it applies to.
func (s *Service) GetEULA(id string, params *GetEULAQuery) (*EndUserLicenseAgreementResponse, *common.Response, error) {
	url := fmt.Sprintf("endUserLicenseAgreements/%s", id)
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetEULAForApp gets the custom end user license agreement (EULA) for a specific app and the territories where the agreement applies.
func (s *Service) GetEULAForApp(id string, params *GetEULAForAppQuery) (*EndUserLicenseAgreementResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/endUserLicenseAgreement", id)
	res := new(EndUserLicenseAgreementResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
