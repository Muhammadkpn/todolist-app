package impl

import "fmt"

func generateTaskByIDKey(id int64) string {
	return fmt.Sprintf("task_%+v", id)
}
