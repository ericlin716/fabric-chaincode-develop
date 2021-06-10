package chaincode

import "fmt"

type CreateError = func(interface{}) error

var (
	ErrFailedGetClientId             CreateError = func(err interface{}) error { return fmt.Errorf("failed to get client id: %v", err) }
	ErrFailedReadWorldState          CreateError = func(err interface{}) error { return fmt.Errorf("failed to read from world state: %v", err) }
	ErrAccountNotExist               CreateError = func(err interface{}) error { return fmt.Errorf("the account %s does not exist", err) }
	ErrFailedRetrieveTotalSupply     CreateError = func(err interface{}) error { return fmt.Errorf("failed to retrieve total token supply: %v", err) }
	ErrFailedGetMSPID                CreateError = func(err interface{}) error { return fmt.Errorf("failed to get MSPID: %v", err) }
	ErrClientIsNotAuthorizedToMint   CreateError = func(err interface{}) error { return fmt.Errorf("client is not authorized to mint new tokens") }
	ErrMintAmountMustBePositive      CreateError = func(err interface{}) error { return fmt.Errorf("mint amount must be a positive integer") }
	ErrBurnAmountMustBePositive      CreateError = func(err interface{}) error { return fmt.Errorf("burn amount must be a positive integer") }
	ErrBalanceNotExist               CreateError = func(err interface{}) error { return fmt.Errorf("the balance does not exist") }
	ErrTotalSupplyNotExist           CreateError = func(err interface{}) error { return fmt.Errorf("totalSupply does not exist") }
	ErrFailedTransfer                CreateError = func(err interface{}) error { return fmt.Errorf("failed to transfer: %v", err) }
	ErrSpenderNotHaveEnoughAllowance CreateError = func(err interface{}) error { return fmt.Errorf("spender does not have enough allowance for transfer") }
	ErrFailedCreateCompositeKey      CreateError = func(err interface{}) error {
		return fmt.Errorf("failed to create the composite key for prefix %s", err)
	}
	ErrFailedUpdateStateForKey CreateError = func(err interface{}) error {
		return fmt.Errorf("failed to update state of smart contract for key %s", err)
	}
	ErrFailedReadAllowance CreateError = func(err interface{}) error {
		return fmt.Errorf("failed to read allowance from world state: %s", err)
	}
	ErrFailedReadMinterAccount CreateError = func(err interface{}) error {
		return fmt.Errorf("failed to read minter account from world state: %s", err)
	}
)
