package archiving

import "github.com/fanie42/sansa"

type projection struct {
    es sansa.EventStore
    pr sansa.Projector
}

// NewProjection TODO
func NewProjection(
    eventstore sansa.EventStore,
    projector sansa.Projector,
) sansa.Projection {
    return &projection{
        es: eventstore,
        pr: projector,
    }
}

func (proj *projection) Project() {

}
