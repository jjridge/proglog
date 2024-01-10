package log

/*import (
	"fmt"
	"os"
	"path"

	api "github.com/travisjeffery/proglog/api/v1"
	"google.golang.org/protobuf/proto"
)*/

type segment struct {
	store                  *store
	index                  *index
	baseOffset, nextOffset uint64
	config                 Config
}
