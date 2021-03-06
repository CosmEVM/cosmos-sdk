package client

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gogo/protobuf/jsonpb"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// Codec defines the x/auth account codec to be used for use with the
// AccountRetriever. The application must be sure to set this to their respective
// codec that implements the Codec interface and must be the same codec that
// passed to the x/auth module.
//
// TODO:/XXX: Using a package-level global isn't ideal and we should consider
// refactoring the module manager to allow passing in the correct module codec.
var Codec codec.Marshaler

// GasEstimateResponse defines a response definition for tx gas estimation.
type GasEstimateResponse struct {
	GasEstimate uint64 `json:"gas_estimate" yaml:"gas_estimate"`
}

func (gr GasEstimateResponse) String() string {
	return fmt.Sprintf("gas estimate: %d", gr.GasEstimate)
}

// EnrichWithGas calculates the gas estimate that would be consumed by the
// transaction and set the transaction's respective value accordingly.
func EnrichWithGas(txBldr authtypes.TxBuilder, clientCtx client.Context, msgs []sdk.Msg) (authtypes.TxBuilder, error) {
	_, adjusted, err := simulateMsgs(txBldr, clientCtx, msgs)
	if err != nil {
		return txBldr, err
	}

	return txBldr.WithGas(adjusted), nil
}

// CalculateGas simulates the execution of a transaction and returns
// the simulation response obtained by the query and the adjusted gas amount.
func CalculateGas(
	queryFunc func(string, []byte) ([]byte, int64, error), cdc *codec.Codec,
	txBytes []byte, adjustment float64,
) (sdk.SimulationResponse, uint64, error) {

	// run a simulation (via /app/simulate query) to
	// estimate gas and update TxBuilder accordingly
	rawRes, _, err := queryFunc("/app/simulate", txBytes)
	if err != nil {
		return sdk.SimulationResponse{}, 0, err
	}

	simRes, err := parseQueryResponse(rawRes)
	if err != nil {
		return sdk.SimulationResponse{}, 0, err
	}

	adjusted := adjustGasEstimate(simRes.GasUsed, adjustment)
	return simRes, adjusted, nil
}

// PrintUnsignedStdTx builds an unsigned StdTx and prints it to os.Stdout.
func PrintUnsignedStdTx(txBldr tx.Factory, clientCtx client.Context, msgs []sdk.Msg) error {
	err := tx.GenerateOrBroadcastTxWithFactory(clientCtx, txBldr, msgs...)
	return err
}

// SignTx appends a signature to a transaction. If appendSig
// is false, it replaces the signatures already attached with the new signature.
// Don't perform online validation or lookups if offline is true.
func SignTx(txFactory tx.Factory, clientCtx client.Context, name string, stdTx client.TxBuilder, offline bool) error {
	info, err := txFactory.Keybase().Key(name)
	if err != nil {
		return err
	}
	addr := sdk.AccAddress(info.GetPubKey().Address())
	if !isTxSigner(addr, stdTx.GetTx().GetSigners()) {
		return fmt.Errorf("%s: %s", sdkerrors.ErrorInvalidSigner, name)
	}
	if !offline {
		txFactory, err = populateAccountFromState(txFactory, clientCtx, addr)
		if err != nil {
			return err
		}
	}

	return tx.Sign(txFactory, name, stdTx)
}

// SignTxWithSignerAddress attaches a signature to a transaction.
// Don't perform online validation or lookups if offline is true, else
// populate account and sequence numbers from a foreign account.
func SignTxWithSignerAddress(txFactory tx.Factory, clientCtx client.Context, addr sdk.AccAddress,
	name string, txBuilder client.TxBuilder, offline bool) (err error) {

	// check whether the address is a signer
	if !isTxSigner(addr, txBuilder.GetTx().GetSigners()) {
		return fmt.Errorf("%s: %s", sdkerrors.ErrorInvalidSigner, name)
	}

	if !offline {
		txFactory, err = populateAccountFromState(txFactory, clientCtx, addr)
		if err != nil {
			return err
		}
	}

	return tx.Sign(txFactory, name, txBuilder)
}

