package repository

import (
	"antrein/bc-dashboard/application/common/resource"
	"antrein/bc-dashboard/internal/repository/configuration"
	"antrein/bc-dashboard/internal/repository/project"
	"antrein/bc-dashboard/internal/repository/tenant"
	"antrein/bc-dashboard/model/config"
)

type CommonRepository struct {
	TenantRepo  *tenant.Repository
	ProjectRepo *project.Repository
	ConfigRepo  *configuration.Repository
}

func NewCommonRepository(cfg *config.Config, rsc *resource.CommonResource) (*CommonRepository, error) {
	tenantRepo := tenant.New(cfg, rsc.Db)
	projectRepo := project.New(cfg, rsc.Db)
	configRepo := configuration.New(cfg, rsc.Db)

	commonRepo := CommonRepository{
		TenantRepo:  tenantRepo,
		ProjectRepo: projectRepo,
		ConfigRepo:  configRepo,
	}
	return &commonRepo, nil
}
