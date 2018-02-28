package cmd

import (
	"fmt"

	"../generator"

	"github.com/spf13/cobra"
)

func upsert(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("At least one argument needed")
		return
	}
	c, err := getClient(cmd)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if args[0] == "all" {
		args, err = c.GetAllModels()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	for _, arg := range args {
		var content map[string]map[string]interface{}
		err = c.DoRequest("fields_get", arg, []interface{}{}, map[string][]string{"attributes": []string{"type"}}, &content)
		if err != nil {
			if err.Error() != "error: \"\" code: 2" {
				fmt.Println(err.Error())
				return
			}
		}
		if len(content) == 0 {
			fmt.Println(fmt.Sprintf("WARN: The model %s was not found", arg))
			continue
		}
		err = generator.GenerateTypes(arg, generateContent(content))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = generator.GenerateAPI(arg)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}

func generateContent(apiContent map[string]map[string]interface{}) map[string]string {
	content := make(map[string]string, len(apiContent))
	for modelName, field := range apiContent {
		for fieldType, fieldContent := range field {
			if fieldType == "type" {
				content[modelName] = fieldContent.(string)
			}
		}
	}
	return content
}
