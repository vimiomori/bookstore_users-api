package errors

// type RestErr interface {
// 	Message() string
// 	Status() int
// 	Error() string
// 	Causes() []interface{}
// }

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}
