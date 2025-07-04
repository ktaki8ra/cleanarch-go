package gateway

import (
    "fmt"
    "gorm.io/gorm"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_model"
    "github.com/ktaki8ra/cleanarch-go/interfaces/config/db_model"
)

type UserRepositoryImpl struct {
    DB *gorm.DB
}

func (ur *UserRepositoryImpl) FindUserById(
    userId domain_model.UserId,
) (domain_model.User, error) {
    var userRow struct {
        UserId string
        Email string
        EncryptedPassword string
    }
    result := ur.DB.Table("users").Where("user_id = ?", userId.Value).First(&userRow)
    if result.Error != nil {
        return domain_model.User{}, result.Error
    }
    return domain_model.User{
        UserId:            domain_model.UserId{Value: userRow.UserId},
        Email:             domain_model.Email{Value: userRow.Email},
        EncryptedPassword: domain_model.EncryptedPassword{Value: userRow.EncryptedPassword},
    }, nil
}

func (ur *UserRepositoryImpl) Save(user domain_model.User) error {
    newUser := db_model.Users{
        UserId:            user.UserId.Value,
        Email:             user.Email.Value,
        EncryptedPassword: user.EncryptedPassword.Value,
    }
    res := ur.DB.Create(&newUser)
    return res.Error
}

func (ur *UserRepositoryImpl) Delete(user domain_model.User) error {
    result := ur.DB.Unscoped().Where("user_id = ?", user.UserId.Value).Delete(&db_model.Users{})
    return result.Error
}

func (ur *UserRepositoryImpl) Update(currentUserId domain_model.UserId, newUser domain_model.User) error {
    updateUser := db_model.Users{
        UserId:            newUser.UserId.Value,
        Email:             newUser.Email.Value,
        EncryptedPassword: newUser.EncryptedPassword.Value,
    }
    res := ur.DB.Model(&db_model.Users{}).Where("user_id = ?", currentUserId.Value).Updates(updateUser)
    if res.Error != nil {
        return res.Error
    }
    if res.RowsAffected == 0 {
        return fmt.Errorf("user not found: user_id = %s", currentUserId.Value)
    }
    return nil
}

