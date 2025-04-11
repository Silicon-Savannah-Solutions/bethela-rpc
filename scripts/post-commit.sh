#!/bin/bash
set -e

echo "Running post-commit hook: auto-tagging..."

# Store the current working directory
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "$REPO_ROOT"

# Get the latest tag (format v0.0.X)
LATEST_TAG=$(git tag -l "v0.0.*" | sort -V | tail -n 1)

if [ -z "$LATEST_TAG" ]; then
  # No tags yet, start with v0.0.1
  NEW_TAG="v0.0.1"
else
  # Extract the version number part
  VERSION_NUMBER=$(echo $LATEST_TAG | sed 's/v0.0.//')
  # Increment it
  NEXT_VERSION=$((VERSION_NUMBER + 1))
  # Create the new tag
  NEW_TAG="v0.0.$NEXT_VERSION"
fi

# Create the new tag
git tag $NEW_TAG
echo "Created new tag: $NEW_TAG"

# Suggest pushing the tag
echo "To make this tag available to other projects, push it with:"
echo "git push origin $NEW_TAG"

# Option to automatically push the tag (uncomment if desired)
# echo "Pushing tag automatically..."
# git push origin $NEW_TAG

exit 0