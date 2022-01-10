




set ORIGINAL_GOPATH_VALUE=%GOPATH%
echo "Original GoPath:" %ORIGINAL_GOPATH_VALUE%
set GOPATH=E:\Go\src\github.com\go-dex-dev\platform\example
echo "Current GoPath for Project:" %GOPATH%

go run main.go

::# go fmt example\build\
::# go fmt example\build\domain\
::# go fmt project\build\database\
::# go fmt project\build\domain\entities\

go build -o %GOPATH%\dist\
set GOPATH=%ORIGINAL_GOPATH_VALUE%
echo "Restored GoPath:" %GOPATH%
