// Code generated by goctl. DO NOT EDIT.
package types

type TowerDetail struct {
	Id          int64    `json:"id"`            // 主键ID
	SubitemId   int64    `json:"subitem_id"`    // 子项ID
	Name        string   `json:"name"`          // 基站名称
	Address     string   `json:"address"`       // 详细地区
	Image       []string `json:"image"`         // 基站图片，考虑是否支持展示多张
	CheckStatus string   `json:"check_status"`  // 检查状态
	CheckTime   string   `json:"check_time"`    // 检查时间
	CheckUserId int64    `json:"check_user_id"` // 检查人ID
	PrincipalId int64    `json:"principal_id"`  // 负责人ID
	PlanTime    string   `json:"plan_time"`     // 计划检查时间
	CreateTime  string   `json:"create_time"`   // 创建时间
	UpdateTime  string   `json:"update_time"`   // 更新时间
}

type CreateTowerDetailReq struct {
	SubitemId int64  `json:"subitem_id"` // 子项ID
	Name      string `json:"name"`       // 基站名称
	Address   string `json:"address"`    // 详细地区
	Image     string `json:"image"`      // 基站图片
	PlanTime  string `json:"plan_time"`  // 计划检查时间
}

type CreateTowerDetailResp struct {
	Id int64 `json:"id"` // 主键ID
}

type UpdateTowerDetailReq struct {
	Id          int64  `json:"id"`            // 主键ID
	SubitemId   int64  `json:"subitem_id"`    // 子项ID
	Name        string `json:"name"`          // 基站名称
	Address     string `json:"address"`       // 详细地区
	Image       string `json:"image"`         // 基站图片
	CheckStatus string `json:"check_status"`  // 检查状态
	CheckTime   string `json:"check_time"`    // 检查时间
	CheckUserId int64  `json:"check_user_id"` // 检查人ID
	PrincipalId int64  `json:"principal_id"`  // 负责人ID
	PlanTime    string `json:"plan_time"`     // 计划检查时间
}

type UpdateTowerDetailResp struct {
	Updated bool `json:"updated"` // 是否更新成功
}

type GetTowerDetailReq struct {
	Id int64 `path:"id"` // 主键ID
}

type GetTowerDetailResp struct {
	TowerDetail TowerDetail `json:"tower_detail"` // 杆塔详情
}

type ListTowerDetailsReq struct {
	PageNo   string `form:"pageNo"`
	PageSize string `form:"pageSize"`
}

type ListTowerDetailsResp struct {
	PageData []TowerDetail `json:"pageData"` // 杆塔详情列表
	Total    int64         `json:"total"`    // 总数
}

type DeleteTowerDetailReq struct {
	Id int64 `json:"id"` // 主键ID
}

type DeleteTowerDetailResp struct {
	Deleted bool `json:"deleted"` // 是否删除成功
}

type ListEquipmentByTowerIdReq struct {
	Id int64 `path:"id"` // 杆塔ID
}

type ListEquipmentByTowerIdResp struct {
	PageData []EquipmentDetail `json:"pageData"` // 设备列表
	Total    int64             `json:"total"`    // 总数
}

type CreateTowerEquipmentReq struct {
	TowerId     int64 `json:"tower_id"`     // 杆塔ID
	EquipmentId int64 `json:"equipment_id"` // 设备管理ID
	PrincipalId int64 `json:"principal_id"` // 负责人ID
}

type CreateTowerEquipmentResp struct {
	Id int64 `json:"id"` // 新增记录的主键ID
}

type DeleteTowerEquipmentReq struct {
	Id int64 `json:"id"` // 主键ID
}

type DeleteTowerEquipmentResp struct {
	Deleted bool `json:"deleted"` // 是否删除成功
}

type EquipmentClass struct {
	Id     int64  `json:"id"`     // 主键ID
	Name   string `json:"name"`   // 名称
	Code   string `json:"code"`   // 代码
	Status int    `json:"status"` // 状态
	PId    int64  `json:"p_id"`   // 父类ID
}

type CreateEquipmentClassReq struct {
	Name   string `json:"name"`   // 名称
	Code   string `json:"code"`   // 代码
	Status int    `json:"status"` // 状态
	PId    int64  `json:"p_id"`   // 父类ID
}

type CreateEquipmentClassResp struct {
	Id int64 `json:"id"` // 主键ID
}

type UpdateEquipmentClassReq struct {
	Id     int64  `json:"id"`     // 主键ID
	Name   string `json:"name"`   // 名称
	Code   string `json:"code"`   // 代码
	Status int    `json:"status"` // 状态
	PId    int64  `json:"p_id"`   // 父类ID
}

type UpdateEquipmentClassResp struct {
	Updated bool `json:"updated"` // 是否更新成功
}

type GetEquipmentClassReq struct {
	Id int64 `path:"id"` // 主键ID
}

type GetEquipmentClassResp struct {
	EquipmentClass EquipmentClass `json:"equipment_class"` // 设备分类
}

type ListEquipmentClassesReq struct {
	Page     int64 `json:"page"`      // 页码
	PageSize int64 `json:"page_size"` // 每页数量
}

type ListEquipmentClassesResp struct {
	EquipmentClasses []EquipmentClass `json:"equipment_classes"` // 设备分类列表
	Total            int64            `json:"total"`             // 总数
}

type DeleteEquipmentClassReq struct {
	Id int64 `json:"id"` // 主键ID
}

type DeleteEquipmentClassResp struct {
	Deleted bool `json:"deleted"` // 是否删除成功
}

