package db

import (
	"context"
	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createAnTransfer(t *testing.T, fromAccount *Account, toAccount *Account) Transfer {
	request := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}

	payload, err := testQueries.CreateTransfer(context.Background(), request)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.NotZero(t, payload.CreatedAt)

	require.Equal(t, payload.FromAccountID, request.FromAccountID)
	require.Equal(t, payload.ToAccountID, request.ToAccountID)
	require.Equal(t, payload.Amount, request.Amount)

	return payload
}

func TestCreateTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	createAnTransfer(t, &fromAccount, &toAccount)
}

func TestGetTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	transfer := createAnTransfer(t, &fromAccount, &toAccount)

	request := GetTransferParams{
		ID:    transfer.ID,
		Limit: 1,
	}

	payload, err := testQueries.GetTransfer(context.Background(), request)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, transfer, payload)
}

func TestListTransfers(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	for i := 0; i < 4; i++ {
		createAnTransfer(t, &fromAccount, &toAccount)
	}

	request := ListTransfersParams{
		FromAccountID: fromAccount.ID, // Can be switched to `toAccount` as well
		// ToAccountID: toAccount.ID,
		Limit:  4,
		Offset: 0,
	}

	payload, err := testQueries.ListTransfers(context.Background(), request)
	require.NoError(t, err)
	require.Len(t, payload, 4)

	for _, transfer := range payload {
		require.NotEmpty(t, transfer)
	}
}
