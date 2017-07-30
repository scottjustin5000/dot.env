package dotenv

import (
  "bufio"
  "log"
  "os"
  "regexp"
  "strings"
)

func isValid(line string) bool {
  cleaned := strings.Trim(line, " \n\t")
  if len(cleaned) == 0 {
    return false
  }
  if strings.HasPrefix(cleaned, "#") {
    return false
  }
  if strings.HasPrefix(cleaned, "/*") {
    return false
  }
  matched, err :=regexp.MatchString(`^\s*([\w\.\-]+)\s*=\s*(.*)?\s*$`, cleaned)
  if err != nil {
    return false
  }
  return matched
}

func processLine(line string) (string, string) {

   data := strings.SplitN(line, "=", 2)
   var key = strings.Trim(data[0], " ")
   var val = strings.Trim(data[1], " ")
   //clean text trim everything after text before /*
   i := strings.Index(val, "#")
   if i > -1 {
    val = val[:i]
   }
   i = strings.Index(val, "/*")
   if i > -1 {
    val = val[:i]
   }
   return key, strings.Trim(val, " ")
}

func readEnv(fileName string) (map[string]string, error) {
  file, err := os.Open(fileName)
  if err != nil {
    return nil, err
  }

  defer file.Close()
  var data []string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    var line = scanner.Text()
    if isValid(line) {
      data = append(data, line)
    }
  }
  if err = scanner.Err(); err != nil {
    log.Fatal(err)
    return nil, err
  }

  var enVars = make(map[string]string)
  for _, line := range data {
    var key, value string
    key, value = processLine(line)

    enVars[key] = value   
  }
  return enVars, nil
}


func Load(fileName string) error {
  var enVars, err = readEnv(fileName)
  if err != nil {
    log.Fatal(err)
    return err
  }
  for key, value := range enVars {
    os.Setenv(key, value)
  }
  return nil
}