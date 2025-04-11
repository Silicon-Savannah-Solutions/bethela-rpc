#!/bin/bash
set -e

echo "Setting up Git hooks for bethela-rpc..."

# Get the root of the repository
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "$REPO_ROOT"

# Ensure the scripts are executable
chmod +x scripts/pre-commit.sh
chmod +x scripts/post-commit.sh

# Create symlinks in the .git/hooks directory
ln -sf ../../scripts/pre-commit.sh .git/hooks/pre-commit
ln -sf ../../scripts/post-commit.sh .git/hooks/post-commit

echo "Hooks installed successfully!"
echo "✓ pre-commit: Automatically compiles Protocol Buffers before each commit"
echo "✓ post-commit: Automatically creates a new version tag after each commit"
echo ""
echo "Remember to push tags after pushing commits:"
echo "git push origin <tag-name>"
echo ""
echo "To push all tags:"
echo "git push origin --tags"