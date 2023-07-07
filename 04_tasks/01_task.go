package tasks


/*

				ВАЖНО СНАЧАЛА ВНИМАТЕЛЬНО ПРОЧТИ ОПИСАНИЕ !!!

Написать функцию которая принимает -> имя, и телефон юзера -> валидирует входные данные - 
(проверяя на наличие пустоты, не корректности длины имены должна быть больше 4 букв, номер 
телефона больше 6 цифр, формат номера телефона), если не валидна, возвр. ошибку -> 
после убирает лишние символы, и приводит в одну форму, если формат номера не корректен то ошибка !!! 
-> и сохраняет в мапу соответсвенно, и имя и номер телефона должны быть уникальны, 
иначе возврашается ошибка !!! -> юзернейм:номер-телефона-юзера**


main structure

#A function that accepts -> username and phone number
function(username, phone) error { 
	
	# Validates input data (checking for the presence of emptiness, incorrect 
	# length of the name must be more than 4 letters, number
	# phone number more than 6 digits, phone number format)
	phone, name := validator(params)
	# error check
	
	#then removes extra characters, and leads in one form!
	phone, name := filter(params)
	# error check

	# saves to the map accordingly -> name:user-phone-number
	saver(params)
	# error check
} 


# Validates input data (checking for the presence of emptiness, incorrect 
# length of the name must be more than 4 letters, number
# phone number more than 6 digits, phone number format)
validator(params) (param, param, error) { }


#then removes extra characters, and leads in one form!
filter(params) (param, param, error)  {}


# saves to the map accordingly -> name:user-phone-number
saver(params) error { }

var phoneNumbers map[string]string -> userName:phoneNumber


*/

// let's gщ