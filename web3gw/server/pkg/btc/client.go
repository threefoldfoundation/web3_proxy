package btc

import (
	"context"

	"github.com/LeeSmet/go-jsonrpc"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	btcRpcClient "github.com/btcsuite/btcd/rpcclient"
	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/3bot/web3gw/server/pkg"
)

const (
	// NostrID is the ID for state of a btc client in the connection state.
	BtcID = "btc"
)

type (
	// Client exposes nostr related functionality
	Client struct {
	}
	// state managed by nostr client
	btcState struct {
		client *btcRpcClient.Client
	}

	Load struct {
		Host string `json:"host"`
		User string `json:"user"`
		Pass string `json:"pass"`
	}

	ImportAddressRescan struct {
		Address string `json:"address"`
		Account string `json:"account"`
		Rescan  bool   `json:"rescan"`
	}

	ImportPrivKeyLabel struct {
		WIF   string `json:"wif"`
		Label string `json:"label"`
	}

	ImportPrivKeyRescan struct {
		WIF    string `json:"wif"`
		Label  string `json:"label"`
		Rescan bool   `json:"rescan"`
	}

	ImportPubKeyRescan struct {
		PubKey string `json:"pub_key"`
		Rescan bool   `json:"rescan"`
	}

	RenameAccount struct {
		OldAccount string `json:"old_account"`
		NewAccount string `json:"new_account"`
	}

	SendToAddress struct {
		Address   string         `json:"address"`
		Amount    btcutil.Amount `json:"amount"`
		Comment   string         `json:"comment"`
		CommentTo string         `json:"comment_to"`
	}

	EstimateSmartFee struct {
		ConfTarget int64                        `json:"conf_target"`
		Mode       btcjson.EstimateSmartFeeMode `json:"mode"`
	}

	GenerateToAddress struct {
		NumBlocks int64  `json:"num_blocks"`
		Address   string `json:"address"`
		MaxTries  int64  `json:"max_tries"`
	}

	GetChainTxStatsNBlocksBlockHash struct {
		AmountOfBlocks int32  `json:"amount_of_blocks"`
		BlockHashEnd   string `json:"block_hash_end"`
	}

	CreateWallet struct {
		Name               string `json:"name"`
		DisablePrivateKeys bool   `json:"disable_private_keys"`
		CreateBlackWallet  bool   `json:"create_blank_wallet"`
		Passphrase         string `json:"passphrase"`
		AvoidReuse         bool   `json:"avoid_reuse"`
	}

	Move struct {
		FromAccount      string         `json:"from_account"`
		ToAccount        string         `json:"to_account"`
		Amount           btcutil.Amount `json:"amount"`
		MinConfirmations int            `json:"min_confirmations"`
		Comment          string         `json:"comment"`
	}
)

// State from a connection. If no state is present, it is initialized
func State(conState jsonrpc.State) *btcState {
	raw, exists := conState[BtcID]
	if !exists {
		ns := &btcState{}
		conState[BtcID] = ns
		return ns
	}
	ns, ok := raw.(*btcState)
	if !ok {
		// This means the invariant is violated, so panic here is ok
		panic("Invalid saved state for btc")
	}
	return ns
}

// Close implements jsonrpc.Closer
func (s *btcState) Close() {}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Load(ctx context.Context, conState jsonrpc.State, args Load) error {
	log.Debug().Msgf("BTC: connecting to btc node %s", args.Host)

	client, err := btcRpcClient.New(
		&btcRpcClient.ConnConfig{
			Host:         args.Host,
			User:         args.User,
			Pass:         args.Pass,
			HTTPPostMode: true,
			DisableTLS:   true,
		}, nil)
	if err != nil {
		return err
	}
	state := State(conState)
	state.client = client

	return nil
}

func (c *Client) ImportAddress(ctx context.Context, conState jsonrpc.State, address string) error {
	log.Debug().Msgf("BTC: importing address %s", address)

	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ImportAddress(address)
}

func (c *Client) ImportAddressRescan(ctx context.Context, conState jsonrpc.State, args ImportAddressRescan) error {
	log.Debug().Msgf("BTC: importing address rescan %s for account %s", args.Address, args.Account)

	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ImportAddressRescan(args.Address, args.Account, args.Rescan)
}

