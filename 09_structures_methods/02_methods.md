# Methods - Методы

В продолжении про моделирование, в нашем реальном мире каждый объект (структура, объект, сущность) помимо свойств может
совершать какие-то действия (функции - методы), например машина может ездить, ускорятся, тормозить, сигналить, у машины
можно открыть двери, капот, багажник и тд. Собственно метод — это функция, связанная с определенной структурой, типом.

 То есть различия обычной функции от метода в том что функция не принадлежит никакой структуре, типу, его можно вызвать и
использовать в любом участке кода без привязки и использование через структуру, а метод мы не можем вызывать без обращения
и вызова непосредственно через структуру, так как метод это функция - действие которое принадлежит (свойственно) определенной
структуре, и вот эта связь структуры и его метода определяется (создается связь, принадлежность) через receiver, запомни
это слово, и определение в целом, это важно. Сейчас увидишь наглядно.

```go
// decl
type Car struct {
    name string
    color string
    engine float64
}

// init car 
car := Car{
    name: "mercedes",
    color "red",
    engine: 5.0,
}

// decl car's method 
func (c Car) run() { // (c Car) <- receiver, это и есть то что связывает, определяет принадлежность
    // some action ...
    fmt.Printf("Car, name: %s color: %s, engine: %f running...", c.name, c.color, c.engine)
}

// call car's method, как видишь вызываем как обычную функцию, но только через точку (dot, point - точка - связь) со
// структурой, которое и подчеркивает связь
car.run() // output: Car, name: mercedes color: red, engine 5.0 running...

```
Не будем повторятся, внимательно прочти комментарии к коду, прокомментировал про receiver, и связь - принадлежность.

```go
func (r Receiver - StructName) MethodName(params...) (returning params...) {
    // instructs, some action
}
```

Есть правило насчет наименования ресивера, ресивер это имя структуры, а первая буква это аляс (и принято первую букву имени структуры
использовать для аляса, как выше в примере показано) для удобного обращения к структуре как показано в первом примере.

### Methods and pointer as method receiver

***На примере выше показан пример когда метод получает и работает с копией объекта т.е если мы как то изменим поля структуры
внутри объекта, то это ни как не изменить значение поля самой структуры, т.е изменен будет только копия структуры внутри метода.***

Если мы хотим чтобы изменения которые мы делаем со структурой внутри метода изменили значения полей самой структуры то
нам нужно указать структуру в качестве ресивера ссылку на структуру, сейчас увидишь, продолжение примера с машиной:

```go
car := Car{
    name: "mercedes",
    color "red",
    engine: 5.0,
}

// working with copy of struct 
func (c Car) Upgrade() {
    // some action 
    c.color = "white" // изминения не затронуть поля самой структуры, а будуть касаться только копии внутри метода
    c.engine = 9.0
}

fmt.Printf("Car, name: %s color: %s, engine: %f running...", c.name, c.color, c.engine)
// output: Car, name: mercedes color: red, engine 5.0 running...


// working with pointer of struct, direct manipulation with struct, pointer as method receiver
func (c *Car) Upgrade() {
    // some action
    c.color = "white" // изминения затронуть поля самой структуры, так как мы работаем напрямую со структурой, ссылка как помнишь
    c.engine = 9.0 // расказывал про тип  * - pointer, вот пример прямой работы со значениямы в оперативной памяты
}


fmt.Printf("Car, name: %s color: %s, engine: %f running...", c.name, c.color, c.engine)
// output: Car, name: mercedes color: white, engine 9.0 running...

```