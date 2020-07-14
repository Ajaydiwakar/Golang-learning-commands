package main
import "fmt"

type Books struct{
	title string
	id int
}
func main (){
	var book1 Books
	var book2 Books

	book1.title = "Philosophy"
	book1.id = 12

	book2.title = "Diwakar empire"
	book2.id = 444


	printBook(book1)
	printBook(book2)

}
func printBook(Book Books){
	fmt.Printf("Book1 tile  %s\n", Book.title)
	fmt.Printf("book1 id  %d\n" , Book.id)
}