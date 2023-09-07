#!/usr/bin/env bash
set -ex
SOURCE=${BASH_SOURCE[0]}
DIR_OF_THIS_SCRIPT="$( dirname "$SOURCE" )"
ABS_DIR_OF_SCRIPT="$( realpath $DIR_OF_THIS_SCRIPT )"
mkdir -p ~/.vmodules/threefoldtech

ln -s $ABS_DIR_OF_SCRIPT/web3gw/client ~/.vmodules/threefoldtech/web3gw
ln -s $ABS_DIR_OF_SCRIPT/3bot/vgrid ~/.vmodules/threefoldtech/vgrid
ln -s $ABS_DIR_OF_SCRIPT/3bot/zos ~/.vmodules/threefoldtech/zos
ln -s $ABS_DIR_OF_SCRIPT/v2go2tgrid/tfgrid ~/.vmodules/threefoldtech/tfgrid

# install crystallib
if !(v list | grep -q 'freeflowuniverse.crystallib'); then
    git clone https://github.com/freeflowuniverse/crystallib.git ~/.vmodules/freeflowuniverse/crystallib
fi
