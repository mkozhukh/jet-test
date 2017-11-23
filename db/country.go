package db

import (
	"net/http"

	"github.com/mholt/binding"
)

// Country is a DB record of a single snippet Country
type Country struct {
	CommonModel
	Name string
	Code string
}

// FieldMap provides a field mapping
func (cf *Country) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&cf.Name: "Name",
		&cf.Code: "Code",
	}
}

// GetCountry returns Country by id
func (db *Context) GetCountry(id string) *Country {
	file := Country{}
	db.First(&file, id)
	return &file
}

// GetAllCountrys returns all snippets
func (db *Context) GetAllCountrys() []Country {
	files := []Country{}
	db.Find(&files)
	return files
}

// DeleteCountry deletes snippet from DB
func (db *Context) DeleteCountry(id string) {
	file := db.GetCountry(id)
	if file.ID != 0 {
		db.Delete(file)
	}
}

// UpdateCountry deletes Country from DB
func (db *Context) UpdateCountry(id string, update *Country) *Country {
	file := db.GetCountry(id)
	if file.ID != 0 {
		db.Model(file).Update(update)
	}
	return file
}

// AddCountry deletes Country from DB
func (db *Context) AddCountry(update *Country) *Country {
	db.Save(&update)
	return update
}
