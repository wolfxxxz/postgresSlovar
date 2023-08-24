package users

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func StartUser(db *sqlx.DB, name string) (int, error) {
	//Создать пользователя
	//newUs := NewUser()
	//newUs.ScanUser()

	us, err := GetUserByName(db, name)
	if err != nil {
		fmt.Println(err)
		fmt.Println("You need to register in the system")
		us.ScanUser()
		err = InsertUser(db, us)
		if err != nil {
			return 0, err
		}
		id, err := StartUser(db, us.Name)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	fmt.Println(us)
	return us.Id, nil

}
