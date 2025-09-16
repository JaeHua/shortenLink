# scripts/gen_api.sh
#!/usr/bin/env bash
set -euo pipefail

# 项目根目录
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

echo "[gen_api] using goctl version:"
goctl --version || { echo "goctl not found, please install: go install github.com/zeromicro/go-zero/tools/goctl@latest"; exit 1; }

# 生成 API（会覆盖 goctl 生成的 routes/types 等文件）
echo "[gen_api] generate convert-api"
goctl api go -api api/convert/convert.api -dir cmd/convert-api

echo "[gen_api] generate show-api"
goctl api go -api api/show/show.api -dir cmd/show-api

echo "[gen_api] done."