package users

import (
	"fmt"
	"postgresTakeWords/internal/repositories"

	"github.com/jmoiron/sqlx"
)

func StartUser(db *sqlx.DB, name string) (int, error) {
	us, err := repositories.GetUserByName(db, name)
	if err != nil {
		fmt.Println(err)
		fmt.Println("You need to register in the system")
		us.ScanUser()
		err = repositories.InsertUser(db, us)
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
