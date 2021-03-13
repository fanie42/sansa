package graphql

import "context"

// AdminController TODO - equivalent to root resolver
type AdminController struct {
    service *admin.Service
}

// NewAdminController TODO
func NewAdminController(
    service *admin.Service,
) *AdminController {
    return &AdminController{
        service: service,
    }
}

// GetSite TODO
func (ctrl *AdminController) GetSite(
    ctx context.Context,
    id graphl.ID,
) error {
    site := ctrl.service.FindSite(ctx, str(id))

    return site
}
