package request

type StoreSearch struct {
	Id     int    `json:"id" form:"id"`
	Search string `json:"search" form:"search"`
}
