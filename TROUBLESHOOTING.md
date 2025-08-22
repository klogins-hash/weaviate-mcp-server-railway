# Railway Deployment Troubleshooting Guide

## Issues Found & Solutions

### 1. **CLI Deployment Issues**
**Problem**: Railway CLI deployments failing with no logs
**Root Cause**: Service configuration mismatch and Docker build issues

**Solution**: Use GitHub-based deployment instead of CLI uploads

### 2. **Weaviate Docker Configuration**
**Problem**: Original Dockerfile had incorrect binary path and permissions
**Fixed**: 
- ✅ Updated binary path to `/bin/weaviate`
- ✅ Fixed user permissions and data directory
- ✅ Added nixpacks.toml for Railway PORT handling

### 3. **Service URLs**
**Current Status**:
- Domain created: `https://weaviate-mcp-server-railway-production.up.railway.app`
- Status: Building/Deploying

## Working Deployment Method

### **Recommended: GitHub Integration**

1. **Go to Railway Dashboard**: https://railway.app/dashboard
2. **New Project** → **Deploy from GitHub repo**
3. **Select Repository**: `klogins-hash/weaviate-mcp-server-railway`
4. **Create Services**:

#### Service 1: Weaviate Database
- **Name**: `weaviate`
- **Root Directory**: `weaviate-railway/`
- **Build**: Uses Dockerfile automatically
- **Environment**: Set via railway.toml

#### Service 2: MCP Server  
- **Name**: `mcp-server`
- **Root Directory**: `mcp-server-weaviate/`
- **Environment Variables**:
  ```
  WEAVIATE_HOST=weaviate.railway.internal
  WEAVIATE_SCHEME=http
  PORT=8675
  ```

## Current Status

✅ **Repository**: Updated with fixes
✅ **CLI**: Authenticated and linked
✅ **Domain**: Generated for service
⏳ **Deployment**: In progress via GitHub integration

## Next Steps

1. **Monitor Build**: Check Railway dashboard for build progress
2. **Test Endpoints**: Once deployed, test:
   - Weaviate: `/v1/.well-known/ready`
   - MCP Server: `/health`
3. **Configure Variables**: Set WEAVIATE_HOST to internal service URL

## Railway CLI Commands (Backup)

```bash
# Check status
railway status

# View logs (when deployment completes)
railway logs

# Get service URL
railway domain

# Set environment variables
railway variables set WEAVIATE_HOST=weaviate.railway.internal
```
