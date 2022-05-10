package services

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spinel/gophkeeper-storage-svc/pkg/db"
	"github.com/spinel/gophkeeper-storage-svc/pkg/models"
	pb "github.com/spinel/gophkeeper-storage-svc/pkg/pb"
)

type Server struct {
	H db.Handler
}

func (s *Server) CreateEntity(ctx context.Context, req *pb.CreateEntityRequest) (*pb.CreateEntityResponse, error) {
	var entity models.Entity

	entity.Identifier = req.Identifier
	entity.TypeID = int(req.TypeID)
	if entity.TypeID == 1 {
		entity.Account = models.Account{
			Login:    req.GetAccount().GetLogin(),
			Password: req.GetAccount().GetPassword(),
		}
	}
	if entity.TypeID == 2 {
		cc := models.CreditCard{
			Number:     req.GetCreditCard().Number,
			HolderName: req.GetCreditCard().GetHolderName(),
		}
		entity.CreditCard = cc
	}

	if result := s.H.DB.Create(&entity); result.Error != nil {
		return &pb.CreateEntityResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateEntityResponse{
		Status: http.StatusCreated,
		Id:     entity.ID,
	}, nil
}

func (s *Server) CreateEntityHttp(ctx *gin.Context) {
	var entity models.Entity

	if err := ctx.BindJSON(&entity); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result := s.H.DB.Create(&entity); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, result.Error.Error())
	}

	ctx.JSON(http.StatusCreated, &entity)
}

func (s *Server) FindOne(ctx *gin.Context) {
	var entity models.Entity

	if result := s.H.DB.First(&entity, ctx.Param("id")); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, result.Error.Error())
	}

	switch entity.TypeID {
	case 1:
		var account models.Account
		if result := s.H.DB.First(&account, "entity_id = ?", entity.ID); result.Error != nil {
			ctx.JSON(http.StatusBadRequest, result.Error.Error())
		}
		entity.Account = account
	}

	ctx.JSON(http.StatusCreated, &entity)
}
