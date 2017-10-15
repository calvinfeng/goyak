// Model contains all the model for our Postgres database, just like ActiveRecord. We create a model and the snake case
// version of the struct name is going to be used as table name. Similarly, the snake cased field names will be used
// column names for the table.
package model

import "github.com/jinzhu/gorm"

// Default table name is users
type User struct {
	gorm.Model
	Name      string
	Email     string     `gorm:"type:varchar(100);unique_index"` // Set SQL type and create unique indexing
	ChatRooms []ChatRoom `gorm:"many2many:memberships"`       // many-to-many
	Messages  []Message  `gorm:"ForeignKey:UserID"`              // has-many
}