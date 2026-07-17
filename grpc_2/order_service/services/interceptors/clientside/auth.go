package clientside

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthInterceptor(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	token, ok := ctx.Value("token").(string)
	if !ok {
		log.Default().Print("could not read the token")
		return invoker(ctx, method, req, reply, cc, opts...)
	}
	ctx = metadata.AppendToOutgoingContext(
		ctx,
		"authorization",
        "Bearer "+token,
	)
	return invoker(ctx, method, req, reply, cc, opts...)
}
