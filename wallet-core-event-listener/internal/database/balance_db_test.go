package database

import (
	"database/sql"
	"github.com/danyukod/wallet-core-event-listener/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BalanceDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	balanceDB *BalanceDB
	balance   *entity.Balance
}

func (s *BalanceDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table balances (id varchar(255), account_id varchar(255), amount int, created_at date)")
	s.balanceDB = NewBalanceDB(db)
}

func (s *BalanceDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE balances")
}

func TestBalanceDBTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceDBTestSuite))
}

func (s *BalanceDBTestSuite) TestSave() {
	balance, err := entity.NewBalance("1", 100)
	err = s.balanceDB.Save(balance)
	s.Nil(err)
}

func (s *BalanceDBTestSuite) TestFindByAccountID() {
	balance, err := entity.NewBalance("1", 100)
	err = s.balanceDB.Save(balance)
	s.Nil(err)
	balanceDB, err := s.balanceDB.FindByAccountID(balance.AccountID)
	s.Nil(err)
	s.Equal(balance.ID, balanceDB.ID)
	s.Equal(balance.AccountID, balanceDB.AccountID)
	s.Equal(balance.Amount, balanceDB.Amount)
}
