package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	helpFlag := flag.Bool("help", false, "显示帮助信息")
	srcFlag := flag.Bool("src", false, "打印源码")
	flag.Parse()

	if *helpFlag {
		fmt.Println("如果直接运行，默认使用80端口")
		fmt.Println("软件名+空格+端口，运行指定端口")
		fmt.Println("如运行8080端口 run.exe 8080")
		fmt.Println("")
		fmt.Println("如果在后面添加多个参数，只会运行第一个端口")
		fmt.Println("--src  使用这个参数时，将会把源码全部打印出来")
		fmt.Println("By：开心ucu-2023-0720")
		return
	}

	if *srcFlag {
		printSourceCode()
		return
	}

	ports := flag.Args()
	if len(ports) == 0 {
		startServer("80")
		return
	}

	for _, port := range ports {
		startServer(port)
	}
}

func startServer(port string) {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))

	go func(p string) {
		fmt.Printf("服务器已启动，访问地址：http://localhost:%s\n", p)
		err := http.ListenAndServe(":"+p, mux)
		if err != nil {
			fmt.Println("启动服务器失败：", err)
		}
	}(port)

	select {}
}

func printSourceCode() {
	sourceCode := `
========================================
package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	helpFlag := flag.Bool("help", false, "显示帮助信息")
	srcFlag := flag.Bool("src", false, "打印源码")
	flag.Parse()

	if *helpFlag {
		fmt.Println("如果直接运行，默认使用80端口")
		fmt.Println("软件名+空格+端口，运行指定端口")
		fmt.Println("如运行8080端口 run.exe 8080")
		fmt.Println("")
		fmt.Println("如果在后面添加多个参数，只会运行第一个端口")
		fmt.Println("--src  使用这个参数时，将会把源码全部打印出来")
		fmt.Println("By：开心ucu-2023-0720")
		return
	}

	if *srcFlag {
		printSourceCode()
		return
	}

	ports := flag.Args()
	if len(ports) == 0 {
		startServer("80")
		return
	}

	for _, port := range ports {
		startServer(port)
	}
}

func startServer(port string) {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))

	go func(p string) {
		fmt.Printf("服务器已启动，访问地址：http://localhost:%s\n", p)
		err := http.ListenAndServe(":"+p, mux)
		if err != nil {
			fmt.Println("启动服务器失败：", err)
		}
	}(port)

	select {}
}

func printSourceCode() {
	sourceCode := 你看到这段内容
	fmt.Println(sourceCode)
}

========================================
将内容保存到main.go后，运行 go build main.go
即可编码为exe（windows系统）
`
	fmt.Println(sourceCode)
}
