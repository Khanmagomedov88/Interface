package main

import "fmt"

type Speaker interface {
	Speak()
}

type Obj interface{}

type Dog struct {
	Name string
	Age  int
}

func (d Dog) Speak() {
	fmt.Println("Гав-гав... мое имя", d.Name)
}

func (c Cat) Speak() {
	fmt.Println("Мяу-мяу... мое имя", c.Name)
}

func (d Dog) String() string {
	return fmt.Sprintf("Меня зовут %s, мне %d", d.Name, d.Age)
}

func MakeSound(sp Speaker) {
	sp.Speak()
}

type Cat struct {
	Name string
}

func Write(o Obj) {
	fmt.Printf("Type: %T, Value: %v\n", o, o) // Использование пустого интерфейса
}

func main() {
	fmt.Println("Привет мир!")

	cat := Cat{Name: "Муизза"}
	dog := Dog{Name: "Барсик", Age: 10}

	// Используем интерфейс напрямую
	cat.Speak()
	dog.Speak()

	// Используем MakeSound для полиморфизма
	MakeSound(cat)
	MakeSound(dog)

	// Используем пустой интерфейс
	Write(cat)
	Write(dog)

}
