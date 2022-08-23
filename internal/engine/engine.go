package engine

import (
	"fmt"
	"net/http"

	"github.com/alwindoss/fragmos"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run(cfg *fragmos.Config) error {
	fragmos.MINING_DIFFICULTY = cfg.MiningDifficulty
	fmt.Println("Mining Difficult set to: ", cfg.MiningDifficulty)
	myBlockchainAddress := "my_fragmos_address"
	bc := NewBlockchain(myBlockchainAddress)
	// bc.Print()
	bc.AddTransaction("PersonA", "PersonB", 1.0)
	bc.Mining()
	// bc.Print()

	bc.AddTransaction("PersonC", "PersonD", 2.0)
	bc.AddTransaction("PersonX", "PersonY", 3.0)
	bc.Mining()
	bc.Print()

	fmt.Printf("PersonA Total: %.1f\n", bc.CalculateTotalAmount("PersonA"))
	fmt.Printf("PersonC Total: %.1f\n", bc.CalculateTotalAmount("PersonC"))
	fmt.Printf("PersonB Total: %.1f\n", bc.CalculateTotalAmount("PersonB"))
	fmt.Printf("PersonD Total: %.1f\n", bc.CalculateTotalAmount("PersonD"))
	fmt.Printf("PersonX Total: %.1f\n", bc.CalculateTotalAmount("PersonX"))
	fmt.Printf("PersonY Total: %.1f\n", bc.CalculateTotalAmount("PersonY"))
	fmt.Printf("my_fragmos_address Total: %.1f\n", bc.CalculateTotalAmount("my_fragmos_address"))
	return nil
}

func RunServer(cfg *fragmos.Config) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	addr := fmt.Sprintf(":%d", cfg.Port)
	http.ListenAndServe(addr, r)
	return nil
}
