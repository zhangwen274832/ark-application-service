#!groovy

library 'fotoable-libs'
    def map = [:]

    //以下参数不需要修改
    //jenkins内置变量 job的名字
    def job_name = env.JOB_NAME.replaceAll("/","-")
    //构建分支，读取多分支构建的分支
    def BRANCH = env.BRANCH_NAME


    //以下参数需要研发人员修改
    //拉取代码库的地址
    map.put('REPO_URL',"git@gitlab.ftsview.com:aircraft/toc/ark-application-service.git")
    //代码的构建分支
    map.put('BRANCH', "${BRANCH}")

    //以下参数为多分支构建参数
 //以下参数为多分支构建参数
    if ("${BRANCH}" == "dev"){
        // 测试环境发版节点
        map.put('node','master')
        // 部署环境
        map.put('DEPENV','test')
        map.put('cluster', "ACK")
    //预发布环境
    } else if("${BRANCH}" == "dev-ack"){
        // 测试环境发版节点
        map.put('node','master')
        // 部署环境
        map.put('DEPENV','ack-test')
        map.put('cluster', "ACK")
    //生产环境
    } else if("${BRANCH}" == "release"){
        // 预发布环境发版节点
        map.put('node','master')
        // 部署环境
        map.put('DEPENV','ha')
        map.put('cluster', "ACK")

    //生产环境
    } else if ("${BRANCH}" == "master"){
        // 生产环境发版节点
        map.put('node','master')
        // 部署环境
        map.put('DEPENV','online')
        map.put('cluster', "ACK")


    }


// 环境使用方法(dev为测试环境请使用k8s;stage为预发布使用ekst;master为生产环境请使用eks)
Betta_ACK_v2 ("cluster",map)
