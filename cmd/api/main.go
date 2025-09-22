package main

import "myapp/internal/fx"

func main() {

	fx.New().Run()
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
