# assert

```go
import (
	"fmt"
	"testing"

	"code.gopub.tech/assert"
)

func TestXxx(t *testing.T){
    assert.True(t, 1 == 1)
    _ = assert.True
    _ = assert.False
    _ = assert.Nil
    _ = assert.NotNil
    _ = assert.Eq
    _ = assert.Ne
    _ = assert.DeepEqual
}

```