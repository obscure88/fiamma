package bitvmstaker

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/fiamma-chain/fiamma/api/fiamma/bitvmstaker"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "AllStakerInfo",
					Use:            "all-staker-info",
					Short:          "Query all-staker-info",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "CommitteeAddress",
					Use:            "committee-address",
					Short:          "Query committee-address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				{
					RpcMethod:      "RegisteredVKList",
					Use:            "registered-vk-list",
					Short:          "Query registered-vk-list",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateStaker",
					Use:            "create-staker [staker-address]",
					Short:          "Send a create-staker tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "staker_address"}},
				},
				{
					RpcMethod:      "RemoveStaker",
					Use:            "remove-staker [remove-address]",
					Short:          "Send a remove-staker tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "staker_address"}},
				},
				{
					RpcMethod:      "UpdateCommitteeAddress",
					Use:            "update-committee-address [new-committee-address]",
					Short:          "Send a update-committee-address tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "new_committee_address"}},
				},
				{
					RpcMethod:      "RegisterVK",
					Use:            "register-vk [vk]",
					Short:          "Send a register-vk tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "vk"}},
				},
				{
					RpcMethod:      "RemoveVK",
					Use:            "remove-vk [vk]",
					Short:          "Send a remove-vk tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "vk"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
