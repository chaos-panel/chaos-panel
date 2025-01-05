package dlock

type DistributedLock struct {
}

func GetDistributedLockByDB() *DistributedLock {
	return &DistributedLock{}
}

func (s *DistributedLock) Lock(key string) bool {
	return true
}

func (s *DistributedLock) Unlock(key string) bool {
	return true
}
