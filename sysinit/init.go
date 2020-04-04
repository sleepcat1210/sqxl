package sysinit

import "sqlx/utils/filecache"

func init(){
	cacheinit()
}
func cacheinit(){
	filecache.BasePath="./cache/filecache"
	filecache.ExpirseSec=10
	filecache.InitCache()
}