package cmds

import (
	"github.com/ProtoconNet/mitum-currency-extension/v2/cmds"
)

type OperationCommand struct {
	CreateAccount         cmds.CreateAccountCommand         `cmd:"" name:"create-account" help:"create new account"`
	KeyUpdater            cmds.KeyUpdaterCommand            `cmd:"" name:"key-updater" help:"update account keys"`
	Transfer              cmds.TransferCommand              `cmd:"" name:"transfer" help:"transfer amounts to receiver"`
	CreateContractAccount cmds.CreateContractAccountCommand `cmd:"" name:"create-contract-account" help:"create new contract account"`
	Withdraw              cmds.WithdrawCommand              `cmd:"" name:"withdraw" help:"withdraw amounts from target contract account"`
	CurrencyRegister      cmds.CurrencyRegisterCommand      `cmd:"" name:"currency-register" help:"register new currency"`
	CurrencyPolicyUpdater cmds.CurrencyPolicyUpdaterCommand `cmd:"" name:"currency-policy-updater" help:"update currency policy"`
	SuffrageInflation     cmds.SuffrageInflationCommand     `cmd:"" name:"suffrage-inflation" help:"suffrage inflation operation"`
	Append                AppendCommand                     `cmd:"" name:"append" help:"append new timestamp to service"`
	SuffrageCandidate     cmds.SuffrageCandidateCommand     `cmd:"" name:"suffrage-candidate" help:"suffrage candidate operation"`
	SuffrageJoin          cmds.SuffrageJoinCommand          `cmd:"" name:"suffrage-join" help:"suffrage join operation"`
	SuffrageDisjoin       cmds.SuffrageDisjoinCommand       `cmd:"" name:"suffrage-disjoin" help:"suffrage disjoin operation"` // revive:disable-line:line-length-limit
}

func NewOperationCommand() OperationCommand {
	return OperationCommand{
		CreateAccount:         cmds.NewCreateAccountCommand(),
		KeyUpdater:            cmds.NewKeyUpdaterCommand(),
		Transfer:              cmds.NewTransferCommand(),
		CreateContractAccount: cmds.NewCreateContractAccountCommand(),
		Withdraw:              cmds.NewWithdrawCommand(),
		CurrencyRegister:      cmds.NewCurrencyRegisterCommand(),
		CurrencyPolicyUpdater: cmds.NewCurrencyPolicyUpdaterCommand(),
		SuffrageInflation:     cmds.NewSuffrageInflationCommand(),
		Append:                NewAppendCommand(),
		SuffrageCandidate:     cmds.NewSuffrageCandidateCommand(),
		SuffrageJoin:          cmds.NewSuffrageJoinCommand(),
		SuffrageDisjoin:       cmds.NewSuffrageDisjoinCommand(),
	}
}
