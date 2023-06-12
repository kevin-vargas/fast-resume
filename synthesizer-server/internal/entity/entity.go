package entity

type Data[T any] struct {
	Data T `json:"data"`
}
