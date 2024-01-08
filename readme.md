## Run via Docker

docker run -ti --name monetachain --restart=always -v ~/monetachain:/root/monetachain -v ~/files/repos/moneta/monetachain/bootstrap:/usr/local/bootstrap -p 0.0.0.0:26657:26657 -p 0.0.0.0:26656:26656 -p 0.0.0.0:1317:1317 -p 0.0.0.0:9091:9091 --env-file .env monetachain:1.0 start

## Run commands via Docker

docker exec monetachain /usr/local/bin/monetachaind query bank total-supply


docker run -ti -v $HOME/sdh:/home/tendermint -v $PWD:/apps ignitehq/cli:v28.1.0 scaffold module tokenfactory --dep account,bank

docker run -ti -v $HOME/sdh:/home/tendermint -v $PWD:/apps ignitehq/cli:v28.1.0 scaffold map Denom description:string ticker:string precision:int url:string maxSupply:int supply:int canChangeMaxSupply:bool --signer owner --index denom --module tokenfactory