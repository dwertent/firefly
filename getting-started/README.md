# Custom contract getting started guide

### Compile 
```
solc --evm-version paris --combined-json  abi,bin simple_storage.sol > simple_storage.json
```

### Deploy
There are two ways for deploying, CLI and REST API

#### CLI
```
firefly deploy ethereum <<stack_name>> <compiled solidity contract in json format>  
```

e.g.
```
firefly deploy ethereum dev simple_storage.json
```
save the address
0xae1b5868001ec5cbe1ba6005746578cbb1a680ec

#### REST API
https://hyperledger.github.io/firefly/latest/tutorials/custom_contracts/ethereum/#contract-deployment



### Interface

#### Generate:

```
wget --method=POST --header="Content-Type: application/json" --body-file="generate-interface.json" http://localhost:5000/api/v1/namespaces/default/contracts/interfaces/generate -O generate-interface-response.json
```

#### Broadcast
```
wget --method=POST --header="Content-Type: application/json" --body-file="broadcast-interface.json" "http://localhost:5000/api/v1/namespaces/default/contracts/interfaces?publish=true" -O broadcast-interface-response.json
```

### API

```
wget --method=POST --header="Content-Type: application/json" --body-file="register-api.json" "http://localhost:5000/api/v1/namespaces/default/apis?publish=true" -O register-api-response.json
```

Try a curl


### Set/Get

#### Set

```
wget --method=POST --header="Content-Type: application/json" --body-data='{"input":{"newValue":3}}' "http://localhost:5000/api/v1/namespaces/default/apis/simple-storage-2/invoke/set" -O set-response.json
```

#### Get

```
wget --method=POST --header="Content-Type: application/json" --body-data='{}' "http://localhost:5000/api/v1/namespaces/default/apis/simple-storage-2/query/get" -O get-response.json
```