package broker

import (
	"github.com/kyma-project/helm-broker/internal"
	"github.com/pkg/errors"
)

// DTOs for Open Service Broker v2.13 API

type contextDTO struct {
	Platform  string             `json:"platform"`
	Namespace internal.Namespace `json:"namespace"`
}

// ProvisionRequestDTO represents provision request
type ProvisionRequestDTO struct {
	ServiceID        internal.ServiceID     `json:"service_id"`
	PlanID           internal.ServicePlanID `json:"plan_id"`
	OrganizationGUID string                 `json:"organization_guid"`
	SpaceGUID        string                 `json:"space_guid"`
	Parameters       map[string]interface{} `json:"parameters,omitempty"`
	Context          contextDTO             `json:"context,omitempty"`
}

// Validate validates necessary provisioning parameters
func (params *ProvisionRequestDTO) Validate() error {
	if params.OrganizationGUID == "" {
		return errors.New("OrganizationGUID must be non-empty string")
	}
	if params.SpaceGUID == "" {
		return errors.New("SpaceGUID must be non-empty string")
	}
	return nil
}

// ProvisionSuccessResponseDTO represents response after successful provisioning
type ProvisionSuccessResponseDTO struct {
	DashboardURL *string               `json:"dashboard_url"`
	Operation    *internal.OperationID `json:"operation,omitempty"`
}

// DeprovisionSuccessResponseDTO represents response after successful deprovisioning
type DeprovisionSuccessResponseDTO struct {
	Operation *internal.OperationID `json:"operation,omitempty"`
}

// CatalogSuccessResponseDTO represents info about successful catalog response
// TODO: implement me based on osb.CatalogResponse
type CatalogSuccessResponseDTO struct{}

// LastOperationSuccessResponseDTO represents info response about last successful operation
type LastOperationSuccessResponseDTO struct {
	State       internal.OperationState `json:"state"`
	Description *string                 `json:"description,omitempty"`
}

// BindSuccessResponseDTO represents response with credentials for service instance after successful binding
type BindSuccessResponseDTO struct {
	// Credentials is a free-form hash of credentials that can be used by
	// applications or users to access the service.
	Credentials map[string]interface{} `json:"credentials,omitempty"`
}

// BindInProgressResponseDTO represents response with operation for service binding in progress
type BindInProgressResponseDTO struct {
	Operation *internal.OperationID `json:"operation,omitempty"`
}

// BindParametersDTO contains parameters sent by Service Catalog in the body of bind request.
type BindParametersDTO struct {
	ServiceID  internal.ServiceID     `json:"service_id"`
	PlanID     internal.ServicePlanID `json:"plan_id"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Context    contextDTO             `json:"context,omitempty"`
}

// Validate checks if bind parameters aren't empty
func (params *BindParametersDTO) Validate() error {
	if params.PlanID == "" {
		return errors.New("PlanID must be non-empty string")
	}
	if params.ServiceID == "" {
		return errors.New("ServiceID must be non-empty string")
	}
	return nil
}