func (c *Client) ImportPrivKey(ctx context.Context, conState jsonrpc.State, wif string) error {
	log.Debug().Msg("BTC: importing private key")

	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	privKeyWIF, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return err
	}

	return state.client.ImportPrivKey(privKeyWIF)
}

func (c *Client) ImportPrivKeyLabel(ctx context.Context, conState jsonrpc.State, args ImportPrivKeyLabel) error {
	log.Debug().Msgf("BTC: importing private key with label %s", args.Label)

	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	privKeyWIF, err := btcutil.DecodeWIF(args.WIF)
	if err != nil {
		return err
	}

	return state.client.ImportPrivKeyLabel(privKeyWIF, args.Label)
}

func (c *Client) ImportPrivKeyRescan(ctx context.Context, conState jsonrpc.State, args ImportPrivKeyRescan) error {
	log.Debug().Msgf("BTC: importing private key rescan with label %s", args.Label)

	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	privKeyWIF, err := btcutil.DecodeWIF(args.WIF)
	if err != nil {
		return err
	}

	return state.client.ImportPrivKeyRescan(privKeyWIF, args.Label, args.Rescan)
}

func (c *Client) ImportPubKey(ctx context.Context, conState jsonrpc.State, pubKey string) error {
	log.Debug().Msg("BTC: importing public key")

	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ImportPubKey(pubKey)
}

func (c *Client) ImportPubKeyRescan(ctx context.Context, conState jsonrpc.State, args ImportPubKeyRescan) error {
	log.Debug().Msg("BTC: importing public key rescan")

	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ImportPubKeyRescan(args.PubKey, args.Rescan)
}

func (c *Client) RenameAccount(ctx context.Context, conState jsonrpc.State, args RenameAccount) error {
	log.Debug().Msgf("BTC: renaming account from %s to %s", args.OldAccount, args.NewAccount)

	state := State(conState)
	if state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.RenameAccount(args.OldAccount, args.NewAccount)
}

func (c *Client) SendToAddress(ctx context.Context, conState jsonrpc.State, args SendToAddress) (string, error) {
	log.Debug().Msgf("BTC: sending %d to address %s with comment %s and commentTo %s", args.Amount, args.Address, args.Comment, args.CommentTo)

	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	address, err := btcutil.DecodeAddress(args.Address, nil)
	if err != nil {
		return "", err
	}

	var blockHash *chainhash.Hash = nil
	if args.Comment != "" {
		blockHash, err = state.client.SendToAddressComment(address, args.Amount, args.Comment, args.CommentTo)
	} else {
		blockHash, err = state.client.SendToAddress(address, args.Amount)
	}
	if err != nil || blockHash == nil {
		return "", err
	}
	return blockHash.String(), err
}

func (c *Client) EstimateSmartFee(ctx context.Context, conState jsonrpc.State, args EstimateSmartFee) (*btcjson.EstimateSmartFeeResult, error) {
	log.Debug().Msgf("BTC: estimating smart fee for %s blocks with estimation mode %s", args.ConfTarget, args.Mode)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.EstimateSmartFee(args.ConfTarget, &args.Mode)
}

func hashesToStrings(hashes []*chainhash.Hash) []string {
	var blockHashes = []string{}
	for _, hash := range hashes {
		if hash == nil {
			blockHashes = append(blockHashes, "")
		} else {
			blockHashes = append(blockHashes, hash.String())
		}
	}
	return blockHashes
}

func (c *Client) GenerateBlocks(ctx context.Context, conState jsonrpc.State, numBlocks uint32) ([]string, error) {
	log.Debug().Msgf("BTC: generating %d blocks", numBlocks)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	hashes, err := state.client.Generate(numBlocks)
	if err != nil {
		return []string{}, err
	}

	return hashesToStrings(hashes), err
}

