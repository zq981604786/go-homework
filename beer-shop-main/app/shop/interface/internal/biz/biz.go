package biz

import (
	"context"

	csV1 "github.com/go-kratos/beer-shop/api/cart/service/v1"
	ctV1 "github.com/go-kratos/beer-shop/api/catalog/service/v1"
	osV1 "github.com/go-kratos/beer-shop/api/order/service/v1"
	psV1 "github.com/go-kratos/beer-shop/api/payment/service/v1"
	usV1 "github.com/go-kratos/beer-shop/api/user/service/v1"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewUserUseCase,
	NewDiscovery,
	NewUserServiceClient,
	NewCartServiceClient,
	NewCatalogServiceClient,
	NewOrderServiceClient,
	NewPaymentServiceClient,
)

func NewDiscovery() registry.Discovery {
	cli, err := consulAPI.NewClient(consulAPI.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(cli)
	return r
}

func NewUserServiceClient(r registry.Discovery) usV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///default/beer.user.service"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return usV1.NewUserClient(conn)
}

func NewCartServiceClient(r registry.Discovery) csV1.CartClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///default/beer.cart.service"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return csV1.NewCartClient(conn)
}

func NewCatalogServiceClient(r registry.Discovery) ctV1.CatalogClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///default/beer.catalog.service"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return ctV1.NewCatalogClient(conn)
}

func NewOrderServiceClient(r registry.Discovery) osV1.OrderClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///default/beer.catalog.service"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return osV1.NewOrderClient(conn)
}

func NewPaymentServiceClient(r registry.Discovery) psV1.PaymentClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///default/beer.payment.service"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return psV1.NewPaymentClient(conn)
}
