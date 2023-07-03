// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/deposit"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/goodbenefit"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/limitation"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/payment"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/platform"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/transfer"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// Deposit is the client for interacting with the Deposit builders.
	Deposit *DepositClient
	// GoodBenefit is the client for interacting with the GoodBenefit builders.
	GoodBenefit *GoodBenefitClient
	// Limitation is the client for interacting with the Limitation builders.
	Limitation *LimitationClient
	// Payment is the client for interacting with the Payment builders.
	Payment *PaymentClient
	// Platform is the client for interacting with the Platform builders.
	Platform *PlatformClient
	// Transfer is the client for interacting with the Transfer builders.
	Transfer *TransferClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.Deposit = NewDepositClient(c.config)
	c.GoodBenefit = NewGoodBenefitClient(c.config)
	c.Limitation = NewLimitationClient(c.config)
	c.Payment = NewPaymentClient(c.config)
	c.Platform = NewPlatformClient(c.config)
	c.Transfer = NewTransferClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Account:     NewAccountClient(cfg),
		Deposit:     NewDepositClient(cfg),
		GoodBenefit: NewGoodBenefitClient(cfg),
		Limitation:  NewLimitationClient(cfg),
		Payment:     NewPaymentClient(cfg),
		Platform:    NewPlatformClient(cfg),
		Transfer:    NewTransferClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Account:     NewAccountClient(cfg),
		Deposit:     NewDepositClient(cfg),
		GoodBenefit: NewGoodBenefitClient(cfg),
		Limitation:  NewLimitationClient(cfg),
		Payment:     NewPaymentClient(cfg),
		Platform:    NewPlatformClient(cfg),
		Transfer:    NewTransferClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.Deposit.Use(hooks...)
	c.GoodBenefit.Use(hooks...)
	c.Limitation.Use(hooks...)
	c.Payment.Use(hooks...)
	c.Platform.Use(hooks...)
	c.Transfer.Use(hooks...)
	c.User.Use(hooks...)
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Create returns a builder for creating a Account entity.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id uuid.UUID) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AccountClient) DeleteOneID(id uuid.UUID) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id uuid.UUID) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id uuid.UUID) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	hooks := c.hooks.Account
	return append(hooks[:len(hooks):len(hooks)], account.Hooks[:]...)
}

// DepositClient is a client for the Deposit schema.
type DepositClient struct {
	config
}

