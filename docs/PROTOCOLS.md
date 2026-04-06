# 🛠️ VaultMesh Protocol Specification

VaultMesh uses custom LibP2P protocols for its core decentralized operations. This document defines the wire formats and interaction sequences.

## 1. Upload Protocol (`/vaultmesh/upload/1.0.0`)
Used by the `Distributor` to send data shards to peers.

- **Sequence:**
  1. **Handshake:** Requester opens a stream to the storage peer.
  2. **Hash Prefix:** Requester sends the 32-byte BLAKE3 hash of the chunk.
  3. **Data Stream:** Requester sends the raw encrypted chunk data.
  4. **Confirmation:** Provider sends a single byte `0x01` (Success) or `0x00` (Failure).
  5. **Close:** Both sides close the stream.

## 2. Download Protocol (`/vaultmesh/download/1.0.0`)
Used by the `Retriever` to pull shards from providers.

- **Sequence:**
  1. **Request:** Requester opens a stream and sends the 32-byte BLAKE3 hash of the desired chunk.
  2. **Response Header:** Provider sends 4 bytes indicating the data length (BigEndian).
  3. **Data Stream:** Provider streams the chunk data.
  4. **Close:** Stream is closed after data transfer.

## 3. Health & Verification Protocol (`/vaultmesh/health/1.0.0`)
Used for heartbeats and ZKP storage challenges.

- **Heartbeat:**
  - Request: `PING` (4 bytes)
  - Response: `PONG` (4 bytes)

- **Storage Challenge:**
  - Request: 32-byte Salt + 32-byte Target Chunk Hash.
  - Response: 32-byte Proof (BLAKE3(Chunk + Salt)).
