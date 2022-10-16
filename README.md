# gin-layout-cli
- https://github.com/LIOU2021/gin-layout

# introduction
Install the gin-layout generator with the command go install github.com/LIOU2021/gin-layout-cli@latest. Go will automatically install it in your $GOPATH/bin directory which should be in your $PATH.

# Usage
- gin-layout-cli create -s
- gin-layout-cli create env -n database
- gin-layout-cli create helper -n printMyTest
- gin-layout-cli create controller -n userController
# ref
- https://github.com/spf13/cobra/blob/main/user_guide.md

# todo list
- create controller
- create middleware
- serve start/stop
- test create other go docker, and go install, can success work cli?