package iteration

func repeat(str string, num int) string {
  repeated := ""
  for i:=0; i < num; i++ {
    repeated += str
  }
  return repeated
}
