#!groovy

// 封装HTTP Post请求
def HttpPost(reqUrl){
    def post = new URL(reqUrl).openConnection();
    def data = '{"msgtype": "text","text": {"content": "' + "${JOB_NAME}" + '构建完成，点击' + "${BUILD_URL}" + '查看详情"}}'

    post.setRequestMethod("POST")
    post.setDoOutput(true)
    post.setRequestProperty("Content-Type", "application/json")
    post.getOutputStream().write(data.getBytes("UTF-8"));
    post.getResponseCode();
}
// 执行企业微信通知
def WorkWechatNotify(reqUrl){
    response = HttpPost(reqUrl)
    return response
}
// 定义邮件内容
def Email(status){
    mail to: "$BUILD_USER_EMAIL",
        subject: "流水线${JOB_NAME}构建失败",
        body: "项目名：${JOB_NAME}构建失败,项目地址:${BUILD_URL},请点击${BUILD_URL}console查看并修复"
}
// 带颜色的格式化输出
def PrintMsg(value,color){
    colors = [
        'red'   : "\033[40;31m >>>>>>>>>>>${value}<<<<<<<<<<< \033[0m",
        'blue'  : "\033[47;34m ${value} \033[0m",
        'green' : "\033[40;32m >>>>>>>>>>>${value}<<<<<<<<<<< \033[0m"
    ]
    ansiColor('xterm') {
        println(colors[color])
    }
}

// 启动容器
def StartContainer(HostsArray){
  if (env.SVC != 'none') {
      echo "HostsArray:${HostsArray}"
      if (env.SVC == 'all') {
          SVC=""
      }
      HostsArray.each{ host->
          sh "make env-obv && rsync -crv ./* ubuntu@${host}:${PROJECT_PATH}"
          sh """ssh ubuntu@${host} << serverssh
              cd ${PROJECT_PATH} && sudo cp env .env
              make start SVC=${SVC}
              make clean-docker
              exit
              serverssh
          """
      }
  }
}

def BuildImage = {
    sh 'make clean-docker && make env && make base'
    if (env.SVC != 'none' && env.SVC != 'all') {
        sh 'make build SVC=${SVC}'
    }
    if (env.SVC == 'all') {
        sh 'make build SVC=""'
    }
}

