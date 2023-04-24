## How create a simple class constructor in go

* First, you have to cretae your struct, that is a objection of a object for our code.

```
type Artist struct {
	id   uint8
	name string
	age  uint8
}
```

* Second, create your get Methods, they will return the data of the private attribute 

```
func (instance *Artist) ID() uint8 {
	return instance.id
}

func (instance *Artist) Name() string {
	return instance.name
}

func (instance *Artist) Age() uint8 {
	return instance.age
}
```

* Third, create your constructor: 

```
func New(id uint8, name string, age uint8) *Artist {
	return &Artist{
		id:   id,
		name: name,
		age:  age,
	}
}
```
