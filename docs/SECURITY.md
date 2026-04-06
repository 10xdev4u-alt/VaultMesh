# 🔒 VaultMesh Security Model

VaultMesh is designed with a defense-in-depth approach to ensure data remains private, tamper-proof, and accessible only to authorized users.

## 1. Multi-Layer Encryption
VaultMesh does not rely on a single cipher. Every chunk of data is passed through a layered encryption pipeline:
- **Layer 1: AES-256-GCM** - Provides high-speed authenticated encryption.
- **Layer 2: ChaCha20-Poly1305** - A robust stream cipher that provides additional security against potential cipher-specific vulnerabilities.

## 2. Decentralized Key Management
VaultMesh avoids central key servers:
- **HKDF Derivation:** Unique per-chunk keys are derived from a master key using HMAC-based Key Derivation.
- **Shamir's Secret Sharing:** The master key can be split into `N` shards where any `K` shards can reconstruct the key. These shards are distributed across the peer network.

## 3. Data Privacy
- **Blind Storage:** Storage providers receive encrypted blobs identified only by their BLAKE3 hashes. Providers have no knowledge of file names, types, or ownership.
- **Request Anonymity:** Retrieval requests can be proxied through intermediate "Onion" nodes, preventing storage providers from tracking which users are accessing specific data.

## 4. Integrity & Verification
- **Cryptographic Hashing:** BLAKE3 is used for all data indexing, providing fast and secure integrity checks.
- **Zero-Knowledge Proofs:** VaultMesh uses PoSt-lite (Proof of Spacetime) to verify that storage nodes are continuously storing the data they committed to without requiring them to reveal the data content.
