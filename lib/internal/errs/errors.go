package errs

import "fmt"

const pkgName = "lib"

func NewErrIndexOutOfRange(length int, index int) error {

	return fmt.Errorf("[%s] index out of range [%d] with length %d", pkgName, length, index)
}
