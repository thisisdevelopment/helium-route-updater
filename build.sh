#!/bin/bash

set -eu
package="github.com/thisisdevelopment/helium-route-updater/pkg"
dir="api/helium"
go_opts="--go_opt=module=${package}"
grpc_opts="--go-grpc_opt=module=${package}"
files=""
for file in $(find ${dir} -not -path "${dir}/service/*" -name "*.proto"); do
  file=${file##$dir/}
  files="${files} ${file}"
  subdir=$(dirname $file)
  file_package=${package}/${dir}
  if [ "${subdir}" != "." ]; then
    file_package="${file_package}/${subdir}"
  fi
  go_opts="${go_opts} --go_opt=M${file}=${file_package}"
  grpc_opts="${grpc_opts} --go-grpc_opt=M${file}=${file_package}"
done
protoc -I=$dir --go_out=pkg --go-grpc_out=pkg ${go_opts} ${grpc_opts} ${files}

service_files=""
for file in ${dir}/service/*.proto; do
  service_file=${file##$dir/service/}
  service_name=${service_file%%.proto}
  service_files="${service_files} service/${service_file}"
  go_opts="${go_opts} --go_opt=Mservice/${service_file}=${package}/${dir}/service/${service_name}"
  grpc_opts="${grpc_opts} --go-grpc_opt=Mservice/${service_file}=${package}/${dir}/service/${service_name} "
done
protoc -I=$dir --go_out=pkg --go-grpc_out=pkg ${go_opts} ${grpc_opts} ${service_files}

for cmd in cmd/*; do
  if [ -d "${cmd}" -a -f "${cmd}/main.go" ]; then
    go build -o ./bin/ "./${cmd}"
  fi
done
