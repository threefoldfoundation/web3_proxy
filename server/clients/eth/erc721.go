package goethclient

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/threefoldtech/web3_proxy/server/clients/eth/erc721"
)

// GetFungibleBalance returns the balance of the given address for the given fungible token contract
func (c *Client) GetFungibleBalance(contractAddress, target string) (*big.Int, error) {
	fungible, err := erc721.NewErc721(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return nil, err
	}

	return fungible.BalanceOf(&bind.CallOpts{}, common.HexToAddress(target))
}

// ownerOfFungible returns the owner of the given fungible token
func (c *Client) OwnerOfFungible(contractAddress string, token int64) (string, error) {
	fungible, err := erc721.NewErc721(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	owner, err := fungible.OwnerOf(&bind.CallOpts{}, big.NewInt(token))
	if err != nil {
		return "", err
	}

	return owner.Hex(), nil
}

// SafeTransferFungible transfers a fungible token from the given address to the given target address
func (c *Client) SafeTransferFungible(contractAddress, from, to string, tokenId int64) (string, error) {
	fungible, err := erc721.NewErc721(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	tx, err := fungible.SafeTransferFrom(&bind.TransactOpts{}, common.HexToAddress(from), common.HexToAddress(to), big.NewInt(tokenId))
	if err != nil {
		return "", err
	}

	return c.sendTransaction(tx)
}

// Transfer transfers the given fungible token from the given address to the given target address
func (c *Client) TransferFungible(contractAddress, from, to string, tokenId int64) (string, error) {
	fungible, err := erc721.NewErc721(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	tx, err := fungible.TransferFrom(&bind.TransactOpts{}, common.HexToAddress(from), common.HexToAddress(to), big.NewInt(tokenId))
	if err != nil {
		return "", err
	}

	return c.sendTransaction(tx)
}

// SetFungibleApproval approves the given address to spend the given amount of the given fungible token
func (c *Client) SetFungibleApproval(contractAddress, from, to string, amount int64) (string, error) {
	fungible, err := erc721.NewErc721(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	tx, err := fungible.Approve(&bind.TransactOpts{}, common.HexToAddress(to), big.NewInt(amount))
	if err != nil {
		return "", err
	}

	return c.sendTransaction(tx)
}

// SetFungibleApprovalForAll approves the given address to spend all the given fungible token
func (c *Client) SetFungibleApprovalForAll(contractAddress, from, to string, approved bool) (string, error) {
	fungible, err := erc721.NewErc721(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return "", err
	}

	tx, err := fungible.SetApprovalForAll(&bind.TransactOpts{}, common.HexToAddress(to), approved)
	if err != nil {
		return "", err
	}

	return c.sendTransaction(tx)
}

// GetApprovalForFungible returns the approval status of the given address for the given fungible token
func (c *Client) GetApprovalForFungible(contractAddress, owner, operator string) (bool, error) {
	fungible, err := erc721.NewErc721(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return false, err
	}

	return fungible.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(owner), common.HexToAddress(operator))
}

// GetApprovalForAllFungible returns the approval status of the given address for the given fungible token
func (c *Client) GetApprovalForAllFungible(contractAddress, owner, operator string) (bool, error) {
	fungible, err := erc721.NewErc721(common.HexToAddress(contractAddress), c.Eth)
	if err != nil {
		return false, err
	}

	return fungible.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(owner), common.HexToAddress(operator))
}
