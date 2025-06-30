#!/bin/bash

# SASTOJ 部署脚本
# 用法: ./deploy.sh [服务名]
# 如果不提供服务名，则部署所有服务

set -e

SERVICES=("admin" "contest" "auth" "gojudge" "freshcup")
SERVICE_TARGETS=("admin" "contest" "public-auth" "gojudge-server" "freshcup-server")

# 获取当前的git commit hash
GIT_COMMIT=$(git rev-parse HEAD | cut -c1-6)

function show_help() {
  echo "SASTOJ 部署脚本"
  echo "用法: ./deploy.sh [服务名]"
  echo "可用的服务:"
  echo "  admin    - 管理后台服务"
  echo "  contest  - 比赛服务"
  echo "  auth     - 认证服务"
  echo "  gojudge  - Go评测服务"
  echo "  freshcup - Freshcup评测服务"
  echo "  all      - 所有服务 (默认)"
  echo ""
  echo "示例:"
  echo "  ./deploy.sh admin     # 只部署admin服务"
  echo "  ./deploy.sh           # 部署所有服务"
}

function deploy_service() {
  local service=$1
  local target=$2
  echo "正在部署 $service 服务..."
  echo "Git commit: $GIT_COMMIT"

  # 1. 构建新镜像
  echo "构建新镜像: sastoj/$target:$GIT_COMMIT"
  docker build --build-arg SERVICE=$service --build-arg GIT_COMMIT=$GIT_COMMIT -t sastoj/$target:$GIT_COMMIT -f Dockerfile --target $target .

  # 2. 停止并移除旧容器
  echo "停止并移除旧容器..."
  docker-compose stop $target || true
  docker-compose rm -f $target || true

  # 3. 部署新容器
  echo "部署新容器..."
  GIT_COMMIT=$GIT_COMMIT docker-compose up -d $target

  # 4. 清理未使用的镜像
  echo "清理未使用的旧镜像..."
  # 获取当前使用的镜像ID
  CURRENT_IMAGE_ID=$(docker images sastoj/$target:$GIT_COMMIT -q)
  # 获取所有同名镜像ID
  ALL_IMAGE_IDS=$(docker images sastoj/$target -q)
  # 移除旧镜像，保留当前使用的
  for IMAGE_ID in $ALL_IMAGE_IDS; do
    if [ "$IMAGE_ID" != "$CURRENT_IMAGE_ID" ]; then
      docker rmi $IMAGE_ID || true
    fi
  done

  echo "$service 服务部署完成! 镜像标签: sastoj/$target:$GIT_COMMIT"
}

# 检查是否需要显示帮助
if [[ "$1" == "-h" || "$1" == "--help" ]]; then
  show_help
  exit 0
fi

# 获取要部署的服务
DEPLOY_SERVICE=${1:-"all"}

# 验证服务名
if [[ "$DEPLOY_SERVICE" != "all" ]]; then
  VALID_SERVICE=false
  for service in "${SERVICES[@]}"; do
    if [[ "$service" == "$DEPLOY_SERVICE" ]]; then
      VALID_SERVICE=true
      break
    fi
  done

  if [[ "$VALID_SERVICE" == "false" ]]; then
    echo "错误: 未知的服务 '$DEPLOY_SERVICE'"
    show_help
    exit 1
  fi
fi

# 部署指定服务或所有服务
if [[ "$DEPLOY_SERVICE" == "all" ]]; then
  echo "部署所有服务..."
  echo "Git commit: $GIT_COMMIT"
  for i in "${!SERVICES[@]}"; do
    deploy_service "${SERVICES[$i]}" "${SERVICE_TARGETS[$i]}"
  done
else
  # 查找服务对应的目标
  for i in "${!SERVICES[@]}"; do
    if [[ "${SERVICES[$i]}" == "$DEPLOY_SERVICE" ]]; then
      deploy_service "$DEPLOY_SERVICE" "${SERVICE_TARGETS[$i]}"
      break
    fi
  done
fi

echo "部署过程完成!"