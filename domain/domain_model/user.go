package domain_model

type User struct {
    UserId UserId
    Email Email
    EncryptedPassword EncryptedPassword
}

func GenerateUser(userId UserId, email Email, encryptedPassword EncryptedPassword) User {
    return User{
        UserId: userId,
        Email: email,
        EncryptedPassword: encryptedPassword,
    }
}
