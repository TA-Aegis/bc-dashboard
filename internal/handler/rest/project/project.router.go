package project

import (
	guard "antrein/bc-dashboard/application/middleware"
	"antrein/bc-dashboard/internal/usecase/configuration"
	"antrein/bc-dashboard/internal/usecase/project"
	validate "antrein/bc-dashboard/internal/utils/validator"
	"antrein/bc-dashboard/model/config"
	"antrein/bc-dashboard/model/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	cfg           *config.Config
	usecase       *project.Usecase
	configUsecase *configuration.Usecase
	vld           *validator.Validate
}

func New(cfg *config.Config, usecase *project.Usecase, configUsecase *configuration.Usecase, vld *validator.Validate) *Router {
	return &Router{
		cfg:           cfg,
		usecase:       usecase,
		configUsecase: configUsecase,
		vld:           vld,
	}
}

func (r *Router) RegisterRoute(app *fiber.App) {
	g := app.Group("/bc/dashboard/project")
	g.Post("", guard.AuthGuard(r.cfg, r.CreateProject)...)
	g.Put("/config", guard.AuthGuard(r.cfg, r.UpdateProjectConfig)...)
	g.Put("/style", guard.AuthGuard(r.cfg, r.UpdateProjectStyle)...)
}

func (r *Router) CreateProject(g *guard.AuthGuardContext) error {
	req := dto.CreateProjectRequest{}

	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "Request tidak sesuai format")
	}

	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "Request tidak sesuai format")
	}

	err = validate.ValidateCreateProject(req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}

	ctx := g.FiberCtx.Context()
	userID := g.Claims.UserID
	resp, errRes := r.usecase.RegisterNewProject(ctx, req, userID)
	if errRes != nil {
		return g.ReturnError(errRes.Status, errRes.Error)
	}

	return g.ReturnCreated(resp)
}

func (r *Router) UpdateProjectConfig(g *guard.AuthGuardContext) error {
	req := dto.UpdateProjectConfig{}

	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "Request tidak sesuai format")
	}

	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "Request tidak sesuai format")
	}

	ctx := g.FiberCtx.Context()
	errRes := r.configUsecase.UpdateProjectConfig(ctx, req)
	if errRes != nil {
		return g.ReturnError(errRes.Status, errRes.Error)
	}

	return g.ReturnSuccess("Berhasil mengupdate konfigurasi project")
}

func (r *Router) UpdateProjectStyle(g *guard.AuthGuardContext) error {
	req := dto.UpdateProjectStyle{}

	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "Request tidak sesuai format")
	}

	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "Request tidak sesuai format")
	}

	ctx := g.FiberCtx.Context()
	errRes := r.configUsecase.UpdateProjectStyle(ctx, req)
	if errRes != nil {
		return g.ReturnError(errRes.Status, errRes.Error)
	}

	return g.ReturnSuccess("Berhasil mengupdate tampilan project")
}
