# VaultMesh 120-Commit Roadmap

## PHASE 1: FOUNDATION (Commits 1-15)
- [x] 001 - init: project scaffold, go.mod, Makefile, .gitignore, README.md, TODO.md
- [x] 002 - config: base configuration system with YAML/env support
- [x] 003 - config: add validation, defaults, and hot-reload support
- [x] 004 - crypto: implement AES-256-GCM encryption/decryption
- [x] 005 - crypto: add ChaCha20-Poly1305 as secondary cipher
- [x] 006 - crypto: implement layered dual encryption pipeline
- [x] 007 - crypto: add Shamir's Secret Sharing (k-of-n key splitting)
- [x] 008 - crypto: key derivation with HKDF + per-chunk unique keys
- [x] 009 - crypto: keystore with encrypted local key management
- [x] 010 - chunker: implement fixed-size chunking algorithm
- [x] 011 - chunker: add content-defined chunking (CDC/Rabin fingerprint)
- [x] 012 - chunker: implement chunk-level deduplication with blake3
- [x] 013 - chunker: add chunk reassembly and integrity verification
- [x] 014 - storage: define storage interface + BadgerDB implementation
- [x] 015 - storage: chunk store with metadata indexing

## PHASE 2: NETWORK LAYER (Commits 16-30)
- [x] 016 - network: LibP2P host initialization with QUIC + TCP
- [x] 017 - network: identity management (PeerID, keypairs)
- [x] 018 - network: Kademlia DHT setup and configuration
- [x] 019 - network: mDNS local peer discovery
- [x] 020 - network: bootstrap peer list + DHT bootstrap
- [x] 021 - network: GossipSub for pub/sub messaging
- [x] 022 - network: custom protocol definitions (upload/download/health)
- [ ] 023 - network: NAT traversal with AutoNAT + relay
- [ ] 024 - network: circuit relay for firewalled peers
- [ ] 025 - network: connection manager with limits + scoring
- [ ] 026 - network: peer scoring and blacklisting
- [ ] 027 - network: WebRTC transport fallback
- [ ] 028 - network: bandwidth measurement + throttling
- [ ] 029 - network: network topology mapper
- [ ] 030 - network: peer exchange protocol (PEX)

## PHASE 3: DISTRIBUTION ENGINE (Commits 31-45)
- [ ] 031 - distributor: Reed-Solomon erasure coding implementation
- [ ] 032 - distributor: configurable redundancy (N data + M parity)
- [ ] 033 - distributor: chunk placement strategy (geo + latency aware)
- [ ] 034 - distributor: parallel chunk upload to multiple peers
- [ ] 035 - distributor: upload progress tracking + resumable uploads
- [ ] 036 - distributor: file manifest creation and DHT publishing
- [ ] 037 - distributor: IPFS-compatible CID generation
- [ ] 038 - distributor: replication manager + target replication factor
- [ ] 039 - distributor: smart peer selection algorithm
- [ ] 040 - distributor: upload pipeline with backpressure
- [ ] 041 - distributor: chunk verification after upload
- [ ] 042 - distributor: metadata encryption before DHT storage
- [ ] 043 - retriever: basic chunk retrieval from DHT
- [ ] 044 - retriever: parallel multi-peer retrieval
- [ ] 045 - retriever: streaming retrieval pipeline

## PHASE 4: PRIVACY & RETRIEVAL (Commits 46-57)
- [ ] 046 - retriever: onion routing for anonymous requests
- [ ] 047 - retriever: Private Information Retrieval (PIR) basic impl
- [ ] 048 - retriever: blind indexing - zero-knowledge storage proofs
- [ ] 049 - retriever: adaptive chunk reassembly from erasure shards
- [ ] 050 - retriever: download resume from checkpoint
- [ ] 051 - retriever: streaming decrypt pipeline (decrypt on-the-fly)
- [ ] 052 - retriever: video/audio progressive streaming support
- [ ] 053 - retriever: IPFS gateway compatibility layer

