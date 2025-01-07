package main

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type payload struct {
	Input string `json:"input"`
}

var jokenpoInputs = []string{"pedra", "tesoura", "papel"}
var inputValid = map[string]struct{}{
	"pedra":   {},
	"tesoura": {},
	"papel":   {},
}

var resultJokenpo = map[string]string{
	"pedra_pedra":     "empate",
	"pedra_tesoura":   "Você venceu",
	"pedra_papel":     "Você perdeu",
	"tesoura_pedra":   "Você perdeu",
	"tesoura_tesoura": "empate",
	"tesoura_papel":   "Você venceu",
	"papel_pedra":     "Você venceu",
	"papel_tesoura":   "Você perdeu",
	"papel_papel":     "empate",
}

func main() {
	e := echo.New()

	e.POST("/", jokenpo)

	if err := e.Start(":3002"); err != nil {
		panic("Erro ao iniciar servidor: " + err.Error())
	}
}

func jokenpo(c echo.Context) error {
	body := payload{}
	if err := c.Bind(&body); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "payload invalido"}
	}

	body.Input = strings.ToLower(body.Input)
	if _, ok := inputValid[body.Input]; !ok {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "input diferente de Pedra, Papel, tesoura"}
	}

	rOption := rand.Intn(2)

	result, ok := resultJokenpo[body.Input+"_"+jokenpoInputs[rOption]]
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Erro ao processar resultado"}
	}

	return c.JSON(http.StatusOK, map[string]any{
		"resultado":        result,
		"opção da maquina": jokenpoInputs[rOption],
	})
}
