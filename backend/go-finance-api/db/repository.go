package db

// Repositories holds all the repository instances
type Repositories struct {
	User          *UserRepository
	Item          *ItemRepository
	Account       *AccountRepository
	Transaction   *TransactionRepository
	PlaidAPIEvent *PlaidAPIEventRepository
	LinkEvent     *LinkEventRepository
}

// NewRepositories creates a new Repositories instance
func NewRepositories(db *Database) *Repositories {
	return &Repositories{
		User:          NewUserRepository(db),
		Item:          NewItemRepository(db),
		Account:       NewAccountRepository(db),
		Transaction:   NewTransactionRepository(db),
		PlaidAPIEvent: NewPlaidAPIEventRepository(db),
		LinkEvent:     NewLinkEventRepository(db),
	}
}
