package main

import(
  "fmt"
  "strings"
  "unsafe"
)

var countries = []string{
    "Argentina", "Bolivia", "Chile", "Colombia", "Ecuador", "French Guiana", "Guyana", "Paraguay", "Peru", "Suriname", "Uruguay", "Venezuela",
  } 

func main() { 
  fmt.Println("Linear search for Peru:")
  index, found := LinearSearch("Peru")
  if found {
    fmt.Println(fmt.Sprintf("Index: %d", index))
  } else {
    fmt.Println("Not found")
  }

  fmt.Println("\nBinary search for Peru:")
  index, found = BinarySearch("Peru")
  if found {
    fmt.Println(fmt.Sprintf("Index: %d", index))
  } else {
    fmt.Println("Not found")
  }
}

func LinearSearch(target string) (int, bool) {
  for index, element := range countries {
    if element == target {
      return index, true
    }
  }
  return -1, false
}

func BinarySearch(target string) (int, bool) {
  ptrSize := unsafe.Sizeof(string(""))
  origin := unsafe.Pointer(&countries[0])
  lower := unsafe.Pointer(&countries[0])
  upper := unsafe.Pointer(uintptr(lower) + ptrSize * uintptr((len(countries) - 1)))
  mid := unsafe.Pointer(uintptr(lower) + ((uintptr(upper) - uintptr(lower)) / ptrSize / 2) * ptrSize)

  for uintptr(lower) < uintptr(upper) {
    index := (uintptr(mid) - uintptr(origin)) / ptrSize
    switch(strings.Compare(GetPointerStringValue(mid), target)) {
      case 0: {
        return int(index), true
      }
      case -1: {
        lower = unsafe.Pointer(uintptr(mid) - ptrSize)
        mid = unsafe.Pointer(uintptr(lower) + ((uintptr(upper) - uintptr(lower)) / ptrSize / 2) * ptrSize)
      }
      case 1: {
        upper = unsafe.Pointer(uintptr(mid) + ptrSize)
        mid = unsafe.Pointer(uintptr(lower) + ((uintptr(upper) - uintptr(lower)) / ptrSize / 2) * ptrSize)
      }
    }
  }
  
  return -1, false
}

func GetPointerStringValue(ptr unsafe.Pointer) string {
  return *(*string)(ptr)
}
