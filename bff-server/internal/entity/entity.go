package entity

type Channel struct {
	ID          string   `json:"id" validate:"required"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Err         []string `json:"error,omitempty"`
}

type Data[T any] struct {
	Data T `json:"data"`
}

type Channels = Data[[]Channel]
