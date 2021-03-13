package admin

import "context"

// SitePresenter TODO
type SitePresenter interface {
    Publish(context.Context, *SiteResponse) error
}

// SiteUseCase TODO
type SiteUseCase interface {
    Create
    Read
    Update
    Delete
}

// SiteGateway TODO
type SiteGateway interface {
    Save(*entities.Site)
}
