# My blog

- Gin
    * Http web框架
- Gorm
    * 操作数据库
    * 软删除
- Go ini
    * 读取配置文件

仅使用gin托管前端页面，需将打包好的文件放入`static`文件中，并适当修改托管配置

## 开发日志

- 工程化开发讲究条理
    * 剥离配置文件，方便修改和运维
    * 读取配置文件调用同一的api，方便你我他
    * 写好模型model，然后再用api负责调用，当然api可可能会有许多版本，也可进行细分
        + 模型里只需关注自己的业务，需要组合逻辑的话再有api负责处理
- 不要明文储存密码，不是每个程序员素养都ok，也防止黑客盗取
    * 不同加密算法开销可能有所不同，还需仔细研究
- 文件上传
    * 文件、图片的传输很耗带宽，和文章一起储存在服务端本地不是一个明智的选择，一般使用第三方来保存图片(这里使用七牛云)
- 日志分割，方便我们查看
    * github.com/lestrrat-go/file-rotatelogs
    * github.com/rifflock/lfshook
- validate验证
    * 防止用户绕过前端做些非法操作，如创建超级用户
    * 也可限制密码长度、用户名字符类型等
- 配置跨域
- 富文本编辑器tinymce
