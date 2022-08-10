## 正常的git提交流程

如下所示：

```shell
echo "# elastic-web-operator" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:Youngpig1998/elasticweb-operator.git
git push -u origin main
```

Note：更新代码之后，在commit后push前要先pull一下代码




如果你本地的分支名称和远程名称相同，则使用以下命令：

```
git push origin [branchName]
```


如果本地的分支跟远程的分支不一样，则指定本地和远程的分支的名称：

```
git push origin localBranchName:remoteBranchName
```





## Git分支的意义和使用方法

https://blog.csdn.net/qq_27674439/article/details/107412097





## Git删除本地分支

​	查看本地分支：

```shell
git branch  #前面带有*号的是当前分支
```


​	删除本地已合并的分支或者没有新改动的分支：

```shell
git branch -d [branchname]
```


​	强行删除分支：

```shell
git branch -D [branchName]
```

​	注意：
​		你是无法删除当前所在的分支的，因此，通常需要先切换到其他分支上：

```shell
git checkedout [branchName]
```

​		如果切换的分支已存在，就切换上去，否则就重新创建一个分支并切换上去



## GitHub删除远程分支

​	远程分支就是GitHub上创建的分支。删除远程分支其实是用push。
​	**注意：删除远程分支并不会删除本地分支**

```shell
git push origin --delete [branchname]
```

清理本地无效分支(远程已删除本地没删除的分支):

```shell
git fetch -p
```





## Git的branch和tag的区别

tag 对应某次 commit, 是一个点，是不可移动的。
branch 对应一系列 commit，是很多点连成的一根线，有一个HEAD 指针，是可以依靠 HEAD 指针移动的。

所以，两者的区别决定了使用方式，改动代码用 branch ,不改动只查看用 tag。

tag 和 branch 的相互配合使用，有时候起到非常方便的效果，例如 已经发布了 v1.0 v2.0 v3.0 三个版本，这个时候，我突然想不改现有代码的前提下，在 v2.0 的基础上加个新功能，作为 v4.0 发布。就可以 检出 v2.0 的代码作为一个 branch ，然后作为开发分支。



## 如何在本地添加删除Tag，并push到Github

```shell
#删除本地tag
git tag -d <tagname> 

#删除远端tag
git push origin --delete <tagname>  

#向远端推送本地tag
git push origin <tagname>

#向远端推送本地所有tag
git push origin --tags

# 查看commit详情，可以看到commit id
git log

#给某一次的commit打上tag
git tag v0.9 6df3aa614a45e687697a628a9174106822f79bfc

#可以用git show <tagname>查看标签信息：
git show testtag

```





## 修改GitHub仓库名称

1. 打开仓库地址，点击Settings
2. 更改仓库名称
3. 回到本地，通过以下命令查看之前的仓库名

```shell
#查看结果：
git remote -v

origin	git://github.com/Chloeqq/learngit.git (fetch)
origin	git://github.com/Chloeqq/learngit.git (push)

#通过以下命令进行修改
git remote set-url origin git@github.com:Chloeqq/Pytest_Auto_APi_Master.git

#查看结果:
git remote -v 

origin	git@github.com:Chloeqq/Pytest_Auto_APi_Master.git (fetch)
origin	git@github.com:Chloeqq/Pytest_Auto_APi_Master.git (push)

#通过以下命令，推送仓库
git push -u origin master   

```





