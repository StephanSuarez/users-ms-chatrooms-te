package http

import (
	"log"
	"os"
)

func SubcribersHandlers(UsersDep *UsersDependencies) {
	// Canal para manejar errores
	errChan := make(chan error)

	// Ejecutar cada suscriptor en su propia goroutine
	go func() {
		errChan <- createUserSubs(os.Stdout, UsersDep)
	}()
	go func() {
		errChan <- deleteUserSubs(os.Stdout, UsersDep)
	}()

	// Manejar errores de las goroutines
	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			log.Fatal(err)
		}
	}
}
