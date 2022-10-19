package routes

import (
	"database/sql"
	"github.com/coba/config"
	"github.com/coba/repository"
	uService "github.com/coba/service/users"
	"gorm.io/gorm"
)

type Payload struct {
	Config   *config.Config
	DBGorm   *gorm.DB
	DBSql    *sql.DB
	repoSql  repository.Database
	uService uService.UserService
}

func (p *Payload) InitUserService() {
	if p.repoSql == nil {
		p.InitRepoMysql()
	}

	p.uService = uService.NewUser(p.repoSql)
}

func (p *Payload) InitRepoMysql() {
	p.repoSql = repository.NewGorm(p.DBGorm)
}

func (p *Payload) GetUserService() uService.UserService {
	if p.uService == nil {
		p.InitUserService()
	}

	return p.uService
}
