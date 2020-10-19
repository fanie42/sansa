package printer

import (
    "fmt"

    "github.com/fanie42/sansa/pkg/lemi011b"
)

// Presenter TODO
type Presenter struct{}

// NewPresenter TODO
func NewPresenter() *Presenter {
    return &Presenter{}
}

// Send TODO
func (pres *Presenter) Send(r *lemi011b.Response) error {
    fmt.Printf(
        "%s, %d, %d, %d, %d\n",
        r.Timestamp.Format("2006-01-02 15:04:05.000000"),
        r.X,
        r.Y,
        r.Z,
        r.Temperature,
    )

    return nil
}
