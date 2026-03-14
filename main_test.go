package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}
```

---

**`go.mod`**
```
module voxdeploy-test

go 1.24