// NewDepositClient returns a client for the Deposit from the given config.
func NewDepositClient(c config) *DepositClient {
	return &DepositClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `deposit.Hooks(f(g(h())))`.
func (c *DepositClient) Use(hooks ...Hook) {
	c.hooks.Deposit = append(c.hooks.Deposit, hooks...)
}

// Create returns a builder for creating a Deposit entity.
func (c *DepositClient) Create() *DepositCreate {
	mutation := newDepositMutation(c.config, OpCreate)
	return &DepositCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Deposit entities.
func (c *DepositClient) CreateBulk(builders ...*DepositCreate) *DepositCreateBulk {
	return &DepositCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Deposit.
func (c *DepositClient) Update() *DepositUpdate {
	mutation := newDepositMutation(c.config, OpUpdate)
	return &DepositUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DepositClient) UpdateOne(d *Deposit) *DepositUpdateOne {
	mutation := newDepositMutation(c.config, OpUpdateOne, withDeposit(d))
	return &DepositUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DepositClient) UpdateOneID(id uuid.UUID) *DepositUpdateOne {
	mutation := newDepositMutation(c.config, OpUpdateOne, withDepositID(id))
	return &DepositUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Deposit.
func (c *DepositClient) Delete() *DepositDelete {
	mutation := newDepositMutation(c.config, OpDelete)
	return &DepositDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DepositClient) DeleteOne(d *Deposit) *DepositDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *DepositClient) DeleteOneID(id uuid.UUID) *DepositDeleteOne {
	builder := c.Delete().Where(deposit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DepositDeleteOne{builder}
}

// Query returns a query builder for Deposit.
func (c *DepositClient) Query() *DepositQuery {
	return &DepositQuery{
		config: c.config,
	}
}

// Get returns a Deposit entity by its id.
func (c *DepositClient) Get(ctx context.Context, id uuid.UUID) (*Deposit, error) {
	return c.Query().Where(deposit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DepositClient) GetX(ctx context.Context, id uuid.UUID) *Deposit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DepositClient) Hooks() []Hook {
	hooks := c.hooks.Deposit
	return append(hooks[:len(hooks):len(hooks)], deposit.Hooks[:]...)
}

// GoodBenefitClient is a client for the GoodBenefit schema.
type GoodBenefitClient struct {
	config
}

// NewGoodBenefitClient returns a client for the GoodBenefit from the given config.
func NewGoodBenefitClient(c config) *GoodBenefitClient {
	return &GoodBenefitClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `goodbenefit.Hooks(f(g(h())))`.
func (c *GoodBenefitClient) Use(hooks ...Hook) {
	c.hooks.GoodBenefit = append(c.hooks.GoodBenefit, hooks...)
}

// Create returns a builder for creating a GoodBenefit entity.
func (c *GoodBenefitClient) Create() *GoodBenefitCreate {
	mutation := newGoodBenefitMutation(c.config, OpCreate)
	return &GoodBenefitCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GoodBenefit entities.
func (c *GoodBenefitClient) CreateBulk(builders ...*GoodBenefitCreate) *GoodBenefitCreateBulk {
	return &GoodBenefitCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GoodBenefit.
func (c *GoodBenefitClient) Update() *GoodBenefitUpdate {
	mutation := newGoodBenefitMutation(c.config, OpUpdate)
	return &GoodBenefitUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GoodBenefitClient) UpdateOne(gb *GoodBenefit) *GoodBenefitUpdateOne {
	mutation := newGoodBenefitMutation(c.config, OpUpdateOne, withGoodBenefit(gb))
	return &GoodBenefitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GoodBenefitClient) UpdateOneID(id uuid.UUID) *GoodBenefitUpdateOne {
	mutation := newGoodBenefitMutation(c.config, OpUpdateOne, withGoodBenefitID(id))
	return &GoodBenefitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GoodBenefit.
func (c *GoodBenefitClient) Delete() *GoodBenefitDelete {
	mutation := newGoodBenefitMutation(c.config, OpDelete)
	return &GoodBenefitDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GoodBenefitClient) DeleteOne(gb *GoodBenefit) *GoodBenefitDeleteOne {
	return c.DeleteOneID(gb.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GoodBenefitClient) DeleteOneID(id uuid.UUID) *GoodBenefitDeleteOne {
	builder := c.Delete().Where(goodbenefit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GoodBenefitDeleteOne{builder}
}

// Query returns a query builder for GoodBenefit.
func (c *GoodBenefitClient) Query() *GoodBenefitQuery {
	return &GoodBenefitQuery{
		config: c.config,
	}
}

// Get returns a GoodBenefit entity by its id.
func (c *GoodBenefitClient) Get(ctx context.Context, id uuid.UUID) (*GoodBenefit, error) {
	return c.Query().Where(goodbenefit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GoodBenefitClient) GetX(ctx context.Context, id uuid.UUID) *GoodBenefit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GoodBenefitClient) Hooks() []Hook {
	hooks := c.hooks.GoodBenefit
	return append(hooks[:len(hooks):len(hooks)], goodbenefit.Hooks[:]...)
}

// LimitationClient is a client for the Limitation schema.
type LimitationClient struct {
	config
}

// NewLimitationClient returns a client for the Limitation from the given config.
func NewLimitationClient(c config) *LimitationClient {
	return &LimitationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `limitation.Hooks(f(g(h())))`.
func (c *LimitationClient) Use(hooks ...Hook) {
	c.hooks.Limitation = append(c.hooks.Limitation, hooks...)
}

// Create returns a builder for creating a Limitation entity.
func (c *LimitationClient) Create() *LimitationCreate {
	mutation := newLimitationMutation(c.config, OpCreate)
	return &LimitationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Limitation entities.
func (c *LimitationClient) CreateBulk(builders ...*LimitationCreate) *LimitationCreateBulk {
	return &LimitationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Limitation.
func (c *LimitationClient) Update() *LimitationUpdate {
	mutation := newLimitationMutation(c.config, OpUpdate)
	return &LimitationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LimitationClient) UpdateOne(l *Limitation) *LimitationUpdateOne {
	mutation := newLimitationMutation(c.config, OpUpdateOne, withLimitation(l))
	return &LimitationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LimitationClient) UpdateOneID(id uuid.UUID) *LimitationUpdateOne {
	mutation := newLimitationMutation(c.config, OpUpdateOne, withLimitationID(id))
	return &LimitationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Limitation.
func (c *LimitationClient) Delete() *LimitationDelete {
	mutation := newLimitationMutation(c.config, OpDelete)
	return &LimitationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LimitationClient) DeleteOne(l *Limitation) *LimitationDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *LimitationClient) DeleteOneID(id uuid.UUID) *LimitationDeleteOne {
	builder := c.Delete().Where(limitation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LimitationDeleteOne{builder}
}

// Query returns a query builder for Limitation.
func (c *LimitationClient) Query() *LimitationQuery {
	return &LimitationQuery{
		config: c.config,
	}
}

// Get returns a Limitation entity by its id.
func (c *LimitationClient) Get(ctx context.Context, id uuid.UUID) (*Limitation, error) {
	return c.Query().Where(limitation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LimitationClient) GetX(ctx context.Context, id uuid.UUID) *Limitation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LimitationClient) Hooks() []Hook {
	hooks := c.hooks.Limitation
	return append(hooks[:len(hooks):len(hooks)], limitation.Hooks[:]...)
}

// PaymentClient is a client for the Payment schema.
type PaymentClient struct {
	config
}

// NewPaymentClient returns a client for the Payment from the given config.
func NewPaymentClient(c config) *PaymentClient {
	return &PaymentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `payment.Hooks(f(g(h())))`.
func (c *PaymentClient) Use(hooks ...Hook) {
	c.hooks.Payment = append(c.hooks.Payment, hooks...)
}

// Create returns a builder for creating a Payment entity.
func (c *PaymentClient) Create() *PaymentCreate {
	mutation := newPaymentMutation(c.config, OpCreate)
	return &PaymentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Payment entities.
func (c *PaymentClient) CreateBulk(builders ...*PaymentCreate) *PaymentCreateBulk {
	return &PaymentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Payment.
func (c *PaymentClient) Update() *PaymentUpdate {
	mutation := newPaymentMutation(c.config, OpUpdate)
	return &PaymentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PaymentClient) UpdateOne(pa *Payment) *PaymentUpdateOne {
	mutation := newPaymentMutation(c.config, OpUpdateOne, withPayment(pa))
	return &PaymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PaymentClient) UpdateOneID(id uuid.UUID) *PaymentUpdateOne {
	mutation := newPaymentMutation(c.config, OpUpdateOne, withPaymentID(id))
	return &PaymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Payment.
func (c *PaymentClient) Delete() *PaymentDelete {
	mutation := newPaymentMutation(c.config, OpDelete)
	return &PaymentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PaymentClient) DeleteOne(pa *Payment) *PaymentDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PaymentClient) DeleteOneID(id uuid.UUID) *PaymentDeleteOne {
	builder := c.Delete().Where(payment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PaymentDeleteOne{builder}
}

// Query returns a query builder for Payment.
func (c *PaymentClient) Query() *PaymentQuery {
	return &PaymentQuery{
		config: c.config,
	}
}

// Get returns a Payment entity by its id.
func (c *PaymentClient) Get(ctx context.Context, id uuid.UUID) (*Payment, error) {
	return c.Query().Where(payment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PaymentClient) GetX(ctx context.Context, id uuid.UUID) *Payment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PaymentClient) Hooks() []Hook {
	hooks := c.hooks.Payment
	return append(hooks[:len(hooks):len(hooks)], payment.Hooks[:]...)
}

// PlatformClient is a client for the Platform schema.
type PlatformClient struct {
	config
}

// NewPlatformClient returns a client for the Platform from the given config.
func NewPlatformClient(c config) *PlatformClient {
	return &PlatformClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `platform.Hooks(f(g(h())))`.
func (c *PlatformClient) Use(hooks ...Hook) {
	c.hooks.Platform = append(c.hooks.Platform, hooks...)
}

// Create returns a builder for creating a Platform entity.
func (c *PlatformClient) Create() *PlatformCreate {
	mutation := newPlatformMutation(c.config, OpCreate)
	return &PlatformCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Platform entities.
func (c *PlatformClient) CreateBulk(builders ...*PlatformCreate) *PlatformCreateBulk {
	return &PlatformCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Platform.
func (c *PlatformClient) Update() *PlatformUpdate {
	mutation := newPlatformMutation(c.config, OpUpdate)
	return &PlatformUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlatformClient) UpdateOne(pl *Platform) *PlatformUpdateOne {
	mutation := newPlatformMutation(c.config, OpUpdateOne, withPlatform(pl))
	return &PlatformUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PlatformClient) UpdateOneID(id uuid.UUID) *PlatformUpdateOne {
	mutation := newPlatformMutation(c.config, OpUpdateOne, withPlatformID(id))
	return &PlatformUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Platform.
func (c *PlatformClient) Delete() *PlatformDelete {
	mutation := newPlatformMutation(c.config, OpDelete)
	return &PlatformDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PlatformClient) DeleteOne(pl *Platform) *PlatformDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PlatformClient) DeleteOneID(id uuid.UUID) *PlatformDeleteOne {
	builder := c.Delete().Where(platform.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlatformDeleteOne{builder}
}

// Query returns a query builder for Platform.
func (c *PlatformClient) Query() *PlatformQuery {
	return &PlatformQuery{
		config: c.config,
	}
}

// Get returns a Platform entity by its id.
func (c *PlatformClient) Get(ctx context.Context, id uuid.UUID) (*Platform, error) {
	return c.Query().Where(platform.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlatformClient) GetX(ctx context.Context, id uuid.UUID) *Platform {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PlatformClient) Hooks() []Hook {
	hooks := c.hooks.Platform
	return append(hooks[:len(hooks):len(hooks)], platform.Hooks[:]...)
}

// TransferClient is a client for the Transfer schema.
type TransferClient struct {
	config
}

// NewTransferClient returns a client for the Transfer from the given config.
func NewTransferClient(c config) *TransferClient {
	return &TransferClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `transfer.Hooks(f(g(h())))`.
func (c *TransferClient) Use(hooks ...Hook) {
	c.hooks.Transfer = append(c.hooks.Transfer, hooks...)
}

// Create returns a builder for creating a Transfer entity.
func (c *TransferClient) Create() *TransferCreate {
	mutation := newTransferMutation(c.config, OpCreate)
	return &TransferCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Transfer entities.
func (c *TransferClient) CreateBulk(builders ...*TransferCreate) *TransferCreateBulk {
	return &TransferCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Transfer.
func (c *TransferClient) Update() *TransferUpdate {
	mutation := newTransferMutation(c.config, OpUpdate)
	return &TransferUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TransferClient) UpdateOne(t *Transfer) *TransferUpdateOne {
	mutation := newTransferMutation(c.config, OpUpdateOne, withTransfer(t))
	return &TransferUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TransferClient) UpdateOneID(id uuid.UUID) *TransferUpdateOne {
	mutation := newTransferMutation(c.config, OpUpdateOne, withTransferID(id))
	return &TransferUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Transfer.
func (c *TransferClient) Delete() *TransferDelete {
	mutation := newTransferMutation(c.config, OpDelete)
	return &TransferDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TransferClient) DeleteOne(t *Transfer) *TransferDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TransferClient) DeleteOneID(id uuid.UUID) *TransferDeleteOne {
	builder := c.Delete().Where(transfer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TransferDeleteOne{builder}
}

// Query returns a query builder for Transfer.
func (c *TransferClient) Query() *TransferQuery {
	return &TransferQuery{
		config: c.config,
	}
}

// Get returns a Transfer entity by its id.
func (c *TransferClient) Get(ctx context.Context, id uuid.UUID) (*Transfer, error) {
	return c.Query().Where(transfer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TransferClient) GetX(ctx context.Context, id uuid.UUID) *Transfer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TransferClient) Hooks() []Hook {
	hooks := c.hooks.Transfer
	return append(hooks[:len(hooks):len(hooks)], transfer.Hooks[:]...)
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	hooks := c.hooks.User
	return append(hooks[:len(hooks):len(hooks)], user.Hooks[:]...)
}
