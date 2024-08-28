package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Contact struct {
	Name  string
	Email string
}

func newContact(name string, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts = []Contact
type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("Soham", "sohamskudalkar@gmail.com"),
			newContact("John", "john.doe@mail.com"),
		},
	}
}

type Form struct {
	Values map[string]string
	Errors map[string]string
}

func newForm() Form {
	return Form{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form Form
}

func newPage() Page {
	return Page{
		Data: newData(),
		Form: newForm(),
	}
}

var page Page

func (p *Page) hasEmail(email string) bool {
	for _, contact := range p.Data.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func init() {
	page = newPage()
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("./views/index.html"))
	tmpl.ExecuteTemplate(w, "index", page)
}

func AddContact(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/index.html"))
	if page.hasEmail(r.FormValue("email")) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		formData := newForm()
		formData.Values["name"] = r.FormValue("name")
		formData.Values["email"] = r.FormValue("email")
		formData.Errors["email"] = r.FormValue("email")
		tmpl.ExecuteTemplate(w, "contactForm", formData)
		return
	}
	contact := newContact(r.FormValue("name"), r.FormValue("email"))
	page.Data.Contacts = append(page.Data.Contacts, contact)
	log.Println(page)
	tmpl.ExecuteTemplate(w, "oob-contact", contact)
	tmpl.ExecuteTemplate(w, "contactForm", newForm())
}
