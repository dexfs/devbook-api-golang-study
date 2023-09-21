package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into publicacoes (titulo, conteudo, author_id) values (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AuthorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		select p.*, u.nick from publicacoes p
		inner join usuarios u on u.id = p.author_id 
		where p.id = ?
	`, publicacaoID)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	var publicacao modelos.Publicacao
	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AuthorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AuthorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}
	return publicacao, nil
}

func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	var query string = `select DISTINCT p.*, u.nick from publicacoes p
		inner join usuarios u on u.id = p.author_id
		inner join seguidores s on p.author_id = s.usuario_id 
		where u.id = ? or s.seguidor_id = ?
		order by 1 desc`

	linhas, erro := repositorio.db.Query(query, usuarioID, usuarioID)
	if erro != nil {
		return []modelos.Publicacao{}, erro
	}
	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AuthorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AuthorNick,
		); erro != nil {
			return []modelos.Publicacao{}, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	var query = "update publicacoes set titulo = ?, conteudo = ? where id = ?"
	statement, erro := repositorio.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	var query = "delete from publicacoes where id = ?"
	statement, erro := repositorio.db.Prepare(query)
	if erro != nil {
		return erro
	}

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	var query = "select p.*, u.nick from publicacoes p inner join usuarios u on u.id = p.author_id where p.author_id = ?"
	linhas, erro := repositorio.db.Query(query, usuarioID)
	if erro != nil {
		return nil, erro
	}

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AuthorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AuthorNick,
		); erro != nil {
			return []modelos.Publicacao{}, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) CurtirPublicacao(publicacaoID uint64) error {
	var query = "update publicacoes set curtidas = curtidas + 1 where id = ?"
	statement, erro := repositorio.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Publicacoes) DescurtirPublicacao(publicacaoID uint64) error {
	var queryStatement = `update publicacoes set curtidas = 
    CASE WHEN curtidas > 0 THEN curtidas - 1 ELSE 0 END 
	where id = ?`

	statement, erro := repositorio.db.Prepare(queryStatement)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}
