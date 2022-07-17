package models

import (
	"loja/db"
)

type Produto struct {
	Nome, Descricao string
	Id, Quantidade  int
	Preco           float64
}

func BuscaTodosProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	selectTodoProtudos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectTodoProtudos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodoProtudos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Quantidade = quantidade
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriaNovoProduto(produto Produto) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)
	defer db.Close()
}

func DeletaProduto(idProduto int) {
	db := db.ConectaComBancoDeDados()
	deletarProduto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())

	}
	deletarProduto.Exec(idProduto)
	defer db.Close()

}

func BuscaProduto(idProduto string) Produto {
	db := db.ConectaComBancoDeDados()
	selectTodoProtudos, err := db.Query("select * from produtos where id = $1", idProduto)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for selectTodoProtudos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodoProtudos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Quantidade = quantidade
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
	}
	defer db.Close()
	return produto

}

func AtualizaProduto(produto Produto) {
	db := db.ConectaComBancoDeDados()
	atualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	atualizaProduto.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.Id)
	defer db.Close()
}
