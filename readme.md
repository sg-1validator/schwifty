# schwifty
**schwifty** is a blockchain built using Cosmos SDK and Tendermint

Get SCHWIFTY

https://www.youtube.com/watch?v=I1188GO4p1E

## Get started

Install [go](https://go.dev/dl/)

## Build and install to go bin path

```
make install
```

## Initialize config

Come up with a moniker for your node, then run:

```
schwiftyd init $MONIKER
```

## Available commands

```
Usage:
  schwiftyd tx schwifty [flags]
  schwiftyd tx schwifty [command]

Available Commands:
  create-collection Create a new collection
  create-nft        Create a new nft
  delete-collection Delete a collection by id
  delete-nft        Delete a nft by id
  update-collection Update a collection
  update-nft        Update a nft

```


## Create a NFT collection

```
Use:   "create-collection [name] [description] [ticker] [uri] [uri-hash] [data]"
Short: "Create a new collection"
```

## Create new NFT within a collection

```
Use:   "create-nft [collection-id] [uri] [uri-hash] [data]"
Short: "Create a new nft"
``` 
 
## Launch with genesis file or run as standalone chain

To launch as a consumer chain, download and save shared genesis file to `~/.schwifty/config/genesis.json`. Additionally add peering information (`persistent_peers` or `seeds`) to `~/.schwifty/config/config.toml`

To instead launch as a standalone, single node chain, run:

```
schwiftyd add-consumer-section
```

## Launch node

```
schwiftyd start
```
