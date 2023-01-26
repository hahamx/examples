# 2 git 命令速查
   
   一般的默认分支  master 

   默认远程版本库  origin 
   
   默认开发分支  Head 

   Head的前一个提交  Head^  

   简要工作流如下：

     ---------------------  ----------------  -----------------
     |                   |  |              |  |               | 
    CREATE-----BROWSE    CHANGE---REVERT    UPDATE---BRANCH---COMMIT---PUBLISH
    (创建)     (查看)     (更改)  (还原)     (更新)     (分支)  (提交)  (推送到远程)
                         |  |     | | |                |     |  |  
                         |   -----  | ------------------     |  |
                         |        - -------------------------   |
                         ----------------------------------------

## 2.0 创建
  
  从已有数据创建 代码仓库

    cd ./existing_data
    git init
    git add .

  从已有仓库 创建新的代码仓库，克隆远程代码仓库

    git clone ~/existing_repo ~/new/repo
    git clone git://host.org/project.git
    git clone ssh://user@host.org/project.git

## 2.1 修改和提交 

状态查询

   git status 

查看变动内容

   git diff

跟踪全部改动过的文件

   git add .

跟踪指定文件

   git add <file>

文件更名

   git mv <old> <new>

删除文件

   git rm <file>

停止跟踪文件但是不删除

   git rm --cached <file>

提交全部更新过的文件

   git commit -m 'commit for what'   

 修改最后一次提交
    
      git commit --amend               

## 2.2 还原操作
  
  返回到最近提交状态   

    git reset --hard  //注意这是不可撤销的

  还原到最新提交   

     git revert HEAD //创建新的提交

  还原到指定提交

     git revert <ID>  //创建新的提交
  
  修复最后一次提交

     git revert <ID>  //创建一个新的提交
  
  检出文件的某个版本

     git checkout <ID> <FILE>

## 2.3 分支管理
  
   检出并切换到分支 B0, 可以是指定分支或者标签为B0 

      git checkout <B0>  
     
   合并分支 B1 到 B2

      1 git checkout <B2>
      2 git merge <B1>     //合并b1到 b2
   
   衍合入指定分支到当前分支

      git rebase <B1>

   创建分支 基于当前分支某个 HEAD 地址 创建一个新的分支 B0

       git branch <B0>

       git checkout -b <B0>

   基于其他分支 创建新的分支 B0，并切换过去

      git checkout -b <B0> <OTHER>

   删除分支 B0

      git branch -d <B0>
   
   * 远程分支 操作

   查看远程代码库信息

      git remote -v  
   
   查看指定远程版本库信息

      git remote show <remote-b>
   
   添加远程代码库

      git remote add <remote-b> <url>

    从远程库获取代码

      git fetch <remote-b>

    快速更新下载代码

      git pull <remote-b> <local-a>

    推送代码及快速合并

      git push <remote-b> <branch>

    删除远程分支或者标签

      git push <remote-b>:<branch/tag-name>

    上传全部标签

      git push --tags
 

## 2.4 解决合并的问题
   
   查看合并冲突问题
     
     git diff

   针对某个文件查看合并的冲突，

     git diff --base <FILE>

   针对更改查看合并冲突

     git diff --ours <FILE>

   针对其他更改查看冲突

     git diff --theirs <FILE
   
   跟踪文件更改记录, 对比两个记录的不同 ID1  ID2

      git diff

      git diff <ID1> <ID2>

   丢弃一个冲突的补丁

     git reset --hard
     git rebase --skip

   解决冲突后，合并

     git add <CONFLICTING, FILE>
     git rebase --continue
   
   撤销工作目录的全部未提交文件的内容
   
     git reset --hard HEAD
   
   撤销指定的未提交文件的修改内容
      
      git checkout HEAD <file>

   撤销指定提交

      git  revert <commit>

## 2.5 查询
  
   工作目录中的文件更改状态

      git status

   历史更改记录

      git log

   指定文件历史更改记录的不同之处

      git log -p <FILE> <DIRECTORY>

    列表方式查看指定文件的更改者和更改内容

      git blame <FILE>

    某个提交ID 的具体记录

      git show <ID> 
   
   某个特定文件的特定ID记录

      git show <ID>:<FILE>

   全部本地分支， 将在控制台以 * 标注本地的分支

      git branch

## 2.6 更新
  
   从远程获取最近更改  --不执行合并

      git fetch 

    从远程拉取最新更改内容  -- 执行获取，并随后执行合并

      git pull 

    应用一个他人发给你的更改

      git am -3 patch.mbox   --- 如果 操作冲突，则处理冲突  
      git am --resolved

## 2.7 发布
    
    提交本地更改

      git commit -a

    为其他开发成员准备一个更改

       git format-patch origin
    
    推送更改到远程

       git push

    制造一个版本或者里程碑

       git tag v1.0