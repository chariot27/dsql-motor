package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Estrutura simplificada apenas com o essencial
type DSQLLite struct {
	Session *discordgo.Session
	GuildID string
}

func main() {
	// 1. Configuração Rápida (Pode usar variáveis de ambiente ou input)
	token := os.Getenv("DSQL_TOKEN")
	guildID := os.Getenv("DSQL_GUILD")

	if token == "" || guildID == "" {
		fmt.Println("❌ Erro: Defina DSQL_TOKEN e DSQL_GUILD no ambiente.")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Erro ao criar sessão:", err)
		return
	}

	db := &DSQLLite{Session: dg, GuildID: guildID}
	fmt.Println("✅ DSQL Lite Conectado! (Digite 'sair' para encerrar)")
	fmt.Println("--------------------------------------------------")

	// 2. Loop de execução (REPL)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("DSQL > ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if strings.ToLower(input) == "sair" {
			break
		}
		if input == "" {
			continue
		}

		// Lógica de execução simplificada
		res, err := db.quickExecute(input)
		if err != nil {
			fmt.Printf("⚠️ Erro: %v\n", err)
		} else {
			fmt.Printf("📝 Resultado: %s\n", res)
		}
	}
}

// Motor de execução "magro"
func (db *DSQLLite) quickExecute(query string) (string, error) {
	upper := strings.ToUpper(query)

	// Exemplo simplificado de INSERT
	if strings.HasPrefix(upper, "DINSERT") {
		// Aqui você manteria seu Regex original do DINSERT
		return "Comando INSERT enviado (Simulação Lite)", nil
	}

	// Exemplo simplificado de SELECT (Apenas lista as IDs das mensagens)
	if strings.HasPrefix(upper, "DSELECT") {
		parts := strings.Fields(query)
		if len(parts) < 2 {
			return "", fmt.Errorf("Use: DSELECT tabela")
		}

		chID := db.findChannel(parts[len(parts)-1])
		msgs, _ := db.Session.ChannelMessages(chID, 10, "", "", "")

		fmt.Println("| ID MENSAGEM          | CONTEÚDO RESUMIDO")
		for _, m := range msgs {
			fmt.Printf("| %s | %s...\n", m.ID, strings.ReplaceAll(m.Content, "\n", " ")[:20])
		}
		return "Fim da consulta", nil
	}

	return "Comando não implementado nesta versão Lite", nil
}

func (db *DSQLLite) findChannel(name string) string {
	channels, _ := db.Session.GuildChannels(db.GuildID)
	for _, c := range channels {
		if strings.EqualFold(c.Name, name) {
			return c.ID
		}
	}
	return ""
}
