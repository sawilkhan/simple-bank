package db

import (
	"context"
	"testing"
	"time"

	"github.com/sawilkhan/simple-bank/db/util"
	"github.com/stretchr/testify/require"
)



func createRandomEntry(t *testing.T) Entry{
	arg1 := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	
	account, err := testQueries.CreateAccount(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	
	arg2 := CreateEntryParams{
		AccountID: account.ID,
		Amount: 1000,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg2.AccountID, entry.AccountID)
	require.NotEmpty(t, entry.CreatedAt)
	require.NotEmpty(t, entry.ID)
	return entry
}

func createMultipleEntries(AccountID int64){
	for i := 0; i < 10; i++{
		arg := CreateEntryParams{
			AccountID: AccountID,
			Amount: util.RandomMoney(),
		}
		testQueries.CreateEntry(context.Background(), arg)
	}
}

func TestCreateEntry(t *testing.T){
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T){
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T){
	account := createRandomAccount(t)
	createMultipleEntries(account.ID)
	
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit: 5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries{
		require.NotEmpty(t, entry)
	}
}