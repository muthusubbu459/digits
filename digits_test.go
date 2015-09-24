package main
import(
"testing"
)
func value()bool{
return true
}
func testdigits(t *testing.T) {
expected :=value()
actual:=digits("ab555")
if actual==expected{
t.Errorf("Test Failed")
}else {
t.Log("first test passed ")
}
//return value
}
