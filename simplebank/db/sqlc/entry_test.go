package db

import (
	"context"
	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createAnEntry(t *testing.T, account *Account) Entry {

	request := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	payload, err := testQueries.CreateEntry(context.Background(), request)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, payload.AccountID, account.ID)
	require.Equal(t, payload.Amount, request.Amount)
	require.NotZero(t, payload.ID)
	require.NotZero(t, payload.CreatedAt)

	return payload
}

func TestCreateEntry(t *testing.T) {
	acc := createRandomAccount(t)
	createAnEntry(t, &acc)
}

func TestGetEntry(t *testing.T) {
	acc := createRandomAccount(t)
	entry := createAnEntry(t, &acc)

	payload, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	// entry.Amount = 90000 // Error check
	// require.Equal(t, payload.ID, entry.ID)
	// require.Equal(t, payload.AccountID, entry.AccountID)
	// require.Equal(t, payload.Amount, entry.Amount)

	require.Equal(t, payload, entry) // Ketimbang validate 1 - 1, tak validate struct to struct aja woawokwaokawokwaok
}

func TestListEntries(t *testing.T) {
	acc := createRandomAccount(t)

	for i := 0; i < 4; i++ {
		createAnEntry(t, &acc)
	}

	request := ListEntriesParams{
		AccountID: acc.ID,
		Limit:     4,
		Offset:    0,
	}

	payload, err := testQueries.ListEntries(context.Background(), request)
	require.NoError(t, err)
	require.Len(t, payload, 4)

	for _, entry := range payload {
		require.NotEmpty(t, entry)
	}
}
