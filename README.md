# gin-layout-cli
- https://github.com/LIOU2021/gin-layout

# introduction
Install the gin-layout generator with the command go install github.com/LIOU2021/gin-layout-cli@latest. Go will automatically install it in your $GOPATH/bin directory which should be in your $PATH.

or

```
git clone https://github.com/LIOU2021/gin-layout-cli.git
cd gin-layout-cli
make build
```

# Usage
- gin-layout-cli create -s
- gin-layout-cli create env -n database
- gin-layout-cli create helper -n printMyTest
- gin-layout-cli create controller -n userController
# ref
- https://github.com/spf13/cobra/blob/main/user_guide.md

# todo list
- create middleware
- serve start/stop