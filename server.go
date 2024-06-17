package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os/exec"
)

type Command struct {
    Cmd string `json:"cmd"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    // Verifica se o método da requisição é POST
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    // Decodifica o corpo da requisição JSON
    var cmd Command
    if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
        http.Error(w, fmt.Sprintf("Erro ao decodificar JSON: %v", err), http.StatusBadRequest)
        return
    }

    // Verifica se o comando foi especificado
    if cmd.Cmd == "" {
        http.Error(w, "Comando não especificado", http.StatusBadRequest)
        return
    }

    // Executa o comando no terminal
    out, err := exec.Command("bash", "-c", cmd.Cmd).CombinedOutput()
    if err != nil {
        log.Printf("Erro ao executar comando: %v", err)
        http.Error(w, fmt.Sprintf("Erro ao executar comando: %v", err), http.StatusInternalServerError)
        return
    }

    // Escreve a saída do comando como resposta HTTP
    fmt.Fprintf(w, "Comando executado com sucesso:\n%s", out)
}

func main() {
    // Configura o handler para a rota /executar-comando
    http.HandleFunc("/executar-comando", handler)

    // Inicia o servidor na porta 8080
    log.Fatal(http.ListenAndServe(":8080", nil))
}
