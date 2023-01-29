https://www.liwenzhou.com/posts/Go/import-local-package/


go module是Go1.11版本之后官方推出的版本管理工具，并且从Go1.13版本开始，go module将是Go语言默认的依赖管理工具。到今天Go1.14版本推出之后Go modules 功能已经被正式推荐在生产环境下使用了。

这几天已经有很多教程讲解如何使用go module，以及如何使用go module导入gitlab私有仓库，我这里就不再啰嗦了。但是最近我发现很多小伙伴在群里问如何使用go module导入本地包，作为初学者大家刚开始接触package的时候肯定都是先在本地创建一个包，然后本地调用一下，然后就被卡住了。。。

这里就详细介绍下如何使用go module导入本地包。


