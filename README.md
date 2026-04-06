# 🔥 VaultMesh - Decentralized Encrypted P2P Storage System

VaultMesh is a decentralized, encrypted P2P file storage system built on LibP2P. It splits files into chunks using Content-Defined Chunking, encrypts each chunk with layered AES-256-GCM + ChaCha20-Poly1305, applies Reed-Solomon erasure coding for redundancy, and then distributes chunks across a Kademlia DHT peer network via LibP2P.

## Key Features

- **Zero Central Servers:** Entirely peer-to-peer over LibP2P and Kademlia DHT.
- **End-to-End Encrypted:** Nodes cannot read stored data. Multi-layer encryption and Shamir's Secret Sharing.
- **Self-Healing:** Automatic health monitoring and re-replication on node failure.
- **IPFS-Compatible:** CID generation and gateway support.
- **Private Retrieval:** Onion routing and Private Information Retrieval (PIR) protocols.
- **Streaming Support:** Real-time progressive decryption pipeline.

## Getting Started

1. Clone the repository
2. Run `make deps`
3. Run `make build`

See [TODO.md](./TODO.md) for the roadmap of this project.
