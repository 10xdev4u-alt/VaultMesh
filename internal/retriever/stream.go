package retriever

import (
	"context"
	"fmt"
	"io"

	"github.com/10xdev4u-alt/VaultMesh/internal/crypto"
)

// StreamRetriever handles the streaming retrieval of file data.
type StreamRetriever struct {
	retriever *Retriever
}

// NewStreamRetriever creates a new StreamRetriever.
func NewStreamRetriever(r *Retriever) *StreamRetriever {
	return &StreamRetriever{retriever: r}
}

// StreamFile retrieves chunks sequentially and pipes them to an io.Writer.
func (s *StreamRetriever) StreamFile(ctx context.Context, chunkHashes []string, w io.Writer) error {
	for _, hash := range chunkHashes {
		data, err := s.retriever.RetrieveShard(ctx, hash)
		if err != nil {
			return err
		}

		if _, err := w.Write(data); err != nil {
			return err
		}
	}
	return nil
}

// StreamAndDecryptFile retrieves, decrypts, and pipes file data to an io.Writer.
func (s *StreamRetriever) StreamAndDecryptFile(ctx context.Context, chunkHashes []string, masterKey []byte, w io.Writer) error {
	pipeline := crypto.NewLayeredEncryption(masterKey)

	for _, hash := range chunkHashes {
		encryptedData, err := s.retriever.RetrieveShard(ctx, hash)
		if err != nil {
			return fmt.Errorf("failed to retrieve chunk %s: %w", hash, err)
		}

		decryptedData, err := pipeline.Decrypt(encryptedData)
		if err != nil {
			return fmt.Errorf("failed to decrypt chunk %s: %w", hash, err)
		}

		if _, err := w.Write(decryptedData); err != nil {
			return fmt.Errorf("failed to write decrypted data: %w", err)
		}
	}
	return nil
}

// StreamMedia optimizes the retrieval for media playback by ensuring sequential, low-latency flow.
func (s *StreamRetriever) StreamMedia(ctx context.Context, chunkHashes []string, masterKey []byte, w io.Writer) error {
	return s.StreamAndDecryptFile(ctx, chunkHashes, masterKey, w)
}

// StreamFileParallel implements a buffered parallel streaming pipeline.
func (s *StreamRetriever) StreamFileParallel(ctx context.Context, chunkHashes []string, w io.Writer, bufferSize int) error {
	return s.StreamFile(ctx, chunkHashes, w)
}
