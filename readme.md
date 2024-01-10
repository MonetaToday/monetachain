## Run via Docker

docker build ./ --tag monetachain:1.0

docker run -ti --name monetachain --rm -v ~/monetachain:/root/monetachain -v ~/files/repos/moneta/monetachain/bootstrap:/usr/local/bootstrap -p 0.0.0.0:26657:26657 -p 0.0.0.0:26656:26656 -p 0.0.0.0:1317:1317 -p 0.0.0.0:9091:9091 --env-file .env monetachain:1.0 start

docker run -ti --name monetachain --rm -v ~/monetadb:/root/monetachain -v ~/monetachain/bootstrap:/usr/local/bootstrap -p 0.0.0.0:26657:26657 -p 0.0.0.0:26656:26656 -p 0.0.0.0:1317:1317 -p 0.0.0.0:9091:9091 --env-file .env monetachain:1.0 start

## Run commands via Docker

docker exec monetachain /usr/local/bin/monetachaind query bank total-supply


docker run -ti -v $HOME/sdh:/home/tendermint -v $PWD:/apps ignitehq/cli:v28.1.0 scaffold module tokenfactory --dep account,bank

docker run -ti -v $HOME/sdh:/home/tendermint -v $PWD:/apps ignitehq/cli:v28.1.0 scaffold map Denom description:string ticker:string precision:int url:string maxSupply:int supply:int canChangeMaxSupply:bool --signer owner --index denom --module tokenfactory

docker run -ti -v $HOME/sdh:/home/tendermint -v $PWD:/apps ignitehq/cli:v28.1.0 scaffold message MintAndSendTokens denom:string amount:int recipient:string --module tokenfactory --signer owner

docker run -ti -v $HOME/sdh:/home/tendermint -v $PWD:/apps ignitehq/cli:v28.1.0 scaffold message UpdateOwner denom:string newOwner:string --module tokenfactory --signer owner


docker exec -ti monetachain /usr/local/bin/monetachaind tx tokenfactory create-denom uignite "My denom" IGNITE 6 "some/url" 3000000000 true --from validator --keyring-backend file --keyring-dir /root/monetachain --chain-id monetachain --fees 200000micromoneta

docker exec -ti monetachain /usr/local/bin/monetachaind tx tokenfactory update-denom uignite "Ignite" "newurl" 2000000000000000000 false --from validator --keyring-backend file --keyring-dir /root/monetachain --chain-id monetachain --fees 200000micromoneta

docker exec -ti monetachain /usr/local/bin/monetachaind tx tokenfactory mint-and-send-tokens uignite 2000000000000000000 moneta1r4rfvqyjj0nycretax744krvrwlh404lp30083 --from validator --keyring-backend file --keyring-dir /root/monetachain --chain-id monetachain --fees 200000micromoneta

docker exec -ti monetachain /usr/local/bin/monetachaind tx tokenfactory update-owner uignite moneta1r4rfvqyjj0nycretax744krvrwlh404lp30083 --from validator --keyring-backend file --keyring-dir /root/monetachain --chain-id monetachain --fees 200000micromoneta