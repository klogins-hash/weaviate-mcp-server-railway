# Weaviate Docker configuration optimized for Railway
FROM semitechnologies/weaviate:1.27.0

# Railway-specific environment variables will be set via Railway dashboard
# Default configurations for Railway deployment
ENV QUERY_DEFAULTS_LIMIT=25
ENV AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED=true
ENV PERSISTENCE_DATA_PATH=/var/lib/weaviate
ENV DEFAULT_VECTORIZER_MODULE=none
ENV ENABLE_MODULES=""
ENV CLUSTER_HOSTNAME=node1
ENV ORIGIN=*

# Expose port for Railway
EXPOSE 8080

# Use Railway's PORT if provided, otherwise default to 8080
CMD ["./weaviate", "--host", "0.0.0.0", "--port", "8080", "--scheme", "http"]
