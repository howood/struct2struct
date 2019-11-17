[![Build Status](https://travis-ci.org/howood/struct2struct.svg?branch=master)](https://travis-ci.org/howood/struct2struct)
[![GitHub release](http://img.shields.io/github/release/howood/struct2struct.svg?style=flat-square)][release]
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/howood/struct2struct)
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
    responseinfo = struct2struct.ConvertStructToStruct(request, &responseinfo, "entitytag", "").(entity.PlanInfoResponse)
    position = struct2struct.ConvertStructToStruct(request.Position, &position, "entitytag", "").(entity.Position)
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

    update := struct2struct.MergeStructToStruct(struct1, struct2, "inouttag", "").(input.Update)

```
