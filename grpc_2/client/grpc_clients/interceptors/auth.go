package interceptors

import (
	"client/storage"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const LOGIN_METHOD string = ""

func AuthInterceptor() grpc.UnaryClientInterceptor{
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		tokenStore := storage.GetTokenStore()
		if method != LOGIN_METHOD {
			if token := tokenStore.Get(); token != "" {
				ctx = metadata.AppendToOutgoingContext(
					ctx,
					"authorization",
	                "Bearer "+token,
				)
			}
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
