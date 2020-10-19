package lemi011b

type service struct {
    repo DataRepo
    pres Presenter
}

// NewService TODO
func NewService(
    repo DataRepo,
    pres Presenter,
) Service {
    // Start the loop
    return &service{
        repo: repo,
        pres: pres,
    }
}

// AddNewData TODO
func (svc *service) AddNewData(
    r *Request,
) error {
    id := UUID("lemi011b:" + r.Timestamp.String())
    data := Data{
        id,
        r.Timestamp,
        r.X,
        r.Y,
        r.Z,
        r.Temperature,
    }
    svc.repo.Save(&data)

    resp := Response{
        r.Timestamp,
        r.X,
        r.Y,
        r.Z,
        r.Temperature,
    }
    svc.pres.Send(&resp)

    return nil
}
