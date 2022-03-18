package grpc

import (
	"context"
	"fmt"

	schemav1 "github.com/francois-poidevin/briefly/pkg/gen/go/briefly/schema/v1"
)

// GrpcService gRPC client structure
type GrpcService struct {
	GrpcClient schemav1.SchemaAPIClient
}

// GetUnshortcodedURL get the unshortCoded URL from the gRPC server
func (g *GrpcService) GetUnshortcodedURL(ctx context.Context, hash string) (string, error) {

	gResp, gErr := g.GrpcClient.GetUnShortCodedURL(ctx, &schemav1.GetUnShortCodedURLRequest{Hash: hash})
	if gErr != nil {
		fmt.Println("Error Dude !!") //TODO: change to real log
		return "", gErr
	}

	return gResp.Url, nil
}

func (g *GrpcService) GetShortCodeHash(ctx context.Context, url string) (string, error) {
	gResp, gErr := g.GrpcClient.GetShortCodeHash(ctx, &schemav1.GetShortCodeHashRequest{Url: url})
	if gErr != nil {
		fmt.Println("Error Dude !!") //TODO: change to real log
		return "", gErr
	}

	return gResp.Hash, nil
}
