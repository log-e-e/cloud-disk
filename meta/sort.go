/**
 *    FILENAME      :   sort.go
 *    AUTHOR        :   jihonghe
 *    DATE          :   19-10-24
 *    DESCRIPTION   :
 */
package meta

import "time"

const baseFormat = "2006-01-02 15:04:05"

type ByUploadTime []FileMeta

func (a ByUploadTime) Len() int {
    return len(a)
}

func (a ByUploadTime) Swap(i int, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a ByUploadTime) Less(i int, j int) bool {
    iTime, _ := time.Parse(baseFormat, a[i].UploadTime)
    jTime, _ := time.Parse(baseFormat, a[j].UploadTime)
    return iTime.UnixNano() > jTime.UnixNano()
}
