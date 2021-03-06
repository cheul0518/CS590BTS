### Hyperledger Fabric 1.1 SetUp: Mac

0. Website: https://hyperledger-fabric.readthedocs.io/en/release-1.1/getting_started.html

1. Prerequsites

   1. cURL: it's already installed if you're using Mac. Skip this process
   <br />
   
   2. Docker and Docker Compose: https://hub.docker.com/editions/community/docker-ce-desktop-mac
   - Don't forget to check its version after you've installed Docker: "docker --version"
   - Check the version of Docker Compose: "docker-compose --version"
   <br />
   
   3. Go Programming Language: https://golang.org/doc/install
   - After you've installed Go package, check the $GOPATH variable: it defaults to a directory named "go" insdie your home directory (e.g. /Users/USERNAME/go/). By the way you can check it by typing "go env" and looking up the enviroment variable there
   - Add the workstation's bin subdirectory to your path: "export PATH=$PATH:$(go env GOPATH)/bin" (for convenience)
   - Type "export GOPATH=$(go env GOPATH)". So you can now use $GOPATH instead of $(go env GOPATH) (for convenience)
   - Feel free to check your GO by compiling and running a simple program given on the website.
   <br />

   4. Node.js Runtime and NPM: https://nodejs.org/en/download/
   - Check Node.js version: "node -v"
   - Check npm version: "npm -v"
   <br />
   
   5. Python: it's already installed if you're using Mac. Skip this process
   <br />
   
2. Install Binaries and Docker Images
```c
curl -sSL https://goo.gl/6wtTN5 | bash -s <fabric> <fabric-ca> <thirdparty>
curl -sSL https://goo.gl/6wtTN5 | bash -s 1.1.0 1.1.0 0.4.6
```
   - You may want to add that to your PATH environment variables: "export PATH=<path to download location>/bin:$PATH"
   <br />
   
3. Hyperledger Fabric Samples
```c
git clone -b master https://github.com/hyperledger/fabric-samples.git
cd fabric-samples
git checkout master
```
   
