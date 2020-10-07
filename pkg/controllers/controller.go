package controllers

import "context"

// Controller TODO
type Controller interface {
    Listen(ctx context.Context) error
    Serve(ctx context.Context) error
    ListenAndServe(ctx context.Context) error
}
