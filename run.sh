case $1 in
  build)
    go build -o portfolioGenerator cmd/portfoliogenerator/main.go
    ;;
  *)
    echo """
Usage:
  build"""
    ;;
esac