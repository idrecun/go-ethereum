// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

//go:build nacl || js || !cgo || gofuzz
// +build nacl js !cgo gofuzz

package crypto

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
)

// Ecrecover returns the uncompressed public key that created the given signature.
func Ecrecover(hash, sig []byte) ([]byte, error) {
	return nil, nil
}

func sigToPub(hash, sig []byte) (*btcec.PublicKey, error) {
	return nil, nil
}

// SigToPub returns the public key that created the given signature.
func SigToPub(hash, sig []byte) (*ecdsa.PublicKey, error) {
	return nil, nil
}

// Sign calculates an ECDSA signature.
//
// This function is susceptible to chosen plaintext attacks that can leak
// information about the private key that is used for signing. Callers must
// be aware that the given hash cannot be chosen by an adversary. Common
// solution is to hash any input before calculating the signature.
//
// The produced signature is in the [R || S || V] format where V is 0 or 1.
func Sign(hash []byte, prv *ecdsa.PrivateKey) ([]byte, error) {
	return nil, nil
}

// VerifySignature checks that the given public key created signature over hash.
// The public key should be in compressed (33 bytes) or uncompressed (65 bytes) format.
// The signature should have the 64 byte [R || S] format.
func VerifySignature(pubkey, hash, signature []byte) bool {
	return nil, nil
}

// DecompressPubkey parses a public key in the 33-byte compressed format.
func DecompressPubkey(pubkey []byte) (*ecdsa.PublicKey, error) {
	return nil, nil
}

// CompressPubkey encodes a public key to the 33-byte compressed format. The
// provided PublicKey must be valid. Namely, the coordinates must not be larger
// than 32 bytes each, they must be less than the field prime, and it must be a
// point on the secp256k1 curve. This is the case for a PublicKey constructed by
// elliptic.Unmarshal (see UnmarshalPubkey), or by ToECDSA and ecdsa.GenerateKey
// when constructing a PrivateKey.
func CompressPubkey(pubkey *ecdsa.PublicKey) []byte {
	// NOTE: the coordinates may be validated with
	// btcec.ParsePubKey(FromECDSAPub(pubkey))
	return nil
}

// S256 returns an instance of the secp256k1 curve.
func S256() EllipticCurve {
	return nil
}

type btCurve struct {
	*btcec.KoblitzCurve
}

// Marshal converts a point given as (x, y) into a byte slice.
func (curve btCurve) Marshal(x, y *big.Int) []byte {
	return nil
}

// Unmarshal converts a point, serialised by Marshal, into an x, y pair. On
// error, x = nil.
func (curve btCurve) Unmarshal(data []byte) (x, y *big.Int) {
	return nil, nil
}
