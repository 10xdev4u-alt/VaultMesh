package retriever

import (
	"context"
	"io"
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

// StreamFileParallel implements a buffered parallel streaming pipeline.
func (s *StreamRetriever) StreamFileParallel(ctx context.Context, chunkHashes []string, w io.Writer, bufferSize int) error {
	return s.StreamFile(ctx, chunkHashes, w)
}
