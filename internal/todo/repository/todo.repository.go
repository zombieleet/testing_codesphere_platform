package todoRepository

import (
	"context"
	"errors"
	"fmt"

	Db "github.com/zombieleet/codesphere-test-todo-app/internal/infrastructure/db"
	todoEntity "github.com/zombieleet/codesphere-test-todo-app/internal/todo/domains/entities"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TodoRepository struct {
	col *mongo.Collection
}

type mongoTodo struct {
	ID    string `bson:"_id"`
	Title string `bson:"title"`
	Done  bool   `bson:"done"`
}

func GetTodoRepository() *TodoRepository {
	return &TodoRepository{
		col: Db.DB.Collection("todo"),
	}
}

func (t *TodoRepository) Save(ctx context.Context, snapshot todoEntity.TodoSnapshot) error {
	data := mongoTodo{
		ID:    snapshot.ID(),
		Title: snapshot.Title(),
		Done:  snapshot.IsDone(),
	}

	result, err := t.col.InsertOne(ctx, data)

	if err != nil {
		return err
	}

	if !result.Acknowledged {
		return fmt.Errorf("SaveOp not acknowledged")
	}
	return nil
}

func (t *TodoRepository) GetTodo(ctx context.Context, todoId string) (todoEntity.TodoSnapshot, error) {

	var dbResult mongoTodo

	err := t.col.FindOne(ctx, bson.M{"_id": todoId}).Decode(&dbResult)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, err
	}

	entity := todoEntity.HydrateTodo(dbResult.ID, dbResult.Title, dbResult.Done)

	if entity == nil {
		return nil, fmt.Errorf("failed to properly hydrate entity %s", todoId)
	}

	return entity, nil
}
