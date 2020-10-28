package model

type NetworkRequest struct {
	NetworkRequestId int    `json:"networkRequestId" db:"network_request_id"`
	UserBaseId       int    `json:"userBaseId" db:"user_base_id"`
	UserRelationId   int    `json:"userRelationId" db:"user_relation_id"`
	Status           int    `json:"status" db:"status"`
	CreatedAt        string `json:"-" db:"created_at"`
	UpdatedAt        string `json:"-" db:"updated_at"`
}
type Network struct {
	NetworkId      int    `json:"networkId" db:"network_id"`
	UserBaseId     int    `json:"userBaseId" db:"user_base_id"`
	UserRelationId int    `json:"userRelationId" db:"user_relation_id"`
	Status         int    `json:"status" db:"status"`
	CreatedAt      string `json:"-" db:"created_at"`
	UpdatedAt      string `json:"-" db:"updated_at"`
}
