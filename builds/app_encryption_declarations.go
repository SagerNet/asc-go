package builds

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/apps"
	"github.com/aaronsky/asc-go/internal"
)

// AppEncryptionDeclarationState defines model for AppEncryptionDeclarationState.
type AppEncryptionDeclarationState string

// List of AppEncryptionDeclarationState
const (
	Approved AppEncryptionDeclarationState = "APPROVED"
	Expired  AppEncryptionDeclarationState = "EXPIRED"
	Invalid  AppEncryptionDeclarationState = "INVALID"
	InReview AppEncryptionDeclarationState = "IN_REVIEW"
	Rejected AppEncryptionDeclarationState = "REJECTED"
)

// AppEncryptionDeclaration defines model for AppEncryptionDeclaration.
type AppEncryptionDeclaration struct {
	Attributes *struct {
		AppEncryptionDeclarationState   *AppEncryptionDeclarationState `json:"appEncryptionDeclarationState,omitempty"`
		AvailableOnFrenchStore          *bool                          `json:"availableOnFrenchStore,omitempty"`
		CodeValue                       *string                        `json:"codeValue,omitempty"`
		ContainsProprietaryCryptography *bool                          `json:"containsProprietaryCryptography,omitempty"`
		ContainsThirdPartyCryptography  *bool                          `json:"containsThirdPartyCryptography,omitempty"`
		DocumentName                    *string                        `json:"documentName,omitempty"`
		DocumentType                    *string                        `json:"documentType,omitempty"`
		DocumentURL                     *string                        `json:"documentUrl,omitempty"`
		Exempt                          *bool                          `json:"exempt,omitempty"`
		Platform                        *apps.Platform                 `json:"platform,omitempty"`
		UploadedDate                    *time.Time                     `json:"uploadedDate,omitempty"`
		UsesEncryption                  *bool                          `json:"usesEncryption,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppEncryptionDeclarationBuildsLinkagesRequest defines model for AppEncryptionDeclarationBuildsLinkagesRequest.
type AppEncryptionDeclarationBuildsLinkagesRequest struct {
	Data []internal.RelationshipsData `json:"data"`
}

// AppEncryptionDeclarationResponse defines model for AppEncryptionDeclarationResponse.
type AppEncryptionDeclarationResponse struct {
	Data     AppEncryptionDeclaration `json:"data"`
	Included *[]apps.App              `json:"included,omitempty"`
	Links    internal.DocumentLinks   `json:"links"`
}

// AppEncryptionDeclarationsResponse defines model for AppEncryptionDeclarationsResponse.
type AppEncryptionDeclarationsResponse struct {
	Data     []AppEncryptionDeclaration  `json:"data"`
	Included *[]apps.App                 `json:"included,omitempty"`
	Links    internal.PagedDocumentLinks `json:"links"`
	Meta     *internal.PagingInformation `json:"meta,omitempty"`
}

type ListAppEncryptionDeclarationsQuery struct {
	FieldsAppEncryptionDeclarations *[]string `url:"fields[appEncryptionDeclarations],omitempty"`
	FieldsApps                      *[]string `url:"fields[apps],omitempty"`
	FilterApp                       *[]string `url:"filter[app],omitempty"`
	FilterBuilds                    *[]string `url:"filter[builds],omitempty"`
	FilterPlatforms                 *[]string `url:"filter[platforms],omitempty"`
	Include                         *[]string `url:"include,omitempty"`
	Limit                           *int      `url:"limit,omitempty"`
	Cursor                          *string   `url:"cursor,omitempty"`
}

type GetAppEncryptionDeclarationQuery struct {
	FieldsAppEncryptionDeclarations *[]string `url:"fields[appEncryptionDeclarations],omitempty"`
	FieldsApps                      *[]string `url:"fields[apps],omitempty"`
	Include                         *[]string `url:"include,omitempty"`
}

type GetAppForEncryptionDeclarationQuery struct {
	FieldsApps *[]string `url:"fields[apps],omitempty"`
}

// ListAppEncryptionDeclarations finds and lists all available app encryption declarations.
func (s *Service) ListAppEncryptionDeclarations(params *ListAppEncryptionDeclarationsQuery) (*AppEncryptionDeclarationsResponse, *internal.Response, error) {
	res := new(AppEncryptionDeclarationsResponse)
	resp, err := s.GetWithQuery("appEncryptionDeclarations", params, res)
	return res, resp, err
}

// GetAppEncryptionDeclaration gets information about a specific app encryption declaration.
func (s *Service) GetAppEncryptionDeclaration(id string, params *GetAppEncryptionDeclarationQuery) (*AppEncryptionDeclarationResponse, *internal.Response, error) {
	url := fmt.Sprintf("appEncryptionDeclarations/%s", id)
	res := new(AppEncryptionDeclarationResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppForAppEncryptionDeclaration gets the app information from a specific app encryption declaration.
func (s *Service) GetAppForAppEncryptionDeclaration(id string, params *GetAppForEncryptionDeclarationQuery) (*apps.AppResponse, *internal.Response, error) {
	url := fmt.Sprintf("appEncryptionDeclarations/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// AssignBuildsToAppEncryptionDeclaration assigns one or more builds to an app encryption declaration.
func (s *Service) AssignBuildsToAppEncryptionDeclaration(id string, body *AppEncryptionDeclarationBuildsLinkagesRequest) (*internal.Response, error) {
	url := fmt.Sprintf("appStoreVersionSubmissions/%s", id)
	return s.Post(url, body, nil)
}