func (c *Client) GenerateBlocksToAddress(ctx context.Context, conState jsonrpc.State, args GenerateToAddress) ([]string, error) {
	log.Debug().Msgf("BTC: generating %d blocks for address %s", args.NumBlocks, args.Address)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	address, err := btcutil.DecodeAddress(args.Address, nil)
	if err != nil {
		return nil, err
	}

	hashes, err := state.client.GenerateToAddress(args.NumBlocks, address, &args.MaxTries)
	if err != nil {
		return nil, err
	}

	return hashesToStrings(hashes), err
}

func (c *Client) GetAccount(ctx context.Context, conState jsonrpc.State, address string) (string, error) {
	log.Debug().Msgf("BTC: getting account name for address %s", address)

	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	decodedAddress, err := btcutil.DecodeAddress(address, nil)
	if err != nil {
		return "", err
	}

	return state.client.GetAccount(decodedAddress)
}

func (c *Client) GetAccountAddress(ctx context.Context, conState jsonrpc.State, account string) (string, error) {
	log.Debug().Msgf("BTC: getting account address of account %s", account)

	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	address, err := state.client.GetAccountAddress(account)
	if err != nil {
		return "", err
	}

	return address.EncodeAddress(), nil
}

func (c *Client) GetAddressInfo(ctx context.Context, conState jsonrpc.State, address string) (*btcjson.GetAddressInfoResult, error) {
	log.Debug().Msgf("BTC: getting address info of address %s", address)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetAddressInfo(address)
}

func (c *Client) GetAddressesByAccount(ctx context.Context, conState jsonrpc.State, account string) ([]string, error) {
	log.Debug().Msgf("BTC: getting address of account %s", account)

	state := State(conState)
	if state.client == nil {
		return []string{}, pkg.ErrClientNotConnected{}
	}

	addresses, err := state.client.GetAddressesByAccount(account)
	if err != nil {
		return []string{}, err
	}

	addressesEncoded := []string{}
	for _, address := range addresses {
		addressesEncoded = append(addressesEncoded, address.EncodeAddress())
	}

	return addressesEncoded, nil
}

func (c *Client) GetBalance(ctx context.Context, conState jsonrpc.State, account string) (btcutil.Amount, error) {
	log.Debug().Msgf("BTC: getting balance of account %s", account)

	state := State(conState)
	if state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetBalance(account)
}

func (c *Client) GetBlockCount(ctx context.Context, conState jsonrpc.State) (int64, error) {
	log.Debug().Msg("BTC: getting block count")

	state := State(conState)
	if state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetBlockCount()
}

func (c *Client) GetBlockHash(ctx context.Context, conState jsonrpc.State, blockHeight int64) (string, error) {
	log.Debug().Msgf("BTC: getting block hash for block at height %d", blockHeight)

	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	blockHash, err := state.client.GetBlockHash(blockHeight)
	if err != nil {
		return "", err
	}

	return blockHash.String(), nil
}

func (c *Client) GetBlockStats(ctx context.Context, conState jsonrpc.State, hash string) (*btcjson.GetBlockStatsResult, error) {
	log.Debug().Msgf("BTC: getting block stats for block with hash %s", hash)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetBlockStats(hash, nil)
}

func (c *Client) GetBlockVerboseTx(ctx context.Context, conState jsonrpc.State, hash string) (*btcjson.GetBlockVerboseTxResult, error) {
	log.Debug().Msgf("BTC: getting block verbose tx for block at height %s", hash)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	blockHash, err := chainhash.NewHashFromStr(hash)
	if err != nil {
		return nil, err
	}

	return state.client.GetBlockVerboseTx(blockHash)
}

func (c *Client) GetChainTxStats(ctx context.Context, conState jsonrpc.State, args GetChainTxStatsNBlocksBlockHash) (*btcjson.GetChainTxStatsResult, error) {
	log.Debug().Msgf("BTC: getting chain transaction statistics (amount_of_blocks:%d, hash_block_end:%s)", args.AmountOfBlocks, args.BlockHashEnd)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	if args.AmountOfBlocks > 0 {
		if args.BlockHashEnd != "" {
			blockHash, err := chainhash.NewHashFromStr(args.BlockHashEnd)
			if err != nil {
				return nil, err
			}
			return state.client.GetChainTxStatsNBlocksBlockHash(args.AmountOfBlocks, *blockHash)
		}
		return state.client.GetChainTxStatsNBlocks(args.AmountOfBlocks)
	}
	return state.client.GetChainTxStats()
}

