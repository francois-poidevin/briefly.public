package cmd

import (
	"context"

	grpcSvc "github.com/francois-poidevin/briefly.public/internal/app/grpc"
	"github.com/francois-poidevin/briefly.public/internal/app/services"
	schemaV1 "github.com/francois-poidevin/briefly/pkg/gen/go/briefly/schema/v1"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ctx, cancel = context.WithCancel(context.Background())
)

var startHttpCmd = &cobra.Command{
	Use:   "startHttp",
	Short: "Allow to start REST API service around tracking of all flights",
	Long:  `The HTTP Rest API service start with config parameters. Several endpoints are available `,
	Run: func(cmd *cobra.Command, args []string) {

		// Initialize config
		// initConfig()

		//setup gRPC connection
		gRPCClient, errGrpc := setup()
		if errGrpc != nil {
			log.WithContext(ctx).WithFields(logrus.Fields{
				"Error": errGrpc,
			}).Error("gRPC connection failed")
			//TODO: do a return to stop the service ?
		}

		svc := services.Services{
			Ctx:        ctx,
			Log:        log,
			GRPCClient: gRPCClient,
		}

		r := gin.Default()
		r.GET("/shortcoder/:url", svc.ShortCode)
		r.GET("/unshortcoder/:hash", svc.UnShortCode)
		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	},
}

func setup() (*grpcSvc.GrpcService, error) {

	gRPCURL := "localhost:5556" //TODO: put gRPC server URL in a configuration file

	// Perform dialing to BB gRPC service
	conn, errGrpcConn := grpc.DialContext(
		ctx,
		gRPCURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
	)

	if errGrpcConn != nil {
		return nil, errGrpcConn
	}

	grpcBrieflyClient := schemaV1.NewSchemaAPIClient(conn)
	log.WithContext(ctx).WithFields(logrus.Fields{
		"URL":  gRPCURL,
		"conn": conn,
	}).Debug("gRPC schema connection initialized")

	gRPCSvc := grpcSvc.GrpcService{
		GrpcClient: grpcBrieflyClient,
	}
	return &gRPCSvc, nil
}
