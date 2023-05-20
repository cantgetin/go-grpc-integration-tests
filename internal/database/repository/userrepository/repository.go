package userrepository

import "context"

func (r *RepositoryImpl) GetUser(ctx context.Context, id int32) (*User, error) {
	var user *User

	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *RepositoryImpl) ListUsers(ctx context.Context) ([]*User, error) {
	var users []*User

	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
