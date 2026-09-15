#!/bin/bash
# Test runner script

set -e

echo "🧪 Running Football Overlay Tests..."
echo ""

# Unit tests
echo "📋 Unit Tests:"
go test ./internal/api/... -v

echo ""
echo "📊 Integration Tests:"
go test ./tests/... -v

echo ""
echo "✅ All tests passed!"
