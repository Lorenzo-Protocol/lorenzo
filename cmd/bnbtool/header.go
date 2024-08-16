package main

import (
	"context"
	"encoding/json"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

type client struct {
	eth *ethclient.Client
}

func newClient(ctx context.Context, url string) (*client, error) {
	c, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, err
	}
	return &client{eth: c}, nil
}

func (c client) genHeaders(ctx context.Context, from, to int64, file string) error {
	headers, err := c.getHeaderBetween(ctx, from, to)
	if err != nil {
		return err
	}
	return c.writeHeaders(ctx, headers, file)
}

func (c client) writeHeaders(ctx context.Context, headers []*Header, file string) error {
	headersBz, err := json.Marshal(headers)
	if err != nil {
		return err
	}
	return os.WriteFile(file, headersBz, 0644)
}

func (c client) getHeaderBetween(ctx context.Context, from, to int64) ([]*Header, error) {
	headers := []*Header{}
	for i := from; i <= to; i++ {
		header, err := c.getLCHeader(ctx, i)
		if err != nil {
			return nil, err
		}
		headers = append(headers, header)
	}
	return headers, nil
}

func (c client) getHeader(ctx context.Context, height int64) (*types.Header, error) {
	header, err := c.eth.HeaderByNumber(context.Background(), big.NewInt(height))
	if err != nil {
		return nil, err
	}
	return header, nil
}

func (c client) getLCHeader(ctx context.Context, height int64) (*Header, error) {
	header, err := c.getHeader(ctx, height)
	if err != nil {
		return nil, err
	}
	headerBz, err := rlp.EncodeToBytes(header)

	return &Header{
		RawHeader:   headerBz,
		ParentHash:  header.ParentHash[:],
		Hash:        header.Hash().Bytes(),
		Number:      header.Number.Uint64(),
		ReceiptRoot: header.ReceiptHash[:],
	}, nil
}

// Header defines the bnb header
type Header struct {
	// header defines the bnb header bytes
	RawHeader []byte `protobuf:"bytes,1,opt,name=raw_header,json=rawHeader,proto3" json:"raw_header,omitempty"`
	// parent_hash defines the previous bnb header hash
	ParentHash []byte `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	// hash defines the bnb header hash
	Hash []byte `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	// number defines the block number
	Number uint64 `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	// receipt_root defines the receipts merkle root hash
	ReceiptRoot []byte `protobuf:"bytes,5,opt,name=receipt_root,json=receiptRoot,proto3" json:"receipt_root,omitempty"`
}


