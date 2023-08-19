package database

type Pet struct {
	Id   int
	Name string
	Age  int
}

func DeletePet(id int) error {
	_, err := Db.Exec("DELETE FROM pets WHERE id = $1", id)
	if err != nil {
		return nil
	}
	return nil
}

func GetPet(_id int) (*Pet, error) {
	res, err := Db.Query("SELECT * FROM pets WHERE id = $1", _id)
	if err != nil {
		return nil, err
	}
	res.Next()
	var id int
	var name string
	var age int
	res.Close()
	return &Pet{
		Id:   id,
		Name: name,
		Age:  age,
	}, nil
}

func GetPets() ([]Pet, error) {
	res, err := Db.Query("SELECT * FROM pets")
	if err != nil {
		return nil, err
	}
	var pets []Pet = make([]Pet, 0)

	for res.Next() {
		var id int
		var name string
		var age int

		err := res.Scan(&id, &name, &age)

		if err != nil {
			return nil, err
		}

		pets = append(pets, Pet{
			Id:   id,
			Name: name,
			Age:  age,
		})
	}
	res.Close()
	return pets, nil
}

func (p *Pet) Save() error {
	if p.Id == -1 {
		_, err := Db.Exec("INSERT INTO pets (name, age) VALUES ($1, $2)", p.Name, p.Age)
		return err

	} else {
		_, err := Db.Exec("UPDATE pets SET name = $1, age = $2 WHERE id = $3", p.Name, p.Age, p.Id)
		return err
	}

}
