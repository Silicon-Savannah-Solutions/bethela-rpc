#!/bin/bash
set -e

echo "Running pre-commit hook: compiling protobuf files..."

# Store the current working directory
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "$REPO_ROOT"

# Run the compile_proto.sh script
./compile_proto.sh

# Check if any .pb.go files were changed by the compilation
if git diff --name-only | grep -q "\.pb\.go$"; then
    echo "Protocol buffer files were recompiled."
    echo "Adding generated files to the commit..."
    git add gen/go/*.pb.go
fi

echo "Pre-commit hook completed successfully"
exit 0