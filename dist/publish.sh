#!/usr/bin/env bash
# brew install go
# go get ./...
# for docs update, uncomment generator in root.go

export VERSION=2.01

if [ -z "${AWS_ACCESS_KEY}" ]; then
      echo "AWS_ACCESS_KEY is empty, aborting"
      exit 1
fi

if [ -z "${AWS_SECRET}" ]; then
      echo "AWS_SECRET is empty, aborting"
      exit 1
fi

echo "Building..."
cd ..
make build_all
cd dist

echo "Uploading to S3..."
s3cmd put --access_key=${AWS_ACCESS_KEY} --secret_key=${AWS_SECRET} *.tgz s3://lokalise-assets/cli2/

echo "Readme update..."
sed -i.bu "19s/.*/- [lokalise2-${VERSION}-darwin-amd64.tgz](https:\/\/s3-eu-west-1.amazonaws.com\/lokalise-assets\/cli2\/lokalise2-${VERSION}-darwin-amd64.tgz)/g" ../README.md
sed -i.bu "20s/.*/- [lokalise2-${VERSION}-darwin-386.tgz](https:\/\/s3-eu-west-1.amazonaws.com\/lokalise-assets\/cli2\/lokalise2-${VERSION}-darwin-386.tgz)/g" ../README.md
sed -i.bu "23s/.*/- [lokalise2-${VERSION}-linux-amd64.tgz](https:\/\/s3-eu-west-1.amazonaws.com\/lokalise-assets\/cli2\/lokalise2-${VERSION}-linux-amd64.tgz)/g" ../README.md
sed -i.bu "24s/.*/- [lokalise2-${VERSION}-linux-386.tgz](https:\/\/s3-eu-west-1.amazonaws.com\/lokalise-assets\/cli2\/lokalise2-${VERSION}-linux-386.tgz)/g" ../README.md
sed -i.bu "27s/.*/- [lokalise2-${VERSION}-freebsd-amd64.tgz](https:\/\/s3-eu-west-1.amazonaws.com\/lokalise-assets\/cli2\/lokalise2-${VERSION}-freebsd-amd64.tgz)/g" ../README.md
sed -i.bu "28s/.*/- [lokalise2-${VERSION}-freebsd-386.tgz](https:\/\/s3-eu-west-1.amazonaws.com\/lokalise-assets\/cli2\/lokalise2-${VERSION}-freebsd-386.tgz)/g" ../README.md
sed -i.bu "31s/.*/- [lokalise2-${VERSION}-windows-amd64.tgz](https:\/\/s3-eu-west-1.amazonaws.com\/lokalise-assets\/cli2\/lokalise2-${VERSION}-windows-amd64.tgz)/g" ../README.md
sed -i.bu "32s/.*/- [lokalise2-${VERSION}-windows-386.tgz](https:\/\/s3-eu-west-1.amazonaws.com\/lokalise-assets\/cli2\/lokalise2-${VERSION}-windows-386.tgz)/g" ../README.md
rm -f ../README.md.bu
cd ..
git add .
git commit -m ${VERSION}
git push
cd dist

echo "Dockerhub update..."
rm -f ../../../../../../lokalise-cli-2/lokalise2
tar zxvf lokalise2-${VERSION}-linux-amd64.tgz -C ../../../../../../lokalise-cli-2
cd ../../../../../../lokalise-cli-2
git add .
git commit -m ${VERSION}
git push
cd ../go/src/github.com/lokalise/lokalise-cli-2-go/dist

echo "Homebrew update..."
SHA=$(shasum -a 256 lokalise2-${VERSION}-darwin-amd64.tgz | awk '{print $1}')
sed -i.bu "5s/.*/  version \"${VERSION}\"/g" ../../../../../../homebrew-cli-2/Formula/lokalise2.rb
sed -i.bu "6s/.*/  sha256 \"${SHA}\"/g" ../../../../../../homebrew-cli-2/Formula/lokalise2.rb
sed -i.bu "8s/.*/  url \"https:\/\/s3-eu-west-1.amazonaws.com\/lokalise-assets\/cli2\/lokalise2-${VERSION}-darwin-amd64.tgz\"/g" ../../../../../../homebrew-cli-2/Formula/lokalise2.rb
rm -f ../../../../../../homebrew-cli-2/Formula/*.bu
cd ../../../../../../homebrew-cli-2
git add .
git commit -m ${VERSION}
git push
cd ../go/src/github.com/lokalise/lokalise-cli-2-go/dist