type EquipmentDetail struct {
	Id            int64  `json:"id"`             // 主键ID
	ClassId       int64  `json:"class_id"`       // 分类ID
	Name          string `json:"name"`           // 设备名称
	Type          string `json:"type"`           // 设备类型
	SpecificModel string `json:"specific_model"` // 具体型号
	Unit          string `json:"unit"`           // 单位
	Status        string `json:"status"`         // 状态
	Image         string `json:"image"`          // 图片
	CreateUserId  int64  `json:"create_user_id"` // 创建用户ID
}

type CreateEquipmentDetailReq struct {
	ClassId       int64  `json:"class_id"`       // 分类ID
	Name          string `json:"name"`           // 设备名称
	Type          string `json:"type"`           // 设备类型
	SpecificModel string `json:"specific_model"` // 具体型号
	Unit          string `json:"unit"`           // 单位
	Status        string `json:"status"`         // 状态
	Image         string `json:"image"`          // 图片
	CreateUserId  int64  `json:"create_user_id"` // 创建用户ID
}

type CreateEquipmentDetailResp struct {
	Id int64 `json:"id"` // 主键ID
}

type UpdateEquipmentDetailReq struct {
	Id            int64  `json:"id"`             // 主键ID
	ClassId       int64  `json:"class_id"`       // 分类ID
	Name          string `json:"name"`           // 设备名称
	Type          string `json:"type"`           // 设备类型
	SpecificModel string `json:"specific_model"` // 具体型号
	Unit          string `json:"unit"`           // 单位
	Status        string `json:"status"`         // 状态
	Image         string `json:"image"`          // 图片
}

type UpdateEquipmentDetailResp struct {
	Updated bool `json:"updated"` // 是否更新成功
}

type GetEquipmentDetailReq struct {
	Id int64 `path:"id"` // 主键ID
}

type GetEquipmentDetailResp struct {
	EquipmentDetail EquipmentDetail `json:"equipment_detail"` // 设备详情
}

type ListEquipmentDetailsReq struct {
	Page     int64 `json:"page"`      // 页码
	PageSize int64 `json:"page_size"` // 每页数量
}

type ListEquipmentDetailsResp struct {
	EquipmentDetails []EquipmentDetail `json:"equipment_details"` // 设备详情列表
	Total            int64             `json:"total"`             // 总数
}

type DeleteEquipmentDetailReq struct {
	Id int64 `json:"id"` // 主键ID
}

type DeleteEquipmentDetailResp struct {
	Deleted bool `json:"deleted"` // 是否删除成功
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}

type AuthPwReq struct {
	NewPassword string `form:"newPassword"`
	OldPassword string `form:"oldPassword"`
}

type PatchUserReq struct {
	Id       int64   `path:"id"`
	Enable   bool    `json:"enable,omitempty"`
	RoleIds  []int64 `json:"roleIds,omitempty"`
	Password string  `json:"password,omitempty"`
	Username string  `json:"username,omitempty"`
}

type PatchProfileUserReq struct {
	Id       int64  `path:"id"`
	Gender   int64  `json:"gender"`
	NickName string `json:"nickName"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

type EnableRoleReq struct {
	Enable bool  `json:"enable"`
	Id     int64 `path:"id"`
}

type AddUserReq struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Enable   bool    `json:"enable"`
	RoleIds  []int64 `json:"roleIds"`
}

type AddRoleReq struct {
	Code          string  `json:"code"`
	Enable        bool    `json:"enable"`
	Name          string  `json:"name"`
	PermissionIds []int64 `json:"permissionIds"`
}

type PatchRoleReq struct {
	Id            int64   `path:"id"`
	Code          string  `json:"code,omitempty"`
	Enable        bool    `json:"enable,omitempty"`
	Name          string  `json:"name,omitempty"`
	PermissionIds []int64 `json:"permissionIds,omitempty"`
}

type PatchRoleOperateUserReq struct {
	Id      int64   `path:"id"`
	UserIds []int64 `json:"userIds"`
}

type AddPermissionReq struct {
	Type      string `json:"type"`
	ParentId  int64  `json:"parentId,omitempty"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Path      string `json:"path,omitempty"`
	Icon      string `json:"icon,omitempty"`
	Layout    string `json:"layout,omitempty"`
	Component string `json:"component,omitempty"`
	Show      bool   `json:"show"`
	Enable    bool   `json:"enable"`
	KeepAlive bool   `json:"keepAlive"`
	Order     int64  `json:"order"`
}

type PatchPermissionReq struct {
	Id        int64  `path:"id"`
	Type      string `json:"type"`
	ParentId  int64  `json:"parentId,optional"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Path      string `json:"path,optional"`
	Icon      string `json:"icon,optional"`
	Layout    string `json:"layout,optional"`
	Component string `json:"component,optional"`
	Show      int64  `json:"show"`
	Enable    int64  `json:"enable"`
	KeepAlive int64  `json:"keepAlive"`
	Order     int64  `json:"order"`
}

type UserListItem struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Enable     bool   `json:"enable"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	Gender     int    `json:"gender"`
	Avatar     string `json:"avatar"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	Roles      []Role `json:"roles"`
}

type Role struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Enable bool   `json:"enable"`
}

type UserListRes struct {
	PageData []UserListItem `json:"pageData"`
	Total    int64          `json:"total"`
}

type UserListReq struct {
	Gender   string `form:"gender,optional"`
	Enable   string `form:"enable,optional"`
	Username string `form:"username,optional"`
	PageNo   string `form:"pageNo"`
	PageSize string `form:"pageSize"`
}
