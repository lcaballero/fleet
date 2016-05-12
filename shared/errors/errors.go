package errors

import "gopkg.in/bufio.v1"

type Errors []error

func (e Errors) Add(r ...error) Errors {
	return append(e, r...)
}
func (e Errors) Len() int {
	return len(e)
}
func (e Errors) Error() string {
	buf := bufio.NewBuffer([]byte{})
	for _, s := range e {
		buf.WriteString(s.Error())
		buf.WriteByte('\n')
	}
	return buf.String()
}
