[![Build Status](https://travis-ci.org/howood/struct2struct.svg?branch=master)](https://travis-ci.org/howood/struct2struct)
[![GitHub release](http://img.shields.io/github/release/howood/struct2struct.svg?style=flat-square)][release]
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/howood/struct2struct)
[![Test Coverage](https://api.codeclimate.com/v1/badges/125c6df22df2c16470ce/test_coverage)](https://codeclimate.com/github/howood/struct2struct/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/howood/struct2struct)](https://goreportcard.com/report/github.com/howood/struct2struct)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[release]: https://github.com/howood/struct2struct/releases
[license]: https://github.com/howood/struct2struct/blob/master/LICENSE

# Struct2Struct

Copy and merge between different structs

# Install

```
$ go get -u github.com/howood/struct2struct
```

# Usage

### Convert Struct To Struct

```
    request := input.Request{
     Name: "name",
     Position: &input.Position {
     ...
     },
     ....
    }

    responseinfo := entity.PlanInfoResponse{}
    position := entity.Position{}
    obj, err = struct2struct.ConvertStructToStruct(request, &responseinfo, "entitytag", "")
    responseinfo = obj.(entity.PlanInfoResponse)
    objsub, err = struct2struct.ConvertStructToStruct(request.Position, &position, "entitytag", "")
    position = objsub.(entity.PlanInfoResponse)
    responseinfo.Position = &position

```

### Merge Struct To Struct

```
    struct1 := input.Register{
     Name: "name",
     ....
    }

    struct2 := input.Update{
     Name: "name",
     ....
    }

    obj, err := struct2struct.MergeStructToStruct(struct1, struct2, "inouttag", "")
    update = obj.(input.Update)
```
