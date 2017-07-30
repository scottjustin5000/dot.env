package dotenv 
import (
"os"
"testing"
)


func TestLoad(t *testing.T) {
  expectedValues := map[string]string{
    "FOO": "BAR",
    "SOME": "THING",
    "ANOTHER": "THING1",
    "ONEMORE": "THING3",
  }
  var file = "test-envs/.env-test"
  er := Load(file)
  if er != nil {
    t.Fatalf("Failed to load %v", file)
  }

  for k := range expectedValues {
    envValue := os.Getenv(k)
    v := expectedValues[k]
    if envValue != v {
      t.Fatalf("Mismatch for key=%q, expected %q got %q", k, v, envValue)
    }
  }

}