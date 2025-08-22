#!/bin/bash

echo "🚀 Railway Deployment Guide for Weaviate MCP Server"
echo "=================================================="
echo ""

echo "📋 Your GitHub Repository:"
echo "https://github.com/klogins-hash/weaviate-mcp-server-railway"
echo ""

echo "🔧 Railway CLI Commands (already authenticated):"
echo ""

echo "1. Deploy Weaviate Instance:"
echo "   cd weaviate-railway/"
echo "   railway up --detach"
echo ""

echo "2. Deploy MCP Server:"
echo "   cd ../mcp-server-weaviate/"
echo "   railway service create mcp-server"
echo "   railway up --detach"
echo ""

echo "3. Set Environment Variables:"
echo "   railway variables set WEAVIATE_HOST=<your-weaviate-url>"
echo "   railway variables set WEAVIATE_SCHEME=https"
echo ""

echo "4. Get Service URLs:"
echo "   railway status"
echo "   railway domain"
echo ""

echo "🌐 Alternative: Deploy via Railway Dashboard"
echo "1. Go to: https://railway.app/dashboard"
echo "2. New Project → Deploy from GitHub"
echo "3. Select: klogins-hash/weaviate-mcp-server-railway"
echo "4. Create two services:"
echo "   - Service 1: Root = weaviate-railway/"
echo "   - Service 2: Root = mcp-server-weaviate/"
echo ""

echo "✅ Railway CLI is ready and authenticated as: klogins@thekollektiv.xyz"
echo "✅ Project linked: enthusiastic-serenity"
