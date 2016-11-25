set -ex

docker run --rm -v "$GOPATH":/go -w /go/src/github.com/amaysim-au/logspout -e "GOPATH=/go" iron/go:dev sh -c 'go build -o logspout'

# Can build the image too
# docker build -t amaysim-au/logspout:latest .
