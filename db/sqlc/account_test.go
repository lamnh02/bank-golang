package db

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
	"strings"
	"time"
	"math/rand"
)

const alphabet = "abcdefghijkmlnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())	
}

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// random string length n
func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i:=0; i<n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func randomOwner() string {
	return randomString(6)
}

func randomMoney() int64 {
	return randomInt(0, 1000)
}

func randomCurrency() string {
	currencies := []string{"EURO", "USD", "VND"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: randomOwner(),
		Balance: randomMoney(),
		Currency: randomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}