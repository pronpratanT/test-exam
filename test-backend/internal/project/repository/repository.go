package repository

import "gorm.io/gorm"

type TestRepository struct {
	DB *gorm.DB
}

func NewTestExamRepository(db *gorm.DB) *TestRepository {
	return &TestRepository{DB: db}
}
