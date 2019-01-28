FROM node
# 基于node镜像 ，如果基于刚才的镜像则 FROM ryzebo/docker-nodejs-test:0.1


MAINTAINER Ryze <hz.bgning@gmail.com>
# 作者信息

RUN git clone https://git.oschina.net/ryzecode/Docker_Nodejs_Test.git \
    && cd ./Docker_Nodejs_Test \
    && npm install . 
# RUN 命令的内容大家应该看得懂

WORKDIR ./Docker_Nodejs_Test
# 设定初始目录


EXPOSE 3000
# 暴露3000端口


CMD ["node","bin/www"]
# 运行 node bin/www命令


