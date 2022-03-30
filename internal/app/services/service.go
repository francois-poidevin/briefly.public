package services

import (
	"context"
	"net/http"

	grpcSvc "github.com/francois-poidevin/briefly.public/internal/app/grpc"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Services struct
type Services struct {
	Ctx        context.Context
	Log        *logrus.Logger
	GRPCClient *grpcSvc.GrpcService
}

// Binding from JSON
type ShortCode struct {
	Url string `form:"url" json:"url" xml:"url" binding:"required"`
}

func (s *Services) UnShortCode(c *gin.Context) {
	hash := c.Param("hash")
	s.Log.WithContext(s.Ctx).Info("hash parameter: " + hash)

	unshortedURL, errUnShrt := s.GRPCClient.GetUnshortcodedURL(s.Ctx, hash)
	if errUnShrt != nil {
		s.Log.WithContext(s.Ctx).WithFields(logrus.Fields{
			"Error": errUnShrt,
		}).Error("gRPC connection failed")
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Redirect(301, unshortedURL)
}

func (s *Services) ShortCode(c *gin.Context) {
	var json ShortCode

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.Log.WithContext(s.Ctx).Info("url parameter: " + json.Url)

	hash, errUnShrt := s.GRPCClient.GetShortCodeHash(s.Ctx, json.Url)
	if errUnShrt != nil {
		s.Log.WithContext(s.Ctx).WithFields(logrus.Fields{
			"Error": errUnShrt,
		}).Error("gRPC connection failed")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.String(http.StatusOK, hash)
}
