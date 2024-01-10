#!/bin/sh

BIN=/usr/local/bin/monetachaind
HOME=/root/monetachain
NODE_NAME=validator
KEYRING_SETTINGS="--keyring-backend file --keyring-dir $HOME"

# $BIN export
# exit

if [ ! -f "$HOME/.monetachain/config/genesis.json" ]; then
  rm -R $HOME/.monetachain
  rm -R $HOME/keyring-file    

  $BIN init validator
  cp /usr/local/bootstrap/genesis/genesis.json $HOME/.monetachain/config/genesis.json

  # if [ -n "$MNEMONIC" ]; then
    printf "$MNEMONIC\n$PASSPHRASE\n$PASSPHRASE" | $BIN keys add $NODE_NAME $KEYRING_SETTINGS --coin-type 707 --recover
  # else
  #   $BIN keys add $NODE_NAME $KEYRING_SETTINGS
  # fi

  $BIN genesis gentx $NODE_NAME 50000000000000micromoneta --gas 1000000 --gas-prices 0.1micromoneta $KEYRING_SETTINGS
  $BIN genesis collect-gentxs

  sed -i -e 's/minimum-gas-prices = ""/minimum-gas-prices = "1micromoneta"/' $HOME/.monetachain/config/app.toml
  sed -i -e 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/' $HOME/.monetachain/config/config.toml
fi

$BIN $@