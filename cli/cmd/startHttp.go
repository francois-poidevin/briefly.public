package cmd

import (
	"context"
	"fmt"
	"os"

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
		initConfig()

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
		r.POST("/shortcoder", svc.ShortCode)
		r.GET("/unshortcoder/:hash", svc.UnShortCode)

		r.Run(conf.Briefly_public.REST.ListenPort) // listen and serve
	},
}

func setup() (*grpcSvc.GrpcService, error) {

	//log handling
	log = logrus.New()

	fmt.Println(fmt.Sprintf("log format json: %t ", conf.Log.JSONFormatter))
	if conf.Log.JSONFormatter {
		fmt.Println("log format: Json")
		log.Formatter = new(logrus.JSONFormatter)
	} else {
		fmt.Println("log format: Text")
		log.Formatter = new(logrus.TextFormatter)                     //default
		log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
		log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	}

	lvl, err := logrus.ParseLevel(conf.Log.Level)
	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Fatal("Not success to parse logrus log level")
		return nil, err
	}
	log.Level = lvl
	log.Out = os.Stdout

	// Perform dialing to BB gRPC service
	conn, errGrpcConn := grpc.DialContext(
		ctx,
		conf.Briefly_public.Briefly.Adress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
	)

	if errGrpcConn != nil {
		return nil, errGrpcConn
	}

	grpcBrieflyClient := schemaV1.NewSchemaAPIClient(conn)

	log.WithContext(ctx).WithFields(logrus.Fields{
		"URL":        conf.Briefly_public.Briefly.Adress,
		"conn State": conn.GetState(),
	}).Debug("gRPC schema connection initialized")

	gRPCSvc := grpcSvc.GrpcService{
		GrpcClient: grpcBrieflyClient,
	}
	return &gRPCSvc, nil
}

func init() {
	startHttpCmd.Flags().StringVar(&cfgFile, "config", "", "config file")
}
