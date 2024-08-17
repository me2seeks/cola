# set shell
set windows-shell := ["powershell.exe", "-c"]

# set `&&` or `;` for different OS
and := if os_family() == "windows" {";"} else {"&&"}

#====================================== alias start ============================================#

#======================================= alias end =============================================#


#===================================== targets start ===========================================#

# default target - `just` 默认目标
default: lint test

# go build
[unix]
build:
    @echo "Building..."
    @GIN_MODE=release go build -tags="jsoniter" -ldflags "-s -w" -o {{bin}} {{main_file}}
    @echo "Build done."

[windows]
build:
    @echo "Building..."
    @$env:GIN_MODE="release" {{and}} go build -tags="jsoniter" -ldflags "-s -w" -o {{bin}} {{main_file}}
    @echo "Build done."

# go run
run: 
    @go run {{main_file}}
    # @go run -ldflags "-X 'config.Mode=debug'" {{main_file}}

# go test
test:
    @go test -v {{join(".", "...")}}

# generate swagger docs - 生成 swagger 文档
# swag: dep-swag
#     @cd {{server}} {{and}} swag init -g swagger.go

# format code - 格式化代码
fmt: dep-gofumpt
    @echo "Formatting..."
    @gofumpt -w  .

# lint - 代码检查
lint: dep-golangci-lint
    @echo "Linting..."
    @go mod tidy 
    @golangci-lint run

# install dependencies - 安装依赖工具
dependencies:  dep-swag dep-golangci-lint dep-gofumpt

# a tool to help you write API docs - 一个帮助你编写 API 文档的工具
dep-swag:
    @go install github.com/swaggo/swag/cmd/swag@latest

# a linter for Go - 一个 Go 语言的代码检查工具
dep-golangci-lint:
    @go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# a stricter gofmt - 一个更严格的 gofmt
dep-gofumpt:
    @go install mvdan.cc/gofumpt@latest

#===================================== targets end ===========================================#

#=================================== variables start =========================================#

# project name - 项目名称
project_name := "cola"

# project root directory - 项目根目录
root := justfile_directory()

# binary path - go build 输出的二进制文件路径
bin := join(root, project_name)

# main.go path - main.go 文件路径
main_file := join(root,"main.go")



#=================================== variables end =========================================#