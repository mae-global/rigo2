#!/bin/sh
echo "building proto buffers from ri.proto..."
protoc --go_out=. ri.proto