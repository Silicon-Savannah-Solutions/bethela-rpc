#!/bin/bash
set -e

# Directory structure
PROTO_DIR="./grpc"
GO_OUT_DIR="./gen/go"
DOC_OUT_DIR="./docs"

# Script header
echo "=========================================="
echo "Bethela RPC Protocol Buffers Compiler"
echo "=========================================="

# Create output directories if they don't exist
mkdir -p $GO_OUT_DIR
mkdir -p $DOC_OUT_DIR

# Check if protoc is installed
if ! command -v protoc &> /dev/null; then
    echo "Error: protoc is not installed."
    echo "Please install Protocol Buffers compiler (protoc) from https://github.com/protocolbuffers/protobuf/releases"
    exit 1
fi

# Check if protoc-gen-go and protoc-gen-go-grpc are installed
if ! command -v protoc-gen-go &> /dev/null || ! command -v protoc-gen-go-grpc &> /dev/null; then
    echo "Installing protoc-gen-go and protoc-gen-go-grpc..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    
    # Verify installation
    if ! command -v protoc-gen-go &> /dev/null || ! command -v protoc-gen-go-grpc &> /dev/null; then
        echo "Error: Failed to install protoc-gen-go or protoc-gen-go-grpc."
        echo "Please ensure your GOPATH/bin is in your PATH."
        exit 1
    fi
fi

echo "Compiling Protocol Buffers..."

# Get the total number of proto files
TOTAL_FILES=$(find $PROTO_DIR -name "*.proto" | wc -l)
CURRENT=0

# Find all proto files and compile them
for proto_file in $(find $PROTO_DIR -name "*.proto"); do
    CURRENT=$((CURRENT + 1))
    echo "[$CURRENT/$TOTAL_FILES] Compiling: $proto_file"

    # Generate Go code
    protoc --proto_path=$PROTO_DIR \
        --go_out=$GO_OUT_DIR --go_opt=paths=source_relative \
        --go-grpc_out=$GO_OUT_DIR --go-grpc_opt=paths=source_relative \
        $proto_file

    # Generate documentation (optional)
    # Uncomment if you have protoc-gen-doc installed
    # protoc --proto_path=$PROTO_DIR \
    #     --doc_out=$DOC_OUT_DIR --doc_opt=markdown,$(basename $proto_file .proto).md \
    #     $proto_file
done

echo ""
echo "Compilation complete! ✓"
echo "Generated files in $GO_OUT_DIR"
echo "Note: Remember to tag a new version after making changes to proto files"
echo "=========================================="
