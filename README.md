<!--
parent:
  order: false
-->
<img align="right" width="150" height="150" top="100" src="./assets/readme.jpg">

# [----------] • [![tests](https://github.com/abigger87/femplate/actions/workflows/tests.yml/badge.svg)](https://github.com/abigger87/femplate/actions/workflows/tests.yml) [![lints](https://github.com/abigger87/femplate/actions/workflows/lints.yml/badge.svg)](https://github.com/abigger87/femplate/actions/workflows/lints.yml) ![GitHub](https://img.shields.io/github/license/abigger87/femplate)  ![GitHub package.json version](https://img.shields.io/github/package-json/v/abigger87/femplate)


[----------] is a scalable, high-throughput Proof-of-Liquidity blockchain that is fully EVM & IBC compatible.

It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/) which runs on top of [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine. **Note**: Requires [Go 1.17.5+](https://golang.org/dl/)

## Installation
### For Node Operators
Install the Node Binary to your system using the provided [setup tool](https://github.com/berachain/osmosis-installer)

### For Developers
For prerequisites and detailed build instructions please read the [Installation](https://evmos.dev/quickstart/installation.html) instructions. Once the dependencies are installed, build an install a node binary:

```bash
make install
```
Or check out the latest [release](https://github.com/tharsis/evmos/releases).

To compile and run a local node in one command:
```bash
./init.sh
```

## Code Structure

```ml
app
├─ Entrypoint for the Node Software
cmd
├─ Setup and management type Daemon commands
contracts
├─ lib
│  └─ ds-test — https://github.com/dapphub/ds-test
│  └─ forge-std — https://github.com/brockelmore/forge-std
│  └─ solmate — https://github.com/Rari-Capital/solmate
├─ compilied_contracts
│  └─ contracts used for module tests
├─ out  
│  └─ ABIs and Bytecoede for Solidity contracts
├─ src
│  └─ core contracts
crypto
├─ keyring
│  └─ signing utilities 
docs
├─ full node documentation
ibctesting
├─ testing utils for IBC related tests
proto
├─ amm
├─ amo
├─ evmos
│  └─ claims
│  └─ epochs
│  └─ erc20
│  └─ incentives
│  └─ inflation
│  └─ vesting
├─ synapse
scripts
├─ testing + deployment scripts
testutil
├─ keeper
├─ network
├─ nullify
├─ sample
third_party
├─ proto
│  └─ cosmos
│  └─ cosmos_proto
│  └─ ethermint
│  └─ gogoproto
│  └─ vesting
│  └─ tendermint
version
├─ track node binary version
x
├─ amm - core berachain automated market maker
├─ amo - stablecoin automated market operations controller
├─ claims - claim airdrops at launch
├─ epochs - epochs for periodic execution of logic
├─ erc20 - relay tokens between ICS20 and ERC20 formats
├─ incentives - incentivize user behaviour
├─ inflation - manage fee token inflation rates
├─ synapse - native Synapse Protocol integration to brige tokens from other EVM chains
```

## Community

The following chat channels and forums are a great spot to ask questions about [----------]:

- [[----------] Twitter](https://twitter.com/EvmosOrg)
- [[----------] Discord](https://discord.gg/evmos)
- [[----------] Forum](https://forum.cosmos.network/c/ethermint)
- [[----------] Twitter](https://twitter.com/TharsisHQ)

## Contributing

Looking for a good place to start contributing? Check out some [`good first issues`](https://github.com/[----------]/[----------]/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22).

For additional instructions, standards and style guides, please refer to the [Contributing](./CONTRIBUTING.md) document.

## Acknowledgements

- [foundry](https://github.com/gakonst/foundry)
- [solmate](https://github.com/Rari-Capital/solmate)
- [forge-std](https://github.com/brockelmore/forge-std)
- [forge-template](https://github.com/FrankieIsLost/forge-template) by [FrankieIsLost](https://github.com/FrankieIsLost).
- [Evmos](https://github.com/tharsis/evmos)
- [Georgios Konstantopoulos](https://github.com/gakonst) for [forge-template](https://github.com/gakonst/forge-template) resource.

## Disclaimer

_These smart contracts and node software are being provided as is. No guarantee, representation or warranty is being made, express or implied, as to the safety or correctness of the user interface or the smart contracts. They have not been audited and as such there can be no assurance they will work as intended, and users may experience delays, failures, errors, omissions, loss of transmitted information or loss of funds. The creators are not liable for any of the foregoing. Users should proceed with caution and use at their own risk._
