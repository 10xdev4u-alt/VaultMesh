package distributor

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

// GenerateCID creates an IPFS-compatible CID for the given data using BLAKE3 multihash.
func GenerateCID(data []byte) (string, error) {
	// We use BLAKE3 for high-performance hashing
	pref := cid.Prefix{
		Version:  1,
		Codec:    cid.Raw,
		MhType:   multihash.BLAKE3,
		MhLength: -1, // default length
	}

	c, err := pref.Sum(data)
	if err != nil {
		return "", fmt.Errorf("failed to generate cid: %w", err)
	}

	return c.String(), nil
}

// GenerateCIDFromHash creates a CID from an existing multihash.
func GenerateCIDFromHash(hash []byte) (string, error) {
	mh, err := multihash.Encode(hash, multihash.BLAKE3)
	if err != nil {
		return "", err
	}
	
	c := cid.NewCidV1(cid.Raw, mh)
	return c.String(), nil
}
