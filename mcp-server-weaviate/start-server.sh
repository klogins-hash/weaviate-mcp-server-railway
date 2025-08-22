#!/bin/bash

# Load environment variables
if [ -f .env ]; then
    export $(cat .env | grep -v '^#' | xargs)
fi

# Start the MCP server
echo "Starting Weaviate MCP Server..."
echo "Host: $WEAVIATE_HOST"
echo "Scheme: $WEAVIATE_SCHEME" 
echo "Port: $PORT"
echo ""

./client/mcp-server
