# scripts/make_release.sh
#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

VERSION="${1:-$(date +%Y%m%d%H%M)}"
DIST_DIR="dist/${VERSION}"
mkdir -p "$DIST_DIR"

echo "[release] building with Go version:"
go version

# 可执行文件清单：api + rpc
declare -a BINS=(
  "convert-api:cmd/convert-api/convert.go"
  "show-api:cmd/show-api/show.go"
  "convert-rpc:cmd/convert-rpc/convert.go"
  "show-rpc:cmd/show-rpc/show.go"
  "sequence-rpc:cmd/sequence-rpc/sequence.go"
)

# 构建（当前平台）
for item in "${BINS[@]}"; do
  name="${item%%:*}"
  path="${item##*:}"
  echo "[release] building ${name}"
  go build -trimpath -ldflags="-s -w" -o "${DIST_DIR}/${name}" "${path}"
done

# 打包配置文件
echo "[release] copying configs"
mkdir -p "${DIST_DIR}/configs"
cp -v cmd/convert-api/etc/convert-api.yaml "${DIST_DIR}/configs/" || true
cp -v cmd/show-api/etc/show-api.yaml "${DIST_DIR}/configs/" || true
cp -v cmd/convert-rpc/etc/convert.yaml "${DIST_DIR}/configs/" || true
cp -v cmd/show-rpc/etc/show.yaml "${DIST_DIR}/configs/" || true
cp -v cmd/sequence-rpc/etc/sequence.yaml "${DIST_DIR}/configs/" || true

# 生成校验和
echo "[release] checksums"
( cd "${DIST_DIR}" && sha256sum * || shasum -a 256 * ) > "${DIST_DIR}/SHA256SUMS" || true

echo "[release] done. output: ${DIST_DIR}"
echo "[release] usage example:"
echo "  ${DIST_DIR}/convert-rpc -f ${DIST_DIR}/configs/convert.yaml"
echo "  ${DIST_DIR}/show-rpc    -f ${DIST_DIR}/configs/show.yaml"
echo "  ${DIST_DIR}/sequence-rpc -f ${DIST_DIR}/configs/sequence.yaml"
echo "  ${DIST_DIR}/convert-api -f ${DIST_DIR}/configs/convert-api.yaml"
echo "  ${DIST_DIR}/show-api    -f ${DIST_DIR}/configs/show-api.yaml"