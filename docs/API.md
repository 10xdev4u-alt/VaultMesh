# 📡 VaultMesh API Reference

VaultMesh provides a multi-protocol interface for interacting with the node.

## REST API
Default Port: `8080`

### Authentication
All requests (except `/health`) require the `X-API-Key` header.

### 📤 Upload File
`POST /upload`
- **Body:** Multipart form with `file` field.
- **Success:** `200 OK` with JSON `{ "cid": "...", "message": "..." }`
- **Example:**
```bash
curl -X POST http://localhost:8080/upload \
  -H "X-API-Key: your-key" \
  -F "file=@document.pdf"
```

### 📥 Download File
`GET /download/:cid`
- **Headers:** Supports standard HTTP `Range` headers.
- **Success:** Byte stream of the file.
- **Example:**
```bash
curl http://localhost:8080/download/bafybeigdyrzt5sfp7udm7hu76uh7m \
  -H "X-API-Key: your-key" \
  -o document.pdf
```

### 📋 List Files
`GET /files`
- **Success:** JSON array of file metadata.

---

## 🔌 WebSocket API
Endpoint: `ws://localhost:8080/ws`

Provides real-time event streaming.
- **Events:**
  - `upload_progress`: `{ "type": "upload", "cid": "...", "progress": 45 }`
  - `peer_found`: `{ "type": "peer_discovered", "id": "..." }`

---

## ⚙️ gRPC API
Service: `VaultMeshService`

- `GetNodeInfo`: Returns `peer_id` and `connected_peers`.
- `GetSystemStats`: Returns CPU, Memory, and Disk usage metrics.
