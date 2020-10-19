package lemi011b

import "time"

// UUID TODO
type UUID string

// Data TODO
type Data struct {
    ID          UUID
    Timestamp   time.Time
    X           int32
    Y           int32
    Z           int32
    Temperature int32
}

// DataRepo TODO
type DataRepo interface {
    Save(*Data) error
}

// Request TODO
type Request struct {
    Timestamp   time.Time
    X           int32
    Y           int32
    Z           int32
    Temperature int32
}

// Service TODO
type Service interface {
    AddNewData(r *Request) error
}

// Response TODO
type Response struct {
    Timestamp   time.Time
    X           int32
    Y           int32
    Z           int32
    Temperature int32
}

// Presenter TODO
type Presenter interface {
    Send(*Response) error
}
