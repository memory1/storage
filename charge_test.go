package storage

import (
	"github.com/stretchr/testify/assert"
    "testing"
    "fmt"
)

func TestStorageCharge(t *testing.T) {
    s1 := Storage{1024, BlockStorageType}
    s2 := Storage{512, FileStorageType}
    s3 := Storage{128, ObjectStorageType}

    fmt.Println("BlockStorageType: ", BlockStorageType)

    l1 := Lease{3,  s1}
    l2 := Lease{4,  s2}
    l3 := Lease{14, s3}

	var tenant Tenant
	tenant = append(tenant, l1)
	tenant = append(tenant, l2)
	tenant = append(tenant, l3)

    total, levels := tenant.charge()
	assert.Equal(t, 3910, total)
    assert.Equal(t, 1, levels)
}
