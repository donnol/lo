package lo

import (
	"reflect"
	"testing"
)

type User struct {
	Id   uint64
	Name string
}

type Book struct {
	Id     uint64
	Title  string
	Author uint64 // User.Id
}

type BookWithUser struct {
	Book
	UserName string
}

func UserBookMatcher(j User, k Book) bool {
	return j.Id == k.Author
}

func TestJoin(t *testing.T) {
	r := Join([]User{
		{Id: 1, Name: "jd"},
		{Id: 2, Name: "jc"},
	}, []Book{
		{Id: 1, Title: "hello", Author: 1},
		{Id: 2, Title: "world", Author: 1},
		{Id: 3, Title: "good", Author: 2},
		{Id: 4, Title: "job", Author: 2},
	}, UserBookMatcher, func(j User, k Book) BookWithUser {
		return BookWithUser{
			Book:     k,
			UserName: j.Name,
		}
	})
	want := []BookWithUser{
		{Book{1, "hello", 1}, "jd"},
		{Book{2, "world", 1}, "jd"},
		{Book{3, "good", 2}, "jc"},
		{Book{4, "job", 2}, "jc"},
	}
	if !reflect.DeepEqual(r, want) {
		t.Errorf("bad case, %+v != %+v", r, want)
	}
}
