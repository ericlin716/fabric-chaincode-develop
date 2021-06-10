package chaincode

import (
	"fmt"
	"log"
	"strconv"
	"token-erc-20/entity/event"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Mint creates new tokens and adds them to minter's account balance
// This function triggers a Transfer event
func (s *SmartContract) Mint(ctx contractapi.TransactionContextInterface, amount int) error {

	// Check minter authorization - this sample assumes Org1 is the central banker with privilege to mint new tokens
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return ErrFailedGetMSPID(err)
	}
	if clientMSPID != "Org1MSP" {
		return ErrClientIsNotAuthorizedToMint(nil)
	}

	// Get ID of submitting client identity
	minter, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return ErrFailedGetClientId(err)
	}

	if amount <= 0 {
		return ErrMintAmountMustBePositive(nil)
	}

	currentBalanceBytes, err := ctx.GetStub().GetState(minter)
	if err != nil {
		return ErrFailedReadMinterAccount(fmt.Sprintf("%s: %v", minter, err))
	}

	var currentBalance int

	// If minter current balance doesn't yet exist, we'll create it with a current balance of 0
	if currentBalanceBytes == nil {
		currentBalance = 0
	} else {
		currentBalance, _ = strconv.Atoi(string(currentBalanceBytes)) // Error handling not needed since Itoa() was used when setting the account balance, guaranteeing it was an integer.
	}

	updatedBalance := currentBalance + amount

	err = ctx.GetStub().PutState(minter, []byte(strconv.Itoa(updatedBalance)))
	if err != nil {
		return err
	}

	// Update the totalSupply
	totalSupplyBytes, err := ctx.GetStub().GetState(totalSupplyKey)
	if err != nil {
		return ErrFailedRetrieveTotalSupply(err)
	}

	var totalSupply int

	// If no tokens have been minted, initialize the totalSupply
	if totalSupplyBytes == nil {
		totalSupply = 0
	} else {
		totalSupply, _ = strconv.Atoi(string(totalSupplyBytes)) // Error handling not needed since Itoa() was used when setting the totalSupply, guaranteeing it was an integer.
	}

	// Add the mint amount to the total supply and update the state
	totalSupply += amount
	err = ctx.GetStub().PutState(totalSupplyKey, []byte(strconv.Itoa(totalSupply)))
	if err != nil {
		return err
	}

	// Emit the Transfer event
	err = event.SetEvent(ctx, "Transfer", event.Event{From: "0x0", To: minter, Value: amount})
	if err != nil {
		return err
	}

	log.Printf("minter account %s balance updated from %d to %d", minter, currentBalance, updatedBalance)

	return nil
}

// Burn redeems tokens the minter's account balance
// This function triggers a Transfer event
func (s *SmartContract) Burn(ctx contractapi.TransactionContextInterface, amount int) error {

	// Check minter authorization - this sample assumes Org1 is the central banker with privilege to burn new tokens
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return ErrFailedGetMSPID(err)
	}
	if clientMSPID != "Org1MSP" {
		return ErrClientIsNotAuthorizedToMint(nil)
	}

	// Get ID of submitting client identity
	minter, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return ErrFailedGetClientId(err)
	}

	if amount <= 0 {
		return ErrBurnAmountMustBePositive(nil)
	}

	currentBalanceBytes, err := ctx.GetStub().GetState(minter)
	if err != nil {
		return ErrFailedReadMinterAccount(fmt.Sprintf("%s: %v", minter, err))
	}

	var currentBalance int

	// Check if minter current balance exists
	if currentBalanceBytes == nil {
		return ErrBalanceNotExist(nil)
	}

	currentBalance, _ = strconv.Atoi(string(currentBalanceBytes)) // Error handling not needed since Itoa() was used when setting the account balance, guaranteeing it was an integer.

	updatedBalance := currentBalance - amount

	err = ctx.GetStub().PutState(minter, []byte(strconv.Itoa(updatedBalance)))
	if err != nil {
		return err
	}

	// Update the totalSupply
	totalSupplyBytes, err := ctx.GetStub().GetState(totalSupplyKey)
	if err != nil {
		return ErrFailedRetrieveTotalSupply(err)
	}

	// If no tokens have been minted, throw error
	if totalSupplyBytes == nil {
		return ErrTotalSupplyNotExist(nil)
	}

	totalSupply, _ := strconv.Atoi(string(totalSupplyBytes)) // Error handling not needed since Itoa() was used when setting the totalSupply, guaranteeing it was an integer.

	// Subtract the burn amount to the total supply and update the state
	totalSupply -= amount
	err = ctx.GetStub().PutState(totalSupplyKey, []byte(strconv.Itoa(totalSupply)))
	if err != nil {
		return err
	}

	// Emit the Transfer event
	err = event.SetEvent(ctx, "Transfer", event.Event{From: minter, To: "0x0", Value: amount})
	if err != nil {
		return err
	}

	log.Printf("minter account %s balance updated from %d to %d", minter, currentBalance, updatedBalance)

	return nil
}
