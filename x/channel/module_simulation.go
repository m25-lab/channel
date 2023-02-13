package channel

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/m25-lab/channel/testutil/sample"
	channelsimulation "github.com/m25-lab/channel/x/channel/simulation"
	"github.com/m25-lab/channel/x/channel/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = channelsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCommitment = "op_weight_msg_commitment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCommitment int = 100

	opWeightMsgWithdrawTimelock = "op_weight_msg_withdraw_timelock"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawTimelock int = 100

	opWeightMsgWithdrawHashlock = "op_weight_msg_withdraw_hashlock"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawHashlock int = 100

	opWeightMsgCloseChannel = "op_weight_msg_close_channel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCloseChannel int = 100

	opWeightMsgOpenChannel = "op_weight_msg_open_channel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOpenChannel int = 100

	opWeightMsgFund = "op_weight_msg_fund"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFund int = 100

	opWeightMsgAcceptfund = "op_weight_msg_acceptfund"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAcceptfund int = 100

	opWeightMsgSendercommit = "op_weight_msg_sendercommit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendercommit int = 100

	opWeightMsgSenderwithdrawtimelock = "op_weight_msg_senderwithdrawtimelock"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSenderwithdrawtimelock int = 100

	opWeightMsgSenderwithdrawhashlock = "op_weight_msg_senderwithdrawhashlock"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSenderwithdrawhashlock int = 100

	opWeightMsgReceiverwithdraw = "op_weight_msg_receiverwithdraw"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReceiverwithdraw int = 100

	opWeightMsgReceivercommit = "op_weight_msg_receivercommit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReceivercommit int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	channelGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&channelGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCommitment int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCommitment, &weightMsgCommitment, nil,
		func(_ *rand.Rand) {
			weightMsgCommitment = defaultWeightMsgCommitment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCommitment,
		channelsimulation.SimulateMsgCommitment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawTimelock int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawTimelock, &weightMsgWithdrawTimelock, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawTimelock = defaultWeightMsgWithdrawTimelock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawTimelock,
		channelsimulation.SimulateMsgWithdrawTimelock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawHashlock int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawHashlock, &weightMsgWithdrawHashlock, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawHashlock = defaultWeightMsgWithdrawHashlock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawHashlock,
		channelsimulation.SimulateMsgWithdrawHashlock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCloseChannel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCloseChannel, &weightMsgCloseChannel, nil,
		func(_ *rand.Rand) {
			weightMsgCloseChannel = defaultWeightMsgCloseChannel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCloseChannel,
		channelsimulation.SimulateMsgCloseChannel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgOpenChannel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOpenChannel, &weightMsgOpenChannel, nil,
		func(_ *rand.Rand) {
			weightMsgOpenChannel = defaultWeightMsgOpenChannel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOpenChannel,
		channelsimulation.SimulateMsgOpenChannel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFund int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFund, &weightMsgFund, nil,
		func(_ *rand.Rand) {
			weightMsgFund = defaultWeightMsgFund
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFund,
		channelsimulation.SimulateMsgFund(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAcceptfund int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAcceptfund, &weightMsgAcceptfund, nil,
		func(_ *rand.Rand) {
			weightMsgAcceptfund = defaultWeightMsgAcceptfund
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAcceptfund,
		channelsimulation.SimulateMsgAcceptfund(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendercommit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendercommit, &weightMsgSendercommit, nil,
		func(_ *rand.Rand) {
			weightMsgSendercommit = defaultWeightMsgSendercommit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendercommit,
		channelsimulation.SimulateMsgSendercommit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSenderwithdrawtimelock int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSenderwithdrawtimelock, &weightMsgSenderwithdrawtimelock, nil,
		func(_ *rand.Rand) {
			weightMsgSenderwithdrawtimelock = defaultWeightMsgSenderwithdrawtimelock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSenderwithdrawtimelock,
		channelsimulation.SimulateMsgSenderwithdrawtimelock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSenderwithdrawhashlock int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSenderwithdrawhashlock, &weightMsgSenderwithdrawhashlock, nil,
		func(_ *rand.Rand) {
			weightMsgSenderwithdrawhashlock = defaultWeightMsgSenderwithdrawhashlock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSenderwithdrawhashlock,
		channelsimulation.SimulateMsgSenderwithdrawhashlock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReceiverwithdraw int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReceiverwithdraw, &weightMsgReceiverwithdraw, nil,
		func(_ *rand.Rand) {
			weightMsgReceiverwithdraw = defaultWeightMsgReceiverwithdraw
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReceiverwithdraw,
		channelsimulation.SimulateMsgReceiverwithdraw(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReceivercommit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReceivercommit, &weightMsgReceivercommit, nil,
		func(_ *rand.Rand) {
			weightMsgReceivercommit = defaultWeightMsgReceivercommit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReceivercommit,
		channelsimulation.SimulateMsgReceivercommit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
