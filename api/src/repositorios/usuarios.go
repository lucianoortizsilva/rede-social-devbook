package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	sql := "insert into usuarios (nome, nick, email, senha) values(?,?,?,?)"
	statement, erro := repositorio.db.Prepare(sql)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar retorna todos os usuários por filtro
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	sql := "select id, nome, nick, email, criadoEm from usuarios  where nome like ? or nick like ?"
	linhas, erro := repositorio.db.Query(sql, nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var u modelos.Usuario
		if erro = linhas.Scan(
			&u.ID,
			&u.Nome,
			&u.Nick,
			&u.Email,
			&u.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, u)
	}

	return usuarios, nil
}