## PHASE 5: SELF-HEALING & ADVANCED (Commits 54-63)
- [ ] 054 - healing: node health monitoring daemon
- [ ] 055 - healing: heartbeat protocol + failure detection
- [ ] 056 - healing: chunk availability scanner
- [ ] 057 - healing: automatic re-replication on node failure
- [ ] 058 - healing: predictive failure scoring per node
- [ ] 059 - healing: repair scheduler with priority queue
- [ ] 060 - healing: cross-shard repair without full decode
- [ ] 061 - sync: delta sync engine (chunk-level diffing)
- [ ] 062 - sync: version history with merkle DAG
- [ ] 063 - sync: conflict detection and resolution

## PHASE 6: COLLABORATIVE VAULTS & INCENTIVE (Commits 64-75)
- [ ] 064 - vault: collaborative vault creation
- [ ] 065 - vault: multi-signature file ownership
- [ ] 066 - vault: threshold decryption (N-of-M)
- [ ] 067 - vault: access control list per file/vault
- [ ] 068 - incentive: node reputation scoring system
- [ ] 069 - incentive: storage proof challenges (PoSt-lite)
- [ ] 070 - incentive: proof verification without blockchain
- [ ] 071 - incentive: bandwidth + storage credit tracking
- [ ] 072 - incentive: reputation-based peer selection boost
- [ ] 073 - incentive: anti-sybil measures
- [ ] 074 - incentive: storage commitment announcements
- [ ] 075 - incentive: credit ledger with cryptographic receipts

## PHASE 7: API LAYER (Commits 76-84)
- [ ] 076 - api: REST server with Gin framework
- [ ] 077 - api: upload endpoint with multipart streaming
- [ ] 078 - api: download endpoint with range requests
- [ ] 079 - api: file listing, search, delete endpoints
- [ ] 080 - api: peer management endpoints
- [ ] 081 - api: WebSocket server for real-time events
- [ ] 082 - api: gRPC server + proto definitions
- [ ] 083 - api: API key authentication middleware
- [ ] 084 - api: rate limiting + request validation

## PHASE 8: TERMINAL UI (Commits 85-93)
- [ ] 085 - tui: Bubbletea app skeleton + routing
- [ ] 086 - tui: main dashboard with live stats
- [ ] 087 - tui: upload view with drag-drop + progress bars
- [ ] 088 - tui: download view with speed graphs
- [ ] 089 - tui: peer list view with health indicators
- [ ] 090 - tui: network topology visualization (ASCII art nodes)
- [ ] 091 - tui: file browser with search
- [ ] 092 - tui: settings/config panel
- [ ] 093 - tui: Lipgloss theming system (dark/light)

## PHASE 9: WEB UI (Commits 94-103)
- [ ] 094 - web: Vite + React + TypeScript setup
- [ ] 095 - web: design system (Tailwind + Framer Motion)
- [ ] 096 - web: dashboard page with animated stats
- [ ] 097 - web: file manager with drag-and-drop upload
- [ ] 098 - web: real-time transfer progress with WebSocket
- [ ] 099 - web: network graph visualization (D3.js)
- [ ] 100 - web: peer explorer with interactive map
- [ ] 101 - web: vault management UI
- [ ] 102 - web: settings + key management UI
- [ ] 103 - web: mobile responsive + PWA support

## PHASE 10: POLISH & RELEASE (Commits 104-120)
- [ ] 104 - cli: complete Cobra CLI with all subcommands
- [ ] 105 - cli: shell completion (bash/zsh/fish)
- [ ] 106 - docs: full README with architecture diagrams
- [ ] 107 - docs: API documentation with examples
- [ ] 108 - docs: security model documentation
- [ ] 109 - docs: protocol specification
- [ ] 110 - test: unit tests for crypto package (>90% coverage)
- [ ] 111 - test: unit tests for chunker + storage
- [ ] 112 - test: integration tests for network layer
- [ ] 113 - test: e2e 3-node test network simulation
- [ ] 114 - test: benchmarks for throughput + latency
- [ ] 115 - ci: GitHub Actions CI/CD + release pipeline
- [ ] 116 - docker: multi-stage Dockerfile + compose
- [ ] 117 - deploy: Kubernetes manifests
- [ ] 118 - perf: profiling + optimization pass
- [ ] 119 - perf: connection pooling + caching layer
- [ ] 120 - release: v0.1.0 tag + changelog
