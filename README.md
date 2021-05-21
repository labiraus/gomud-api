# gomud-api

This repo reprepresents an externally facing json REST api to allow tool to access and manipulate the gomud space from outside of the cluster

## Build

The steps in the build chain are:

* Download proto files from gomud-common/proto
* Build proto files into go
* Go generate
* Go test
* Build into a docker container


## Next steps

Build a docker container that will pull the gomud-common repo, get the code files required and 