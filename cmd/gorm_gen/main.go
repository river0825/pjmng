package main

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func snakeToCamel(s string) string {
	caser := cases.Title(language.English)
	parts := strings.Split(s, "_")
	for i := 0; i < len(parts); i++ {
		parts[i] = caser.String(parts[i])
	}
	return strings.Join(parts, "")
}

//go:generate go run ./main.go

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../core/dependency/storage/gorm/models",
		// if you want the nullable field generation property to be pointer type, set FieldNullable true
		FieldNullable: true,
		// if you want to assign field which has a default value in the `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		FieldCoverable: true,
		// if you want to generate field with unsigned integer type, set FieldSignable true
		FieldSignable: true,
		// if you want to generate index tags from database, set FieldWithIndexTag true
		FieldWithIndexTag: true,
		//// if you want to generate type tags from database, set FieldWithTypeTag true
		FieldWithTypeTag: true,
		ModelPkgPath:     "table",
		Mode: gen.WithoutContext |
			gen.WithDefaultQuery |
			gen.WithQueryInterface,
	})

	g.WithOpts(gen.FieldModify(func(f gen.Field) gen.Field {
		// if name consist '-' or '__', replace it to '_'
		switch f.ColumnName {
		case "teacher-recommendation_id":
			f.Name = "TeacherRecommendationID"
		case "teacher_recommendation_id":
			f.Name = "TeacherRecommendationID2"
		default:
			f.Name = strings.ReplaceAll(strings.ReplaceAll(f.ColumnName, "-", "_"), "__", "_")

		}

		if len(f.GORMTag["default"]) > 0 && f.Name == "id" {
			delete(f.GORMTag, "default")
		}

		// if type is jsonb, change it to datatypes.JSON, f.GORMTag["type"] is an array. check one by one
		if len(f.GORMTag["type"]) > 0 && f.GORMTag["type"][0] == "jsonb" {
			if len(f.GORMTag["type"]) > 1 && f.GORMTag["type"][1] == "not null" {
				f.Type = "datatypes.JSON"
			} else {
				f.Type = "*datatypes.JSON"
			}
		}
		return f
	}))

	g.WithFileNameStrategy(func(tableName string) (fileName string) {
		// replace '-' to '_' and replace '__' to '_'
		fileName = strings.ReplaceAll(strings.ReplaceAll(tableName, "-", "_"), "__", "_")
		return fileName
	})

	g.WithModelNameStrategy(func(tableName string) (modelName string) {
		// snake case to camel case and replace '-' to ''
		modelName = snakeToCamel(strings.ReplaceAll(strings.ReplaceAll(tableName, "-", "_"), "__", "_"))
		// remove tail 's'
		modelName = strings.TrimSuffix(modelName, "s")
		return modelName
	})

	gormdb, _ := gorm.Open(postgres.Open("host=localhost port=5430 user=postgres password=postgres dbname=postgres search_path=dev sslmode=disable"))

	g.UseDB(gormdb) // reuse your gorm db

	//// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.ApplyBasic(
	//	// Generate struct `User` based on table `users`
	//	g.GenerateModel("users"),
	//
	//	// Generate struct `Employee` based on table `users`
	//	g.GenerateModelAs("users", "Employee"),
	//
	//	// Generate struct `User` based on table `users` and generating options
	//	g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64")),
	//
	//	// Generate struct `Customer` based on table `customer` and generating options
	//	// customer table may have a tags column, it can be JSON type, gorm/gen tool can generate for your JSON data type
	//	g.GenerateModel("customer", gen.FieldType("tags", "datatypes.JSON")),
	//)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)

	// Generate the code
	g.Execute()
}
