// package models
package archive

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body  string
}
