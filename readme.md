# versioner -  versioning tool to help at releasing projects

[![Build](https://github.com/cbuschka/versioner/workflows/build/badge.svg)](https://github.com/cbuschka/versioner) [![Go Report Card](https://goreportcard.com/badge/github.com/cbuschka/versioner)](https://goreportcard.com/report/github.com/cbuschka/versioner) [![codecov](https://codecov.io/gh/cbuschka/versioner/branch/main/graph/badge.svg)](https://codecov.io/gh/cbuschka/versioner) [![License](https://img.shields.io/github/license/cbuschka/versioner.svg)](https://github.com/cbuschka/versioner/blob/main/license.txt)

## Prerequisites
* go 1.16.x
* GNU make
* git

## Usage

### Get latest version by git tags
```
versioner latest-version
```

### Get next version for latest by git tags
```
versioner next-version
```

### Get versioner build version info
```
versioner version
```

## Build

```
make build
```

## License
Copyright (c) 2020-2021 by [Cornelius Buschka](https://github.com/cbuschka).

[Apache License, Version 2.0](./license.txt)