func (c *Client) GetDifficulty(ctx context.Context, conState jsonrpc.State) (float64, error) {
	log.Debug().Msg("BTC: getting difficulty")

	state := State(conState)
	if state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetDifficulty()
}

func (c *Client) GetMiningInfo(ctx context.Context, conState jsonrpc.State) (*btcjson.GetMiningInfoResult, error) {
	log.Debug().Msg("BTC: getting mining info")

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	return state.client.GetMiningInfo()
}

func (c *Client) GetNewAddress(ctx context.Context, conState jsonrpc.State, account string) (string, error) {
	log.Debug().Msgf("BTC: getting new address for account %s", account)

	state := State(conState)
	if state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	address, err := state.client.GetNewAddress(account)
	if err != nil {
		return "", err
	}

	return address.EncodeAddress(), nil
}

func (c *Client) GetNodeAddresses(ctx context.Context, conState jsonrpc.State) ([]btcjson.GetNodeAddressesResult, error) {
	log.Debug().Msg("BTC: getting node addresses")

	state := State(conState)
	if state.client == nil {
		return []btcjson.GetNodeAddressesResult{}, pkg.ErrClientNotConnected{}
	}

	return state.client.GetNodeAddresses(nil)
}

func (c *Client) GetPeerInfo(ctx context.Context, conState jsonrpc.State) ([]btcjson.GetPeerInfoResult, error) {
	log.Debug().Msg("BTC: getting peer info")

	state := State(conState)
	if state.client == nil {
		return []btcjson.GetPeerInfoResult{}, pkg.ErrClientNotConnected{}
	}

	return state.client.GetPeerInfo()
}

func (c *Client) GetRawTransaction(ctx context.Context, conState jsonrpc.State, txHash string) (*btcutil.Tx, error) {
	log.Debug().Msgf("BTC: getting raw transaction with hash %s", txHash)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	txHashDecoded, err := chainhash.NewHashFromStr(txHash)
	if err != nil {
		return nil, err
	}

	return state.client.GetRawTransaction(txHashDecoded)
}

func (c *Client) CreateWallet(ctx context.Context, conState jsonrpc.State, args CreateWallet) (*btcjson.CreateWalletResult, error) {
	log.Debug().Msgf("BTC: creating wallet with name %s (AvoidReuse:%t, CreateBlackWallet:%t, DisablePrivateKeys:%t)", args.Name, args.AvoidReuse, args.CreateBlackWallet, args.DisablePrivateKeys)

	state := State(conState)
	if state.client == nil {
		return nil, pkg.ErrClientNotConnected{}
	}

	options := []btcRpcClient.CreateWalletOpt{}
	if args.DisablePrivateKeys {
		options = append(options, btcRpcClient.WithCreateWalletDisablePrivateKeys())
	}
	if args.AvoidReuse {
		options = append(options, btcRpcClient.WithCreateWalletAvoidReuse())
	}
	if args.CreateBlackWallet {
		options = append(options, btcRpcClient.WithCreateWalletBlank())
	}
	if args.Passphrase != "" {
		options = append(options, btcRpcClient.WithCreateWalletPassphrase(args.Passphrase))
	}

	return state.client.CreateWallet(args.Name, options...)
}

func (c *Client) Move(ctx context.Context, conState jsonrpc.State, args Move) (bool, error) {
	log.Debug().Msgf("BTC: moving %d from account %s to %s", args.Amount, args.FromAccount, args.ToAccount)

	state := State(conState)
	if state.client == nil {
		return false, pkg.ErrClientNotConnected{}
	}

	if args.MinConfirmations > 0 {
		if args.Comment != "" {
			return state.client.MoveComment(args.FromAccount, args.ToAccount, args.Amount, args.MinConfirmations, args.Comment)
		}
		return state.client.MoveMinConf(args.FromAccount, args.ToAccount, args.Amount, args.MinConfirmations)
	}
	return state.client.Move(args.FromAccount, args.ToAccount, args.Amount)

}
