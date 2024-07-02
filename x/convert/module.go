package convert

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {
	// TODO implement me
	panic("implement me")
}

func (a AppModuleBasic) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	// TODO implement me
	panic("implement me")
}

func (a AppModuleBasic) RegisterInterfaces(registry types.InterfaceRegistry) {
	// TODO implement me
	panic("implement me")
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {
	// TODO implement me
	panic("implement me")
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	// TODO implement me
	panic("implement me")
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	// TODO implement me
	panic("implement me")
}

type AppModule struct{}

func (a AppModule) Name() string {
	// TODO implement me
	panic("implement me")
}

func (a AppModule) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	// TODO implement me
	panic("implement me")
}

func (a AppModule) RegisterInterfaces(registry types.InterfaceRegistry) {
	// TODO implement me
	panic("implement me")
}

func (a AppModule) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {
	// TODO implement me
	panic("implement me")
}

func (a AppModule) GetTxCmd() *cobra.Command {
	// TODO implement me
	panic("implement me")
}

func (a AppModule) GetQueryCmd() *cobra.Command {
	// TODO implement me
	panic("implement me")
}
