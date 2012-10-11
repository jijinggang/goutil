PACKAGE

package file
    import "."


FUNCTIONS

func CopyFile(srcName, dstName string) (written int64, err error)
    复制文件，把srcName复制到dstName

func CreateDir(file string)
    递归创建文件所在的目录

func Exists(path string) bool
    判断文件或路径是否存在

func FormatPath(path string) string
    格式化路径，统一用/替代\，并且去掉最后面的/

func GetAllFiles(dir string) []string
    获取一个目录的所有子文件

func GetFileMd5(file string) string
    获取文件的Md码

func GetPath(file string) string
    获取一个文件的路径，不带最后的/

func GetRelativePath(file string, dir string) string
    获取file或目录相对其一个父目录dir的相对路径名

func WriteStringFile(file string, str string) (written int, err error)
    写文本文件，如果文件存在，则覆盖


