package db

import (
	"context"
	"database/sql"
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	// fmt.Println(account)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	// require.Equal(t, arg.Owner, account.Owner)
	// require.Equal(t, arg.Balance, account.Balance)
	// require.Equal(t, arg.Currency, account.Currency)

	// require.NotZero(t, account.ID)
	// require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	payload, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, account.ID, payload.ID)
	require.Equal(t, account.Owner, payload.Owner)
	require.Equal(t, account.Balance, payload.Balance)
	require.Equal(t, account.Currency, payload.Currency)
	require.WithinDuration(t, account.CreatedAt, payload.CreatedAt, time.Second) // This line is to validate 2 timestamps within delta duration (in this case is second)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	request := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	payload, err := testQueries.UpdateAccount(context.Background(), request)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, account.ID, payload.ID)
	require.Equal(t, account.Owner, payload.Owner)
	require.Equal(t, request.Balance, payload.Balance)
	require.Equal(t, account.Currency, payload.Currency)
	require.WithinDuration(t, account.CreatedAt, payload.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	validateAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error()) // To validate specific type of error (in this case, no rows returned)
	require.Empty(t, validateAccount)
}

func TestListAccounts(t *testing.T) {
	// Optional : Add loop to create multiple random accounts

	request := ListAccountsParams{
		Limit:  3,
		Offset: 1,
	}

	payload, err := testQueries.ListAccounts(context.TODO(), request)
	require.NoError(t, err)
	require.Len(t, payload, 3)

	// Check individual records
	for _, record := range payload {
		require.NotEmpty(t, record)
	}
}
