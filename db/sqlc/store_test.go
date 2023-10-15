package db

import (
	"context"
	"testing"

	"github.com/pawarpranav83/golang-chat/db/util"
	"github.com/stretchr/testify/require"
)

func TestRoomCreationTx(t *testing.T) {
	store := NewStore(testDB)

	user := createRandomUser(t)

	// Testing the transaction by running n concurrent go routines performing that tx.
	n := 5

	// We cannot use testify require to check the errors, because the func is running inside diff go routine from the one that the TestRoomCreationTx func is running on, so no guarantee that it will stop whole test if condition is not satisfied.
	// So we send the err to the main go routine where the test is actually running on.
	errs := make(chan error)
	results := make(chan RoomCreationTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.RoomCreationTx(context.Background(), RoomCreationTxParams{
				Roomname: util.RandomString(4),
				UserID:   user.ID,
			})

			errs <- err
			results <- result
		}()
	}

	// checking results from results chan
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check if room is created
		room := result.Room
		require.NotEmpty(t, room)
		// require.Equal(t, "room"+fmt.Sprint(i+1), room.Roomname)
		require.NotZero(t, room.ID)

		_, err = store.GetRoom(context.Background(), room.ID)
		require.NoError(t, err)

		// check if userroom is created
		userroom := result.Userroom
		require.Equal(t, room.ID, userroom.RoomID)
		require.Equal(t, user.ID, userroom.UserID)

		_, err = store.GetRoomusers(context.Background(), userroom.RoomID)
		require.NoError(t, err)
	}

}
