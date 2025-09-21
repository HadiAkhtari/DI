package main

import (
	"go.uber.org/fx"
	"time"
)

func main() {

	fx.New(
		fx.StartTimeout(time.Minute*2),
		fx.StopTimeout(time.Minute*2),
		fx.Provide(),
	).Run()
	/*tp := service.RealTime{}
	svc := service.NewGreeterService(tp)
	h := controller.NewHandler(svc)
	http.HandleFunc("/", h.Greet)
	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)*/

	/*handler := controller.NewHandler()
	http.HandleFunc("/", handler.Greet)
	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)*/
}
