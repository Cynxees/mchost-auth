package api

import (
	manager "mchost/auth/server/jwt"
	"mchost/auth/server/pb"
	"mchost/auth/server/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	pb.UserServiceServer
	Db      *gorm.DB
	Logger  *logrus.Logger
	Manager *manager.JWTManager
	Config *config.Config
}