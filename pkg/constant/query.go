package constant

const (
	CreateTaskQuery = `
		INSERT INTO task (
			title,
			status
		) VALUES (
			$1,
			1
		)
	`
	GetAllTaskQuery = `
		SELECT
			id,
			title,
			status
		FROM 
			task
		WHERE 
			status = 1
	`
	GetTaskByIDQuery = `
		SELECT
			id, 
			title,
			status
		FROM 
			task 
		WHERE 
			id = $1 
			AND status = 1
	`
	UpdateTaskQuery = `
		UPDATE
			task
		SET 
			title = $2 
		WHERE 
			id = $1 
			AND status = 1
	`
	DeleteTaskQuery = `
		UPDATE 
			task 
		SET 
			status = 0 
		WHERE 
			id = $1
	`
)
