package cqrs

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// --- COMMAND SIDE (Write Model) ---

type CreateAccountCommand struct {
	ID      string
	Owner   string
	Balance float64
}

type DepositMoneyCommand struct {
	ID     string
	Amount float64
}

type AccountWriteModel struct {
	ID        string
	Owner     string
	Balance   float64
	Version   int
	UpdatedAt time.Time
}

type WriteStore struct {
	mu       sync.Mutex
	accounts map[string]*AccountWriteModel
}

func NewWriteStore() *WriteStore {
	return &WriteStore{
		accounts: make(map[string]*AccountWriteModel),
	}
}

// --- QUERY SIDE (Read Model View) ---

type AccountReadView struct {
	ID               string  `json:"id"`
	Owner            string  `json:"owner"`
	CurrentBalance   float64 `json:"current_balance"`
	TotalDeposits    int     `json:"total_deposits"`
	LastActivityTime string  `json:"last_activity_time"`
}

type ReadStore struct {
	mu    sync.RWMutex
	views map[string]*AccountReadView
}

func NewReadStore() *ReadStore {
	return &ReadStore{
		views: make(map[string]*AccountReadView),
	}
}

// --- CQRS SERVICE & PROJECTION ENGINE ---

type CQRSService struct {
	writeStore *WriteStore
	readStore  *ReadStore
}

func NewCQRSService(writeStore *WriteStore, readStore *ReadStore) *CQRSService {
	return &CQRSService{
		writeStore: writeStore,
		readStore:  readStore,
	}
}

// HandleCreateAccount processes a write command and projects it to the read model.
func (s *CQRSService) HandleCreateAccount(cmd CreateAccountCommand) error {
	if cmd.ID == "" || cmd.Owner == "" {
		return errors.New("invalid command parameters")
	}

	s.writeStore.mu.Lock()
	if _, exists := s.writeStore.accounts[cmd.ID]; exists {
		s.writeStore.mu.Unlock()
		return fmt.Errorf("account ID %s already exists", cmd.ID)
	}

	writeModel := &AccountWriteModel{
		ID:        cmd.ID,
		Owner:     cmd.Owner,
		Balance:   cmd.Balance,
		Version:   1,
		UpdatedAt: time.Now(),
	}
	s.writeStore.accounts[cmd.ID] = writeModel
	s.writeStore.mu.Unlock()

	// Project Event to Read Store (Query View)
	s.projectReadView(writeModel)
	return nil
}

// HandleDepositMoney processes a state-modifying write command.
func (s *CQRSService) HandleDepositMoney(cmd DepositMoneyCommand) error {
	if cmd.Amount <= 0 {
		return errors.New("deposit amount must be positive")
	}

	s.writeStore.mu.Lock()
	acc, exists := s.writeStore.accounts[cmd.ID]
	if !exists {
		s.writeStore.mu.Unlock()
		return errors.New("account not found in write store")
	}

	acc.Balance += cmd.Amount
	acc.Version++
	acc.UpdatedAt = time.Now()
	s.writeStore.mu.Unlock()

	// Update Read Model Projection
	s.projectReadView(acc)
	return nil
}

func (s *CQRSService) projectReadView(model *AccountWriteModel) {
	s.readStore.mu.Lock()
	defer s.readStore.mu.Unlock()

	view, exists := s.readStore.views[model.ID]
	if !exists {
		s.readStore.views[model.ID] = &AccountReadView{
			ID:               model.ID,
			Owner:            model.Owner,
			CurrentBalance:   model.Balance,
			TotalDeposits:    0,
			LastActivityTime: model.UpdatedAt.Format(time.RFC3339),
		}
	} else {
		view.CurrentBalance = model.Balance
		view.TotalDeposits++
		view.LastActivityTime = model.UpdatedAt.Format(time.RFC3339)
	}
}

// GetAccountView queries the read store (fast, lock-free projection).
func (s *CQRSService) GetAccountView(id string) (*AccountReadView, error) {
	s.readStore.mu.RLock()
	defer s.readStore.mu.RUnlock()

	view, exists := s.readStore.views[id]
	if !exists {
		return nil, errors.New("account view not found")
	}

	// Return copy to prevent caller mutation
	copyView := *view
	return &copyView, nil
}
