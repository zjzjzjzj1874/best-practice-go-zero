#!/bin/bash
# 脚本替换deployment的命名空间 需要传入参数,第一个参数:命名空间,第二个参数:镜像Tag

# 替换部署文件的命名空间
# shellcheck disable=SC2120
ns_image_replace() {

  # 命名空间使用传入变量
  ns=$1

  if [ -z "${ns}" ]; then
    echo "请传入命名空间"
    exit
  fi

  env=$3
  if [ -z "${env}" ]; then
    echo "请传入环境"
    exit
  fi

  tag=$2 #  dev和test分支自动分配一个 => prod则根据用户填写的
  pro_prefix="pro"
  if [[ ${env} =~ ${pro_prefix} ]]; then # main分支需要传入镜像标签
    if [ -z "${tag}" ]; then
      echo "请传入镜像标签"
      exit
    fi
  fi

  echo " ================= 命名空间:${ns},镜像Tag:${tag},env:${env} ================"

  #  deployment替换
  dep_path="./scripts/zero/"
  files=$(ls ${dep_path} | grep ".yaml")

  for file in ${files}; do
    sed -i "s/{{ns}}/${ns}/g" "${dep_path}${file}"   # 替换ns
    sed -i "s/{{tag}}/${tag}/g" "${dep_path}${file}" # 替换tag
    sed -i "s/{{env}}/${env}/g" "${dep_path}${file}" # 替换tag
  done
}

# 替换配置文件
regen_configmap() {
  # 命名空间使用传入变量
  ns=$1
  if [ -z "${ns}" ]; then
    echo "请传入命名空间"
    exit
  fi

  BRANCH_ENV=$2 # 部署分支
  svc=$3        # 启动服务
  # 查找当前文件夹中所有k8s配置文件
  files=$(find "$PWD" | xargs ls -d | grep ".${BRANCH_ENV}.yaml") # 查找所有k8s的配置文件,以后如果需要k8s部署,那么grep将会替换为{{dev}}.yaml,不同的分支用不同的配置

  echo " BRANCH_ENV ======== ${BRANCH_ENV},SVC ======== ${svc}"
  oldIfs=$IFS
  IFS=$'\n'
  for file in ${files}; do
    name=${file##*/}
    name=${name%%.*} # 截取configmap的名字

    if [ -z "${svc}" ]; then # 空字符串 => 部署所有服务
      echo "重新生成configmap文件名 conf-${name}"
      kubectl delete configmap -n "${ns}" "conf-common-${name}"                       # 删除原来configmap
      kubectl create configmap -n "${ns}" "conf-common-${name}" --from-file="${file}" # 创建新的configmap
      continue
    fi

    if [[ ${name} =~ ${svc} ]]; then
      echo "单个服务:重新生成configmap文件名 conf-${name}"
      kubectl delete configmap -n "${ns}" "conf-common-${name}"                       # 删除原来configmap
      kubectl create configmap -n "${ns}" "conf-common-${name}" --from-file="${file}" # 创建新的configmap
    fi
  done
  IFS=${oldIfs} # 恢复默认的分割
}

# 部署的deployment
apply_deployment() {
  echo " =============== 部署deployment  =============== "
  dep_path="./scripts/zero/"
  files=$(ls ${dep_path} | grep ".yaml")

  svc=$1 # 待部署服务
  for file in ${files}; do
    if [ -z "${svc}" ]; then # 空字符串 => 部署所有服务
      echo "部署deployment:${file}"
      kubectl apply -f "${dep_path}${file}"
      continue
    fi

    # 部署单个服务
    if [[ ${file} =~ ${svc} ]]; then
      echo "部署服务:${svc}"
      kubectl apply -f "${dep_path}${file}"
    fi
  done
}

ns_image_replace "$1" "$2" "$3" # 调用命名空间替换方法 第一个参数:ns 第二个参数:tag 第三个参数:分支参数 dev/test/pro 第四个参数:服务名
regen_configmap "$1" "$3" "$4"  # 替换configmap 第三个参数:分支参数 dev/test/pro
apply_deployment "$4"           # 部署deployment
