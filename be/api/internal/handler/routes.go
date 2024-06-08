// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	equipment "be/api/internal/handler/equipment"
	tower "be/api/internal/handler/tower"
	towerequipment "be/api/internal/handler/towerequipment"
	"be/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/tower-detail",
				Handler: tower.CreateTowerDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/tower-detail/:id",
				Handler: tower.UpdateTowerDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/tower-detail/:id",
				Handler: tower.GetTowerDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/tower-details",
				Handler: tower.ListTowerDetailsHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/tower-detail/:id",
				Handler: tower.DeleteTowerDetailHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/tower-equipment/list",
				Handler: towerequipment.ListEquipmentByTowerIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/tower-equipment",
				Handler: towerequipment.CreateTowerEquipmentHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/tower-equipment/:id",
				Handler: towerequipment.DeleteTowerEquipmentHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/equipment-class",
				Handler: equipment.CreateEquipmentClassHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/equipment-class/:id",
				Handler: equipment.UpdateEquipmentClassHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/equipment-class/:id",
				Handler: equipment.GetEquipmentClassHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/equipment-classes",
				Handler: equipment.ListEquipmentClassesHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/equipment-class/:id",
				Handler: equipment.DeleteEquipmentClassHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/equipment-detail",
				Handler: equipment.CreateEquipmentDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/equipment-detail/:id",
				Handler: equipment.UpdateEquipmentDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/equipment-detail/:id",
				Handler: equipment.GetEquipmentDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/equipment-details",
				Handler: equipment.ListEquipmentDetailsHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/equipment-detail/:id",
				Handler: equipment.DeleteEquipmentDetailHandler(serverCtx),
			},
		},
	)
}
