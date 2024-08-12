package main


type Todo struct {
	Title string
	Completed bool
	CreatedAt time.time
	CompletedAt *time.Time
}
type Todos []Todo

func (todos *Todos) add(title string){
	todo := Todo{
		Title: title,
		Completed: false,
        CompletedAt: nil,
		CreatedAt: time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error{
	if index < 0 || index >- len(*todos){
		
	}
}