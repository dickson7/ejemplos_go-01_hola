"# proyectogocomentarios"
Este proyecto abarca el Backend para el manejo de comentarios. Es una api  donde se valida el ingreso de  los usuarios los comentarios realizados y los votados tanto positivos como negativos.


#Crear usuarios
#POST
localhost:8080/api/users/

{
"username":"dickson",
"email":"garciadickson258@gmail.com",
"fullname":"Dickson Garcia",
"password": "admin1234",
"confirmPassword":"admin1234"
}

#Hacer Login
#POST
localhost:8080/api/login

{
"email":"garciadickson258@gmail.com",
"password":"admin1234"
}
