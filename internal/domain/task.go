package domain

import "github.com/google/uuid"

type Task struct {
	Id          uuid.UUID  `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	TagId       uuid.UUID  `json:"tagId" db:"tag_id"`
	Tags        Tags       `json:"tags" db:"tags"`
	Difficulty  int        `json:"difficulty" db:"difficulty"`
	ExecutorId  uuid.UUID  `json:"executorId" db:"executor_id"`
	Executor    UserPicker `json:"executor" db:"executor"`
	AuthorId    UserPicker `json:"authorId" db:"author_id"`
	Author      UserPicker `json:"author" db:"author"`
}

type Tags []Tag

type Tag struct {
	Id    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
}
