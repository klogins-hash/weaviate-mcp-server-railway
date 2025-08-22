# Weaviate Docker configuration optimized for Railway
FROM semitechnologies/weaviate:1.27.0

# Create weaviate user and data directory
USER root
RUN mkdir -p /var/lib/weaviate && chown -R weaviate:weaviate /var/lib/weaviate

# Railway-specific environment variables
ENV QUERY_DEFAULTS_LIMIT=25
ENV AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED=true
ENV PERSISTENCE_DATA_PATH=/var/lib/weaviate
ENV DEFAULT_VECTORIZER_MODULE=none
ENV ENABLE_MODULES=""
ENV CLUSTER_HOSTNAME=node1
ENV ORIGIN=*
ENV LOG_LEVEL=info

# Switch back to weaviate user
USER weaviate

# Expose port for Railway (Railway will override with PORT env var)
EXPOSE 8080

# Start Weaviate with proper configuration
CMD ["/bin/weaviate", "--host", "0.0.0.0", "--port", "8080", "--scheme", "http"]
