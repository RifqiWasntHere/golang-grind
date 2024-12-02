package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransaction(t *testing.T) {
	store := Newstore(testDb)

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	fmt.Println(">> Before:", fromAccount.Balance, toAccount.Balance)

	// run n concurrent transactions
	n := 5
	amount := int64(10) // amount used in transaction

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccount.ID,
				ToAccountID:   toAccount.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// Check result from channel

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// Check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
		require.Equal(t, toAccount.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), GetTransferParams{ID: transfer.ID, Limit: 1})
		require.NoError(t, err)

		// Check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, fromAccount.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, toEntry.AccountID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// Check accounts
		checkFromAccount := result.FromAccount
		require.NotEmpty(t, checkFromAccount)
		require.Equal(t, fromAccount.ID, checkFromAccount.ID)

		checkToAccount := result.ToAccount
		require.NotEmpty(t, checkToAccount)
		require.Equal(t, toAccount.ID, checkToAccount.ID)

		// checks account balance
		fmt.Println(">> tx:", checkFromAccount.Balance, checkToAccount.Balance)
		diff1 := fromAccount.Balance - checkFromAccount.Balance // check if the diff between transaction is == amount
		diff2 := checkToAccount.Balance - toAccount.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	//  check the final updated balance
	updatedFromAccount, err := testQueries.GetAccount(context.Background(), fromAccount.ID)
	require.NoError(t, err)

	updatedToAccount, err := testQueries.GetAccount(context.Background(), toAccount.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedFromAccount.Balance, updatedToAccount.Balance)
	require.Equal(t, fromAccount.Balance-int64(n)*amount, updatedFromAccount.Balance)
	require.Equal(t, toAccount.Balance+int64(n)*amount, updatedToAccount.Balance)
}
