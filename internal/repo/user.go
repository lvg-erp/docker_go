package repo

import "context"

type UserRepo struct {
	db *Database
}

func NewUserRepo(db *Database) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (d *UserRepo) CreateUser(ctx context.Context, user *User) error {
	newUser := &User{
		Username: user.Username,
		Email:    user.Email,
		IsActive: user.IsActive,
	}

	if err := d.db.Client.WithContext(ctx).Create(newUser).Error; err != nil {
		return err
	}

	return nil
}

func (d *UserRepo) GetUserByID(ctx context.Context, id int64) (User, error) {
	user := User{}
	if err := d.db.Client.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return User{}, err
	}

	return User{
		Username: user.Username,
		Email:    user.Email,
		IsActive: user.IsActive,
	}, nil

}

func (d *UserRepo) UpdateUser(ctx context.Context, uUser *User, id uint) error {
	var existingUser User
	if err := d.db.Client.WithContext(ctx).Where("id = ?", id).First(&existingUser).Error; err != nil {
		return err
	}

	existingUser.Username = uUser.Username
	existingUser.Email = uUser.Email
	existingUser.IsActive = uUser.IsActive
	//
	if err := d.db.Client.WithContext(ctx).Save(&existingUser).Error; err != nil {
		return err
	}

	return nil
}

func (d *UserRepo) DeleteUser(ctx context.Context, id uint) error {
	var existingUser User
	if err := d.db.Client.WithContext(ctx).Where("id = ?", id).First(&existingUser).Error; err != nil {
		return err
	}

	if err := d.db.Client.WithContext(ctx).Delete(&existingUser).Error; err != nil {
		return err
	}

	return nil

}
