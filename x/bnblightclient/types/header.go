package types

import (
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"

	errorsmod "cosmossdk.io/errors"
)

// hasherPool holds LegacyKeccak256 hashers for rlpHash.
var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

//go:generate go run github.com/fjl/gencodec -type Header -field-override headerMarshaling -out gen_header_json.go
//go:generate go run ../../rlp/rlpgen -type Header -out gen_header_rlp.go

// BNBHeader represents a block header in the Ethereum blockchain.
type BNBHeader struct {
	ParentHash  common.Hash      `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash      `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address   `json:"miner"`
	Root        common.Hash      `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash      `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash      `json:"receiptsRoot"     gencodec:"required"`
	Bloom       types.Bloom      `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int         `json:"difficulty"       gencodec:"required"`
	Number      *big.Int         `json:"number"           gencodec:"required"`
	GasLimit    uint64           `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64           `json:"gasUsed"          gencodec:"required"`
	Time        uint64           `json:"timestamp"        gencodec:"required"`
	Extra       []byte           `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash      `json:"mixHash"`
	Nonce       types.BlockNonce `json:"nonce"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
	BlobGasUsed *uint64 `json:"blobGasUsed" rlp:"optional"`

	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessBlobGas *uint64 `json:"excessBlobGas" rlp:"optional"`

	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
	ParentBeaconRoot *common.Hash `json:"parentBeaconBlockRoot" rlp:"optional"`

	// caches
	hash atomic.Value `rlp:"-"`
}

// Hash returns the block hash of the header, which is simply the keccak256 hash of its
// RLP encoding.
func (h *BNBHeader) Hash() common.Hash {
	if hash := h.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := rlpHash(h)
	h.hash.Store(v)
	return v
}

// ConvertToBNBHeader decodes the input data into a BNBHeader struct and validates it against the provided Header.
//
// It takes a pointer to a Header struct as input and returns a pointer to a BNBHeader struct and an error.
// The function decodes the RawHeader field of the Header struct into a BNBHeader struct using the rlp package.
// It then checks if the Number field of the BNBHeader struct is equal to the Number field of the Header struct.
// If not, it returns an ErrInvalidHeader error with the message "number not equal".
// It also checks if the Hash() method of the BNBHeader struct is equal to the Hash field of the Header struct.
// If not, it returns an ErrInvalidHeader error with the message "hash not equal".
// It also checks if the ReceiptHash field of the BNBHeader struct is equal to the ReceiptRoot field of the Header struct.
// If not, it returns an ErrInvalidHeader error with the message "receipt hash not equal".
// If all checks pass, it returns the BNBHeader struct and a nil error.
func ConvertToBNBHeader(header *Header) (*BNBHeader, error) {
	bnbHeader, err := UnmarshalBNBHeader(header.RawHeader)
	if err != nil {
		return nil, err
	}

	if bnbHeader.Number.Uint64() != header.Number {
		return nil, errorsmod.Wrap(ErrInvalidHeader, "number not equal")
	}

	if bnbHeader.Hash() != common.Hash(header.Hash) {
		return nil, errorsmod.Wrap(ErrInvalidHeader, "hash not equal")
	}

	if bnbHeader.ParentHash != common.Hash(header.ParentHash) {
		return nil, errorsmod.Wrap(ErrInvalidHeader, "parentHash not equal")
	}

	if bnbHeader.ReceiptHash != common.Hash(header.ReceiptRoot) {
		return nil, errorsmod.Wrap(ErrInvalidHeader, "receipt hash not equal")
	}
	return bnbHeader, nil
}

// VerifyHeaders checks the validity of a sequence of BNBHeaders.
//
// It takes a slice of Header pointers as input and returns an error.
// The function iterates over the headers and checks if the current header's
// hash is equal to the next header's parent hash. It also checks if the
// current header's number is one less than the next header's number.
// If any of these conditions are not met, the function returns an
// ErrInvalidHeader error.
func VerifyHeaders(headers []*Header) error {
	if len(headers) == 0 {
		return nil
	}

	preHeader, err := ConvertToBNBHeader(headers[0])
	if err != nil {
		return err
	}

	for i := 1; i < len(headers)-1; i++ {
		nextHeader, err := ConvertToBNBHeader(headers[i])
		if err != nil {
			return err
		}

		if preHeader.Hash() != nextHeader.ParentHash {
			return errorsmod.Wrap(ErrInvalidHeader, "hash not equal")
		}

		if preHeader.Number.Uint64()+1 != nextHeader.Number.Uint64() {
			return errorsmod.Wrap(ErrInvalidHeader, "hash not equal")
		}

		preHeader = nextHeader
	}
	return nil
}

func rlpHash(x interface{}) (h common.Hash) {
	sha, _ := hasherPool.Get().(crypto.KeccakState)
	defer hasherPool.Put(sha)
	sha.Reset()
	_ = rlp.Encode(sha, x)
	_, _ = sha.Read(h[:])
	return h
}
