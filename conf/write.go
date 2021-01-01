package conf

import (
	"os"
)

func WriteConfDB(cate string, news []string) error {
	f, err := os.Create("conf/" + cate + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range news {
		f.WriteString(line + "\n")
	}
	return nil
}
