package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v6"
)

// ReputationConfig is an object representing the database table.
type ReputationConfig struct {
	GuildID                int64       `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	PointsName             string      `boil:"points_name" json:"points_name" toml:"points_name" yaml:"points_name"`
	Enabled                bool        `boil:"enabled" json:"enabled" toml:"enabled" yaml:"enabled"`
	Cooldown               int         `boil:"cooldown" json:"cooldown" toml:"cooldown" yaml:"cooldown"`
	MaxGiveAmount          int64       `boil:"max_give_amount" json:"max_give_amount" toml:"max_give_amount" yaml:"max_give_amount"`
	RequiredGiveRole       null.String `boil:"required_give_role" json:"required_give_role,omitempty" toml:"required_give_role" yaml:"required_give_role,omitempty"`
	RequiredReceiveRole    null.String `boil:"required_receive_role" json:"required_receive_role,omitempty" toml:"required_receive_role" yaml:"required_receive_role,omitempty"`
	BlacklistedGiveRole    null.String `boil:"blacklisted_give_role" json:"blacklisted_give_role,omitempty" toml:"blacklisted_give_role" yaml:"blacklisted_give_role,omitempty"`
	BlacklistedReceiveRole null.String `boil:"blacklisted_receive_role" json:"blacklisted_receive_role,omitempty" toml:"blacklisted_receive_role" yaml:"blacklisted_receive_role,omitempty"`
	AdminRole              null.String `boil:"admin_role" json:"admin_role,omitempty" toml:"admin_role" yaml:"admin_role,omitempty"`

	R *reputationConfigR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L reputationConfigL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// reputationConfigR is where relationships are stored.
type reputationConfigR struct {
}

// reputationConfigL is where Load methods for each relationship are stored.
type reputationConfigL struct{}

var (
	reputationConfigColumns               = []string{"guild_id", "points_name", "enabled", "cooldown", "max_give_amount", "required_give_role", "required_receive_role", "blacklisted_give_role", "blacklisted_receive_role", "admin_role"}
	reputationConfigColumnsWithoutDefault = []string{"guild_id", "points_name", "enabled", "cooldown", "max_give_amount", "required_give_role", "required_receive_role", "blacklisted_give_role", "blacklisted_receive_role", "admin_role"}
	reputationConfigColumnsWithDefault    = []string{}
	reputationConfigPrimaryKeyColumns     = []string{"guild_id"}
)

type (
	// ReputationConfigSlice is an alias for a slice of pointers to ReputationConfig.
	// This should generally be used opposed to []ReputationConfig.
	ReputationConfigSlice []*ReputationConfig

	reputationConfigQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	reputationConfigType                 = reflect.TypeOf(&ReputationConfig{})
	reputationConfigMapping              = queries.MakeStructMapping(reputationConfigType)
	reputationConfigPrimaryKeyMapping, _ = queries.BindMapping(reputationConfigType, reputationConfigMapping, reputationConfigPrimaryKeyColumns)
	reputationConfigInsertCacheMut       sync.RWMutex
	reputationConfigInsertCache          = make(map[string]insertCache)
	reputationConfigUpdateCacheMut       sync.RWMutex
	reputationConfigUpdateCache          = make(map[string]updateCache)
	reputationConfigUpsertCacheMut       sync.RWMutex
	reputationConfigUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single reputationConfig record from the query, and panics on error.
func (q reputationConfigQuery) OneP() *ReputationConfig {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single reputationConfig record from the query.
func (q reputationConfigQuery) One() (*ReputationConfig, error) {
	o := &ReputationConfig{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for reputation_configs")
	}

	return o, nil
}

// AllP returns all ReputationConfig records from the query, and panics on error.
func (q reputationConfigQuery) AllP() ReputationConfigSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all ReputationConfig records from the query.
func (q reputationConfigQuery) All() (ReputationConfigSlice, error) {
	var o ReputationConfigSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ReputationConfig slice")
	}

	return o, nil
}

// CountP returns the count of all ReputationConfig records in the query, and panics on error.
func (q reputationConfigQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all ReputationConfig records in the query.
func (q reputationConfigQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count reputation_configs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q reputationConfigQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q reputationConfigQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if reputation_configs exists")
	}

	return count > 0, nil
}

// ReputationConfigsG retrieves all records.
func ReputationConfigsG(mods ...qm.QueryMod) reputationConfigQuery {
	return ReputationConfigs(boil.GetDB(), mods...)
}

// ReputationConfigs retrieves all the records using an executor.
func ReputationConfigs(exec boil.Executor, mods ...qm.QueryMod) reputationConfigQuery {
	mods = append(mods, qm.From("\"reputation_configs\""))
	return reputationConfigQuery{NewQuery(exec, mods...)}
}

// FindReputationConfigG retrieves a single record by ID.
func FindReputationConfigG(guildID int64, selectCols ...string) (*ReputationConfig, error) {
	return FindReputationConfig(boil.GetDB(), guildID, selectCols...)
}

// FindReputationConfigGP retrieves a single record by ID, and panics on error.
func FindReputationConfigGP(guildID int64, selectCols ...string) *ReputationConfig {
	retobj, err := FindReputationConfig(boil.GetDB(), guildID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindReputationConfig retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindReputationConfig(exec boil.Executor, guildID int64, selectCols ...string) (*ReputationConfig, error) {
	reputationConfigObj := &ReputationConfig{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"reputation_configs\" where \"guild_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, guildID)

	err := q.Bind(reputationConfigObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from reputation_configs")
	}

	return reputationConfigObj, nil
}

// FindReputationConfigP retrieves a single record by ID with an executor, and panics on error.
func FindReputationConfigP(exec boil.Executor, guildID int64, selectCols ...string) *ReputationConfig {
	retobj, err := FindReputationConfig(exec, guildID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ReputationConfig) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *ReputationConfig) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *ReputationConfig) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *ReputationConfig) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no reputation_configs provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(reputationConfigColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	reputationConfigInsertCacheMut.RLock()
	cache, cached := reputationConfigInsertCache[key]
	reputationConfigInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			reputationConfigColumns,
			reputationConfigColumnsWithDefault,
			reputationConfigColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(reputationConfigType, reputationConfigMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(reputationConfigType, reputationConfigMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"reputation_configs\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into reputation_configs")
	}

	if !cached {
		reputationConfigInsertCacheMut.Lock()
		reputationConfigInsertCache[key] = cache
		reputationConfigInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single ReputationConfig record. See Update for
// whitelist behavior description.
func (o *ReputationConfig) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single ReputationConfig record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *ReputationConfig) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the ReputationConfig, and panics on error.
// See Update for whitelist behavior description.
func (o *ReputationConfig) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the ReputationConfig.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *ReputationConfig) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	reputationConfigUpdateCacheMut.RLock()
	cache, cached := reputationConfigUpdateCache[key]
	reputationConfigUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(reputationConfigColumns, reputationConfigPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update reputation_configs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"reputation_configs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, reputationConfigPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(reputationConfigType, reputationConfigMapping, append(wl, reputationConfigPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update reputation_configs row")
	}

	if !cached {
		reputationConfigUpdateCacheMut.Lock()
		reputationConfigUpdateCache[key] = cache
		reputationConfigUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q reputationConfigQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q reputationConfigQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for reputation_configs")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ReputationConfigSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ReputationConfigSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ReputationConfigSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ReputationConfigSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reputationConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"reputation_configs\" SET %s WHERE (\"guild_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(reputationConfigPrimaryKeyColumns), len(colNames)+1, len(reputationConfigPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in reputationConfig slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ReputationConfig) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *ReputationConfig) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *ReputationConfig) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ReputationConfig) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no reputation_configs provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(reputationConfigColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	reputationConfigUpsertCacheMut.RLock()
	cache, cached := reputationConfigUpsertCache[key]
	reputationConfigUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			reputationConfigColumns,
			reputationConfigColumnsWithDefault,
			reputationConfigColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			reputationConfigColumns,
			reputationConfigPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert reputation_configs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(reputationConfigPrimaryKeyColumns))
			copy(conflict, reputationConfigPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"reputation_configs\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(reputationConfigType, reputationConfigMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(reputationConfigType, reputationConfigMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert reputation_configs")
	}

	if !cached {
		reputationConfigUpsertCacheMut.Lock()
		reputationConfigUpsertCache[key] = cache
		reputationConfigUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single ReputationConfig record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *ReputationConfig) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single ReputationConfig record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ReputationConfig) DeleteG() error {
	if o == nil {
		return errors.New("models: no ReputationConfig provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single ReputationConfig record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *ReputationConfig) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single ReputationConfig record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ReputationConfig) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no ReputationConfig provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), reputationConfigPrimaryKeyMapping)
	sql := "DELETE FROM \"reputation_configs\" WHERE \"guild_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from reputation_configs")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q reputationConfigQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q reputationConfigQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no reputationConfigQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from reputation_configs")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ReputationConfigSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ReputationConfigSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no ReputationConfig slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ReputationConfigSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ReputationConfigSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no ReputationConfig slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reputationConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"reputation_configs\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, reputationConfigPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(reputationConfigPrimaryKeyColumns), 1, len(reputationConfigPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from reputationConfig slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *ReputationConfig) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *ReputationConfig) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ReputationConfig) ReloadG() error {
	if o == nil {
		return errors.New("models: no ReputationConfig provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ReputationConfig) Reload(exec boil.Executor) error {
	ret, err := FindReputationConfig(exec, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ReputationConfigSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ReputationConfigSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReputationConfigSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ReputationConfigSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReputationConfigSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	reputationConfigs := ReputationConfigSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reputationConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"reputation_configs\".* FROM \"reputation_configs\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, reputationConfigPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(reputationConfigPrimaryKeyColumns), 1, len(reputationConfigPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&reputationConfigs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ReputationConfigSlice")
	}

	*o = reputationConfigs

	return nil
}

// ReputationConfigExists checks if the ReputationConfig row exists.
func ReputationConfigExists(exec boil.Executor, guildID int64) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"reputation_configs\" where \"guild_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, guildID)
	}

	row := exec.QueryRow(sql, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if reputation_configs exists")
	}

	return exists, nil
}

// ReputationConfigExistsG checks if the ReputationConfig row exists.
func ReputationConfigExistsG(guildID int64) (bool, error) {
	return ReputationConfigExists(boil.GetDB(), guildID)
}

// ReputationConfigExistsGP checks if the ReputationConfig row exists. Panics on error.
func ReputationConfigExistsGP(guildID int64) bool {
	e, err := ReputationConfigExists(boil.GetDB(), guildID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ReputationConfigExistsP checks if the ReputationConfig row exists. Panics on error.
func ReputationConfigExistsP(exec boil.Executor, guildID int64) bool {
	e, err := ReputationConfigExists(exec, guildID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
