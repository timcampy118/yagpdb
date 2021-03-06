package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testReputationUsers(t *testing.T) {
	t.Parallel()

	query := ReputationUsers(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testReputationUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = reputationUser.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := ReputationUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReputationUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ReputationUsers(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := ReputationUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReputationUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ReputationUserSlice{reputationUser}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := ReputationUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testReputationUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ReputationUserExists(tx, reputationUser.UserID, reputationUser.GuildID)
	if err != nil {
		t.Errorf("Unable to check if ReputationUser exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ReputationUserExistsG to return true, but got false.")
	}
}
func testReputationUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	reputationUserFound, err := FindReputationUser(tx, reputationUser.UserID, reputationUser.GuildID)
	if err != nil {
		t.Error(err)
	}

	if reputationUserFound == nil {
		t.Error("want a record, got nil")
	}
}
func testReputationUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ReputationUsers(tx).Bind(reputationUser); err != nil {
		t.Error(err)
	}
}

func testReputationUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := ReputationUsers(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testReputationUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUserOne := &ReputationUser{}
	reputationUserTwo := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUserOne, reputationUserDBTypes, false, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}
	if err = randomize.Struct(seed, reputationUserTwo, reputationUserDBTypes, false, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUserOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = reputationUserTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := ReputationUsers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testReputationUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	reputationUserOne := &ReputationUser{}
	reputationUserTwo := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUserOne, reputationUserDBTypes, false, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}
	if err = randomize.Struct(seed, reputationUserTwo, reputationUserDBTypes, false, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUserOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = reputationUserTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ReputationUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testReputationUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ReputationUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReputationUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx, reputationUserColumns...); err != nil {
		t.Error(err)
	}

	count, err := ReputationUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReputationUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = reputationUser.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testReputationUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ReputationUserSlice{reputationUser}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testReputationUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := ReputationUsers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	reputationUserDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `GuildID`: `bigint`, `Points`: `bigint`, `UserID`: `bigint`}
	_                     = bytes.MinRead
)

func testReputationUsersUpdate(t *testing.T) {
	t.Parallel()

	if len(reputationUserColumns) == len(reputationUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ReputationUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true, reputationUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	if err = reputationUser.Update(tx); err != nil {
		t.Error(err)
	}
}

func testReputationUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(reputationUserColumns) == len(reputationUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	reputationUser := &ReputationUser{}
	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ReputationUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, reputationUser, reputationUserDBTypes, true, reputationUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(reputationUserColumns, reputationUserPrimaryKeyColumns) {
		fields = reputationUserColumns
	} else {
		fields = strmangle.SetComplement(
			reputationUserColumns,
			reputationUserPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(reputationUser))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ReputationUserSlice{reputationUser}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testReputationUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(reputationUserColumns) == len(reputationUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	reputationUser := ReputationUser{}
	if err = randomize.Struct(seed, &reputationUser, reputationUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reputationUser.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert ReputationUser: %s", err)
	}

	count, err := ReputationUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &reputationUser, reputationUserDBTypes, false, reputationUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ReputationUser struct: %s", err)
	}

	if err = reputationUser.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert ReputationUser: %s", err)
	}

	count, err = ReputationUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
