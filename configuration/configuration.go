package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Configuration va a tener la misma estructura que el archivo json que es la confi para acceder a la BD
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// GetConfiguration funcion para leer el archivo json y mapearla con la struct configuration
func GetConfiguration() Configuration {
	var c Configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//poblamos la info del archivo json en la variable c de tipo Configuration
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

//GetConnection voy a obtener la conexion a la base de datos
func GetConnection() *gorm.DB {
	//primero obtenemos la configuracion del archivo json que tiene la info de la BD
	c := GetConfiguration()
	//user:password@tcp(server: port)/database?charset=utf8&parseTime=true&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	//abrimos esa conexión
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
