package http

import (
	"flag"
	"log"
	"net/http"
	"route256/cart/internal/pkg/clients/loms"
	"route256/cart/internal/pkg/clients/product"
	hitem "route256/cart/internal/pkg/handlers/item"
	sitem "route256/cart/internal/pkg/services/item"
)

type config struct {
	addr        string
	lomsAddr    string
	productAddr string
}

func newConfigFromFlags() config {
	const (
		defaultAddr        = ":8080"
		defaultLomsAddr    = "http://loms:8080"
		defaultProductAddr = "http://route256.pavl.uk:8080"
	)

	result := config{}
	flag.StringVar(&result.addr, "addr", defaultAddr, "server adress default: "+defaultAddr)
	flag.StringVar(&result.lomsAddr, "loms_addr", defaultLomsAddr, "server adress default: "+defaultLomsAddr)
	flag.StringVar(&result.productAddr, "product_addr", defaultProductAddr, "server adress default: "+defaultProductAddr)
	flag.Parse()
	return result
}

type App struct {
	config config
}

func NewApp() *App {
	return &App{
		config: newConfigFromFlags(),
	}
}

func (app *App) Start() error {
	lomsClient, err := loms.New("loms client", app.config.lomsAddr)
	if err != nil {
		log.Fatal("Error creating loms client")
	}

	productClient, err := product.New("product client", app.config.productAddr)
	if err != nil {
		log.Fatal("Error creating product client")
	}
	itemAddHandler := hitem.NewItemsAddHandler(sitem.NewAddService(lomsClient, productClient))
	http.HandleFunc("/item/add", itemAddHandler.Handle)

	itemDeleteHAndler := hitem.NewItemsDeleteHandler(sitem.NewDeleteService())
	http.HandleFunc("/item/delete", itemDeleteHAndler.Handle)

	log.Fatal(http.ListenAndServe(app.config.addr, nil))
	return nil
}
