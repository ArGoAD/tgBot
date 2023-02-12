package files

import (
	e "awesomeProject/lib/error"
	"awesomeProject/storage"
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

//this type init the interface

type Storage struct {
	basePath string
}

const defaultPerm = 0774

var ErrNoSavedPages = errors.New("no saved page")

//Create new Storage struct

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

//Init methods

func (s Storage) Save(page *storage.Page) (err error) {
	//For errors
	defer func() { err = e.WrapIfErr("can't save page", err) }()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	//Get fileName
	fName, err := fileName(page)
	if err != nil {
		return err
	}
	fPath = filepath.Join(fPath, fName)

	//Create file
	file, err := os.Create(fPath)
	if err != nil {
		return err
	}

	//Close created file
	defer func() { _ = file.Close() }()

	//refactor page for format gob and convert in current file
	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}
	return nil

}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = e.WrapIfErr("can't pick random page", err) }()

	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, ErrNoSavedPages
	}

	// create list of files
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	//get random file
	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))

}

func (s Storage) Remove(p *storage.Page) error {
	fileName, err := fileName(p)
	if err != nil {
		return e.Wrap("can't remove file", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)
	if err := os.Remove(path); err != nil {
		msg := fmt.Sprintf("can't remove file %s", path)
		return e.Wrap(msg, err)
	}
	return nil
}

func (s Storage) IsExists(p *storage.Page) (bool, error) {
	fileName, err := fileName(p)
	if err != nil {
		return false, e.Wrap("can't check if file exists", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)
	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, err
	case err != nil:
		msg := fmt.Sprintf("can't check if file %s exists", path)

		return false, e.Wrap(msg, err)
	}
	return true, nil
}

// Open decode page
func (s Storage) decodePage(filePath string) (*storage.Page, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, e.Wrap("can't decode page", err)
	}

	defer func() { _ = f.Close() }()

	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, e.Wrap("can't decode page", err)
	}

	return &p, nil

}

// definition name file

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
