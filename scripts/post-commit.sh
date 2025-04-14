#!/bin/bash
set -e

echo "Running post-commit hook: auto-tagging..."

# Store the current working directory
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "$REPO_ROOT"

# Get the latest tag that matches v0.0.X pattern
LATEST_TAG=$(git tag -l "v0.0.*" | sort -V | tail -n 1)

if [ -z "$LATEST_TAG" ]; then
  # No tags yet, start with v0.0.1
  NEW_TAG="v0.0.1"
else
  # Check if the tag matches our expected format
  if [[ $LATEST_TAG =~ ^v0\.0\.([0-9]+)$ ]]; then
    # Extract the version number part
    VERSION_NUMBER=${BASH_REMATCH[1]}
    # Increment it
    NEXT_VERSION=$((VERSION_NUMBER + 1))
    # Create the new tag
    NEW_TAG="v0.0.$NEXT_VERSION"
  else
    # Tag doesn't match expected format, create a new one
    echo "Warning: Latest tag format unexpected, starting new tag series"
    NEW_TAG="v0.0.1"
  fi
fi

# Create the new tag
git tag $NEW_TAG
echo "Created new tag: $NEW_TAG"

# Determine whether to push tag automatically
AUTO_PUSH=${BETHELA_AUTO_PUSH_TAGS:-0}  # Default is 0 (don't push)

if [ "$AUTO_PUSH" == "1" ]; then
  echo "Auto-pushing tag to remote (BETHELA_AUTO_PUSH_TAGS=1)..."
  git push origin $NEW_TAG
  if [ $? -eq 0 ]; then
    echo "Tag pushed successfully! ✓"
    echo "Other projects can now use version $NEW_TAG"
  else
    echo "Error pushing tag. Please push manually with:"
    echo "git push origin $NEW_TAG"
  fi
else
  echo "Tag created locally. To make it available to other projects, push it with:"
  echo "git push origin $NEW_TAG"
  echo ""
  echo "To enable automatic tag pushing, set environment variable BETHELA_AUTO_PUSH_TAGS=1"
  echo "  export BETHELA_AUTO_PUSH_TAGS=1"
fi

exit 0