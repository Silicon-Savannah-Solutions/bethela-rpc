# Bethela RPC

This repository contains the Protocol Buffers definitions for the Bethela microservices architecture.

## Overview

Bethela RPC provides the interface definitions for communication between microservices in the Bethela platform. It uses Protocol Buffers and gRPC for efficient, strongly-typed service interfaces.

## Services

- **Identity Service**: User management, authentication, and authorization
- **Wallet Service**: Virtual wallet and transaction management
- **Provider Interface**: External service provider integration

## Development

### Prerequisites

- Protocol Buffers compiler (`protoc`) - [Installation Guide](https://github.com/protocolbuffers/protobuf/releases)
- Go 1.24 or higher

### Compiling Protocol Buffers

To compile the Protocol Buffer definitions to Go code, run:

```bash
./compile_proto.sh
```

This script generates the Go files in the `gen/go` directory.

### Git Hooks

This repository includes Git hooks that automate various tasks. When you clone the repository, you need to set up these hooks:

#### Pre-commit Hook

The pre-commit hook automatically compiles the Protocol Buffers before each commit:

```bash
# Make the pre-commit hook executable
chmod +x .git/hooks/pre-commit

# Or create a symlink from the provided hook
ln -sf ../../scripts/pre-commit.sh .git/hooks/pre-commit
```

#### Post-commit Hook (Auto-tagging)

The post-commit hook automatically creates a new version tag after each commit:

```bash
# Make the post-commit hook executable
chmod +x .git/hooks/post-commit

# Or create a symlink from the provided hook
ln -sf ../../scripts/post-commit.sh .git/hooks/post-commit
```

This automatic tagging ensures that each commit has a corresponding version that other subprojects can immediately reference. After pushing your commit, remember to also push the tag:

```bash
git push origin <tag-name>
```

##### Automatic Tag Pushing

You can configure the hook to automatically push tags by setting an environment variable:

```bash
# Enable automatic tag pushing
export BETHELA_AUTO_PUSH_TAGS=1
```

When this variable is set, the post-commit hook will automatically push the newly created tag to the remote repository, making it immediately available for other projects to use.

You can set up both hooks at once with the provided setup script:

```bash
# Run the setup script
./scripts/setup-hooks.sh
```

This script will make the hook scripts executable and create the appropriate symlinks in your `.git/hooks` directory.

## Versioning

When making changes to the Protocol Buffers definitions, please follow semantic versioning:

- **MAJOR**: Incompatible API changes
- **MINOR**: Add functionality in a backwards compatible manner
- **PATCH**: Backwards compatible bug fixes

After making changes, tag the repository with the new version:

```bash
git tag v1.0.0
git push origin v1.0.0
```

## Usage

To use these definitions in your service, add this repository as a dependency:

```bash
go get github.com/Silicon-Savannah-Solutions/bethela-rpc@v0.0.2
```