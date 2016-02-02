package app

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"net"
	"sync"

	"github.com/codegangsta/cli"
	client "github.com/tendermint/go-rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tmsp/types"
)

const ()

type ETCDApplication struct {

	// TODO: hold etcd state

	// client to the tendermint core rpc
	client *client.ClientURI
}

func NewETCDApplication(ctx *cli.Context) *ETCDApplication {

	// TODO...

	etcdApp := &ETCDApplication{
		client: client.NewClientURI(fmt.Sprintf("http://%s", ctx.String(TendermintCoreHostFlag.Name))),
	}

	// TODO: Start RPC server, etc.

	return etcdApp
}

/* For broadcasting txs/requests received on etcd's rpc to the tendermint core
func (app *ETCDApplication) BroadcastTx(tx someType) error {
	var result ctypes.TMResult
	buf := new(bytes.Buffer)

	TODO: decode tx ...


	// Send tx to tendermint core
	params := map[string]interface{}{
		"tx": hex.EncodeToString(buf.Bytes()),
	}
	_, err := app.client.Call("broadcast_tx", params, &result)
	return err
}
*/

//--------------------------------------------------------
// Implements TMSP App

func (app *ETCDApplication) Info() string {
	// TODO: cache some info about the state
	return ""
}

func (app *ETCDApplication) Query(query []byte) (result []byte, log string) {
	return nil, ""
}

func (app *ETCDApplication) SetOption(key string, value string) string {
	// TODO: gas limits, other params
	return ""
}

func (app *ETCDApplication) AppendTx(txBytes []byte) (retCode types.RetCode, result []byte, log string) {
	// TODO: decode and run tx

	return types.RetCodeOK, result, log
}

func (app *ETCDApplication) CheckTx(tx []byte) (retCode types.RetCode, result []byte, log string) {
	// TODO: check nonce, balance, etc
	return retCode, result, log
}

// Commit
func (app *ETCDApplication) GetHash() (hashBytes []byte, log string) {

	//TODO: this is the "commit" function - the current state should be committed
	// etcd has no merkle tree so either we put one in or we just return some common string here

	hashBytes = []byte{"etcd-has-no-merkle-tree"}

	return hashBytes, log
}
