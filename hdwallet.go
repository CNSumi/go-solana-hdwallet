package hdwallet

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"

	"github.com/btcsuite/btcutil/base58"
	"github.com/tyler-smith/go-bip39"
)

const FirstHardenedChildIndex = 0x80000000

var (
	rootNodeKey = []byte("ed25519 seed")

	defaultDerivationPath = []uint32{
		FirstHardenedChildIndex + 44,
		FirstHardenedChildIndex + 501,
		FirstHardenedChildIndex,
		FirstHardenedChildIndex,
	}
)

type Node struct {
	key, chainCode []byte
}

type newNodeCfg struct {
	password string
	index    uint32
	path     []uint32
}

type NewNodeOption func(cfg *newNodeCfg)

func WithPassword(password string) NewNodeOption {
	return func(cfg *newNodeCfg) {
		cfg.password = password
	}
}

func WithPath(path []uint32) NewNodeOption {
	return func(cfg *newNodeCfg) {
		cfg.path = path
	}
}

// WithIndex only work if no path preset
func WithIndex(index uint32) NewNodeOption {
	return func(cfg *newNodeCfg) {
		cfg.index = index
	}
}

func NewNode(mnemonic string, opts ...NewNodeOption) *Node {
	cfg := &newNodeCfg{
		path: defaultDerivationPath,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	if cfg.index != 0 && cfg.path[2] == FirstHardenedChildIndex {
		cfg.path[2] += cfg.index
	}

	hash := hmac.New(sha512.New, rootNodeKey)
	hash.Write(bip39.NewSeed(mnemonic, cfg.password))
	ans := hash.Sum(nil)

	node := &Node{
		key:       ans[:32],
		chainCode: ans[32:],
	}

	for _, index := range cfg.path {
		node = node.Derive(index)
	}

	return node
}

func (node *Node) DeriveChild(index uint32) *Node {
	data := make([]byte, 37)
	data[0] = 0x00
	copy(data[1:], node.key[:])
	binary.BigEndian.PutUint32(data[33:], index)

	hash := hmac.New(sha512.New, node.chainCode)
	hash.Write(data)
	ans := hash.Sum(nil)

	return &Node{
		key:       ans[:32],
		chainCode: ans[32:],
	}
}

func (node *Node) Derive(path ...uint32) *Node {
	ans := node
	for _, index := range path {
		ans = ans.DeriveChild(index)
	}
	return ans
}

func (node *Node) SecretKey() string {
	secretKey := make([]byte, 64)

	copy(secretKey[:32], node.key)
	copy(secretKey[32:], node.PublicKey())

	return base58.Encode(secretKey)
}

func (node *Node) PublicKey() ed25519.PublicKey {
	return ed25519.NewKeyFromSeed(node.key).Public().(ed25519.PublicKey)
}

func (node *Node) Address() string {
	return base58.Encode(node.PublicKey())
}
