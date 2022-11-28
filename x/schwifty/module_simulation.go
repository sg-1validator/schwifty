package schwifty

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/toschdev/schwifty/testutil/sample"
	schwiftysimulation "github.com/toschdev/schwifty/x/schwifty/simulation"
	"github.com/toschdev/schwifty/x/schwifty/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = schwiftysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateCollection = "op_weight_msg_collection"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCollection int = 100

	opWeightMsgUpdateCollection = "op_weight_msg_collection"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCollection int = 100

	opWeightMsgDeleteCollection = "op_weight_msg_collection"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCollection int = 100

	opWeightMsgCreateNft = "op_weight_msg_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateNft int = 100

	opWeightMsgUpdateNft = "op_weight_msg_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateNft int = 100

	opWeightMsgDeleteNft = "op_weight_msg_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteNft int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	schwiftyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		CollectionList: []types.Collection{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		CollectionCount: 2,
		NftList: []types.Nft{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		NftCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&schwiftyGenesis)
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

	var weightMsgCreateCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCollection, &weightMsgCreateCollection, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCollection = defaultWeightMsgCreateCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCollection,
		schwiftysimulation.SimulateMsgCreateCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateCollection, &weightMsgUpdateCollection, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCollection = defaultWeightMsgUpdateCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCollection,
		schwiftysimulation.SimulateMsgUpdateCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCollection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteCollection, &weightMsgDeleteCollection, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCollection = defaultWeightMsgDeleteCollection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCollection,
		schwiftysimulation.SimulateMsgDeleteCollection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateNft int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateNft, &weightMsgCreateNft, nil,
		func(_ *rand.Rand) {
			weightMsgCreateNft = defaultWeightMsgCreateNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateNft,
		schwiftysimulation.SimulateMsgCreateNft(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateNft int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateNft, &weightMsgUpdateNft, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateNft = defaultWeightMsgUpdateNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateNft,
		schwiftysimulation.SimulateMsgUpdateNft(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteNft int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteNft, &weightMsgDeleteNft, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteNft = defaultWeightMsgDeleteNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteNft,
		schwiftysimulation.SimulateMsgDeleteNft(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
