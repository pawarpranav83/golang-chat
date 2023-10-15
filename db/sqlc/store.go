package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions.
// We only have one transaction, that is, whenever a room is created, we must add it in rooms table, and one entry in userroom table, with room id from rooms table and user id of user that created that room.
type Store struct {
	// All methods for Queries will be accessbile for Store.
	// Provides functions for db queries.
	*Queries

	// Required to create a new db tx.
	db *sql.DB
}

// Creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a databse transaction.
// Second argument is a callback function.
// It starts a new db tx and create a new queries object with that tx, and call the callback function with the created queries and finally commit or rollback the tx based on the error returned.
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// second arg is TxOptions allows to set a custom isolation level for the tx. If not specified uses default isolation level.
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// This struct contains all params required for room creation, that is, params for room create and create userroom
type RoomCreationTxParams struct {
	Roomname string `json:"roomname"`
	UserID   int64  `json:"user_id"`
}

type RoomCreationTxResult struct {
	Room     Room     `json:"room"`
	Userroom Userroom `json:"userroom"`
}

// RoomCreationTx creates a room, adds teh room along with userid that created the room in the userroom.
func (store *Store) RoomCreationTx(ctx context.Context, arg RoomCreationTxParams) (RoomCreationTxResult, error) {
	var result RoomCreationTxResult

	// Doubt - Implements Closure #6
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Room, err = q.CreateRoom(ctx, arg.Roomname)

		if err != nil {
			return err
		}

		result.Userroom, err = q.AddUsertoRoom(ctx, AddUsertoRoomParams{
			RoomID: result.Room.ID,
			UserID: arg.UserID,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
