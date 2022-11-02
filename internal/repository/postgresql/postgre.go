package repository

import (
	"fmt"
	"log"

	authErrors "github.com/AdiKhoironHasan/golangProject1/pkg/errors"
	"github.com/AdiKhoironHasan/matkul/internal/models"
	"github.com/AdiKhoironHasan/matkul/internal/repository"
	"github.com/jmoiron/sqlx"
)

const (
	SaveMahasiswa       = `INSERT INTO kampus.mahasiswas (nama, nim, created_at) VALUES ($1, $2, now()) RETURNING id`
	SaveMahasiswaAlamat = `INSERT INTO kampus.mahasiswa_alamats (jalan, no_rumah, created_at, id_mahasiswas) VALUES ($1,$2, now(), $3)`
	Login               = `SELECT id FROM kampus.users WHERE email = %s AND password = %s`
)

var statement PreparedStatement

type PreparedStatement struct {
	// login *sqlx.Stmt //membungkus query untuk melindungi dari sql inject
}

type PostgreSQLRepo struct {
	Conn *sqlx.DB
}

func NewRepo(Conn *sqlx.DB) repository.Repository {

	repo := &PostgreSQLRepo{Conn}
	InitPreparedStatement(repo)
	return repo
}

func (p *PostgreSQLRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Conn.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *PostgreSQLRepo) {
	statement = PreparedStatement{
		// login: m.Preparex(Login),
	}
}

func (p *PostgreSQLRepo) Login(dataLogin *models.UserModels) (bool, error) {
	var idUser int64
	var query string
	query = fmt.Sprintf(Login, dataLogin.Email, dataLogin.Password)

	err := p.Conn.Select(&idUser, query)

	if err != nil {
		log.Println("Failed Query GetMahasiswaAlamat: ", err.Error())
		return false, fmt.Errorf(authErrors.ErrorDB)
	}

	if idUser < 1 {
		return false, fmt.Errorf(authErrors.ErrorDataNotFound)
	}

	return true, nil
}
