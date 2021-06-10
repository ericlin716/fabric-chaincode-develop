package chaincode

import (
	"fmt"
	"log"
	"strconv"
	"token-erc-20/entity/event"
	"token-erc-20/entity/permission"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Approve allows the spender to withdraw from the calling client's token account
// The spender can withdraw multiple times if necessary, up to the value amount
// This function triggers an Approval event
func (s *SmartContract) Approve(ctx contractapi.TransactionContextInterface, input permission.ApproveInput) error {

	// Get ID of submitting client identity
	owner, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return ErrFailedGetClientId(err)
	}

	// Create allowanceKey
	allowanceKey, err := ctx.GetStub().CreateCompositeKey(allowancePrefix, []string{owner, input.Spender})
	if err != nil {
		return ErrFailedCreateCompositeKey(fmt.Sprintf("%s: %v", allowancePrefix, err))
	}

	// Update the state of the smart contract by adding the allowanceKey and value
	err = ctx.GetStub().PutState(allowanceKey, []byte(strconv.Itoa(input.Value)))
	if err != nil {
		return ErrFailedUpdateStateForKey(fmt.Sprintf("%s: %v", allowanceKey, err))
	}

	// Emit the Approval event
	err = event.SetEvent(ctx, "Approval", event.Event{From: owner, To: input.Spender, Value: input.Value})
	if err != nil {
		return err
	}

	log.Printf("client %s approved a withdrawal allowance of %d for spender %s", owner, input.Value, input.Spender)

	return nil
}

// Allowance returns the amount still available for the spender to withdraw from the owner
func (s *SmartContract) Allowance(ctx contractapi.TransactionContextInterface, input permission.AllowanceInput) (int, error) {

	// Create allowanceKey
	allowanceKey, err := ctx.GetStub().CreateCompositeKey(allowancePrefix, []string{input.Owner, input.Spender})
	if err != nil {
		return 0, ErrFailedCreateCompositeKey(fmt.Sprintf("%s: %v", allowancePrefix, err))
	}

	// Read the allowance amount from the world state
	allowanceBytes, err := ctx.GetStub().GetState(allowanceKey)
	if err != nil {
		return 0, ErrFailedReadAllowance(fmt.Sprintf("%s: %v", allowanceKey, err))
	}

	var allowance int

	// If no current allowance, set allowance to 0
	if allowanceBytes == nil {
		allowance = 0
	} else {
		allowance, _ = strconv.Atoi(string(allowanceBytes)) // Error handling not needed since Itoa() was used when setting the totalSupply, guaranteeing it was an integer.
	}

	log.Printf("The allowance left for spender %s to withdraw from owner %s: %d", input.Spender, input.Owner, allowance)

	return allowance, nil
}
