//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type BookTitle string
type MemberName string

type lendAudit struct {
	checkOutTime time.Time
	checkInTime  time.Time
}

type Member struct {
	name  MemberName
	books map[BookTitle]lendAudit
}

type BookEntry struct {
	total  int
	lended int
}
type library struct {
	books   map[BookTitle]BookEntry
	members map[MemberName]Member
}

func printMemberAudit(m *Member) {
	for book, audit := range m.books {
		var returnTime string
		if audit.checkInTime.IsZero() {
			returnTime = "not returned"
		} else {
			returnTime = audit.checkInTime.String()
		}
		fmt.Println(m.name, "checked out", book, "at", audit.checkOutTime, "and returned at", returnTime)
	}
}
func printAudit(library *library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}

}
func printLibraryBooks(library *library) {
	for book, entry := range library.books {
		fmt.Println(book, "/ total", entry.total, "/ lended", entry.lended)
	}
	fmt.Println()
}
func checkoutBook(library *library, book BookTitle, member Member) {
	entry, found := library.books[book]
	if !found {
		panic("book not found")
	}
	if entry.lended >= entry.total {
		panic("no more copies")
	}
	entry.lended += 1
	library.books[book] = entry
	member.books[book] = lendAudit{checkOutTime: time.Now()}

}
func returnBook(library *library, book BookTitle, member Member) {
	entry, found := library.books[book]
	if !found {
		panic("book not found")
	}
	if entry.lended <= 0 {
		panic("no more copies")
	}
	entry.lended -= 1
	library.books[book] = entry
	audit, found := member.books[book]
	if !found {
		panic("book not found")
	}
	audit.checkInTime = time.Now()
	member.books[book] = audit
}
func main() {
	library := library{
		books:   make(map[BookTitle]BookEntry),
		members: make(map[MemberName]Member),
	}
	library.books["The Hobbit"] = BookEntry{total: 1, lended: 0}
	library.books["The Lord of the Rings"] = BookEntry{total: 3, lended: 0}
	library.books["The Silmarillion"] = BookEntry{total: 2, lended: 0}
	library.books["The Children of Hurin"] = BookEntry{total: 1, lended: 0}

	library.members["John"] = Member{name: "John", books: make(map[BookTitle]lendAudit)}
	library.members["Jane"] = Member{name: "Jane", books: make(map[BookTitle]lendAudit)}
	library.members["Jack"] = Member{name: "Jack", books: make(map[BookTitle]lendAudit)}

	printLibraryBooks(&library)

	checkoutBook(&library, "The Hobbit", library.members["John"])
	checkoutBook(&library, "The Lord of the Rings", library.members["John"])
	checkoutBook(&library, "The Silmarillion", library.members["Jane"])
	checkoutBook(&library, "The Children of Hurin", library.members["Jack"])

	printAudit(&library)

	returnBook(&library, "The Hobbit", library.members["John"])
	returnBook(&library, "The Lord of the Rings", library.members["John"])
	returnBook(&library, "The Silmarillion", library.members["Jane"])
	returnBook(&library, "The Children of Hurin", library.members["Jack"])

	printAudit(&library)

	printLibraryBooks(&library)

}
