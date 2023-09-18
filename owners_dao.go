package main

import (
	"database/sql"
)

type OwnerDao interface {
	GetAll() ([]*Owner, error)
	Get(id string) (*Owner, error)
	Create(owner *Owner) error
	Update(owner *Owner) error
	Delete(id string) error
	Done(id string) error
}

type OwnerDaoImpl struct {
	conn *sql.DB
}


func NewOwnerDao(conn *sql.DB) OwnerDao {
	return &OwnerDaoImpl{conn}
}

func (dao *OwnerDaoImpl) GetAll() ([]*Owner, error) {
	rows, err := dao.conn.Query("SELECT id, name, created_at, updated_at FROM owners")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	owners := []*Owner{}
	for rows.Next() {
		owner := &Owner{}
		err := rows.Scan(&owner.ID, &owner.Name, &owner.CreatedAt, &owner.UpdatedAt)
		if err != nil {
			return nil, err
		}
		owners = append(owners, owner)
	}
	return owners, nil
}

func (dao *OwnerDaoImpl) Get(id string) (*Owner, error) {
	owner := &Owner{}
	err := dao.conn.QueryRow("SELECT id, name, created_at, updated_at FROM owners WHERE id = $1", id).Scan(&owner.ID, &owner.Name, &owner.CreatedAt, &owner.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return owner, nil
}

func (dao *OwnerDaoImpl) Create(owner *Owner) error {
	_, err := dao.conn.Exec("INSERT INTO owners (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)", owner.ID, owner.Name, owner.CreatedAt, owner.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (dao *OwnerDaoImpl) Update(owner *Owner) error {
	_, err := dao.conn.Exec("UPDATE owners SET name = $1, updated_at = now() WHERE id = $2", owner.Name, owner.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *OwnerDaoImpl) Delete(id string) error {
	_, err := dao.conn.Exec("DELETE FROM owners WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// Done implements TodoDao.
func (dao *OwnerDaoImpl) Done(id string) error {
	_, err := dao.conn.Exec("UPDATE owners SET completed = true, updated_at = now() WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
