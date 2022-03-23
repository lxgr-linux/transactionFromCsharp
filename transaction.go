package main

import (
  "context"
  "C"
  "log"

  sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

func getLogger(name string) (*log.Logger) {
  logger := log.Default()
  logger.SetPrefix("\033[1m[" + name + "]\033[0m ")
  return logger
}

func getAddr(logger *log.Logger, cosmos cosmosclient.Client, user string) sdktypes.AccAddress {
  address, err := cosmos.Address(user)
	if err != nil {
		logger.Fatal("Error:", err)
	}
  return address
}

func broadcastMsg(logger *log.Logger, cosmos cosmosclient.Client, creator string, msg sdktypes.Msg) {
  go logger.Println("Message:", msg)

  txResp, err := cosmos.BroadcastTx(creator, msg)
  if err != nil {
    logger.Fatal("Error:", err)
  }
  go logger.Println("Response:", txResp)
}

func getClient() (cosmosclient.Client, error) {
  config := sdktypes.GetConfig()
	config.SetBech32PrefixForAccount("cc", "ccpub")

  return cosmosclient.New(context.Background(), cosmosclient.WithAddressPrefix("cc"))
}

//export makeConfirmMatchRequest
func makeConfirmMatchRequest(creator *C.char, matchId int, rawOutcome *C.char) {
  logger := getLogger("make_add_artwork_request")
  cosmos, err := getClient()
	if err != nil {
		logger.Fatal("Error:", err)
	}
  address := getAddr(logger, cosmos, C.GoString(creator))

  outcome := types.Outcome(types.Outcome_value[C.GoString(rawOutcome)])

  msg := types.NewMsgConfirmMatch(
    address.String(),
		uint64(matchId),
		outcome,
	)

  broadcastMsg(logger, cosmos, C.GoString(creator), msg)
}

func main() {
}