def DeployK8S = {
    sh '''
        echo "====== docker tag: ${DOCKER_TAG} ======"
        make clean-docker && make env && make build SVC=${SVC}
        make env && pwd
        sudo bash ./scripts/k8s_deploy.sh ${NS} ${DOCKER_TAG} ${BRANCH_ENV} ${SVC}
    '''
}
pipeline {

  options {
    timestamps() // 显示日志时间戳 => 使用前先安装插件
    skipDefaultCheckout() // 隐式删checkout scm语句
    timeout(time: 20, unit: 'MINUTES') // 流水线超时设置:20min
  }

  agent none

  parameters {
    choice (
          name: 'SVC',
          choices: [
              'none', 'all','my-zero','access-api', 'access-rpc'
          ],
          description: '请选择需要发布的服务'
    )
    choice (
          name: 'SVC_DEV_IP',
          choices: [
              'all' ,'127.0.0.1'
          ],
          description: '开发环境发布服务IP'
    )
    choice (
          name: 'SVC_TEST_IP',
          choices: [
              'all', '127.0.0.1'
          ],
          description: '测试环境发布服务IP'
    )
    choice (
          name: 'SVC_PRO_IP',
          choices: [
              'all', '127.0.0.1', '127.0.0.1'
          ],
          description: '生产环境服务IP'
    )
    string (name: 'HUB_DOMAIN', defaultValue: 'docker.hub.com', description: 'docker-hub私有仓库地址')
    string (name: 'PROJECT', defaultValue: 'zero', description: '项目')
    string (name: 'DOCKER_TAG', defaultValue: 'latest', description: 'docker镜像tag，默认latest')
    string (name: 'PROJECT_PATH', defaultValue: '/home/ubuntu/zero', description: '项目位置')
    string (name: 'LOG_PATH', defaultValue: '/home/ubuntu/logs', description: '日志落地位置')
  }

  environment { // 全局变量定义
    HUB_DOMAIN = "${params.HUB_DOMAIN}"
    DOCKER_TAG = "${params.DOCKER_TAG}"
    PROJECT = "${params.PROJECT}"
    BRANCH_ENV = "${params.BRANCH_ENV}"
    SVC = "${params.SVC}"
    LOG_PATH = "${params.LOG_PATH}"
    PROJECT_PATH = "${params.PROJECT_PATH}"
    SVC_DEV_IP = "${params.SVC_DEV_IP}"
    SVC_TEST_IP = "${params.SVC_TEST_IP}"
    SVC_PRO_IP = "${params.SVC_PRO_IP}"
  }
  stages {
    stage('Check on Controller') {
        agent { label 'HOST' }
        stages {
            stage('Cleaning workspace') {
              steps { sh 'ls -l && sudo rm -rf ./*' }
            }

            stage('SCE Checkout') {
                steps { checkout scm }
            }

            stage('Stash artifacts') {
              steps {
                stash name: 'source', includes: '**', excludes: '**/.git,**/.git/**'
              }
            }

            stage('Tag Version Prepare') {
              steps {
                script {
                    // 部署环境
                    echo "当前版本:${DOCKER_TAG}"
                    if (env.DOCKER_TAG == '' || env.DOCKER_TAG == 'latest') {
                        echo "替换默认的latest版本"
                        DOCKER_TAG=sh(script: "git log -n 1 --pretty=format:%H", returnStdout: true)
                    }
                    echo "替换后版本:${DOCKER_TAG}"
                }
              }
            }
        }
    }
    stage('Transfer on JUMP_SERVER') {
        agent { label 'JUMPER' }
        stages {
            stage('Cleaning workspace') {
              steps { sh 'ls -l && sudo rm -rf ./*' }
            }
            stage('Unstash artifacts') {
              steps {
                unstash 'source'
              }
            }
        }
    }

    stage ('Deploy on Dev') {
      when { branch 'dev' }
      environment {
        PROJECT_PATH = "/home/ubuntu/jenkins/workspace/zero_dev"
        BRANCH_ENV = '.dev'
        NS = 'zero-dev'
      }
      agent { label 'JUMPER' }
      stages {
          stage('Building && Starting containers') {
            steps {
              script {
                  BuildImage()
                  def HostsArray = [
                      "127.0.0.1"
                  ]
                  if (env.SVC_DEV_IP != 'all') {
                     echo "env.SVC_DEV_IP != 'all' && ip = ${SVC_DEV_IP}"
                     HostsArray = [
                        env.SVC_DEV_IP
                    ]
                  }
                  StartContainer(HostsArray)
              }
            }
          }
        }
    }

    stage ('Deploy on TEST') {
      when { branch 'test' }
      environment {
        PROJECT_PATH = "/home/ubuntu/jenkins/workspace/zero_test"
        BRANCH_ENV = '.test'
        NS = 'zero-test'
      }
      agent { label 'JUMPER' }
      stages {
          stage('Building && Starting containers') {
            steps {
              script {
                  BuildImage()
                  def HostsArray = [
                      "127.0.0.1"
                  ]
                  if (env.SVC_TEST_IP != 'all') {
                     echo "env.SVC_TEST_IP != 'all' && ip = ${SVC_TEST_IP}"
                     HostsArray = [
                        env.SVC_TEST_IP
                    ]
                  }
                  StartContainer(HostsArray)
              }
            }
          }
        }
    }

    stage ('Deploy on Prod') {
      when { branch 'main' }
      environment {
        PROJECT_PATH = "/home/ubuntu/zero"
        BRANCH_ENV = '.pro'
        NS = 'zero-pro'
      }
      agent { label 'JUMPER' }
      stages {
        stage('Building && Starting containers') {
          steps {
            script {
                BuildImage()
                def HostsArray = [
                    "127.0.0.1",
                    "127.0.0.1"
                ]
                if (env.SVC_PRO_IP != 'all') {
                   echo "env.SVC_PRO_IP != 'all' && ip = ${SVC_PRO_IP}"
                   HostsArray = [
                      env.SVC_PRO_IP
                  ]
                }
                StartContainer(HostsArray)
            }
          }
        }
      }
    }
  }

  // 构建后操作
  post {
    always { script { PrintMsg("构建完成","green") }}
    success { script { currentBuild.description = "构建成功！" }}
    failure { script {
        currentBuild.description = "构建失败！"
        PrintMsg("构建失败,发送邮件和企业微信推送","red")
        Email("本次构建失败")
        WorkWechatNotify("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=your-key")
    }}
    aborted { script { currentBuild.description = "取消本次构建！" }}
  }
}
