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
	url := c.Param("url")
	s.Log.WithContext(s.Ctx).Info("url parameter: " + url)

	hash, errUnShrt := s.GRPCClient.GetShortCodeHash(s.Ctx, url)
	if errUnShrt != nil {
		s.Log.WithContext(s.Ctx).WithFields(logrus.Fields{
			"Error": errUnShrt,
		}).Error("gRPC connection failed")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.String(http.StatusOK, hash)
}
