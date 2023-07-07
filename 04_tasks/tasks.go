package tasks


/*

Написать функцию которая принимает -> имя, и телефон юзера -> валидирует входные данные - 
(проверяя на наличие пустоты, не корректности длины имены должна быть больше 4 букв, номер 
телефона больше 6 цифр, формат номера телефона) -> после убирает лишние символы, и приводит 
в одну форму !!! -> и сохраняет в мапу соответсвенно -> имя:номер-телефона-юзера**


main structure

#A function that accepts -> username and phone number
main function(...params) { 
	
	# Validates input data (checking for the presence of emptiness, incorrect 
	# length of the name must be more than 4 letters, number
	# phone number more than 6 digits, phone number format)
	phone, name := validator(params)
	
	#then removes extra characters, and leads in one form!
	phone, name := filter(params)
	
	# saves to the map accordingly -> name:user-phone-number
	saver(params)
} 


# Validates input data (checking for the presence of emptiness, incorrect 
# length of the name must be more than 4 letters, number
# phone number more than 6 digits, phone number format)
validator(params) params { }


#then removes extra characters, and leads in one form!
filter(params) {}


# saves to the map accordingly -> name:user-phone-number
saver(params) { }

var phoneNumbers map[string]string -> userName:phoneNumber


*/

// let's gщ