// Read and decode a StdTx from the given filename.  Can pass "-" to read from stdin.
func ReadTxFromFile(ctx client.Context, filename string) (tx sdk.Tx, err error) {
	var bytes []byte

	if filename == "-" {
		bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, err = ioutil.ReadFile(filename)
	}

	if err != nil {
		return
	}

	return ctx.TxConfig.TxJSONDecoder()(bytes)
}

// NewBatchScanner returns a new BatchScanner to read newline-delimited StdTx transactions from r.
func NewBatchScanner(cfg client.TxConfig, r io.Reader) *BatchScanner {
	return &BatchScanner{Scanner: bufio.NewScanner(r), cfg: cfg}
}

// BatchScanner provides a convenient interface for reading batch data such as a file
// of newline-delimited JSON encoded StdTx.
type BatchScanner struct {
	*bufio.Scanner
	theTx        sdk.Tx
	cfg          client.TxConfig
	unmarshalErr error
}

// Tx returns the most recent Tx unmarshalled by a call to Scan.
func (bs BatchScanner) Tx() sdk.Tx { return bs.theTx }

// UnmarshalErr returns the first unmarshalling error that was encountered by the scanner.
func (bs BatchScanner) UnmarshalErr() error { return bs.unmarshalErr }

// Scan advances the Scanner to the next line.
func (bs *BatchScanner) Scan() bool {
	if !bs.Scanner.Scan() {
		return false
	}

	tx, err := bs.cfg.TxJSONDecoder()(bs.Bytes())
	bs.theTx = tx
	if err != nil && bs.unmarshalErr == nil {
		bs.unmarshalErr = err
		return false
	}

	return true
}

func populateAccountFromState(
	txBldr tx.Factory, clientCtx client.Context, addr sdk.AccAddress,
) (tx.Factory, error) {

	num, seq, err := clientCtx.AccountRetriever.GetAccountNumberSequence(clientCtx, addr)
	if err != nil {
		return txBldr, err
	}

	return txBldr.WithAccountNumber(num).WithSequence(seq), nil
}

// GetTxEncoder return tx encoder from global sdk configuration if ones is defined.
// Otherwise returns encoder with default logic.
func GetTxEncoder(cdc *codec.Codec) (encoder sdk.TxEncoder) {
	encoder = sdk.GetConfig().GetTxEncoder()
	if encoder == nil {
		encoder = authtypes.DefaultTxEncoder(cdc)
	}

	return encoder
}

// simulateMsgs simulates the transaction and returns the simulation response and
// the adjusted gas value.
func simulateMsgs(txBldr authtypes.TxBuilder, clientCtx client.Context, msgs []sdk.Msg) (sdk.SimulationResponse, uint64, error) {
	txBytes, err := txBldr.BuildTxForSim(msgs)
	if err != nil {
		return sdk.SimulationResponse{}, 0, err
	}

	return CalculateGas(clientCtx.QueryWithData, clientCtx.Codec, txBytes, txBldr.GasAdjustment())
}

func adjustGasEstimate(estimate uint64, adjustment float64) uint64 {
	return uint64(adjustment * float64(estimate))
}

func parseQueryResponse(bz []byte) (sdk.SimulationResponse, error) {
	var simRes sdk.SimulationResponse
	if err := jsonpb.Unmarshal(strings.NewReader(string(bz)), &simRes); err != nil {
		return sdk.SimulationResponse{}, err
	}

	return simRes, nil
}

func isTxSigner(user sdk.AccAddress, signers []sdk.AccAddress) bool {
	for _, s := range signers {
		if bytes.Equal(user.Bytes(), s.Bytes()) {
			return true
		}
	}

	return false
}
