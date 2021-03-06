// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package questing

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// EnrollQuestor is the `enrollQuestor` instruction.
type EnrollQuestor struct {

	// [0] = [WRITE, SIGNER] initializer
	//
	// [1] = [WRITE] questor
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewEnrollQuestorInstructionBuilder creates a new `EnrollQuestor` instruction builder.
func NewEnrollQuestorInstructionBuilder() *EnrollQuestor {
	nd := &EnrollQuestor{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetInitializerAccount sets the "initializer" account.
func (inst *EnrollQuestor) SetInitializerAccount(initializer ag_solanago.PublicKey) *EnrollQuestor {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(initializer).WRITE().SIGNER()
	return inst
}

// GetInitializerAccount gets the "initializer" account.
func (inst *EnrollQuestor) GetInitializerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetQuestorAccount sets the "questor" account.
func (inst *EnrollQuestor) SetQuestorAccount(questor ag_solanago.PublicKey) *EnrollQuestor {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(questor).WRITE()
	return inst
}

// GetQuestorAccount gets the "questor" account.
func (inst *EnrollQuestor) GetQuestorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *EnrollQuestor) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *EnrollQuestor {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *EnrollQuestor) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst EnrollQuestor) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_EnrollQuestor,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst EnrollQuestor) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *EnrollQuestor) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Initializer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Questor is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *EnrollQuestor) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("EnrollQuestor")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  initializer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      questor", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj EnrollQuestor) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *EnrollQuestor) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewEnrollQuestorInstruction declares a new EnrollQuestor instruction with the provided parameters and accounts.
func NewEnrollQuestorInstruction(
	// Accounts:
	initializer ag_solanago.PublicKey,
	questor ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *EnrollQuestor {
	return NewEnrollQuestorInstructionBuilder().
		SetInitializerAccount(initializer).
		SetQuestorAccount(questor).
		SetSystemProgramAccount(systemProgram)
}
