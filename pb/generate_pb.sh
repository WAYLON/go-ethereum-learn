#! /bin/bash

echo "----------------------------------------"
echo "start generate pb file"

dirs=("v1")
module="ChainCollect/pb"

for d in ${dirs[@]};
do
    echo "--------------------"
    echo "deal dir: $d"
    if [ ! -d ${d} ]; then
        echo "not exist dir: $d, continue"
        echo "----------"
        continue
    fi
    rm -f $d/*.json
    rm -f $d/*.go

    echo "----------"
    echo "generate golang code"
    protoc -I . -I ../third_party  --go_out=./ --go_opt=module="$module" --go-grpc_out . --go-grpc_opt=module="$module" --validate_out=lang=go:. --validate_opt=module="$module" --grpc-gateway_out . --grpc-gateway_opt=logtostderr=true,module=$module,generate_unbound_methods=true --openapiv2_out=${d} --openapiv2_opt=logtostderr=true,allow_merge=true,openapi_naming_strategy=fqn ${d}/pbfile/*.proto
    echo "--------------------"
done

echo "end generate pb file"
echo "----------------------------------------"

