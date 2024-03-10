package dbError

import "fmt"

type CreateTableError struct {
	TableName string
	Query string
}

func (e *CreateTableError) Error() string {
	return fmt.Sprintf("Failed to create table.\nTableName: %v\nQuery: %v", e.TableName, e.Query)
}

func NewCreateTableError(tableName string, query string) *CreateTableError {
	return &CreateTableError{
		TableName: tableName,
		Query: query,
	}
}



type EncryptPasswordError struct {}

func (e *EncryptPasswordError) Error() string {
	return fmt.Sprintln("Failed to encrypt password");
}

func NewEncryptPasswordError() *EncryptPasswordError {
	return &EncryptPasswordError{}
}



type UserAlreadyExistsError struct {
	Username string
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("User already Exists.\nUsername: %v\n", e.Username)
}

func NewUserAlreadyExistsError(username string) *UserAlreadyExistsError {
	return &UserAlreadyExistsError{
		Username: username,
	}
}


type DbRowScanError struct {
	DbName string
}

func (e *DbRowScanError) Error() string {
	return fmt.Sprintf("Failed to scan db row scan.\nDatabase name: %v\n", e.DbName)
}


func NewDbRowScanError(dbName string) *DbRowScanError {
	return &DbRowScanError{
		DbName: dbName,
	}
}


type DbQueryError struct {
	Query string
}

func (e *DbQueryError) Error() string {
	return fmt.Sprintf("Failed to run db query.\nQueryName: %v\n", e.Query)
}

func NewDbQueryError(query string) *DbQueryError {
	return &DbQueryError{
		Query: query,
	}
}


type InsertUserError struct {
	Username string
}

func (e *InsertUserError) Error() string {
	return fmt.Sprintf("Failed to insert user into db.\n Username: %v\n", e.Username)
}

func NewInsertUserError(username string) *InsertUserError {
	return &InsertUserError{
		Username: username,
	}
}


type PasswordDoesNotMatchError struct {}

func (e *PasswordDoesNotMatchError) Error() string {
	return fmt.Sprintln("Password does not match.")
}

func NewPasswordDoesNotMatchError() *PasswordDoesNotMatchError {
	return &PasswordDoesNotMatchError{}
}


type GenerateJwtTokenError struct {}

func (e *GenerateJwtTokenError) Error() string {
	return fmt.Sprintln("Failed to generate jwt token.")
}

func NewGenerateJwtTokenError() *GenerateJwtTokenError {
	return &GenerateJwtTokenError{}
}


type FindUserError struct {
	Username string
}

func (e *FindUserError) Error() string {
	return fmt.Sprintf("Couldn't find such user.\nUsername: %v\n", e.Username)
}

func NewFindUserError(username string) *FindUserError {
	return &FindUserError{
		Username: username,
	}
}


type MapToStructError struct {
	MappedStruct interface{}
}

func (e *MapToStructError) Error() string {
	return fmt.Sprintf("Failed to map data into struct. Struct: %v\n", e.MappedStruct)
}

func NewMapToStructError(mappedStruct interface{}) *MapToStructError{
	return &MapToStructError{
		MappedStruct: mappedStruct,
	}
} 