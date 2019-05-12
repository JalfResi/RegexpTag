# RegexpTag

Given the following text file:

``` text
"FirstName: firstname
Surname: SURNAME
EmailAddress: firstname.surname@example.com
Organisation: Some County Council (Unitary)
Team: N/A
Username: firstname.surname
Status: Active
ShortCode: N/A
"
```

``` go
type User struct {
	State  string `regexp:"Status: (.+)"`
	User   struct {
		Email        string `regexp:"EmailAddress: (.+)"`
		FirstName    string `regexp:"FirstName: (.+)"`
		Organisation string `regexp:"Organisation: (.+)"`
		SecondName   string `regexp:"Surname: (.+)"`
		ShortCode    string `regexp:"ShortCode: (.+)"`
		Team         string `regexp:"Team: (.+)"`
		Username     string `regexp:"Username: (.+)"`
	}
}

u := &User{}
Parse(u, s)

assert.Equal(t, "Active", u.State)
assert.Equal(t, "firstname.surname@example.com", u.User.Email)
assert.Equal(t, "firstname", u.User.FirstName)
assert.Equal(t, "Some County Council (Unitary)", u.User.Organisation)
assert.Equal(t, "SURNAME", u.User.SecondName)
assert.Equal(t, "N/A", u.User.ShortCode)
assert.Equal(t, "N/A", u.User.Team)
assert.Equal(t, "firstname.surname", u.User.Username)
```

