#UPDATER
##1 创建UPDATE：
`func CreateUpdate(content UpdateContent, outputWriter io.Writer) error`<br>
<br>
**其中:**<br>
`type UpdateContent struct {`<br>`
`&nbsp;&nbsp;`Name`&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`string `&nbsp;&nbsp;` // app可执行文件所在目录名`<br>`
`&nbsp;&nbsp;`Paths   `&nbsp;&nbsp;&nbsp;`[]string  `&nbsp;&nbsp;` //要添加的更新文件路径`<br>`
`&nbsp;&nbsp;`Scripts []string     `&nbsp;&nbsp;`//脚本路径`<br>`
`&nbsp;&nbsp;`Version VersionControl  `&nbsp;&nbsp;`// 允许从哪些版本升级过来`<br>
}`<br>
<br>

##2 执行更新:  
`func Start(inputUpdate io.Reader, currentVersion string) error`<br>

##3 手动回滚:
`func PerformRollback(manager , backupPath, appName string) error `<br>
`manager`:&nbsp;辅助进程路径，默认放在app目录下<br>
`backupPath`:&nbsp;备份文件目录，默认放于app/BACKUPS<br>
`appName`:&nbsp;app可执行文件所在目录名<br>