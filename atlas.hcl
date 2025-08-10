data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/atlas-loader",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "sqlite://dev.db"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}