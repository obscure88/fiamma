#!/bin/bash

set -e

if [ $# -ne 1 ]; then
  echo "Usage: $0 <account> " 
  echo "accepts 1 arg, received $#"
  exit 1
else
  ACCOUNT=$1
fi

: ${CHAIN_ID:="fiamma-testnet-1"}
: ${NODE:="https://testnet-rpc.fiammachain.io"}
: ${GAS:=60000000}
: ${FEE:=120000ufia}
: ${PROOF_FILE:=../../prover_examples/bitvm_challenge/proof.bitvm}
: ${PUBLIC_INPUT_FILE:=../../prover_examples/bitvm_challenge/public_input.bitvm}
: ${VK_FILE:=../../prover_examples/bitvm_challenge/vk.bitvm}
: ${NAMESPACE:="TEST"}
: ${PROOF_SYSTEM:="GROTH16_BN254_BITVM"}
: ${DATA_LOCATION:="FIAMMA"}

fiammad tx zkpverify submit-proof \
  --from $ACCOUNT --chain-id $CHAIN_ID  \
  --gas $GAS  \
  --node $NODE \
  --keyring-backend test \
  --fees $FEE \
  $NAMESPACE \
  $PROOF_SYSTEM \
  $PROOF_FILE \
	$PUBLIC_INPUT_FILE \
	$VK_FILE \
	$DATA_LOCATION
