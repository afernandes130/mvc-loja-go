package controllers

import (
	"log"
	"loja/models"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := models.Produto{
			Nome:      r.FormValue("nome"),
			Descricao: r.FormValue("descricao"),
		}

		// nome := r.FormValue("nome")
		// descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidParFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Err na conversão do preço", err)
		}
		produto.Preco = precoConvertidParFloat

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Err na conversão da quantidade", err)
		}
		produto.Quantidade = quantidadeConvertidaParaInt

		models.CriaNovoProduto(produto)
		http.Redirect(w, r, "/", 301)

	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.BuscaProduto(idProduto)
	templates.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := models.Produto{
			Nome:      r.FormValue("nome"),
			Descricao: r.FormValue("descricao"),
		}

		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		id := r.FormValue("id")

		precoConvertidParFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Err na conversão do preço", err)
		}
		produto.Preco = precoConvertidParFloat

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Err na conversão da quantidade", err)
		}
		produto.Quantidade = quantidadeConvertidaParaInt

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Err na conversão do ID", err)
		}
		produto.Id = idConvertidaParaInt
		models.AtualizaProduto(produto)
		http.Redirect(w, r, "/", 301)
	}

}
