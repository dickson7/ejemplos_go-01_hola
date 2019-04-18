"# proyectogocomentarios"
Este proyecto abarca el Backend para el manejo de comentarios. Es una api  donde se valida el ingreso de  los usuarios los comentarios realizados y los votados tanto positivos como negativos.

#Instalación
para ejecutar la aplicacion y ejecutar la migracion
./proyectogocomentarios --migrate yes
proyectogocomentarios.exe --migrate yes
de esta manera se generan las tablas en la BD
si no ejecuta migrate el valor seria no y no se llama a migrate

#Ejecución
./proyectogocomentarios
proyectogocomentarios.exe


#Crear usuarios
#POST

localhost:8080/api/users/

#json
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

#json
{
"email":"garciadickson258@gmail.com",
"password":"admin1234"
}

#Registrar comentarios
Se realiza el Login
#POST
ruta_http://localhost:8080/api/comments/

Headers
Content-type: application/json
Authorization: Token del usuario

#json
{
	"userId": 1,
	"parentId": 0, // es cero porque no es un comentario hijo
	"votes": 0,  // aun no tiene votos
	"content": "Nuestro primer comentario"
}

#Comentario sin utilizar usuario ya que lo extrae del token 
{
	"content": "otror comentario sin ID de usuario"
}

#consulta de proyectogocomentarios
#GET
localhost:8080/api/comments/

Debe contener el token
