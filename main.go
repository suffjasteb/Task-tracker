package main

func main() {
	todos := Todos{}
	Storange := NewStorange[Todos]("todos.json")
	Storange.Load(&todos)
	cmdFlags := newCmdFlags()
	cmdFlags.Execute(&todos)
	Storange.Save(todos)
}