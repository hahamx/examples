# 1 启动动物园服务

    go run http_service.go

## 1.1   访问动物园一 

        访问猫头鹰
        http://127.0.0.1:3010/owls
            miao,miao...
    
        访问大象
            http://127.0.0.1:3010/elephants
                ang,ang...
        访问老虎
            http://127.0.0.1:3010/tiger
                aoo,aoo...

 ## 1.1    访问动物园二
    访问老虎

        http://127.0.0.1:3020/tiger
            aoo,aoo,aoo,aoo,...

    访问猫头鹰

        http://127.0.0.1:3020/owl
            wu,wu,wu,wu,...
    
    访问大象    

        http://127.0.0.1:3020/ele/
            ang,ang,ang,ang,...  
    喂食大象

        http://127.0.0.1:3020/ele/food
            corn,apple,rice,grass...
    

# 2 启动视频服务

    go mod init webService
    go mod tidy
    go mod vendor
    go run http_gin.go

## 2.1 使用web服务


    主界面
        curl  http://127.0.0.1:3040/
            
            message "pong"
    
    欢迎界面

        curl  http://127.0.0.1:3040/resource

            "message": "success", 
            "data": "welcome!", 
            "code": 200,

    查询列表

        {"code":200,
        "data":[
                    {"id":"en00029","name":"infinsh war III","path":"video.asdaliyun.com/oss/usubda","times":"180min"},
                    {"id":"en00028","name":"infinsh war II","path":"video.asdaliyun.com/oss/usubde","times":"170min"}],
        "message":"success"}
        
    具体信息

        curl  http://127.0.0.1:3040/resource/en00029
        {"code":200,
        "data":
            {"id":"en00029","name":"infinsh war III","path":"video.asdaliyun.com/oss/usubda","times":"180min"},
        "message":"success"}

        

