# scripts/gen_rpc.sh
#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

echo "[gen_rpc] using goctl version:"
goctl --version || { echo "goctl not found, please install: go install github.com/zeromicro/go-zero/tools/goctl@latest"; exit 1; }

# convert-rpc
echo "[gen_rpc] generate convert-rpc"
goctl rpc protoc rpc/convert/convert.proto \
  --go_out=cmd/convert-rpc \
  --go-grpc_out=cmd/convert-rpc \
  --zrpc_out=cmd/convert-rpc

# show-rpc
echo "[gen_rpc] generate show-rpc"
goctl rpc protoc rpc/show/show.proto \
  --go_out=cmd/show-rpc \
  --go-grpc_out=cmd/show-rpc \
  --zrpc_out=cmd/show-rpc

# sequence-rpc（如果存在）
if [[ -f "rpc/sequence/sequence.proto" ]]; then
  echo "[gen_rpc] generate sequence-rpc"
  goctl rpc protoc rpc/sequence/sequence.proto \
    --go_out=cmd/sequence-rpc \
    --go-grpc_out=cmd/sequence-rpc \
    --zrpc_out=cmd/sequence-rpc
fi

echo "[gen_rpc] done."