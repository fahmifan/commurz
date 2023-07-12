package preloads_test

import (
	"testing"

	"github.com/fahmifan/commurz/pkg/preloads"
)

func TestPreload(t *testing.T) {
	type Address struct {
		UserID int
		Detail string
	}

	type User struct {
		ID int

		Addresses Address
	}

	users := []User{
		{ID: 1},
		{ID: 2},
	}

	addresses := []Address{
		{UserID: 1, Detail: "Address 11"},
		{UserID: 2, Detail: "Address 22"},
	}

	preloadedUsers := preloads.Preload(users, addresses, preloads.PreloadArg[User, Address, int]{
		KeyByItem:   func(item Address) int { return item.UserID },
		KeyByTarget: func(target User) int { return target.ID },
		SetItem:     func(target *User, item Address) { target.Addresses = item },
	})

	expected := []User{
		{ID: 1, Addresses: Address{UserID: 1, Detail: "Address 11"}},
		{ID: 2, Addresses: Address{UserID: 2, Detail: "Address 22"}},
	}

	for i := range expected {
		got := preloadedUsers[i]
		want := expected[i]

		if got.ID != want.ID {
			t.Errorf("got %d, want %d", got.ID, want.ID)
		}

		if got.Addresses.UserID != want.Addresses.UserID {
			t.Errorf("got %d, want %d", got.Addresses.UserID, want.Addresses.UserID)
		}

		if got.Addresses.Detail != want.Addresses.Detail {
			t.Errorf("got %s, want %s", got.Addresses.Detail, want.Addresses.Detail)
		}
	}

}

func TestPreloadMany(t *testing.T) {
	type Address struct {
		UserID int
		Detail string
	}

	type User struct {
		ID int

		Addresses []Address
	}

	users := []User{
		{ID: 1},
		{ID: 2},
	}

	addresses := []Address{
		{UserID: 1, Detail: "Address 11"},
		{UserID: 1, Detail: "Address 22"},
		{UserID: 2, Detail: "Address 33"},
		{UserID: 2, Detail: "Address 44"},
	}

	preloadedUsers := preloads.PreloadsMany(users, addresses,
		preloads.PreloadManyArg[User, Address, int]{
			KeyByItem:   func(item Address) int { return item.UserID },
			KeyByTarget: func(target User) int { return target.ID },
			SetItem: func(target *User, items []Address) {
				target.Addresses = items
			},
		},
	)

	expected := []User{
		{
			ID: 1,
			Addresses: []Address{
				{UserID: 1, Detail: "Address 11"},
				{UserID: 1, Detail: "Address 22"},
			},
		},
		{
			ID: 2,
			Addresses: []Address{
				{UserID: 2, Detail: "Address 33"},
				{UserID: 2, Detail: "Address 44"},
			},
		},
	}

	for i := range expected {
		got := preloadedUsers[i]
		want := expected[i]

		if got.ID != want.ID {
			t.Errorf("got %d, want %d", got.ID, want.ID)
		}

		if len(got.Addresses) != len(want.Addresses) {
			t.Errorf("got %d, want %d", len(got.Addresses), len(want.Addresses))
		}

		for addrIdx := range got.Addresses {
			gotAddr := got.Addresses[addrIdx]
			wantAddr := want.Addresses[addrIdx]

			if gotAddr.UserID != wantAddr.UserID {
				t.Errorf("got %d, want %d", gotAddr.UserID, wantAddr.UserID)
			}

			if gotAddr.Detail != wantAddr.Detail {
				t.Errorf("got %s, want %s", gotAddr.Detail, wantAddr.Detail)
			}
		}

	}
}
