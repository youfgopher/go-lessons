package main

import (
	"fmt"
	"log"
	"os"
)

/*
File manipulation
	- 1. Create file
	- 2. Open file
	- 3. Write
	- 4. Read file
	- 5. Close
*/

func main() {
	// 1.
	// Создание фала, метод Create принимает название фала которое ты хочешь создать
	// и возвращает (дескриптор) структуру только что созданного файла, и ошибку если что-то идет не так

	//file, err := os.Create("some_file.txt") // файл создастся в корне репозитория - проекта
	//if err != nil {
	//	fmt.Println(err)
	//}

	//fmt.Println(file) // output: &{0xc00005e180} - this is address of the file in RAM
	// после того как ты увидел адрес созданного фала в оперативке, идем дальше

	// 2. Open file
	// Открытые файла, метод Open принимает название фала который нужно открыть, и возвращает его *File
	// и ошибку если не удается его открыть
	file, err := os.Open("some_file.txt")
	if err != nil {
		fmt.Println(err)
	}
	// тут можно было сразу отложено закрыть файл с помощью defer, напомни я расскажу что и как
	//defer file.Close()

	// Важно запомни, любой файл который ты открыл, создал, нужно ЗАКРЫТЬ, иначе они так
	// и будут стоять открытые в оперативное помяты, а если этот файл весить 3 гигабайта ? твоя оперативка как комп
	// скажут спасибо и зависнуть, утечка будет, если выцдя из программы не закроешь, запомни, это важно

	// 3. Writing in file
	// Запись в файл, метод write принимает слайс байтов, поэтому нам нужжно привести данные к байту
	// метод возвращает количество записанных байт, и ошибку

	// first way
	//_, err = file.Write([]byte("some text for writing in file\n"))
	//if err != nil {
	//	fmt.Println(err)
	//}

	// second way
	//err = os.WriteFile("some_file.txt", []byte("some text for writing in file\n"), 0666)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 4. Read file
	// Чтение из фала, метод принимает наз. файла для чтения и возвращает данные в виде слайса байтов
	// собственно нам нужно привести их к строке
	rbytes, err := os.ReadFile("some_file.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(rbytes))

	// 5. Closing the file
	// у структуры файла вызываем метод Close, закрывает файл, и возвращает ошибку если не удалось
	// можешь обработать, можешь не делать
	file.Close()
}
