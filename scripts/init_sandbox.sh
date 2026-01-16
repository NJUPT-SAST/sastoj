#!/bin/bash
set -e

# Disable default sources
if [ -f /etc/apt/sources.list.d/debian.sources ]; then
    mv /etc/apt/sources.list.d/debian.sources /etc/apt/sources.list.d/debian.sources.bak
fi

if [ -f /etc/apt/sources.list ]; then
    mv /etc/apt/sources.list /etc/apt/sources.list.bak
fi

echo "Replacing APT sources with USTC mirrors..."
# Using http to avoid potential ssl issues initially
cat > /etc/apt/sources.list <<EOF
deb http://mirrors.ustc.edu.cn/debian/ trixie main contrib non-free non-free-firmware
deb http://mirrors.ustc.edu.cn/debian/ trixie-updates main contrib non-free non-free-firmware
deb http://mirrors.ustc.edu.cn/debian/ trixie-backports main contrib non-free non-free-firmware
deb http://mirrors.ustc.edu.cn/debian-security trixie-security main contrib non-free non-free-firmware
EOF

echo "Updating package list..."
apt-get update

echo "Installing toolchain..."
# Install essential compilers and interpreters
# gcc, g++: C/C++
# python3: Python
# openjdk-17-jdk: Java
# golang: Go
# rustc: Rust
# ghc: Haskell
# ruby: Ruby
apt-get install -y \
    wget \
    curl \
    python3 \
    python3-pip \
    gcc \
    g++ \
    make \
    libtool \
    rustc \
    ruby \
    golang \
    ghc \
    openjdk-21-jdk

echo "Verifying installation..."
g++ --version
python3 --version
go version
javac -version

echo "Toolchain installation complete."
