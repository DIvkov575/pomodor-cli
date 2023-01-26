package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
  time.Sleep(time.Second)
  fmt.Fprint(os.Stdout, "a")
}
