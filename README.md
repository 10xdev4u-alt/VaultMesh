# 🔥 VaultMesh - Decentralized Encrypted P2P Storage System

VaultMesh is a production-grade, decentralized file storage system built on LibP2P. It splits files into content-defined chunks, encrypts them with layered ciphers, applies Reed-Solomon erasure coding for redundancy, and distributes them across a global Kademlia DHT.

## 🏗️ Architecture

```text
┌─────────────────┐      ┌─────────────────┐      ┌─────────────────┐
│   FILE INPUT    │─────▶│  CDC CHUNKER    │─────▶│ LAYERED ENCRYPT │
└─────────────────┘      └─────────────────┘      └─────────────────┘
                                                           │
                                                           ▼
┌─────────────────┐      ┌─────────────────┐      ┌─────────────────┐
│  DHT DISCOVERY  │◀─────│   DISTRIBUTOR   │◀─────│ REED-SOLOMON EC │
└─────────────────┘      └─────────────────┘      └─────────────────┘
         │                        │
         ▼                        ▼
┌───────────────────────────────────────────────────────────────────┐
│                        LIBP2P MESH NETWORK                        │
└───────────────────────────────────────────────────────────────────┘
```

## 🚀 Key Features
- **Adaptive Chunking:** Content-Defined Chunking (CDC) for efficient deduplication.
- **Layered Encryption:** AES-256-GCM + ChaCha20-Poly1305 + Shamir's.
- **Self-Healing:** Proactive repair of missing shards using predictive failure scoring.
- **Privacy First:** Onion-routed retrieval, blind DHT indexing, and Private Information Retrieval (PIR).
- **Modern UIs:** Sleek Terminal UI (Bubbletea) and Web Dashboard (React/Vite).

## 🛠️ Getting Started

### Prerequisites
- Go 1.22+
- Node.js (for Web UI)

### Installation
1. Clone the repository
2. Run `make deps` to install Go dependencies
3. Run `make build` to compile the `vaultmesh` binary

### Usage
- Start the daemon: `./vaultmesh daemon start`
- Upload a file: `./vaultmesh upload myfile.dat`
- Access the Web UI: `http://localhost:3000`

## 📜 Roadmap
See [TODO.md](./TODO.md) for the detailed 120-commit development plan.
