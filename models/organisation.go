package models

type Organisation struct {
	BaseAggregateSourced

	Avatar string `json:"avatar" required:"true"`
	Name   string `json:"name" required:"true"`
	Url    string `json:"url" required:"true"`
}
