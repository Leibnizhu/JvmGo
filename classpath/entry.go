package classpath
import "os"
import "strings"
const pathListSeperator = string(os.PathListSeparator)
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

//工厂类
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeperator) {
		return newCompositeEntry(path)
	} else if strings.HasSuffix(path,"*"){
		return newWildcardEntry(path)
	} else if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR") || 
	strings.HasSuffix(path,".zip") || strings.HasSuffix(path,".ZIP"){
		return newZipEntry(path)
	} else {	
		return newDirEntry(path)
	}
}