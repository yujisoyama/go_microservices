package interceptor

import (
	"context"
	"encoding/json"

	"github.com/yujisoyama/go_microservices/pkg/logger"
	"google.golang.org/grpc"
)

func LoggingInterceptor(l *logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, h grpc.UnaryHandler) (interface{}, error) {
		reqJSON, err := json.Marshal(req)
		if err == nil {
			// a little less than the GCP limit
			if len(string(reqJSON)) > 250000 {
				reqJSON = []byte(string(reqJSON)[:250000])
			}

			l.Infof("%s request: %s", info.FullMethod, reqJSON)
		} else {
			l.Warnf("json serialization failed for %s request, logging as go value: %v", info.FullMethod, err)
		}
		result, err := h(ctx, req)

		return result, err
	}
}
