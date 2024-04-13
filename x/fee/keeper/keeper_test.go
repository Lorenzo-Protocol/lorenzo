package keeper

import (
	"reflect"
	"testing"

	"github.com/Lorenzo-Protocol/lorenzo/x/fee/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types"
)

func TestKeeper_GetParams(t *testing.T) {
	type fields struct {
		cdc       codec.BinaryCodec
		storeKey  types.StoreKey
		authority string
	}
	type args struct {
		ctx types.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantParams types.Params
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Keeper{
				cdc:       tt.fields.cdc,
				storeKey:  tt.fields.storeKey,
				authority: tt.fields.authority,
			}
			if gotParams := k.GetParams(tt.args.ctx); !reflect.DeepEqual(gotParams, tt.wantParams) {
				t.Errorf("GetParams() = %v, want %v", gotParams, tt.wantParams)
			}
		})
	}
}

func TestKeeper_SetParams(t *testing.T) {
	type fields struct {
		cdc       codec.BinaryCodec
		storeKey  types.StoreKey
		authority string
	}
	type args struct {
		ctx    types.Context
		params types.Params
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Keeper{
				cdc:       tt.fields.cdc,
				storeKey:  tt.fields.storeKey,
				authority: tt.fields.authority,
			}
			if err := k.SetParams(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("SetParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewKeeper(t *testing.T) {
	type args struct {
		cdc       codec.BinaryCodec
		storeKey  types.StoreKey
		authority string
	}
	tests := []struct {
		name string
		args args
		want *Keeper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKeeper(tt.args.cdc, tt.args.storeKey, tt.args.authority); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeeper() = %v, want %v", got, tt.want)
			}
		})
	}
}
