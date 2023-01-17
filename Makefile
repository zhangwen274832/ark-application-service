ver=$(ver)
buildNo=$(buildNo)
config=$(config)
OVERLAYS=deployments/kustomize/overlays
# Docker 测试地址
DockerRegistryTest := harbor.nuclearport.com/aircraft/ark-application-service-test
# Docker 预发地址
DockerRegistryStage := harbor.nuclearport.com/aircraft/ark-application-service-stage
# Docker 生产地址
DockerRegistry := harbor.nuclearport.com/aircraft/ark-application-service-prod

# ==============================================================================
# OPTIONS
define USAGE_OPTIONS

Options:
  ver       	Docker image version.
            	Example:
              	  test: make ci-test ver=12
              	  stage: make ci-stage ver=12
              	  prod: make ci-prod ver=v1.2.0

  buildNo   	Build No, jenkins build number.
            	Example: Reference config

  config    	Profiles deployed to k8s.
            	Example:
              	  test: make cd-test ver=12 config=~/.kube/config.test
              	  stage: make cd-stage ver=12 config=~/.kube/config.stage
              	  prod: make cd-prod ver=v1.2.0 buildNo=12 config=~/.kube/config.prod
endef
export USAGE_OPTIONS

.PHONY:help
help: Makefile
	@echo "Usage: make <Targets> <Options> ... \n\nTargets"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"

## run: start service.
.PHONY:run
run:
	@echo "start service"
	go run -race main.go  --aman_addr=https://aman-internal.akgoo.net --aman_project_id=ark-application-service --aman_env_id=1 --version=latest


## ci-test: build test image and push.
.PHONY:ci-test
ci-test:
	@echo "ci-test"
	docker build -t $(DockerRegistryTest):$(ver) -f build/Dockerfile .
	docker push $(DockerRegistryTest):$(ver)


## ci-stage: build stage image and push.
.PHONY:ci-stage
ci-stage:
	@echo "ci-stage"
	docker build -t $(DockerRegistryStage):$(ver) -f build/Dockerfile .
	docker push $(DockerRegistryStage):$(ver)


## ci-prod: build prod image and push.
.PHONY:ci-prod
ci-prod:
	@echo "ci-prod"
	docker build -t $(DockerRegistry):$(ver) -f build/Dockerfile .
	docker push $(DockerRegistry):$(ver)

## cd-test: deploy test image and push.
.PHONY:cd-test
cd-test:
	@echo "cd-test start"
	@echo $(ver)
	@echo $(config)
	cd $(OVERLAYS)/test && kustomize edit set image $(DockerRegistryTest):$(ver) && kustomize edit add annotation kubesphere.io/description:'ark-application-service-测试环境-'$(ver)
	kustomize build $(OVERLAYS)/test | kubectl --kubeconfig $(config) apply -f -
	@echo "success"



## cd-stage: deploy stage image and push.
.PHONY:cd-stage
cd-stage:
	@echo "cd-stage start"
	@echo deploy-cd-stage
	@echo $(ver)
	@echo $(config)
	cd $(OVERLAYS)/stage && kustomize edit set image $(DockerRegistryStage):$(ver) && kustomize edit add annotation kubesphere.io/description:'ark-application-service-预发环境-'$(ver)
	kustomize build $(OVERLAYS)/stage | kubectl --kubeconfig $(config) apply -f -
	@echo "success"



## cd-prod: deploy prod image and push.
.PHONY:cd-prod
cd-prod:
	@echo "cd-prod start"
	@echo $(ver)
	@echo $(buildNo)
	@echo $(config)
	cd $(OVERLAYS)/prod && kustomize edit set namesuffix -- -$(buildNo) && kustomize edit set label version:$(ver) && kustomize edit set image $(DockerRegistry):$(ver) && kustomize edit add annotation kubesphere.io/description:'ark-application-service-生产环境-'$(ver)
	kustomize build $(OVERLAYS)/prod | kubectl --kubeconfig $(config) apply -f -
	@echo "success"

