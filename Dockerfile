FROM centos
WORKDIR /app
COPY . .
CMD ["./RedEnv_Simple"]