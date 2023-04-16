# meta-egg-layout
MetaEgg Framework Template

## Framework Generator Tool
https://github.com/Jinglever/meta-egg/releases


## Usage

### Prepare
Install protoc:
```bash
wget https://github.com/protocolbuffers/protobuf/releases/download/v21.8/protoc-21.8-linux-x86_64.zip
mkdir -p /usr/local/protoc
sudo unzip protoc-21.8-linux-x86_64.zip -d /usr/local/protoc
sudo ls -s /usr/local/protoc/bin/protoc /usr/local/bin/protoc
```

### Init
```bash
make init
```

### Generate pb/wier/mock
```bash
make generate
```
P.S. Maybe you need to run `go mod tidy` several times.

### Build
```bash
make build
```

### Run
```bash
make run
```
P.S. You need to copy and modify `configs/conf.yml` to `configs/conf-local.yml` before run.