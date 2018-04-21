package orm

import (
	"fmt"
)

// Framework initialize
func init() {
	fmt.Println("Gear init...")
}

func GetContext() OrmContext {
	return singleOrmContext()
}

// Create Orm
func GetDao() *Orm {
	return &Orm{}
}

