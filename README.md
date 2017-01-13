restgen
=======

[![Build Status](https://travis-ci.org/doozer-de/restgen.svg?branch=master)](https://travis-ci.org/doozer-de/restgen)
[![GoDoc](https://godoc.org/github.com/doozer-de/restgen?status.svg)](https://godoc.org/github.com/doozer-de/restgen)

restgen is a plugin for [protocol buffer compiler](https://github.com/google/protobuf).
It generates go code for REST endpoints, based on [Google's Protocol Buffers](https://github.com/google/protobuf) (a.k.a. protobuf).

Installation
------------

This tool is written in go. You need a working go installation to install
it. Execute the following command in a terminal:

```bash
go install bitbucket.org/doozer-de/restgen/protoc-gen-bff
```

This installs the protobuf compiler plugin `protoc-gen-bff` (where "bff"
is an abbreviation for "backend for frontend"). It must be in your `$PATH`
for the protocol compiler to find it.

Usage
-----

1. Define a service in your protobuf. Every service will be defined with a service map and rpc methods.

A service map definition looks like the following:

```protobuf
	option (pbmap.service_map) = {
		version: "1"
		base_uri: "/base_uri/"
		target_package: "samplesvc"
	};
```

This service map is defined in pbmap/map.proto:

- `version`:
- `base_uri`: defines the base URI, e. g. `http://[address]/[base_uri]/[path]...`
- `target_package`: defines the name of the package, where the generated
  file will be placed. This option doesn't affect the location of the
  file - this has to be defined with `OUT_DIR` in the protoc command operand
  `--bff_out=OUT_DIR`

A rpc method would be defined like the following:

```protobuf
	rpc Methodname(RequestMethod) returns (ResponseMethod) {
		option (pbmap.method_map) = {
            method: "method"
            path: "/[path]/:id"
            query_string: "key=[message].[field]:[data_type]"
            body: "*"
        };
    }
```

The method map is defined in pbmap/map.proto.

Explanation:

- `method`: defines the used HTTP method, supported are: Get, Put, Post, Delete
- `path`: identifies the method, has to be unique in this service and
  has to be compatible with Julian Schmidt's HTTP Router; Format `a/b/:c/:d`
  where `a/b/` are fixed and `:c` and `:d` are parameters mapped to C and D in
  the underlying GRPC Request
- `query_string`: Format "key1=:var1&key=:var2": they keys will be mapped
  to the variables `var1`, `var2` of the unterlying GRPC Request
- `body`: indicates, if the request has a body or not

2. [The messages are defined as usual](https://developers.google.com/protocol-buffers/docs/proto#simple)

3. [Compile the protobuf](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers).

4. Execute the following command in a terminal:

```bash
protoc [OPTION] --bff_out=OUT_DIR PROTO_FILES
```

Features
--------

## Supported

- atomic datatypes
- maps
- user defined messages; inside of the same package and outside of this package)
- HTTP Methods: Get, Put, Post, Delete
- path parameter; also with deep paths
- query strings; also with deep paths

## Not supported

- enums in path parameter and query strings

Example
-------

You can find a sample protobuf definition in
`samplesvc/pb/samplesvc.proto` and the generated code in
`samplesvc/samplesvc_gen.go`.


