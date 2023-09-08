package lock

import (
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func key(id string) string {
	return fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixAccountLock, id)
}

func Lock(id string) error {
	return redis2.TryLock(key(id), 0)
}

func Unlock(id string) error {
	return redis2.Unlock(key(id))
}
