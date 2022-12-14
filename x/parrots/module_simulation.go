package parrots

import (
	"math/rand"

	"parrots/testutil/sample"
	parrotssimulation "parrots/x/parrots/simulation"
	"parrots/x/parrots/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = parrotssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgSetProfile = "op_weight_msg_set_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetProfile int = 100

	opWeightMsgUploadBeak = "op_weight_msg_upload_beak"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUploadBeak int = 100

	opWeightMsgSendRespect = "op_weight_msg_send_respect"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendRespect int = 100

	opWeightMsgCreateComment = "op_weight_msg_create_comment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateComment int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	parrotsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&parrotsGenesis)
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

	var weightMsgSetProfile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetProfile, &weightMsgSetProfile, nil,
		func(_ *rand.Rand) {
			weightMsgSetProfile = defaultWeightMsgSetProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetProfile,
		parrotssimulation.SimulateMsgSetProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUploadBeak int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUploadBeak, &weightMsgUploadBeak, nil,
		func(_ *rand.Rand) {
			weightMsgUploadBeak = defaultWeightMsgUploadBeak
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUploadBeak,
		parrotssimulation.SimulateMsgUploadBeak(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendRespect int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendRespect, &weightMsgSendRespect, nil,
		func(_ *rand.Rand) {
			weightMsgSendRespect = defaultWeightMsgSendRespect
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendRespect,
		parrotssimulation.SimulateMsgSendRespect(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateComment int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateComment, &weightMsgCreateComment, nil,
		func(_ *rand.Rand) {
			weightMsgCreateComment = defaultWeightMsgCreateComment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateComment,
		parrotssimulation.SimulateMsgCreateComment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
