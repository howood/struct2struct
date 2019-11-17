[![Build Status](https://img.shields.io/travis/hjson/hjson-go.svg?style=flat-square)](https://travis-ci.org/howood/struct2struct)
[![Go Pkg](https://img.shields.io/github/release/hjson/hjson-go.svg?style=flat-square&label=go-pkg)](https://github.com/howood/struct2struct/releases)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/howood/struct2struct)

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
