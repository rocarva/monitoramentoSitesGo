// principal pacote na aplicacao
// para rodar go run main.go
// para buildar o condigo go build main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()
	leSitesDoArquivo()

	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs")
			imprimeLogs()
		case 0:
			fmt.Println("Sair do programa")
			os.Exit(0)
		default:
			fmt.Println("Não reconheco o programa")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Rodrigo"
	versao := 1.1

	fmt.Println("Olá,sr", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	// lista de opções
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	// capturar o que o usuario escreveu
	var comandoLido int
	fmt.Scan(&comandoLido)

	fmt.Println("O valor da variável comando é:", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {

	fmt.Println(" Iniciar Monitoramento")
	// sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			fmt.Println("Ocorreu um erro", err)
			break
		}
	}

	arquivo.Close()

	return sites

}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(arquivo)

	arquivo.WriteString(time.Now().Format("02/01/2006") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(string(arquivo))
}
