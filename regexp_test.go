package regexptag

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type User struct {
	State  string `regexp:"Status: (.+)"   conform:"lower"   yaml:"state"`
	User   struct {
		Email        string `regexp:"EmailAddress: (.+)" conform:"email"                yaml:"email"`
		FirstName    string `regexp:"FirstName: (.+)"    conform:"trim,name"            yaml:"first_name"`
		Organisation string `regexp:"Organisation: (.+)" conform:"lower,snake"          yaml:"organisation"`
		SecondName   string `regexp:"Surname: (.+)"      conform:"trim,name"            yaml:"second_name"`
		ShortCode    string `regexp:"ShortCode: (.+)"    conform:"trim,shortcode,snake" yaml:"shortcode"`
		Team         string `regexp:"Team: (.+)"         conform:"trim,team,snake"      yaml:"team"`
		Username     string `regexp:"Username: (.+)"     conform:"trim,lower"           yaml:"username"`
	}
}

func TestParse(t *testing.T) {
	// Arrange
	buf, err := ioutil.ReadFile("testdata/example.source.golden")
	if err != nil {
		t.Error(err)
	}
	s := string(buf)

	// Act
	u := &User{}
	Parse(u, s)

	// Assert
	assert.Equal(t, "Active", u.State)
	assert.Equal(t, "firstname.surname@example.com", u.User.Email)
	assert.Equal(t, "firstname", u.User.FirstName)
	assert.Equal(t, "Some County Council (Unitary)", u.User.Organisation)
	assert.Equal(t, "SURNAME", u.User.SecondName)
	assert.Equal(t, "N/A", u.User.ShortCode)
	assert.Equal(t, "N/A", u.User.Team)
	assert.Equal(t, "firstname.surname", u.User.Username)
}
