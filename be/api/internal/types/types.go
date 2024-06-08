// Code generated by goctl. DO NOT EDIT.
package types

type BaseStation struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Image       string `json:"image,omitempty"` // Assuming images are stored as JSON array strings
	CheckStatus string `json:"check_status,omitempty"`
	CheckTime   string `json:"check_time,omitempty"` // Using string for date-time to simplify
	CheckUserId int64  `json:"check_user_id,omitempty"`
	PrincipalId int64  `json:"principal_id,omitempty"`
	PlanTime    string `json:"plan_time,omitempty"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

type CreateBaseStationReq struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Image    string `json:"image,omitempty"`
	PlanTime string `json:"plan_time,omitempty"`
}

type CreateBaseStationResp struct {
	Id int64 `json:"id"`
}

type UpdateBaseStationReq struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Image       string `json:"image,omitempty"`
	CheckStatus string `json:"check_status,omitempty"`
	CheckTime   string `json:"check_time,omitempty"`
	CheckUserId int64  `json:"check_user_id,omitempty"`
	PrincipalId int64  `json:"principal_id,omitempty"`
	PlanTime    string `json:"plan_time,omitempty"`
}

type UpdateBaseStationResp struct {
	Updated bool `json:"updated"`
}

type GetBaseStationReq struct {
	Id int64 `path:"id"`
}

type GetBaseStationResp struct {
	BaseStation BaseStation `json:"base_station"`
}

type ListBaseStationsReq struct {
	Page      int64 `json:"page"`
	PageSizse int64 `json:"page_size"`
}

type ListBaseStationsResp struct {
	BaseStations []BaseStation `json:"base_stations"`
	Total        int64         `json:"total"`
}

type DeleteBaseStationReq struct {
	Id int64 `json:"id"`
}

type DeleteBaseStationResp struct {
	Deleted bool `json:"deleted"`
}
