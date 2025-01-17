package models

import (
	"log"
	"shvx/config"
)

type UserSignUPInfo struct {
	Name        string
	Email       string
	Password    string
	PhoneNumber string
}

func UserSignUP(userInfo UserSignUPInfo) error {
	db, err := config.GetDB2()
	if err != nil {
		log.Println("Failed to  connnect with database with error :", err)
		return err
	}
	defer db.Close()
	query := `
				insert into shvx_user
				(user_name,phone_number,email,password
				)values(
					$1,$2,$3,$4)`
	log.Print(query, userInfo.Name, userInfo.PhoneNumber, userInfo.Email, userInfo.Password)
	_, err = db.Exec(query, userInfo.Name, userInfo.PhoneNumber, userInfo.Email, userInfo.Password)
	if err != nil {
		log.Println("Failed to  execute the query into database with error :", err)
		return err
	}
	return nil
}
func GetUserPassword(email string) (string, string, error) {
	var password, name string
	db, err := config.GetDB2()
	if err != nil {
		log.Println("Failed to  connnect with database with error :", err)
		return password, name, err
	}
	defer db.Close()
	query := `select password,user_name from shvx_user where email = $1`
	err = db.QueryRow(query, email).Scan(&password, &name)
	if err != nil {
		log.Println("Failed to  execute the query into database with error :", err)
		return password, name, err
	}
	return password, name, nil
}
