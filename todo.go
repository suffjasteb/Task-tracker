package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo {
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex (index int) error {
	if index < 0 || index > len(*todos) {
		err := errors.New("invalid Index")
		fmt.Println(err)
		return err
	}
	return nil
}


func (todos *Todos) delete(index int) error {
	t := *todos // Mengambil slice asli yang ditunjuk oleh pointer todos
	if err := t.validateIndex(index); err != nil {
		return err // Jika index tidak valid, kembalikan error
	}

	// Menghapus item pada index yang diberikan
	*todos = append(t[:index], t[index+1:]...)
	return nil // Tidak ada error, berarti sukses
}

func (todos *Todos) toggle(index int) error {
	t := *todos// Mengambil slice asli yang ditunjuk oleh pointer todos
	if err := t.validateIndex(index); err != nil {
		return err // Jika index tidak valid, kembalikan error
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completedTime := time.Now()
		t[index].CompletedAt = &completedTime
	}

	 t[index].Completed = !isCompleted
	
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos// Mengambil slice asli yang ditunjuk oleh pointer todos
	if err := t.validateIndex(index); err != nil {
		return err // Jika index tidak valid, kembalikan error
	}

	

	 t[index].Title = title
	
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed,t.CreatedAt.Format(time.RFC1123), completedAt )
	}
	table.Render()
}