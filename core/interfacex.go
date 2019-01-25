package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokia NokiaPhone) call() {
	fmt.Print("my name is Nokia Phone, you can call me!\n")
}

type Ipone struct {
}

func (ipone Ipone) call() {
	fmt.Print("my name is  IPhone, you can call me!\n")
}

type Books struct {
	title   string
	subject string
	author  string
	price   string
	book_id int
}

func show_books() {

	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "golang", price: "10 USD", book_id: 1})

	var book Books
	var book2 Books

	book.title = "my go"
	book.subject = "my go  2"
	book.author = "golang"
	book.book_id = 2

	book2.title = "you go"
	book2.subject = "you go  2"
	book2.author = "golang"
	book2.book_id = 3

	printBook(book)

	printBook_pointer(&book2)

}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBook_pointer(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(Ipone)
	phone.call()

	show_books()

}
