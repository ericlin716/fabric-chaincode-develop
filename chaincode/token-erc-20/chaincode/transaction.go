package chaincode

import (
	"fmt"
	"log"
	"strconv"
	"token-erc-20/entity/event"
	"token-erc-20/entity/transaction"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Transfer transfers tokens from client account to recipient account
// recipient account must be a valid clientID as returned by the ClientID() function
// This function triggers a Transfer event
func (s *SmartContract) Transfer(ctx contractapi.TransactionContextInterface, input transaction.TransferInput) error {

	// Get ID of submitting client identity
	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return ErrFailedGetClientId(err)
	}

	err = transaction.TransferHelper(ctx, clientID, input.Recipient, input.Amount)
	if err != nil {
		return ErrFailedTransfer(err)
	}

	// Emit the Transfer event
	err = event.SetEvent(ctx, "Transfer", event.Event{From: clientID, To: input.Recipient, Value: input.Amount})
	if err != nil {
		return err
	}

	return nil
}

// TransferFrom transfers the value amount from the "from" address to the "to" address
// This function triggers a Transfer event
func (s *SmartContract) TransferFrom(ctx contractapi.TransactionContextInterface, input transaction.TransferFromInput) error {

	// Get ID of submitting client identity
	spender, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return ErrFailedGetClientId(err)
	}

	// Create allowanceKey
	allowanceKey, err := ctx.GetStub().CreateCompositeKey(allowancePrefix, []string{input.From, spender})
	if err != nil {
		return ErrFailedCreateCompositeKey(fmt.Sprintf("%s: %v", allowancePrefix, err))
	}

	// Retrieve the allowance of the spender
	currentAllowanceBytes, err := ctx.GetStub().GetState(allowanceKey)
	if err != nil {
		return ErrFailedReadAllowance(fmt.Sprintf("%s: %v", allowanceKey, err))
	}

	var currentAllowance int
	currentAllowance, _ = strconv.Atoi(string(currentAllowanceBytes)) // Error handling not needed since Itoa() was used when setting the totalSupply, guaranteeing it was an integer.

	// Check if transferred value is less than allowance
	if currentAllowance < input.Value {
		return ErrSpenderNotHaveEnoughAllowance(nil)
	}

	// Initiate the transfer
	err = transaction.TransferHelper(ctx, input.From, input.To, input.Value)
	if err != nil {
		return ErrFailedTransfer(err)
	}

	// Decrease the allowance
	updatedAllowance := currentAllowance - input.Value
	err = ctx.GetStub().PutState(allowanceKey, []byte(strconv.Itoa(updatedAllowance)))
	if err != nil {
		return err
	}

	// Emit the Transfer event
	err = event.SetEvent(ctx, "Transfer", event.Event{From: input.From, To: input.To, Value: input.Value})
	if err != nil {
		return err
	}

	log.Printf("spender %s allowance updated from %d to %d", spender, currentAllowance, updatedAllowance)

	return nil
}
