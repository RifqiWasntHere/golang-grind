package db

import (
	"context"
	"database/sql"
	"fmt"
)

// `Store` provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// `NewStore` creates a new `Store`
func Newstore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// `execTx` a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error : %v, rb error : %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// Money transferring mechanism : 1. Creates a transfer record, 2. Add account entries, 3. Update account balances
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// Create sender's entry
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// Create receiver's entry
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		// get account, then update its balance
		// fromAccount, err := q.GetAccountForUpdate(context.Background(), arg.FromAccountID)
		// if err != nil {
		// 	return err
		// }

		// result.FromAccount, err = q.UpdateAccount(context.Background(), UpdateAccountParams{
		// 	ID:      fromAccount.ID,
		// 	Balance: fromAccount.Balance - arg.Amount,
		// })
		// if err != nil {
		// 	return err
		// }

		// toAccount, err := q.GetAccountForUpdate(context.Background(), arg.ToAccountID)
		// if err != nil {
		// 	return err
		// }

		// result.ToAccount, err = q.UpdateAccount(context.Background(), UpdateAccountParams{
		// 	ID:      toAccount.ID,
		// 	Balance: toAccount.Balance + arg.Amount,
		// })
		// if err != nil {
		// 	return err
		// }
		result.FromAccount, err = q.AddAccountBalance(context.Background(), AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalance(context.Background(), AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
