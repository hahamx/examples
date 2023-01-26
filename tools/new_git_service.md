# 1 git服务器建立
	ubuntu 20  官方安装说明 https://about.gitlab.com/install/#ubuntu
	注意: 在安装时，最后一步 https 如果没有域名可用，或者不需要使用加密https协议，需要替换域名为ip地址
	sudo EXTERNAL_URL="https://gitlab.example.com" apt-get install gitlab-ee
	替换为
	sudo EXTERNAL_URL="http://192.168.136.130" apt-get install gitlab-ee

	# 如果安装完成后，gitlab服务器仍然不能工作
	sudo gitlab-ctl restart
	# 配置和启用外部ip地址访问服务，编辑配置,并重启服务
	sudo vim /etc/gitlab/gitlab.rb
		external_url "http://192.168.10.1"
	sudo gitlab-ctl reconfigure 

	# 相对路径的禁用
	sudo gitlab-ctl restart unicorn

	# 登录并设置密码

	# 生成ssh密钥对 RSA算法 至少2048位，-C标志带有带引号的注释，例如电子邮件地址，是标记SSH密钥的一种可选方式
	ssh-keygen -t rsa -b 4096 -C "autocommsky@gmail.com"

	# 更新密钥
	ssh-keygen -p -f /path/to/ssh_key

	# 设置密钥
		在web界面。 用户头像--设置--ssh密钥，将生成的公钥pub添加到页面
		在git ssh 客户端添加私钥地址
		设置ssh代理环境
		eval $(ssh-agent -s)
		添加(如果ssh产生私钥在当前目录)
		ssh-add ./id_rsa
## 生成一个 32位随机密钥
	
	openssl rand -hex 32
	
## 调试
	综合安装的调试指令 sudo gitlab-rails console
	启动控制台后添加日志监控 ActiveRecord::Base.logger = Logger.new(STDOUT)
## gitlab 服务器添加和管理用户
	run: alertmanager:

	run: gitaly: 

	run: gitlab-exporter: 

	run: gitlab-workhorse:
	run: grafana: 
	run: logrotate: 
	run: nginx: 
	run: node-exporter: 
	run: postgres-exporter: 
	run: postgresql: 
	run: prometheus: 
	run: puma: 
	run: redis:
	run: redis-exporter: 
	run: sidekiq: 


## 语言设置
	用户登录后，在右上角 设置--用户设置--偏好设置

## CI/CD 持续集成，持续发布
	安装gitrunner参考链接:	
	https://docs.gitlab.com/runner/install/linux-manually.html
	注册gitrunner，绑定gitlab地址url，token和执行者
	注册参考文档
	https://docs.gitlab.com/runner/register/index.html#docker
		在gitlab服务器的设置-CICD-展开可以查到当前gitlab服务的
		url:http://192.168.136.130
		token:Qi1yC28qsoFVzQtRyaus

	continue delopy， comtinue intergration
	Step 1:  add .gitlab-ci.yml in the root folder of your project/repo
	Step 2: git commit and git push to git repo
	Step 3: Create Gitlab runner for the project   # 本机配置失败#  需要安装，注册runner到你的gitlab服务器 http://192.168.136.130
	Step 4: Start the runner
	Step 5: Make any change in the project -> commit -> push

	# gitlab runner 注册
	sudo gitlab-runner register --config /tmp/test-config.toml --template-config /tmp/test-config.template.toml --non-interactive --url http://192.168.136.130 --registration-token "your gitlab token" --name test-runner --tag-list bash,test --locked --paused --executor shell

	或者
	sudo gitlab-runner start   # 一步步配置
	--url http://192.168.136.130 
	--registration-token "your gitlab token" 
	--name test-runner 
	--tag-list bash,test 
	--locked 
	--paused 
	--executor shell

# 总结

	yml文件放在项目的根目录下，格式也正确，runner也注册成功跑起来了但是项目CICD中没有用户的项目提交记录。
 