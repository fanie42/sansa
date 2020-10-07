package domains

import "time"

// Model TODO
type Model interface {
    // encoding.BinaryMarshaler
    // encoding.BinaryUnmarshaler
    // encoding.TextMarshaler
    // encoding.TextUnmarshaler
    ID() string
    Timestamp() time.Time
    String() string
}
