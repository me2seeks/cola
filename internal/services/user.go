package services

import (
	"errors"
	"strconv"

	"github.com/me2seeks/cola/internal/models"
	"github.com/me2seeks/cola/internal/pkg/logger"
	"github.com/me2seeks/cola/internal/pkg/mysql"
	"github.com/me2seeks/cola/internal/pkg/redis"
	"github.com/me2seeks/cola/internal/pkg/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/bwmarrin/snowflake"
	"github.com/me2seeks/cola/internal/types"
	goRedis "github.com/redis/go-redis/v9"
)

type UserService struct {
	Redis  *goRedis.Client
	Node   *snowflake.Node
	Logger *logrus.Entry
	Model  *models.UserModel
}

func NewUserService() *UserService {
	node, err := snowflake.NewNode(1)
	if err != nil {
		logger.Logger.Fatalf("snowflake node init failed, err: %v", err)
	}
	return &UserService{
		Redis:  redis.Client,
		Node:   node,
		Logger: logger.Logger.WithField("service", "user"),
		Model:  models.NewUserModel(mysql.Connect(), redis.Connect()),
	}
}

func (us *UserService) Register(req *types.RegisterReq) (string, error) {
	if _, err := us.Model.FindByEmail(req.Email); err != gorm.ErrRecordNotFound {
		us.Logger.Errorf("user already exists email: %s", req.Email)
		return "", errors.New("user already exists")
	}
	user := &models.User{
		ID:     us.Node.Generate().Int64(),
		Name:   req.Name,
		Email:  req.Email,
		Sex:    req.Sex,
		Avatar: req.Avatar,
		Info:   req.Info,
	}
	user.Salt = strconv.FormatInt(user.ID, 10)
	// encrypt password
	user.Password = utils.Encrypt(req.Password, user.Salt)

	if err := us.Model.Insert(user); err != nil {
		us.Logger.Errorf("insert user failed, err: %v", err)
		return "", err
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		us.Logger.Errorf("generate token failed, err: %v", err)
		return "", err
	}
	return token, nil
}

func (us *UserService) Login(req *types.LoginReq) (string, error) {
	user, err := us.Model.FindByEmail(req.Email)
	if err != nil {
		us.Logger.Errorf("find user failed, err: %v", err)
		return "", err
	}

	if user.Password != utils.Encrypt(req.Password, user.Salt) {
		us.Logger.Errorf("password is incorrect")
		return "", errors.New("password is incorrect")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		us.Logger.Errorf("generate token failed, err: %v", err)
		return "", err
	}

	return token, nil
}

func (us *UserService) GetUser(id string) (*models.User, error) {
	user, err := us.Model.FindByID(id)
	if err != nil {
		us.Logger.Errorf("find user failed, err: %v", err)
		return nil, err
	}
	return user, nil
}

func (us *UserService) UpdateUser(id string, req *types.UpdateUserReq) error {
	user, err := us.Model.FindByID(id)
	if err != nil {
		us.Logger.Errorf("find user failed, err: %v", err)
		return err
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = utils.Encrypt(req.Password, user.Salt)
	}

	// sex

	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}

	if req.Info != "" {
		user.Info = req.Info
	}

	if err := us.Model.UpdateWithVersion(user); err != nil {
		us.Logger.Errorf("update user failed, err: %v", err)
		return err
	}
	return nil
}

func (us *UserService) DeleteUser(id string) error {
	user, err := us.Model.FindByID(id)
	if err != nil {
		us.Logger.Errorf("find user failed, err: %v", err)
		return err
	}
	if err := us.Model.DeleteSoft(user); err != nil {
		us.Logger.Errorf("delete user failed, err: %v", err)
		return err
	}
	return nil
}

func (us *UserService) ListUser() ([]*models.User, error) {
	users, err := us.Model.FindAll()
	if err != nil {
		us.Logger.Errorf("list user failed, err: %v", err)
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := us.Model.FindByEmail(email)
	if err != nil {
		us.Logger.Errorf("find user failed, err: %v", err)
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetUserByToken(token string) (*models.User, error) {
	claims, err := utils.ParseToken(token)
	if err != nil {
		us.Logger.Errorf("parse token failed, err: %v", err)
		return nil, err
	}

	user, err := us.Model.FindByID(claims.Subject)
	if err != nil {
		us.Logger.Errorf("find user failed, err: %v", err)
		return nil, err
	}
	return user, nil
}

func (us *UserService) ListUserByPage(page, pageSize int) ([]*models.User, error) {
	users, err := us.Model.FindPageListByPage(page, pageSize)
	if err != nil {
		us.Logger.Errorf("list user failed, err: %v", err)
		return nil, err
	}
	return users, nil
}

func (us *UserService) ListUserByPageWithTotal(page, pageSize int) ([]*models.User, int64, error) {
	users, total, err := us.Model.FindPageListByPageWithTotal(page, pageSize)
	if err != nil {
		us.Logger.Errorf("list user failed, err: %v", err)
		return nil, 0, err
	}
	return users, total, nil
}
