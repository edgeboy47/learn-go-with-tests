package array

func Sum(numbers [] int) int {
  sum := 0
  for _, number := range(numbers) {
    sum += number
  }
  return sum
}

func SumAll(arrs ...[]int) [] int {
  result := [] int {}
  
  for _, nums := range arrs {
    result = append(result, Sum(nums))
  }

  return result
}
