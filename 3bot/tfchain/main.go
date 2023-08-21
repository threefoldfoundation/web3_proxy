package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
	"github.com/urfave/cli"
)

func main() {
	log.SetOutput(io.Discard)
	zerolog.SetGlobalLevel(zerolog.Disabled)

	app := &cli.App{
		Name:  "tfchain",
		Usage: "Example: tfchain [COMMAND]",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "substrate",
				Value: "wss://tfchain.grid.tf/ws",
				Usage: "substrate URL",
			},
		},
		Commands: []cli.Command{
			{
				Name: "create_node_contract",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.UintFlag{
						Name:     "node_id",
						Required: true,
						Usage:    "node id to create the contract on",
					},
					cli.StringFlag{
						Name:  "body",
						Value: "",
						Usage: "contract body",
					},
					cli.StringFlag{
						Name:     "hash",
						Required: true,
						Usage:    "deployment hash",
					},
					cli.UintFlag{
						Name:  "public_ips",
						Value: 0,
						Usage: "number of reserved public ips for this deployment",
					},
					cli.Uint64Flag{
						Name:  "solution_provider",
						Value: 0,
						Usage: "twin id for the solution provider",
					},
				},
				Action: substrateDecorator(createNodeContract),
			},
			{
				Name:  "update_node_contract",
				Usage: "update node contract",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.UintFlag{
						Name:     "contract_id",
						Required: true,
						Usage:    "id of contract to update",
					},
					cli.StringFlag{
						Name:  "body",
						Value: "",
						Usage: "contract body",
					},
					cli.StringFlag{
						Name:     "hash",
						Required: true,
						Usage:    "deployment hash",
					},
				},
				Action: substrateDecorator(updateNodeContract),
			},
			{
				Name:  "cancel_contract",
				Usage: "cancel any type of contract",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.UintFlag{
						Name:     "contract_id",
						Required: true,
						Usage:    "id of contract to update",
					},
				},
				Action: substrateDecorator(cancelContract),
			},
			{
				Name:  "create_name_contract",
				Usage: "create name contract",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.UintFlag{
						Name:     "name",
						Required: true,
						Usage:    "contract name",
					},
				},
				Action: substrateDecorator(createNameContract),
			},
			{
				Name:  "create_rent_contract",
				Usage: "create rent contract",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.UintFlag{
						Name:     "node_id",
						Required: true,
						Usage:    "id of node to rent",
					},
					cli.UintFlag{
						Name:  "solution_provider",
						Value: 0,
						Usage: "solution provider twin id",
					},
				},
				Action: substrateDecorator(createRentContract),
			},
			{
				Name:  "sign",
				Usage: "sign a deployment",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:     "hash",
						Required: true,
						Usage:    "deployment hash",
					},
				},
				Action: func(c *cli.Context) error {
					mnemonics := c.String("mnemonics")
					identity, err := substrate.NewIdentityFromSr25519Phrase(mnemonics)
					if err != nil {
						return errors.Wrap(err, "failed to create identity from provided mnemonics")
					}

					hash := c.String("hash")
					signatureBytes, err := identity.Sign([]byte(hash))
					if err != nil {
						return errors.Wrap(err, "failed to sign deployment hash")
					}

					sig := hex.EncodeToString(signatureBytes)
					fmt.Printf("%s", sig)

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func substrateDecorator(action func(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error)) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		substrateURL := ctx.Parent().String("substrate")

		manager := substrate.NewManager(substrateURL)
		sub, err := manager.Substrate()
		if err != nil {
			return errors.Wrap(err, "failed to create substrate connection")
		}
		defer sub.Close()

		mnemonics := ctx.String("mnemonics")
		identity, err := substrate.NewIdentityFromSr25519Phrase(mnemonics)
		if err != nil {
			return errors.Wrap(err, "failed to create identity from provided mnemonics")
		}

		ret, err := action(ctx, sub, identity)
		if err != nil {
			return err
		}

		fmt.Printf("%v", ret)
		return nil
	}
}

func createNameContract(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error) {
	name := ctx.String("name")

	contractID, err := sub.CreateNameContract(identity, name)
	if err != nil {
		return nil, err
	}

	return contractID, nil
}

func createRentContract(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error) {
	nodeID := ctx.Uint("node_id")
	solutionProvider := ctx.Uint64("solution_provider")
	spp := &solutionProvider
	if solutionProvider == 0 {
		spp = nil
	}

	contractID, err := sub.CreateRentContract(identity, uint32(nodeID), spp)
	if err != nil {
		return nil, err
	}

	return contractID, nil
}

func cancelContract(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error) {
	contractID := ctx.Uint64("contract_id")

	if err := sub.CancelContract(identity, contractID); err != nil {
		return nil, err
	}

	return "", nil
}

func createNodeContract(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error) {
	nodeID := ctx.Uint("node_id")
	body := ctx.String("body")
	hash := ctx.String("hash")
	publicIPs := ctx.Uint("public_ips")
	solutionProvider := ctx.Uint64("solution_provider")
	spp := &solutionProvider
	if solutionProvider == 0 {
		spp = nil
	}

	contractID, err := sub.CreateNodeContract(identity, uint32(nodeID), body, hash, uint32(publicIPs), spp)
	if err != nil {
		return nil, err
	}

	return contractID, nil
}

func updateNodeContract(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error) {
	contractID := ctx.Uint64("contract_id")
	body := ctx.String("body")
	hash := ctx.String("hash")

	_, err := sub.UpdateNodeContract(identity, contractID, body, hash)
	if err != nil {
		return nil, err
	}

	return "", nil
}
