package repository

import (
	"fmt"
	"github.com/arthur-teixeira/imersao/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDB struct {
	model.TransactionRepositoryInterface
	Db *gorm.DB
}

func (t *TransactionRepositoryDB) Register(transaction *model.Transaction) error {
	return t.Db.Create(transaction).Error
}

func (t *TransactionRepositoryDB) Save(transaction *model.Transaction) error {
	return t.Db.Save(transaction).Error
}

func (t *TransactionRepositoryDB) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	t.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}
	return &transaction, nil
}
