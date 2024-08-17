package models

import (
	"strconv"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// User represents the user table
type (
	// userModel interface {
	// 	Insert(data *User) error
	// 	FindOne(id int64) (*User, error)
	// 	FindOneByEmail(email string) (*User, error)
	// 	Update(data *User) error
	// 	UpdateWithVersion(data *User) error
	// 	DeleteSoft(data *User) error
	// 	Delete(id int64) error
	// 	FindAll() ([]*User, error)
	// 	FindPageListByPage(page, pageSize int) ([]*User, error)
	// 	FindPageListByPageWithTotal(page, pageSize int) ([]*User, int64, error)
	// 	FindPageListByIdDESC(preMinId, pageSize int) ([]*User, error)
	// 	FindPageListByIdASC(preMaxId, pageSize int) ([]*User, error)
	// 	FindSum(field string) (float64, error)
	// 	FindCount(field string) (int64, error)
	// }

	UserModel struct {
		*gorm.DB
		*redis.Client
	}

	User struct {
		Base
		ID       int64  `gorm:"column:id;primaryKey;autoIncrement"`
		Name     string `gorm:"column:name;type:varchar(255);not null"`
		Email    string `gorm:"column:email;type:char(254);not null;unique"`
		Password string `gorm:"column:password;type:varchar(255);not null"`
		Salt     string `gorm:"column:salt;type:varchar(255);not null"`
		Sex      int8   `gorm:"column:sex;default:0;comment:性别 0:男 1:女"`
		Avatar   string `gorm:"column:avatar;type:varchar(255);not null"`
		Info     string `gorm:"column:info;type:varchar(255);not null"`
		Version  int8   `gorm:"column:version;default:0"`
	}
)

// TODO add redis cache

func NewUserModel(db *gorm.DB, redis *redis.Client) *UserModel {
	return &UserModel{
		DB:     db,
		Client: redis,
	}
}

// Insert inserts a new user record
func (m *UserModel) Insert(data *User) error {
	return m.DB.Create(data).Error
}

// FindOne finds a user record by id
func (m *UserModel) FindByID(id string) (*User, error) {
	var user User

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	err = m.DB.Where("id = ?", i).First(&user).Error
	return &user, err
}

// FindOneByEmail finds a user record by email
func (m *UserModel) FindByEmail(email string) (*User, error) {
	var user User
	err := m.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// Update updates a user record
func (m *UserModel) Update(data *User) error {
	return m.DB.Save(data).Error
}

// UpdateWithVersion updates a user record with version
func (m *UserModel) UpdateWithVersion(data *User) error {
	return m.DB.Model(data).Where("version = ?", data.Version).Update("version", data.Version+1).Error
}

// DeleteSoft deletes a user record
func (m *UserModel) DeleteSoft(data *User) error {
	return m.DB.Delete(data).Error
}

// Delete deletes a user record
func (m *UserModel) Delete(id int64) error {
	return m.DB.Where("id = ?", id).Delete(&User{}).Error
}

// FindAll finds all user records
func (m *UserModel) FindAll() ([]*User, error) {
	var users []*User
	err := m.DB.Find(&users).Error
	return users, err
}

// FindPageListByPage finds user records by page
func (m *UserModel) FindPageListByPage(page, pageSize int) ([]*User, error) {
	var users []*User
	err := m.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error
	return users, err
}

// FindPageListByPageWithTotal finds user records by page with total
func (m *UserModel) FindPageListByPageWithTotal(page, pageSize int) ([]*User, int64, error) {
	var users []*User
	var total int64
	err := m.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Count(&total).Error
	return users, total, err
}

// FindPageListByIdDESC finds user records by id desc
func (m *UserModel) FindPageListByIDDESC(preMinID, pageSize int) ([]*User, error) {
	var users []*User
	err := m.DB.Where("id < ?", preMinID).Order("id desc").Limit(pageSize).Find(&users).Error
	return users, err
}

// FindPageListByIdASC finds user records by id asc
func (m *UserModel) FindPageListByIDASC(preMaxID, pageSize int) ([]*User, error) {
	var users []*User
	err := m.DB.Where("id > ?", preMaxID).Order("id asc").Limit(pageSize).Find(&users).Error
	return users, err
}

// FindSum finds the sum of a field
func (m *UserModel) FindSum(field string) (float64, error) {
	var sum float64
	err := m.DB.Model(&User{}).Select("sum(" + field + ")").Scan(&sum).Error
	return sum, err
}

// FindCount finds the count of a field
func (m *UserModel) FindCount(field string) (int64, error) {
	var count int64
	err := m.DB.Model(&User{}).Select("count(" + field + ")").Scan(&count).Error
	return count, err
}

// TableName returns the table name of the user model
func (User) TableName() string {
	return "user"
}
