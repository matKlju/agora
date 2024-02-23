package sqlite

import "Forum/internal/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Login(email, password string) error {
	
	return nil
}

func (r *UserRepository) Register(user *model.User) error {
	queryInsert := "INSERT INTO users(UUID, email, username, password) VALUES(?, ?, ?, ?) "
	_, err := r.store.Db.Exec(queryInsert, user.UUID, user.Email, user.UserName, user.Password)
	if err != nil {
		return err
	}
	return nil
}
