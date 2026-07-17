package serverside

import (
	"context"
	"errors"
	"order_service/common/jwt"
	"order_service/data/models"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
)

var publicMethods map[string]struct{} = map[string]struct{}{}

func Authenticate(ctx context.Context) (context.Context, error) {
	method, _ := grpc.Method(ctx)
	if _, ok := publicMethods[method]; ok {
		return ctx, nil
	}

	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return ctx, errors.Join(err , errors.New("could not find jwt"))
	}

	user, err := jwt.DecodeJWT[models.User](token)
	if err != nil {
		return ctx, errors.Join(err, errors.New("could not decode token"))
	}

	ctx = context.WithValue(ctx, "userdata", user)
	ctx = context.WithValue(ctx, "token", token)

	return ctx, nil
}
